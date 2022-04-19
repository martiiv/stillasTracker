//
//  FilterProjectData.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 14/04/2022.
//

import UIKit
import SwiftUI

enum FilterType {
    case none,
         period,
         startBeforePeriod,
         startAfterPeriod,
         endBeforePeriod,
         endAfterPeriod,
         sizeEqualTo,
         sizeLessThan,
         sizeGreaterThan,
         state,
         county
}

struct FilterProjectData: View {
    @State var projects = [Project]()
    @State private var showFilterModalView: Bool = false
    @State private var showAddProjectModalView: Bool = false
    
    let filter: FilterType
    
    // TODO: Make these values operable
    @State var projectStartDate = ""
    @State var projectEndDate = ""
    @State var projectSize = ""
    @State var projectState = ""
    @State var projectCounty = ""
    
    var body: some View {
        VStack {
            NavigationView {
                Form {
                    Section(header: Text("All Projects")) {
                        List(filteredProjects, id: \.projectID) { project in
                            Text(project.projectName)
                        }
                        .navigationTitle("Projects")
                        //.listStyle(.grouped)
                    }
                }
                .listStyle(.grouped)
                .toolbar {
                    ToolbarItemGroup(placement: .navigationBarLeading) {
                        Button(action: {
                            print("Filter tapped!")
                            self.showFilterModalView.toggle()
                            
                        }) {
                            Label("Filter", systemImage: "line.3.horizontal.decrease.circle")
                        }
                    }
                    
                    ToolbarItemGroup(placement: .navigationBarTrailing) {
                        Button(action: {
                            print("Add project tapped!")
                            self.showAddProjectModalView.toggle()
                        }) {
                            Label("Add", systemImage: "plus.circle")
                        }
                    }
                }
                .sheet(isPresented: $showFilterModalView,
                       onDismiss: didDismiss) {
                    FilterView()
                }
               .sheet(isPresented: $showAddProjectModalView, onDismiss: didDismiss) {
                   AddProjectView()
               }
            }
        }
        .task {
            await ProjectData().loadData { (projects) in
                 self.projects = projects
            }
        }
    }
    
    func didDismiss() {
        
        // Handle the dismissing action.
    }
    
    var filteredProjects: [Project] {
        switch filter {
        case .none:
            return projects
        case .period:
            return projects.filter { $0.period.startDate > projectStartDate && $0.period.endDate < projectEndDate }
        case .startBeforePeriod:
            return projects.filter { $0.period.startDate < projectStartDate }
        case .startAfterPeriod:
            return projects.filter { $0.period.startDate > projectStartDate }
        case .endBeforePeriod:
            return projects.filter { $0.period.endDate < projectEndDate }
        case .endAfterPeriod:
            return projects.filter { $0.period.endDate > projectEndDate }
        case .sizeEqualTo:
            return projects.filter { $0.size == Int(projectSize) }
        case .sizeLessThan:
            return projects.filter { $0.size < Int(projectSize) ?? 0 }
        case .sizeGreaterThan:
            return projects.filter { $0.size > Int(projectSize) ?? 0 }
        case .state:
            return projects.filter { $0.state == projectState }
        case .county:
            return projects.filter { $0.address.county == projectCounty }
        }
    }
}

struct FilterView: View {
    @State private var filterItems = ["Område", "Prosjekt periode", "Størrelse", "Status"]
        
    var body: some View {
        NavigationView {
            List {
                ForEach(filterItems, id: \.self) { filterItem in
                    NavigationLink {
                        switch filterItem {
                        case "Område":
                            FilterProjectArea()
                        /*case "Prosjekt periode":
                            print("Add period view")
                            // ADD period view*/
                        case "Størrelse":
                            IntSlider()
                            //print("Add size view")
                            // ADD size view
                            /*
                        case "Status":
                            print("Add status view")
                            // ADD status view
                        default:
                            print("Did not find any")
                        */
                        default:
                            AddProjectView()
                        }
                    } label: {
                        Text(filterItem)
                    }
                }
            }
            .navigationTitle(Text("Filter"))
            .navigationViewStyle(StackNavigationViewStyle())
        }
    }
}


struct IntSlider: View {
    private enum Field: Int, CaseIterable {
            case input
        }
    
    @ObservedObject var input = NumbersOnly()
    
    @State var score: Int = 0
    @FocusState private var focusedField: Field?

    var intProxy: Binding<Double>{
        Binding<Double>(
            get: {
            //returns the score as a Double
                return Double(score)
        }, set: {
            //rounds the double to an Int
            print($0.description)
            score = Int($0)
            input.value = "\(Int($0))"
        })
    }
    
    var body: some View {
        VStack{
            /*if(input.value != "\(0)" && Int(input.value) != score) {
                input.value = "\(score)"
                AddProjectView()
            } else {
                AddProjectView()
            }*/
            
            VStack {
                TextField("Input", text: $input.value)
                    .onChange(of: input.value) { value in
                        score = Int(value) ?? 0
                    }
                    .padding()
                    .keyboardType(.numberPad)
                    .focused($focusedField, equals: .input)

                Text(score.description)
                Text("Størrelse")
            }
            .toolbar {
                ToolbarItem(placement: .keyboard) {
                    Button("Done") {
                        focusedField = nil
                    }
                }
            }
            .font(.headline)
            .font(Font.system(size: 60, design: .default))
            
            Slider(value: intProxy , in: 100.0...1000.0, step: 50.0, onEditingChanged: {_ in
                print(score.description)
            })
            .frame(width: 350, alignment: .center)
        }
    }
}

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

struct AddProjectView: View {
    var body: some View {
        VStack {
            Text("Add Project SheetView")
        }
    }
}

struct FilterProjectData_Previews: PreviewProvider {
    static var previews: some View {
        FilterProjectData(filter: .none)
    }
}
