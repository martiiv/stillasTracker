//
//  SignUpView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 05/05/2022.
//

import SwiftUI
import FirebaseAuth

// TODO: Add body where user inputs data in addition to username and password
struct SignUpView: View {
    @State var email = ""
    @State var password = ""
    
    @EnvironmentObject var viewModel: AppViewModel
    
    var body: some View {
        VStack {
            TextField("Email", text: $email)
                .autocapitalization(.none)
                .disableAutocorrection(true)
            SecureField("Password", text: $password)
                .autocapitalization(.none)
                .disableAutocorrection(true)
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
        
        //NavigationBarBottom()
    }
}

struct SignUpView_Previews: PreviewProvider {
    static var previews: some View {
        SignUpView()
    }
}
