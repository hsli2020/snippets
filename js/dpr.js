const DEBUG = 1;

function dpr(...args) {
    if (DEBUG) {
        console.log(args);
    }
}

var a = 1, b = 2, c = [ 'a', 'b', 'c' ];
dpr(a, b, c);
