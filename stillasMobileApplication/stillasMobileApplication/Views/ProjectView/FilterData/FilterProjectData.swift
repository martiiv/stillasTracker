//
//  FilterProjectData.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 14/04/2022.
//

import UIKit
import SwiftUI

enum FilterType {
    case none,
         period,
         startBeforePeriod,
         startAfterPeriod,
         endBeforePeriod,
         endAfterPeriod,
         sizeEqualTo,
         sizeLessThan,
         sizeGreaterThan,
         state,
         county
}

struct FilterProjectData: View {
    @State var projects = [Project]()
    @State private var showFilterModalView: Bool = false
    @State private var showAddProjectModalView: Bool = false
    
    let filter: FilterType
    
    // TODO: Make these values operable
    @State var projectStartDate = Date.distantPast
    @State var projectEndDate = Date.distantFuture
    @State var projectSize = 99999
    @State var projectState = "Active"
    @State var projectCounty = "Innlandet"
    
    var body: some View {
        VStack {
            NavigationView {
                Form {
                    Section(header: Text("All Projects")) {
                        List(filteredProjects, id: \.projectID) { project in
                            Text(project.projectName)
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
                            self.showAddProjectModalView.toggle()
                        }) {
                            Label("Add", systemImage: "plus.circle")
                        }
                    }
                }
                .sheet(isPresented: $showFilterModalView,
                       onDismiss: didDismiss) {
                    FilterView(selStartDateBind: $projectStartDate, selEndDateBind: $projectEndDate, projectArea: $projectCounty, projectSize: $projectSize, projectStatus: $projectState)
                }
               .sheet(isPresented: $showAddProjectModalView, onDismiss: didDismiss) {
                   AddProjectView()
               }
            }
        }
        .task {
            await ProjectData().loadData { (projects) in
                 self.projects = projects
            }
        }
    }
    
    func didDismiss() {
        
        // Handle the dismissing action.
    }
    
    var filteredProjects: [Project] {
        let dateFormatter = DateFormatter()
        dateFormatter.dateFormat = "dd/MM/yy"
        
        switch filter {
        case .none:
            return projects
        case .period:
            //return projects.filter { $0.period.startDate > projectStartDate && $0.period.endDate < projectEndDate }
            return projects.filter { dateFormatter.date(from: $0.period.startDate)! > projectStartDate && dateFormatter.date(from: $0.period.endDate)! < projectEndDate }
        case .startBeforePeriod:
            //return projects.filter { $0.period.startDate < projectStartDate }
            return projects.filter { dateFormatter.date(from: $0.period.startDate)! < projectStartDate }
        case .startAfterPeriod:
            //return projects.filter { $0.period.startDate > projectStartDate }
            return projects.filter { dateFormatter.date(from: $0.period.startDate)! > projectStartDate }
        case .endBeforePeriod:
            //return projects.filter { $0.period.endDate < projectEndDate }
            return projects.filter { dateFormatter.date(from: $0.period.endDate)! < projectEndDate }
        case .endAfterPeriod:
            //return projects.filter { $0.period.endDate > projectEndDate }
            return projects.filter { dateFormatter.date(from: $0.period.endDate)! > projectEndDate }
        case .sizeEqualTo:
            return projects.filter { $0.size == Int(projectSize) }
        case .sizeLessThan:
            return projects.filter { $0.size < Int(projectSize) }
        case .sizeGreaterThan:
            return projects.filter { $0.size > Int(projectSize) }
        case .state:
            return projects.filter { $0.state == projectState }
        case .county:
            return projects.filter { $0.address.county == projectCounty }
        }
    }
}

struct FilterView: View {
    @State private var filterItems = ["Område", "Prosjekt periode", "Størrelse", "Status"]
    
    @Binding var selStartDateBind: Date
    @Binding var selEndDateBind: Date
    @Binding var projectArea: String
    @Binding var projectSize: Int
    @Binding var projectStatus: String
    
    @State var selStartDate = Date()
    @State var selEndDate = Date()
        
    var body: some View {
        NavigationView {
            List {
                ForEach(filterItems, id: \.self) { filterItem in
                    NavigationLink {
                        switch filterItem {
                        case "Område":
                            FilterProjectArea()
                        case "Prosjekt periode":
                            FilterProjectPeriod(selStartDateBind: $selStartDate, selEndDateBind: $selEndDate)
                                .onChange(of: selStartDate) { value in
                                    selStartDateBind = $selStartDate.wrappedValue
                                    print("______")
                                    print(value)
                                    print("______")
                                }
                                .onChange(of: selEndDate) { value in
                                    selEndDateBind = $selEndDate.wrappedValue
                                    print("______")
                                    print(value)
                                    print("______")
                                }
                        case "Størrelse":
                            FilterProjectSize()
                            /*
                        case "Status":
                            print("Add status view")
                            // ADD status view
                        default:
                            print("Did not find any")
                        */
                        default:
                            AddProjectView()
                        }
                    } label: {
                        Text(filterItem)
                    }
                }
            }
            .navigationTitle(Text("Filter"))
            .navigationViewStyle(StackNavigationViewStyle())
        }
    }
}

struct AddProjectView: View {
    var body: some View {
        VStack {
            Text("Add Project SheetView")
        }
    }
}

struct FilterProjectData_Previews: PreviewProvider {
    static var previews: some View {
        FilterProjectData(filter: .none)
    }
}
