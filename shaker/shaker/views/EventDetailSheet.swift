//
//  EventDetail.swift
//  shaker
//
//  Created by LangHui He on 3/2/22.
//

import SwiftUI

struct EventDetailSheet: View {
    var screenWidth  = UIScreen.main.bounds.width
    var screenHeight = UIScreen.main.bounds.height
    
    let eventInfo : EventInfo
    @EnvironmentObject var events : Events
    @State var doubleTapLike : Bool = false
    
    @Environment(\.dismiss) var dismiss
    
    var body: some View {
        VStack {
            HStack {
                Text(eventInfo.Title)
                    .font(.title)
                    .multilineTextAlignment(.center)
                Spacer()
                Button {
                    dismiss()
                } label : {
                    Image(systemName: "arrow.down.right.and.arrow.up.left")
                }
            }

            ScrollView {
                /*
                if poster == emptyImage {
                    poster
                    AsyncImage(url: URL(string: self.event.ImageUrl)) { phase in
                        if let poster = phase.image {
                                poster
                                    .resizable()
                            } else if phase.error != nil {
                                Text("There was an error loading the image.")
                            } else {
                                ProgressView()
                            }
                    }
                    .shadow(color: .black, radius: screenWidth/50)
                    .frame(width: screenWidth/1.3, height: screenHeight/2.2)
                } else {
                    poster.resizable()
                        .shadow(color: .black, radius: screenWidth/50)
                        .frame(width: screenWidth/1.3, height: screenHeight/2.2)
                }
                 */
                AsyncImage(url: URL(string: self.eventInfo.ImageUrl)) { phase in
                    if let poster = phase.image {
                            poster
                                .resizable()
                        } else if phase.error != nil {
                            Text("There was an error loading the image.")
                        } else {
                            ProgressView()
                        }
                }
                .shadow(color: .black, radius: screenWidth/50)
                .frame(width: screenWidth/1.3, height: screenHeight/2.2)
                
                
                Text(FormatTimePeriod(startDate: eventInfo.StartDate ,endDate:eventInfo.EndDate)).font(.title2).multilineTextAlignment(.center)
                
                
                HStack {
                    if eventInfo.FreeFood == 1{
                        Text("Free Food \(Image(systemName: "takeoutbag.and.cup.and.straw"))").padding()
                    }
                    if eventInfo.Merchandise == 1{
                        Text("Free Gift \(Image(systemName: "gift"))").padding()
                    }
                }
 
                Text(eventInfo.EventDescription).multilineTextAlignment(.leading).padding()
                
                // Text("Location:")
                VStack {
                    Image(systemName: "location")
                    Text(eventInfo.Address)
                    Text(eventInfo.City)
                    Text(eventInfo.State)
                }.padding()
                HStack {
                    Text(String(eventInfo.Id))
                    Spacer()
                    Text(eventInfo.Department)
                }.padding()

                HStack(spacing: screenWidth / 4) {
                    switch events.dict[eventInfo] {
                    case -1:
                        Button {
                            events.dict[eventInfo] = 0
                        } label : {
                            Image(systemName: "heart.slash.fill")
                                .resizable()
                                .scaledToFit()
                                .foregroundColor(Color("DixieRed"))
                                .frame(width: screenWidth / 10, height: screenWidth / 10)
                        }
                        
                    case 1:
                        Button {
                            events.dict[eventInfo] = 0
                        } label : {
                            Image(systemName: "heart.fill")
                                .resizable()
                                .scaledToFit()
                                .foregroundColor(Color("DixieRed"))
                                .frame(width: screenWidth / 10, height: screenWidth / 10)
                        }
                    default:
                        Button {
                            events.dict[eventInfo] = -1
                        } label : {
                            Image(systemName: "heart.slash")
                                .resizable()
                                .scaledToFit()
                                .foregroundColor(Color("DixieRed"))
                                .frame(width: screenWidth / 10, height: screenWidth / 10)
                        }
                        
                        Button {
                            events.dict[eventInfo] = 1
                        } label : {
                            Image(systemName: "heart")
                                .resizable()
                                .scaledToFit()
                                .foregroundColor(Color("DixieRed"))
                                .frame(width: screenWidth / 10, height: screenWidth / 10)
                        }
                    }
                }.padding()
                 
            }.onTapGesture(count: 2) {
                events.dict[eventInfo] = 1
                self.doubleTapLike.toggle()
                DispatchQueue.main.asyncAfter(deadline: .now()+0.5) {
                    self.doubleTapLike.toggle()
                }
            }.overlay(Image(systemName: "heart.fill")
                        .resizable()
                        .frame(width: screenWidth / 3, height: screenWidth / 3, alignment: .center)
                        .foregroundColor(Color("DixieRed"))
                        .opacity(self.doubleTapLike ? 1.0:0)
                        .scaleEffect(self.doubleTapLike ? 1.2:0.1)
                        .animation(.easeInOut, value: 1)
            )
        }.padding()
    }
}
/*
struct EventDetail_Previews: PreviewProvider {
    static var event = EventInfo()
    static var previews: some View {
        EventDetail(event: event)
    }
}
*/
