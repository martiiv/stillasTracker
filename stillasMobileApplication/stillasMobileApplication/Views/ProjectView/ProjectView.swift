//
//  ProjectView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/03/2022.
//

import SwiftUI

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

struct ProjectView: View {
    @State var searchQuery = ""

    let projectsArr = [
        Project(id: 420, projectName: "Ntnu i gjøvik", latitude: 60.7905060889568, longitude: 10.681777071532371, period: Period.init(startDate: "20.02.2020", endDate: "10.05.2020"), size: 240, state: "Active", adresse: Adresse.init(gate: "Piazza del Colosseo 1", postnummer: "0184", kommune: "Gjøvik kommune", fylke: "Innlandet"), leier: Leier(name: "NTNU", number: 639967700), scaffolding: Scaffolding(units: [Unit(type: "Spire", quantity: Quantity(expected: 3241, registered: 3241)), Unit(type: "Flooring", quantity: Quantity(expected: 500000, registered: 499211))]), geofence: Geofence(wPosition: Position(latitude: 60.79077759591496, longitude: 10.683249543160402), xPosition: Position(latitude: 60.79077759591496, longitude: 10.683249543160402), yPosition: Position(latitude: 60.79077759591496, longitude: 10.683249543160402), zPosition: Position(latitude: 60.79077759591496, longitude: 10.683249543160402))),
        Project(id: 321, projectName: "Ntnu i gjøvik", latitude: 60.7905060889568, longitude: 10.681777071532371, period: Period.init(startDate: "20.02.2020", endDate: "10.05.2020"), size: 240, state: "Active", adresse: Adresse.init(gate: "Piazza del Colosseo 1", postnummer: "0184", kommune: "Gjøvik kommune", fylke: "Innlandet"), leier: Leier(name: "NTNU", number: 639967700), scaffolding: Scaffolding(units: [Unit(type: "Spire", quantity: Quantity(expected: 3241, registered: 3241)), Unit(type: "Flooring", quantity: Quantity(expected: 500000, registered: 499211))]), geofence: Geofence(wPosition: Position(latitude: 60.79077759591496, longitude: 10.683249543160402), xPosition: Position(latitude: 60.79077759591496, longitude: 10.683249543160402), yPosition: Position(latitude: 60.79077759591496, longitude: 10.683249543160402), zPosition: Position(latitude: 60.79077759591496, longitude: 10.683249543160402))),
    ]

    var body: some View {
        ZStack {
            NavigationView {
                List (searchResults) { project in
                NavigationLink(destination: DetailView(project: project)) { ProjectRow (project: project) }
                }
                .listStyle(PlainListStyle())
                .searchable(text: $searchQuery)
                .navigationTitle("Projects")
            }
            .navigationViewStyle(.stack)
        }
    }

    var searchResults: [Project] {
        if searchQuery.isEmpty {
            return projectsArr.sorted { $0.projectName < $1.projectName }
        } else {
            return projectsArr.filter { $0.projectName.contains(searchQuery) }.sorted { $0.projectName < $1.projectName }
        }
    }
}

struct ProjectView_Previews: PreviewProvider {
    static var previews: some View {
        ProjectView()
    }
}
