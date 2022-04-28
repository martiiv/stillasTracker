//
//  FilterProjectSize.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 19/04/2022.
//

import SwiftUI

struct FilterProjectSize: View {
    @State var scoreFrom: Int = 100
    @State var scoreTo: Int = 1000
    
    @Binding var scoreFromBind: Int
    @Binding var scoreToBind: Int
    
    @Binding var sizeFilterActive: Bool
    
    var body: some View {
        VStack {
            IntSlider(sizeFilterActive: $sizeFilterActive, scoreFrom: scoreFrom, scoreFromBind: $scoreFromBind, scoreTo: scoreTo, scoreToBind: $scoreToBind)
                .onChange(of: scoreTo) { val in
                    scoreToBind = val
                    sizeFilterActive = true
                }
                .onChange(of: scoreFrom) { val in
                    scoreFromBind = val
                    sizeFilterActive = true
                }
        }
    }
}

extension UIScreen {
   static let screenWidth = UIScreen.main.bounds.size.width
   static let screenHeight = UIScreen.main.bounds.size.height
   static let screenSize = UIScreen.main.bounds.size
}

struct CornerRadiusStyle: ViewModifier {
    var radius: CGFloat
    var corners: UIRectCorner
    
    struct CornerRadiusShape: Shape {

        var radius = CGFloat.infinity
        var corners = UIRectCorner.allCorners

        func path(in rect: CGRect) -> Path {
            let path = UIBezierPath(roundedRect: rect, byRoundingCorners: corners, cornerRadii: CGSize(width: radius, height: radius))
            return Path(path.cgPath)
        }
    }

    func body(content: Content) -> some View {
        content
            .clipShape(CornerRadiusShape(radius: radius, corners: corners))
    }
}

extension View {
    func cornerRadius(_ radius: CGFloat, corners: UIRectCorner) -> some View {
        ModifiedContent(content: self, modifier: CornerRadiusStyle(radius: radius, corners: corners))
    }
}

struct IntSlider: View {
    enum Field: Int, CaseIterable {
            case input
    }

    @Binding var sizeFilterActive: Bool
    
    @FocusState var focusedField: Field?
    
    @State var scoreFrom: Int = 100
    @Binding var scoreFromBind: Int
    @ObservedObject var input = NumbersOnly()
    
    @State var scoreTo: Int = 1000
    @Binding var scoreToBind: Int

    @ObservedObject var input2 = NumbersOnly()
    
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
    
    var intProxyS2: Binding<Double>{
        Binding<Double>(
            get: {
            //returns the score as a Double
                return Double(scoreTo)
        }, set: {
            //rounds the double to an Int
            print($0.description)
            scoreTo = Int($0)
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
                        TextField("\(Int(sliderSizeMin))", text: $input.value)
                            .font(Font.system(size: 30, design: .default))
                            .onChange(of: input.value) { value in
                                scoreFrom = Int(value) ?? Int(sliderSizeMax + sliderSizeMin) / 2
                            }
                            .keyboardType(.numberPad)
                            .focused($focusedField, equals: .input)
                            .frame(alignment: .center)
                            .multilineTextAlignment(.center)
                        
                        HStack {
                            Text("m")
                                .font(Font.system(size: 30, design: .default))
                            Text("2")
                                .baselineOffset(6.0)
                        }
                    }
                    
                    Divider()
                        .frame(height: 1)
                        .padding(.horizontal, 20)
                        .background(Color.gray)
                }
                
                Text(" - ")
                    .font(Font.system(size: 35, design: .default))
                    .offset(y: 10)
                
                VStack {
                    if (Int(input2.value) == Int(sliderSizeMax)) {
                        Text("Over")
                    } else {
                        Text("Maksimum")
                    }
                    HStack {
                        TextField("\(Int(sliderSizeMax))", text: $input2.value)
                            .font(Font.system(size: 30, design: .default))
                            .onChange(of: input2.value) { value in
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
                            .multilineTextAlignment(.center)
                        
                        HStack {
                            Text("m")
                                .font(Font.system(size: 30, design: .default))
                            Text("2")
                                .baselineOffset(6.0)
                        }
                    }
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
                        print(scoreFrom.description)
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
                        print(scoreTo.description)
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
            scoreTo = Int(input2.value) ?? 1000
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
/*
struct FilterProjectSize_Previews: PreviewProvider {
    static var previews: some View {
        FilterProjectSize()
    }
}*/
