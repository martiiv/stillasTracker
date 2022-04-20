//
//  CacheImplementation.swift
//  stillasMobileApplication
//
//  Created by Tormod Mork Muller on 07/04/2022.
//

import Foundation

actor InMemoryCache<V> {
    private let cache: NSCache<NSString, CacheEntry<V>> = .init()
    private let expirationInterval: TimeInterval
    
    init(expirationInterval: TimeInterval) {
        self.expirationInterval = expirationInterval
    }
    
    func removeValue(forKey key: String) {
        cache.removeObject(forKey: key as NSString)
    }
    
    func removeAllValues() {
        cache.removeAllObjects()
    }
    
    func setValue(_ value: V?, forKey key: String) {
        if let value = value {
            let expiredTimestamp = Date().addingTimeInterval(expirationInterval)
            let cacheEntry = CacheEntry(key: key, value: value, expiredTimestamp: expiredTimestamp)
            cache.setObject(cacheEntry, forKey: key as NSString)
        } else {
            removeValue(forKey: key)
        }
    }
    
    func value(forKey key: String) -> V? {
        guard  let entry = cache.object(forKey: key as NSString) else {
            return nil
        }
        
        guard !entry.isCacheExpired(after: Date()) else {
            removeValue(forKey: key)
            return nil
        }
        
        return entry.value
    }
}
