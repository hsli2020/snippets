function $id(sel)  { return document.getElementById(sel) }

function $(sel)    { return document.querySelector(sel) }
function $$(sel)   { return document.querySelectorAll(sel) }

function $e(sel)   { return document.querySelector(sel) }
function $all(sel) { return document.querySelectorAll(sel) }

var $ = document;

$.getElementById("demo1").addEventListener('click', function(){ alert('click'); });

console.log("Height: " + $.body.clientHeight);
console.log("Width: "  + $.body.clientWidth);

// Returns first element that matches CSS selector {expr}.
// Querying can optionally be restricted to {container}’s descendants
function $(expr, container) {
	return typeof expr === "string"? (container || document).querySelector(expr) : expr || null;
}

// Returns all elements that match CSS selector {expr} as an array.
// Querying can optionally be restricted to {container}’s descendants
function $$(expr, container) {
	return [].slice.call((container || document).querySelectorAll(expr));
}

function fadeIn(el) {
  el.style.opacity = 0;

  var last = +new Date();
  var tick = function() {
    el.style.opacity = +el.style.opacity + (new Date() - last) / 400;
    last = +new Date();

    if (+el.style.opacity < 1) {
      (window.requestAnimationFrame && requestAnimationFrame(tick)) || setTimeout(tick, 16);
    }
  };

  tick();
}

function outerWidth(el) {
  var width = el.offsetWidth;
  var style = getComputedStyle(el);

  width += parseInt(style.marginLeft) + parseInt(style.marginRight);
  return width;
}

function ready(fn) {
  if (document.attachEvent ? document.readyState === "complete" : document.readyState !== "loading"){
    fn();
  } else {
    document.addEventListener('DOMContentLoaded', fn);
  }
}

var deepExtend = function(out) {
  out = out || {};

  for (var i = 1; i < arguments.length; i++) {
    var obj = arguments[i];

    if (!obj)
      continue;

    for (var key in obj) {
      if (obj.hasOwnProperty(key)) {
        if (typeof obj[key] === 'object')
          out[key] = deepExtend(out[key], obj[key]);
        else
          out[key] = obj[key];
      }
    }
  }

  return out;
};

// deepExtend({}, objA, objB);

var extend = function(out) {
  out = out || {};

  for (var i = 1; i < arguments.length; i++) {
    if (!arguments[i])
      continue;

    for (var key in arguments[i]) {
      if (arguments[i].hasOwnProperty(key))
        out[key] = arguments[i][key];
    }
  }

  return out;
};

// extend({}, objA, objB);

/* bling.js */

window.$ = document.querySelectorAll.bind(document);

Node.prototype.on = window.on = function (name, fn) {
  this.addEventListener(name, fn);
}

NodeList.prototype.__proto__ = Array.prototype;

NodeList.prototype.on = NodeList.prototype.addEventListener = function (name, fn) {
  this.forEach(function (elem, i) {
    elem.on(name, fn);
  });
}

// comment to bling.js
Node.prototype.on = window.on = function (name, delegate, fn) {
  if(arguments.length !== 3) {
    return this.addEventListener(name, arguments[1]);
  }

  return this.addEventListener(name, function (e) {
    if(e.target.matches(delegate)){
      return fn.apply(e.target, arguments);
    }
  })
}

window.$.Deferred = function() {
  var resolver, rejector;
  var promise = new Promise(function(resolve, reject) {
    resolver = resolve;
    rejector = reject;
  });
  var deferred = {
    resolve: resolver,
    reject: rejector,
    then: function(a, b) { promise = promise.then(a, b); return deferred; },
    pipe: function(a, b) { promise = promise.then(a, b); return deferred; },
    done: function(a) { promise = promise.then(a, b); return deferred; },
    fail: function(a) { promise = promise.catch(a); return deferred; },
    always: function(a) { promise = promise.finally(a); return deferred; },
    promise: function() { return deferred; }
  };
  return deferred;
};

Node.prototype.on = window.on = function(names, fn) {
  var self = this

  names.split(' ').forEach(function(name) {
    self.addEventListener(name, fn)
  })

  return this
}

NodeList.prototype.on = NodeList.prototype.addEventListener = function(names, fn) {
  this.forEach(function(elem) {
    elem.on(names, fn)
  })

  return this
}

function simulateClick (elem) {
  var evt = document.createEvent('MouseEvents')
  evt.initMouseEvent('click', true, true, window, 0, 0, 0, 0, 0, false, false, false, false, 0, null)
  var canceled = !elem.dispatchEvent(evt)
}

// simulateClick(input1)
// setTimeout(function () { simulateClick(input2)  }, 500)

// VanillaJs
function docReady(cb) {
  if (document.readyState != 'loading'){
    cb(); 
  } else {
    document.addEventListener('DOMContentLoaded', cb);
  }
}
