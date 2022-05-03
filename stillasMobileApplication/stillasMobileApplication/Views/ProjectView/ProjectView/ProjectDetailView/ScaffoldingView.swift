//
//  ScaffoldingView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 29/04/2022.
//

import SwiftUI

struct ScaffoldingView: View {
    var projects: [Project]
    var scaffolding: [Scaffolding]
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
