//
//  SignInView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 05/05/2022.
//

import SwiftUI

struct SignInView: View {
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
                viewModel.signIn(email: email, password: password)
                
            }) {
                Text("Sign in")
            }
            NavigationLink("Sign up", destination: SignUpView())
                .padding()
        }
        .padding()
        
        //NavigationBarBottom()
    }
}

struct SignInView_Previews: PreviewProvider {
    static var previews: some View {
        SignInView()
    }
}
