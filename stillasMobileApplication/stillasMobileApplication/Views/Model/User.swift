//
//  User.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/03/2022.
//

import SwiftUI
import Foundation

struct User: Codable, Hashable, Identifiable {
    var id: Int
    var name: String
    var dateOfBirth: String
    var role: String
    var admin: Bool

    private var imageName: String
    var image: Image {
        Image(imageName)
    }
    
    static var formatter = LengthFormatter()
}
