//
//  ScaffoldingView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 29/04/2022.
//

import SwiftUI

struct ScaffoldingView: View {
    var scaffolding: [Scaffolding]
    @Binding var isShowingSheet: Bool

    var body: some View {
        VStack {
            TransfereScaffoldingButton(isShowingSheet: $isShowingSheet)
            
            ScaffoldingItems(scaffolding: scaffolding)
        }
    }
}

/*
struct ScaffoldingView_Previews: PreviewProvider {
    static var previews: some View {
        ScaffoldingView()
    }
}*/
