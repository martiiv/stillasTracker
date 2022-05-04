//
//  ContentView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Müller on 24/03/2022.
//

import SwiftUI
import FirebaseAuth

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
        NavigationView {
            if viewModel.signedIn {
                VStack {
                    Text("You are signed in!")
                    
                    Button (action: {
                        viewModel.signOut()
                    }, label: {
                        Text("Sign out")
                            .frame(width: 150, height: 50)
                            .background(Color.gray)
                            .foregroundColor(Color.blue)
                            .padding()
                    })
                }
            } else {
                SignInView()
            }
            
        }
        .onAppear {
            viewModel.signedIn = viewModel.isSignedIn
        }
    }
    /*func login() {
     Auth.auth().signIn(withEmail: email, password: password) { (result, error) in
     if error != nil {
     print(error?.localizedDescription ?? "")
     } else {
     print("success")
     }
     }
     }*/
}

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
    /*func login() {
     Auth.auth().signIn(withEmail: email, password: password) { (result, error) in
     if error != nil {
     print(error?.localizedDescription ?? "")
     } else {
     print("success")
     }
     }
     }*/
}

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
    /*func login() {
     Auth.auth().signIn(withEmail: email, password: password) { (result, error) in
     if error != nil {
     print(error?.localizedDescription ?? "")
     } else {
     print("success")
     }
     }
     }*/
}

class AppViewModel: ObservableObject {
    
    let auth = Auth.auth()
    
    @Published var signedIn: Bool = false
    
    var isSignedIn: Bool {
        return auth.currentUser != nil
    }
    
    func signIn(email: String, password: String) {
        auth.signIn(withEmail: email, password: password) { [weak self ] (result, error) in
            guard result != nil, error == nil else {
                return
            }
            DispatchQueue.main.async {
                // Success
                self?.signedIn = true
            }
        }
    }
    
    func signUp(email: String, password: String) {
        auth.createUser(withEmail: email, password: password) { [weak self ] (result, error) in
            guard result != nil, error == nil else {
                return
            }
            DispatchQueue.main.async {
                // Success
                self?.signedIn = true
            }
        }
    }
    
    func signOut() {
        try? auth.signOut()
        
        self.signedIn = false
    }
}

/*
 func signUp() {
 Auth.auth().signIn(withEmail: email, password: password) { (result, error) in
 if error != nil {
 print(error?.localizedDescription ?? "")
 } else {
 print("success")
 }
 }
 }
 */

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
