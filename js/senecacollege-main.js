$(document).ready(function () {

    $('#mobileNavIcon').click(function () {
        $('nav').slideToggle();
        return false;
    });

    $('#searchIcon').click(function () {
        $('form#search').slideToggle();
        return false;
    });

    $('nav ul li').hoverIntent({
        over: megaMenuIn,
        timeout: 500,
        out: megaMenuOut
    });

    $('.megaMenu li').hoverIntent({
        over: megaMenuSubIn,
        out: megaMenuSubOut
    });

    function megaMenuIn() {
        $(this).children('.megaMenu').fadeIn();
        $(this).children('a').addClass('hovered');
    }

    function megaMenuOut() {
        $(this).children('.megaMenu').fadeOut();
        $(this).children('a').removeClass('hovered');
    }

    function megaMenuSubIn() {
        $(this).children('ul').fadeIn();
        $(this).children('a').addClass('hovered');
    }

    function megaMenuSubOut() {
        $(this).children('ul').fadeOut();
        $(this).children('a').removeClass('hovered');
    }

    $('#innerNav li.expandable > a').click(function() {
        $(this).siblings('ul').fadeToggle(); 
        return false;
    });

    $('#accordian h2').click(function () {
        $(this).next().slideToggle('fast');
        return false;
    }).next().hide();

    $('table tr:nth-child(odd)').addClass('odd');

    Shadowbox.init();

    $('.browser').treeview({
        collapsed: true,
        animated: 'medium',
        control: '#sidetreecontrol',
        prerendered: true,
        persist: 'location'
    });

    var seneSliderDefaults = {
        effect: 'fade',
        pauseTime: 8000,
        directionNav: false,
        directionNavHide: false,
        controlNav: false,
        keyboardNav: false,
        pauseOnHover: false,
        manualAdvance: false,
        captionOpacity: 0.8,
    };
    
    var seneSliderOptions = (typeof slider!="undefined") ? $.extend({}, seneSliderDefaults, slider) : seneSliderDefaults;
    $('#slider').nivoSlider(seneSliderOptions);

    $("#program-list-tbl").tablesorter();

    $('#ptlink').bind('mouseover', function () {
        $('img#fcet-image').attr("src", "/.assets/drop-down-menus/programs-pm.jpg");
    });
    $('#ptlink').bind('mouseout', function () {
        $('img#fcet-image').attr("src", "/.assets/drop-down-menus/programs-am.jpg");
    });

var count = 0;
$("#leadForm").submit(function(event) {

	event.preventDefault();
	$("#result").hide();

	var fields = $(this).serializeObject();
	
	var firstName = fields['first_name'];
	var lastName = fields['last_name'];
	var emailAddress = fields['email'];
	var caslOptIn = fields['00Ni000000ESpAm'];
	var eventLabel = fields['00Ni000000Bsatc'];
	
	if(typeof firstName == 'undefined' || firstName == '') {
		$("#result").show(300).html('Please specify your First Name.');
		$('html, body').animate({ scrollTop: $("#first_name").prev().offset().top }, 300);
		$("#first_name").focus();
		return false;
	}
	
	if(typeof lastName == 'undefined' || lastName == '') {
		$("#result").show(300).html('Please specify your Last Name.');
		$('html, body').animate({ scrollTop: $("#last_name").prev().offset().top }, 300);
		$("#last_name").focus();
		return false;
	}

	var filter = /^([A-Za-z0-9._-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,4})$/;
	
	if(typeof emailAddress == 'undefined' || !filter.test(emailAddress)) {
		$("#result").show(300).html('Please specify your Email Address.');
		$('html, body').animate({ scrollTop: $("#email").prev().offset().top }, 300);
		$("#email").focus();
		return false;
	}
	
	if(typeof caslOptIn == 'undefined' && count == 0) {
		$("#result").show(300).html('Reminder:<br>We need your consent. Please check the box above to agree to receive our electronic communications.');
		$('html, body').animate({ scrollTop: $("#00Ni000000ESpAm").offset().top }, 300);
		count++;
		return false;
	}
	
	$('#leadfrmsbmt').hide(300);
	
	$.ajax({
		type: "POST",
		url: "/leads.jsp",
		data: $(this).serialize(),
		
		success: function(){
			dataLayer.push({'event':'GAevent', 'eventCategory':'Leadform', 'eventAction':'Click', 'eventLabel':eventLabel});
			$("#result").show().html('Thanks! Your request has been sent!');
			setTimeout(function(){location.reload();},3000);
		},
		error: function(data){
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
        if (locationPath == thisPath && (location.hostname == this.hostname || !this.hostname) && this.hash.replace(/#/, '')) {
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
});