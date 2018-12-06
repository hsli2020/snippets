/*!
 * jQuery custom-events plugin boilerplate
 * Author: DevPatch
 * Further changes: @addyosmani
 * Licensed under the MIT license
 */

// In this pattern, we use jQuery's custom events to add 
// pub/sub (publish/subscribe) capabilities to widgets. 
// Each widget would publish certain events and subscribe 
// to others. This approach effectively helps to decouple 
// the widgets and enables them to function independently.

;(function ( $, window, document, undefined ) {
    $.widget("ao.eventStatus", {
        options: {

        },
        
        _create : function() {
            var self = this;

            //self.element.addClass( "my-widget" );

            //subscribe to 'myEventStart'
            self.element.bind( "myEventStart", function( e ) {
                console.log("event start");
            });

            //subscribe to 'myEventEnd'
            self.element.bind( "myEventEnd", function( e ) {
                console.log("event end");
            });

            //unsubscribe to 'myEventStart'
            //self.element.unbind( "myEventStart", function(e){
                ///console.log("unsubscribed to this event"); 
            //});
        },

        destroy: function(){
            $.Widget.prototype.destroy.apply( this, arguments );
        },
    });
})( jQuery, window , document );

//Publishing event notifications
//usage: 
// $(".my-widget").trigger("myEventStart");
// $(".my-widget").trigger("myEventEnd");
