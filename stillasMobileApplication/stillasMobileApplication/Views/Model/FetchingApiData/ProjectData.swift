//
//  ProjectData.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 07/04/2022.
//

import SwiftUI
import Foundation

/// **ProjectData**
/// Retrieves all the projects and its data from the API
class ProjectData: ObservableObject {
    //@Published var projects = [Project]() // TODO: REMOVE IF NOT USED IN THE END
    /// Is data loading?
    @Published private var _isLoading: Bool = false

    /// Getter for the _isLoading
    var isLoading: Bool {
        get { return _isLoading}
    }
    
    
    /// Responsible for getting the data about projects from the API
    /// - Parameter completion: completion handler
    func loadData(completion:@escaping ([Project]) -> ()) async {
        _isLoading = true
        print("One = \(_isLoading)")
        
        guard let url = URL(string: "http://10.212.138.205:8080/stillastracking/v1/api/project?scaffolding=true") else {
            print("Invalid url...")
            return
        }
        /// Sends the request and gets the data
        URLSession.shared.dataTask(with: url) { [self] data, response, error in
            let projects = try! JSONDecoder().decode([Project].self, from: data!)
            DispatchQueue.main.async {
                completion(projects)
                self._isLoading = false
                print("Two = \(self._isLoading)")
            }
        }.resume()
    }
}
