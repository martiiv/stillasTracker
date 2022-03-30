//
//  MapView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/03/2022.
//

import SwiftUI
import UIKit
import MapKit

/**
    A MapView responsible for displaying the Apple Maps in the application.
 */
struct MapView: View {
    @StateObject private var viewModel = MapViewModel()
    @State private var searchText = ""
    
    var body: some View {
        VStack {
            GeometryReader { proxy in
                /// MapDisplay responsible for the map view
                MapDisplay()
                    .ignoresSafeArea()
                    .onAppear {
                        /// Check if locationservices are enabled when you open the map
                        viewModel.checkIfLocationServicesIsEnabled()
                    }
                ///Inspired from:
                ///https://www.hackingwithswift.com/forums/swiftui/getting-error-when-trying-to-change-location-authorisation/9216
                    .alert(isPresented: $viewModel.locationPermissionDenied, content: {
                                Alert(title: Text("Location Services Disabled"),
                                      message: Text("Please Enable Location Services For The App In App Settings For The Best Experience."),
                                      primaryButton: .default(Text("Go To Settings"),
                                                              action: {
                                                                UIApplication.shared.open(URL(string: UIApplication.openSettingsURLString)!)
                                      }),
                                      secondaryButton: .cancel(Text("Dismiss"), action: { setLocationPermissionFalse()
                                }))
                            })
            }
        }
    }
    
    /**
        Makes the app
     */
    func setLocationPermissionFalse() {
        viewModel.locationPermissionDenied = false
    }

    struct MapView_Previews: PreviewProvider {
        static var previews: some View {
            MapView()
        }
    }
}
