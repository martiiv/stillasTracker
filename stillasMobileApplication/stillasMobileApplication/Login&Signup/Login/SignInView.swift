//
//  SignInView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 05/05/2022.
//

import SwiftUI
import MapKit

/// **SignInView**
/// The view responsible for the sign in for users
/// Code are inspired from:
/// https://medium.com/@success.anil.kk/login-screen-demo-with-swiftui-2f711e0c657d
/// https://www.youtube.com/watch?v=vPCEIPL0U_k
/// https://firebase.google.com/docs/auth/ios/start
struct SignInView: View {
    @State var email = ""
    @State var password = ""
    
    /// The model responsible for sign in
    @EnvironmentObject var viewModel: AppViewModel
    
    let verticalPaddingForForm = 40.0
    
    var body: some View {
        NavigationView{
            ZStack {
                /// Darkened background of the login View with the map in the shadows
                ZStack {
                    /// Blur effect
                    Color.black.frame(width: UIScreen.screenWidth, height: UIScreen.screenHeight, alignment: .center).zIndex(1).opacity(0.5).ignoresSafeArea()
                    MapView()
                        .blur(radius: 10)
                        .allowsHitTesting(false)
                }
                
                VStack(spacing: CGFloat(verticalPaddingForForm)) {
                    Text("Welcome To MBStillas ST")
                        .font(.title)
                        .bold()
                    HStack {
                        Image(systemName: "person")
                            .foregroundColor(.secondary)
                        
                        /// The login email
                        TextField("Enter your email", text: $email)
                            .foregroundColor(Color.black)
                            .autocapitalization(.none)
                            .disableAutocorrection(true)
                    }
                    .padding()
                    .background(Color.white)
                    .cornerRadius(10)
                    
                    HStack {
                        Image(systemName: "lock")
                            .foregroundColor(.secondary)
                        
                        /// The login password masked
                        SecureField("Enter password", text: $password)
                            .foregroundColor(Color.black)
                            .autocapitalization(.none)
                            .disableAutocorrection(true)
                    }
                    .padding()
                    .background(Color.white)
                    .cornerRadius(10)
                    
                    /// Sends the login request given that both email and password are filled out
                    Button(action: {
                        guard !email.isEmpty, !password.isEmpty else {
                            return
                        }
                        viewModel.signIn(email: email, password: password)
                    }) {
                        Text("Sign in")
                            .frame(width: 150, height: 50, alignment: .center)
                    }
                    .foregroundColor(.white)
                    .background(Color.blue)
                    .cornerRadius(10)
                    
                    /// Button to redirect to Sign up
                    NavigationLink("Sign up", destination: SignUpView())
                        .foregroundColor(.white)
                        .padding()
                        .background(Color.black.opacity(0.2))
                        .cornerRadius(10)
                    
                }
                .padding(.horizontal, CGFloat(verticalPaddingForForm))
            }
        }
    }
}

struct SignInView_Previews: PreviewProvider {
    static var previews: some View {
        SignInView()
    }
}
