let start = Date.now();
let obj = {};
for (var i = 0; i< 1000000; i++) {
    var time = Date.now();
    obj[i +'_' + time] = time;
}
console.log((Date.now() - start) / 1000);
