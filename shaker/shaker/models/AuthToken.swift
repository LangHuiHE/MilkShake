//
//  AuthToken.swift
//  shaker
//
//  Created by LangHui He on 2/5/22.
//

import Foundation

func SwapToken(token : String) {
    UserDefaults.standard.removeObject(forKey: localTokeName)
    SetToken(token: token)
}

func GetToken() -> String {
    return UserDefaults.standard.string(forKey: localTokeName) ?? ""
}

func SetToken (token : String){
    UserDefaults.standard.set(token, forKey: localTokeName)
}

func RemoveToken() {
    UserDefaults.standard.removeObject(forKey: localTokeName)
}

