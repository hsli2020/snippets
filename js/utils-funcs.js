function isValidDate(value, userFormat) {
 
  // Set default format if format is not provided
  userFormat = userFormat || 'mm/dd/yyyy';
 
  // Find custom delimiter by excluding
  // month, day and year characters
  var delimiter = /[^mdy]/.exec(userFormat)[0];
 
  // Create an array with month, day and year
  // so we know the format order by index
  var theFormat = userFormat.split(delimiter);
 
  // Create array from user date
  var theDate = value.split(delimiter);
 
  function isDate(date, format) {
    var m, d, y, i = 0, len = format.length, f;
    for (i; i < len; i++) {
      f = format[i];
      if (/m/.test(f)) m = date[i];
      if (/d/.test(f)) d = date[i];
      if (/y/.test(f)) y = date[i];
    }
    return (
      m > 0 && m < 13 &&
      y && y.length === 4 &&
      d > 0 &&
      // Check if it's a valid day of the month
      d <= (new Date(y, m, 0)).getDate()
    );
  }
 
  return isDate(theDate, theFormat);
}

function highlight(text, words, tag) {
 
  // Default tag if no tag is provided
  tag = tag || 'span';
 
  var i, len = words.length, re;
  for (i = 0; i < len; i++) {
    // Global regex to highlight all matches
    re = new RegExp(words[i], 'g');
    if (re.test(text)) {
      text = text.replace(re, '<'+ tag +' class="highlight">$&</'+ tag +'>');
    }
  }
 
  return text;
}

$('p').html( highlight(
    $('p').html(), // the text
    ['foo', 'bar', 'baz', 'hello world'], // list of words or phrases to highlight
    'strong' // custom tag
));

function unhighlight(text, tag) {
  // Default tag if no tag is provided
  tag = tag || 'span';
  var re = new RegExp('(<'+ tag +'.+?>|<\/'+ tag +'>)', 'g');
  return text.replace(re, '');
}

