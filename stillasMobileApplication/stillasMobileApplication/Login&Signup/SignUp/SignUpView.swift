//
//  SignUpView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 05/05/2022.
//

import SwiftUI
import FirebaseAuth

/// **SignUpView**
/// The view responsible for the sign up page
/// Code taken and inspired from this youtube video and firebases authentication startup
/// https://www.youtube.com/watch?v=vPCEIPL0U_k
/// https://firebase.google.com/docs/auth/ios/start
// TODO: Add body where user inputs data in addition to username and password
struct SignUpView: View {
    @State var email = ""
    @State var password = ""
    
    /// The model responsible for sign up
    @EnvironmentObject var viewModel: AppViewModel
    
    var body: some View {
        VStack {
            TextField("Email", text: $email)
                .autocapitalization(.none)
                .disableAutocorrection(true)
            
            /// SecureField to mask the password
            SecureField("Password", text: $password)
                .autocapitalization(.none)
                .disableAutocorrection(true)
            
            /// Sends the request to sign up to the system
            Button(action: {
                guard !email.isEmpty, !password.isEmpty else {
                    return
                }
                viewModel.signUp(email: email, password: password)
            }) {
                Text("Sign up")
            }
        }
        .padding()
    }
}

struct SignUpView_Previews: PreviewProvider {
    static var previews: some View {
        SignUpView()
    }
}
