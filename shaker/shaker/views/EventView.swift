//
//  EventView.swift
//  shaker
//
//  Created by LangHui He on 2/5/22.
//

import SwiftUI

struct EventView: View {
    @EnvironmentObject var user: User
    @EnvironmentObject var events : Events
    
    var screenWidth  = UIScreen.main.bounds.width
    var screenHeight = UIScreen.main.bounds.height
    
    var body: some View {
        NavigationView{
            ScrollViewReader { scrollView in
                ScrollView(.horizontal, showsIndicators: false) {
                    HStack(spacing: 0) {
                        ForEach([EventInfo](events.dict.keys.sorted(by: <)), id:\EventInfo.Id) { event in
                            EventCellView(eventInfo: event).environmentObject(events)
                        }
                        
                        // funny meme
                        VStack {
                            ZStack {
                                // Back View
                                Text("Back View")
                                    .frame(width: screenWidth/1.2, height: screenHeight/2.2)
                                    .background(Color("Light Grayish Yellow"))
                                    .offset(y: screenHeight/15)
                                
                                VStack {
                                // Front View
                                    GifImage("mad-angry")
                                        .shadow(color: .black, radius: screenWidth/50)
                                        .frame(width: screenWidth/1.3, height: screenHeight/2.2)
                                        .scaledToFit()
                                        .offset(y: -screenHeight/15)
                                }.padding()
                            }.padding()
                        }
                        
                     // END: funny meme
                        
                    }
                }.onChange(of: events.dict.count, perform: { value in
                    // withAnimation(.easeInOut) {
                        scrollView.scrollTo(events.getFirstFollowingEvent().Id, anchor: .center)
                    // }
                })
                Button("What Next?"){
                    withAnimation(.easeInOut) {
                        scrollView.scrollTo(events.getFirstFollowingEvent().Id, anchor: .center)
                    }
                }.padding(.bottom)
            }.onAppear(perform: getAllEvents)
                
            .toolbar {
                ToolbarItem(placement: .navigationBarTrailing) {
                    NavigationLink(destination: LikedEventList()){
                        Label("likedEvent", systemImage: "heart.text.square")
                    }
                    
                }
                ToolbarItem(placement: .navigation){
                    Text("Events").font(.title).bold()
                }
            }
        }
    }
}

extension EventView {
    func getAllEvents() {
        let request = ConstructApi(method: "Get", endPoint: "/mobile/events", bodyData: nil, addAuthTokenAndDeviceInfo: false, header: nil)
        let session = URLSession.shared
        session.dataTask(with: request) {(data, response, error) in
            if let data = data {
                if let response_obj = try?
                 JSONDecoder().decode([EventInfo].self, from: data) {
                    DispatchQueue.main.async {
                        for event in response_obj {
                            if self.events.dict[event] == nil {
                                self.events.dict[event] = 0
                            }
                        }
                        self.events.restoreRecord()
                    }
                }
            }
        }.resume()
    }
}

struct EventView_Previews: PreviewProvider {
    static var previews: some View {
        EventView()
    }
}
