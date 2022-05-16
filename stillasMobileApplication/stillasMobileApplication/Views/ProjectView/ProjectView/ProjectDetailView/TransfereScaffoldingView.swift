//
//  TransfereScaffoldingView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 09/05/2022.
//

import SwiftUI
import Combine

/// **TransfereScaffoldingView**
/// The View for transfering scaffolding
struct TransfereScaffoldingView: View {
    /// Input field
    enum Field: Int, CaseIterable {
        case input
    }
    
    /// Is Modal View active
    @Binding var isShowingSheet: Bool

    /// Field is selected --> Focus
    @FocusState var focusedField: Field?

    /// All projects
    var projects: [Project]
    
    /// Scaffolding type
    var scaffolding: Scaffolding
    
    /// Value of project picker selection
    @State private var pickerSelectionFrom: String = ""
    @State private var pickerSelectionTo: String = ""

    /// Checks if "from project" is not selected yet
    @State private var pickerSelectionFromNotPicked: Bool = true
    
    /// Search for projects
    @State private var searchTerm: String = ""

    /// Is amount of scaffolding empty?
    @State private var empty = false

    /// The number of scaffolding (as string because of textfield)
    @State var anumber: String = ""

    /// Template numbers to be picked from
    var commonNumbers: [String] = ["1", "5", "10", "25", "50"]
    
    @State var selectedIndex: Int = -1
    @State private var confirmationMessage = ""
    @State private var confirmationTitle = ""
    @State private var showingConfirmation = false
    @State private var transfereMessage = ""
    @State private var transfereConfirmation: Bool = false
    
    /// ProjectID's
    @State var pickedFromID: Int = 0
    @State var pickedToID: Int = 0

    @State var tempIntFrom: Int = 0
    @State var tempIntTo: Int = 0

    /// Checks the expected count of the selected project
    @State var tem: Int = 0
    
    /// The filtered projects
    var filteredProjects: [Project] {
        projects.filter {
            searchTerm.isEmpty ? true : $0.projectName.lowercased().contains(searchTerm.lowercased())
        }
    }

    var body: some View {
        NavigationView {
            Form {
                Section(header: Text("Fra prosjekt")) {
                    /// Picker for the project to transfere scaffolding from
                    Picker(selection: $pickerSelectionFrom, label: Text("Fra prosjekt")) {
                        SearchBar(text: $searchTerm, placeholder: "Søk etter prosjekt")
                        .onAppear {
                            searchTerm = ""
                            /// Reset number of scaffolding
                            anumber.removeAll()
                        }
                        .onChange(of: pickerSelectionFrom) { pick in
                            tempIntFrom = projects.firstIndex(where: { $0.projectName == pick })!
                            pickedFromID = projects[tempIntFrom].projectID
                            pickerSelectionFromNotPicked = false
                        }
                        /// List all projects
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
                
                Section(header: Text("Til prosjekt")) {
                    /// Picker for the project to transfere scaffolding to
                    Picker(selection: $pickerSelectionTo, label: Text("Til prosjekt")) {
                        SearchBar(text: $searchTerm, placeholder: "Søk etter prosjekt")
                        .onAppear {
                            searchTerm = ""
                        }
                        .onChange(of: pickerSelectionTo) { pick in
                            tempIntTo = projects.firstIndex(where: { $0.projectName == pick })!
                            pickedToID = projects[tempIntTo].projectID
                        }
                        /// List all projects
                        ForEach(filteredProjects, id: \.projectID) { project in
                                Text("\(project.projectName)").tag("\(project.projectName)")
                        }
                    }
                }
                
                Section (header: Text("Antall \(scaffolding.type)")){
                VStack(alignment: .leading) {
                    Text("Forhåndsdefinert mengde")
                    /// Picker for the predefined scaffolding amount
                    Picker("Pick a number", selection: $anumber) {
                        ForEach(commonNumbers, id: \.self) { aNumber in
                            Image(systemName: "\(aNumber).circle.fill")
                        }
                    }
                    .pickerStyle(SegmentedPickerStyle())
                }
                .padding()

                HStack {
                    Text("Manuell innfylling")
                    TextField("Input", text: $anumber)
                        .disabled(pickerSelectionFromNotPicked)
                        //.modifier(ClearButton(text: $anumber))
                        .onChange(of: anumber) {
                            tem = (projects[tempIntFrom].scaffolding![projects[tempIntFrom].scaffolding!.firstIndex(where: { $0.type == scaffolding.type}) ?? 0].quantity.expected)
                            
                            if $0.isEmpty || Int($0) ?? -1 > tem {
                                empty = true
                            } else {
                                empty = false
                            }
                        }
                        .textFieldStyle(TextFieldEmpty(empty: $empty))
                        .keyboardType(.numberPad)
                        .focused($focusedField, equals: .input)
                    /// Code taken from:
                    /// https://stackoverflow.com/questions/58733003/how-to-create-textfield-that-only-accepts-numbers
                    /// to make sure input is of type Int and update picker
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
                /// Button to start transaction for transfering scaffolding
                Button(action: {
                    Task {
                        if !empty {
                            /// Starts transfere process if alert is confirmed
                            transfereMessage = "Overfør \(Int(anumber)!)x \(scaffolding.type) fra \(pickerSelectionFrom) til \(pickerSelectionTo)?"
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
            .navigationTitle(Text("Overfør \(scaffolding.type)"))
            .alert(isPresented: $showingConfirmation) {
                /// Transfere was successful alert
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
            /// Checks if the user wants to proceed. If yes, execute request to the API
            Alert(
                title: Text("Er du sikker på at du vil fortsette?"),
                message: Text(transfereMessage),
                primaryButton: .default(Text("Ja")) {
                    transfereConfirmation = false
                    Task {
                        if !empty {
                            /// Starts request to transfere scaffoling
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
    
    /// Establishes connection to the API and sendts a PUT request to the API with the scaffolding amount from which project to which project.
    /// The code is partially taken and inspired from:
    /// https://www.appsdeveloperblog.com/http-post-request-example-in-swift/
    /// - Parameters:
    ///   - pickedFromID: The ID of the project to transfere scaffolding from
    ///   - pickedToID: The ID of the project to transfere scaffolding to
    ///   - transfereAmount: The amount of scaffolding to transfere
    func transfereScaffoldingUnit(pickedFromID: Int, pickedToID: Int, transfereAmount: Int) async {
        let url = URL(string: "http://10.212.138.205:8080/stillastracking/v1/api/project/scaffolding")
        guard let requestUrl = url else { fatalError() }
        // Prepare URL Request Object
        var request = URLRequest(url: requestUrl)
        request.httpMethod = "PUT"
         
        /// Sets the body to be of type Scaff
        let body : Scaff = Scaff(scaffold: [Move(type: "\(scaffolding.type.capitalizingFirstLetter())", quantity: transfereAmount)], toProjectID: pickedToID, fromProjectID: pickedFromID)
        
        /// Tries to encode body
        guard let jsonData = try? JSONEncoder().encode(body) else {
            print("Failed to encode order")
            return
        }
        
        /// Makes a jsonString of the encoded data
        let jsonString = String(data: jsonData, encoding: .utf8)!
        
        /// Creates the request body with the jsonStrings data
        request.httpBody = jsonString.data(using: String.Encoding.utf8);
        
        /// Perform HTTP Request
        let task = URLSession.shared.dataTask(with: request) { (data, response, error) in
                
            /// Check for Error
            if let error = error {
                print("Error took place \(error)")
                return
            }
     
            /// Convert HTTP Response Data to a String
            if let data = data, let dataString = String(data: data, encoding: .utf8) {
                print("Response data string:\n \(dataString)")
            }
            
            /// Checks the response's statuscode and displays the alert based on the code
            if let response = response as? HTTPURLResponse {
                if response.statusCode == 200 {
                    confirmationTitle = "Suksess!"
                    confirmationMessage = "Overføringen av \(body.scaffold[0].quantity)x \(body.scaffold[0].type) fra prosjekt \(body.fromProjectID) til prosjekt \(body.toProjectID) var en suksess!"
                } else {
                    confirmationTitle = "Mislykket!"
                    confirmationMessage = "Overføringen av \(body.scaffold[0].quantity)x \(body.scaffold[0].type) fra prosjekt \(body.fromProjectID) til prosjekt \(body.toProjectID) feilet!"
                }
                showingConfirmation = true
            }
        }
        task.resume()
    }
}

/*
struct TransfereScaffoldingView_Previews: PreviewProvider {
    static var previews: some View {
        TransfereScaffoldingView()
    }
}*/
