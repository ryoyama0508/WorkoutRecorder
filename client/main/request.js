function postRequestStay(url) {

    var xmlHttp = new XMLHttpRequest();//connection

    xmlHttp.onreadystatechange = () => {//func that check if the server over there is ready or not
        if (xmlHttp.readyState == 4 && xmlHttp.status == 200) {
            //onreadystatechangeで「読み込み途中⁠」⁠，および「読み込み完了⁠」⁠・「⁠失敗」時の処理を行います。
            //まず，読み込みの状態を管理するのはreadyStateプロパティです。readyStateが4の時，読み込みの完了を示します。

            //read the text right here
            document.getElementById("demo").innerHTML = xhttp.responseText;
        };
    };

    xmlHttp.open("GET", url, true); // true for asynchronous

    xmlHttp.send(null);
}

function sendJSON() {

    let result = document.querySelector('.result');
    let name = document.querySelector('#name');
    let email = document.querySelector('#email');

    // Creating a XHR object 
    let xhr = new XMLHttpRequest();
    let url = "submit.php";

    // open a connection 
    xhr.open("POST", url, true);

    // Set the request header i.e. which type of content you are sending 
    xhr.setRequestHeader("Content-Type", "application/json");

    // Create a state change callback 
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {

            // Print received data from server 
            result.innerHTML = this.responseText;

        }
    };

    // Converting JSON data to string 
    var data = JSON.stringify({ "name": name.value, "email": email.value });

    // Sending data with the request 
    xhr.send(data);
} 