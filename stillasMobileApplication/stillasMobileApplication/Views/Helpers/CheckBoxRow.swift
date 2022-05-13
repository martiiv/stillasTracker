//
//  CheckBoxList.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 27/04/2022.
//

import SwiftUI

/// **CheckBoxRow**
/// Creates a checkbox row to be used in lists etc.
struct CheckBoxRow: View {
    var title: String
    @Binding var selectedItems: Set<String>
    @State var isSelected: Bool
    
    var body: some View {
        GeometryReader { geometry in
            HStack {
                CheckBoxView(checked: $isSelected, title: title)
                    .onChange(of: isSelected) { _ in
                        if isSelected {
                            selectedItems.insert(title)
                            print(selectedItems)
                        } else {
                            selectedItems.remove(title)// or
                        }
                    }
            }
        }
    }
}

/// **CheckBoxView**
/// Creates the view of the checkboxes
struct CheckBoxView: View {
    @Binding var checked: Bool
    @State var title: String
    
    var body: some View {
        HStack {
            Image(systemName: checked ? "checkmark.square.fill" : "square")
                .foregroundColor(checked ? Color(UIColor.systemBlue) : Color.secondary)
            Text(title)
                .padding(.leading)
        }
        .onTapGesture {
            self.checked.toggle()
        }
    }
}

/*
struct CheckBoxList_Previews: PreviewProvider {
    static var previews: some View {
        CheckBoxList()
    }
}*/
