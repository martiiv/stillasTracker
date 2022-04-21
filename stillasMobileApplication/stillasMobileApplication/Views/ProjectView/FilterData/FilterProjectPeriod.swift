//
//  FilterProjectPeriod.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 20/04/2022.
//

import SwiftUI

struct FilterProjectPeriod: View {
    //@State private var date = Date()
    
    var body: some View {
        VStack {
            CalendarView()
        }
        .navigationTitle(Text("Prosjekt periode"))
        //.ignoresSafeArea(edges: .top)
    }
}

struct CalendarView: View {
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
                            //.background(RoundedRectangle(cornerRadius: 4.0).stroke(Color.blue).padding(-3))
                        //.datePickerStyle(.graphical)
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
            Button(action: { print("Bruk") }) {
                Text("Bruk")
                    .frame(width: 300, height: 50, alignment: .center)
            }
            .foregroundColor(.white)
            //.padding(.vertical, 10)
            .background(Color.blue)
            .cornerRadius(10)
            
            Spacer()
                .frame(height:50)  // limit spacer size by applying a frame
        }
    }
}

struct FilterProjectPeriod_Previews: PreviewProvider {
    static var previews: some View {
        FilterProjectPeriod()
    }
}
