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
                .navigationTitle(Text("Overfør \(scaffolding.type)"))
            }
    }
    
    func didDismiss() {
        // Handle the dismissing action.
    }
}

/*
struct TransfereScaffolding_Previews: PreviewProvider {
    static var previews: some View {
        TransfereScaffolding()
    }
}*/
