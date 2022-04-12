//
//  DeviceInfo.swift
//  shaker
//
//  Created by LangHui He on 1/31/22.
//

import SwiftUI

class DeviceInfo: UIDevice, Codable {
    var Name : String = UIDevice.current.name
    var Model : String = UIDevice.current.model
    var UUID : String = UIDevice.current.identifierForVendor!.uuidString
}
