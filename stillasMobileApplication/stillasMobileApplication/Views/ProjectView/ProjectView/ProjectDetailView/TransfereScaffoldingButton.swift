//
//  TransfereScaffoldingButton.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 01/05/2022.
//

import SwiftUI

/// **TransfereScaffoldingButton**
/// Button for activating of transfere scaffolding Modal View
struct TransfereScaffoldingButton: View {
    /// All projects
    var projects: [Project]
    
    /// Darkmode or lightmode activated?
    @Environment(\.colorScheme) var colorScheme
    
    /// Specific scaffolding type
    var scaffolding: Scaffolding
    
    /// Transfere scaffolding Modal View is showing
    @Binding var isShowingSheet: Bool
    
    var body: some View {
        /// Button for opening transfere scaffolding Modal View
        Button {
            isShowingSheet.toggle()
        } label: {
            Text("Overfør Stillas")
                .padding(12)
                .font(.system(size: 20))
                .foregroundColor(colorScheme == .dark ? Color(UIColor.white) : Color(UIColor.white))
                .frame(width: 300, height: 50, alignment: .center)
        }
        .foregroundColor(.white)
        .background(Color.blue)
        .cornerRadius(10)
        .padding(.bottom, 50)
        
        .contentShape(Rectangle())
        //.background(colorScheme == .dark ? Color.blue : Color.blue).cornerRadius(7)
        .shadow(color: Color(UIColor.black).opacity(0.1), radius: 5, x: 0, y: 2)
        .shadow(color: Color(UIColor.black).opacity(0.2), radius: 20, x: 0, y: 10)
        .sheet(isPresented: $isShowingSheet,
               onDismiss: didDismiss) {
            /// Transfere scaffolding Modal View
            TransfereScaffolding(projects: projects, scaffolding: scaffolding, isShowingSheet: $isShowingSheet)
        }
    }
    
    /// Modal view got dismissed
    func didDismiss() {
        // Handle the dismissing action.
    }
}

/*
struct TransfereScaffoldingButton_Previews: PreviewProvider {
    static var previews: some View {
        TransfereScaffoldingButton()
    }
}
*/
