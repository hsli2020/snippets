var rels = {
    Viola: {
        Orsino: {
            Olivia: {
                Cesario: null
            }
        }
    }
};

// Outputs: undefined
// console.log(rels.Viola.Harry);

// TypeError: Cannot read property 'Sally' of undefined
// console.log(rels.Viola.Harry.Sally);

function deepGet (obj, properties) {
    // If we have reached an undefined/null property
    // then stop executing and return undefined.
    if (obj === undefined || obj === null) {
        return;
    }

    // If the path array has no more elements, we've reached
    // the intended property and return its value.
    if (properties.length === 0) {
        return obj;
    }

    // Prepare our found property and path array for recursion
    var foundSoFar = obj[properties[0]];
    var remainingProperties = properties.slice(1);

    return deepGet(foundSoFar, remainingProperties);
}

// Outputs: { Cesario: null }
console.log(deepGet(rels, ["Viola", "Orsino", "Olivia"]));

// Outputs: undefined
console.log(deepGet(rels, ["Viola", "Harry"]));

// Outputs: undefined
console.log(deepGet(rels, ["Viola", "Harry", "Sally"]));

/*
// defaultValue
function deepGet (obj, props, defaultValue) {
    // If we have reached an undefined/null property
    // then stop executing and return the default value.
    // If no default was provided it will be undefined.
    if (obj === undefined || obj === null) {
        return defaultValue;
    }

    // If the path array has no more elements, we've reached
    // the intended property and return its value
    if (props.length === 0) {
        return obj;
    }

    // Prepare our found property and path array for recursion
    var foundSoFar = obj[props[0]];
    var remainingProps = props.slice(1);

    return deepGet(foundSoFar, remainingProps, defaultValue);
}

sallyRel = deepGet(rels, ["Viola", "Harry", "Sally"], {});

// Will output a graph based on the empty object
graph(sallyRel);
*/

// https://medium.com/javascript-inside/safely-accessing-deeply-nested-values-in-javascript-99bf72a0855a

const props = {
  user: {
    posts: [
      { title: 'Foo', comments: [ 'Good one!', 'Interesting...' ] },
      { title: 'Bar', comments: [ 'Ok' ] },
      { title: 'Baz', comments: [] },
    ]
  }
}

/*
// access deeply nested values...
props.user &&
props.user.posts &&
props.user.posts[0] &&
props.user.posts[0].comments

props.user &&
props.user.posts &&
props.user.posts[0] &&
props.user.posts[0].comments &&
props.user.posts[0].comments[0]
*/

const get = (p, o) =>
  p.reduce((xs, x) => (xs && xs[x]) ? xs[x] : null, o)

// let's pass in our props object...

console.log(get(['user', 'posts', 0, 'comments'], props))
// [ 'Good one!', 'Interesting...' ]

console.log(get(['user', 'post', 0, 'comments'], props))
// null


// https://codereview.stackexchange.com/questions/72253/safe-navigating-function-for-nested-object-properties

var foo = {
    bar: {
        baz: 1
    }
};

function getDeepProp(obj, properties) {

    if (typeof properties === 'string')
        properties = properties.split('.');

    if (typeof obj === 'undefined')
        return;

    if (!properties.length)
        return obj;

    return getDeepProp(obj[properties[0]], properties.slice(1));
}

console.log(getDeepProp(foo, 'bar.baz')); // 1
console.log(getDeepProp(foo, 'baz.baz')); // undefined

// The function was created to avoid exceptions like:
console.log(foo.baz.baz); // Uncaught TypeError: Cannot read property 'baz' of undefined    

