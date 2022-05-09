//
//  TransfereScaffolding.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 05/04/2022.
//

import SwiftUI
import Foundation

struct TransfereScaffolding: View {
    var projects: [Project]
    @Environment(\.colorScheme) var colorScheme

    var scaffolding: Scaffolding
    @Binding var isShowingSheet: Bool
    
    @State private var quantity: Int = 1
    @State private var name: String = "Tim"
    @State private var projectFrom: String = ""
    @State private var projectTo: String = ""

    
    var body: some View {
        VStack {
            TransfereScaffoldingView(isShowingSheet: $isShowingSheet, projects: projects, scaffolding: scaffolding)
                .navigationTitle(Text("Transfere \(scaffolding.type)"))
            }
    }
    
    func didDismiss() {
        // Handle the dismissing action.
    }
}

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

/*
struct TransfereScaffolding_Previews: PreviewProvider {
    static var previews: some View {
        TransfereScaffolding()
    }
}*/
