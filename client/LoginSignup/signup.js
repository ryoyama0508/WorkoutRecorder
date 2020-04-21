function httpGetAsync(theUrl, callback) {

    var xmlHttp = new XMLHttpRequest();//connection

    xmlHttp.open("GET", theUrl, true); // true for asynchronous

    xmlHttp.onreadystatechange = () => {//func that check if the server over there is ready or not
        if (xmlHttp.readyState == 4 && xmlHttp.status == 200)
            //onreadystatechangeで「読み込み途中⁠」⁠，および「読み込み完了⁠」⁠・「⁠失敗」時の処理を行います。
            //まず，読み込みの状態を管理するのはreadyStateプロパティです。readyStateが4の時，読み込みの完了を示します。
            callback(xmlHttp.responseText);//read the text right here
    }

    xmlHttp.send(null);
}