//
//  CalendarView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 27/04/2022.
//

import SwiftUI

/// **CalendarView**
/// Displays the period filter View with two calendars for the user to interact with.
struct CalendarView: View {
    /// The selected start date and end date
    @Binding var selStartDate: Date
    @Binding var selEndDate: Date
    
    /// Checks if the period filter is active
    @Binding var periodFilterActive: Bool

    /// Base values for the calendar
    @State private var startDate = Date()
    @State private var endDate = Date()
    
    var body: some View {
        VStack {
            ScrollView {
                VStack {
                    Section {
                        VStack {
                            /// First calendar
                            DatePicker(
                                "Start dato",
                                selection: $startDate,
                                in: ...$endDate.wrappedValue,
                                displayedComponents: [.date]
                            )
                            
                            Divider()
                            
                            /// Second calendar
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
            
            /// Returns the selected dates to the parent View
            Button(action: {
                selStartDate = $startDate.wrappedValue
                selEndDate = $endDate.wrappedValue
                periodFilterActive = true
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
struct CalendarView_Previews: PreviewProvider {
    static var previews: some View {
        CalendarView()
    }
}*/ 
