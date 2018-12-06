http://stackoverflow.com/questions/14128446/call-methods-using-jquery-plugin-design-pattern

;(function ( $, window, document, undefined ) {

    var pluginName = 'test';
    var defaults;

    function Plugin(element, options) {
        this.element = element;

        this.options = $.extend( {}, defaults, options) ;

        this._name = pluginName;

        this.init();
    }

    Plugin.prototype = {
        init: function() {
            this.hello();
        },
        hello : function() {
            document.write('hello');
        },
        goodbye : function() {
            document.write('goodbye');
        }
    }

    $.fn[pluginName] = function ( options ) {
        return this.each(function () {
            if (!$.data(this, 'plugin_' + pluginName)) {
                $.data(this, 'plugin_' + pluginName, 
                new Plugin( this, options ));
            }
        });
    }

})( jQuery, window, document );

$(document).ready(function() {
    $("#foo").test();
    $("#foo").test('goodbye');
});

Good

;
(function($, window, document, undefined) {

    var pluginName = 'test';
    var defaults;

    function Plugin(element, options) {
        this.element = element;

        this.options = $.extend({}, defaults, options);

        this._name = pluginName;

        this.init();
    }

    Plugin.prototype = {
        init: function(name) {
            this.hello();
        },
        hello: function(name) {
            console.log('hello');
        },
        goodbye: function(name) {
            console.log('goodbye');
        }
    }

    $.fn[pluginName] = function(options) {
        return this.each(function() {
            if (!$.data(this, 'plugin_' + pluginName)) {
                $.data(this, 'plugin_' + pluginName, new Plugin(this, options));
            }
            else if ($.isFunction(Plugin.prototype[options])) {
                $.data(this, 'plugin_' + pluginName)[options]();
            }
        });
    }

})(jQuery, window, document);

$(document).ready(function() {
    $("#foo").test();
    $("#foo").test('goodbye');
});
