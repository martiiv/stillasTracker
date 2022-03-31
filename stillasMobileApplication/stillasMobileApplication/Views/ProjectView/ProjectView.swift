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

struct ProjectView: View {
    let items: [BookMark] = [.example1, .example2, .example3]
    
    @State var searchQuery = ""

    let projectsArr = [
        Project(
            id: 420,
            projectName: "Ntnu i gjøvik",
            latitude: 60.7905060889568,
            longitude: 10.681777071532371,
            period:
                Period.init(startDate: "20.02.2020", endDate: "10.05.2020"),
            size: 240,
            state: "Active",
            adresse:
                Adresse.init(gate: "Piazza del Colosseo 1", postnummer: "0184", kommune: "Gjøvik kommune", fylke: "Innlandet"),
            leier:
                Leier(name: "NTNU", number: 639967700),
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
            id: 321,
            projectName: "Ntnu i gjøvik",
            latitude: 60.7905060889568,
            longitude: 10.681777071532371,
            period:
                Period.init(startDate: "20.02.2020", endDate: "10.05.2020"),
            size: 240,
            state: "Active",
            adresse:
                Adresse.init(gate: "Piazza del Colosseo 1", postnummer: "0184", kommune: "Gjøvik kommune", fylke: "Innlandet"),
            leier:
                Leier(name: "NTNU", number: 639967700),
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
    ]

    var body: some View {
        NavigationView {
            Form {
                Section(header: Text("All Projects")) {
                    List(searchResults) { project in
                        NavigationLink(destination: DetailView(project: project), label: {
                            ProjectRow(project: project) }
                        )
                    }
                    .navigationTitle("Projects")
                    //.listStyle(.grouped)
                }

                Section(header: Text("Second List")) {
                    List(items, children: \.items) { row in
                        Image(systemName: row.icon)
                        Text(row.name)
                    }
                }
                
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
            return projectsArr.filter { $0.projectName.contains(searchQuery) }.sorted { $0.projectName < $1.projectName }
        }
    }
}

struct ProjectRow: View {
    var project: Project
    
    var body: some View {
        HStack {
            VStack(alignment: .leading) {
                Text(project.projectName)
                Text(project.period.startDate + " until " + project.period.endDate).font(.subheadline).foregroundColor(.gray)
            }
            Spacer()
            Text(String(format: "%d", project.id))
                .foregroundColor(.gray)
        }
    }
}

struct DetailView: View {
    var project: Project

    var body: some View {
        VStack {
            Text(project.projectName).font(.title)
            
            HStack {
                Text("\(project.projectName) - \(String(format: "%d", project.state))")
            }
            
            Spacer()
        }
    }
}

struct ProjectView_Previews: PreviewProvider {
    static var previews: some View {
            ProjectView()
    }
}


extension BookMark {
    static let apple = BookMark(name: "Apple", icon: "1.circle")
    static let bbc = BookMark(name: "BBC", icon: "square.and.pencil")
    static let swift = BookMark(name: "Swfit", icon: "bolt.fill")
    static let twitter = BookMark(name: "Twitter", icon: "mic")

    static let example1 = BookMark(name: "Favorites", icon: "star", items: [BookMark.apple, BookMark.bbc, BookMark.swift, BookMark.twitter])
    static let example2 = BookMark(name: "Recent", icon: "timer", items: [BookMark.apple, BookMark.bbc, BookMark.swift, BookMark.twitter])
    static let example3 = BookMark(name: "Recommended", icon: "hand.thumbsup", items: [BookMark.apple, BookMark.bbc, BookMark.swift, BookMark.twitter])
}
