//
//  AppViewModel.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 05/05/2022.
//

import FirebaseAuth

/// **AppViewModel**
/// The prosessing class responsible for login and sign up through Firebase Authentication
/// Code taken from and inspired from the following sources:
/// https://www.youtube.com/watch?v=vPCEIPL0U_k
/// https://firebase.google.com/docs/auth/ios/start
class AppViewModel: ObservableObject {
    
    let auth = Auth.auth()
    
    /// Is user signed in?
    @Published var signedIn: Bool = false
    
    /// Getter and Setter for the userID
    var userID: String {
        get { return auth.currentUser?.uid ?? "" }
        set { self.userID = newValue }
    }
    
    /// Getter for is user signed in?
    var isSignedIn: Bool {
        return auth.currentUser != nil
    }
    
    
    /// Attempts to sign in user to the application by authorizing the user through the Firebase Authentication
    /// - Parameters:
    ///   - email: The email of the user
    ///   - password: The password of the user
    func signIn(email: String, password: String) {
        
        /// Attempts to authorize the credentials in Firebase Authentication
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
    
    
    /// Attempts to register a new user to the application by adding them to the Firebase Authentication
    /// - Parameters:
    ///   - email: The email of the new user
    ///   - password: The password of the new user
    func signUp(email: String, password: String) {
        /// Attempts to create a new user
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
    
    
    /// Signs the user out of the application
    func signOut() {
        try? auth.signOut()
        
        self.signedIn = false
    }
}
