//
//  ProjectView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/03/2022.
//

import SwiftUI

struct BookMark: Identifiable {
    let id = UUID()
    let name: String
    let icon: String
    var items: [BookMark]?
}
/*
struct ProjectView: View {
    //let items: [BookMark] = [.example1, .example2, .example3]
    
    @State var searchQuery = ""

    let projectsArr = [
        Project(projectID: 145, projectName: "Project name", size: 130, state: "Active", latitude: 60.7905060889568, longitude: 10.681777071532371, period: Period.init(startDate: "20.03.2021", endDate: "23.12.2022"), address: Address(street: "Nice gate", zipcode: "2821", municipality: "Gjøvik kommune", county: "Innlandet"), customer: Customer(name: "Kai Rune Pettersen", number: 94250982, email: "KaiRune@hotmail.com"), geofence: Geofence(wPosition: Position(latitude: 60.79077759591496, longitude: 10.683249543160402), xPosition: Position(latitude: 60.79077759591496, longitude: 10.683249543160402), yPosition: Position(latitude: 60.79077759591496, longitude: 10.683249543160402), zPosition: Position(latitude: 60.79077759591496, longitude: 10.683249543160402)), scaffolding: [Scaffolding(type: "Flooring", quantity: Quantity(expected: 290, registered: 175))]),
        
        /*
        Project(
            projectID: 420,
            projectName: "Ntnu i gjøvik",
            latitude: 60.7905060889568,
            longitude: 10.681777071532371,
            period:
                Period.init(startDate: "20.02.2020", endDate: "10.05.2020"),
            size: 240,
            state: "Active",
            address:
                Address.init(street: "Piazza del Colosseo 1", zipcode: "0184", municipality: "Gjøvik kommune", county: "Innlandet"),
            customer:
                Customer(name: "NTNU", number: 639967700, email: "NTNU@ntnu.no"),
            scaffolding:
                Scaffolding(units: [
                    Unit(type: "Spire",
                        quantity:
                            Quantity(expected: 3241, registered: 3241)),
                    Unit(type: "Flooring",
                        quantity:
                            Quantity(expected: 500000, registered: 499211))]),
            geofence:
                Geofence(
                    wPosition:
                            Position(latitude: 60.79077759591496, longitude: 10.683249543160402),
                    xPosition:
                            Position(latitude: 60.79077759591496, longitude: 10.683249543160402),
                    yPosition:
                            Position(latitude: 60.79077759591496, longitude: 10.683249543160402),
                    zPosition:
                            Position(latitude: 60.79077759591496, longitude: 10.683249543160402))),
        
        Project(
            projectID: 321,
            projectName: "CC Gjøvik",
            latitude: 60.799530,
            longitude: 10.693144,
            period:
                Period.init(startDate: "20.02.2020", endDate: "10.05.2020"),
            size: 240,
            state: "Active",
            address:
                Address.init(street: "Piazza del Colosseo 1", zipcode: "0184", municipality: "Gjøvik kommune", county: "Innlandet"),
            customer:
                Customer(name: "NTNU", number: 639967700, email: "NTNU@ntnu.no"),
            scaffolding:
                Scaffolding(units: [
                    Unit(type: "Spire",
                        quantity:
                            Quantity(expected: 3241, registered: 3241)),
                    Unit(type: "Flooring",
                         quantity:
                            Quantity(expected: 500000, registered: 499211))]),
            geofence:
                Geofence(
                    wPosition:
                        Position(latitude: 60.799530, longitude: 10.693144),
                    xPosition:
                        Position(latitude: 60.799530, longitude: 10.693144),
                    yPosition:
                        Position(latitude: 60.799530, longitude: 10.693144),
                    zPosition:
                        Position(latitude: 60.799530, longitude: 10.693144))),
        Project(
            projectID: 510,
            projectName: "Studenten Gjøvik",
            latitude: 60.798036,
            longitude: 10.681777071532371,
            period:
                Period.init(startDate: "20.02.2020", endDate: "10.05.2020"),
            size: 240,
            state: "Active",
            address:
                Address.init(street: "Piazza del Colosseo 1", zipcode: "0184", municipality: "Gjøvik kommune", county: "Innlandet"),
            customer:
                Customer(name: "NTNU", number: 639967700, email: "NTNU@ntnu.no"),
            scaffolding:
                Scaffolding(units: [
                    Unit(type: "Spire",
                        quantity:
                            Quantity(expected: 3241, registered: 3241)),
                    Unit(type: "Flooring",
                         quantity:
                            Quantity(expected: 500000, registered: 499211))]),
            geofence:
                Geofence(
                    wPosition:
                        Position(latitude: 60.798036, longitude: 10.685283),
                    xPosition:
                        Position(latitude: 60.798036, longitude: 10.685283),
                    yPosition:
                        Position(latitude: 60.798036, longitude: 10.685283),
                    zPosition:
                        Position(latitude: 60.798036, longitude: 10.685283))),
        Project(
            projectID: 124,
            projectName: "Sit Barnehage",
            latitude: 60.787788,
            longitude: 10.680136,
            period:
                Period.init(startDate: "20.02.2020", endDate: "10.05.2020"),
            size: 240,
            state: "Active",
            address:
                Address.init(street: "Piazza del Colosseo 1", zipcode: "0184", municipality: "Gjøvik kommune", county: "Innlandet"),
            customer:
                Customer(name: "NTNU", number: 639967700, email: "NTNU@ntnu.no"),
            scaffolding:
                Scaffolding(units: [
                    Unit(type: "Spire",
                        quantity:
                            Quantity(expected: 3241, registered: 3241)),
                    Unit(type: "Flooring",
                         quantity:
                            Quantity(expected: 500000, registered: 499211))]),
            geofence:
                Geofence(
                    wPosition:
                        Position(latitude: 60.787788, longitude: 10.680136),
                    xPosition:
                        Position(latitude: 60.787788, longitude: 10.680136),
                    yPosition:
                        Position(latitude: 60.787788, longitude: 10.680136),
                    zPosition:
                        Position(latitude: 60.787788, longitude: 10.680136))),*/
    ]
    
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

                /*
                Section(header: Text("Second List")) {
                    List(items, children: \.items) { row in
                        Image(systemName: row.icon)
                        Text(row.name)
                    }
                }
                */
                
            }
            .listStyle(.grouped)
        }
        .searchable(text: $searchQuery)
        .navigationViewStyle(.stack)
    }
    var searchResults: [Project] {
        if searchQuery.isEmpty {
            return projectsArr.sorted { $0.projectName < $1.projectName }
        } else {
            return projectsArr.filter { $0.projectName.lowercased().contains(searchQuery.lowercased()) }.sorted { $0.projectName < $1.projectName }
        }
    }
}*/

struct ProjectRow: View {
    var project: Project
    
    var body: some View {
        HStack {
            VStack(alignment: .leading) {
                Text(project.projectName).font(.headline)
                Text(project.period.startDate + " until " + project.period.endDate).font(.subheadline).foregroundColor(.gray)
            }
            Spacer()
            Text(String(format: "%d", project.projectID))
                .foregroundColor(.gray)
        }
    }
}
/*
extension BookMark {
    static let apple = BookMark(name: "Apple", icon: "1.circle")
    static let bbc = BookMark(name: "BBC", icon: "square.and.pencil")
    static let swift = BookMark(name: "Swfit", icon: "bolt.fill")
    static let twitter = BookMark(name: "Twitter", icon: "mic")

    static let example1 = BookMark(name: "Favorites", icon: "star", items: [BookMark.apple, BookMark.bbc, BookMark.swift, BookMark.twitter])
    static let example2 = BookMark(name: "Recent", icon: "timer", items: [BookMark.apple, BookMark.bbc, BookMark.swift, BookMark.twitter])
    static let example3 = BookMark(name: "Recommended", icon: "hand.thumbsup", items: [BookMark.apple, BookMark.bbc, BookMark.swift, BookMark.twitter])
}
 */
struct ProjectView: View {
    @State var results = [Project]()
    @State var searchQuery = ""
    @State var hasFetchedData = false

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
            if (hasFetchedData) {
                print("Already gotten data.")
            } else {
                await loadData()
                hasFetchedData = true
            }
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
        ProjectView()
    }
}
