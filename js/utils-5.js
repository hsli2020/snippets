
function httpPost (url, parameters, callback) {
    if (parameters && typeof parameters !== "object") {
        console.error("Parameters is not an object");
        return;
    }
    if (callback && typeof callback !== "function") {
        console.error("Callback is not a function");
        return;
    }

    var query = [];
    if (!parameters) parameters = {};

    for (var p in parameters) {
        query.push(encodeURIComponent(p) + "=" + encodeURIComponent(parameters[p]));
    }
    query = query.join("&");

    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function(){
        if (this.readyState == 4) {
            callback && callback(this.status, this.responseText);
        }
    }

    xhttp.open("POST", url, true);
    xhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xhttp.send(query);
}

function getURLQuery(findKey) {
    var queryParts = window.location.search.substring(1).split("&");
    var queryData  = {};

    for (var i = 0; i < queryParts.length; i++) {
        var queryPair = queryParts[i].split("=");

        try {
            var key = decodeURIComponent(queryPair[0]);
        }
        catch (e) {
            continue;
        }

        try {
            var value = decodeURIComponent(queryPair[1] || "");
        }
        catch (e) {
            var value = null;
        }

        if (key) queryData[key] = value;
    }

    return findKey ? queryData[findKey] : queryData;
}

function updateURLQuery(queryData) {
    var currentQuery = getURLQuery();

    var keys = Object.keys(queryData);
    for (var i = 0; i < keys.length; i++) {
        var key = keys[i];

        if (queryData[key] === null) {
            delete currentQuery[key];
        }
        else {
            currentQuery[key] = queryData[key];
        }
    }

    setURLQuery(currentQuery);
}

function setURLQuery(queryData) {
    var queryArray = [];

    var keys = Object.keys(queryData);
    for (var i = 0; i < keys.length; i++) {
        var key = keys[i];
        var value = queryData[key];

        if (value) {
            queryArray.push(encodeURIComponent(key) + "=" + encodeURIComponent(value));
        }
        else {
            queryArray.push(encodeURIComponent(key));
        }
    }

    var queryString = queryArray.join("&");

    if (window.history.pushState) {
        var newURL = window.location.protocol + "//" + window.location.host + window.location.pathname;
        if (queryString) {
            newURL += "?" + queryString;
        }
        if (window.location.hash) {
            newURL += window.location.hash;
        }

        window.history.replaceState(null, "", newURL);
    }
}

function IE_version() {
    var ua = window.navigator.userAgent;

    var msie = ua.indexOf('MSIE ');
    if (msie > 0) {
        return parseInt(ua.substring(msie + 5, ua.indexOf('.', msie)), 10);
    }

    var trident = ua.indexOf('Trident/');
    if (trident > 0) {
        var rv = ua.indexOf('rv:');
        return parseInt(ua.substring(rv + 3, ua.indexOf('.', rv)), 10);
    }

    var edge = ua.indexOf('Edge/');
    if (edge > 0) {
        return parseInt(ua.substring(edge + 5, ua.indexOf('.', edge)), 10);
    }

    return false;
}

function getCookie(name) {
    var cookieName = name + "=";
    var cookiesArray = document.cookie.split(';');
    var cookiesLength = cookiesArray.length;

    for (var i = 0; i < cookiesLength; i++) {
        var cookie = cookiesArray[i];
        while (cookie.charAt(0) === ' ') {
            cookie = cookie.substring(1);
        }
        if (cookie.indexOf(cookieName) !== -1) {
            return decodeURIComponent(cookie.substring(cookieName.length, cookie.length));
        }
    }
    return "";
}

function setCookie(key, value, days) {
    var expires = new Date();
    expires.setTime(expires.getTime() + (days * 24 * 60 * 60 * 1000));
    document.cookie = key + '=' + value + ';expires=' + expires.toUTCString() + ';path=/';
}

function validateEmail(email) {
    var valid = false;

    // Filters everything without @
    // e.g 'userdomain'
    if (email.indexOf("@") !== -1) {

        // Filters multiple @ characters
        // e.g 'user@do@main'
        email = email.split("@");
        if (email.length == 2) {
            var user   = email[0];
            var domain = email[1];

            // Filters '@domain.com'
            if (user.length > 0) {
                domain    = domain.split(".");
                var size  = domain.length >= 2;
                var first = domain[0] && domain[0].length > 0;
                var last  = domain[size - 1] && domain[size - 1].length > 0;

                // Filters 'user@domain', 'user@.com', 'user@domain.'
                valid = (size && first && last);
            }
        }
    }

    return valid;
}

function templatePolyfill() {
    if (!("content" in document.createElement("template"))) {
        var templates = document.getElementsByTagName("template");
        var length = templates.length;

        for (var i = 0; i < length; ++i) {
            var template = templates[i];
            var content  = template.childNodes;
            var fragment = document.createDocumentFragment();

            while (content[0]) {
                fragment.appendChild(content[0]);
            }

            template.content = fragment;
        }
    }
}

function each(element, query, callback) {
    var elements = element.querySelectorAll(query);
    if (elements) {
        for (var i = 0; i < elements.length; i++) {
            callback(elements[i]);
        }
    }
}

var id = function(name) { 
    var element = document.getElementById(name);

    if (!element) {
        console.error("Couldn't find an element with id \"" + name + "\"");
    }

    return element;
}

var getScrollTop = function() {
    return parseInt(document.documentElement.scrollTop) || 0;
}

var setScrollTop = function(position) {
    document.documentElement.scrollTop = parseInt(position) || 0;
}

function registerServiceWorker() {
    if ("serviceWorker" in navigator) {
        navigator.serviceWorker.register("/js/service-worker.js");
        
        window.addEventListener("beforeinstallprompt", function(e) {
            e.preventDefault();
            page.addToHomeScreenPrompt = e;
        });
    }
}
