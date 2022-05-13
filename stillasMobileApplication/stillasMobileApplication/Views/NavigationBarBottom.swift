//
//  NavigationBarBottom.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/03/2022.
//

import SwiftUI

/// **NavigationBarBottom**
/// The navigationbar View responsible for the navigation between the three main views Project, Map and Profile
struct NavigationBarBottom: View {
    /// Sets selection to project
    @State private var selection: Tab = .project
    
    /// All projects
    @State var projects = [Project]()
    
    /// Enum for the different pages
    enum Tab {
        case project
        case map
        case profile
    }
    
    var body: some View {
        TabView(selection: $selection) {
            ProjectListView()
                .tabItem {
                    Label("Prosjekt", systemImage: "square.grid.2x2")
                }
                .tag(Tab.project)
            
            MapView()
                .tabItem {
                    Label("Kart", systemImage: "map")
                }
                .tag(Tab.map)
                
            ProfileView()
                .tabItem {
                    Label("Profil", systemImage: "person.crop.circle")
                }
                .tag(Tab.profile)
        }
        .onAppear() {
            /// The transparrent effect is taken from: https://www.bigmountainstudio.com/community/public/posts/86559-how-to-customize-the-background-of-the-tabview-in-swiftui
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
            //.environmentObject(ModelData())
    }
}
