//
//  ActiveAreaFilterView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/04/2022.
//

import SwiftUI


/// **ActiveAreaFilterView**
/// The view presented on top of the area-navigation row to display a preview of the selected areafilter.
struct ActiveAreaFilterView: View {
    @Binding var filterArr: [String]
    
    @Binding var areaFilterActive: Bool

    var body: some View {

        /// if there is a filter applied, display a preview of the selected filter
        if areaFilterActive {
            HStack {
                HStack {
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
                }
                .font(.system(size: 11).bold())
                .padding(.vertical, 5)
                .lineLimit(1)
                
                /// Deletes the selected filter and removes it from the preview
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
    
    /// Removes the filter from the array with filters
    /// - Parameter filterItem: the selected filter you want to remove
    func deleteFilterItem(filterItem: String) {
        if let i = filterArr.firstIndex(of: filterItem) {
            filterArr.remove(at: i)
        }
    }
}

/*
struct ActiveAreaFilterView_Previews: PreviewProvider {
    static var previews: some View {
        ActiveAreaFilterView()
    }
}
*/
