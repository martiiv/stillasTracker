//
//  FilterProjectData.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 14/04/2022.
//

import UIKit
import SwiftUI
/*
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
}*/

struct FilterProjectData: View {
    @State var projects = [Project]()
    @State private var showFilterModalView: Bool = false
    @State private var showAddProjectModalView: Bool = false
    @State var filterArrArea: [String] = []

    @State var filter: FilterType = .none
    @State var filterArr: [String] = []

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
                    FilterView(selStartDateBind: $projectStartDate, selEndDateBind: $projectEndDate, projectArea: $projectCounty, projectSize: $projectSize, projectStatus: $projectState, filterArr: $filterArr, filterArrArea: $filterArrArea)
                        .onChange(of: projectStartDate) { value in
                            filter = .period
                        }
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

struct AddProjectView: View {
    var body: some View {
        VStack {
            Text("Add Project SheetView")
        }
    }
}
/*
struct FilterProjectData_Previews: PreviewProvider {
    static var previews: some View {
        FilterProjectData()
    }
}*/
