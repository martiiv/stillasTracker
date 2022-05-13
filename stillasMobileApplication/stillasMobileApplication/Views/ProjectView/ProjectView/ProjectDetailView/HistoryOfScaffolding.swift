//
//  HistoryOfScaffolding.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 09/05/2022.
//

import SwiftUI

/// **HistoryOfScaffolding**
/// A View for showing the history of a scaffolding type for a project
/// As of now, it is hard coded due to missing implementation in the API, however, the design and concept is the same.
struct HistoryOfScaffolding: View {
    /// All projects
    var projects: [Project]
    
    /// Darkmode or lightmode activated?
    @Environment(\.colorScheme) var colorScheme

    /// Scaffolding type
    var scaffolding: Scaffolding
    
    /// Transfere scaffolding Modal View showing
    @Binding var isShowingSheet: Bool
    
    /// Used to show dates in history
    let dateNow = Date()

    var body: some View {
        ScrollView(.vertical) {
            VStack (alignment: .leading){
                HStack {
                    VStack {
                        /// Each "blob" with history info
                        RoundedRectangle(cornerRadius: 8, style: .continuous)
                        .fill(colorScheme == .dark ? Color(UIColor.white) : Color(UIColor.white)).cornerRadius(7)
                        .frame(width: UIScreen.screenWidth / 2.3 , height: 60)
                        .shadow(color: Color(UIColor.black).opacity(0.1), radius: 5, x: -2, y: 2)
                        .shadow(color: Color(UIColor.black).opacity(0.2), radius: 20, x: -10, y: 10)
                        .overlay(VStack {
                            /// Preview data for scaffolding on the project on a given date
                            Text(Date(), style: .date)
                                .foregroundColor(Color.gray)
                            HStack {
                                VStack {
                                    Text(String(format: "%d", scaffolding.quantity.expected))
                                        .font(.system(size: 15))
                                    Text("FORVENTET")
                                        .foregroundColor(.gray)
                                        .font(.system(size: 10))
                                }
                                VStack {
                                    /// Changes color of registered depending on how close it is to be equal to expected
                                    amountOfScaffoldingRegistered(expected: scaffolding.quantity.expected, registered: scaffolding.quantity.registered)
                                    Text("REGISTRERT")
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
                            /// Sets date to go back in time for each history "blob"
                            Text(Calendar.current.date(byAdding: .day, value: -2, to: dateNow)!, style: .date)
                                .foregroundColor(Color.gray)
                            HStack {
                                VStack {
                                    Text(String(format: "%d", scaffolding.quantity.expected))
                                        .font(.system(size: 15))
                                    Text("FORVENTET")
                                        .foregroundColor(.gray)
                                        .font(.system(size: 10))
                                }
                                VStack {
                                    amountOfScaffoldingRegistered(expected: scaffolding.quantity.expected, registered: scaffolding.quantity.registered)
                                    Text("REGISTRERT")
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
                                    Text("FORVENTET")
                                        .foregroundColor(.gray)
                                        .font(.system(size: 10))
                                }
                                VStack {
                                    amountOfScaffoldingRegistered(expected: scaffolding.quantity.expected, registered: scaffolding.quantity.registered)
                                    Text("REGISTRERT")
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
                                    Text("FORVENTET")
                                        .foregroundColor(.gray)
                                        .font(.system(size: 10))
                                }
                                VStack {
                                    amountOfScaffoldingRegistered(expected: scaffolding.quantity.expected, registered: scaffolding.quantity.registered)
                                        .font(.system(size: 15))
                                    Text("REGISTRERT")
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
                                    Text("FORVENTET")
                                        .foregroundColor(.gray)
                                        .font(.system(size: 10))
                                }
                                VStack {
                                    amountOfScaffoldingRegistered(expected: scaffolding.quantity.expected, registered: scaffolding.quantity.registered)
                                        .font(.system(size: 15))
                                    Text("REGISTRERT")
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
                                    Text("FORVENTET")
                                        .foregroundColor(.gray)
                                        .font(.system(size: 10))
                                }
                                VStack {
                                    amountOfScaffoldingRegistered(expected: scaffolding.quantity.expected, registered: scaffolding.quantity.registered)
                                        .font(.system(size: 15))
                                    Text("REGISTRERT")
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
        .navigationTitle(Text("Historie for \(scaffolding.type)".capitalizingFirstLetter()))
    }
}

/*
struct HistoryOfScaffolding_Previews: PreviewProvider {
    static var previews: some View {
        HistoryOfScaffolding()
    }
}*/
