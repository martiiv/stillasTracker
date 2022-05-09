//
//  HistoryOfScaffolding.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 09/05/2022.
//

import SwiftUI

struct HistoryOfScaffolding: View {
    var projects: [Project]
    @Environment(\.colorScheme) var colorScheme

    var scaffolding: Scaffolding
    @Binding var isShowingSheet: Bool
    
    let dateNow = Date()

    var body: some View {
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
        .navigationTitle(Text("History of \(scaffolding.type)".capitalizingFirstLetter()))
    }
}

/*
struct HistoryOfScaffolding_Previews: PreviewProvider {
    static var previews: some View {
        HistoryOfScaffolding()
    }
}*/
