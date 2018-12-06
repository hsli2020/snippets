function pr(d) { console.log(d); }

function objectKeys(input) {
    var output = [];

    if (input) {
        // Loop all keys in input.
        for (var key in input) {
            if (input.hasOwnProperty(key)) {
                output.push(key);
            }
        }
    }

    return output;
}

function objectValues(input) {
    var values = [];

    for (var key in input) {
        if (input.hasOwnProperty(key)) {
            values.push(input[key]);
        }
    }

    return values;
}


