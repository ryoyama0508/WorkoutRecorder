function test(arg1, arg2) {
    var val1 = document.getElementById(arg1);
    var val2 = document.getElementById(arg2);
    console.log(val1.value, val2.value)
}

function testRecord(
    cr_reps, cr_sets,
) {
    var data = []
    getAndInsertData(data, "", cr_reps, cr_sets, "crunch")
    /* let reps = document.getElementById(cr_reps);
    let sets = document.getElementById(cr_sets);

    checkUnfilled("", reps.value, sets.value)

    var isFilled = nothingOrFilled("", reps.value, sets.value)

    if (isFilled == true) {
        var datum = {
            exercise: "crunch",
            reps: reps.value,
            sets: sets.value
        }
        data.push(datum);
    } */

    console.log(data)
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
    return data
}

function getAndInsertData(data, argweight, argsets, argreps, exercise) {
    let sets = document.getElementById(argsets);
    let reps = document.getElementById(argreps);

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