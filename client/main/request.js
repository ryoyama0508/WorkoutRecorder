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
