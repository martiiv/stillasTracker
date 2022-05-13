//
//  ScaffoldingItem.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 01/05/2022.
//

import SwiftUI

/// **ScaffoldingItem**
/// The preview used for the scaffolding buttons with scaffolding preview data
struct ScaffoldingItem: View {
    /// Checks if the device has light- or dark mode activated
    @Environment(\.colorScheme) var colorScheme
    
    /// The scaffolding unit
    var scaffolding: Scaffolding
    
    var body: some View {
        VStack {
            /// Preview data for the button in use
            Text("\(scaffolding.type)".capitalizingFirstLetter()).font(.title2)
                .lineLimit(1)
            
            Image("\(scaffolding.type)".capitalizingFirstLetter()).resizable().scaledToFit()
            
            Spacer()

            HStack {
                VStack {
                    Text(String(format: "%d", scaffolding.quantity.expected)).foregroundColor(.black)
                        .font(.system(size: 15))
                    Text("FORVENTET")
                        .foregroundColor(.gray)
                        .font(.system(size: 10))
                }
                .frame(alignment: .center)

                VStack {
                    amountOfScaffoldingRegistered(expected: scaffolding.quantity.expected, registered: scaffolding.quantity.registered)
                    Text("REGISTRERT")
                        .foregroundColor(.gray)
                        .font(.system(size: 10))
                }
                .frame(alignment: .center)
            }
        }
        .padding(.vertical, 5)
        .frame(width: (UIScreen.screenWidth / 2) - 40, height: (UIScreen.screenWidth / 2) - 40, alignment: .center)
        .contentShape(RoundedRectangle(cornerRadius: 5))
        .background(colorScheme == .dark ? Color(UIColor.white) : Color(UIColor.white)).cornerRadius(7)
        .shadow(color: Color(UIColor.black).opacity(0.1), radius: 5, x: 0, y: 2)
        .shadow(color: Color(UIColor.black).opacity(0.2), radius: 20, x: 0, y: 10)
    }
}


/// A method for checking the registered scaffolding up against the expected scaffolding and assigning color to the result text depending on how close they are to be equal to eachother
/// - Parameters:
///   - expected: The expected amount of the scaffolding type for the project
///   - registered: The registered amount of the scaffolding type for the project at a given time
/// - Returns: The text of registered units with either red, yellow, green or purple color depending on count
func amountOfScaffoldingRegistered(expected: Int, registered: Int) -> Text {
    if (registered >= Int(Double(expected) * 0.95) && registered <= Int(Double(expected))) {
        return Text(String(format: "%d", registered)).foregroundColor(Color.green)
            .font(.system(size: 15))
    } else if ((registered < Int(Double(expected) * 0.95)) && (registered >= Int(Double(expected) * 0.8))) {
        return Text(String(format: "%d", registered)).foregroundColor(Color.yellow)
            .font(.system(size: 15))
    } else if (registered > Int(Double(expected))) {
        return Text(String(format: "%d", registered)).foregroundColor(Color.purple)
            .font(.system(size: 15))
    } else {
        return Text(String(format: "%d", registered)).foregroundColor(Color.red)
            .font(.system(size: 15))
    }
}

/// An extension to the String struct which adds a function for capitalizing the first letter in a word.
/// This was taken in use from the following URL:
/// https://www.hackingwithswift.com/example-code/strings/how-to-capitalize-the-first-letter-of-a-string
extension String {
    
    /// Capitalizes the first letter of a word
    /// - Returns: The word with the capitalized first letter
    func capitalizingFirstLetter() -> String {
        return prefix(1).capitalized + dropFirst()
    }

    mutating func capitalizeFirstLetter() {
        self = self.capitalizingFirstLetter()
    }
}

/*
struct ScaffoldingItem_Previews: PreviewProvider {
    static var previews: some View {
        ScaffoldingItem()
    }
}
*/
