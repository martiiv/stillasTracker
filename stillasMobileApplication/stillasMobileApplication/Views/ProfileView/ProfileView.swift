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
            ProfileDetails()
        }
    }
}

/**
    ProfileDetails - A view responsible for the layout of the user information and showing the details about the user
 */
struct ProfileDetails: View {
    @EnvironmentObject var viewModel: AppViewModel
    @ObservedObject var profileModel: ProfileData = ProfileData()

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
        
            if (!user.isEmpty) {
            /// A VStack used to display all the user profile data
            VStack(alignment: .leading) {
                HStack {
                    Text("\(user[0].name.firstName) \(user[0].name.lastName)")
                       .font(.largeTitle)
               }
                
                HStack {
                    Text("MBStillas")
                        //.font(.subheadline)
                    Spacer()
                    Text("Rolle: \(user[0].role)")
                        //.font(.subheadline)
                }
                //.font(.subheadline)
                .foregroundColor(.secondary)
                
                Divider()

                VStack {
                    Text("Fødselsdato")
                        .font(.title2)
                    Text("\(user[0].dateOfBirth)")
                        .foregroundColor(.secondary)
                }
            }
            .padding()
            
            Spacer()
            
            Button (action: {
                viewModel.signOut()
            }) {
                Text("Logg av")
                    .frame(width: 150, height: 50, alignment: .center)
            }
            .foregroundColor(.white)
            .background(Color.blue)
            .cornerRadius(10)
            
            Spacer()
                .frame(height:50)  // limit spacer size by applying a frame
            }
        }
        .task {
            await profileModel.loadData(userID: viewModel.userID) { (user) in
                self.user.append(user)
            }
        }
        .ignoresSafeArea(edges: .top)
    }
}

/*
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
    @EnvironmentObject var viewModel: AppViewModel

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
                    Text("Rolle: \(user.role)")
                        //.font(.subheadline)
                }
                //.font(.subheadline)
                .foregroundColor(.secondary)
                
                Divider()

                VStack {
                    Text("Fødselsdato")
                        .font(.title2)
                    Text("\(user.dateOfBirth)")
                        .foregroundColor(.secondary)
                }
            }
            .padding()
            
            Spacer()
            
            Button (action: {
                viewModel.signOut()
            }) {
                Text("Logg av")
                    .frame(width: 150, height: 50, alignment: .center)
            }
            .foregroundColor(.white)
            .background(Color.blue)
            .cornerRadius(10)
            
            Spacer()
                .frame(height:50)  // limit spacer size by applying a frame
        }
        .navigationTitle(user.name)
        .navigationBarTitleDisplayMode(.inline)
        .ignoresSafeArea(edges: .top)
    }
}*/

struct ProfileView_Previews: PreviewProvider {
    static var previews: some View {
        ProfileView()
    }
}
