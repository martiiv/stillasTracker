//
//  MapViewModel.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 29/03/2022.
//

import Foundation
import MapKit
import SwiftUI

/// **MapDetails**
/// Enum values that are used multiple places.
/// Used to abstract reused code.
enum MapDetails {
    /// Sets the starting location of the map to be Gjøvik (latitude: 60.79574, longitude: 10.69155)
    static let startingLocation = CLLocationCoordinate2D(
        latitude: 60.79574,
        longitude: 10.69155
    )
    /// Sets zoom level of map on initialization
    /// Closer to zero is more zoomed in
    static let defaultSpan = MKCoordinateSpan(
        latitudeDelta: 0.03,
        longitudeDelta: 0.03
    )
}

/// **MapViewModel**
/// Class responsible for checking if the user has enabled location services.
/// This class is inspired a lot by the Apple Development documentation as well as this youtube video:
/// https://www.youtube.com/watch?v=hWMkimzIQoU
final class MapViewModel: NSObject, ObservableObject, CLLocationManagerDelegate {
    @Published var locationPermissionDenied = false
    @Published var dismissCount = 0
    // TODO: use this region to update on region change for instance in project
    
    /// A changable variable 'region' which is responsible for the maps start location
    @Published var region = MKCoordinateRegion (
        center: MapDetails.startingLocation,
        span: MapDetails.defaultSpan
    )
    
    var locationManager: CLLocationManager?
    
    /// checkIfLocationServicesIsEnabled() - Checks if the user has enabled location services.
    func checkIfLocationServicesIsEnabled() {
        if CLLocationManager.locationServicesEnabled() {
            locationManager = CLLocationManager()
            locationManager?.delegate = self
            locationManager?.desiredAccuracy = kCLLocationAccuracyBest
        } else {
            locationPermissionDenied = true
            //throw LocationError.denied
            // TODO: Add error thingy where you tell user to turn it on
        }
    }
    
    // TODO: Check -> fatal error when location services are off?
    /// checkLocationAuthorization
    /// Checks which authorization the application is assigned to.
    private func checkLocationAutorization() {
        guard let locationManager = locationManager else { return }
  
            switch locationManager.authorizationStatus {
                case .notDetermined:
                    locationManager.requestWhenInUseAuthorization()
                
                case .restricted:
                    locationPermissionDenied = true
                    print("Restricted location")
                // TODO: Add error thingy location restricted
                
                case .denied:
                    locationPermissionDenied = true
                    print("Denied location")
                
                // TODO: Add error thingy location denied

                case .authorizedAlways, .authorizedWhenInUse:
                    region = MKCoordinateRegion(center: locationManager.location!.coordinate,
                                                span: MapDetails.defaultSpan)
                @unknown default:
                    break
            }
    }
    
    /// locationManagerDidChangeAuthorization() - Checks the autorization on locationManager creation as well as if the apps authorization changes
    func locationManagerDidChangeAuthorization(_ manager: CLLocationManager) {
        checkLocationAutorization()
    }
}
