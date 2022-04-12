//
//  EventCellView.swift
//  shaker
//
//  Created by LangHui He on 2/28/22.
//

import SwiftUI

struct EventCellView: View {
    var screenWidth  = UIScreen.main.bounds.width
    var screenHeight = UIScreen.main.bounds.height
    
    @State private var showingDetail = false
    // how far the circle has been dragged
    @State private var offset = CGSize.zero
    
    @State var doubleTapLike : Bool = false
    
    var textColor = Color("Very Dark Grayish Red")
    var backViewColor = Color("Tan")
    
    let eventInfo : EventInfo
    @EnvironmentObject var events : Events
    
    var body: some View {
        VStack {
            ZStack {
                // Back View
                Text("")
                    .frame(width: screenWidth/1.2, height: screenHeight/2.2)
                    .background(backViewColor)
                    .offset(y: screenHeight/15)
                
                VStack {
                // Front View
                    AsyncImage(url: URL(string: self.eventInfo.ImageUrl)) { phase in
                        if let poster = phase.image {
                                poster
                                    .resizable()
                            } else {
                                ProgressView()
                            }
                    }
                    // .glow(color: .black, radius: 5)
                    .shadow(color: .black, radius: screenWidth/50)
                    .frame(width: screenWidth/1.3, height: screenHeight/2.2)
                    //.scaledToFit()
                    // .position(x: 0, y: 0)
                    
                    /*
                    Image(uiImage: self.poster)
                            .resizable()
                            // .scaledToFill()
                            // .glow(color: .black, radius: 5)
                            .shadow(color: .black, radius: screenWidth/50)
                            .frame(width: screenWidth/1.3, height: screenHeight/2.2)
                            // .position(x: 0, y: 0)
                            .onAppear{
                                guard let url = URL(string: self.event.ImageUrl) else { return }
                                    URLSession.shared.dataTask(with: url) { (data, response, error) in
                                        guard let data = data else { return }
                                        guard let image = UIImage(data: data) else { return }
                                        DispatchQueue.main.async {
                                            self.poster = image
                                        }
                                }.resume()
                            }
                    */
                    VStack {
                        HStack {
                            Text(eventInfo.Title)
                                .foregroundColor(.black)
                                .font(.title2)
                                .padding(.leading)
                                .padding(.trailing)
                                .multilineTextAlignment(.center)
                        }
                        HStack {
                            Text(eventInfo.Department)
                                .foregroundColor(textColor)
                            Spacer()
                        }
                        HStack {
                            if events.dict[eventInfo] == 1 {
                                Image(systemName: "heart.fill")
                                    .foregroundColor(Color("DixieRed"))
                            } else if events.dict[eventInfo] == -1 {
                                Image(systemName: "heart.slash.fill")
                                    .foregroundColor(Color("DixieRed"))
                            }
                            Spacer()
                            Text(eventInfo.Category)
                                .foregroundColor(textColor)
                        }
                        Text(RelativeDateTime(startDate: LocalizeDateTime(s :eventInfo.StartDate), endDate: LocalizeDateTime(s :eventInfo.EndDate)))
                            .textCase(.uppercase)
                            .foregroundColor(.black)
                    }.padding().frame(width: screenWidth/1.2).background(backViewColor)
                }.padding()
            }.padding()
        }.onTapGesture(count: 2) {
            events.dict[eventInfo] = 1
            events.updateRecord()
            self.doubleTapLike.toggle()
            DispatchQueue.main.asyncAfter(deadline: .now()+0.5) {
                self.doubleTapLike.toggle()
            }
        }.overlay(Image(systemName: "heart.fill")
                    .resizable()
                    .frame(width: screenWidth / 3, height: screenWidth / 3, alignment: .center)
                    .foregroundColor(Color("DixieRed"))
                    .opacity(self.doubleTapLike ? 1.0:0)
                    .scaleEffect(self.doubleTapLike ? 1.2:0.1 )
                    .animation(.easeInOut, value: 1)
        ).onTapGesture(count: 1)  {
            showingDetail = true
        }.gesture(
            DragGesture()
            .onEnded { gesture in
                // print(gesture.translation.height)
                // print(abs(gesture.translation.width))
                
                 if ((gesture.translation.height <= -50) && (abs(gesture.translation.width) <= 20)) {
                     showingDetail = true
                 }
            }
        ).sheet(isPresented: $showingDetail) {
            EventDetailSheet(eventInfo: eventInfo).environmentObject(events)
        }
    }
}


extension View {
    func glow(color: Color, radius: CGFloat) -> some View {
        self
            .overlay(self.blur(radius: radius / 6))
            .shadow(color: color, radius: radius / 3)
            .shadow(color: color, radius: radius / 3)
            .shadow(color: color, radius: radius / 3)
    }
}

/*
struct EventCellView_Previews: PreviewProvider {
    static var event = EventInfo()
    static var previews: some View {
        EventCellView(event: event)
    }
}
*/
