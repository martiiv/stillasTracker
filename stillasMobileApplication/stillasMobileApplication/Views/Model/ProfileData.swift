//
//  ProfileData.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 10/05/2022.
//

import SwiftUI
import Foundation

class ProfileData: ObservableObject {
    @Published private var _isLoadingProfile: Bool = false
    
    var isLoadingProfile: Bool {
        get { return _isLoadingProfile}
    }
    
    func loadData(userID: String, completion:@escaping (Profile) -> ()) async {
        _isLoadingProfile = true
        print("One = \(_isLoadingProfile)")
        
        guard let url = URL(string: "http://10.212.138.205:8080/stillastracking/v1/api/user?id=\(userID)") else {
            print("Invalid url...")
            return
        }
        URLSession.shared.dataTask(with: url) { [self] data, response, error in
            let profile = try! JSONDecoder().decode(Profile.self, from: data!)
            DispatchQueue.main.async {
                completion(profile)
                self._isLoadingProfile = false
                print("Two = \(self._isLoadingProfile)")
            }
        }.resume()
    }
}
