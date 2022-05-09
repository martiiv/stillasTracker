//
//  Scaffolding.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/03/2022.
//

import Combine

// MARK: - Scaff
struct Scaff: Codable {
    let scaffold: [Move]
    let toProjectID, fromProjectID: Int
}

// MARK: - Move
struct Move: Codable {
    let type: String
    let quantity: Int
}
