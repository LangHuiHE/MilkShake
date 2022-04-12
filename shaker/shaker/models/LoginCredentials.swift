//
//  Credentials.swift
//  shaker
//
//  Created by LangHui He on 1/31/22.
//

import Foundation

class LoginCredentials:Codable {
    var Id : String = UserDefaults.standard.string(forKey: "userID") ?? ""
    var UserPassword : String = "milkshake"
}
