//
//  LoginView.swift
//  shaker
//
//  Created by LangHui He on 2/1/22.
//

import SwiftUI


struct LoginView: View {
    @StateObject private var loginVM = LoginViewModel()
    @EnvironmentObject var authentication: ShakerError
    
    enum Field {
        case UserID
        case Password
    }
    
    @FocusState private var foucesdField: Field?

    var body: some View {
        VStack {
            WelcomeText()
            UserImage()
            TextField("UserID",text: $loginVM.credentials.Id)
                .focused($foucesdField, equals: .UserID)
                .submitLabel(.next)
                .keyboardType(.numberPad)
                .multilineTextAlignment(.center)
                .padding()
                .foregroundColor(.black)
                .background(Color("Light Grayish Yellow"))
                .cornerRadius(500)
                .padding()

            
            SecureField("Password", text: $loginVM.credentials.UserPassword)
                .focused($foucesdField, equals: .Password)
                .submitLabel(.go)
                .multilineTextAlignment(.center)
                .padding()
                .foregroundColor(.black)
                .background(Color("Light Grayish Yellow"))
                .cornerRadius(500)
                .padding()
            
                .onSubmit {
                    switch foucesdField {
                    case .UserID:
                        foucesdField = .Password
                    default:
                        loginVM.showProgressView = true
                        loginVM.login {
                            success in authentication.updateValidation(success: success)
                        }
                    }
                }

            
            if loginVM.showProgressView {
                ProgressView()
            }

            Button("LOGIN"){
                loginVM.login {
                    success in authentication.updateValidation(success: success)
                }
            }
            .font(.headline)
            .foregroundColor(Color("Light Grayish Yellow"))
            .padding()
            .frame(width: 220, height: 60)
            .background(Color("DixieBlue"))
            .cornerRadius(15)
            .disabled(loginVM.loginDisabled)
            .padding()
            

        }.padding()
        .autocapitalization(.none)
        .onTapGesture {
                hideKeyboard()
        }
        // .disabled(loginVM.showProgressView)
        .alert(item: $loginVM.error) { error in
            Alert(title: Text("Invlid Login"), message: Text(error.localizedDescription))
        }
    }
}

extension View {
    func hideKeyboard() {
        let resign = #selector(UIResponder.resignFirstResponder)
        UIApplication.shared.sendAction(resign, to: nil, from: nil, for: nil)
    }
}

struct LoginView_Previews: PreviewProvider {
    static var previews: some View {
        LoginView()
    }
}

struct WelcomeText: View {
    var body: some View {
        Text("Dixie State University")
            .font(.system(size: 30))
            .font(.subheadline)
        Text("Good \(getMAE())!")
            .font(.system(size: 30))
            .font(.subheadline)
            .fontWeight(.semibold)
        Text("Trailblazers")
            .font(.title2)
            .fontWeight(.semibold)
            .padding(.bottom, 20)
    }
}

struct UserImage: View {
    var body: some View {
        Image("LoginImage")
            .resizable()
            .aspectRatio(contentMode: .fit)
            .frame(width: 150, height: 150)
            .clipped()
            // .cornerRadius(100)
            .padding(.bottom, 75)
    }
}
