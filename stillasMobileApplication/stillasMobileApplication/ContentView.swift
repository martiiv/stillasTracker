//
//  ContentView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Müller on 24/03/2022.
//

import SwiftUI

/// https://www.youtube.com/watch?v=vPCEIPL0U_k
/// https://firebase.google.com/docs/auth/ios/start
/**
 ContentView is responsible for the views in the application.
 This will need enum and TabView on a later stage to switch between views.
 */
struct ContentView: View {
    @State var email = ""
    @State var password = ""
    
    @EnvironmentObject var viewModel: AppViewModel
    
    var body: some View {
        ZStack {
            if viewModel.signedIn {
                NavigationBarBottom()
            } else {
                SignInView()
            }
        }
        .onAppear {
            viewModel.signedIn = viewModel.isSignedIn
        }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
