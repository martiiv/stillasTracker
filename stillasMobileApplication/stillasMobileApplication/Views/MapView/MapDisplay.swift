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
    MapDisplay - makes a MKMapView and defines its properties like userTrackingMode and setRegion etc.
 
    This way to define a map was inspired by this resource, as it creates a MKMapView in a neat way. It also makes the process of displaying CheckPoints and GeoFences easier alongside with other map functionality.
    https://iosapptemplates.com/blog/swiftui/map-view-swiftui
 */
struct MapDisplay: UIViewRepresentable {
    /// A property wrapper type that instantiates an observable object of type MapViewModel()
    @StateObject private var viewModel = MapViewModel()

    /**
        makeUIView() - Makes the MKMapView
            Allows to show user location, sets tracking mode and region of interest on "open"
     */
    func makeUIView(context: Context) -> MKMapView {
        let mapView = MKMapView(frame: UIScreen.main.bounds)
        mapView.showsUserLocation = true
        mapView.userTrackingMode = .follow
        mapView.setRegion(viewModel.region, animated: true)
        updateUIView(mapView)
        return mapView
    }
    
    func updateUIView(_ uiView: MKMapView) {
        let projVi = ProjectView()
        
        let annotations = projVi.projectsArr.map { project -> MKAnnotation in
            let annotation = MKPointAnnotation()
            annotation.title = project.projectName
            annotation.subtitle = "\(project.id)"
            annotation.coordinate = CLLocationCoordinate2D(latitude: project.latitude, longitude: project.longitude)
            return annotation
        }
        uiView.addAnnotations(annotations)
    }

    /**
        updateUIView() - Updates the state of the MKMapView with the changed information from SwiftUI
     */
    func updateUIView(_ uiView: MKMapView, context: Context) {
    }
}

