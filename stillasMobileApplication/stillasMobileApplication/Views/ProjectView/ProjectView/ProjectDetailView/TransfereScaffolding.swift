//
//  TransfereScaffolding.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 05/04/2022.
//

import SwiftUI
import Foundation

struct TransfereScaffolding: View {
    var projects: [Project]
    @Environment(\.colorScheme) var colorScheme

    var scaffolding: Scaffolding
    @Binding var isShowingSheet: Bool
    
    @State private var quantity: Int = 1
    @State private var name: String = "Tim"
    @State private var projectFrom: String = ""
    @State private var projectTo: String = ""

    
    var body: some View {
        VStack {
            TransfereScaffoldingView(isShowingSheet: $isShowingSheet, projects: projects, scaffolding: scaffolding)
                .navigationTitle(Text("Transfere \(scaffolding.type)"))
            }
    }
    
    func didDismiss() {
        // Handle the dismissing action.
    }
}

func amountOfScaffoldingRegistered(expected: Int, registered: Int) -> Text {
    if (registered >= Int(Double(expected) * 0.95) && registered <= Int(Double(expected))) {
        return Text(String(format: "%d", registered)).foregroundColor(Color.green)
            .font(.system(size: 15))
    } else if ((registered < Int(Double(expected) * 0.95)) && (registered >= Int(Double(expected) * 0.8))) {
        return Text(String(format: "%d", registered)).foregroundColor(Color.yellow)
            .font(.system(size: 15))
    } else if (registered > Int(Double(expected))) {
        return Text(String(format: "%d", registered)).foregroundColor(Color.purple)
            .font(.system(size: 15))
    } else {
        return Text(String(format: "%d", registered)).foregroundColor(Color.red)
            .font(.system(size: 15))
    }
}

/// https://roddy.io/2020/09/07/add-search-bar-to-swiftui-picker/
struct SearchBar: UIViewRepresentable {

    @Binding var text: String
    var placeholder: String

    func makeUIView(context: UIViewRepresentableContext<SearchBar>) -> UISearchBar {
        let searchBar = UISearchBar(frame: .zero)
        searchBar.delegate = context.coordinator

        searchBar.placeholder = placeholder
        searchBar.autocapitalizationType = .none
        searchBar.searchBarStyle = .minimal
        return searchBar
    }

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

/// https://stackoverflow.com/questions/60379010/how-to-change-swiftui-textfield-style-after-tapping-on-it
///
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

// https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&ved=2ahUKEwjzjJC2tMP3AhUnSfEDHSwjC-0QFnoECAUQAQ&url=https%3A%2F%2Fsanzaru84.medium.com%2Fswiftui-how-to-add-a-clear-button-to-a-textfield-9323c48ba61c&usg=AOvVaw1aPoAd3QYr5ByERti3mGWj
struct ClearButton: ViewModifier
{
    @Binding var text: String

    public func body(content: Content) -> some View
    {
        ZStack(alignment: .trailing)
        {
            content
            if !text.isEmpty
            {
                Button(action:
                {
                    self.text = ""
                })
                {
                    Image(systemName: "delete.left")
                        .foregroundColor(Color(UIColor.opaqueSeparator))
                }
                .padding(.trailing, 20)
            }
        }
    }
}

/*
struct TransfereScaffolding_Previews: PreviewProvider {
    static var previews: some View {
        TransfereScaffolding()
    }
}*/
