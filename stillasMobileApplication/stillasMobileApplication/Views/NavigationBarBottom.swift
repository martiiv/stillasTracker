//
//  NavigationBarBottom.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/03/2022.
//

import SwiftUI

struct NavigationBarBottom: View {
    @State private var selection: Tab = .map
    
    enum Tab {
        case project
        case map
        case profile
        case projectViewAPI
    }
    
    
    var body: some View {
        
        TabView(selection: $selection) {
            ProjectView()
                .tabItem {
                    Label("Project", systemImage: "square.grid.2x2")
                }
                .tag(Tab.project)
            
            MapView()
                .tabItem {
                    Label("Map", systemImage: "map")
                }
                .tag(Tab.map)
                
            ProfileView()
                .tabItem {
                    Label("Profile", systemImage: "person.crop.circle")
                }
                .tag(Tab.profile)
            /*ProjectViewN()
                .tabItem {
                    Label("Project API", systemImage: "map")
                }*/
        }
        .onAppear() {
            /// https://www.bigmountainstudio.com/community/public/posts/86559-how-to-customize-the-background-of-the-tabview-in-swiftui
            let appearance = UITabBarAppearance()
                        appearance.backgroundEffect = UIBlurEffect(style: .systemThinMaterial)
                        // Use this appearance when scrolling behind the TabView:
                        UITabBar.appearance().standardAppearance = appearance
                        // Use this appearance when scrolled all the way up:
                        UITabBar.appearance().scrollEdgeAppearance = appearance
        }
    }
}

struct NavigationBarBottom_Previews: PreviewProvider {
    static var previews: some View {
        NavigationBarBottom()
            .environmentObject(ModelData())
    }
}
