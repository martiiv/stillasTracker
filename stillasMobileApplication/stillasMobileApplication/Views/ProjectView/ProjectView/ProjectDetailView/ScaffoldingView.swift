//
//  ScaffoldingView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 29/04/2022.
//

import SwiftUI

/// **ScaffoldingView**
/// The view responsible for showing scaffolding units
struct ScaffoldingView: View {
    var projects: [Project]
    
    /// All scaffolding units
    var scaffolding: [Scaffolding]
    
    /// Scaffolding item Modal View
    @State var isShowingSheet: Bool = false
    var body: some View {
        VStack {
            ScaffoldingItems(projects: projects, scaffolding: scaffolding, isShowingSheet: $isShowingSheet)
        }
    }
}

/*
struct ScaffoldingView_Previews: PreviewProvider {
    static var previews: some View {
        ScaffoldingView()
    }
}*/
