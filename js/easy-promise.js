// https://blog.csdn.net/linminghe/article/details/80491745

function easyPromise (fn) {
    var that = this

    // 第一步,定义 then()
    this.then = function (cb) {
        //先将 then() 括号里面的参数(回调函数)保存起来
        that.cb = cb
    }

    // 定义一个 resolve
    this.resolve = function(data) {
        that.cb(data)
    }

    // 将 resolve 作为回调函数,传给fn
    fn(this.resolve)
}

/*
function easyPromise (fn) {
    this.then = cb => this.cb = cb
    this.resolve = data => this.cb(data)
    fn(this.resolve)
}*/

new easyPromise((resolve) => {
    setTimeout(() => {
        resolve("延时执行")
    }, 1000)
}).then((data) => {
    console.log(data)
})
