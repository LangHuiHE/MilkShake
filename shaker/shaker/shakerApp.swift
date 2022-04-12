//
//  shakerApp.swift
//  shaker
//
//  Created by LangHui He on 1/31/22.
//

import SwiftUI

@main
struct shakerApp: App {
    @StateObject var authentication = ShakerError()
    
    var body: some Scene {
        WindowGroup {
            if authentication.isValidated {
                MainView()
                    .environmentObject(authentication)
            } else {
                LoginView()
                    .environmentObject(authentication)
                
            }
        }
    }
}
