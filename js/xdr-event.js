// https://github.com/remy/libraries

// xhr.js

function request(type, url, opts, callback) {
  var xhr = new XMLHttpRequest(),
      pd;

  if (typeof opts === 'function') {
    callback = opts;
    opts = null;
  }

  xhr.open(type, url);

  if (type === 'POST' && opts) {
    pd = JSON.stringify(opts);

    xhr.setRequestHeader('Content-Type', 'application/json');
  }

  xhr.setRequestHeader('X-Requested-With', 'XMLHttpRequest');

  xhr.onload = function () {
    callback.call(xhr, null, JSON.parse(xhr.response));
  };

  xhr.onerror = function () {
    callback.call(xhr, true);
  };

  xhr.send(opts ? pd : null);

  return xhr;
}

var get = request.bind(this, 'GET');
var post = request.bind(this, 'POST');

// events.js

var ev = (function (window, document, undefined) {
  function each(el, fn) {
    if (el && el.nodeName || el === window) {
      fn(el);
    } else if (el && el.length) {
      for (var i = 0, l = el.length; i < l; i++) {
        fn(el[i]);
      }
    }
  }
  
  var handlers = {},
      guid = 0,
      expando = 'e' + new Date;
  
  var addEvent = (function () {
    if (document.addEventListener) {
      return function (el, type, fn) {
        each(el, function (el) {
          var ret = el.addEventListener(type, function (event) {
            if (fn.call(el, event) === false) {
              event.preventDefault();
              event.stopPropagation();
            }
          }, false);
        });
      };
    } else {
      return function (el, type, fn) {
        each(el, function (el) {
          if (el['on' + type] === null) {
            el.attachEvent('on' + type, function () { 
              return fn.call(el, window.event); 
            });
            return;            
          }
          
          if (handlers[type] === undefined) {
            handlers[type] = {};
          }
          
          var id = el[expando] = guid++;
          
          if (handlers[type][id] !== undefined) {
            var old = handlers[type][id];
            handlers[type][id] = function (event) {
              if (fn.call(el, event) !== false) {
                return old();
              }
              return false;
            };
          } else {
            handlers[type][id] = function (event) {
              return fn.call(el, event);
            };            
          }
          
          el = null; // <3 IE
        });
      };
    }
  })();
  
  var fireEvent = (function () {
    if (document.createEvent) {
      return function (el, type, data) {
        each(el, function (el) {
          var event = document.createEvent('HTMLEvents');
          event.initEvent(type, true, true);
          event.data = data || {};
          event.eventName = type;

          el.dispatchEvent(event);          
        });
      };
    } else {
      return function (el, type, data) {
        each(el, function (el) {
          try {
            el['on' + type].call(el, data);
          } catch (e) {
            var id;
            do {
              id = el[expando];
              if (!el || el.nodeType === 3 || el.nodeType === 8) {
                return; // we're done - this shouldn't happen
              } else {
                handlers[type][id] && handlers[type][id].call(el, { type: type, data : data || {} });
              }
            } while (el = el.parentNode)
          }
        });
      };
    }
  })();
  
  return function (el) {
    if (this === window) {
      return new ev(el);
    }
    this.el = el;
    this.addEvent = addEvent;
    this.fireEvent = fireEvent;
  };
})(this, document);

ev.prototype = {
  on: function (type, fn) {
    if (fn === undefined) {
      for (var t in type) {
        this.addEvent(this.el, t, type[t]);
      }
    } else {
      this.addEvent(this.el, type, fn);
    }
    
    return this;
  },
  fire: function (type) {
    this.fireEvent(this.el, type);
    return this;
  }
};
