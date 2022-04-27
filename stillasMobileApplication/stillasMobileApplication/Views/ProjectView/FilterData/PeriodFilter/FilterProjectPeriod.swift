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

/*
struct FilterProjectPeriod_Previews: PreviewProvider {
    @State private var selStartDate = Date()
    @State private var selEndDate = Date()
    
    static var previews: some View {
        FilterProjectPeriod(selStartDateBind: selStartDate, selEndDateBind: selEndDate)
    }
}*/

