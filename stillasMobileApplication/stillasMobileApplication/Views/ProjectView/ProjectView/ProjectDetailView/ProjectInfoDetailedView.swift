//
//  ProjectInfoDetailedView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 29/04/2022.
//

import SwiftUI
import UIKit

struct ProjectInfoDetailedView: View {
    /// Selected project
    var project: Project

    /// Darkmode or lightmode?
    @Environment(\.colorScheme) var colorScheme
    
    /// Predefined to be easily maintianed and changed
    let duration = "VARIGHET"
    let customer = "KUNDE"
    let amountScaff = "MENGDE"
    let scaffoldingSize = "STØRRELSE"
    let state = "STATUS"
    let address = "ADRESSE"
    let contactPerson = "KONTAKT PERSON"
    let phoneNumber = "MOBILNUMMER"
    let email = "EMAIL"
    
    var body: some View {
        VStack {
            VStack(alignment: .leading) {
                /// Project info card
                VStack(alignment: .leading) {
                    VStack {
                        VStack {
                            Image(systemName: "square.text.square")
                                .resizable()
                                .frame(width: 30, height: 30)
                                .foregroundColor(.blue)
                            
                            Text("Prosjekt info")
                                .font(Font.system(size: 20).bold())
                                .padding(.bottom, 2)
                            
                            Text("Nedenfor finner du informasjon om dette prosjektet.")
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
                
                /// Contact info card
                VStack {
                    VStack {
                        Image(systemName: "person.circle")
                            .resizable()
                            .frame(width: 30, height: 30)
                            .foregroundColor(.blue)
                        
                        Text("Kontakt info")
                            .font(Font.system(size: 20).bold())
                            .padding(.bottom, 2)
                        
                        Text("Nedenfor finner du kontaktinformasjonen til kunden.")
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
}

/*
struct ProjectInfoDetailedView_Previews: PreviewProvider {
    static var previews: some View {
        ProjectInfoDetailedView()
    }
}*/
