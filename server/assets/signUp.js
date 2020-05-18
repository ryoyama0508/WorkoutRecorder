function postRecord(
    url,
    username,
    email,
    password,
) {
    data = []

    getAndInsertData(data, username, email, password)

    // Creating a XHR object 
    let xhr = new XMLHttpRequest();

    // open a connection 
    xhr.open("POST", url, true);

    // Set the request header i.e. which type of content you are sending 
    xhr.setRequestHeader("Content-Type", "application/json");

    //erase

    // Converting JSON data to string 
    var json = JSON.stringify(data);

    console.log(json)

    // Sending data with the request 
    xhr.send(json);
}

function checkUnfilled(username, email, password) {
    if (username != "") {
        if (email != "") {
            if (password == "") {
                alert("fill out correctly")
            }
        } else if (password != "") {
            if (email == 0) {
                alert("fill out correctly")
            }
        }
    } else {
        if (sets != 0) {
            if (reps == 0) {
                alert("fill out correctly")
            }
        } else if (reps != 0) {
            if (sets == 0) {
                alert("fill out correctly")
            }
        }
    }
}

function getAndInsertData(data, username, email, password) {
    let gotUsername = document.getElementById(username);
    let gotEmail = document.getElementById(email);
    let gotPassword = document.getElementById(password);

    var userData = {
        name: gotUsername,
        mailaddr: gotEmail,
        pw: gotPassword,
    }
    data.push(userData);

    return data
}