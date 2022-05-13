//
//  ActiveSizeFilterView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/04/2022.
//

import SwiftUI

/// **ActiveSizeFilterView**
/// The view presented on top of the size-navigation row to display a preview of the selected sizefilter.
struct ActiveSizeFilterView: View {
    @Binding var filterArr: [String]
    @Binding var projectMinSize: Int
    @Binding var projectMaxSize: Int
    @Binding var sizeFilterActive: Bool
    @Binding var selection: String

    var body: some View {
        /// if there is a filter applied, display a preview of the selected filter
        if sizeFilterActive {
            HStack {
                HStack {
                    ScrollView (.horizontal, showsIndicators: false) {
                        if selection == "Between" { /// display both minimum size and max size
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
                        } else if selection == "Less Than" { /// display only minimum size
                            HStack {
                                Text("Under ") +
                                Text("\(projectMinSize) m")
                                + Text("2")
                                    .baselineOffset(6)
                                    .font(Font.system(size: 10))
                            }
                            .padding(.leading, 5)
                        } else if selection == "Greater Than" { /// display only maximum size
                            HStack {
                                Text("Over ") +
                                Text("\(projectMaxSize) m")
                                + Text("2")
                                    .baselineOffset(6)
                                    .font(Font.system(size: 10))
                            }
                            .padding(.leading, 5)
                        }
                    }
                }
                .font(.system(size: 11).bold())
                .padding(.vertical, 5)
                .lineLimit(1)
                
                /// Deletes the selected filter and removes it from the preview
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
    
    
    /// Removes the filter from the array with filters
    /// - Parameter filterItem: the selected filter you want to remove
    func deleteFilterItem(filterItem: String) {
        if let i = filterArr.firstIndex(of: filterItem) {
            filterArr.remove(at: i)
        }
    }
}

/*
struct ActiveSizeFilterView_Previews: PreviewProvider {
    static var previews: some View {
        ActiveSizeFilterView()
    }
}*/
