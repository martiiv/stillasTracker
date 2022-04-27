//
//  FilterProjectPeriod.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 20/04/2022.
//

import SwiftUI

struct FilterProjectPeriod: View {
    //@State private var date = Date()
    @Binding var selStartDateBind: Date
    @Binding var selEndDateBind: Date
    @Binding var periodFilterActiveBind: Bool
    @State var periodFilterActive: Bool = true
    @State private var selStartDate = Date()
    @State private var selEndDate = Date()
    
    
    var body: some View {
        VStack {
            CalendarView(selStartDate: $selStartDate, selEndDate: $selEndDate, periodFilterActive: $periodFilterActive)
                .onAppear {
                    selStartDateBind = selStartDate
                    selEndDateBind = selEndDate
                }
                .onChange(of: selStartDate) { selectedStartDate in
                    selStartDateBind = selectedStartDate
                    periodFilterActiveBind = true
                }
                .onChange(of: selEndDate) { selectedEndDate in
                    selEndDateBind = selectedEndDate
                    periodFilterActiveBind = true
                }
        }
        .navigationTitle(Text("Prosjekt periode"))
        //.ignoresSafeArea(edges: .top)
    }
}

struct CalendarView: View {
    @Binding var selStartDate: Date
    @Binding var selEndDate: Date
    @Binding var periodFilterActive: Bool

    @State private var startDate = Date()
    @State private var endDate = Date()
    
    var body: some View {
        VStack {
            ScrollView {
                VStack {
                    Section {
                        VStack {
                            DatePicker(
                                "Start dato",
                                selection: $startDate,
                                in: ...$endDate.wrappedValue,
                                displayedComponents: [.date]
                            )
                            
                            Divider()
                            
                            DatePicker(
                                "Slutt dato",
                                selection: $endDate,
                                in: $startDate.wrappedValue...,
                                displayedComponents: [.date]
                            )
                        }
                    }
                }
                .padding(.horizontal, 20)
                .padding(.top, 40)
            }
            Spacer()
            Button(action: {
                selStartDate = $startDate.wrappedValue
                selEndDate = $endDate.wrappedValue
                periodFilterActive = true
                print("______")
                print(selStartDate)
                print(selEndDate)
                print("______")
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
}

/*
struct FilterProjectPeriod_Previews: PreviewProvider {
    @State private var selStartDate = Date()
    @State private var selEndDate = Date()
    
    static var previews: some View {
        FilterProjectPeriod(selStartDateBind: selStartDate, selEndDateBind: selEndDate)
    }
}*/

