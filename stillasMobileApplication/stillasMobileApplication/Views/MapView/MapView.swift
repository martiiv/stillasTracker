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
 
    Inspiration taken from this youtube video:
    https://www.youtube.com/watch?v=CyMtjSspJZA
 */
struct MapView: View {
    @State private var searchText = ""
    
    var body: some View {
        ZStack {
            GeometryReader { proxy in
                /// MapViewMap responsible for the map view
                MapViewMap()
            }
        }
    }
    /**
        MapViewMap creates a view containing the map
     */
    struct MapViewMap: View {
        /// Sets the starting location to be Gjøvik (latitude: 60.79574, longitude: 10.69155)
        @State var region = MKCoordinateRegion (
            center: CLLocationCoordinate2D(
                latitude: 60.79574,
                longitude: 10.69155
            ),
            /// The zoom level of the application when opened
            /// Closer to 0 means greater zoom level
            span: MKCoordinateSpan(
                latitudeDelta: 0.03,
                longitudeDelta: 0.03
            )
        )
        var body: some View {
            Map(coordinateRegion: $region)
                .ignoresSafeArea()
        }
    }

    struct MapView_Previews: PreviewProvider {
        static var previews: some View {
            MapView()
        }
    }
}
