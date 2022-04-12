//
//  ShakerError.swift
//  shaker
//
//  Created by LangHui He on 1/31/22.
//

import SwiftUI

class ShakerError: ObservableObject {
    @Published var isValidated = UserDefaults.standard.bool(forKey: "login")
    
    enum AuthenticationError: Error, LocalizedError, Identifiable {
        case invalidCredentials
        case invalidAuth
        
        var id:String{
            self.localizedDescription
        }
        
        var errorDescription: String? {
            switch self {
            case .invalidCredentials:
                return NSLocalizedString("Either your id or password are incorrect. Please try agian", comment: "")
            case .invalidAuth:
                return NSLocalizedString("Either your auth token is expired or your device infomation has changed. Please relogin", comment: "")
            }
        }
    }
    func updateValidation(success:Bool) {
        DispatchQueue.main.async {
            withAnimation {
                self.isValidated = success
                UserDefaults.standard.set(success, forKey: "login")
            }
        }
    }
}


func removeAllUserDefaults() {
    UserDefaults.standard.removeObject(forKey: localTokeName)
    UserDefaults.standard.removeObject(forKey: "login")
}
