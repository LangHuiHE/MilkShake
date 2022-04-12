//
//  LockedView.swift
//  shaker
//
//  Created by LangHui He on 2/13/22.
//

import SwiftUI
import LocalAuthentication

struct LockedView: View {
    @Binding var locked : Bool
    @Binding var timeSlot : String
    
    @EnvironmentObject var authentication: ShakerError
    @EnvironmentObject var user: User
    
    var width  = UIScreen.main.bounds.width
    var height = UIScreen.main.bounds.height
    
    var body: some View {
        NavigationView {
            ZStack {
                VStack {
                    Text("Good \(timeSlot)")
                        .font(.system(size: 30))
                        .font(.subheadline)
                        .fontWeight(.semibold)
                    if user.userInfo.FirstName != "" {
                        Text(user.userInfo.FirstName)
                            .font(.system(size: 40))
                            .font(.headline)
                            .fontWeight(.bold)
                    }
                    
                    if user.userInfo.EnrollDate != "" {
                        Text("Class \(user.getClassYear())")
                            .font(.system(size: 20))
                            .font(.subheadline)
                            .fontWeight(.light)
                    }

                    if user.userInfo.StudentInfo.Credits != 0 {
                        Text("\(user.userInfo.StudentInfo.Department)")
                            .font(.system(size: 20))
                            .font(.subheadline)
                            .fontWeight(.light)
                        Text("\(user.userInfo.StudentInfo.Major)")
                            .font(.system(size: 20))
                            .font(.subheadline)
                            .fontWeight(.light)
                    }
                    
                    Button {
                        locked.toggle()
                        user.loadAccounts()
                    } label: {
                        Image(systemName: "faceid")
                            .resizable()
                            .scaledToFit()
                            .padding()
                            .frame(width: width/4, height: width/4)
                            //.background(Color(verdigris))
                    }
                    Image("LoginImage")
                        .resizable()
                        .scaledToFit()
                    Spacer()
                    }
                }
                
                .alert(item: $user.error) { error in
                    authentication.updateValidation(success: false)
                    return Alert(title: Text("Unauthorized Operation"), message: Text(error.localizedDescription))
                }

                .toolbar {
                    ToolbarItem(placement: .navigationBarTrailing) {
                        NavigationLink(destination: QrCodeView()){
                            Label("qrCode", systemImage: "qrcode")
                        }
                    }
                    ToolbarItem(placement: .principal) { // <3>
                        VStack {
                            Text(user.getUserTypeName())
                                .font(.headline)
                            if user.userInfo.Id != 0 {
                                Text(user.getFullID())
                                    .font(.subheadline)
                            }
                        }
                    }
                    ToolbarItem(placement: .navigationBarLeading) {
                        Button("Log out") {authentication.updateValidation(success: false)
                           removeAllUserDefaults()
                        }.foregroundColor(Color("DixieRed"))
                    }
            }
        }
    }
}
