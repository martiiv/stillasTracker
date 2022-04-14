//
//  FilterProjectData.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 14/04/2022.
//

import SwiftUI

struct FilterProjectData: View {
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

    @State var projects = [Project]()
    
    let filter: FilterType
    
    // TODO: Make these values operable
    @State var projectStartDate = ""
    @State var projectEndDate = ""
    @State var projectSize = ""
    @State var projectState = ""
    @State var projectCounty = ""
    
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
            }
            VStack {
                Text("Nice")
                Text("\(projects.count)")
            }
        }
        .task {
            await ProjectData().loadData { (projects) in
                 self.projects = projects
            }
        }
    }
    
    var filteredProjects: [Project] {
        switch filter {
        case .none:
            return projects
        case .period:
            return projects.filter { $0.period.startDate > projectStartDate && $0.period.endDate < projectEndDate }
        case .startBeforePeriod:
            return projects.filter { $0.period.startDate < projectStartDate }
        case .startAfterPeriod:
            return projects.filter { $0.period.startDate > projectStartDate }
        case .endBeforePeriod:
            return projects.filter { $0.period.endDate < projectEndDate }
        case .endAfterPeriod:
            return projects.filter { $0.period.endDate > projectEndDate }
        case .sizeEqualTo:
            return projects.filter { $0.size == Int(projectSize) }
        case .sizeLessThan:
            return projects.filter { $0.size < Int(projectSize) ?? 0 }
        case .sizeGreaterThan:
            return projects.filter { $0.size > Int(projectSize) ?? 0 }
        case .state:
            return projects.filter { $0.state == projectState }
        case .county:
            return projects.filter { $0.address.county == projectCounty }
        }
    }
}

struct FilterProjectData_Previews: PreviewProvider {
    static var previews: some View {
        FilterProjectData(filter: .none)
    }
}
