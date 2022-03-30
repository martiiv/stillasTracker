//
//  Project.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 30/03/2022.
//

import Foundation

struct Project: Identifiable {
    let id: Int
    let projectName: String
    let latitude, longitude: Double
    let period: Period
    let size: Int
    let state: String
    let adresse: Adresse
    let leier: Leier
    let scaffolding: Scaffolding
    let geofence: Geofence

    enum CodingKeys: String, CodingKey {
        case projectID, projectName, latitude, longitude, period, size
        case state = "State"
        case adresse
        case leier = "Leier"
        case scaffolding = "Scaffolding"
        case geofence
    }
}

// MARK: - Adresse
struct Adresse: Codable {
    let gate, postnummer, kommune, fylke: String
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

// MARK: - Leier
struct Leier: Codable {
    let name: String
    let number: Int
}

// MARK: - Period
struct Period: Codable {
    let startDate, endDate: String
}

// MARK: - Scaffolding
struct Scaffolding: Codable {
    let units: [Unit]
}

// MARK: - Unit
struct Unit: Codable {
    let type: String
    let quantity: Quantity
}

// MARK: - Quantity
struct Quantity: Codable {
    let expected, registered: Int
}

