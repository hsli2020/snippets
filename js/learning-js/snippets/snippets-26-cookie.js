var checkCookieWritable = function(domain) {
    try {
        // Create cookie
        document.cookie = 'cookietest=1' + (domain ? '; domain=' + domain : '');
        var ret = document.cookie.indexOf('cookietest=') != -1;
        // Delete cookie
        document.cookie = 'cookietest=1; expires=Thu, 01-Jan-1970 00:00:01 GMT' + (domain ? '; domain=' + domain : '');
        return ret;
    } catch (e) {
        return false;
    }
};

var cookie = {
    write: function(name, value, days, domain, path) {
        var date = new Date();
        days = days || 730; // two years
        path = path || '/';
        date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
        var expires = '; expires=' + date.toGMTString();
        var cookieValue = name + '=' + value + expires + '; path=' + path;
        if (domain) {
            cookieValue += '; domain=' + domain;
        }
        document.cookie = cookieValue;
    },
    read: function(name) {
        var allCookie = '' + document.cookie;
        var index = allCookie.indexOf(name);
        if (name === undefined || name === '' || index === -1) return '';
        var ind1 = allCookie.indexOf(';', index);
        if (ind1 == -1) ind1 = allCookie.length;
        return unescape(allCookie.substring(index + name.length + 1, ind1));
    },
    remove: function(name) {
        if (this.read(name)) {
            this.write(name, '', -1, '/');
        }
    }
};

不是每个浏览器都支持window.localStorage，在使用之前必须确认是否可用。

var testCanLocalStorage = function() {
   var mod = 'modernizr';
   try {
       localStorage.setItem(mod, mod);
       localStorage.removeItem(mod);
       return true;
   } catch (e) {
       return false;
   }
};

检查SessionStorage可写性

var checkCanSessionStorage = function() {
  var mod = 'modernizr';
  try {
    sessionStorage.setItem(mod, mod);
    sessionStorage.removeItem(mod);
    return true;
  } catch (e) {
    return false;
  }
}

// handle IE8+
function ready (fn) {
    if (document.readyState != 'loading') {
        fn();
    } else if (window.addEventListener) {
        // window.addEventListener('load', fn);
        window.addEventListener('DOMContentLoaded', fn);
    } else {
        window.attachEvent('onreadystatechange', function() {
            if (document.readyState != 'loading')
                fn();
            });
    }
}

