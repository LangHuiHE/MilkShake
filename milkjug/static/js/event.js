// var button = document.querySelector("#submit")
// console.log("this is button", button);

// button.onclick = function () {
//     var name = document.querySelector("#input-name");
//     console.log("check", name);
//     var age = document.querySelector("#input-age");
//     console.log("check", age);
//     var guest = document.querySelector("#input-guest");
//     console.log("check", guest);
//     createNewTicket(name.value, age.value, guest.value);

//     name.value = "";
//     age.value = "";
//     guest.value = "";
// };

// function createNewTicket(name, age, guest) {
//     var data = "name=" + encodeURIComponent(name);
//     data += "&age=" + encodeURIComponent(age);
//     data += "&guest=" + encodeURIComponent(guest);

//     fetch("http://localhost:8080/ticket", {
//         method: "POST",
//         body: data,
//         credentials: "include",
//         headers: {
//             "Content-Type": "application/x-www-form-urlencoded"
//         }
//     }).then(function (response) {
//         // console.log(response)
//         if (response.status == 403) {
//             response.text().then(function (msg) {
//                 alert(msg);
//             })

//         } else {
//             console.log("This worked! Reload the list.");
//             loadTheTicket();
//         }
//     });
// }



function loadTheEvents() {
    var theList = document.querySelector("#the-list");
    // console.log(lst)

    var d = new Date();
    var today = d.getDate();

    // console.log("today is ", today);
    theList.innerHTML = "";

    fetch("http://localhost:8080/mobile/events", {
        method: "GET",
        credentials: "include",
    }).then(function (response) {
        // console.log(response)

        if (response.status != 200) {
            response.text().then(function (msg) {
                console.log(msg);
            })

        }
        else {
            response.json().then(function (data) {
                lst = data;
                // console.log("list loaded form server: ", lst);

                data.forEach(function (t) {
                    var newID = document.createElement("h3");
                    var newTitle = document.createElement("h3");
                    var newDate = document.createElement("h3");

                    newID.innerHTML = t.Id;
                    newID.className = "column";

                    newTitle.innerHTML = t.Title;
                    newTitle.className = "column";

                    var ed = new Date(t.StartDate);
                    newDate.innerHTML = ed.toLocaleString();
                    newDate.className = "column";
                    // console.log(ed.toLocaleString());

                    if (today == ed.getDate()) {
                        newID.style.background = "#FFD700";
                        newTitle.style.background = "#FFD700";
                        newDate.style.background = "#FFD700";
                    }

                    var selectButton = document.createElement("button");
                    selectButton.innerHTML = "X";
                    selectButton.onclick = function () {
                        console.log("selected", t.Id, t.Title);
                        var website = document.getElementById("website");
                        console.log(website)
                        website.value = t.Id + "%20" + t.Title;
                        generateQRCode();
                    };
                    selectButton.className = "column";

                    var container = document.createElement("div");
                    container.appendChild(newID);
                    container.appendChild(newTitle);
                    container.appendChild(newDate);
                    container.appendChild(selectButton);

                    var newEvent = document.createElement("li");
                    newEvent.appendChild(container);
                    theList.appendChild(newEvent);
                });
            })
        }
    });
}

loadTheEvents();