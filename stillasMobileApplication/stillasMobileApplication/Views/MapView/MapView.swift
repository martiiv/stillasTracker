//
//  MapView.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 24/03/2022.
//

import SwiftUI
import MapKit

/**
    A MapView responsible for displaying the Apple Maps in the application.
 
    Inspiration taken from this youtube video:
    https://www.youtube.com/watch?v=CyMtjSspJZA
 */
struct MapView: View {
    @State var searchText = ""
    @State var offset: CGFloat = 0
    @State var lastOffset: CGFloat = 0
    
    @GestureState var gestureOffset: CGFloat = 0
    
    @State private var isInitialOffsetSet = false
    
    var body: some View {
    ZStack {
        GeometryReader { proxy in
            /// MapViewMap responsible for the map view
            MapViewMap()
            
            /// DrawerView responsible for the drawer slide
            DrawerView()
        }
    }
}
    
    struct DrawerView: View {
        @State var searchText = ""
        @State var offset: CGFloat = 0
        @State var lastOffset: CGFloat = 0
        @GestureState var gestureOffset: CGFloat = 0
        @State private var isInitialOffsetSet = false
        
        var body: some View {
            GeometryReader { proxy in
                let height = proxy.frame(in: .global).height
                    ZStack {
                        BlurView(style:
                                .systemMaterial)
                        .clipShape(CustomCorners(corners: [.topLeft, .topRight], radius: 30))
                        
                        VStack {
                            Capsule()
                                .fill(Color.gray)
                                .frame(width: 40, height: 5)
                                .padding(.top, 7)
                            
                            TextField("Search field", text: $searchText)
                                .padding(.vertical, 7)
                                .padding(.horizontal)
                                .background(BlurView(style: .dark))
                                .cornerRadius(10)
                                .colorScheme(.dark)
                                .padding(.top, 5)
                        }
                        .padding(.horizontal)
                        .frame(maxHeight: .infinity, alignment: .top)
                        
                    }
                    .offset(y: height - 100)
                    .offset(y: -offset > 0 ? -offset <= (height - 100) ? offset : -(height - 100) : 0)
                    .gesture(DragGesture().updating($gestureOffset, body: {value, out, _ in
                        out = value.translation.height
                        /// onChangeDrawer() updates the offset when a gesture was performed
                        onChangeDrawer()
                    }).onEnded({ value in
                        let maxHeight = height - 100
                        /// When the gesture ends, update the placement of the drawer view to fixed position
                        withAnimation {
                            if -offset > 100 && -offset < maxHeight / 2 {
                                offset = -(maxHeight / 3)
                            }
                            else if (-offset > maxHeight / 2) {
                                 offset = -maxHeight
                            }
                            else {
                                offset = 0
                            }
                        }
                        lastOffset = offset
                    }))
            }
            .ignoresSafeArea(.all, edges: .bottom)
        }
        
        /**
            onChangeDrawer() resposible for updating the offset when a gesture is performed
         */
        func onChangeDrawer (){
            DispatchQueue.main.async {
                self.offset = gestureOffset + lastOffset
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
