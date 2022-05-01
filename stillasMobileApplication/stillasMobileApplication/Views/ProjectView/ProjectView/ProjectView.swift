//
//  ProjectView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/03/2022.
//

import SwiftUI
import UIKit
/*
struct BookMark: Identifiable {
    let id = UUID()
    let name: String
    let icon: String
    var items: [BookMark]?
}

struct ProjectView1: View {
    let items: [BookMark] = [.example1, .example2, .example3]
    var body: some View {
    Section(header: Text("Second List")) {
         List(items, children: \.items) { row in
             Image(systemName: row.icon)
             Text(row.name)
         }
     }
    }
}*/
 

enum FilterType {
    case none,
         period,
         startBeforePeriod,
         startAfterPeriod,
         endBeforePeriod,
         endAfterPeriod,
         sizeBetween,
         sizeEqualTo,
         sizeLessThan,
         sizeGreaterThan,
         state,
         county
}

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
    @State private var showFilterModalView: Bool = false
    
    @State var sizeSortType: String = "Between"
    @State var filter: FilterType = .none
    @State var filterArr: [String] = []
    @State var filterArrArea: [String] = []

    // TODO: REMOVE?
    @State var projectStartDate = Date.distantPast
    @State var projectEndDate = Date.distantFuture
    @State var projectSize = 99999
    @State var minProjectSize = 100
    @State var maxProjectSize = 1000
    @State var projectState = "Active"
    @State var projectCounty = "Innlandet"
    
    var body: some View {
        VStack {
        NavigationView {
            Form {
                Section(header: Text("All Projects")) {
                    /*List(searchResults, id: \.projectID) { project in
                        NavigationLink(destination: ProjectDetailView(project: project), label: {
                            ProjectRow(project: project) }
                        )
                    }
                    .navigationTitle("Projects")*/
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
            .toolbar {
                ToolbarItemGroup(placement: .navigationBarLeading) {
                    Button(action: {
                        print("Filter tapped!")
                        self.showFilterModalView.toggle()
                    }) {
                        Label("Filter", systemImage: "line.3.horizontal.decrease.circle")
                    }
                }
                
                ToolbarItemGroup(placement: .navigationBarTrailing) {
                    Button(action: {
                        print("Add project tapped!")
                    }) {
                        Label("Add", systemImage: "plus.circle")
                    }
                }
            }
        }
        .task {
            await ProjectData().loadData { (projects) in
                 self.projects = projects
            }
        }
        .sheet(isPresented: $showFilterModalView,
               onDismiss: didDismiss) {
            FilterView(selStartDateBind: $projectStartDate, selEndDateBind: $projectEndDate, projectArea: $projectCounty, projectSize: $projectSize, projectStatus: $projectState, minProjectSize: $minProjectSize, maxProjectSize: $maxProjectSize, sizeSortType: $sizeSortType, filterArr: $filterArr, filterArrArea: $filterArrArea)
                .onChange(of: filterArr) { filterVal in
                    if filterArr.contains("period") {
                        filter = .period
                    }
                    if filterArr.contains("area") {
                        filter = .county
                    }
                    if filterArr.contains("size") && sizeSortType == "Between" {
                        filter = .sizeBetween
                    } else if filterArr.contains("size") && sizeSortType == "Less Than" {
                        filter = .sizeLessThan
                    } else if filterArr.contains("size") && sizeSortType == "Greater Than" {
                        filter = .sizeGreaterThan
                    }
                    if filterArr.contains("status") {
                        filter = .state
                    }
                    if filterArr.isEmpty {
                        filter = .none
                    }
                }
        }
        .navigationViewStyle(.stack)
        .searchable(text: $searchQuery, placement: .navigationBarDrawer(displayMode: .always))
        .refreshable {
            print("Refreshed")
            // TODO: Add action here (retrieve from API again or something)
        }
            //ProjectView1()
        }
    }
    
    // TODO: Add support for search of filtered items
    var searchResults: [Project] {
        if searchQuery.isEmpty {
            return filteredProjects.sorted { $0.projectName < $1.projectName }
        } else {
            return filteredProjects.filter { $0.projectName.lowercased().contains(searchQuery.lowercased()) }.sorted { $0.projectName < $1.projectName }
        }
    }
    
    func didDismiss() {
        // Handle the dismissing action.
    }
    
    var filteredProjects: [Project] {
        let dateFormatter = DateFormatter()
        dateFormatter.dateFormat = "dd/MM/yyyy"
        
        switch filter {
        case .none:
            return projects
        case .period:
            //return projects.filter { $0.period.startDate > projectStartDate && $0.period.endDate < projectEndDate }
            return projects.filter { dateFormatter.date(from: $0.period.startDate)! >= projectStartDate && dateFormatter.date(from: $0.period.endDate)! <= projectEndDate }
        case .startBeforePeriod:
            //return projects.filter { $0.period.startDate < projectStartDate }
            return projects.filter { dateFormatter.date(from: $0.period.startDate)! <= projectStartDate }
        case .startAfterPeriod:
            //return projects.filter { $0.period.startDate > projectStartDate }
            return projects.filter { dateFormatter.date(from: $0.period.startDate)! >= projectStartDate }
        case .endBeforePeriod:
            //return projects.filter { $0.period.endDate < projectEndDate }
            return projects.filter { dateFormatter.date(from: $0.period.endDate)! <= projectEndDate }
        case .endAfterPeriod:
            //return projects.filter { $0.period.endDate > projectEndDate }
            return projects.filter { dateFormatter.date(from: $0.period.endDate)! >= projectEndDate }
        case .sizeBetween:
            return projects.filter { $0.size >= Int(minProjectSize) && $0.size <= Int(maxProjectSize)}
        case .sizeEqualTo:
            return projects.filter { $0.size == Int(projectSize) }
        case .sizeLessThan:
            return projects.filter { $0.size < Int(minProjectSize) }
        case .sizeGreaterThan:
            return projects.filter { $0.size > Int(maxProjectSize) }
        case .state:
            return projects.filter { $0.state == projectState }
        case .county:
            return projects.filter { filterArrArea.contains($0.address.county)}
        }
    }
}

struct Project_Previews: PreviewProvider {
    static var previews: some View {
        ProjectView()
    }
}
