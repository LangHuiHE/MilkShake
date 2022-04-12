//
//  Scanner.swift
//  shaker
//
//  Created by LangHui He on 3/10/22.
//

import Foundation


func ProcessScannerResult(result: [String])->(String, String){
    switch result[0] {
    case "/web/loginQrCode":
        return ("Confirm Login on Website", "")
    case "/mobile/events":
        return ("Event Sign Up", result[2])
    default:
        print(result)
        return ("Can't Recognize Result", "")
    }
}

enum lazyError: Error {
    case general
}

func ExcecuteScannerResult(result: [String], completion: @escaping (Result<String,Error>)-> Void) {
    let err = lazyError.general
    let path = result[0]
    if path ==  "/web/loginQrCode" {
        let request = ConstructApi(method: "Post", endPoint: path, bodyData: nil, addAuthTokenAndDeviceInfo: true, header: ["QrToken": result[1]])
        let session = URLSession.shared
        session.dataTask(with: request){(_, response, error) in
            if error == nil, let response = response as? HTTPURLResponse {
                if response.statusCode == 200 && CatchAuthToken(response: response){
                    completion(.success("Login Success"))
                } else {
                    completion(.failure(err))
                }
            } else {
                completion(.failure(err))
            }
        }.resume()
    }
    if path == "/mobile/events" {
        let request = ConstructApi(method: "Post", endPoint: path+"/"+result[1], bodyData: nil, addAuthTokenAndDeviceInfo: true, header: [:])
        let session = URLSession.shared
        session.dataTask(with: request){(_, response, error) in
            if error == nil, let response = response as? HTTPURLResponse {
                if response.statusCode == 200 && CatchAuthToken(response: response){
                    completion(.success("Sign Up Success"))
                } else {
                    completion(.failure(err))
                }
            } else {
                completion(.failure(err))
            }
        }.resume()
    }
}
