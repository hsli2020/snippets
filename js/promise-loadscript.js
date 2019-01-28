// Here's the callback-based variant:

function loadScript(src, callback) {
    let script = document.createElement('script');
    script.src = src;

    script.onload = () => callback(null, script);
    script.onerror = () => callback(new Error(`Script load error ` + src));

    document.head.append(script);
}

// Let?s rewrite it using Promises.

// The new function loadScript will not require a callback. Instead, it will create and return
// a Promise object that resolves when the loading is complete. The outer code can add handlers
// (subscribing functions) to it using .then:

function loadScript(src) {
    return new Promise(function(resolve, reject) {
        let script = document.createElement('script');
        script.src = src;

        script.onload = () => resolve(script);
        script.onerror = () => reject(new Error("Script load error: " + src));

        document.head.append(script);
    });
}

// Usage:

let promise = loadScript("https://cdnjs.cloudflare.com/ajax/libs/lodash.js/3.2.0/lodash.js");

promise.then(
    script => alert(`${script.src} is loaded!`),
    error => alert(`Error: ${error.message}`)
);

promise.then(script => alert('One more handler to do something else!'));
