//
//  FilterProjectPeriod.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 20/04/2022.
//

import SwiftUI

/// **FilterProjectPeriod**
/// The View for selecting a period filter
struct FilterProjectPeriod: View {
    /// Selected start date and end date
    @Binding var selStartDateBind: Date
    @Binding var selEndDateBind: Date
    
    /// Tells if filter is activated or not
    @Binding var periodFilterActiveBind: Bool
    @State var periodFilterActive: Bool = true
    
    /// Initializes the start date and end date to be the date of the day
    @State private var selStartDate = Date()
    @State private var selEndDate = Date()
    
    var body: some View {
        VStack {
            /// CalendarView with calendars for both start date and end date
            CalendarView(selStartDate: $selStartDate, selEndDate: $selEndDate, periodFilterActive: $periodFilterActive)
                .onAppear {
                    /// Resets the calendars selected date
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

