//
//  FilterProjectArea.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 19/04/2022.
//

import SwiftUI

/// **FilterProjectArea**
/// The View for selecting a project with area set to a value
struct FilterProjectArea: View {
    
    /// Is the box checked?
    @State private var checked: [Bool]
    
    /// All the selected counties
    @Binding var selArr: [String]
    
    /// Area filter active
    @Binding var areaFilterActive: Bool

    /// All counties to filter based of
    let counties = ["Agder", "Innlandet", "Møre og Romsdal", "Nordland", "Oslo", "Rogaland", "Vestfold og Telemark", "Troms og Finnmark", "Trøndelag", "Vestlandet", "Viken"]

    /// selectedItems gets updated by the CheckBoxRow as it changes
    @State var selectedItems: Set<String> = [] /// Use a Set to keep track of multiple check boxes
    
    /// Initializes the selections to false so the boxes are unchecked as the user accesses the filter
    init(selArr: Binding<[String]>, areaFilterActive: Binding<Bool>) {
        self._selArr = selArr
        _checked = State(initialValue: [Bool](repeating: false, count: counties.count))
        self._areaFilterActive = areaFilterActive
    }
    
    var body: some View {
        VStack {
            VStack {
                List {
                    /// For each county, add it to the list with a checkbox and description
                    ForEach(counties, id: \.self) { county in
                        HStack {
                            CheckBoxRow(title: county, selectedItems: $selectedItems, isSelected: selectedItems.contains(county))
                                .padding(.top)
                                .padding(.bottom)
                        }
                    }
                }
                .navigationTitle(Text("Område"))
                }
            .padding(.bottom, 110)
            }
        .overlay(alignment: .bottom) {
            /// Updates the parent View with the selected counties
            Button(action: {
                print(self.selectedItems)
                for selectedItem in selectedItems {
                    if(!selArr.contains(selectedItem)) {
                        selArr.append(selectedItem)
                    }
                }
                if !selArr.isEmpty {
                    areaFilterActive = true
                } else {
                    areaFilterActive = false
                }
            }) {
                Text("Bruk")
                    .frame(width: 300, height: 50, alignment: .center)
            }
            .foregroundColor(.white)
            .background(Color.blue)
            .cornerRadius(10)
            .padding(.bottom, 50)
        }
    }
}
/*
struct FilterProjectArea_Previews: PreviewProvider {
    static var previews: some View {
        FilterProjectArea()
    }
}*/
