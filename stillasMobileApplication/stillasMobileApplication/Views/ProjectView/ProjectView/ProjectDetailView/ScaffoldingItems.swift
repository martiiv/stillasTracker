//
//  ScaffoldingItems.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 01/05/2022.
//

import SwiftUI

/// **ScaffoldingItems**
/// A grid containing all the scaffolding items as buttons
struct ScaffoldingItems: View {
    /// Projects needed for transfering scaffolding
    var projects: [Project]
    
    /// A two-column grid layout which is flexible
    var gridItemLayout = [GridItem(.flexible()), GridItem(.flexible())]
    
    /// All scaffolding items
    var scaffolding: [Scaffolding]
    
    /// Modal View of transfere scaffolding
    @Binding var isShowingSheet: Bool

    var body: some View {
        ScrollView (.vertical) {
            LazyVGrid (columns: gridItemLayout, spacing: 10) {
                /// For each scaffolding type on the project, add scaffolding button
                ForEach(scaffolding, id: \.type) { scaffolding in
                    NavigationLink(destination: ScaffoldingDetailedView(projects: projects, scaffolding: scaffolding, isShowingSheet: $isShowingSheet),
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
