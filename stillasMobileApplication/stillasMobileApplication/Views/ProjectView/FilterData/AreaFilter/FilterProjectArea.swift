//
//  FilterProjectArea.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 19/04/2022.
//

import SwiftUI

struct FilterProjectArea: View {
    @State private var checked: [Bool]
    @Binding var selArr: [String]
    @Binding var areaFilterActive: Bool

    let counties = ["Agder", "Innlandet", "Møre og Romsdal", "Nordland", "Oslo", "Rogaland", "Vestfold og Telemark", "Troms og Finnmark", "Trøndelag", "Vestlandet", "Viken"]

    // selectedItems gets updated by the CheckBoxRow as it changes
    @State var selectedItems: Set<String> = [] // Use a Set to keep track of multiple check boxes
    
    init(selArr: Binding<[String]>, areaFilterActive: Binding<Bool>) {
        self._selArr = selArr
        _checked = State(initialValue: [Bool](repeating: false, count: counties.count))
        self._areaFilterActive = areaFilterActive
    }
    
    var body: some View {
        VStack {
            VStack {
                List {
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
