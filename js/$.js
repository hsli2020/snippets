/* eslint-disable no-multi-assign */
const $ = document.querySelector.bind(document);
const $$ = document.querySelectorAll.bind(document);

Node.prototype.on = window.on = function(name, fn) {
    this.addEventListener(name, fn);
};

// eslint-disable-next-line no-proto
NodeList.prototype.__proto__ = Array.prototype;

NodeList.prototype.on = NodeList.prototype.addEventListener = function(name, fn) {
    this.forEach(function(elem, i) {
        elem.on(name, fn);
    });
};

export { $, $$ };


// prevent enter key press
document.on('keypress', e => {
    const keyCode = e.which || e.keyCode;
    if (keyCode === 13) {
        e.preventDefault();
    }
});