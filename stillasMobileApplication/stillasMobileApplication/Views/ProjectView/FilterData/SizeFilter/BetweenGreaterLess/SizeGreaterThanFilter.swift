//
//  SizeGreaterThanFilter.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/04/2022.
//

import SwiftUI

/// **SizeGreaterThanFilter**
/// The View for selecting a project with size greater than a values
struct SizeGreaterThanFilter: View {
   
    /// Enum for input
    enum Field: Int, CaseIterable {
        case input
    }
    
    /// Size filter active
    @Binding var sizeFilterActive: Bool
    
    /// Has textfield been activated and should be in focus?
    @FocusState var focusedField: Field?
    
    /// Slider values
    @State var scoreTo: Int = 1000
    @Binding var scoreToBind: Int
    
    /// The input of minimum size
    @ObservedObject var input = NumbersOnly()
    
    /// Slider data
    var sliderSizeMin = 100.0
    var sliderSizeMax = 1000.0
    var stepLength = 50.0
    
    /// Slider value
    var intProxyS2: Binding<Double>{
        Binding<Double>(
            get: {
            //returns the score as a Double
                return Double(scoreTo)
        }, set: {
            //rounds the double to an Int
            print($0.description)
            scoreTo = Int($0)
            input.value = "\(Int($0))"
        })
    }
    
    var body: some View {
        ScrollView {
            HStack {
                VStack {
                    Text("Over")
                    HStack {
                        /// Adds textfield with bind to slider
                        TextField("\(Int(sliderSizeMax))", text: $input.value)
                            .font(Font.system(size: 30, design: .default))
                            .onChange(of: input.value) { value in
                                if (Int(value) ?? Int(sliderSizeMax)) >= Int(sliderSizeMax) {
                                    scoreTo = Int(sliderSizeMax)
                                    
                                    // TODO: Update textfield value to slider value or max value
                                } else if (Int(value) ?? Int(sliderSizeMin)) <= Int(sliderSizeMin) {
                                    scoreTo = Int(sliderSizeMin)
                                    // TODO: Update textfield value to slider value or min value
                                } else {
                                    scoreTo = Int(value) ?? Int(sliderSizeMax + sliderSizeMin) / 2
                                }
                            }
                        //.frame(height: 100)
                            .keyboardType(.numberPad)
                            .focused($focusedField, equals: .input)
                            .frame(width: UIScreen.screenWidth / 3, alignment: .center)
                            .multilineTextAlignment(.center)
                        
                        HStack {
                            Text("m")
                                .font(Font.system(size: 30, design: .default))
                            Text("2")
                                .baselineOffset(6.0)
                        }
                        .padding(.leading, -35)

                    }
                    Divider()
                        .frame(width: UIScreen.screenWidth / 3, height: 1, alignment: .center)
                        .padding(.horizontal, 20)
                        .background(Color.gray)
                }
            }
            .toolbar {
                ToolbarItem(placement: .keyboard) {
                    Button("Done") {
                        focusedField = nil
                    }
                }
            }
            .frame(width: 350, alignment: .center)
            .font(.headline)
            .font(Font.system(size: 60, design: .default))
            
            VStack {
                VStack (alignment: .leading) {
                    HStack {
                        Text("Fra")
                    }
                    .foregroundColor(.secondary)
                    .font(.subheadline)
                    .font(Font.system(size: 20, design: .default))
                    .padding(.top, 20)
                    
                    /// Adds slider for minimum size
                    Slider(value: intProxyS2 , in: sliderSizeMin...sliderSizeMax, step: stepLength, onEditingChanged: {_ in
                        print(scoreTo.description)
                    })
                    .frame(width: 350, alignment: .center)
                    .padding(.vertical, 20)
                }
            }
        }
        Spacer()
        
        /// Returnerer den brukte størrelsedataen til parent Viewen
        Button(action: {
            print("Bruk")
            scoreTo = Int(input.value) ?? 1000
            scoreToBind = scoreTo
            sizeFilterActive = true
        }) {
            Text("Bruk")
                .frame(width: 300, height: 50, alignment: .center)
        }
        .foregroundColor(.white)
        //.padding(.vertical, 10)
        .background(Color.blue)
        .cornerRadius(10)
        
        Spacer()
            .frame(height:50)  // limit spacer size by applying a frame
    }
}

/*
struct SizeGreaterThanFilter_Previews: PreviewProvider {
    static var previews: some View {
        SizeGreaterThanFilter()
    }
}*/
