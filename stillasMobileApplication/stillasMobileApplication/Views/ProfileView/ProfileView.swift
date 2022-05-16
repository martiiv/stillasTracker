//
//  ProfileView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/03/2022.
//

import SwiftUI

/// **ProfileView**
/// Calls the ProfileDetails View containing information about a user
struct ProfileView: View {
    var body: some View {
        VStack {
            ProfileDetails()
        }
    }
}

/// **ProfileDetails**
/// A View responsible for the layout of the user information and showing the details about the user
struct ProfileDetails: View {
    /// Darkmode or lightmode?
    @Environment(\.colorScheme) var colorScheme

    /// Models
    @EnvironmentObject var viewModel: AppViewModel
    @ObservedObject var profileModel: ProfileData = ProfileData()

    /// Initializes a user object
    @State var user: [Profile] = [Profile]()
    
    var body: some View {
        ScrollView {
            /// MapView displaying the map in the top of the screen
            MapView()
                .ignoresSafeArea(edges: .top)
                .frame(height: 300)
            /// CircleImage responsible for displaying the user profile image
            CircleImage(image: Image("UserProfile"))
                .offset(y: -130)
                .padding(.bottom, -130)
        
            /// If there are user data
            if (!user.isEmpty) {
                
                VStack {
                    VStack {
                        Image(systemName: "person.crop.circle.badge.checkmark")
                            .resizable()
                            .frame(width: 35, height: 30)
                            .foregroundColor(.blue)
                        
                        Text("Bruker info")
                            .font(Font.system(size: 20).bold())
                            .padding(.bottom, 2)
                        
                        Text("Nedenfor finner du brukerinformasjonen din.")
                            .font(.caption)
                            .foregroundColor(Color.gray)
                            .padding(.bottom, 5)
                    }
                                            
                    VStack {
                        HStack {
                            Text("\(user[0].name.firstName) \(user[0].name.lastName)")
                       }
                        .font(.title3.bold())
                    }
                    .padding(.bottom, 5)
                    
                    VStack {
                        Text("\(user[0].employeeID)")
                                .font(.body)
                        
                        Text("ANSATT NUMMER")
                            .foregroundColor(.gray)
                            .font(.system(size: 15))
                    }
                    .padding(.bottom, 5)
                    
                    VStack {
                        Text("\(user[0].dateOfBirth)")
                                .font(.body)
                        
                        Text("FØDSELSNUMMER")
                            .foregroundColor(.gray)
                            .font(.system(size: 15))
                    }
                    .padding(.bottom, 5)

                    VStack {
                        Text("\(user[0].role)")
                                .font(.body)
                        
                        Text("ROLLE")
                            .foregroundColor(.gray)
                            .font(.system(size: 15))
                    }
                    .padding(.bottom, 5)

                    VStack {
                        Text("\(user[0].admin.description)")
                                .font(.body)
                        
                        Text("ADMIN")
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
            
                VStack {
                    VStack {
                        Image(systemName: "person.circle")
                            .resizable()
                            .frame(width: 30, height: 30)
                            .foregroundColor(.blue)
                        
                        Text("Kontakt info")
                            .font(Font.system(size: 20).bold())
                            .padding(.bottom, 2)
                        
                        Text("Nedenfor finner du kontaktinformasjonen din.")
                            .font(.caption)
                            .foregroundColor(Color.gray)
                            .padding(.bottom, 5)
                    }
                    
                    VStack {
                        Text("\(user[0].phone)")
                                .font(.body)
                        
                        Text("TELEFONNUMMER")
                            .foregroundColor(.gray)
                            .font(.system(size: 15))
                    }
                    .padding(.bottom, 5)

                    VStack {
                        Text("\(user[0].email)")
                                .font(.body)
                        
                        Text("EMAIL")
                            .foregroundColor(.gray)
                            .font(.system(size: 15))
                    }
                    .padding(.bottom, 5)
                }
                .padding()
                .frame(width: (UIScreen.screenWidth / 1.2), alignment: .center)
                .contentShape(RoundedRectangle(cornerRadius: 5))
                .background(colorScheme == .dark ? Color(UIColor.white) : Color(UIColor.white)).cornerRadius(7)
                .shadow(color: Color(UIColor.black).opacity(0.1), radius: 5, x: 0, y: 2)
                .shadow(color: Color(UIColor.black).opacity(0.2), radius: 20, x: 0, y: 10)
            }
        }
        .task {
            /// Get user data from the API with the logged in users ID
            await profileModel.loadData(userID: viewModel.userID) { (user) in
                self.user.append(user)
            }
        }
        .ignoresSafeArea(edges: .top)
    }
}

struct ProfileView_Previews: PreviewProvider {
    static var previews: some View {
        ProfileView()
    }
}
