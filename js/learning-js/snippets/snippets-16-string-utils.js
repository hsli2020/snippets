function pr(d) { console.log(d); }

String.prototype.repeat= function(n) {
    n = n || 1;
    return Array(n+1).join(this);
}
//pr("--- ".repeat(3));

if (!Array.isArray) {
  Array.isArray = function(arg) {
    return Object.prototype.toString.call(arg) === '[object Array]';
  };
}

if (!String.prototype.startsWith) {
    String.prototype.startsWith = function(searchString, position){
      position = position || 0;
      return this.substr(position, searchString.length) === searchString;
  };
}

if (!String.prototype.endsWith) {
  String.prototype.endsWith = function(searchString, position) {
      var subjectString = this.toString();
      if (typeof position !== 'number' || !isFinite(position) || 
          Math.floor(position) !== position || position > subjectString.length) {
          position = subjectString.length;
      }
      position -= searchString.length;
      var lastIndex = subjectString.indexOf(searchString, position);
      return lastIndex !== -1 && lastIndex === position;
  };
}

if (!String.prototype.includes) {
  String.prototype.includes = function(search, start) {
    'use strict';
    if (typeof start !== 'number') {
      start = 0;
    }
    
    if (start + search.length > this.length) {
      return false;
    } else {
      return this.indexOf(search, start) !== -1;
    }
  };
}

function strStartsWith(str, prefix) {
    return str.indexOf(prefix) === 0;
}
function strEndsWith(str, suffix) {
    return str.match(suffix+"$")==suffix;
}

String.prototype.startsWith = function(prefix) {
    return this.indexOf(prefix) === 0;
}

String.prototype.endsWith = function(suffix) {
    return this.match(suffix+"$") == suffix;
};

////////////////////////////////////////////////////////////

function randomString(len) {
  // Just an array of the characters we want in our random string
  var chrs = [
      'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
      'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',

      'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
      'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',

      '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'
  ];
 
  // Check that a length has been supplied and if not default to 32
  len = (isNaN(len)) ? 32 : len;
 
  // The following section shuffles the array just to further randomise the output
  var tmp, current, top = chrs.length; 
  if (top) {
    while (--top) { 
      current = Math.floor(Math.random() * (top + 1)); 
      tmp = chrs[current]; 
      chrs[current] = chrs[top]; 
      chrs[top] = tmp; 
    }
  }
 
  // Just a holder for our random string
  var randomStr = '';
 
  // Loop through the required number of characters grabbing one at random from the array each time
  for(i=0; i<len; i++) {
    randomStr = randomStr + chrs[Math.floor(Math.random()*chrs.length)];
  }
 
  // Return our random string
  return randomStr;
}

////////////////////////////////////////////////////////////
