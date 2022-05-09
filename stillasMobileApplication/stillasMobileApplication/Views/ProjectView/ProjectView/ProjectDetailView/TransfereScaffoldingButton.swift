//
//  TransfereScaffoldingButton.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 01/05/2022.
//

import SwiftUI

struct TransfereScaffoldingButton: View {
    var projects: [Project]
    @Environment(\.colorScheme) var colorScheme
    var scaffolding: Scaffolding
    @Binding var isShowingSheet: Bool
    
    var body: some View {
        Button {
            isShowingSheet.toggle()
        } label: {
            Text("Transfere Scaffolding")
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
            TransfereScaffolding(projects: projects, scaffolding: scaffolding, isShowingSheet: $isShowingSheet)
        }
    }
    
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
