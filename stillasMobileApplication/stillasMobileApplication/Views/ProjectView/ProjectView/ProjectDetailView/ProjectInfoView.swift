//
//  ProjectDetailView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 04/04/2022.
//

import SwiftUI
import MapKit

struct ProjectInfoView: View {
    @Environment(\.colorScheme) var colorScheme
    @State private var isShowingSheet = false
        
    var projects: [Project]
    var project: Project
    
    let sizeSelections = ["Stillas", "Prosjekt Info"]
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
                        ForEach(sizeSelections, id: \.self) {
                            Text($0)
                        }
                    }
                    .pickerStyle(SegmentedPickerStyle())
                    .padding(.bottom, 15)
                    
                    Spacer()
                    
                    switch selection {
                    case "Prosjekt Info":
                        VStack {
                            //ProjectInfoDetailedView(project: project)
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
