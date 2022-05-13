//
//  ProjectView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/03/2022.
//

import SwiftUI
import UIKit

/// Enum for the different filtertypes
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

/// **ProjectRow**
/// The project lists information preview
struct ProjectRow: View {
    /// The relevant project
    var project: Project
    
    var body: some View {
        HStack {
            VStack(alignment: .leading) {
                Text(project.projectName).font(.headline).bold().italic()
                Text(project.period.startDate + "  -  " + project.period.endDate).font(.subheadline).foregroundColor(.gray)
            }
            Spacer()
            Text(String(format: "%d", project.projectID))
                .foregroundColor(.gray)
        }
    }
}

/// **ProjectListView**
/// The project list
struct ProjectListView: View {
    /// Initializes an instance of ProjectData to call the loadData method
    @ObservedObject var projectData: ProjectData = ProjectData()
    
    /// Data of searchQuery
    @State var searchQuery = ""
    
    @State var hasFetchedData = false
    
    /// Initializes new Project instance to get project details
    @State var projects = [Project]()
    
    /// Filter- or add project Modal Views active
    @State private var showFilterModalView: Bool = false
    @State private var showAddProjectModalView: Bool = false
    
    /// Initializes sort on size filter page
    @State var sizeSortType: String = "Mellom"
    
    /// Initializes filter to be no filter
    @State var filter: FilterType = .none
    @State var filterArr: [String] = []
    @State var filterArrArea: [String] = []

    // TODO: REMOVE?
    /// Initialized values passed into filter
    @State var projectStartDate = Date.distantPast
    @State var projectEndDate = Date.distantFuture
    @State var projectSize = 99999
    @State var minProjectSize = 100
    @State var maxProjectSize = 1000
    @State var projectState = "Active"
    @State var projectCounty = "Innlandet"
    
    /// Data is loading
    @State private var isLoading = true

    var body: some View {
        VStack {
            NavigationView {
                ZStack {
                    /// Data is loading --> show ProgressView
                    if projectData.isLoading {
                        Spacer().frame(height:100)
                        ProgressView("Laster inn...")
                            .progressViewStyle(CircularProgressViewStyle(tint: .gray))
                            .scaleEffect(x: 1.2, y: 1.2, anchor: .center)
                    } else {
                        /// Displays all projects
                        Form {
                            Section(header: Text("Alle Prosjekter")) {
                                /// List with all projects, filtered projects or searched for projects
                                List(searchResults, id: \.projectID) { project in
                                    NavigationLink(destination: ProjectInfoView(projects: projects, project: project), label: {
                                        ProjectRow(project: project) }
                                    )
                                }
                                .navigationTitle("Prosjekter")
                            }
                        }
                        .listStyle(.grouped)
                        .toolbar {
                            /// Button for adding filter
                            ToolbarItemGroup(placement: .navigationBarLeading) {
                                Button(action: {
                                    print("Filter tapped!")
                                    self.showFilterModalView.toggle()
                                }) {
                                    Label("Filter", systemImage: "line.3.horizontal.decrease.circle")
                                }
                            }
                            
                            /// Button for adding project
                            ToolbarItemGroup(placement: .navigationBarTrailing) {
                                Button(action: {
                                    print("Add project tapped!")
                                    self.showAddProjectModalView.toggle()
                                }) {
                                    Label("Legg til", systemImage: "plus.circle")
                                }
                            }
                        }
                    }
                }
            }
            .task {
                /// On first time opening --> Load in project data
                await projectData.loadData { (projects) in
                     self.projects = projects
                }
            }
            .sheet(isPresented: $showFilterModalView,
                   onDismiss: didDismiss) {
                /// Modal View for filtering projects
                FilterView(selStartDateBind: $projectStartDate, selEndDateBind: $projectEndDate, projectArea: $projectCounty, projectSize: $projectSize, projectStatus: $projectState, minProjectSize: $minProjectSize, maxProjectSize: $maxProjectSize, sizeSortType: $sizeSortType, filterArr: $filterArr, filterArrArea: $filterArrArea)
                    .onChange(of: filterArr) { filterVal in
                        if filterArr.contains("period") {
                            filter = .period
                        }
                        if filterArr.contains("area") {
                            filter = .county
                        }
                        if filterArr.contains("size") && sizeSortType == "Mellom" {
                            filter = .sizeBetween
                        } else if filterArr.contains("size") && sizeSortType == "Mindre enn" {
                            filter = .sizeLessThan
                        } else if filterArr.contains("size") && sizeSortType == "Større enn" {
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
                /// Projectdata is fetches again from the API
                await projectData.loadData { (projects) in
                     self.projects = projects
                }
                print("Refreshed")
            }
        }
        .sheet(isPresented: $showAddProjectModalView, onDismiss: didDismiss){
            /// Modal View for adding project
            /// Not implemented yet
            Text("501")
                .font(.system(size: 40).bold())
            Text("Not Implemented")
                .font(.system(size: 30).bold())
        }
    }
    
    /// Returns all the projects sorted and with the filter if applied
    var searchResults: [Project] {
        if searchQuery.isEmpty {
            return filteredProjects.sorted { $0.projectName < $1.projectName }
        } else {
            return filteredProjects.filter { $0.projectName.lowercased().contains(searchQuery.lowercased()) }.sorted { $0.projectName < $1.projectName }
        }
    }
    
    
    /// Modal View dismissed
    func didDismiss() {
        // Handle the dismissing action.
    }
    
    /// The projects filtered based on the passed in filter
    var filteredProjects: [Project] {
        let dateFormatter = DateFormatter()
        dateFormatter.dateFormat = "dd/MM/yyyy"
        
        switch filter {
        case .none:
            return projects
        case .period:
            return projects.filter { dateFormatter.date(from: $0.period.startDate)! >= projectStartDate && dateFormatter.date(from: $0.period.endDate)! <= projectEndDate }
        case .startBeforePeriod:
            return projects.filter { dateFormatter.date(from: $0.period.startDate)! <= projectStartDate }
        case .startAfterPeriod:
            return projects.filter { dateFormatter.date(from: $0.period.startDate)! >= projectStartDate }
        case .endBeforePeriod:
            return projects.filter { dateFormatter.date(from: $0.period.endDate)! <= projectEndDate }
        case .endAfterPeriod:
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

/*
struct Project_Previews: PreviewProvider {
    static var previews: some View {
        ProjectListView()
    }
}
*/
