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
    @Binding var selection: String
    
    let sizeSelections = ["Less Than", "Between", "Greater Than"]

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
                case "Less Than":
                    SizeLessThanFilter(sizeFilterActive: $sizeFilterActive, scoreFrom: scoreFrom, scoreFromBind: $scoreFromBind)
                        .onChange(of: scoreFrom) { val in
                            scoreFromBind = val
                            sizeFilterActive = true
                        }
                case "Between":
                    SizeBetweenFilter(sizeFilterActive: $sizeFilterActive, scoreFrom: scoreFrom, scoreFromBind: $scoreFromBind, scoreTo: scoreTo, scoreToBind: $scoreToBind)
                        .onChange(of: scoreTo) { val in
                            scoreToBind = val
                            sizeFilterActive = true
                        }
                        .onChange(of: scoreFrom) { val in
                            scoreFromBind = val
                            sizeFilterActive = true
                        }
                case "Greater Than":
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
