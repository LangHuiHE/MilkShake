# MilkShake

**MilkShake** is my senior project for my bachelor's degree in Computer Science at Utah Tech University (Dixie State University). 

Also, it's the first project involving different technologies, tools, and platforms I have made, and it's fully designed and developed by myself. 


## Deployment
* MilkJug
    1. Create Database
    Using sql files at `setup/database` to create Database
    **All records are generated randomly**
    2. Cache: Redis

    Don't forget update config setting at `milkjug/conf/app.conf`
* Shaker
    1. Update server address at `shaker/conf.swift`

## Demo Screenshot at `display`


## Problem Description
Students and facilities in Dixie State University (Utah Tech University) couldn't find a good mobile application for their campus life. The applications developed or published by the university are old, which means most of them had the last update more than a year ago. Meanwhile, those applications don't have the functionality and information that are interesting to students and facilities.

## Solution Description
The purpose of this project is to empower every student & staff to become all they can be with a great mobile application. The application makes the hectic student life more structured and ensures students and staff know what is happening around campus.

####  Feature of Solution
* Event Rundown (like, dislike event)
* Student Information
* Account Information
* Qr Code for Authentication (Displaying or Scanning)
    * Login on Website
    * Sign up for event
 
### Technical Overview
* Server (Golang, Beego)
* Cache (Redis)
* Database (MySQL)
* Mobile Application (Swift, SwiftUI)
* Website (HTML, CSS, JS)

#### Overall Architecture
![](/display/overall_architecture.png)

#### Using Device to login on Browser Architecture
![](/display/coolthing.png)

### Research Summary
* Beego (the framework for server)
    1. https://beego.vip
    2. https://github.com/astaxie/beego

* Bee (tool for managing beego framework and project)
    https://github.com/astaxie/bee

* SwiftUI (the language for the application)
    1. https://www.hackingwithswift.com (cover basically most of the thing in the application: Qr Code, updating views, etc)
    2. https://github.com/twostraws (the author of the website above, **Paul Hudson**)
    3. https://github.com/twostraws/CodeScanner (for the pop up (slide over) QR Code scanner)

* Idea for the events display
    https://github.com/Ramotion/expanding-collection


### Further Work
* Performance
    * Database
        * Re-construction Tables
            * keys, indexes, foreign keys, grants
    * Cache
        * pre-load data
            * events and others command static information
    * Server
        * refactor controllers(handlers)
        * organize router(path)
    * Application
        * refactor the flow of views
        * let server to process and finalize data instead do it on fly
    * Website
        * using websocket for QR code login

* Facilities Version on App
add other user types' version on the App. Like having the event QR code on the mobile device.

* Security
    * using other more secure authorization like OAuth
    * shorten the expiry time for tokens
    * using different tokens for the QR token image and actual token
    * encrypting Information on communication

* UI/UX
    * making it better and nicer
