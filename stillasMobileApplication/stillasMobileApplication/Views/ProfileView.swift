//
//  ProfileView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/03/2022.
//

import SwiftUI

struct ProfileView: View {
    var body: some View {
        VStack {
            ProfileDetails(user: ModelData().users[0])
        }
    }
}

struct ProfileView_Previews: PreviewProvider {
    static var previews: some View {
        ProfileView()
    }
}




struct ProfileDetails: View {
    @EnvironmentObject var modelData: ModelData
    var user: User

    var userIndex: Int {
        modelData.users.firstIndex(where: { $0.id == user.id })!
    }
    
    var body: some View {
        ScrollView {
            MapView()
                .ignoresSafeArea(edges: .top)
                .frame(height: 300)
            
            CircleImage(image: user.image)
                .offset(y: -130)
                .padding(.bottom, -130)
            
            VStack(alignment: .leading) {
                HStack {
                   Text(user.name)
                       .font(.title)
                   //FavoriteButton(isSet: $modelData.landmarks[landmarkIndex].isFavorite)
               }


                HStack {
                    Text("Role: " + user.role)
                        .font(.subheadline)
                    Spacer()
                    Text(user.name)
                        .font(.subheadline)
                }
                .font(.subheadline)
                .foregroundColor(.secondary)
                
                Divider()

                Text("Date of birth of user: \(user.name)")
                    .font(.title2)
                Text(user.dateOfBirth)
            }
            .padding()
            
        Spacer()
        }
        .navigationTitle(user.name)
        .navigationBarTitleDisplayMode(.inline)
    }
}

