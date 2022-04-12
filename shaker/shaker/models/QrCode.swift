//
//  QrCode.swift
//  shaker
//
//  Created by LangHui He on 2/12/22.
//

import SwiftUI
import CoreImage.CIFilterBuiltins

let context = CIContext()

func generateBarcode(from string: String) -> UIImage {
    let data = string.data(using: String.Encoding.ascii)
    if let filter = CIFilter(name: "CICode128BarcodeGenerator") {
        filter.setValue(data, forKey: "inputMessage")
        if let outputImage = filter.outputImage {
            if let cgimg = context.createCGImage(outputImage, from: outputImage.extent) {
                return UIImage(cgImage: cgimg)
            }
        }
    }
    return UIImage(systemName: "xmark.circle") ?? UIImage()
}


func generateQRCode(from string: String) -> UIImage {
    let filter = CIFilter.qrCodeGenerator()
    
    filter.message = Data(string.utf8)

    if let outputImage = filter.outputImage {
        if let cgimg = context.createCGImage(outputImage, from: outputImage.extent) {
            return UIImage(cgImage: cgimg)
        }
    }

    return UIImage(systemName: "xmark.circle") ?? UIImage()
}
