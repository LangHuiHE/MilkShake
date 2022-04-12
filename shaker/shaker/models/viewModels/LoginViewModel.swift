//
//  LoginViewModel.swift
//  shaker
//
//  Created by LangHui He on 1/31/22.
//

import Foundation

class LoginViewModel : ObservableObject {
    @Published var credentials = LoginCredentials()
    @Published var showProgressView = false
    @Published var error: ShakerError.AuthenticationError?
    
    var loginDisabled: Bool {
        credentials.Id.isEmpty || credentials.UserPassword.isEmpty
    }
    
   func login(completion: @escaping (Bool) -> Void) {
        showProgressView = true
       DispatchQueue.main.async {
           Api.shared.login(credentials: self.credentials) {
               [unowned self] (result: Result<Bool, ShakerError.AuthenticationError>) in
               switch result {
               case .success:
                   completion(true)
               case .failure(let authError):
                   credentials = LoginCredentials()
                   error = authError
                   completion(false)
               }
               
           }
       }
       showProgressView = false
    }
}
