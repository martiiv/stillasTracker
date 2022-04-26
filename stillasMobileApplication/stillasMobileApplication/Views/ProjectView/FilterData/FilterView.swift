//
//  FilterView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 25/04/2022.
//

import SwiftUI

// TODO: Add enum for switch case instead of hard-coded values

struct FilterView: View {
    @State private var filterItems = ["Område", "Periode", "Størrelse", "Status"]
    
    @Binding var selStartDateBind: Date
    @Binding var selEndDateBind: Date
    @Binding var projectArea: String
    @Binding var projectSize: Int
    @Binding var projectStatus: String
    
    @Binding var filterArr: [String]
    
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
                        case "Periode":
                            FilterProjectPeriod(selStartDateBind: $selStartDate, selEndDateBind: $selEndDate)
                                .onChange(of: selStartDate) { selectedStartDate in
                                    selStartDateBind = selectedStartDate
                                }
                                .onChange(of: selEndDate) { selectedEndDate in
                                    selEndDateBind = selectedEndDate
                                }
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
                        HStack {
                            Text(filterItem)
                            Spacer()
                            HStack {
                                ActiveFilterView(startDate: selStartDate, endDate: selEndDate)
                            }
                            .lineLimit(1)
                        }
                    }
                }
            }
            .navigationTitle(Text("Filter"))
            .navigationViewStyle(StackNavigationViewStyle())
            .overlay(alignment: .bottom) {
                Button(action: {
                    addFilterItem(filterItem: "period")
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
    
    func addFilterItem(filterItem: String){
        if !filterArr.contains(filterItem) {
            filterArr.append(filterItem)
        }
    }
    
    func deleteFilterItem(filterItem: String) {
        if let i = filterArr.firstIndex(of: filterItem) {
            filterArr.remove(at: i)
        }
    }
}

struct ActiveFilterView: View {
    @State var startDate: Date
    @State var endDate: Date
    
    var body: some View {
        HStack {
            Text("\(formatStringDate(date: startDate.description)) - \(formatStringDate(date: endDate.description))")
                .font(.system(size: 11).bold())
                .padding(.vertical, 5)
                .padding(.leading, 5)
            
            Button(action: {
                print("It works")
            }) {
                Image(systemName: "x.circle.fill")
                    .foregroundColor(Color.secondary)
            }
            .padding(.trailing, 5)
            .buttonStyle(PlainButtonStyle())

        }
        .foregroundColor(.white)
        .background(Color.blue)
        .cornerRadius(5)
        .padding(.vertical, 5)
    }
    
    func formatStringDate(date: String) -> String {
        let dateFormatter = DateFormatter()
        dateFormatter.dateFormat = "yy-MM-dd'T'HH:mm:ss.SSSZ"
        let newDate = dateFormatter.date(from: date)
        dateFormatter.setLocalizedDateFormatFromTemplate("MMMM d, yy")
        return dateFormatter.string(from: newDate ?? Date.now)
    }
}

/*
struct FilterView_Previews: PreviewProvider {
    static var previews: some View {
        FilterView()
    }
}*/
