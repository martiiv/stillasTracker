//
//  StatusFilter.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/04/2022.
//

import SwiftUI

/// **FilterProjectStatus**
/// The View for selecting a status filter
struct FilterProjectStatus: View {
    /// Array of all the filters active
    @Binding var filterArr: [String]
    
    /// The selected filter
    @Binding var selection: String

    /// The different project states
    let states = ["Inactive", "Active", "Upcomming"]

    var body: some View {
        VStack {
            Text("Select project status")
                .font(.headline)
                .padding(.vertical, 5)
            
            /// Picker for choosing a project state
            Picker("Select a state: ", selection: $selection) {
                ForEach(states, id: \.self) {
                    Text($0)
                }
            }
            .pickerStyle(SegmentedPickerStyle())
            
            HStack {
                Text("Selected state: ")
                +
                Text("\(selection)")
                    .bold()
                    
            }
            .padding(.vertical, 10)

            Spacer()
        }
        .navigationTitle(Text("Status"))
        .overlay(alignment: .bottom) {
            /// Adds the selection and adds filter to filter array
            Button(action: {
                selection = selection
                filterArr.append("status")
            }) {
                Text("Bruk")
                    .frame(width: 300, height: 50, alignment: .center)
            }
            .foregroundColor(.white)
            .background(Color.blue)
            .cornerRadius(10)
            .padding(.bottom, 50)
        }
    }
}

/*
struct StatusFilter_Previews: PreviewProvider {
    static var previews: some View {
        FilterProjectStatus()
    }
}*/
