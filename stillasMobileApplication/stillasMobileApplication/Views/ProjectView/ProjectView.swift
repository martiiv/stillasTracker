//
//  ProjectView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/03/2022.
//

import SwiftUI

/*
struct BookMark: Identifiable {
    let id = UUID()
    let name: String
    let icon: String
    var items: [BookMark]?
}

struct ProjectView: View {
    //let items: [BookMark] = [.example1, .example2, .example3]
 Section(header: Text("Second List")) {
     List(items, children: \.items) { row in
         Image(systemName: row.icon)
         Text(row.name)
     }
 }
 */

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
    @State var searchQuery = ""
    @State var hasFetchedData = false
    @State var projects = [Project]()
    
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
        .task {
            await ProjectData().loadData { (projects) in
                 self.projects = projects
            }
        }
        .navigationViewStyle(.stack)
        .searchable(text: $searchQuery)
    }
    var searchResults: [Project] {
        if searchQuery.isEmpty {
            return projects.sorted { $0.projectName < $1.projectName }
        } else {
            return projects.filter { $0.projectName.lowercased().contains(searchQuery.lowercased()) }.sorted { $0.projectName < $1.projectName }
        }
    }
}

struct Project_Previews: PreviewProvider {
    static var previews: some View {
        ProjectView()
    }
}
