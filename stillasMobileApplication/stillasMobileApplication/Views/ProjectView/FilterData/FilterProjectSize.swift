//
//  FilterProjectSize.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 19/04/2022.
//

import SwiftUI

struct FilterProjectSize: View {
    var body: some View {
        VStack {
            IntSlider()
        }
    }
}

struct IntSlider: View {
    private enum Field: Int, CaseIterable {
            case input
        }

    @FocusState private var focusedField: Field?

    
    @State var score: Int = 0
    @ObservedObject var input = NumbersOnly()
    
    @State var score2: Int = 0
    @ObservedObject var input2 = NumbersOnly()
    
    private var sliderSizeMin = 100.0
    private var sliderSizeMax = 1000.0
    private var stepLength = 50.0

    var intProxyS1: Binding<Double>{
        Binding<Double>(
            get: {
            //returns the score as a Double
                return Double(score)
        }, set: {
            //rounds the double to an Int
            print($0.description)
            score = Int($0)
            input.value = "\(Int($0))"
        })
    }
    
    var intProxyS2: Binding<Double>{
        Binding<Double>(
            get: {
            //returns the score as a Double
                return Double(score2)
        }, set: {
            //rounds the double to an Int
            print($0.description)
            score2 = Int($0)
            input2.value = "\(Int($0))"
        })
    }
    
    var body: some View {
        ScrollView {
        HStack {
            VStack {
                if (Int(input.value) == Int(sliderSizeMin)) {
                    Text("Under")
                } else {
                    Text("Minimum")
                }
                HStack {
                    TextField("Input", text: $input.value)
                        .font(Font.system(size: 35, design: .default))
                        .onChange(of: input.value) { value in
                            score = Int(value) ?? Int(sliderSizeMax + sliderSizeMin) / 2
                        }
                        //.frame(height: 100)
                        .keyboardType(.numberPad)
                        .focused($focusedField, equals: .input)
                        .frame(alignment: .center)
                        .multilineTextAlignment(.center)
                    
                    HStack {
                        Text("m")
                            .font(Font.system(size: 30, design: .default))
                         Text("2")
                            .baselineOffset(5.0)
                    }
                }
                
                Divider()
                 .frame(height: 1)
                 .padding(.horizontal, 20)
                 .background(Color.gray)
            }
            
            Text(" - ")
                .font(Font.system(size: 35, design: .default))
            
            VStack {
                if (Int(input2.value) == Int(sliderSizeMax)) {
                    Text("Over")
                } else {
                    Text("Maksimum")
                }
                
                TextField("Input", text: $input2.value)
                    .font(Font.system(size: 35, design: .default))
                    .onChange(of: input2.value) { value in
                        score2 = Int(value) ?? Int(sliderSizeMax + sliderSizeMin) / 2
                    }
                    //.frame(height: 100)
                    .keyboardType(.numberPad)
                    .focused($focusedField, equals: .input)
                    .multilineTextAlignment(.center)
                
                Divider()
                 .frame(height: 1)
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
                        Text("Fra")
                    }
                    .foregroundColor(.secondary)
                    .font(.subheadline)
                    .font(Font.system(size: 20, design: .default))
                    .padding(.top, 20)
                    
                    Slider(value: intProxyS1 , in: sliderSizeMin...sliderSizeMax, step: stepLength, onEditingChanged: {_ in
                        print(score.description)
                    })
                    .frame(width: 350, alignment: .center)
                    .padding(.vertical, 20)
                }
                VStack (alignment: .leading) {
                    HStack {
                        Text("Til")
                    }
                    .foregroundColor(.secondary)
                    .font(.subheadline)
                    .font(Font.system(size: 20, design: .default))
                    .padding(.top, 20)
                    
                    Slider(value: intProxyS2 , in: sliderSizeMin...sliderSizeMax, step: stepLength, onEditingChanged: {_ in
                        print(score2.description)
                    })
                    .frame(width: 350, alignment: .center)
                    .padding(.vertical, 20)
                }
            }
        }
        Spacer()
        Button(action: { print("Bruk") }) {
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

class NumbersOnly: ObservableObject {
    @Published var value = "" {
        didSet {
            let filtered = value.filter { $0.isNumber }
            
            if value != filtered {
                value = filtered
            }
        }
    }
}

struct FilterProjectSize_Previews: PreviewProvider {
    static var previews: some View {
        FilterProjectSize()
    }
}
