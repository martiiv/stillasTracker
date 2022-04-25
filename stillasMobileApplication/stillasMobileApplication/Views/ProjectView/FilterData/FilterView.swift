//
//  FilterView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 25/04/2022.
//

import SwiftUI

struct FilterView: View {
    @State private var filterItems = ["Område", "Prosjekt periode", "Størrelse", "Status"]
    
    @Binding var selStartDateBind: Date
    @Binding var selEndDateBind: Date
    @Binding var projectArea: String
    @Binding var projectSize: Int
    @Binding var projectStatus: String
    
    @State var selStartDate = Date()
    @State var selEndDate = Date()
        
    var body: some View {
        NavigationView {
            List {
                ForEach(filterItems, id: \.self) { filterItem in
                    NavigationLink {
                        switch filterItem {
                        case "Område":
                            FilterProjectArea()
                        case "Prosjekt periode":
                            FilterProjectPeriod(selStartDateBind: $selStartDate, selEndDateBind: $selEndDate)
                        case "Størrelse":
                            FilterProjectSize()
                            /*
                        case "Status":
                            print("Add status view")
                            // ADD status view
                        default:
                            print("Did not find any")
                        */
                        default:
                            AddProjectView()
                        }
                    } label: {
                        Text(filterItem)
                    }
                }
            }
            .navigationTitle(Text("Filter"))
            .navigationViewStyle(StackNavigationViewStyle())
        }
        Spacer()
        Button(action: {
            selStartDateBind = $selStartDate.wrappedValue
            selEndDateBind = $selEndDate.wrappedValue
            print("______")
            
            print(selStartDateBind)
            print(selEndDateBind)

            
        }) {
            Text("Bruk")
                .frame(width: 300, height: 50, alignment: .center)
        }
        .foregroundColor(.white)
        .background(Color.blue)
        .cornerRadius(10)
        
        Spacer()
            .frame(height:50)  // limit spacer size by applying a frame
    }
}
/*
struct FilterView_Previews: PreviewProvider {
    static var previews: some View {
        FilterView()
    }
}*/
