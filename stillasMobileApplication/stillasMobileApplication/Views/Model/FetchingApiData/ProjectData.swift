//
//  ProjectData.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 07/04/2022.
//

import SwiftUI
import Foundation

class ProjectData: ObservableObject {
    @Published var projects = [Project]() // TODO: REMOVE IF NOT USED IN THE END
    @Published private var _isLoading: Bool = false

    var isLoading: Bool {
        get { return _isLoading}
    }
    
    func loadData(completion:@escaping ([Project]) -> ()) async {
        _isLoading = true
        print("One = \(_isLoading)")
        
        guard let url = URL(string: "http://10.212.138.205:8080/stillastracking/v1/api/project?scaffolding=true") else {
            print("Invalid url...")
            return
        }
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
