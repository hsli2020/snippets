(function (root, factory) {
    "use strict";

    if (typeof define === "function" && define.amd) {
        // AMD. Register as an anonymous module.
        define(["jquery"], factory);
    } else if (typeof exports === "object") {
        // Node. Does not work with strict CommonJS, but only 
        // CommonJS-like environments that support module.exports, like Node.
        module.exports = factory(require("jquery"));
    } else {
        // Browser globals (root is window)
        root.bootbox = factory(root.jQuery);
    }

}(this, function init($, undefined) {
    "use strict";

    var defaults = {
        locale: "en",   // default language
        /* .. */
    };

    var exports = {};

    function each(collection, iterator) {
        var index = 0;
        $.each(collection, function(key, value) {
            iterator(key, value, index++);
        });
    }

    exports.alert = function() { /* ... */ };
    exports.confirm = function() { /* ... */ };
    exports.prompt = function() { /* ... */ };
    exports.dialog = function(options) { /* ... */ };

    exports.setDefaults = function() {
        var values = {};

        if (arguments.length === 2) {
            // allow passing of single key/value...
            values[arguments[0]] = arguments[1];
        } else {
            // ... and as an object too
            values = arguments[0];
        }

        $.extend(defaults, values);
    };

    exports.init = function(_$) {
        return init(_$ || $);
    };

    return exports;

}));
