//
//  MapDisplay.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 29/03/2022.
//

import SwiftUI
import MapKit
import CoreLocation


/**
 https://iosapptemplates.com/blog/swiftui/map-view-swiftui
 */
struct MapDisplay: UIViewRepresentable {
    @StateObject private var viewModel = MapViewModel()

    func makeUIView(context: Context) -> MKMapView {
        let mapView = MKMapView(frame: UIScreen.main.bounds)
        mapView.showsUserLocation = true
        mapView.userTrackingMode = .follow
        mapView.setRegion(viewModel.region, animated: true)
        return mapView
    }

    func updateUIView(_ uiView: MKMapView, context: Context) {
    }
}
