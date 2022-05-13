//
//  TextFieldModifiers.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 09/05/2022.
//

import SwiftUI

/// **SearchBar**
/// UIViewRepresentable which adds a searchbar to a list in transfere scaffolding
/// This is taken from:
/// https://roddy.io/2020/09/07/add-search-bar-to-swiftui-picker/
struct SearchBar: UIViewRepresentable {

    @Binding var text: String
    var placeholder: String

    /// Makes a searchbar
    func makeUIView(context: UIViewRepresentableContext<SearchBar>) -> UISearchBar {
        let searchBar = UISearchBar(frame: .zero)
        searchBar.delegate = context.coordinator

        searchBar.placeholder = placeholder
        searchBar.autocapitalizationType = .none
        searchBar.searchBarStyle = .minimal
        return searchBar
    }

    /// Updates the searchbar
    func updateUIView(_ uiView: UISearchBar, context: UIViewRepresentableContext<SearchBar>) {
        uiView.text = text
    }

    func makeCoordinator() -> SearchBar.Coordinator {
        return Coordinator(text: $text)
    }

    class Coordinator: NSObject, UISearchBarDelegate {

        @Binding var text: String

        init(text: Binding<String>) {
            _text = text
        }

        func searchBar(_ searchBar: UISearchBar, textDidChange searchText: String) {
            text = searchText
        }
    }
}

// TODO: Make this happen/work - optionally remove
struct NoProjectSelected: TextFieldStyle {
    @Binding var focused: Bool
    func _body(configuration: TextField<Self._Label>) -> some View {
        configuration
        .padding(10)
        .background(
            RoundedRectangle(cornerRadius: 10, style: .continuous)
                .stroke(focused ? Color.red : Color.gray, lineWidth: 1)
        ).padding()
    }
}

/// **TextFieldEmpty**
/// Checks if the textfield is empty, if so - make the border red
/// Code taken from and inspired from:
/// https://stackoverflow.com/questions/60379010/how-to-change-swiftui-textfield-style-after-tapping-on-it
struct TextFieldEmpty: TextFieldStyle {
    @Binding var empty: Bool
    func _body(configuration: TextField<Self._Label>) -> some View {
        configuration
        .padding(10)
        .background(
            RoundedRectangle(cornerRadius: 10, style: .continuous)
                .stroke(empty ? Color.red : Color.gray, lineWidth: 1)
        ).padding()
    }
}

/// **ClearButton**
/// The clear button for the textfield
/// Code taken from and inspired from: https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&ved=2ahUKEwjzjJC2tMP3AhUnSfEDHSwjC-0QFnoECAUQAQ&url=https%3A%2F%2Fsanzaru84.medium.com%2Fswiftui-how-to-add-a-clear-button-to-a-textfield-9323c48ba61c&usg=AOvVaw1aPoAd3QYr5ByERti3mGWj
struct ClearButton: ViewModifier {
    /// The searchfield text
    @Binding var text: String

    public func body(content: Content) -> some View {
        ZStack(alignment: .trailing) {
            content
            if !text.isEmpty {
                Button(action: {
                    self.text = ""
                }) {
                    Image(systemName: "delete.left")
                        .foregroundColor(Color(UIColor.opaqueSeparator))
                }
                .padding(.trailing, 20)
            }
        }
    }
}

/// **NumbersOnly**
/// Validates textfield input to be Int
class NumbersOnly: ObservableObject {
    @Published var value = "" {
        didSet {
            let filtered = value.filter { $0.isNumber }
            
            if value != filtered {
                value = filtered
            }
        }
    }
}
