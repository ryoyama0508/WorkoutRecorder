function postRecord(
    url,
    cr_reps, cr_sets,
    dl_weight, dl_reps, dl_sets,
    chup_reps, chup_sets,
    dc_weight, dc_reps, dc_sets,
    bp_weight, bp_reps, bp_sets,
    ht_weight, ht_reps, ht_sets,
    sq_weight, sq_reps, sq_sets,
    shpre_weight, shpre_reps, shpre_sets,
    cablepre_weight, cablepre_reps, cablepre_sets,
) {
    data = []

    getAndInsertData(data, "", cr_reps, cr_sets, "crunch")

    getAndInsertData(data, dl_weight, dl_reps, dl_sets, "dead lift")

    getAndInsertData(data, "", chup_reps, chup_sets, "chin up")

    getAndInsertData(data, dc_weight, dc_reps, dc_sets, "dumbell curl")

    getAndInsertData(data, bp_weight, bp_reps, bp_sets, "bench press")

    getAndInsertData(data, ht_weight, ht_reps, ht_sets, "hip thrust")

    getAndInsertData(data, sq_weight, sq_reps, sq_sets, "squat")

    getAndInsertData(data, shpre_weight, shpre_reps, shpre_sets, "shoulder press")

    getAndInsertData(data, cablepre_weight, cablepre_reps, cablepre_sets, "cable press down")


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
    var json = JSON.stringify(data);

    // Sending data with the request 
    xhr.send(json);
}

function checkUnfilled(weight, sets, reps) {
    if (weight != "") {
        if (weight != 0) {
            if (sets == 0) {
                alert("fill out correctly")
            } else if (reps == 0) {
                alert("fill out correctly")
            }
        } else if (sets != 0) {
            if (weight == 0) {
                alert("fill out correctly")
            } else if (reps == 0) {
                alert("fill out correctly")
            }
        } else if (reps != 0) {
            if (weight == 0) {
                alert("fill out correctly")
            } else if (sets == 0) {
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

function nothingOrFilled(weight, sets, reps) {
    if (weight != "") {
        if (weight == 0 && sets == 0 && reps == 0) {
            return false
        }
        if (weight != 0 && sets != 0 && reps != 0) {
            return true
        }
    } else {
        if (sets == 0 && reps == 0) {
            return false
        }
        if (sets != 0 && reps != 0) {
            return true
        }
    }
}

function insertData(
    isFilled, data,
    name, weight, sets, reps
) {
    if (isFilled == true && weight != "") {
        var datum = {
            exercise: name,
            weight: weight,
            reps: reps,
            sets: sets
        }
        data.push(datum);
    } else if (isFilled == true && weight == "") {
        var datum = {
            exercise: name,
            reps: reps,
            sets: sets
        }
        data.push(datum);
    }
}

function getAndInsertData(data, argweight, argreps, argsets, exercise) {
    let reps = document.getElementById(argreps);
    let sets = document.getElementById(argsets);

    if (argweight != "") {
        let weight = document.getElementById(argweight);
        checkUnfilled(weight.value, sets.value, reps.value)
        var isFilled = nothingOrFilled(weight.value, sets.value, reps.value)
        data = insertData(isFilled, data, exercise, weight.value, sets.value, reps.value)
    } else {
        checkUnfilled("", reps.value, sets.value)
        var isFilled = nothingOrFilled("", sets.value, reps.value)
        data = insertData(isFilled, data, exercise, "", sets.value, reps.value)
    }
    return data
}