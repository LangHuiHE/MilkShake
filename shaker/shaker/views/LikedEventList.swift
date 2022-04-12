//
//  LikedEventList.swift
//  shaker
//
//  Created by LangHui He on 3/3/22.
//

import SwiftUI

struct LikedEventList: View {
    @EnvironmentObject var events : Events
    @State var liked = [EventInfo]()
    @State var maybe = [EventInfo]()
    @State var nah   = [EventInfo]()
    
    var body: some View {
        List{
            if liked.count > 0 {
                Section(header: Text("Liked Events")) {
                    ForEach(liked, id: \.Id) {event in
                        NavigationLink(destination: EventDetail(eventInfo: event).environmentObject(events)) {
                            Text(event.Title)
                        }
                    }
                }
            }

            if maybe.count > 0 {
                Section(header: Text("Maybe Events")) {
                    ForEach(maybe, id: \.Id) {event in
                        NavigationLink(destination: EventDetail(eventInfo: event).environmentObject(events)) {
                            Text(event.Title)
                        }
                    }
                }
            }

            if nah.count > 0 {
                Section(header: Text("Nah")) {
                    ForEach(nah, id: \.Id) {event in
                        NavigationLink(destination: EventDetail(eventInfo: event).environmentObject(events)) {
                            Text(event.Title)
                        }
                    }
                }
            }
        }
        .navigationBarTitle("Events")
        .navigationBarTitleDisplayMode(.inline)
        .onAppear(perform: formateEventList)
           
    }
}


extension LikedEventList {
    func formateEventList(){
        for event in self.events.dict.keys.sorted(by: <){
            switch self.events.dict[event] {
            case 1: 
                if !self.liked.contains(event) {
                    self.liked.append(event)
                }
            case -1 :
                if !self.nah.contains(event) {
                    self.nah.append(event)
                }
            default:
                if !self.maybe.contains(event) {
                    self.maybe.append(event)
                }
            }
        }
    }
}
