//
//  ProjectDetailView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 04/04/2022.
//

import SwiftUI
import MapKit

struct ProjectDetailView: View {
    @Environment(\.colorScheme) var colorScheme
    @State private var isShowingSheet = false
        
    var project: Project

    var body: some View {
        ScrollView {
            VStack {
                /// MapView displaying the map in the top of the screen
                MapView()
                    .frame(height: 300)
                
                Text("\(project.projectName)")
                    .font(.title).bold()
                    .foregroundColor(colorScheme == .dark ? Color(UIColor.darkGray) : Color(UIColor.darkGray))
                
                DetailView(project: project)
                
                Button {
                    isShowingSheet.toggle()
                } label: {
                    Text("Transfere Scaffolding")
                        .padding(12)
                        .font(.system(size: 20))
                        .foregroundColor(colorScheme == .dark ? Color(UIColor.black) : Color(UIColor.darkGray))
                }
                .contentShape(Rectangle())
                .background(colorScheme == .dark ? Color(UIColor.white) : Color(UIColor.white)).cornerRadius(7)
                .sheet(isPresented: $isShowingSheet,
                       onDismiss: didDismiss) {
                    TransfereScaffolding()
                }
             }
        }
        .ignoresSafeArea(edges: .top)
    }
    func didDismiss() {
        // Handle the dismissing action.
    }
}

struct DetailView: View {
    var project: Project

    @Environment(\.colorScheme) var colorScheme

    var body: some View {
        let projectInfoTitle = "Project Information"
        let duration = "Duration:"
        let customer = "Customer:"
        let amountScaff = "Amount:"
        
        VStack {
            VStack(alignment: .leading) {
                Text(projectInfoTitle)
                    .font(.title).bold()
                
                HStack {
                    Text(duration)
                        .font(.body).bold()
                    
                    Text("\(project.period.startDate) - \(project.period.endDate)")
                        .font(.body)
                }
                
                HStack {
                    Text(customer)
                        .font(.body).bold()
                    
                    Text("\(project.customer.name)")
                        .font(.body)
                }
                
                HStack {
                    Text(amountScaff)
                        .font(.body).bold()
                    
                    Text("\("ADD INFO TO API")")
                        .font(.body)
                }
            }
            .foregroundColor(Color(UIColor.darkGray))
            .lineLimit(1)
            .layoutPriority(100)
            .frame(width: 350, height: 125)
            .background(colorScheme == .dark ? Color.white : Color(UIColor.white))
            .cornerRadius(15)
            .shadow(color: Color(UIColor.black).opacity(0.1), radius: 5, x: 0, y: 2)
            .shadow(color: Color(UIColor.black).opacity(0.2), radius: 20, x: 0, y: 10)
            .overlay(
                    RoundedRectangle(cornerRadius: 15)
                        .stroke(colorScheme == .dark ? Color.gray.opacity(0.1) : Color.gray.opacity(0.1), lineWidth: 1)
                )
        }
    }
}
