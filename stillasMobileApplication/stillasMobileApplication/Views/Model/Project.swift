//
//  Project.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 30/03/2022.
//

import Foundation
import SwiftUI

/*
struct Response: Codable{
    var results: [Project]
}
 */

// MARK: - Project
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

// MARK: - Address
struct Address: Codable {
    let street, zipcode, municipality, county: String
}

// MARK: - Customer
struct Customer: Codable {
    let name: String
    let number: Int
    let email: String
}

// MARK: - Geofence
struct Geofence: Codable {
    let wPosition, xPosition, yPosition, zPosition: Position

    enum CodingKeys: String, CodingKey {
        case wPosition = "w-position"
        case xPosition = "x-position"
        case yPosition = "y-position"
        case zPosition = "z-position"
    }
}

// MARK: - Position
struct Position: Codable {
    let latitude, longitude: Double
}

// MARK: - Period
struct Period: Codable {
    let startDate, endDate: String
}

// MARK: - Scaffolding
struct Scaffolding: Codable {
    let type: String
    let quantity: Quantity

    enum CodingKeys: String, CodingKey {
        case type
        case quantity = "Quantity"
    }
}

// MARK: - Quantity
struct Quantity: Codable {
    let expected, registered: Int
}
