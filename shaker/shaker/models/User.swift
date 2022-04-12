//
//  User.swift
//  shaker
//
//  Created by LangHui He on 1/31/22.
//

import SwiftUI

struct Student : Codable {
    var Credits: Int
    var Department : String
    var Major : String
    
    init() {
        Credits = 0
        Department = ""
        Major = ""
    }
}

struct UserInfo : Codable{
    var ContactAddress  : String
    var EmailAddress    : String
    var EnrollDate      : String
    var FirstName       : String
    var LastName        : String
    var PermanentAddress: String
    var PhoneNumber     : String
    var Id              : Int
    var UserType        : Int
    var StudentInfo     : Student
    
    init(){
        ContactAddress = ""
        EmailAddress = ""
        EnrollDate   = ""
        FirstName  = ""
        LastName = ""
        PermanentAddress = ""
        PhoneNumber = ""
        Id = 0
        UserType = 0
        StudentInfo = Student()
    }
}

class User : ObservableObject {
    @Published var userInfo = UserInfo()
    @Published var error: ShakerError.AuthenticationError?
    @Published var userAccounts = Accounts()
    
    init(){
        getUserInfo()
    }
    
    func getUserInfo(){
        Api.shared.getUserInfo() {
            [unowned self] (result: Result<UserInfo, ShakerError.AuthenticationError>) in
            switch result {
            case .success (let info):
                userInfo = info
            case .failure(let authError):
                error = authError
            }
        }
    }
    
    
    func getUserTypeName() -> String {
        switch userInfo.UserType {
        case 1:
            return "Admin"
        case 2:
            return "Faculty"
        case 3:
            return "Student"
        case 4:
            return "Visitor"
        default:
            return "Trailblazer"
        }
    }
    
    func getClassYear()->String {
        let y = Calendar.current.component(.year, from: Date())
        return String((Int(String(userInfo.EnrollDate.prefix(4))) ?? Int(y)) + 4)
    }
    
    func getFullID()->String{
        var full = String(userInfo.Id)
        while full.count < 8{
            full = "0" + full
        }
        full = "D" + full
        return full
    }
    
    func loadAccounts(){
        Api.shared.getAllAccounts() {
            [unowned self] (result: Result<Accounts, ShakerError.AuthenticationError>) in
            switch result {
            case.success(let data):
                userAccounts = data
                // print(userAccounts)
            case .failure(let authError):
                error = authError
            }
        }
    }
    
    func calculateAllBalance() -> [String : Float] {
        var balances: [String : Float] = ["Tuition":0, "Meal Plan":0, "PrintPoint":0]
        
        // Tuition
        var sum : Float = 0
        for ay in userAccounts.Tuition {
            for ad in ay.Dates {
                for payment in ad.Payments {
                    sum += payment.Amount
                }
            }
        }
        balances["Tuition"] = sum
        
        // Meal Plan
        sum = 0
        for ay in userAccounts.MealPlan {
            for ad in ay.Dates {
                for payment in ad.Payments {
                    sum += payment.Amount
                }
            }
        }
        balances["Meal Plan"] = sum
        
        // PrintPoint
        sum = 0
        for ay in userAccounts.PrintPoint{
            for ad in ay.Dates {
                for payment in ad.Payments {
                    sum += payment.Amount
                }
            }
        }
        balances["PrintPoint"] = sum
        
        return balances
    }
    /*
    func timeConverter(t : String) -> String {
        let dateNtime = t.components(separatedBy: "T")
        let date = dateNtime[0].components(separatedBy: "-")
        var month : String
        switch date[1] {
        case "01":
            month = "Jan"
        case "02":
            month = "Feb"
        case "03":
            month = "Mar"
        case "04":
            month = "Apr"
        case "05":
            month = "May"
        case "06":
            month = "Jun"
        case "07":
            month = "Jul"
        case "08":
            month = "Aug"
        case "09":
            month = "Sep"
        case "10":
            month = "Oct"
        case "11":
            month = "Nov"
        case "12":
            month = "Dec"
        default:
            month = "Error!"
        }
        
        return month + " " + date[2]
    }
     */
}
