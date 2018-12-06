http://stackoverflow.com/questions/1877475/repeat-character-n-times

String.prototype.repeat= function(n){
    n= n || 1;
    return Array(n+1).join(this);
}

alert('Are we there yet?\nNo.\n'.repeat(10))


The most performance-wice way is 
https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/repeat

Short version is below.

String.prototype.repeat = function(count) {
    if (count < 1) return '';
    var result = '', pattern = this.valueOf();
    while (count > 1) {
        if (count & 1) result += pattern;
        count >>>= 1, pattern += pattern;
    }
    return result + pattern;
};

var a = "a";
console.debug(a.repeat(10));

/**  
 * Repeat a string `n`-times (recursive)
 * @param {String} s - The string you want to repeat.
 * @param {Number} n - The times to repeat the string.
 * @param {String} d - A delimiter between each string.
 */
var repeat = function (s, n, d) {
    return --n ? s + (d || "") + repeat(s, n, d) : "" + s;
};

var foo = "foo";
console.log(
    "%s\n%s\n%s\n%s",

    repeat(foo),        // "foo"
    repeat(foo, 2),     // "foofoo"
    repeat(foo, "2"),   // "foofoo"
    repeat(foo, 2, "-") // "foo-foo"
);


Another interesting way to quickly repeat n character is to use idea from quick exponentiation algorithm:

var repeatString = function(string, n) {
    var result = '', i;

    for (i = 1; i <= n; i *= 2) {
        if ((n & i) === i) {
            result += string;
        }
        string = string + string;
    }

    return result;
};

The following function will perform a lot faster than the option suggested in the accepted answer :

var repeat = function(str, count) {
    var array = [];
    for(var i = 0; i <= count;)
        array[i++] = str;
    return array.join('');
}

String.prototype.repeat = String.prototype.repeat || 
function(n) { return new Array(n + 1).join(this.toString()) }

function repeat(str, num) {
    var holder = [];
    for(var i=0; i<num; i++) {
        holder.push(str);
    }
    return holder.join('');
}


You can also try this simple function I found at the node.js core source code:

function repeatString(str, len) {
    return Array.apply(null, {
        length: len + 1
    }).join(str).slice(0, len)
}


<script type="text/javascript">
    (function($) {
        $('#wikiArticle').children('h2[id]').each(function() {
            var id = $(this).attr('id');
            var href = '/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/TrimLeft$edit#' + id;

            $('<a>').attr({
                href: href,
                'class': 'button section-edit only-icon'
            })
            .append($('<i aria-hidden="true" class="icon-pencil"></i>'))
            .append($('<span>'))
            .find('span')
            .text(gettext('Edit'))
            .end()
            .on('click', function(e) {
                e.preventDefault();

                mdn.analytics.trackEvent({
                    category: 'Section Edit',
                    action: id
                }, function() {
                    window.location = href;
                });
            })
            .appendTo(this);
        });
    })(jQuery);
</script>


(function (globals) {

  var django = globals.django || (globals.django = {});
  
  django.pluralidx = function (count) { return (count == 1) ? 0 : 1; };
  
  /* gettext identity library */

  django.gettext = function (msgid) { return msgid; };
  django.ngettext = function (singular, plural, count) { return (count == 1) ? singular : plural; };
  django.gettext_noop = function (msgid) { return msgid; };
  django.pgettext = function (context, msgid) { return msgid; };
  django.npgettext = function (context, singular, plural, count) { return (count == 1) ? singular : plural; };

  django.interpolate = function (fmt, obj, named) {
    if (named) {
      return fmt.replace(/%\(\w+\)s/g, function(match){return String(obj[match.slice(2,-2)])});
    } else {
      return fmt.replace(/%s/g, function(match){return String(obj.shift())});
    }
  };

  /* formatting library */

  django.formats = {
    "DATETIME_FORMAT": "N j, Y, P", 
    "DATETIME_INPUT_FORMATS": [
      "%Y-%m-%d %H:%M:%S", 
      "%Y-%m-%d %H:%M:%S.%f", 
      "%Y-%m-%d %H:%M", 
      "%Y-%m-%d", 
      "%m/%d/%Y %H:%M:%S", 
      "%m/%d/%Y %H:%M:%S.%f", 
      "%m/%d/%Y %H:%M", 
      "%m/%d/%Y", 
      "%m/%d/%y %H:%M:%S", 
      "%m/%d/%y %H:%M:%S.%f", 
      "%m/%d/%y %H:%M", 
      "%m/%d/%y"
    ], 
    "DATE_FORMAT": "N j, Y", 
    "DATE_INPUT_FORMATS": [
      "%Y-%m-%d", 
      "%m/%d/%Y", 
      "%m/%d/%y"
    ], 
    "DECIMAL_SEPARATOR": ".", 
    "FIRST_DAY_OF_WEEK": "0", 
    "MONTH_DAY_FORMAT": "F j", 
    "NUMBER_GROUPING": "3", 
    "SHORT_DATETIME_FORMAT": "m/d/Y P", 
    "SHORT_DATE_FORMAT": "m/d/Y", 
    "THOUSAND_SEPARATOR": ",", 
    "TIME_FORMAT": "P", 
    "TIME_INPUT_FORMATS": [
      "%H:%M:%S", 
      "%H:%M:%S.%f", 
      "%H:%M"
    ], 
    "YEAR_MONTH_FORMAT": "F Y"
  };

  django.get_format = function (format_type) {
    var value = django.formats[format_type];
    if (typeof(value) == 'undefined') {
      return format_type;
    } else {
      return value;
    }
  };

  /* add to global namespace */
  globals.pluralidx = django.pluralidx;
  globals.gettext = django.gettext;
  globals.ngettext = django.ngettext;
  globals.gettext_noop = django.gettext_noop;
  globals.pgettext = django.pgettext;
  globals.npgettext = django.npgettext;
  globals.interpolate = django.interpolate;
  globals.get_format = django.get_format;

}(this));


