# MilkShake

**MilkShake** is the senior project for my bachelor's degree in Computer Science at Utah Tech University (aka Dixie State University). 

Also, it's my first personal and self-designed project that integrates with multiple technologies, tools, and platforms and has completed a dataflow circle. This project could be considered a real product because it contains a database, server, cache management, webpages, and mobile application in the IOS platform. 

**MilkShake** was born with the mission of solving the problem of students and faculties who couldn't find a good product for their campus life. 



## Deployment
1. Setup
    * MilkJug
        1. Database 
        Using SQL files at `setup/database` to create a Database and run it in the background
        **All data are generated randomly**
        2. Cache: Redis
        Run Redis in the background
        Don't forget to update the config setting at `milkjug/conf/app.conf`
    * Shaker
        1. Update server address at `shaker/conf.swift`
        2. Build and install on your device
2. Run
    * Using [bee](https://github.com/astaxie/bee) `bee run` to start server


## Demo Screenshot at `display`

## Problem Description
Students and facilities in Utah Tech University (aka Dixie State University) couldn't find a good mobile application for their campus life. The applications developed or published by the university are old, which means most of them had the last update more than a year ago. Meanwhile, those applications don't have the functionality and information that are interesting to users.

## Solution Description
In short, the solution is to make the applications or products we like. 

#### Feature of Application
* Event Rundown (like, dislike event)
* Student Information
* Account Information
* Qr Code for Authentication (Displaying or Scanning)
    * Login to Website
    * Sign up for an event
 
### Technical Overview
* Server (Golang, Beego)
* Cache (Redis)
* Database (MySQL)
* Mobile Application (Swift, SwiftUI)
* Website (HTML, CSS, JS)

#### Architecture Chart
##### Overall
![](/display/overall_architecture.png)

##### Using Device to login on Browser
![](/display/lazy-version.png)
![](/display/coolthing.png)

### Research Summary
* Beego (the framework for server)
    1. https://beego.vip
    2. https://github.com/astaxie/beego

* Bee (tool for managing the Beego framework and project)
    https://github.com/astaxie/bee

* SwiftUI (the language for the application)
    1. https://www.hackingwithswift.com (cover most of the things in the application: Qr Code, updating views, etc)
    2. https://github.com/twostraws (the author of the website above, **Paul Hudson**)
    3. https://github.com/twostraws/CodeScanner (for the pop up (slide over) QR Code scanner)

* Idea for the events display
    https://github.com/Ramotion/expanding-collection


### Further Work
* Performance
    * Database
        * Reconstructe Tables
            * keys, indexes, foreign keys, grants
    * Cache
        * Pre-load data
            * events and others command static data
    * Server
        * Refactor controllers(handlers)
        * Organize router(path)
    * Application
        * refactor the flow of views
        * let the server process and finalize data instead do it on fly
    * Website
        * using WebSocket for the communication in the QR code login step

* Faculty Version on the application
add other user types' versions with functions like having the event QR code on mobile devices.

* Security
    * Using other more secure authorization like OAuth?
    * Shorten the expiry time for tokens
    * Using different tokens for the QR token image and actual token
    * Encrypting Information in communication

* UI/UX
    * making it better and nicer
