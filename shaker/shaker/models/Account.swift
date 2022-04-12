//
//  Account.swift
//  shaker
//
//  Created by LangHui He on 2/20/22.
//

import SwiftUI

struct AccountPayment : Codable, Hashable {
    var PaymentMethod   : String
    var ProcessDate     : String
    var Semester        : String
    var Amount          : Float
    var Location        : String
    init() {
        PaymentMethod   = ""
        ProcessDate     = ""
        Semester        = ""
        Amount          = 0
        Location        = ""
    }
}

struct AccountDate : Codable, Hashable {
    var Date    : String
    var Payments : [AccountPayment]
    init() {
        Date = ""
        Payments = [AccountPayment]()
    }
}
    
struct AccountYear : Codable, Hashable {
    var Year    : String
    var Dates   : [AccountDate]
    init() {
        Year = ""
        Dates = [AccountDate]()
    }
}

struct Accounts : Codable {
    var Tuition            : [AccountYear]
    var MealPlan           : [AccountYear]
    var PrintPoint         : [AccountYear]
    
    init(){
        Tuition     = [AccountYear]()
        MealPlan    = [AccountYear]()
        PrintPoint  = [AccountYear]()
    }
}

