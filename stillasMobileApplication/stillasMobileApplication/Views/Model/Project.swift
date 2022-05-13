//
//  Project.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 30/03/2022.
//

import Foundation
import SwiftUI

struct Project: Codable {
    let projectID: Int
    let projectName: String
    let size: Int
    let state: String
    let latitude, longitude: Double
    let period: Period
    let address: Address
    let customer: Customer
    let geofence: Geofence
    let scaffolding: [Scaffolding]?
}

struct Address: Codable {
    let street, zipcode, municipality, county: String
}

struct Customer: Codable {
    let name: String
    let number: Int
    let email: String
}

struct Geofence: Codable {
    let wPosition, xPosition, yPosition, zPosition: Position

    enum CodingKeys: String, CodingKey {
        case wPosition = "w-position"
        case xPosition = "x-position"
        case yPosition = "y-position"
        case zPosition = "z-position"
    }
}

struct Position: Codable {
    let latitude, longitude: Double
}

struct Period: Codable {
    let startDate, endDate: String
}

struct Scaffolding: Codable {
    let type: String
    let quantity: Quantity

    enum CodingKeys: String, CodingKey {
        case type
        case quantity = "Quantity"
    }
}

struct Quantity: Codable {
    let expected, registered: Int
}
