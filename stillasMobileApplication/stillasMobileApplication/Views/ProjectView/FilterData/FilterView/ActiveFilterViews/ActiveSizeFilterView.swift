//
//  ActiveSizeFilterView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/04/2022.
//

import SwiftUI

struct ActiveSizeFilterView: View {
    @Binding var filterArr: [String]
    @Binding var projectMinSize: Int
    @Binding var projectMaxSize: Int
    @Binding var sizeFilterActive: Bool

    var body: some View {

        if sizeFilterActive {
            HStack {
                HStack {
                    ScrollView (.horizontal, showsIndicators: false) {
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
                    }
                }
                .font(.system(size: 11).bold())
                .padding(.vertical, 5)
                .lineLimit(1)
                
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
