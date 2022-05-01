//
//  ScaffoldingItems.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 01/05/2022.
//

import SwiftUI

struct ScaffoldingItems: View {
    var gridItemLayout = [GridItem(.flexible()), GridItem(.flexible())]
    var scaffolding: [Scaffolding]

    var body: some View {
        ScrollView (.vertical) {
            LazyVGrid (columns: gridItemLayout, spacing: 10) {
                ForEach(scaffolding, id: \.type) { scaffolding in
                    NavigationLink(destination: Text("Scaffolding detail view"), label: {
                        ScaffoldingItem(scaffolding: scaffolding)
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
