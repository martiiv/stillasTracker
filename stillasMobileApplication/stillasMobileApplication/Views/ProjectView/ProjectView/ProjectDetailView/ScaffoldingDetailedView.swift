//
//  ScaffoldingDetailedView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 09/05/2022.
//

import SwiftUI

struct ScaffoldingDetailedView: View {
    var projects: [Project]
    var scaffolding: Scaffolding
    @Binding var isShowingSheet: Bool
    
    var body: some View {
        VStack {
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
