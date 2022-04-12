//
//  QrCodeView.swift
//  shaker
//
//  Created by LangHui He on 2/12/22.
//

import SwiftUI
import SlideOverCard
import CodeScanner

struct QrCodeView: View {
    @EnvironmentObject var user: User
    
    private var screenWidth  = UIScreen.main.bounds.width
    private var screenHeight = UIScreen.main.bounds.height
    
    @State var isPresented = false
    @State var results : [String] = []
    @State var message : String = ""
    @State var messageHeader : String = ""
    @State private var showAlert: Bool = false
    

    var body: some View {
        ZStack {
            VStack{
                VStack {
                    Image(uiImage: generateBarcode(from: "\(user.getFullID())"))
                        .resizable()
                        .interpolation(.none)
                        .scaledToFit()
                        .frame(width: screenWidth/1.05, height: screenHeight/4)
                    
                    Image(uiImage: generateQRCode(from: "\(user.getFullID())"))
                        .resizable()
                        .interpolation(.none)
                        .scaledToFit()
                        .frame(width: screenWidth/1.25, height: screenWidth/1.25)
                }.background(Color(.white))

                
                Button{
                    isPresented.toggle()
                } label: {
                    Label("Scan", systemImage: "qrcode.viewfinder")
                }.padding()
            }
        }
        .slideOverCard(isPresented: $isPresented) {
            PlaceholderContent(isPresented: $isPresented, screenWidth: screenWidth, screenHeight: screenHeight, results: $results, message: $message, messageHeader : $messageHeader, showAlert: $showAlert)
        }.alert(messageHeader, isPresented: $showAlert) {
            if messageHeader == "Can't Recognize Result" || messageHeader == "Login Fail" ||  messageHeader == "Sign Up Fail"{
                Button("ReScan") { isPresented.toggle()}
            }
            if messageHeader == "Confirm Login on Website" || messageHeader == "Event Sign Up"{
                Button("Yes") {
                    ExcecuteScannerResult(result: results) {
                        response in
                        switch response {
                        case.success(let mess):
                            messageHeader = mess
                            message = ""
                            showAlert = true
                            DispatchQueue.main.asyncAfter(deadline: .now()+2.5) {
                                showAlert.toggle()
                            }
                        case.failure(_):
                            message = "Please ReScan and Try Again"
                            messageHeader = "Fail!"
                            showAlert = true
                    }
                }
            }
                Button("Cancel", role: .cancel) {showAlert = false}
            }
        } message: {
            Text(message)
        }
    }
}


struct PlaceholderContent: View {
    @Binding var isPresented: Bool
    let screenWidth : CGFloat
    let screenHeight : CGFloat
    
    @Binding var results : [String]
    
    @Binding var message : String
    @Binding var messageHeader : String
    @Binding var showAlert : Bool
    
    var body: some View {
        VStack(alignment: .center) {
            VStack {
                Text("Scan QR code").bold()
            }
            //.alert()
            CodeScannerView(codeTypes: [.qr], scanMode: .continuous){ response in
                switch response {
                case.success(let result):
                    isPresented.toggle()
                    results = result.string.components(separatedBy: "%20")
                    // print(details)
                   (messageHeader, message) = ProcessScannerResult(result: results)
                    showAlert.toggle()
                    
                case.failure(let err):
                    print("Scanning failed:\(err.localizedDescription)")
                }
                
            }.frame(width: screenWidth/1.4, height: screenWidth/1.4).cornerRadius(25)
            Spacer()
        }.frame(width: screenWidth/1.3, height: screenHeight/2)
    }
    
//    func handleScan(result: Result<ScanResult, ScanError>) {
//        switch result {
//        case.success(let result):
//            isPresented.toggle()
//            let details = result.string.components(separatedBy: "%20")
//            print(details)
//            ProcessScannerResult(result: details)
//
//
//        case.failure(let err):
//            print("Scanning failed:\(err.localizedDescription)")
//        }
//    }
}

/*
struct QrCodeView_Previews: PreviewProvider {
    static var previews: some View {
        QrCodeView()
    }
}
*/

