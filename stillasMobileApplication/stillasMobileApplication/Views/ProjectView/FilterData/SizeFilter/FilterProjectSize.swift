//
//  FilterProjectSize.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 19/04/2022.
//

import SwiftUI

/*enum SwipeHorizontalDirection: String {
        case left, right, none
    }*/

struct FilterProjectSize: View {
    @State var scoreFrom: Int = 100
    @State var scoreTo: Int = 1000
    
    @Binding var scoreFromBind: Int
    @Binding var scoreToBind: Int
    
    @Binding var sizeFilterActive: Bool
    @Binding var selection: String
    //@State var swipeHorizontalDirection: SwipeHorizontalDirection = .none { didSet { print(swipeHorizontalDirection) } }

    let sizeSelections = ["Mindre enn", "Mellom", "Større enn"]

    var body: some View {
        
        VStack {
            VStack {
                Picker("Select a state: ", selection: $selection) {
                    ForEach(sizeSelections, id: \.self) {
                        Text($0)
                    }
                }
                .pickerStyle(SegmentedPickerStyle())
                .padding(.bottom, 15)
                
                Spacer()
                
                switch selection {
                case "Mindre enn":
                    SizeLessThanFilter(sizeFilterActive: $sizeFilterActive, scoreFrom: scoreFrom, scoreFromBind: $scoreFromBind)
                        .onChange(of: scoreFrom) { val in
                            scoreFromBind = val
                            sizeFilterActive = true
                        }
                        /*.gesture(DragGesture().onChanged {
                            if $0.startLocation.x == $0.location.x {
                                                    self.swipeHorizontalDirection = .none
                                                    selection = selection
                            } else if $0.startLocation.x > $0.location.x {
                                                    self.swipeHorizontalDirection = .right
                                                    selection = "Between"
                                                }
                        }).transition(.asymmetric(insertion: .scale, removal: .opacity))*/
                case "Mellom":
                    SizeBetweenFilter(sizeFilterActive: $sizeFilterActive, scoreFrom: scoreFrom, scoreFromBind: $scoreFromBind, scoreTo: scoreTo, scoreToBind: $scoreToBind)
                        .onChange(of: scoreTo) { val in
                            scoreToBind = val
                            sizeFilterActive = true
                        }
                        .onChange(of: scoreFrom) { val in
                            scoreFromBind = val
                            sizeFilterActive = true
                        }
                case "Større enn":
                    SizeGreaterThanFilter(sizeFilterActive: $sizeFilterActive, scoreTo: scoreTo, scoreToBind: $scoreToBind)
                        .onChange(of: scoreFrom) { val in
                            scoreFromBind = val
                            sizeFilterActive = true
                        }
                default:
                    Text("Found none")
                }
            }
        }
        .navigationTitle(Text("Størrelse"))
    }
}

extension UIScreen {
   static let screenWidth = UIScreen.main.bounds.size.width
   static let screenHeight = UIScreen.main.bounds.size.height
   static let screenSize = UIScreen.main.bounds.size
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
