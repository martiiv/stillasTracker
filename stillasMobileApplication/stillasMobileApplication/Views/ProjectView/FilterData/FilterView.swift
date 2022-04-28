//
//  FilterView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 25/04/2022.
//

import SwiftUI

// TODO: Add enum for switch case instead of hard-coded values

struct FilterView: View {
    // TODO: Add buttons for switching between start before only, period between etc.
    @State private var filterItems = ["Område", "Periode", "Størrelse", "Status"]
    
    @Binding var selStartDateBind: Date
    @Binding var selEndDateBind: Date
    @Binding var projectArea: String
    @Binding var projectSize: Int
    @Binding var projectStatus: String
    @Binding var minProjectSize: Int
    @Binding var maxProjectSize: Int
    // TODO: DENNA TINGEN HER ISTEDENFOR scoreFrom
    
    @State var periodFilterActive: Bool = false
    @State var areaFilterActive: Bool = false
    @State var sizeFilterActive: Bool = false
    
    @Binding var filterArr: [String]
    @Binding var filterArrArea: [String]

    @State var selStartDate = Date()
    @State var selEndDate = Date()
        
    var body: some View {
        NavigationView {
            List {
                ForEach(filterItems, id: \.self) { filterItem in
                    NavigationLink {
                        switch filterItem {
                        case "Område":
                            FilterProjectArea(selArr: $filterArrArea, areaFilterActive: $areaFilterActive)
                                .onAppear {
                                    filterArrArea.removeAll()
                                }
                        case "Periode":
                            FilterProjectPeriod(selStartDateBind: $selStartDate, selEndDateBind: $selEndDate, periodFilterActiveBind: $periodFilterActive)
                                .onAppear {
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
                            FilterProjectSize(scoreFromBind: $minProjectSize, scoreToBind: $maxProjectSize, sizeFilterActive: $sizeFilterActive)
                                .onChange(of: minProjectSize) { selectedMinSize in
                                    minProjectSize = selectedMinSize
                                    sizeFilterActive = true
                                }
                                .onChange(of: maxProjectSize) { selectedMaxSize in
                                    maxProjectSize = selectedMaxSize
                                    sizeFilterActive = true
                                }
                            /*
                        case "Status":
                            print("Add status view")
                            // ADD status view
                        default:
                            print("Did not find any")
                        */
                        default:
                            Text("Default")
                            //AddProjectView()
                        }
                    } label: {
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
                                    ActiveSizeFilterView(filterArr: $filterArr, projectMinSize: $minProjectSize, projectMaxSize: $maxProjectSize, sizeFilterActive: $sizeFilterActive)
                                }
                                .lineLimit(1)
                            case "Status":
                                Text("")
                            default:
                                Text("")
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
                    print(periodFilterActive)
                    periodFilterActive = true
                    print(periodFilterActive)
                    print(filterArr)
                }
                if (filterArr.contains("size")) {
                    sizeFilterActive = true
                }
            }
            .navigationTitle(Text("Filter"))
            .navigationViewStyle(StackNavigationViewStyle())
            .overlay(alignment: .bottom) {
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
                    
                    print(filterArr)
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

struct ActivePeriodFilterView: View {
    @Binding var startDate: Date
    @Binding var endDate: Date

    @Binding var filterArr: [String]
    
    @Binding var periodFilterActive: Bool

    var body: some View {

        if periodFilterActive {
            HStack {
                HStack {
                    Text(startDate, style: .date)
                        .padding(.leading, 5)
                        .padding(.trailing, -5)
                    
                    Text("-")

                    Text(endDate, style: .date)
                        .padding(.trailing, -5)
                        .padding(.leading, -5)
                }
                .font(.system(size: 11).bold())
                .padding(.vertical, 5)
                
                Button(action: {
                    deleteFilterItem(filterItem: "period")
                    self.periodFilterActive.toggle()
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
    }
    
    func deleteFilterItem(filterItem: String) {
        if let i = filterArr.firstIndex(of: filterItem) {
            filterArr.remove(at: i)
        }
    }
}

struct ActiveAreaFilterView: View {
    @Binding var filterArr: [String]
    
    @Binding var areaFilterActive: Bool

    var body: some View {

        if areaFilterActive {
            HStack {
                HStack {
                    //ScrollViewReader { scrollView in
                        ScrollView (.horizontal, showsIndicators: false) {
                            HStack {
                                Text("(\(filterArr.count))")
                                    .padding(.leading, 4)
                                ForEach(filterArr.indices, id: \.self) { index in
                                    HStack {
                                        Text("\(filterArr[index])")
                                            .lineLimit(1)
                                            .padding(-3)
                                        if index != filterArr.count-1 {
                                            Text(",")
                                        }
                                    }
                                }
                            }
                        }
                    /*
                        .onAppear {
                            scrollView.scrollTo(filterArr[filterArr.endIndex])
                        }
                    }*/
                }
                .font(.system(size: 11).bold())
                .padding(.vertical, 5)
                .lineLimit(1)
                
                Button(action: {
                    deleteFilterItem(filterItem: "area")
                    self.areaFilterActive = false
                }) {
                    Image(systemName: "x.circle.fill")
                        .foregroundColor(Color.secondary)
                }
                .padding(.trailing, 5)
                .buttonStyle(PlainButtonStyle())
            }
            .frame(width: 150, alignment: .trailing)
            .foregroundColor(.white)
            .background(Color.blue)
            .cornerRadius(5)
            .padding(.vertical, 5)
        }
    }
    
    func deleteFilterItem(filterItem: String) {
        if let i = filterArr.firstIndex(of: filterItem) {
            filterArr.remove(at: i)
        }
    }
}

struct ActiveSizeFilterView: View {
    @Binding var filterArr: [String]
    @Binding var projectMinSize: Int
    @Binding var projectMaxSize: Int
    @Binding var sizeFilterActive: Bool

    var body: some View {

        if sizeFilterActive {
            HStack {
                HStack {
                    ScrollView (.horizontal, showsIndicators: false) {
                        HStack {
                            Text("\(projectMinSize) m")
                            + Text("2")
                                .baselineOffset(6)
                                .font(Font.system(size: 10))

                            + Text(" - ")
                            + Text("\(projectMaxSize) m")
                            + Text("2")
                                .baselineOffset(6)
                                .font(Font.system(size: 10))
                        }
                        .padding(.leading, 5)
                    }
                }
                .font(.system(size: 11).bold())
                .padding(.vertical, 5)
                .lineLimit(1)
                
                Button(action: {
                    deleteFilterItem(filterItem: "size")
                    self.sizeFilterActive = false
                }) {
                    Image(systemName: "x.circle.fill")
                        .foregroundColor(Color.secondary)
                }
                .padding(.trailing, 5)
                .buttonStyle(PlainButtonStyle())
            }
            .frame(alignment: .trailing)
            .scaledToFit()
            .foregroundColor(.white)
            .background(Color.blue)
            .cornerRadius(5)
            .padding(.vertical, 5)
        }
    }
    
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
