//
//  ActiveStatusFilterView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/04/2022.
//

import SwiftUI

struct ActiveStatusFilterView: View {
    @Binding var filterArr: [String]
    @Binding var projectStatus: String
    @Binding var statusFilterActive: Bool

    var body: some View {

        if statusFilterActive {
            HStack {
                HStack {
                    Text(projectStatus)
                        .padding(.leading, 5)
                        .padding(.trailing, -5)
                }
                .font(.system(size: 11).bold())
                .padding(.vertical, 5)
                
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
