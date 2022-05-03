//
//  ProfileView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/03/2022.
//

import SwiftUI

/**
    ProfileView - Calls the ProfileDetails view containing information about a user
 */
struct ProfileView: View {
    var body: some View {
        VStack {
            // TODO: Change input to not take the ModelData's first element only, but get info from API
            ProfileDetails(user: ModelData().users[0])
        }
    }
}

/**
    ProfileDetails - A view responsible for the layout of the user information and showing the details about the user
 */
struct ProfileDetails: View {
    @EnvironmentObject var modelData: ModelData
    
    var user: User
    
    /// Retrieves the user from the json object with ID equal to the object passed into the voew
    var userIndex: Int {
        modelData.users.firstIndex(where: { $0.id == user.id })!
    }
    
    var body: some View {
        ScrollView {
            /// MapView displaying the map in the top of the screen
            MapView()
                .ignoresSafeArea(edges: .top)
                .frame(height: 300)
        
            /// CircleImage responsible for displaying the user profile image
            CircleImage(image: user.image)
                .offset(y: -130)
                .padding(.bottom, -130)
        
            /// A VStack used to display all the user profile data
            VStack(alignment: .leading) {
                HStack {
                   Text(user.name)
                       .font(.largeTitle)
               }
                
                HStack {
                    // TODO: Change to not hard coded values when API is updated
                    Text("MBStillas")
                        //.font(.subheadline)
                    Spacer()
                    Text("Role: \(user.role)")
                        //.font(.subheadline)
                }
                //.font(.subheadline)
                .foregroundColor(.secondary)
                
                Divider()

                VStack {
                    Text("Date of birth")
                        .font(.title2)
                    Text("\(user.dateOfBirth)")
                        .foregroundColor(.secondary)
                }
        
                Spacer()
            }
            .padding()
            
        Spacer()
        }
        .navigationTitle(user.name)
        .navigationBarTitleDisplayMode(.inline)
    }
}

struct ProfileView_Previews: PreviewProvider {
    static var previews: some View {
        ProfileView()
    }
}
