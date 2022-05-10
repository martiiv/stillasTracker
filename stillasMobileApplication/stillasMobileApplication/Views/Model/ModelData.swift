//
//  ModelData.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 28/03/2022.
//

import Foundation
import Combine

/**
    ModelData responsible for loading/decoding the user information from the data/json object
    To update views when data changes, you make your data model classes observable objects.
 */
final class ModelData: ObservableObject {
    @Published var users: [User] = load("userData.json")
}

/// load<T: Decodable>() - function for decoding the data into json object
/// Uses guard and do/catch methods to make sure to catch potential errors while retreiving or decoding the data
func load<T: Decodable>(_ filename: String) -> T {
    let data: Data

    guard let file = Bundle.main.url(forResource: filename, withExtension: nil)
    else {
        fatalError("Couldn't find \(filename) in main bundle.")
    }

    do {
        data = try Data(contentsOf: file)
    } catch {
        fatalError("Couldn't load \(filename) from main bundle:\n\(error)")
    }

    do {
        let decoder = JSONDecoder()
        return try decoder.decode(T.self, from: data)
    } catch {
        fatalError("Couldn't parse \(filename) as \(T.self):\n\(error)")
    }
}
