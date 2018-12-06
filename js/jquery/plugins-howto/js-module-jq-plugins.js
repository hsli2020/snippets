// mymodule.js - a naive example module with 2 public functions,
// 1 private function and 1 private variable
var MYMODULE = (function () {

    // variables and functions private unless attached to API below
    // 'this' refers to global window

    // private array
    var array = [];

    // add a number into array
    function add(a) {
        log("add "+a);
        array.push(a);
    }

    // return copy of the array
    function get_array() {
        log("copy_array");
        return array.slice();
    }

    // a private debug function
    function log(msg) {
        console.debug(msg);
    }

    // define the public API
    var API = {};
    API.add = add;
    API.get_array = get_array;

    return API;
}());

The above module is used like this:

MYMODULE.add(1);
MYMODULE.add(3);
var arr = MYMODULE.get_array();

;(function($) {
    // wrap code within anonymous function: a private namespace

    // replace 'MyPlugin' and 'myplugin' below in your own plugin...

    // constructor function for the logical object bound to a
    // single DOM element
    function MyPlugin(elem, options) {

        // remember this object as self
        var self = this;
        // remember the DOM element that this object is bound to
        self.$elem = $(elem);

        // default options
        var defaults = {
            msg: "You clicked me!"
        };
        // mix in the passed-in options with the default options
        self.options = $.extend({}, defaults, options);

        // just some private data
        self.count = 1;

        init();

        // initialize this plugin
        function init() {

            // set click handler
            self.$elem.click(click_handler);
        }

        // private click handler
        function click_handler(event) {
            alert(self.options.msg + " " + self.count);
            self.count += 1;
        }

        // public method to change msg
        function change_msg(msg) {
            self.options.msg = msg;
        }

        // define the public API
        var API = {};
        API.change_msg = change_msg;
        return API;
    }

    // attach the plugin to jquery namespace
    $.fn.myplugin = function(options) {
        return this.each(function() {
            // prevent multiple instantiation
            if (!$(this).data('myplugin'))
                $(this).data('myplugin', new MyPlugin(this, options));
        });
    };
})(jQuery);

The plugin is used like this:

var options = {msg : "Click!"};
$('#elem').myplugin(options);

// get the API to the plugin object
var api = $('#elem').data('myplugin');
api.change_msg("Hello there!");

