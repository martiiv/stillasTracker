//
//  TransfereScaffoldingView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 09/05/2022.
//

import SwiftUI
import Combine

struct TransfereScaffoldingView: View {
    enum Field: Int, CaseIterable {
        case input
    }
    @Binding var isShowingSheet: Bool

    @FocusState var focusedField: Field?

    var projects: [Project]
    var scaffolding: Scaffolding
    @State private var pickerSelectionFrom: String = ""
    @State private var pickerSelectionTo: String = ""

    @State private var pickerSelectionFromNotPicked: Bool = true
    
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
                            pickerSelectionFromNotPicked = false
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

/*
struct TransfereScaffoldingView_Previews: PreviewProvider {
    static var previews: some View {
        TransfereScaffoldingView()
    }
}*/
