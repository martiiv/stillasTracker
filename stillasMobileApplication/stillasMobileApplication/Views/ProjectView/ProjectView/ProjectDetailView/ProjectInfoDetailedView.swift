//
//  ProjectInfoDetailedView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 29/04/2022.
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
        VStack {
            Section(header: Text("Second List")) {
                List(items, children: \.items) { row in
                    Image(systemName: row.icon)
                    Text(row.name)
                }
            }
        }
    }
}


extension BookMark {
    static let apple = BookMark(name: "Apple", icon: "1.circle")
    static let bbc = BookMark(name: "BBC", icon: "square.and.pencil")
    static let swift = BookMark(name: "Swfit", icon: "bolt.fill")
    static let twitter = BookMark(name: "Twitter", icon: "mic")

    static let example1 = BookMark(name: "Favorites", icon: "star", items: [BookMark.apple, BookMark.bbc, BookMark.swift, BookMark.twitter])
    static let example2 = BookMark(name: "Recent", icon: "timer", items: [BookMark.apple, BookMark.bbc, BookMark.swift, BookMark.twitter])
    static let example3 = BookMark(name: "Recommended", icon: "hand.thumbsup", items: [BookMark.apple, BookMark.bbc, BookMark.swift, BookMark.twitter])
}*/

struct ProjectInfoDetailedView: View {
    var project: Project

    @Environment(\.colorScheme) var colorScheme
    
    let projectInfoTitle = "Project Information"
    let duration = "Duration:"
    let customer = "Customer:"
    let amountScaff = "Amount:"
    let scaffoldingSize = "Size:"
    let state = "Status:"
    let address = "Address:"
    
    var body: some View {
        VStack {
        VStack(alignment: .leading) {
            Text(projectInfoTitle)
                .font(.title).bold()
                .padding(.bottom, 5)
            
            VStack(alignment: .leading) {
                VStack(alignment: .leading) {
                    HStack {
                        Text(duration)
                            .font(.body).bold()
                        
                        Text("\(project.period.startDate) - \(project.period.endDate)")
                            .font(.body)
                    }
                    
                    Divider()
                }
                
                VStack(alignment: .leading) {
                    HStack {
                        // TODO: Add dropdown list?
                        Text(customer)
                            .font(.body).bold()
                        
                        Text("\(project.customer.name)")
                            .font(.body)
                    }
                
                    Divider()
                }
                
                VStack(alignment: .leading) {
                    HStack {
                        Text(amountScaff)
                            .font(.body).bold()
                        
                        Text("\("ADD INFO TO API")")
                            .font(.body)
                    }
                    
                    Divider()
                }
                
                VStack(alignment: .leading) {
                    HStack {
                        Text(scaffoldingSize)
                            .font(.body).bold()
                        
                        HStack {
                            Text("\(project.size) m")
                            + Text("2")
                                .baselineOffset(6)
                                .font(.system(size: 12))
                        }
                        .font(.body)
                    }
                    
                    Divider()
                }
                
                VStack(alignment: .leading) {
                    HStack {
                        Text(state)
                            .font(.body).bold()
                        
                        Text("\(project.state)")
                            .font(.body)
                    }
                    
                    Divider()
                }
                
                VStack(alignment: .leading) {
                    HStack {
                        Text(address)
                            .font(.body).bold()
                        Text("\(project.address.street), \(project.address.zipcode) \(project.address.municipality)")
                                .font(.body)
                    }
                }
            }
        }
        .padding(.horizontal, 20)
        }
        //ProjectView1()

        }
    
    /*
    var body: some View {
        let projectInfoTitle = "Project Information"
        let duration = "Duration:"
        let customer = "Customer:"
        let amountScaff = "Amount:"
        let scaffoldingSize = "Size:"
        let state = "Status:"
        let address = "Address:"
        
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
                
                HStack {
                    Text(scaffoldingSize)
                        .font(.body).bold()
                    
                    HStack {
                        Text("\(project.size) m")
                        + Text("2")
                            .baselineOffset(6)
                            .font(.system(size: 12))
                    }
                    .font(.body)
                }
                
                HStack {
                    Text(state)
                        .font(.body).bold()
                    
                    Text("\(project.state)")
                        .font(.body)
                }
                
                HStack {
                    Text(address)
                        .font(.body).bold()
                    Text("\(project.address.street), \(project.address.zipcode) \(project.address.municipality)")
                            .font(.body)
                }
            }
            .foregroundColor(Color(UIColor.darkGray))
            .lineLimit(1)
            .layoutPriority(100)
            .frame(width: 350)
            .scaledToFit()
            .background(colorScheme == .dark ? Color.white : Color(UIColor.white))
            .cornerRadius(15)
            .shadow(color: Color(UIColor.black).opacity(0.1), radius: 5, x: 0, y: 2)
            .shadow(color: Color(UIColor.black).opacity(0.2), radius: 20, x: 0, y: 10)
            .overlay(
                    RoundedRectangle(cornerRadius: 15)
                        .stroke(colorScheme == .dark ? Color.gray.opacity(0.1) : Color.gray.opacity(0.1), lineWidth: 1)
                )
        }
    }*/
}
/*
struct ProjectInfoDetailedView_Previews: PreviewProvider {
    static var previews: some View {
        ProjectInfoDetailedView()
    }
}*/
