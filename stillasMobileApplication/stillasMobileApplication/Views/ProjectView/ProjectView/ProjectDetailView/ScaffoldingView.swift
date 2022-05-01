//
//  ScaffoldingView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 29/04/2022.
//

import SwiftUI

struct ScaffoldingView: View {
    var scaffolding: [Scaffolding]
    @Binding var isShowingSheet: Bool

    var body: some View {
        VStack {
            TransfereScaffoldingButton(isShowingSheet: $isShowingSheet)
            
            ScaffoldingItems(scaffolding: scaffolding)
        }
    }
}

struct ScaffoldingItems: View {
    var gridItemLayout = [GridItem(.flexible()), GridItem(.flexible())]
    var scaffolding: [Scaffolding]

    var body: some View {
        ScrollView (.vertical) {
            LazyVGrid (columns: gridItemLayout, spacing: 10) {
                ForEach(scaffolding, id: \.type) { scaffolding in
                    NavigationLink(destination: Text("Scaffolding detail view"), label: {
                        ScaffoldingItem(scaffolding: scaffolding)
                    })
                    .listStyle(.grouped)
                }
            }
        }
    }
}

struct ScaffoldingItem: View {
    @Environment(\.colorScheme) var colorScheme
    var scaffolding: Scaffolding
    
    var body: some View {
        VStack {
            Text(scaffolding.type).font(.title2)
                .lineLimit(1)
            
            Image("\(scaffolding.type)").resizable().scaledToFit()
            
            Spacer()

            HStack {
                VStack {
                    Text(String(format: "%d", scaffolding.quantity.expected)).foregroundColor(.black)
                        .font(.system(size: 15))
                    Text("EXPECTED")
                        .foregroundColor(.gray)
                        .font(.system(size: 10))
                }
                .frame(alignment: .center)

                VStack {
                    if (scaffolding.quantity.registered >= Int(Double(scaffolding.quantity.expected) * 0.95) && scaffolding.quantity.registered <= Int(Double(scaffolding.quantity.expected))) {
                        Text(String(format: "%d", scaffolding.quantity.registered)).foregroundColor(Color.green)
                            .font(.system(size: 15))
                    } else if ((scaffolding.quantity.registered < Int(Double(scaffolding.quantity.expected) * 0.95)) && (scaffolding.quantity.registered >= Int(Double(scaffolding.quantity.expected) * 0.8))) {
                        Text(String(format: "%d", scaffolding.quantity.registered)).foregroundColor(Color.yellow)
                            .font(.system(size: 15))
                    } else if (scaffolding.quantity.registered > Int(Double(scaffolding.quantity.expected))) {
                        Text(String(format: "%d", scaffolding.quantity.registered)).foregroundColor(Color.purple)
                            .font(.system(size: 15))
                    } else {
                        Text(String(format: "%d", scaffolding.quantity.registered)).foregroundColor(Color.red)
                            .font(.system(size: 15))
                    }
                    Text("REGISTERED")
                        .foregroundColor(.gray)
                        .font(.system(size: 10))
                }
                .frame(alignment: .center)
            }
        }
        .padding(.vertical, 5)
        .frame(width: (UIScreen.screenWidth / 2) - 40, height: (UIScreen.screenWidth / 2) - 40, alignment: .center)
        .contentShape(RoundedRectangle(cornerRadius: 5))
        .background(colorScheme == .dark ? Color(UIColor.white) : Color(UIColor.white)).cornerRadius(7)
        .shadow(color: Color(UIColor.black).opacity(0.1), radius: 5, x: 0, y: 2)
        .shadow(color: Color(UIColor.black).opacity(0.2), radius: 20, x: 0, y: 10)
    }
}

struct TransfereScaffoldingButton: View {
    @Environment(\.colorScheme) var colorScheme
    @Binding var isShowingSheet: Bool
    
    var body: some View {
        Button {
            isShowingSheet.toggle()
        } label: {
            Text("Transfere Scaffolding")
                .padding(12)
                .font(.system(size: 20))
                .foregroundColor(colorScheme == .dark ? Color(UIColor.black) : Color(UIColor.darkGray))
        }
        .contentShape(Rectangle())
        .background(colorScheme == .dark ? Color(UIColor.white) : Color(UIColor.white)).cornerRadius(7)
        .shadow(color: Color(UIColor.black).opacity(0.1), radius: 5, x: 0, y: 2)
        .shadow(color: Color(UIColor.black).opacity(0.2), radius: 20, x: 0, y: 10)
        .sheet(isPresented: $isShowingSheet,
               onDismiss: didDismiss) {
            TransfereScaffolding()
        }
    }
    
    func didDismiss() {
        // Handle the dismissing action.
    }
}

/*
struct ScaffoldingView_Previews: PreviewProvider {
    static var previews: some View {
        ScaffoldingView()
    }
}*/
