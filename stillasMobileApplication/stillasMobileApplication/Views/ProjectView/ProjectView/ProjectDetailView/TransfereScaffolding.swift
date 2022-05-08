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
            }
    }
    
    func didDismiss() {
        // Handle the dismissing action.
    }
}

struct ScaffoldingDetails: View {
    var projects: [Project]
    @Environment(\.colorScheme) var colorScheme

    var scaffolding: Scaffolding
    @Binding var isShowingSheet: Bool
    
    let dateNow = Date()

    var body: some View {
        
        Divider()
            .padding(.vertical, 10)
        
        TransfereScaffoldingButton(projects: projects, scaffolding: scaffolding, isShowingSheet: $isShowingSheet)
        
        Divider()
            .padding(.vertical, 10)
        
        Text("History of \(scaffolding.type)".capitalizingFirstLetter())
            .font(Font.title.bold())
            .frame(alignment: .leading)
        
        ScrollView(.vertical) {
            VStack (alignment: .leading){
                HStack {
                    VStack {
                        RoundedRectangle(cornerRadius: 8, style: .continuous)
                        .fill(colorScheme == .dark ? Color(UIColor.white) : Color(UIColor.white)).cornerRadius(7)
                        .frame(width: UIScreen.screenWidth / 2.3 , height: 60)
                        .shadow(color: Color(UIColor.black).opacity(0.1), radius: 5, x: -2, y: 2)
                        .shadow(color: Color(UIColor.black).opacity(0.2), radius: 20, x: -10, y: 10)
                        .overlay(VStack {
                            Text(Date(), style: .date)
                                .foregroundColor(Color.gray)
                            HStack {
                                VStack {
                                    Text(String(format: "%d", scaffolding.quantity.expected))
                                        .font(.system(size: 15))
                                    Text("EXPECTED")
                                        .foregroundColor(.gray)
                                        .font(.system(size: 10))
                                }
                                VStack {
                                    amountOfScaffoldingRegistered(expected: scaffolding.quantity.expected, registered: scaffolding.quantity.registered)
                                    Text("REGISTERED")
                                        .foregroundColor(.gray)
                                        .font(.system(size: 10))
                                }
                            }
                        })
                        .padding(.leading)
                        .padding(.vertical, 35)
                        .ignoresSafeArea()
                        
                        RoundedRectangle(cornerRadius: 8, style: .continuous)
                        .fill(colorScheme == .dark ? Color(UIColor.white) : Color(UIColor.white)).cornerRadius(7)
                        .frame(width: UIScreen.screenWidth / 2.3 , height: 60)
                        .shadow(color: Color(UIColor.black).opacity(0.1), radius: 5, x: -2, y: 2)
                        .shadow(color: Color(UIColor.black).opacity(0.2), radius: 20, x: -10, y: 10)
                        .overlay(VStack {
                            Text(Calendar.current.date(byAdding: .day, value: -2, to: dateNow)!, style: .date)
                                .foregroundColor(Color.gray)
                            HStack {
                                VStack {
                                    Text(String(format: "%d", scaffolding.quantity.expected))
                                        .font(.system(size: 15))
                                    Text("EXPECTED")
                                        .foregroundColor(.gray)
                                        .font(.system(size: 10))
                                }
                                VStack {
                                    amountOfScaffoldingRegistered(expected: scaffolding.quantity.expected, registered: scaffolding.quantity.registered)
                                    Text("REGISTERED")
                                        .foregroundColor(.gray)
                                        .font(.system(size: 10))
                                }
                            }
                        })
                        .padding(.leading)
                        .padding(.vertical, 35)
                        .ignoresSafeArea()
                        
                        RoundedRectangle(cornerRadius: 8, style: .continuous)
                        .fill(colorScheme == .dark ? Color(UIColor.white) : Color(UIColor.white)).cornerRadius(7)
                        .frame(width: UIScreen.screenWidth / 2.3 , height: 60)
                        .shadow(color: Color(UIColor.black).opacity(0.1), radius: 5, x: -2, y: 2)
                        .shadow(color: Color(UIColor.black).opacity(0.2), radius: 20, x: -10, y: 10)
                        .overlay(VStack {
                            Text(Calendar.current.date(byAdding: .day, value: -4, to: dateNow)!, style: .date)
                                .foregroundColor(Color.gray)
                            HStack {
                                VStack {
                                    Text(String(format: "%d", scaffolding.quantity.expected))
                                        .font(.system(size: 15))
                                    Text("EXPECTED")
                                        .foregroundColor(.gray)
                                        .font(.system(size: 10))
                                }
                                VStack {
                                    amountOfScaffoldingRegistered(expected: scaffolding.quantity.expected, registered: scaffolding.quantity.registered)
                                    Text("REGISTERED")
                                        .foregroundColor(.gray)
                                        .font(.system(size: 10))
                                }
                            }
                        })
                        .padding(.leading)
                        .padding(.vertical, 35)
                        .ignoresSafeArea()
                    }
                    
                    Divider()
                        .frame(width: 3, height: 400, alignment: .center)
                        .background(Color.gray)
                        .padding(.top)
                        .offset(y: 30)
                    
                    VStack {
                        RoundedRectangle(cornerRadius: 8, style: .continuous)
                        .fill(colorScheme == .dark ? Color(UIColor.white) : Color(UIColor.white)).cornerRadius(7)
                        .frame(width: UIScreen.screenWidth / 2.3 , height: 60)
                        .shadow(color: Color(UIColor.black).opacity(0.1), radius: 5, x: 2, y: 2)
                        .shadow(color: Color(UIColor.black).opacity(0.2), radius: 20, x: 10, y: 10)
                        .overlay(VStack {
                            Text(Calendar.current.date(byAdding: .day, value: -1, to: dateNow)!, style: .date)
                                .foregroundColor(Color.gray)
                            HStack {
                                VStack {
                                    Text(String(format: "%d", scaffolding.quantity.expected))
                                        .font(.system(size: 15))
                                    Text("EXPECTED")
                                        .foregroundColor(.gray)
                                        .font(.system(size: 10))
                                }
                                VStack {
                                    amountOfScaffoldingRegistered(expected: scaffolding.quantity.expected, registered: scaffolding.quantity.registered)
                                        .font(.system(size: 15))
                                    Text("REGISTERED")
                                        .foregroundColor(.gray)
                                        .font(.system(size: 10))
                                }
                            }
                        })
                        .offset(y: 70)
                        .padding(.trailing)
                        .padding(.vertical, 35)
                        .ignoresSafeArea()
                        
                        RoundedRectangle(cornerRadius: 8, style: .continuous)
                        .fill(colorScheme == .dark ? Color(UIColor.white) : Color(UIColor.white)).cornerRadius(7)
                        .frame(width: UIScreen.screenWidth / 2.3 , height: 60)
                        .shadow(color: Color(UIColor.black).opacity(0.1), radius: 5, x: 2, y: 2)
                        .shadow(color: Color(UIColor.black).opacity(0.2), radius: 20, x: 10, y: 10)
                        .overlay(VStack {
                            Text(Calendar.current.date(byAdding: .day, value: -3, to: dateNow)!, style: .date)
                                .foregroundColor(Color.gray)
                            HStack {
                                VStack {
                                    Text(String(format: "%d", scaffolding.quantity.expected))
                                        .font(.system(size: 15))
                                    Text("EXPECTED")
                                        .foregroundColor(.gray)
                                        .font(.system(size: 10))
                                }
                                VStack {
                                    amountOfScaffoldingRegistered(expected: scaffolding.quantity.expected, registered: scaffolding.quantity.registered)
                                        .font(.system(size: 15))
                                    Text("REGISTERED")
                                        .foregroundColor(.gray)
                                        .font(.system(size: 10))
                                }
                            }
                        })
                        .offset(y: 70)
                        .padding(.trailing)
                        .padding(.vertical, 35)
                        .ignoresSafeArea()
                        
                        RoundedRectangle(cornerRadius: 8, style: .continuous)
                        .fill(colorScheme == .dark ? Color(UIColor.white) : Color(UIColor.white)).cornerRadius(7)
                        .frame(width: UIScreen.screenWidth / 2.3 , height: 60)
                        .shadow(color: Color(UIColor.black).opacity(0.1), radius: 5, x: 2, y: 2)
                        .shadow(color: Color(UIColor.black).opacity(0.2), radius: 20, x: 10, y: 10)
                        .overlay(VStack {
                            Text(Calendar.current.date(byAdding: .day, value: -5, to: dateNow)!, style: .date)
                                .foregroundColor(Color.gray)
                            HStack {
                                VStack {
                                    Text(String(format: "%d", scaffolding.quantity.expected))
                                        .font(.system(size: 15))
                                    Text("EXPECTED")
                                        .foregroundColor(.gray)
                                        .font(.system(size: 10))
                                }
                                VStack {
                                    amountOfScaffoldingRegistered(expected: scaffolding.quantity.expected, registered: scaffolding.quantity.registered)
                                        .font(.system(size: 15))
                                    Text("REGISTERED")
                                        .foregroundColor(.gray)
                                        .font(.system(size: 10))
                                }
                            }
                        })
                        .offset(y: 70)
                        .padding(.trailing)
                        .padding(.vertical, 35)
                        .ignoresSafeArea()
                    }
                }
            }
        }
        .navigationTitle(Text("\(scaffolding.type)".capitalizingFirstLetter()))
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
