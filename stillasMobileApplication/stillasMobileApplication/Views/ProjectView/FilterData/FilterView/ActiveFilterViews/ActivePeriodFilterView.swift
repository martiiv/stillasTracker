//
//  ActivePeriodFilterView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/04/2022.
//

import SwiftUI

/// **ActivePeriodFilterView**
/// The view presented on top of the size-navigation row to display a preview of the selected periodfilter.
struct ActivePeriodFilterView: View {
    @Binding var startDate: Date
    @Binding var endDate: Date

    @Binding var filterArr: [String]
    
    @Binding var periodFilterActive: Bool

    var body: some View {

        /// if there is a filter applied, display a preview of the selected filter
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
                
                /// Deletes the selected filter and removes it from the preview
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
    
    /// Removes the filter from the array with filters
    /// - Parameter filterItem: the selected filter you want to remove
    func deleteFilterItem(filterItem: String) {
        if let i = filterArr.firstIndex(of: filterItem) {
            filterArr.remove(at: i)
        }
    }
}
/*
struct ActivePeriodFilterView_Previews: PreviewProvider {
    static var previews: some View {
        ActivePeriodFilterView()
    }
}*/
