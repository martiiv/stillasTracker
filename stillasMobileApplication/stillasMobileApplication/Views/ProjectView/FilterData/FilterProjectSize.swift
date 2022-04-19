//
//  FilterProjectSize.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 19/04/2022.
//

import SwiftUI

struct FilterProjectSize: View {
    var body: some View {
        IntSlider()
    }
}

struct IntSlider: View {
    private enum Field: Int, CaseIterable {
            case input
        }
    
    @ObservedObject var input = NumbersOnly()
    
    @State var score: Int = 0
    @FocusState private var focusedField: Field?

    var intProxy: Binding<Double>{
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
    
    var body: some View {
        VStack{
            /*if(input.value != "\(0)" && Int(input.value) != score) {
                input.value = "\(score)"
                AddProjectView()
            } else {
                AddProjectView()
            }*/
            
            VStack {
                TextField("Input", text: $input.value)
                    .onChange(of: input.value) { value in
                        score = Int(value) ?? 0
                    }
                    .padding()
                    .keyboardType(.numberPad)
                    .focused($focusedField, equals: .input)

                Text(score.description)
                Text("Størrelse")
            }
            .toolbar {
                ToolbarItem(placement: .keyboard) {
                    Button("Done") {
                        focusedField = nil
                    }
                }
            }
            .font(.headline)
            .font(Font.system(size: 60, design: .default))
            
            Slider(value: intProxy , in: 100.0...1000.0, step: 50.0, onEditingChanged: {_ in
                print(score.description)
            })
            .frame(width: 350, alignment: .center)
        }
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
