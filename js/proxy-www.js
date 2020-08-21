// https://github.com/justjavac/proxy-www
//
// 学会 Proxy 就可以为所欲为吗？
// 对，学会 Proxy 就可以为所欲为！

const www = new Proxy(new URL('https://www'), {
    get: function get(target, prop) {
        let o = Reflect.get(target, prop);
        if (typeof o === 'function') {
            return o.bind(target)
        }
        if (typeof prop !== 'string') {
            return o;
        }
        if (prop === 'then') {
            return Promise.prototype.then.bind(fetch(target));
        }
        target = new URL(target);
        target.hostname += `.${prop}`;
        return new Proxy(target, { get });
    }
});

// 访问百度
www.baidu.com.then(response => {
    console.log(response.status);
    // ==> 200
})
