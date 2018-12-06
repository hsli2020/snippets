var helper = {

  // Returns first match in regex; or ""
  regex: function(regex, string) {
    regex.lastIndex = 0;
    var match = regex.exec(string);

    if (match) return match[1];
    return "";
  },


  // Returns all matches in a regex;
  regexAll: function(regex, string) {
    regex.lastIndex = 0;
    var matches = [];
    while (match = regex.exec(string)) {
        matches.push(match[1]);
    }
    return matches;
  },


  normaliseUrl: function(url, domain) {
    return this.toHttps(domain + this.toRelative(url));
  },


  // Used to filter array to unique values only
  // filtered = array.filter( onlyUnique )
  onlyUnique: function(value, index, self) {
    return self.indexOf(value) === index;
  },


  // Returns Chrome Extension base URL
  extUrl: function() {
    return chrome.extension.getURL("");
  },


  // Returns decoded URI string
  decodeURI: function(string) {
    return decodeURIComponent(string.replace(/\+/g, " "));
  },


  // Converts string to js number format
  toNumber: function(x) {
    x = x.replace(/\$/g, '');
    x = x.replace(/,/g, '');
    if ($.isNumeric(x)) return Number(x);
    return 0;
  },


  //replace , to . for de,fr,es,it
  replaceDot: function(x) {
    x = x.replace(/,/g, '.');
    return x;
  },


  replacePoint: function(x) {
    if (!x) return x;
    x = x.replace(/\./g, ',');
    return x;
  },


  isNum: function(x) {
    var reg = /^(\d+.,?)+$/;
    return reg.test(x);
  },


  trimStr: function(str){
    return str.replace(/(^\s*)|(\s*$)/g,"");
  },


  clearHTMLComments: function(html){
    return html.replace(/<!--[\s\S]*?-->/g, '');
  },

  isCaptchaPage: function(html){
    if (html.match(/>Robot Check<\/title>/)) {
      return true;
    }
    return false;
  },

  // Adds comma thousands separator
  numberWithCommas: function(x) {
    if (!$.isNumeric(x)) return x;
      return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
  },


  // Returns count of active ajax calls
  ajaxCount: function() {
    return $.xhrPool.length;
  },


  // Returns date/time from timestamp
  timestampToDate: function(timestamp) {
    var d = new Date(timestamp);
    return d.toLocaleString();
  },


  // Array to Object
  arrayToObject: function(arr) {
    var rv = {};
    for (var i = 0; i < arr.length; ++i)
      if (arr[i] !== undefined) rv[i] = arr[i];
      return rv;
    }
};


$.xhrPool = [];
$.xhrPool.abortAll = function() {
  $(this).each(function(i, jqXHR) {
    jqXHR.abort();
    $.xhrPool.splice(i, 1);
  });
};
