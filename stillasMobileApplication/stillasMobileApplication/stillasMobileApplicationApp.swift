//
//  stillasMobileApplicationApp.swift
//  stillasMobileApplication
//
//  Created by Aleksander Aaboen on 08/03/2022.
//

import SwiftUI
import Firebase
/*
@main
struct stillasMobileApplicationApp: App {
    var body: some Scene {
        WindowGroup {
            ContentView()
        }
    }
}*/

class AppDelegate: NSObject, UIApplicationDelegate {
    func application(_ application: UIApplication, didFinishLaunchingWithOptions launchOptions: [UIApplication.LaunchOptionsKey : Any]? = nil) -> Bool {
        FirebaseApp.configure()
        return true
    }
}

@main
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
