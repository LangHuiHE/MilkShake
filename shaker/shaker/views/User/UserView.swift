//
//  ContentView.swift
//  shaker
//
//  Created by LangHui He on 1/31/22.
//

import SwiftUI

struct UserView: View {
    @State var timeSlot : String = getMAE()
    @State var locked :Bool = true
    
    var body: some View{
        if locked {
            LockedView(locked: $locked, timeSlot: $timeSlot)
        } else {
            ProtectedView(locked: $locked, timeSlot: $timeSlot)
        }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        UserView()
    }
}

