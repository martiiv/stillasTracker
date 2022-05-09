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
    
    //let projectInfoTitle = "Project Information"
    let duration = "DURATION"
    let customer = "CUSTOMER"
    let amountScaff = "AMOUNT"
    let scaffoldingSize = "SIZE"
    let state = "STATUS"
    let address = "ADDRESS"
    let contactPerson = "CONTACT PERSON"
    let phoneNumber = "PHONE NUMBER"
    let email = "EMAIL"
    
    var body: some View {
        VStack {
            VStack(alignment: .leading) {
                //Text(projectInfoTitle)
                    //.font(.title).bold()
                    //.padding(.bottom, 5)
                
                    VStack(alignment: .leading) {
                        VStack {
                            VStack {
                                Image(systemName: "square.text.square")
                                    .resizable()
                                    .frame(width: 30, height: 30)
                                    .foregroundColor(.blue)
                                
                                Text("Project info")
                                    .font(Font.system(size: 20).bold())
                                    .padding(.bottom, 2)
                                
                                Text("Below you find the project info for this project.")
                                    .font(.caption)
                                    .foregroundColor(Color.gray)
                                    .padding(.bottom, 5)
                            }
                            
                            VStack {
                                Text("\(project.customer.name)")
                                    .font(.body)

                                Text(customer)
                                    .foregroundColor(.gray)
                                    .font(.system(size: 15))
                            }
                            .padding(.bottom, 5)

                            VStack(alignment: .leading) {
                                VStack {
                                    Text("\(project.period.startDate)  -  \(project.period.endDate)")
                                        .font(.body)
                                    
                                    Text(duration)
                                        .foregroundColor(.gray)
                                        .font(.system(size: 15))
                                }
                            }
                            .padding(.bottom, 5)

                            VStack {
                                HStack {
                                    Text("\(project.size) m")
                                    + Text("2")
                                        .baselineOffset(6)
                                        .font(.system(size: 12))
                                }
                                .font(.body)

                                Text(scaffoldingSize)
                                    .foregroundColor(.gray)
                                    .font(.system(size: 15))
                                    .padding(.bottom, 5)
                            }
                            
                            VStack {
                                Text("\(project.state)")
                                    .font(.body)
                                
                                Text(state)
                                    .foregroundColor(.gray)
                                    .font(.system(size: 15))
                            }
                        }
                    }
                    .padding()
                    .frame(width: (UIScreen.screenWidth / 1.2), alignment: .center)
                    .contentShape(RoundedRectangle(cornerRadius: 5))
                    .background(colorScheme == .dark ? Color(UIColor.white) : Color(UIColor.white)).cornerRadius(7)
                    .shadow(color: Color(UIColor.black).opacity(0.1), radius: 5, x: 0, y: 2)
                    .shadow(color: Color(UIColor.black).opacity(0.2), radius: 20, x: 0, y: 10)
                    
                    VStack {
                        VStack {
                            Image(systemName: "person.circle")
                                .resizable()
                                .frame(width: 30, height: 30)
                                .foregroundColor(.blue)
                            
                            Text("Contact info")
                                .font(Font.system(size: 20).bold())
                                .padding(.bottom, 2)
                            
                            Text("Below you find the contact info for this project.")
                                .font(.caption)
                                .foregroundColor(Color.gray)
                                .padding(.bottom, 5)
                        }
                                                
                        VStack {
                            Text("\(project.customer.name)")
                                    .font(.body)
                            
                            Text(contactPerson)
                                .foregroundColor(.gray)
                                .font(.system(size: 15))
                        }
                        .padding(.bottom, 5)
                        
                        VStack {
                            Text("\(project.address.street), \(project.address.zipcode) \(project.address.municipality)")
                                    .font(.body)
                            
                            Text(address)
                                .foregroundColor(.gray)
                                .font(.system(size: 15))
                        }
                        .padding(.bottom, 5)
                        
                        VStack {
                            Text("\(project.customer.number)")
                                    .font(.body)
                            
                            Text(phoneNumber)
                                .foregroundColor(.gray)
                                .font(.system(size: 15))
                        }
                        .padding(.bottom, 5)

                        VStack {
                            Text("\(project.customer.email)")
                                    .font(.body)
                            
                            Text(email)
                                .foregroundColor(.gray)
                                .font(.system(size: 15))
                        }
                    }
                    .padding()
                    .frame(width: (UIScreen.screenWidth / 1.2), alignment: .center)
                    .contentShape(RoundedRectangle(cornerRadius: 5))
                    .background(colorScheme == .dark ? Color(UIColor.white) : Color(UIColor.white)).cornerRadius(7)
                    .shadow(color: Color(UIColor.black).opacity(0.1), radius: 5, x: 0, y: 2)
                    .shadow(color: Color(UIColor.black).opacity(0.2), radius: 20, x: 0, y: 10)
                }
            }
            .padding(.horizontal, 20)
        }
        //ProjectView1()

}

/*
struct ProjectInfoDetailedView_Previews: PreviewProvider {
    static var previews: some View {
        ProjectInfoDetailedView()
    }
}*/
