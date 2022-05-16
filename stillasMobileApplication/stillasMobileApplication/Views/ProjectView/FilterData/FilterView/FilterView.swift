//
//  FilterView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 25/04/2022.
//

import SwiftUI

// TODO: Add enum for switch case instead of hard-coded values

/// **FilterView**
/// Creates a View displaying the a NavigationView for the different project filters.
struct FilterView: View {
    /// Filter options
    @State private var filterItems = ["Område", "Periode", "Størrelse", "Status"]
    
    /// All the different filtervalues
    @Binding var selStartDateBind: Date
    @Binding var selEndDateBind: Date
    @Binding var projectArea: String
    @Binding var projectSize: Int
    @Binding var projectStatus: String
    @Binding var minProjectSize: Int
    @Binding var maxProjectSize: Int
    @Binding var sizeSortType: String
    
    /// Returns whether a filter is active or not
    @State var periodFilterActive: Bool = false
    @State var areaFilterActive: Bool = false
    @State var sizeFilterActive: Bool = false
    @State var statusFilterActive: Bool = false
    
    /// Array of all the filters active
    @Binding var filterArr: [String]
    
    /// All selected areas to be filtered on
    @Binding var filterArrArea: [String]

    /// Selected filter start date and end date
    @State var selStartDate = Date()
    @State var selEndDate = Date()
    
    var body: some View {
        NavigationView {
            List {
                /// For each filter type, add it to the List with its respective navigation destination
                ForEach(filterItems, id: \.self) { filterItem in
                    NavigationLink {
                        switch filterItem {
                        case "Område":
                            FilterProjectArea(selArr: $filterArrArea, areaFilterActive: $areaFilterActive)
                                .onAppear {
                                    /// Resets the filter
                                    filterArrArea.removeAll()
                                }
                        case "Periode":
                            FilterProjectPeriod(selStartDateBind: $selStartDate, selEndDateBind: $selEndDate, periodFilterActiveBind: $periodFilterActive)
                                .onAppear {
                                    /// Resets the filter
                                    selStartDateBind = Date.distantPast
                                    selEndDateBind = Date.distantFuture
                                    if (selStartDateBind != Date.distantPast || selEndDateBind != Date.distantFuture) {
                                        periodFilterActive = true
                                    }
                                }
                                .onChange(of: selStartDate) { selectedStartDate in
                                    selStartDateBind = selectedStartDate
                                }
                                .onChange(of: selEndDate) { selectedEndDate in
                                    selEndDateBind = selectedEndDate
                                }
                        case "Størrelse":
                            FilterProjectSize(scoreFromBind: $minProjectSize, scoreToBind: $maxProjectSize, sizeFilterActive: $sizeFilterActive, selection: $sizeSortType)
                                .onChange(of: minProjectSize) { selectedMinSize in
                                    minProjectSize = selectedMinSize
                                    sizeFilterActive = true
                                }
                                .onChange(of: maxProjectSize) { selectedMaxSize in
                                    maxProjectSize = selectedMaxSize
                                    sizeFilterActive = true
                                }
                        case "Status":
                            FilterProjectStatus(filterArr: $filterArr, selection: $projectStatus)
                                .onChange(of: projectStatus) { status in
                                    projectStatus = status
                                    statusFilterActive = true
                                }
                        default:
                            Text("No views available")
                        }
                    } label: {
                        /// Adds label to list as well as adding the filter preview to the list item if it is active
                        HStack {
                            Text(filterItem)
                            Spacer()
                            switch filterItem {
                            case "Område":
                                HStack {
                                    ActiveAreaFilterView(filterArr: $filterArrArea, areaFilterActive: $areaFilterActive)
                                }
                                .lineLimit(1)

                            case "Periode":
                                HStack {
                                    ActivePeriodFilterView(startDate: $selStartDateBind, endDate: $selEndDateBind, filterArr: $filterArr, periodFilterActive: $periodFilterActive)
                                }
                                .lineLimit(1)
                            case "Størrelse":
                                HStack {
                                    ActiveSizeFilterView(filterArr: $filterArr, projectMinSize: $minProjectSize, projectMaxSize: $maxProjectSize, sizeFilterActive: $sizeFilterActive, selection: $sizeSortType)

                                }
                                .lineLimit(1)
                            case "Status":
                                HStack {
                                    ActiveStatusFilterView(filterArr: $filterArr, projectStatus: $projectStatus, statusFilterActive: $statusFilterActive)
                                }
                                .lineLimit(1)
                            default:
                                Text("No views available")
                            }
                        }
                    }
                }
            }
            .onAppear {
                if !filterArrArea.isEmpty {
                    areaFilterActive = true
                }
                if ((selStartDateBind != Date.distantPast || selEndDateBind != Date.distantFuture) && filterArr.contains("period")) {
                    periodFilterActive = true
                }
                if (filterArr.contains("size")) {
                    sizeFilterActive = true
                }
                if (filterArr.contains("status")) {
                    statusFilterActive = true
                }
            }
            .navigationTitle(Text("Filter"))
            .navigationViewStyle(StackNavigationViewStyle())
            .overlay(alignment: .bottom) {
                /// Adds and removes active filters based on current selection
                Button(action: {
                    // TODO: Change to use for loop?
                    /*for filterItem in filterArr {
                        addFilterItem(filterItem: filterItem)
                    }*/
                    if areaFilterActive {
                        addFilterItem(filterItem: "area")
                    } else {
                        deleteFilterItem(filterItem: "area")
                    }
                    if periodFilterActive {
                        addFilterItem(filterItem: "period")
                    } else {
                        deleteFilterItem(filterItem: "period")
                    }
                    if sizeFilterActive {
                        addFilterItem(filterItem: "size")
                    } else {
                        deleteFilterItem(filterItem: "size")
                    }
                    if statusFilterActive {
                        addFilterItem(filterItem: "status")
                    } else {
                        deleteFilterItem(filterItem: "status")
                    }
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
    
    
    /// If the selected filter is newly added and not only updated, add it to the array of filters
    /// - Parameter filterItem: the selected filter you want to add
    func addFilterItem(filterItem: String){
        if !filterArr.contains(filterItem) {
            filterArr.append(filterItem)
        }
    }
    
    /// Remove filter item from the array of filters
    /// - Parameter filterItem: the selected filter you want to remove
    func deleteFilterItem(filterItem: String) {
        if let i = filterArr.firstIndex(of: filterItem) {
            filterArr.remove(at: i)
        }
    }
}

/*
struct FilterView_Previews: PreviewProvider {
    static var previews: some View {
        FilterView()
    }
}*/
