$(document).ready(function () {

    $('html, body').animate({ scrollTop: $("#00Ni000000ESpAm").offset().top }, 300);
    $('html, body').animate({ scrollTop: $("#email").prev().offset().top }, 300);
    $('html, body').animate({ scrollTop: $("#last_name").prev().offset().top }, 300);

    $("#result").show(300).html('Please specify your First Name.');

    $('table tr:nth-child(odd)').addClass('odd');

    $('#ptlink').bind('mouseover', function () {
        $('img#fcet-image').attr("src", "/assets/drop-down-menus/programs-pm.jpg");
    });

    $('#ptlink').bind('mouseout', function () {
        $('img#fcet-image').attr("src", "/assets/drop-down-menus/programs-am.jpg");
    });

    $("#leadForm").submit(function(event) {

        event.preventDefault();

        $("#result").hide();

        var fields = $(this).serializeObject();
        
        var firstName = fields['first_name'];
        var lastName = fields['last_name'];

    });

    var filter = /^([A-Za-z0-9._-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,4})$/;

	$.ajax({
		type: "POST",
		url:  "/leads.jsp",
		data: $(this).serialize(),
		
		success: function(){
			dataLayer.push({'event':'GAevent', 'eventCategory':'Leadform', 
                            'eventAction':'Click', 'eventLabel':eventLabel});

			$("#result").show().html('Thanks! Your request has been sent!');

			setTimeout(function() { location.reload(); }, 3000 );
		},
		error: function(data) {
			$('#leadfrmsbmt').show(300);
			$("#result").show().html('Please try again.');
		}
	});
});

/* ----------------- Smooth Anchor Scroll ----------------- */

function filterPath(string) {
    return string.replace(/^\//, '')
                 .replace(/(index|default).[a-zA-Z]{3,4}$/, '')
                 .replace(/\/$/, '');
}


var locationPath = filterPath(location.pathname);
var scrollElem = scrollableElement('html', 'body');

$('a[href*=#]').each(function () {

    var thisPath = filterPath(this.pathname) || locationPath;

    if (locationPath == thisPath && (location.hostname == this.hostname || !this.hostname) && 
        this.hash.replace(/#/, '')) {

        var $target = $(this.hash),
            target = this.hash;

        if (target) {
            var targetOffset = $target.offset().top;
            $(this).click(function (event) {
                event.preventDefault();
                $(scrollElem).animate({
                    scrollTop: targetOffset
                }, 400, function () {
                    location.hash = target;
                });
            });
        }
    }
});

// use the first element that is "scrollable"
function scrollableElement(els) {
    for (var i = 0, argLength = arguments.length; i < argLength; i++) {

        var el = arguments[i],
            $scrollElement = $(el);

        if ($scrollElement.scrollTop() > 0) {
            return el;
        } else {
            $scrollElement.scrollTop(1);
            var isScrollable = $scrollElement.scrollTop() > 0;
            $scrollElement.scrollTop(0);

            if (isScrollable) {
                return el;
            }
        }
    }
    return [];
}
