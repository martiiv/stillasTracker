//
//  CacheEntry.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 07/04/2022.
//

import Foundation
import UIKit

final class CacheEntry <V> {
    let key: String
    let value: V
    let expiredTimestamp: Date
    
    init(key: String, value: V, expiredTimestamp: Date) {
        self.key = key
        self.value = value
        self.expiredTimestamp = expiredTimestamp
    }
    
    func isCacheExpired(after date: Date = .now) -> Bool {
        date > expiredTimestamp
    }
}
