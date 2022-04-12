//
//  Event.swift
//  shaker
//
//  Created by LangHui He on 3/1/22.
//

import SwiftUI

struct EventInfo : Decodable, Hashable, Equatable, Comparable, Encodable {
    static func < (lhs: EventInfo, rhs: EventInfo) -> Bool {
        return lhs.StartDate < rhs.StartDate
    }
    
    var Address         : String
    var Category        : String
    var City            : String
    var Department      : String
    var EndDate         : String
    var EventDescription: String
    var Id              : Int
    var FreeFood        : Int
    var Merchandise     : Int
    var HostUserID      : Int
    var ImageUrl        : String
    var Location        : String
    var Note            : String
    var StartDate       : String
    var State           : String
    var Title           : String
    /*
    init() {
        Address = "225 S 700 E"
        Category = "Social"
        City = "ST GEORGE"
        Department = "DSU"
        EndDate = "2022-03-03T22:00:00-07:00"
        EventDescription = "DSU Esports brings Southern Utah Smash on campus!Join us for weekly Smash Bros. Ultimate tournaments"
        Id = 7678979
        FreeFood = 1
        Merchandise = 1
        HostUserID = 205928
        ImageUrl = "https://se-images.campuslabs.com/clink/images/99845968-8027-4215-a781-bf480037c0a30a845568-6da0-4cc4-9f45-51ac53cfd8ad.png"
        Location = "Atwood Innovation Plaza - Lovesac Lounge"
        Note = ""
        StartDate = "2022-03-03T17:30:00-07:00"
        State = "UT"
        Title = "Red Rock Rumble"
    }
     */
    
    static func == (lhs: EventInfo, rhs: EventInfo) -> Bool {
        return
            lhs.Address          == rhs.Address             &&
            lhs.Category         == rhs.Category            &&
            lhs.City             == rhs.City                &&
            lhs.Department       == rhs.Department          &&
            lhs.EndDate          == rhs.EndDate             &&
            lhs.EventDescription == rhs.EventDescription    &&
            lhs.Id               == rhs.Id                  &&
            lhs.FreeFood         == rhs.FreeFood            &&
            lhs.Merchandise      == rhs.Merchandise         &&
            lhs.HostUserID       == rhs.HostUserID          &&
            lhs.ImageUrl         == rhs.ImageUrl            &&
            lhs.Location         == rhs.Location            &&
            lhs.Note             == rhs.Note                &&
            lhs.StartDate        == rhs.StartDate           &&
            lhs.State            == rhs.State               &&
            lhs.Title            == rhs.Title
    }
}


class Events : ObservableObject, Equatable {
    @Published var dict = [EventInfo:Int]()
    
    static func == (lhs: Events, rhs: Events) -> Bool {
        return lhs.dict == rhs.dict
    }
    
    func getFirstFollowingEvent() -> EventInfo {
        for event in dict.keys.sorted(by: <){
            if LocalizeDateTime(s: event.StartDate) > Date(){
                return event
            }
        }
        return Array(dict.keys)[0]
    }
    
    func updateRecord(){
        /*
        do {
            // Create JSON Encoder
            let encoder = JSONEncoder()

            // Encode Note
            let data = try encoder.encode(dict)

            // Write/Set Data
            UserDefaults.standard.set(data, forKey: "eventRecord")

        } catch {
            print("Unable to Encode Note (\(error))")
        }
        
        UserDefaults.standard.set(dict, forKey: "eventRecord")
         */
    }
    
    func restoreRecord(){
        // Read/Get Data
        /*
        if let data = UserDefaults.standard.data(forKey: "eventRecord") {
            do {
                // Create JSON Decoder
                let decoder = JSONDecoder()

                // Decode Note
                let record = try decoder.decode([EventInfo:Int].self, from: data)
                for event in record.keys {
                    if dict[event] != nil {
                        dict[event] = record[event]
                    }
                }

            } catch {
                print("Unable to Decode Note (\(error))")
            }
        }
         */
    }
}


func LocalizeDateTime (s : String) -> Date {
    let localISOFormatter = ISO8601DateFormatter()
    localISOFormatter.timeZone = TimeZone.current

    // Printing a Date
    // let date = Date()
    // print(localISOFormatter.string(from: date))

    // Parsing a string timestamp representing a date
    return localISOFormatter.date(from: s)!
}

func RelativeDateTime(startDate : Date, endDate : Date) -> String {
    let formatter = RelativeDateTimeFormatter()
    formatter.unitsStyle = .full
    let now = Date()
    let calendar = Calendar.current
    if calendar.isDateInYesterday(startDate) { return "Yesterday" }
    else if calendar.isDateInTomorrow(startDate) { return "Tomorrow" }
    else if calendar.isDateInToday(startDate) {
        if endDate > now && startDate <= now { // Started but not ended yet
            return "Started " + formatter.localizedString(for: startDate, relativeTo: now)
        } else if endDate <= now && startDate <= now{ // Alreadly ended
            return "Ended " + formatter.localizedString(for: endDate, relativeTo: now)
        }
    }
    // get exampleDate relative to the current date
    // let relativeDate = formatter.localizedString(for: date, relativeTo: Date())

    // print it out
    return formatter.localizedString(for: startDate, relativeTo: now)
}

func FormatTimePeriod(startDate : String, endDate : String) -> String {
    let calendar = Calendar.current
    let sd = LocalizeDateTime(s: startDate)
    let ed = LocalizeDateTime(s: endDate)
    
    if calendar.isDateInToday(sd) {
        if calendar.isDateInToday(ed) {
            return "Today " + sd.formatted(date: .omitted, time: .shortened) + " ~ " + ed.formatted(date: .omitted, time: .shortened)
        }
        return "Today " + sd.formatted(date: .omitted, time: .shortened) + " ~ " + ed.formatted(date: .long, time: .shortened)
    }
    if calendar.isDate(sd, inSameDayAs: ed) {
        return sd.formatted(date: .long, time: .shortened) + " ~ " + ed.formatted(date: .omitted, time: .shortened)
    }
    return sd.formatted(date: .long, time: .shortened) + " ~ " + ed.formatted(date: .long, time: .shortened)
}
