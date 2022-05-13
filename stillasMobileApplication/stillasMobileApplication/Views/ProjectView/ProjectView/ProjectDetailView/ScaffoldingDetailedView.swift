//
//  ScaffoldingDetailedView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 09/05/2022.
//

import SwiftUI

/// **ScaffoldingDetailedView**
/// The detailed view of each scaffolding type
/// Snowes history of the type as well as a button which redirects to transfering of scaffolding
struct ScaffoldingDetailedView: View {
    
    /// All projects used in transfere scaffolding
    var projects: [Project]
    
    /// The scaffolding type
    var scaffolding: Scaffolding
    
    /// Transfere scaffolding Modal View is showing
    @Binding var isShowingSheet: Bool
    
    var body: some View {
        VStack {
            /// History of scaffolding unit type for project and transfere scaffolding button
            HistoryOfScaffolding(projects: projects, scaffolding: scaffolding, isShowingSheet: $isShowingSheet)
            
            TransfereScaffoldingButton(projects: projects, scaffolding: scaffolding, isShowingSheet: $isShowingSheet)
        }
    }
}

/*
struct ScaffoldingDetailedView_Previews: PreviewProvider {
    static var previews: some View {
        ScaffoldingDetailedView()
    }
}*/
