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
        if (email == "") {
            alert("fill out correctly")
        } else if (password == "") {
            alert("fill out correctly")
        }
    } else if (email != "") {
        if (username == "") {
            alert("fill out correctly")
        } else if (password == "") {
            alert("fill out correctly")
        }
    } else if (password != "") {
        if (username == "") {
            alert("fill out correctly")
        } else if (email == "") {
            alert("fill out correctly")
        }
    }
}

function getAndInsertData(data, username, email, password) {
    let gotUsername = document.getElementById(username);
    let gotEmail = document.getElementById(email);
    let gotPassword = document.getElementById(password);

    let isFilled = checkUnfilled(username, email, password);

    if (isFilled == true) {
        var userData = {
            userName: gotUsername,
            mailAddr: gotEmail,
            passWord: gotPassword,
        }
        data.push(userData);
    }

    return data
}