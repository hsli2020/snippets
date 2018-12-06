var sprintf = function (str) {
    var args = arguments,
        flag = true,
        i = 1;
    str = str.replace(/%s/g, function () {
        var arg = args[i++];
        if (typeof arg === 'undefined') {
            flag = false;
            return '';
        }
        return arg;
    });
    return flag ? str : '';
};

console.log(sprintf("The page number is %s", 12));
