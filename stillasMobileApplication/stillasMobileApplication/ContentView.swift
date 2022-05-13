//
//  ContentView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Müller on 24/03/2022.
//

import SwiftUI

/// **ContentView**
/// Responsible for the views in the application.
/// This will need enum and TabView on a later stage to switch between views.
/// https://www.youtube.com/watch?v=vPCEIPL0U_k
/// https://firebase.google.com/docs/auth/ios/start
struct ContentView: View {
    @State var email = ""
    @State var password = ""
    
    /// The model responsible for sign in and sign up
    @EnvironmentObject var viewModel: AppViewModel
    
    var body: some View {
        ZStack {
            /// If user is signed in, give the user access to the application, if not prompt the user with the login and sign up view
            if viewModel.signedIn {
                NavigationBarBottom()
            } else {
                SignInView()
            }
        }
        .onAppear {
            /// Remembers if the user was signed in and closes the application
            viewModel.signedIn = viewModel.isSignedIn
        }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
