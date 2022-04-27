//
//  CalendarView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 27/04/2022.
//

import SwiftUI

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
struct CalendarView_Previews: PreviewProvider {
    static var previews: some View {
        CalendarView()
    }
}*/ 
