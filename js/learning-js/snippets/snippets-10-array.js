function each(collection, iterator) {
  if (Array.isArray(collection)) {
    for (var i = 0; i < collection.length; i++) {
      iterator(collection[i], i, collection);
    }
  } else {
    for (var key in collection) {
      iterator(collection[key], key, collection);
    }
  }
 }

function filter(collection, test) {
  var filtered = [];
  each(collection, function(item) {
    if (test(item)) {
      filtered.push(item);
    }
  });
  return filtered;
}

function map(collection, iterator) {
  var mapped = [];
  each(collection, function(value, key, collection) {
    mapped.push(iterator(value));
  });
  return mapped;
}

function reduce(collection, iterator, accumulator) {
    var startingValueMissing = accumulator === undefined;

    each(collection, function(item) {
      if(startingValueMissing) {
        accumulator = item;
        startingValueMissing = false;
      } else {
        accumulator = iterator(accumulator, item);
      }
    });

    return accumulator;
}
