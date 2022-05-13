//
//  MapView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/03/2022.
//

import SwiftUI
import UIKit
import MapKit

/// **MapView**
/// Is responsible for displaying the Apple Maps in the application.
struct MapView: View {
    /// A property wrapper type that instantiates an observable object of type MapViewModel()
    @StateObject private var viewModel = MapViewModel()
    @State private var searchText = ""
    @State private var dismissedAlready = false
    //@State var projects = [Project]()

    var body: some View {
        VStack {
            GeometryReader { proxy in
                /// MapDisplay responsible for the map view
                MapDisplay()
                    .ignoresSafeArea()
                    .onAppear {
                        viewModel.checkIfLocationServicesIsEnabled()
                    }
                
                /// Displays an alert to the user if the location services are disabled, recommending the user to enable them and suggesting a redirect to the location service settings for the application.
                /// This alert is inspired from: https://www.hackingwithswift.com/forums/swiftui/getting-error-when-trying-to-change-location-authorisation/9216
                    .alert(isPresented: $viewModel.locationPermissionDenied,
                           content: {
                                Alert(title: Text("Location Services Disabled"),
                                      message: Text("Please Enable Location Services For The App In App Settings For The Best Experience."),
                                      primaryButton: .default(Text("Go To Settings"),
                                                              action: {
                                                                UIApplication.shared.open(URL(string: UIApplication.openSettingsURLString)!)
                                      }),
                                      secondaryButton: .cancel(Text("Dismiss"), action: { setLocationPermissionFalse()
                                }))
                            })
                    /*.task {
                        await ProjectData().loadData { (projects) in
                             self.projects = projects
                        }
                    }*/
            }
        }

    }
    
    
    /// Dismisses the alert
    func setLocationPermissionFalse() {
        viewModel.locationPermissionDenied = false
    }
}

struct MapView_Previews: PreviewProvider {
    static var previews: some View {
        MapView()
    }
}
