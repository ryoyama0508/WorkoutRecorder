function postRecord(
    url,
    crunch_reps, crunch_sets,
    dl_weight, dl_reps, dl_sets,
    cu_weight, cu_reps, cu_sets,
    dc_weight, dc_reps, dc_sets,
    bp_weight, bp_reps, bp_sets,
    ht_weight, ht_reps, ht_sets,
    sq_weight, sq_reps, sq_sets,
    sp_weight, sp_reps, sp_sets,
    cp_weight, cp_reps, cp_sets,
) {

    let result = document.querySelector('.result');
    let name = document.querySelector('#name');
    let email = document.querySelector('#email');

    // Creating a XHR object 
    let xhr = new XMLHttpRequest();

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