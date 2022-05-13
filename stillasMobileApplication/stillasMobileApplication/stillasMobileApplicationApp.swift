//
//  stillasMobileApplicationApp.swift
//  stillasMobileApplication
//
//  Created by Aleksander Aaboen on 08/03/2022.
//

import SwiftUI
import Firebase


/// **AppDelegate**
/// Initializes the application with FirebaseApp
class AppDelegate: NSObject, UIApplicationDelegate {
    func application(_ application: UIApplication, didFinishLaunchingWithOptions launchOptions: [UIApplication.LaunchOptionsKey : Any]? = nil) -> Bool {
        FirebaseApp.configure()
        return true
    }
}

@main
/// **stillasMobileApplication**
/// Create an app by declaring a structure that conforms to the App protocol.
/// Assigns the AppDelegate.
struct stillasMobileApplication: App {
    @UIApplicationDelegateAdaptor(AppDelegate.self) var appDelegate

    var body: some Scene {
        WindowGroup {
            let viewModel = AppViewModel()
            ContentView()
                .environmentObject(viewModel)
        }
    }
}
