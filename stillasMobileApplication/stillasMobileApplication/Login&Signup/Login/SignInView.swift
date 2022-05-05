//
//  SignInView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 05/05/2022.
//

import SwiftUI
import MapKit

/// https://medium.com/@success.anil.kk/login-screen-demo-with-swiftui-2f711e0c657d
/// https://www.youtube.com/watch?v=vPCEIPL0U_k
/// https://firebase.google.com/docs/auth/ios/start
struct SignInView: View {
    @State var email = ""
    @State var password = ""
    
    @EnvironmentObject var viewModel: AppViewModel
    
    let verticalPaddingForForm = 40.0
        var body: some View {
            NavigationView{
                ZStack {
                    
                    ZStack {
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
                             SecureField("Enter password", text: $password)
                                               .foregroundColor(Color.black)
                                               .autocapitalization(.none)
                                               .disableAutocorrection(true)
                        }
                        .padding()
                        .background(Color.white)
                        .cornerRadius(10)
                        
                        
                        
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
                        
                        
                        NavigationLink("Sign up", destination: SignUpView())
                            .foregroundColor(.white)
                            .padding()
                            .background(Color.black.opacity(0.2))
                            .cornerRadius(10)
                        
                    }.padding(.horizontal, CGFloat(verticalPaddingForForm))
                    
                }
            }
        }
    /*
    var body: some View {
        NavigationView {
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
        }
        
        //NavigationBarBottom()
    }*/
}

struct SignInView_Previews: PreviewProvider {
    static var previews: some View {
        SignInView()
    }
}
