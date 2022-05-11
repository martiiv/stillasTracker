//
//  AppViewModel.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 05/05/2022.
//

import FirebaseAuth

/// https://www.youtube.com/watch?v=vPCEIPL0U_k
/// https://firebase.google.com/docs/auth/ios/start
class AppViewModel: ObservableObject {
    
    let auth = Auth.auth()
    
    @Published var signedIn: Bool = false
    
    var userID: String {
        get { return auth.currentUser?.uid ?? "" }
        set { self.userID = newValue }
    }
    
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
