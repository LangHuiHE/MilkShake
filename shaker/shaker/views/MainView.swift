//
//  MainView.swift
//  shaker
//
//  Created by LangHui He on 2/5/22.
//

import SwiftUI

struct MainView: View {
    @StateObject var user = User()
    @State var selection = 0
    @StateObject var events = Events()
    
    init() {
        UITabBar.appearance().barTintColor = .systemBackground
        UITabBar.appearance().backgroundColor = UIColor(Color("DixieBlue"))
        
        // UINavigationBar.appearance().backgroundColor = tan
    }
    
    var body: some View {
        TabView(selection: $selection){
            ExtractedEventView()
                .environmentObject(user)
                .environmentObject(events)
                .tag(0)
            ExtractedInfoView()
                .environmentObject(user)
                .tag(1)
        }
    }
}

struct MainView_Previews: PreviewProvider {
    static var previews: some View {
        MainView()
    }
}

struct ExtractedEventView: View {
    var body: some View {
        EventView()
            .tabItem {
                Label("Event", systemImage: "list.dash")
            }
    }
}

struct ExtractedInfoView: View {
    var body: some View {
        UserView()
            .tabItem {
                Label("Info", systemImage: "person")
            }
    }
}
