function pr(d) { console.log(d); }

function createGUID() {
    function s4() {
        return Math.floor((1 + Math.random()) * 0x10000)
            .toString(16)
            .substring(1);
    }
    return s4() + s4() + '-' + s4() + '-' + s4() + '-' + s4() + '-' + s4() + s4() + s4();
}
pr(createGUID());
pr(createGUID());

function parseNumber(dirty) {
    var clean, number;

    if (typeof dirty == 'string') {
        // Try to clean it up.
        clean = dirty.replace(/[^\d.-]/g, '');

        // Parse it.
        number = parseFloat(clean);
    }
    else if (typeof dirty == 'number') {
        // Use it as is.
        number = dirty;
    }

    return number;
}

/**
 * Adds commas to number.
 * http://stackoverflow.com/questions/3883342/add-commas-to-a-number-in-jquery
 *
 * @param {Number} number
 * @param {Number} [decimals]
 * @param {String} [dec_point]
 * @param {String} [thousands_sep]
 * @returns {String}
 */
function formatNumber (number, decimals, dec_point, thousands_sep) {

    number = (number + '').replace(/[^0-9+\-Ee.]/g, '');

    var n = !isFinite(+number) ? 0 : +number,
        prec = !isFinite(+decimals) ? 0 : Math.abs(decimals),
        sep = (typeof thousands_sep === 'undefined') ? ',' : thousands_sep,
        dec = (typeof dec_point === 'undefined') ? '.' : dec_point,
        s = '',
        toFixedFix = function (n, prec) {
            var k = Math.pow(10, prec);
            return '' + (Math.round(n * k) / k).toFixed(prec);
        };

    // Fix for IE parseFloat(0.55).toFixed(0) = 0;
    s = (prec ? toFixedFix(n, prec) : '' + Math.round(n)).split('.');

    if (decimals == undefined) s = ('' + n).split('.');

    if (s[0].length > 3) {
        s[0] = s[0].replace(/\B(?=(?:\d{3})+(?!\d))/g, sep);
    }
    if ((s[1] || '').length < prec) {
        s[1] = s[1] || '';
        s[1] += new Array(prec - s[1].length + 1).join('0');
    }
    return s.join(dec);
}

//pr(formatNumber(12.56, 0));
//pr(formatNumber(12.56, 1));
//pr(formatNumber(12.56));

function mtRand(min, max) {
    var argc = arguments.length;
    if (argc === 0) {
        min = 0;
        max = 2147483647;
    } else if (argc === 1) {
        throw new Error('Warning: mt_rand() expects exactly 2 parameters, 1 given');
    } else {
        min = parseInt(min, 10);
        max = parseInt(max, 10);
    }
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

//pr(mtRand(10,100));

function stringToCamel(input) {

    var output = input;

    // Ensure we have a string.
    if (input) {
        // Parse the input.
        output = input
            .replace(/\/|_|-/gi, '_')
            .split('_')
            .map(function (text, index) {
                // Only capitalize after the first part.
                return index ? text.capitalize() : text;
            })
            .join('');
    }

    return output;
}
//pr(stringToCamel('this_is_a_string'));

function stringToAcronym(input) {
    var output = input;

    // Ensure have input.
    if (input) {
        // Parse the input.
        output = input
            .replace(/\/|_|-|'/gi, ' ')
            .match(/\b(\w)/gi)
            .join('')
            .toUpperCase();
    }

    return output;
}

//pr(stringToAcronym('this_is a_string'));

function stringPadLeft(input, char, length) {
    // Ensure input is string.
    input = '' + input;
    char = '' + char;

    var pad = char.repeat(length);

    return pad.substring(0, pad.length - input.length) + input;
}
//pr(stringPadLeft('short', '_', 10));

function range(start, end, step) {
  if (step == null) step = 1;
  var array = [];

  if (step > 0) {
    for (var i = start; i <= end; i += step)
      array.push(i);
  } else {
    for (var i = start; i >= end; i += step)
      array.push(i);
  }
  return array;
}

function sum(array) {
  var total = 0;
  for (var i = 0; i < array.length; i++)
    total += array[i];
  return total;
}

function repeat(string, times) {
  var result = "";
  for (var i = 0; i < times; i++)
    result += string;
  return result;
}

//pr(repeat('-', 40));
//pr(range(1, 10, 2));
//pr(range(10, 1, -2));
//pr(sum(range(1, 10, 2)));
