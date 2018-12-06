// https://github.com/remy/min.js

// $.js

/*globals Node:true, NodeList:true*/
$ = (function (document, window, $) {
  // Node covers all elements, but also the document objects
  var node = Node.prototype,
      nodeList = NodeList.prototype,
      forEach = 'forEach',
      trigger = 'trigger',
      each = [][forEach],
      // note: createElement requires a string in Firefox
      dummy = document.createElement('i');

  nodeList[forEach] = each;

  // we have to explicitly add a window.on as it's not included
  // in the Node object.
  window.on = node.on = function (event, fn) {
    this.addEventListener(event, fn, false);

    // allow for chaining
    return this;
  };

  nodeList.on = function (event, fn) {
    this[forEach](function (el) {
      el.on(event, fn);
    });
    return this;
  };

  // we save a few bytes (but none really in compression)
  // by using [trigger] - really it's for consistency in the
  // source code.
  window[trigger] = node[trigger] = function (type, data) {
    // construct an HTML event. This could have
    // been a real custom event
    var event = document.createEvent('HTMLEvents');
    event.initEvent(type, true, true);
    event.data = data || {};
    event.eventName = type;
    event.target = this;
    this.dispatchEvent(event);
    return this;
  };

  nodeList[trigger] = function (event) {
    this[forEach](function (el) {
      el[trigger](event);
    });
    return this;
  };

  $ = function (s) {
    // querySelectorAll requires a string with a length
    // otherwise it throws an exception
    var r = document.querySelectorAll(s || '☺'),
        length = r.length;
    // if we have a single element, just return that.
    // if there's no matched elements, return a nodeList to chain from
    // else return the NodeList collection from qSA
    return length == 1 ? r[0] : r;
  };

  // $.on and $.trigger allow for pub/sub type global
  // custom events.
  $.on = node.on.bind(dummy);
  $[trigger] = node[trigger].bind(dummy);

  return $;
})(document, this);

// delegate.js

// usage: $('body').delegate('li > a', 'click', fn);

Node.prototype.delegate = function (selector, event, fn) {
  var matches = this.mozMatchesSelector || this.webkitMatchesSelector || 
                this.oMatchesSelector || this.matchesSelector || (function (selector) {
    // support IE10 (basically)
    var target = this,
        elements = $(selector),
        match = false;
    if (elements instanceof NodeList) {
      elements.forEach(function (el) {
        if (el === target) match = true;
      });
    } else if (elements === target) {
      match = true;
    }

    return match;
  });

  this.on(event, function (event) {
    if (matches.call(event.target, selector)) {
      fn.call(event.target, event);
    }
  });

  return this;
};

# min.js
========

A super tiny JavaScript library to execute simple DOM querying and hooking event listeners.
Aims to return the raw DOM node for you to manipulate directly, using HTML5 (et al) tech 
like element.classList or element.innerHTML, etc.

# Query elements

var links = $('p:first-child a');

If there is more than one link, the return value is NodeList, if there's only a single match,
you have an Element object. So you need to have an idea of what to expect if you want to 
modify the DOM.

# Events

# Bind events

$('p:first-child a').on('click', function (event) {
  event.preventDefault();
  // do something else
});

Note: the on and trigger methods are on both Node objects and NodeList objects, 
which also means this affects the document node, so document.on(type, callback) will also work.

# Custom events

$('a').on('foo', function () {
  // foo was fired
});

$('a:first-child').trigger('foo');

# Arbitrary events / pubsub

$.on('foo', function () {
  // foo was fired, but doesn't require a selector
});

# Turning off events?

Current min.js has no support for turning off events (beyond .removeEventListener -- but 
even then you don't have the reference function to work with). Currently there's no plans to 
implement this (as I find I don't disable events very often at all) -- but I'm not closed to
the idea. There's an issue open, but it adds quite a bit more logic to a very small file. If 
there's enough 👍 on the issue, I'll add it in. Equally, if you think min.js should stay simple,
please 👎 -- this is useful too.

# Looping

$('p').forEach(function (el, index) {
  console.log(el.innerHTML);
});

Note: jQuery-like libraries tend to make the context this the element. Since we're borrowing 
forEach from the array object, this does not refer to the element.

# Chaining events

$('a').on('foo', bar).on('click', doclick).trigger('foobar');

Also when a single element is matched, you have access to it:

$('a').href = '/some-place.html';

# Silent failing

Like jQuery, this tiny library silently fails when it doesn't match any elements.
As you might expect.
