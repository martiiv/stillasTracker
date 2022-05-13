//
//  Scaffolding.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/03/2022.
//

import Combine

struct Scaff: Codable {
    let scaffold: [Move]
    let toProjectID, fromProjectID: Int
}

struct Move: Codable {
    let type: String
    let quantity: Int
}
