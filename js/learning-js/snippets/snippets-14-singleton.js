////////////////////////////////////////////////////////////

function mockMask() {
    pr("mockMask()");
    return { "mask": "created" }
    //return document.body.appendChild(document.createElement('div'));
}

var createMask = function() { // singleton
    var mask;
    return function() {
        return mask || (mask = mockMask());
    }
}();

var mask1 = createMask();
var mask2 = createMask();
//pr(mask1 == mask2);

////////////////////////////////////////////////////////////

var singleton = function(fn) {
    var result;
    return function() {
        return result || (result = fn.apply(this, arguments));
    }
}

var createMask = singleton(function() {
    return mockMask();
})

//var mask1 = createMask();
//var mask2 = createMask();
//pr(mask1 == mask2);

////////////////////////////////////////////////////////////

function once(fn, context) {
    var result;

    return function() {
        if (fn) {
            result = fn.apply(context || this, arguments);
            fn = null;
        }

        return result;
    };
}

// 用法
var canOnlyFireOnce = once(function() {
    console.log('Fired!');
});

//canOnlyFireOnce(); // "Fired!"
//canOnlyFireOnce(); // 没有执行指定函数

////////////////////////////////////////////////////////////
