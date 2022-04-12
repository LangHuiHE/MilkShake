//
//  helper.swift
//  shaker
//
//  Created by LangHui He on 2/9/22.
//

import SwiftUI

// let hour = NSCalendar.currentCalendar().component(.Hour, fromDate: NSDate()) Swift 2 legacy
let hour = Calendar.current.component(.hour, from: Date())

func getMAE() -> String {
    let hour = Calendar.current.component(.hour, from: Date())
    switch hour {
        case 6..<12 : return "Morning"
        case 12..<17 : return "Afternoon"
        case 17..<22 : return "Evening"
        default: return "Night"
    }
}




// Bounds
/*
var screenBounds : CGRect = UIScene.main.bounds
var bWidth : CGFloat = UIScene.main.bounds.width
var bHeight : CGFloat = UIScene.main.bounds.height
*/
/*
// nativeBounds
var screenBounds : CGRect = UIScene.main.bounds
var bWidth : CGFloat = UIScene.main.bounds.width
var bHeight : CGFloat = UIScene.main.bounds.height
*/

import WebKit

struct GifImage: UIViewRepresentable {
    private let name: String
    
    init(_ name : String) {
        self.name = name
    }
    
    func makeUIView(context: Context) -> WKWebView {
        let webView = WKWebView()
        let url = Bundle.main.url(forResource: name, withExtension: "gif")!
        let data = try! Data(contentsOf: url)
        webView.load(
            data,
            mimeType: "image/gif",
            characterEncodingName: "UTF-8",
            baseURL: url.deletingLastPathComponent()
        )
        return webView
    }
    
    func updateUIView(_ uiView: WKWebView, context: Context) {
        uiView.reload()
    }
}
