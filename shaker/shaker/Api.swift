//
//  Api.swift
//  shaker
//
//  Created by LangHui He on 2/7/22.
//

import SwiftUI

func ConstructURL (endPoint : String) -> URL {
    return URL(string: prefix + host + port + endPoint)!
}

func CatchAuthToken (response: HTTPURLResponse)->Bool{
    if let token = response.value(forHTTPHeaderField: ApiTokenKey) {
        // print(token)
        if token.isEmpty {
            print("There is no token in the header")
        } else {
            SwapToken(token: token)
        }
        return true
    }
    return false
}

func ConstructApi(method : String, endPoint: String, bodyData : Data?, addAuthTokenAndDeviceInfo: Bool, header : [String:String]?)-> (URLRequest){
    let url = ConstructURL(endPoint: endPoint)
    var request = URLRequest(url: url)
    request.httpMethod = method
    if bodyData?.isEmpty == false{
        request.addValue("application/json", forHTTPHeaderField: "Content-Type")
        request.httpBody = bodyData
    }
    if addAuthTokenAndDeviceInfo {
        let deviceInfo = DeviceInfo()
        request.addValue(deviceInfo.Name, forHTTPHeaderField: "DeviceName")
        request.addValue(deviceInfo.Model, forHTTPHeaderField: "DeviceModel")
        request.addValue(deviceInfo.UUID, forHTTPHeaderField: "DeviceUUID")
        request.addValue(GetToken(), forHTTPHeaderField: ApiTokenKey)
    }
    
    for (k, v) in header ?? [:] {
        request.addValue(v, forHTTPHeaderField: k)
    }
    
    return request
}

class Api {
    static let shared = Api()
    
    func login (credentials: LoginCredentials, completion: @escaping (Result<Bool, ShakerError.AuthenticationError>) -> Void) {
        DispatchQueue.main.async {
            let intID = Int(credentials.Id) ?? 0
            if intID == 0 {
                completion(.failure(.invalidCredentials))
                return
            }
            let body = ["Id": intID,
                        "UserPassword":credentials.UserPassword
            ] as [String : Any]
            let bodyData = try? JSONSerialization.data(
                withJSONObject: body,
                options: []
            )
            
            let request = ConstructApi(method: "Post", endPoint: "/mobile/login", bodyData: bodyData, addAuthTokenAndDeviceInfo: true, header: nil)
            
            let session = URLSession.shared

            let task = session.dataTask(with: request) {(data, response, error) in
                
                if error == nil, let response = response as? HTTPURLResponse {
                    if response.statusCode == 200 && CatchAuthToken(response: response){
                            UserDefaults.standard.set(credentials.Id, forKey: "userID")
                            completion(.success(true))
                    } else {
                        completion(.failure(.invalidCredentials))
                    }
            } else if let error = error {
                    // Handle HTTP request error
                    print("This is error")
                    print(error)
                    completion(.failure(.invalidCredentials))
                } else {
                    // Handle unexpected error
                    print("unexpected error")
                    completion(.failure(.invalidCredentials))
                }
            }
            task.resume()
        }
    }
    
    
    func getUserInfo (completion: @escaping (Result<UserInfo, ShakerError.AuthenticationError>) -> Void) {
        // DispatchQueue.main.async {
            
        let request = ConstructApi(method: "Get", endPoint: "/mobile/user", bodyData: nil, addAuthTokenAndDeviceInfo: true, header: nil)
            
            let session = URLSession.shared

            let task = session.dataTask(with: request) {(data, response, error) in
                
                if error == nil, let data = data, let response = response as? HTTPURLResponse {
                    if response.statusCode == 200 && CatchAuthToken(response: response){
                        let userInfo = try! JSONDecoder().decode(UserInfo.self, from: data)
                        //  inprint(userInfo)
                        // completion(.success(true))
                        DispatchQueue.main.async {
                            completion(.success(userInfo))
                        }
                    } else {
                        completion(.failure(.invalidAuth))
                    }
            } else if let error = error {
                    // Handle HTTP request error
                    print("This is error")
                    print(error)
                    completion(.failure(.invalidAuth))
                } else {
                    // Handle unexpected error
                    print("unexpected error")
                    completion(.failure(.invalidAuth))
                }
            }
            task.resume()
        }
    // }
    
    func getAllAccounts (completion: @escaping (Result<Accounts, ShakerError.AuthenticationError>) -> Void) {
        let request = ConstructApi(method: "Get", endPoint: "/mobile/accounts", bodyData: nil, addAuthTokenAndDeviceInfo: true, header: nil)
        let session = URLSession.shared
        let task = session.dataTask(with: request) {(data, response, error) in
            
            if error == nil, let data = data, let response = response as? HTTPURLResponse {
                if response.statusCode == 200 && CatchAuthToken(response: response){
                    let accounts = try? JSONDecoder().decode(Accounts.self, from: data)
                    //  inprint(userInfo)
                    // completion(.success(true))
                    
                    // TODO:
                    // need to test if the body doesn't match what next?
                    DispatchQueue.main.async {
                        completion(.success(accounts ?? Accounts()))
                    }
                } else {
                    completion(.failure(.invalidAuth))
                }
        } else if let error = error {
                // Handle HTTP request error
                print("This is error")
                print(error)
                completion(.failure(.invalidAuth))
            } else {
                // Handle unexpected error
                print("unexpected error")
                completion(.failure(.invalidAuth))
            }
        }
        task.resume()
    }
}
