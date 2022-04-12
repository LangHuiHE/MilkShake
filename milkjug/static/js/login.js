var baseURL = "http://localhost:8080/web"

var form = document.getElementById('login-form');
form.onsubmit = function (event) {
    var xhr = new XMLHttpRequest();
    var formData = new FormData(form);

    console.log(data)

    //open the request
    var url = baseURL + "/login"
    xhr.open('POST', url)
    xhr.setRequestHeader("Content-Type", "application/json");

    var data = JSON.stringify(Object.fromEntries(formData))

    console.log(data)

    //send the form data
    xhr.send(JSON.stringify(Object.fromEntries(formData)));

    xhr.onreadystatechange = function () {
        if (xhr.readyState == XMLHttpRequest.DONE) {
            form.reset(); //reset form after AJAX success or do something else
            location.reload();
        }
    }
    //Fail the onsubmit to avoid page refresh.
    return false;
}


var qrToken
var in_progress = false

var showQrCodeImagButton = document.getElementById("showQrCodeImage");
showQrCodeImagButton.onclick = function () {
    fetch(baseURL + "/loginQrCode")
        .then(res => {
            qrToken = res.headers.get("QrToken")
            return res.json()
        })
        .then(result => {
            // console.log(result)
            if (qrToken != "") {
                var img = document.getElementById('QrCodeImage')
                img.src = "data:image/png;base64, " + result;
                img.style.display = "block";
                showQrCodeImagButton.style.display = "None";

                in_progress = true
                checkStatus(qrToken)

                setTimeout(function () {
                    in_progress = false
                    img.style.display = "None";
                    img.src = "";
                    qrToken = "";
                    showQrCodeImagButton.style.display = "block";

                }, 10000); // 50000
            }
        })
        .catch(err => console.log(err))
}

const checkStatus = async (token) => {
    if (in_progress) {
        try {
            const response = await fetch(baseURL + "/loginQrCode", {
                method: 'GET',
                headers: {
                    'QrToken': token,
                },
            });
            if (response.status === 200) {
                console.log("successfully get jwt.");
                in_progress = false;
                location.reload();
                return;
            }
        } catch (err) {
            // catches errors both in fetch and response.json
            console.log(err);
        } finally {
            // do it again in 1 seconds
            // setTimeout(checkStatus(token), 1000);
            setTimeout(function () { checkStatus(token); }, 250);
        }
    }
};