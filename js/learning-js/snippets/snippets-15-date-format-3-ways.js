function pr(d) { console.log(d); }

function pad(s) {
    return ('0' + s).slice(-2);
}

// for old browser
var dt = new Date();
var date = dt.getFullYear() + '-' + pad(dt.getMonth() + 1) + '-' + pad(dt.getDate());
date += ' ';
date += pad(dt.getHours()) + ':' + pad(dt.getMinutes()) + ':' + pad(dt.getSeconds());
//pr(date);


// better way
var dt = new Date();
var date = [
    [ dt.getFullYear(), dt.getMonth() + 1, dt.getDate() ].join('-'),
    [ dt.getHours(), dt.getMinutes(), dt.getSeconds()].join(':')
].join(' ').replace(/(?=\b\d\b)/g, '0');
//pr(date);


// for modern browser
var dt = new Date();
dt.setMinutes(dt.getMinutes() - dt.getTimezoneOffset());
var date = dt.toISOString().slice(0, -5).replace(/[T]/g, ' ');
//pr(date);
