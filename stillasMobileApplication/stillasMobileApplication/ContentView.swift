//
//  ContentView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Müller on 24/03/2022.
//

import SwiftUI

/**
    ContentView is responsible for the views in the application.
    This will need enum and TabView on a later stage to switch between views.
 */
struct ContentView: View {
    var body: some View {
        MapView()
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
