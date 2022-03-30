//
//  MapViewModel.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 29/03/2022.
//

import Foundation
import MapKit
import SwiftUI

enum LocationError: Error {
    case denied
    case restricted
}

enum MapDetails {
    static let startingLocation = CLLocationCoordinate2D(
        latitude: 60.79574,
        longitude: 10.69155
    )
    static let defaultSpan = MKCoordinateSpan(
        latitudeDelta: 0.03,
        longitudeDelta: 0.03
    )
}

/**
 https://www.youtube.com/watch?v=hWMkimzIQoU
 */
final class MapViewModel: NSObject, ObservableObject, CLLocationManagerDelegate {
    @Published var locationPermissionDenied = false
    @Published var dismissCount = 0
    // TODO: use this region to update on region change for instance in project
    /// Sets the starting location to be Gjøvik (latitude: 60.79574, longitude: 10.69155)
    @Published var region = MKCoordinateRegion (
        center: MapDetails.startingLocation,
        /// The zoom level of the application when opened
        /// Closer to 0 means greater zoom level
        span: MapDetails.defaultSpan
    )
    
    var locationManager: CLLocationManager?
    
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
    
    /// Checks the autorization on locationManager creation as well as if the apps authorization changes
    func locationManagerDidChangeAuthorization(_ manager: CLLocationManager) {
        checkLocationAutorization()
    }
}
