//
//  ProtectedView.swift
//  shaker
//
//  Created by LangHui He on 2/13/22.
//

import SwiftUI

struct ProtectedView: View {
    @EnvironmentObject var authentication: ShakerError
    @EnvironmentObject var user: User
    
    @Binding var locked : Bool
    @Binding var timeSlot : String
    
    var width  = UIScreen.main.bounds.width
    var height = UIScreen.main.bounds.height
    
    var gridItemLayout = [GridItem(.flexible()), GridItem(.flexible()), GridItem(.flexible())]
    
    var accountType = ["Tuition", "Meal Plan", "PrintPoint"]
    @State var selectedAccountTypeIndex = 0
    
    var body: some View {
        NavigationView {
            ScrollView{
                ZStack {
                    VStack {
                        Text("Good \(timeSlot)")
                            .font(.system(size: 30))
                            .font(.subheadline)
                            .fontWeight(.semibold)
                        
                        Text(user.userInfo.FirstName)
                            .font(.system(size: 40))
                            .font(.headline)
                            .fontWeight(.bold)
                       
                        Text("Class \(user.getClassYear())")
                            .font(.system(size: 20))
                            .font(.subheadline)
                            .fontWeight(.light)
                        
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
                        
                        Spacer()
                        
                        Text("Total Blance")
                            .font(.system(size: 20))
                            .font(.headline)
                            .fontWeight(.bold)
                            .padding(.top)
                        
                        Picker(selection: $selectedAccountTypeIndex, label: Text("")) {
                                    ForEach(0 ..< accountType.count) {
                                       Text(self.accountType[$0])
                                            .fontWeight(.bold)
                                    }
                        }
                        .pickerStyle(SegmentedPickerStyle())
                        .padding()
                        
                        Text("$\(user.calculateAllBalance()[self.accountType[selectedAccountTypeIndex]]!, specifier: "%.2f")")
                            .font(.largeTitle)
                        
                        Button("Make Payment"){}
                            .frame(width: 220, height: 60)
                            .background(Color.green)
                            .cornerRadius(15)
                            .padding()
                                            
                     
                        VStack(spacing: 15){
                            switch selectedAccountTypeIndex {
                            case 0:
                                ForEach (user.userAccounts.Tuition, id: \.self) {
                                    ay in
                                    Text(ay.Year)
                                        .font(.title2)
                                    
                                    ForEach (ay.Dates,  id: \.self) {
                                        ad in
                                        VStack(alignment: .leading, spacing: 5){
                                            Text(ad.Date)
                                                .font(.title2)
                                            
                                            VStack{
                                                ForEach(0..<ad.Payments.count, id: \.self) {
                                                    Spacer()
                                                    
                                                    let payment = ad.Payments[$0]
                                                    HStack{
                                                        Text(payment.Location)
                                                        Spacer()
                                                        Text("$\(payment.Amount, specifier: "%.2f")")
                                                    }
                                                    
                                                    Spacer()
                                                    
                                                    /*
                                                    if $0 < ad.Payments.count - 1{
                                                        Text("------------------------")
                                                        Spacer()
                                                    }
                                                    */
                                                }
                                            }
                                            .padding()
                                            .frame(width: width/1.1, height: 50 * CGFloat(ad.Payments.count))
                                            .background(Color("DixieRed"))
                                            .cornerRadius(15)
                                        }
                                    }
                                }
                                
                            case 1:
                                ForEach (user.userAccounts.MealPlan, id: \.self) {
                                    ay in
                                    Text(ay.Year)
                                        .font(.title2)
                                    
                                    ForEach (ay.Dates,  id: \.self) {
                                        ad in
                                        VStack(alignment: .leading, spacing: 5){
                                            Text(ad.Date)
                                                .font(.title2)
                                            
                                            VStack{
                                                ForEach(ad.Payments, id: \.self) {
                                                    payment in
                                                    Spacer()
                                                    HStack{
                                                        Text(payment.Location)
                                                        Spacer()
                                                        Text("$\(payment.Amount, specifier: "%.2f")")
                                                    }
                                                    Spacer()
                                                }
                                            }
                                            .padding()
                                            .frame(width: width/1.1, height: 35 * CGFloat(ad.Payments.count))
                                            .background(Color("DixieRed"))
                                            .cornerRadius(15)
                                        }
                                    }
                                }
                                
                            case 2:
                                ForEach (user.userAccounts.PrintPoint, id: \.self) {
                                    ay in
                                    Text(ay.Year)
                                        .font(.title2)
                                    
                                    ForEach (ay.Dates,  id: \.self) {
                                        ad in
                                        VStack(alignment: .leading, spacing: 5){
                                            Text(ad.Date)
                                                .font(.title2)
                                            
                                            VStack{
                                                ForEach(ad.Payments, id: \.self) {
                                                    payment in
                                                    Spacer()
                                                    HStack{
                                                        Text(payment.Location)
                                                        Spacer()
                                                        Text("$\(payment.Amount, specifier: "%.2f")")
                                                    }
                                                    Spacer()
                                                }
                                            }
                                            .padding()
                                            .frame(width: width/1.1, height: 35 * CGFloat(ad.Payments.count))
                                            .background(Color("DixieRed"))
                                            .cornerRadius(15)
                                        }
                                    }
                                }
                                
                            default:
                                Text("Something is wrong QAQ!")
                            }
                        }.padding()
                       
                    }
                }
                    
                .alert(item: $user.error) {
                    error in
                        authentication.updateValidation(success: false)
                        return Alert(title: Text("Unauthorized Operation"), message: Text(error.localizedDescription))
                    }

                    .toolbar {
                        ToolbarItem(placement: .navigationBarTrailing) {
                            NavigationLink(destination: QrCodeView()){
                                Label("qrCode", systemImage: "qrcode")
                            }
                        }
                        ToolbarItem(placement: .principal) {
                            VStack {
                                Text(user.getUserTypeName())
                                    .font(.headline)
                                Text(user.getFullID())
                                    .font(.subheadline)
                            }
                        }
                        ToolbarItem(placement: .navigationBarLeading) {
                            Button {
                                locked.toggle()
                            } label : {
                                Image(systemName: "lock")
                            }
                        }
                }
            }
        }
    }
}
