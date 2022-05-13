//
//  FilterProjectSize.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 19/04/2022.
//

import SwiftUI

/// **FilterProjectSize**
/// Responsible for switching between the different filtering Views
struct FilterProjectSize: View {
    /// Slider values
    @State var scoreFrom: Int = 100
    @State var scoreTo: Int = 1000
    
    /// Slider selection to be returned to parent View
    @Binding var scoreFromBind: Int
    @Binding var scoreToBind: Int
    
    /// Size filter active
    @Binding var sizeFilterActive: Bool
    
    /// The selected size filtering method (Mindre enn, Mellom, Større enn)
    @Binding var selection: String

    /// Filtrering metoder
    let sizeSelections = ["Mindre enn", "Mellom", "Større enn"]

    var body: some View {
        VStack {
            VStack {
                /// Picker for velging av filtrering metode
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
                    /// Redirects to the SizeLessThanFilter View
                    SizeLessThanFilter(sizeFilterActive: $sizeFilterActive, scoreFrom: scoreFrom, scoreFromBind: $scoreFromBind)
                        .onChange(of: scoreFrom) { val in
                            scoreFromBind = val
                            sizeFilterActive = true
                        }
                case "Mellom":
                    /// Redirects to the SizeBetweenFilter View
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
                    /// Redirects to the SizeGreaterThanFilter View
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

/*
struct FilterProjectSize_Previews: PreviewProvider {
    static var previews: some View {
        FilterProjectSize()
    }
}*/
