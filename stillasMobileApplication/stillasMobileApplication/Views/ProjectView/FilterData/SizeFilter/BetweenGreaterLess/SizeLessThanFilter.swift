//
//  SizeLessThanFilter.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/04/2022.
//

import SwiftUI

struct SizeLessThanFilter: View {
    enum Field: Int, CaseIterable {
        case input
    }

    @Binding var sizeFilterActive: Bool
    
    @FocusState var focusedField: Field?
    
    @State var scoreFrom: Int = 100
    @Binding var scoreFromBind: Int
    @ObservedObject var input = NumbersOnly()
    
    var sliderSizeMin = 100.0
    var sliderSizeMax = 1000.0
    var stepLength = 50.0

    var intProxyS1: Binding<Double>{
        Binding<Double>(
            get: {
            //returns the score as a Double
                return Double(scoreFrom)
        }, set: {
            //rounds the double to an Int
            print($0.description)
            scoreFrom = Int($0)
            input.value = "\(Int($0))"
        })
    }
    
    var body: some View {
        ScrollView {
            HStack {
                VStack {
                    Text("Under")
                    
                    HStack {
                        TextField("\(Int(sliderSizeMin))", text: $input.value)
                            .font(Font.system(size: 30, design: .default))
                            .onChange(of: input.value) { value in
                                scoreFrom = Int(value) ?? Int(sliderSizeMax + sliderSizeMin) / 2
                            }
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
                VStack (alignment: .leading){
                    HStack {
                        Text("Til")
                    }
                    .foregroundColor(.secondary)
                    .font(.subheadline)
                    .font(Font.system(size: 20, design: .default))
                    .padding(.top, 20)
                    
                    Slider(value: intProxyS1 , in: sliderSizeMin...sliderSizeMax, step: stepLength, onEditingChanged: {_ in
                        print(scoreFrom.description)
                    })
                    .frame(width: 350, alignment: .center)
                    .padding(.vertical, 20)
                }
            }
        }
        Spacer()
        Button(action: {
            print("Bruk")
            scoreFrom = Int(input.value) ?? 100
            scoreFromBind = scoreFrom
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
struct SizeLessThanFilter_Previews: PreviewProvider {
    static var previews: some View {
        SizeLessThanFilter()
    }
}*/
