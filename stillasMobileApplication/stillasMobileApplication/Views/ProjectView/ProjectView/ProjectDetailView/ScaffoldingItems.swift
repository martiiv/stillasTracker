//
//  ScaffoldingItems.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 01/05/2022.
//

import SwiftUI

struct ScaffoldingItems: View {
    var projects: [Project]
    var gridItemLayout = [GridItem(.flexible()), GridItem(.flexible())]
    var scaffolding: [Scaffolding]
    @Binding var isShowingSheet: Bool

    var body: some View {
        ScrollView (.vertical) {
            LazyVGrid (columns: gridItemLayout, spacing: 10) {
                ForEach(scaffolding, id: \.type) { scaffolding in
                    NavigationLink(destination: HistoryOfScaffolding(projects: projects, scaffolding: scaffolding, isShowingSheet: $isShowingSheet),
                                   label: { ScaffoldingItem(scaffolding: scaffolding)
                    })
                    .listStyle(.grouped)
                }
            }
            .padding(.vertical, 20)
        }
    }
}

/*
struct ScaffoldingItems_Previews: PreviewProvider {
    static var previews: some View {
        ScaffoldingItems()
    }
}*/
