//
//  ProjectData.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 07/04/2022.
//

import SwiftUI
import Foundation

/**
 ProjectData responsible for loading/decoding the project information from the API
    To update views when data changes, you make your data model classes observable objects.
 */
class ProjectData: ObservableObject {
    @Published var projects = [Project]()
    
    func loadData(completion:@escaping ([Project]) -> ()) async {
        guard let url = URL(string: "http://10.212.138.205:8080/stillastracking/v1/api/project") else {
            print("Invalid url...")
            return
        }
        URLSession.shared.dataTask(with: url) { data, response, error in
            let projects = try! JSONDecoder().decode([Project].self, from: data!)
            DispatchQueue.main.async {
                completion(projects)
            }
        }.resume()
    }
}
