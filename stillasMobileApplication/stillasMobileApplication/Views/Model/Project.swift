//
//  Project.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 30/03/2022.
//

import Foundation
import SwiftUI

/*
struct Response: Codable{
    var results: [Project]
}
 */

// MARK: - Project
struct Project:  Codable {
    //var id = UUID()
    let projectID: Int
    let projectName: String
    let size: Int
    let state: String
    let latitude, longitude: Double
    let period: Period
    let address: Address
    let customer: Customer
    let geofence: Geofence
    let scaffolding: [Scaffolding]?
}

// MARK: - Address
struct Address: Codable {
    let street, zipcode, municipality, county: String
}

// MARK: - Customer
struct Customer: Codable {
    let name: String
    let number: Int
    let email: String
}

// MARK: - Geofence
struct Geofence: Codable {
    let wPosition, xPosition, yPosition, zPosition: Position

    enum CodingKeys: String, CodingKey {
        case wPosition = "w-position"
        case xPosition = "x-position"
        case yPosition = "y-position"
        case zPosition = "z-position"
    }
}

// MARK: - Position
struct Position: Codable {
    let latitude, longitude: Double
}

// MARK: - Period
struct Period: Codable {
    let startDate, endDate: String
}

// MARK: - Scaffolding
struct Scaffolding: Codable {
    let type: String
    let quantity: Quantity

    enum CodingKeys: String, CodingKey {
        case type
        case quantity = "Quantity"
    }
}

// MARK: - Quantity
struct Quantity: Codable {
    let expected, registered: Int
}

struct ProjectViewN: View {
    @State private var results = [Project]()
    @State var searchQuery = ""

    var body: some View {
        NavigationView {
            Form {
                Section(header: Text("All Projects")) {
                    List(searchResults, id: \.projectID) { project in
                        NavigationLink(destination: ProjectDetailView(project: project), label: {
                            ProjectRow(project: project) }
                        )
                    }
                    .navigationTitle("Projects")
                    //.listStyle(.grouped)
                }
            }
            .listStyle(.grouped)
        }
        .navigationViewStyle(.stack)
        .searchable(text: $searchQuery)
        .task {
            await loadData()
        }
    }
    var searchResults: [Project] {
        if searchQuery.isEmpty {
            return results.sorted { $0.projectName < $1.projectName }
        } else {
            return results.filter { $0.projectName.lowercased().contains(searchQuery.lowercased()) }.sorted { $0.projectName < $1.projectName }
        }
    }
    
    
    func loadData() async {
        guard let url = URL(string:
                   "http://10.212.138.205:8080/stillastracking/v1/api/project/") else {
           print("Invalid URL")
           return
        }

        do {
           let (data, response) = try await URLSession.shared.data(from: url)
            print(response)
           if let decodedResponse = try? JSONDecoder().decode([Project].self, from: data){
               results = decodedResponse
           }
        }catch {
           print("Invalid URL")
        }
    }
}

struct Project_Previews: PreviewProvider {
    static var previews: some View {
        ProjectViewN()
    }
}
