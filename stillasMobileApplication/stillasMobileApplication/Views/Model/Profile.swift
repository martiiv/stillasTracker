//
//  Profile.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 10/05/2022.
//

import Foundation

struct Profile: Codable {
    let employeeID: String
    var name: Name
    let dateOfBirth, role: String
    let phone: Int
    let email: String
    let admin: Bool
}

struct Name: Codable {
    let firstName, lastName: String
}
