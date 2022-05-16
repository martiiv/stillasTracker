//
//  ActiveStatusFilterView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/04/2022.
//

import SwiftUI

/// **ActiveStatusFilterView**
/// The view presented on top of the size-navigation row to display a preview of the selected statusfilter.
struct ActiveStatusFilterView: View {
    @Binding var filterArr: [String]
    @Binding var projectStatus: String
    @Binding var statusFilterActive: Bool

    var body: some View {
        /// if there is a filter applied, display a preview of the selected filter
        if statusFilterActive {
            HStack {
                HStack {
                    Text(projectStatus)
                        .padding(.leading, 5)
                        .padding(.trailing, -5)
                }
                .font(.system(size: 11).bold())
                .padding(.vertical, 5)
                
                /// Deletes the selected filter and removes it from the preview
                Button(action: {
                    deleteFilterItem(filterItem: "status")
                    self.statusFilterActive.toggle()
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
struct ActiveStatusFilterView_Previews: PreviewProvider {
    static var previews: some View {
        ActiveStatusFilterView()
    }
}
*/
