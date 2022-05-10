//
//  TransfereScaffolding.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 05/04/2022.
//

import SwiftUI

struct TransfereScaffolding: View {
    @State private var isShowingSheet = false
    
    var body: some View {
        VStack {
            Text("Transfere scaffolding")
                .font(.title)
                .padding(50)
            Text("""
                    Add transfere scaffolding functionality here.
                """)
                .padding(50)
            Button("Dismiss",
                   action: { isShowingSheet.toggle() })
        }
    }

    func didDismiss() {
        // Handle the dismissing action.
    }
}

struct TransfereScaffolding_Previews: PreviewProvider {
    static var previews: some View {
        TransfereScaffolding()
    }
}
