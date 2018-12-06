// 函数防抖实现
function debounce(fn, delay) {
    let timer = null;
    return function () {
        if (timer) clearTimeout(timer);
        timer = setTimeout(() => {
            fn.apply(this, arguments);
        }, delay);
    }
}

// 函数节流实现
function throttle(fn, cycle) {
    let start = Date.now();
    let now;
    let timer;
    return function () {
        now = Date.now();
        clearTimeout(timer);
        if (now - start >= cycle) {
            fn.apply(this, arguments);
            start = now;
        } else {
            timer = setTimeout(() => {
                fn.apply(this, arguments);
            }, cycle);
        }
    }
}

//lodash对这两个方法的实现更灵活一些，有第三个参数，可以去参观学习。
