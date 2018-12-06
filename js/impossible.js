/**
 * https://stackoverflow.com/questions/48270127/can-a-1-a-2-a-3-ever-evaluate-to-true
 */

// ##### Answer 1 #####
/*
const a = {
  i: 1,
  toString: function () {
    return a.i++;
  }
}

if (a == 1 && a == 2 && a == 3) {
  console.log('Hello World!');
}
*/

// The reason this works is due to the use of the loose equality operator.
// When using loose equality, if one of the operands is of a different type
// than the other, the engine will attempt to convert one to the other.
// In the case of an object on the left and a number on the right, it will
// attempt to convert the object to a number by first calling valueOf if it
// is callable, and failing that, it will call toString. I used toString in
// this case simply because it's what came to mind, valueOf would make more 
// sense. If I instead returned a string from toString, the engine would have
// then attempted to convert the string to a number giving us the same end 
// result, though with a slightly longer path.

// ##### Answer 2 #####
/*
var aﾠ = 1;
var a = 2;
var ﾠa = 3;
if(aﾠ==1 && a== 2 &&ﾠa==3) {
    console.log("Why hello there!")
}

// Note the weird spacing in the if statement (that I copied from your question).
// It is the half-width Hangul (that's Korean for those not familiar) which is
// an Unicode space character that is not interpreted by ECMA script as a space
// character - this means that it is a valid character for an identifier. Therefore
// there are three completely different variables, one with the Hangul after the a,
// one with it before and the last one with just a. Replacing the space with _ for
// readability, the same code would look like this:

var a_ = 1;
var a = 2;
var _a = 3;
if(a_==1 && a== 2 &&_a==3) {
    console.log("Why hello there!")
}
*/

// ##### Answer 3 #####
/*
var i = 0;

with({
  get a() {
    return ++i;
  }
}) {
  if (a == 1 && a == 2 && a == 3)
    console.log("wohoo");
}
*/

// ##### Answer 4 #####
/*
a = [1,2,3];
a.join = a.shift;
console.log(a == 1 && a == 2 && a == 3);
*/

// ##### Answer 5 #####
/*
let a = {[Symbol.toPrimitive]: ((i) => () => ++i) (0)};

console.log(a == 1 && a == 2 && a == 3);
*/

// ##### Answer 6 #####
/*
var val = 0;
Object.defineProperty(global, 'a', { // For nodejs use global instead of window
  get: function() {
    return ++val;
  }
});
if (a == 1 && a == 2 && a == 3) {
  console.log('yay');
}
*/

// ##### Answer 7 #####
/*
var a = {
  r: /\d/g, 
  valueOf: function(){
    return this.r.exec(123)[0]
  }
}

if (a == 1 && a == 2 && a == 3) {
    console.log("!")
}
*/

// ##### Answer 8 #####
/*
(() => {
    "use strict";
    Object.defineProperty(this, "a", {
        "get": () => {
            Object.defineProperty(this, "a", {
                "get": () => {
                    Object.defineProperty(this, "a", {
                        "get": () => {
                            return 3;
                        }
                    });
                    return 2;
                },
                configurable: true
            });
            return 1;
        },
        configurable: true
    });
    if (a == 1 && a == 2 && a == 3) {
        console.log("Yes, it’s possible.");
    }
})();
*/

// ##### Answer 9 #####
/*
function A() {
    var value = 0;
    this.valueOf = function () { return ++value; };
}

var a = new A;

if (a == 1 && a == 2 && a == 3) {
    console.log('bingo!');
}
*/

// ##### Answer 10 #####
/*
(function() {
    var i = 0;
    Object.defineProperty(global, "a", {
        get: function() {
            return ++i;
        }
    });
})();

if( a == 1 && a == 2 && a == 3 ) {
    console.log('Oh dear, what have we done?');
}
*/

// ##### Answer 10 #####
/*
a = 100000000000000000
if (a == a+1 && a == a+2 && a == a+3){
  console.log("Precision loss!");
}
*/

// ##### Answer 11 #####
/*
var  a = 1;
var ﾠ1 = a;
var ﾠ2 = a;
var ﾠ3 = a;
console.log( a ==ﾠ1 && a ==ﾠ2 && a ==ﾠ3 );
*/

// ##### Answer 12 #####
/*
var a = 1;
var a‌ = 2;
var a‍ = 3;
console.log(a == 1 && a‌ == 2 && a‍ == 3);
*/

/****
var a = 1;
var a\u200c = 2;
var a\u200d = 3;
console.log(a == 1 && a\u200c == 2 && a\u200d == 3);
****/

// ##### Answer 13 #####
/*
const a = {
  n: [3,2,1],
  toString: function () {
    return a.n.pop();
  }
}

if(a == 1 && a == 2 && a == 3) {
  console.log('Yes');
}
*/

// ##### Answer 14 #####
/*
const value = function* () {
  let i = 0;
  while(true) yield ++i;
}();

Object.defineProperty(this, 'a', {
  get() {
    return value.next().value;
  }
});

if (a === 1 && a === 2 && a === 3) {
  console.log('yo!');
}
*/

// ##### Answer 14 #####
/*
const a = { valueOf: () => this.n = (this.n || 0) % 3 + 1}
    
if(a == 1 && a == 2 && a == 3) {
  console.log('Hello World!');
}

if(a == 1 && a == 2 && a == 3) {
  console.log('Hello World!');
}
*/

// ##### Answer 15 #####
/*
var _a = 1

Object.defineProperty(this, "a", {
  "get": () => {
    return _a++;
  },
  configurable: true
});

console.log(a)
console.log(a)
console.log(a)
*/

// ##### Answer 15 #####
/*
const a = {value: 1};
a[Symbol.toPrimitive] = function() { return this.value++ };
console.log((a == 1 && a == 2 && a == 3));
*/

// ##### Answer 15 #####
/*
let foo = function* () {
    yield 1
    yield 2
    yield 3
}
let bar = foo()
Object.defineProperty(global, 'a', {
    get() {
        return bar.next().value
    }
})
console.log((a == 1 && a == 2 && a == 3));
*/
