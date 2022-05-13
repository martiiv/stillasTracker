//
//  TransfereScaffolding.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 05/04/2022.
//

import SwiftUI
import Foundation

/// **TrandsfereScaffolding**
/// Background view of TransfereScaffolding
struct TransfereScaffolding: View {
    /// All projects used for transfering
    var projects: [Project]
    
    /// Lightmode or darkmode?
    @Environment(\.colorScheme) var colorScheme

    /// Scaffolding type
    var scaffolding: Scaffolding
    
    /// Transfere Modal View is showing
    @Binding var isShowingSheet: Bool
    
    var body: some View {
        VStack {
            /// Transfere scaffolding view
            TransfereScaffoldingView(isShowingSheet: $isShowingSheet, projects: projects, scaffolding: scaffolding)
                .navigationTitle(Text("Overfør \(scaffolding.type)"))
            }
    }
    
    /// Modal view dismissed
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
