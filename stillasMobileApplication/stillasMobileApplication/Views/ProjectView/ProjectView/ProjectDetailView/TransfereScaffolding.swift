//
//  TransfereScaffolding.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 05/04/2022.
//

import SwiftUI
import Foundation
import Combine

// MARK: - Scaff
struct Scaff: Codable {
    let scaffold: [Move]
    let toProjectID, fromProjectID: Int
}

// MARK: - Move
struct Move: Codable {
    let type: String
    let quantity: Int
}


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
            FormView(isShowingSheet: $isShowingSheet, projects: projects, scaffolding: scaffolding)
                .navigationTitle(Text("Transfere \(scaffolding.type)"))

                /*
                Text("Transfere \(scaffolding.type) from \(projectFrom) to \(projectTo)")
                VStack(alignment: .leading) {
                    Text("From project")
                        .font(.system(size: 20, weight: .bold, design: .default))
                    
                    FormView(projects: projects)

                    
                    Section {
                        Picker("Project", selection: $projectFrom) {
                            ForEach(projects.sorted { $0.projectName < $1.projectName }, id: \.projectID) { project in
                                Text(project.projectName)
                            }
                        }
                        .border(.secondary)
                    
                    }
                    
                    Text("To project")
                        .font(.system(size: 20, weight: .bold, design: .default))
                    Section {

                        Picker("Project", selection: $projectTo) {
                            ForEach(projects.sorted { $0.projectName < $1.projectName }, id: \.projectID) { project in
                                Text(project.projectName)
                            }
                        }
                        .padding(.horizontal, 5)
                        .background(RoundedRectangle(cornerRadius: 5).stroke(Color.secondary, lineWidth: 2))
                        .foregroundColor(Color.white)
                        .contentShape(Rectangle())
                        .background(colorScheme == .dark ? Color(UIColor.white) : Color(UIColor.white)).cornerRadius(7)
                        .shadow(color: Color(UIColor.black).opacity(0.1), radius: 5, x: 0, y: 2)
                        .shadow(color: Color(UIColor.black).opacity(0.2), radius: 20, x: 0, y: 10)
                    }
                    
                    Section {
                        TextField("Amount", value: $quantity, format: .number)
                            .keyboardType(.decimalPad)
                    }

                    
                }
                .padding(.horizontal, 30)
                
                Button("Transfere scaffolding unit") {
                    Task {
                        await transfereScaffoldingUnit()
                    }
                }
                .padding()
                
                //FormView(projects: projects)*/
            }
    }
    
    func didDismiss() {
        // Handle the dismissing action.
    }
}

struct ScaffoldingDetails: View {
    var projects: [Project]

    var scaffolding: Scaffolding
    @Binding var isShowingSheet: Bool

    var body: some View {
        TransfereScaffoldingButton(projects: projects, scaffolding: scaffolding, isShowingSheet: $isShowingSheet)
        
        VStack {
            ScrollView(.vertical) {
                HStack {
                    Text("Left")
                    Divider()
                        .frame(width: 3, height: 400, alignment: .center)
                        .background(Color.gray)
                    Text("Right")
                }
            }
        }
    }
}


struct FormView: View {
    enum Field: Int, CaseIterable {
        case input
    }
    @Binding var isShowingSheet: Bool

    @FocusState var focusedField: Field?

    var projects: [Project]
    var scaffolding: Scaffolding
    @State private var pickerSelectionFrom: String = ""
    @State private var pickerSelectionTo: String = ""

    @State private var searchTerm: String = ""

    @State private var empty = false

    @State var anumber: String = ""

    var commonNumbers: [String] = ["1", "5", "10", "25", "50"]
    
    @State var selectedIndex: Int = -1

    @State private var confirmationMessage = ""
    @State private var confirmationTitle = ""
    @State private var showingConfirmation = false
    @State private var transfereMessage = ""
    @State private var transfereConfirmation: Bool = false
    @State var pickedFromID: Int = 0
    @State var pickedToID: Int = 0

    @State var tempIntFrom: Int = 0
    @State var tempIntTo: Int = 0

    @State var tem: Int = 0
    
    var filteredProjects: [Project] {
        projects.filter {
            searchTerm.isEmpty ? true : $0.projectName.lowercased().contains(searchTerm.lowercased())
        }
    }

    var body: some View {
        NavigationView {
            Form {
                Section(header: Text("From project")) {
                    Picker(selection: $pickerSelectionFrom, label: Text("From project")) {
                        SearchBar(text: $searchTerm, placeholder: "Search project")
                        .onAppear {
                            searchTerm = ""
                            anumber.removeAll()
                        }
                        .onChange(of: pickerSelectionFrom) { pick in
                            tempIntFrom = projects.firstIndex(where: { $0.projectName == pick })!
                            pickedFromID = projects[tempIntFrom].projectID
                        }
                        ForEach(filteredProjects, id: \.projectID) { project in
                            Text("\(project.projectName)").tag("\(project.projectName)")
                        }
                    }
                    .onAppear{
                        if pickerSelectionFrom.isEmpty {
                            // TODO: Fix border and error handling here
                        }
                    }
                }
                
                Section(header: Text("To project")) {
                    Picker(selection: $pickerSelectionTo, label: Text("To project")) {
                        SearchBar(text: $searchTerm, placeholder: "Search project")
                        .onAppear {
                            searchTerm = ""
                        }
                        .onChange(of: pickerSelectionTo) { pick in
                            tempIntTo = projects.firstIndex(where: { $0.projectName == pick })!
                            pickedToID = projects[tempIntTo].projectID
                        }
                        ForEach(filteredProjects, id: \.projectID) { project in
                                Text("\(project.projectName)").tag("\(project.projectName)")
                        }
                    }
                }
                
                Section (header: Text("Number of \(scaffolding.type)")){
                VStack(alignment: .leading) {
                        Text("Predefined quantity")
                        Picker("Pick a number", selection: $anumber) {
                            ForEach(commonNumbers, id: \.self) { aNumber in
                                Image(systemName: "\(aNumber).circle.fill")
                            }
                        }
                        .pickerStyle(SegmentedPickerStyle())
                    }
                    .padding()

                    HStack {
                        Text("Manual entry")
                        TextField("Input number", text: $anumber)
                            //.modifier(ClearButton(text: $anumber))
                            .onChange(of: anumber) {
                                tem = projects[tempIntFrom].scaffolding![projects[tempIntFrom].scaffolding!.firstIndex(where: { $0.type == scaffolding.type})!].quantity.expected
                                if $0.isEmpty || Int($0) ?? 0 > tem {
                                    empty = true
                                } else {
                                    empty = false
                                }
                            }
                            .textFieldStyle(TextFieldEmpty(empty: $empty))
                            .keyboardType(.numberPad)
                            .focused($focusedField, equals: .input)
                        /// https://stackoverflow.com/questions/58733003/how-to-create-textfield-that-only-accepts-numbers
                            .onReceive(Just(anumber)) { newValue in
                                            let filtered = newValue.filter { "0123456789".contains($0) }
                                            if filtered != newValue {
                                                self.anumber = filtered
                                            }
                            }
                            .toolbar {
                                ToolbarItem(placement: .keyboard) {
                                    Button("Done") {
                                        focusedField = nil
                                    }
                                }
                            }
                    }
                    .padding()
                }
            }
            .overlay(alignment: .bottom) {
                Spacer()
            
                Button(action: {
                    Task {
                        if !empty {
                            transfereMessage = "Transfere \(Int(anumber)!)x \(scaffolding.type) from \(pickerSelectionFrom) to \(pickerSelectionTo)?"
                            transfereConfirmation = true
                        }
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
            .ignoresSafeArea(.keyboard)
            .navigationTitle(Text("Transfere \(scaffolding.type)"))
            .alert(isPresented: $showingConfirmation) {
                Alert(
                    title: Text(confirmationTitle),
                    message: Text(confirmationMessage),
                    dismissButton: .default(Text("OK")) {
                        isShowingSheet = false
                        showingConfirmation = false
                    }
                )
            }
        }
        .alert(isPresented: $transfereConfirmation) {
            Alert(
                title: Text("Are you sure you want to proceed?"),
                message: Text(transfereMessage),
                primaryButton: .default(Text("Yes")) {
                    transfereConfirmation = false
                    Task {
                        if !empty {
                            await transfereScaffoldingUnit(pickedFromID: Int(exactly: pickedFromID)!, pickedToID: Int(exactly: pickedToID)!, transfereAmount: Int(anumber)!)
                        }
                    }
                    // showingConfirmation = true
                },
                secondaryButton: .cancel() {
                    transfereConfirmation = false
                }
            )
        }
    }
    
    /// https://www.appsdeveloperblog.com/http-post-request-example-in-swift/
    func transfereScaffoldingUnit(pickedFromID: Int, pickedToID: Int, transfereAmount: Int) async {
        
        let url = URL(string: "http://10.212.138.205:8080/stillastracking/v1/api/project/scaffolding")
        guard let requestUrl = url else { fatalError() }
        // Prepare URL Request Object
        var request = URLRequest(url: requestUrl)
        request.httpMethod = "PUT"
         
        
        let body : Scaff = Scaff(scaffold: [Move(type: "\(scaffolding.type.capitalizingFirstLetter())", quantity: transfereAmount)], toProjectID: pickedToID, fromProjectID: pickedFromID)
        
        print(body)
        
        guard let jsonData = try? JSONEncoder().encode(body) else {
            print("Failed to encode order")
            return
        }
        
        let jsonString = String(data: jsonData, encoding: .utf8)!
        
        request.httpBody = jsonString.data(using: String.Encoding.utf8);
        // Perform HTTP Request
        let task = URLSession.shared.dataTask(with: request) { (data, response, error) in
                
            // Check for Error
            if let error = error {
                print("Error took place \(error)")
                return
            }
     
            // Convert HTTP Response Data to a String
            if let data = data, let dataString = String(data: data, encoding: .utf8) {
                print("Response data string:\n \(dataString)")
            }
            
            if let response = response as? HTTPURLResponse {
                if response.statusCode == 200 {
                    confirmationTitle = "Success!"
                    confirmationMessage = "Successfully transfered \(body.scaffold[0].quantity)x \(body.scaffold[0].type) from project \(body.fromProjectID) to project \(body.toProjectID)!"
                } else {
                    confirmationTitle = "Failed!"
                    confirmationMessage = "Failed to transfere \(body.scaffold[0].quantity)x \(body.scaffold[0].type) from project \(body.fromProjectID) to project \(body.toProjectID)!"
                }
                showingConfirmation = true
            }
        }
        task.resume()
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
