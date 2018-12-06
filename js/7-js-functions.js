7 Essential JavaScript Functions
================================

https://davidwalsh.name/essential-javascript-functions

I remember the early days of JavaScript where you needed a simple function for just about 
everything because the browser vendors implemented features differently, and not just edge 
features, basic features, like addEventListener and attachEvent.  Times have changed but 
there are still a few functions each developer should have in their arsenal, for performance 
for functional ease purposes.

debounce
--------

The debounce function can be a game-changer when it comes to event-fueled performance.  If 
you aren't using a debouncing function with a scroll, resize, key* event, you're probably 
doing it wrong.  Here's a debounce function to keep your code efficient:

// Returns a function, that, as long as it continues to be invoked, will not
// be triggered. The function will be called after it stops being called for
// N milliseconds. If `immediate` is passed, trigger the function on the
// leading edge, instead of the trailing.
function debounce(func, wait, immediate) {
	var timeout;
	return function() {
		var context = this, args = arguments;
		var later = function() {
			timeout = null;
			if (!immediate) func.apply(context, args);
		};
		var callNow = immediate && !timeout;
		clearTimeout(timeout);
		timeout = setTimeout(later, wait);
		if (callNow) func.apply(context, args);
	};
};

// Usage
var myEfficientFn = debounce(function() {
	// All the taxing stuff you do
}, 250);
window.addEventListener('resize', myEfficientFn);

The debounce function will not allow a callback to be used more than once per given time frame.  
This is especially important when assigning a callback function to frequently-firing events.

poll
----

As I mentioned with the debounce function, sometimes you don't get to plug into an event to 
signify a desired state -- if the event doesn't exist, you need to check for your desired 
state at intervals:

// The polling function
function poll(fn, timeout, interval) {
    var endTime = Number(new Date()) + (timeout || 2000);
    interval = interval || 100;

    var checkCondition = function(resolve, reject) {
        // If the condition is met, we're done! 
        var result = fn();
        if(result) {
            resolve(result);
        }
        // If the condition isn't met but the timeout hasn't elapsed, go again
        else if (Number(new Date()) < endTime) {
            setTimeout(checkCondition, interval, resolve, reject);
        }
        // Didn't match and too much time, reject!
        else {
            reject(new Error('timed out for ' + fn + ': ' + arguments));
        }
    };

    return new Promise(checkCondition);
}

// Usage:  ensure element is visible
poll(function() {
	return document.getElementById('lightbox').offsetWidth > 0;
}, 2000, 150).then(function() {
    // Polling done, now do something else!
}).catch(function() {
    // Polling timed out, handle the error!
});

Polling has long been useful on the web and will continue to be in the future!

once
----

There are times when you prefer a given functionality only happen once, similar to the way 
you'd use an onload event.  This code provides you said functionality:

function once(fn, context) { 
	var result;

	return function() { 
		if(fn) {
			result = fn.apply(context || this, arguments);
			fn = null;
		}

		return result;
	};
}

// Usage
var canOnlyFireOnce = once(function() {
	console.log('Fired!');
});

canOnlyFireOnce(); // "Fired!"
canOnlyFireOnce(); // nada

The once function ensures a given function can only be called once, thus prevent duplicate 
initialization!

getAbsoluteUrl
--------------

Getting an absolute URL from a variable string isn't as easy as you think. There's the URL 
constructor but it can act up if you don't provide the required arguments (which sometimes 
you can't).  Here's a suave trick for getting an absolute URL from and string input:

var getAbsoluteUrl = (function() {
	var a;

	return function(url) {
		if(!a) a = document.createElement('a');
		a.href = url;

		return a.href;
	};
})();

// Usage
getAbsoluteUrl('/something'); // https://davidwalsh.name/something

The "burn" element href handles and URL nonsense for you, providing a reliable absolute URL in return.

isNative
--------

Knowing if a given function is native or not can signal if you're willing to override it.  
This handy code can give you the answer:

;(function() {

  // Used to resolve the internal `[[Class]]` of values
  var toString = Object.prototype.toString;
  
  // Used to resolve the decompiled source of functions
  var fnToString = Function.prototype.toString;
  
  // Used to detect host constructors (Safari > 4; really typed array specific)
  var reHostCtor = /^\[object .+?Constructor\]$/;

  // Compile a regexp using a common native method as a template.
  // We chose `Object#toString` because there's a good chance it is not being mucked with.
  var reNative = RegExp('^' +
    // Coerce `Object#toString` to a string
    String(toString)
    // Escape any special regexp characters
    .replace(/[.*+?^${}()|[\]\/\\]/g, '\\$&')
    // Replace mentions of `toString` with `.*?` to keep the template generic.
    // Replace thing like `for ...` to support environments like Rhino which add extra info
    // such as method arity.
    .replace(/toString|(function).*?(?=\\\()| for .+?(?=\\\])/g, '$1.*?') + '$'
  );
  
  function isNative(value) {
    var type = typeof value;
    return type == 'function'
      // Use `Function#toString` to bypass the value's own `toString` method
      // and avoid being faked out.
      ? reNative.test(fnToString.call(value))
      // Fallback to a host object check because some environments will represent
      // things like typed arrays as DOM methods which may not conform to the
      // normal native pattern.
      : (value && type == 'object' && reHostCtor.test(toString.call(value))) || false;
  }
  
  // export however you want
  module.exports = isNative;
}());

// Usage
isNative(alert); // true
isNative(myCustomFunction); // false

The function isn't pretty but it gets the job done!

insertRule
----------

We all know that we can grab a NodeList from a selector (via document.querySelectorAll) and 
give each of them a style, but what's more efficient is setting that style to a selector 
(like you do in a stylesheet):

var sheet = (function() {
	// Create the <style> tag
	var style = document.createElement('style');

	// Add a media (and/or media query) here if you'd like!
	// style.setAttribute('media', 'screen')
	// style.setAttribute('media', 'only screen and (max-width : 1024px)')

	// WebKit hack :(
	style.appendChild(document.createTextNode(''));

	// Add the <style> element to the page
	document.head.appendChild(style);

	return style.sheet;
})();

// Usage
sheet.insertRule("header { float: left; opacity: 0.8; }", 1);

This is especially useful when working on a dynamic, AJAX-heavy site.  If you set the style 
to a selector, you don't need to account for styling each element that may match that selector 
(now or in the future).

matchesSelector
---------------

Oftentimes we validate input before moving forward; ensuring a truthy value, ensuring forms 
data is valid, etc.  But how often do we ensure an element qualifies for moving forward?  You 
can use a matchesSelector function to validate if an element is of a given selector match:

function matchesSelector(el, selector) {
	var p = Element.prototype;
	var f = p.matches || p.webkitMatchesSelector || p.mozMatchesSelector || p.msMatchesSelector || function(s) {
		return [].indexOf.call(document.querySelectorAll(s), this) !== -1;
	};
	return f.call(el, selector);
}

// Usage
matchesSelector(document.getElementById('myDiv'), 'div.someSelector[some-attribute=true]')

There you have it:  seven JavaScript functions that every developer should keep in their toolbox.  
Have a function I missed?  Please share it!
