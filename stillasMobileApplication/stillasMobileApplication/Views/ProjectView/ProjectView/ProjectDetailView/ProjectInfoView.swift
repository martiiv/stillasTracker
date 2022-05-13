//
//  ProjectDetailView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 04/04/2022.
//

import SwiftUI
import MapKit

/// **ProjectInfoView**
/// The View responsible for displaying the project info and scaffolding info Views
struct ProjectInfoView: View {
    /// Darkmode or light mode?
    @Environment(\.colorScheme) var colorScheme
    
    /// Transfere scaffolding Modal View showing?
    @State private var isShowingSheet = false
    
    /// All projects
    var projects: [Project]
    
    /// Specific project
    var project: Project
    
    /// The two views available
    let siteSelections = ["Stillas", "Prosjekt Info"]
    
    /// Initialize selection
    @State var selection: String = "Prosjekt Info"

    var body: some View {
        ScrollView {
            VStack {
                /// MapView displaying the map in the top of the screen
                MapView()
                    .frame(height: 300)
                
                Text("\(project.projectName)")
                    .font(.title).bold()
                    .foregroundColor(colorScheme == .dark ? Color(UIColor.darkGray) : Color(UIColor.darkGray))
                
                VStack {
                    Picker("Select a state: ", selection: $selection) {
                        ForEach(siteSelections, id: \.self) {
                            Text($0)
                        }
                    }
                    .pickerStyle(SegmentedPickerStyle())
                    .padding(.bottom, 15)
                    
                    Spacer()
                    
                    switch selection {
                    case "Prosjekt Info":
                        VStack {
                            ProjectInfoDetailedView(project: project)
                        }
                    case "Stillas":
                        VStack {
                            ScaffoldingView(projects: projects, scaffolding: project.scaffolding!)
                        }
                    default:
                        Text("Found none")
                    }
                }
             }
        }
        .ignoresSafeArea(edges: .top)
    }
}
