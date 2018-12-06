'use strict';

/* exported smoothScrollLink, setNGBreadcrumb */

// small per page jquery actions
$(function () {

    var doc = document.documentElement;
    var isIE11 = !!navigator.userAgent.match(/Trident\/7\./);
    var isIE = !!navigator.userAgent.match(/MSIE/);
    doc.setAttribute('data-useragent', navigator.userAgent);

    if (isIE11) {
        $('body').addClass('ie11');
    }
    if (isIE) {
        $('body').addClass('ie');
    }

    $(window).on('resizeEnd', resizeWindow);
    $(window).on('scroll', siteNav.mobile.scrollWindow);

    // Navigation
    $('.nav.navbar-nav li, .breadcrumb li.dropdown').on('mouseover', siteNav.openSub);
    $('.nav.navbar-nav li, .breadcrumb li.dropdown').on('mouseout', siteNav.openSub);
    $('.breadcrumb li.dropdown').on('mouseout', siteNav.openSub);

    $('.nav.navbar-nav li a, .breadcrumb li.dropdown a, .nav.navbar-nav li input[type="text"]').on('focus',
        function (e) {
            $(e.target).parents('li').trigger('mouseover');
        }
    );
    $('.nav.navbar-nav li a, .breadcrumb li.dropdown a, .nav.navbar-nav li input[type="text"]').on('blur',
        function (e) {
            $(e.target).parents('li').trigger('mouseout');
        }
    );

    // removing tabbing
    removeSlideMenuTabs();

    $('.dropdown-toggle').click(function (e) {
        e.preventDefault();
        e.stopPropagation();

        return false;
    });

    $('.navbar-toggle').on('click', siteNav.mobile.toggleNav);
    $('#site-menu [data-target]').on('click', siteNav.mobile.toSubMenu);
    $('.breadcrumb.visible-xs').on('click', '[data-target]', siteNav.mobile.toSubMenu);
    $('.faq_category_nav.selected').on('click', faqToggle);

    $('.sub-back.topLevel').on('click', siteNav.mobile.closeSub);
    $('.search-item').on('click', siteNav.toggleSearchMenu);

    $('#menu-close-overlay').on('click', siteNav.mobile.toggleNav);
    $('#menu-close-overlay').on('swipeleft', siteNav.mobile.swipeLeftHandler);
    $('#site-menu').on('swipeleft', siteNav.mobile.swipeLeftHandler);
    $('#menu-close-overlay').on('swiperight', siteNav.mobile.swipeRightHandler);
    $('.service-disruption').on('click', 'button', siteNav.serviceDisruptions);
    $(".service-disruption a").click(siteNav.serviceDisruptions);
    $('.searchMenu, #mobileSearchMenu').on('submit', 'form', siteNav.siteSearch);

    if ($('.search-results').length > 0) {
        siteNav.toggleSearchMenu();
    }

    $('.equal-height').responsiveEqualHeightGrid();
    siteNav.affixHeader();

    $('a[href*=#]:not([href=#])').click(smoothScrollHashAnchors);

    // accessibility
    $(".textResizer a").click(resizeText);
    checkResizedText();

    // decorators
    $("table:not(.ng-scope table)").wrap('<div class="table-responsive"></div>');

    $(".load-more-results a").click(loadMoreResults);

    if (window.sessionStorage !== undefined) {
        var serviceDisruptionEL = $('.service-disruption');

        if (window.sessionStorage.getItem('display_serviceDisruptions') !== null) {
            if (window.JSON.parse(window.sessionStorage.getItem('display_serviceDisruptions')) === false) {
                serviceDisruptionEL.find('button').trigger('click');
            } else {
                siteNav.getServiceDisruptions();
            }
        } else {
            window.sessionStorage.setItem('display_serviceDisruptions', true);
            siteNav.getServiceDisruptions();
        }
    }
});

$(function () {

    var resizeEnd;

    $(window).on('resize', function () {

        clearTimeout(resizeEnd);

        resizeEnd = setTimeout(function () {
            $(window).trigger('resizeEnd');
        }, 250);
    });

});

function closeModals() {
    var element = document.getElementById('eligibility-warning-close');
    element && element.click();
}

function removeSlideMenuTabs() {

    $('.mobile_footer a, #slide_menu_wrapper a').each(function (obj) {
        $(this).attr("tabindex", "-1");
    });
}

function checkResizedText() {
    // check local storage to set size
    if (sessionStorage.Textsize) {
        $("body").addClass(sessionStorage.Textsize);
        $('.equal-height').responsiveEqualHeightGrid();
    }
}

function resizeText(e) {

    e.preventDefault();
    var body = $("body");
    body.removeClass('textSize1').removeClass('textSize2').removeClass('textSize3');
    var newSize;

    // click event to resize text
    switch ($(e.target)[0].classList[0]) {
        case('size1'):
            newSize = 'textSize1';
            break;
        case('size2'):
            newSize = 'textSize2';
            break;
        case('size3'):
            newSize = 'textSize3';
            break;
    }

    body.addClass(newSize);

    // set local storage
    sessionStorage.setItem("Textsize", newSize);
    $('.equal-height').responsiveEqualHeightGrid();
}

function smoothScrollHashAnchors(event) {

    var anchor = event.currentTarget;

    var height = $("div.navigation-wrapper").outerHeight();
    height += 10; //Buffer of 10px to place above the linked target

    if (location.pathname.replace(/^\//, '') === anchor.pathname.replace(/^\//, '') && location.hostname === anchor.hostname && !/^#\//.test(anchor.hash)) {
        var target = $(anchor.hash);

        if (target && target.selector === '#top') {
            $('html,body').animate({
                scrollTop: 0
            }, 300);
            return false;
        }

        target = target.length ? target : $('[name=' + anchor.hash.slice(1) + ']');
        if (target.length) {
            $('html,body').animate({
                scrollTop: target.offset().top - height
            }, 300);
            return false;
        }
    }

}

function isElementVisible(id) {
    var el =  document.getElementById(id);
    var rect     = el.getBoundingClientRect(),
        vWidth   = window.innerWidth || doc.documentElement.clientWidth,
        vHeight  = window.innerHeight || doc.documentElement.clientHeight,
        efp      = function (x, y) { return document.elementFromPoint(x, y) };

    // Return false if it's not in the viewport
    if (rect.right < 0 || rect.bottom < 0
        || rect.left > vWidth || rect.top > vHeight)
        return false;

    // Return true if any of its four corners are visible
    return (
        el.contains(efp(rect.left,  rect.top))
        ||  el.contains(efp(rect.right, rect.top))
        ||  el.contains(efp(rect.right, rect.bottom))
        ||  el.contains(efp(rect.left,  rect.bottom))
    );
}

function focusOnSubmit(window, element) {
    var submitButton = $(element).children().find('button[type=submit]');
    window.smoothScroll(submitButton.attr('id'));
    submitButton.focus();
}

function smoothScroll(targetId) {
    var target = $('#' + targetId);

    if(!isElementVisible(targetId)) {

        $('html,body').animate({
            scrollTop: target.offset().top - $(window).height() // scroll just within the window
        });
    }
}

function smoothScrollLink(targetId, smooth, topOffset) {

    if (smooth === undefined) {
        smooth = true;
    }
    if (topOffset === undefined) {
        topOffset = 0;
    }

    window.scrollTarget = targetId;

    if (topOffset !== 0) {
        var checkExist = setInterval(function () {   // deferred animation

            if ($(window.scrollTarget).children().length > 0) {
                $('html,body').animate({
                    scrollTop: $(window.scrollTarget).offset().top - topOffset
                }, {
                    duration: 1000,
                    complete: function () {
                        // close the gap
                        if (Math.abs($('body').scrollTop() - ($(window.scrollTarget).offset().top - topOffset)) > 5) {
                            $('html,body').animate({
                                scrollTop: $(window.scrollTarget).offset().top - topOffset
                            }, {duration: 300});
                        }
                    }
                });
                clearInterval(checkExist);
            }
        }, 100);

    } else {
        $('html,body').scrollTop(0);
    }
    return false;

}

var dtUtils = {
    /**
     * Checks value/string for a subset of numbers
     * @param {string} user input string.
     * @return {boolean} The result whether a credit card number is present or not.
     */
    creditCardSubString: function (value) {
        var regex = /\d+/g, returnedResult = false, m;

        while ((m = regex.exec(value)) !== null) {
            if (dtUtils.luhnChk(m[0]) && m[0].length >= 13 && m[0].length <= 16) {
                returnedResult = true;
                break;
            }
        }

        return returnedResult;
    },

    /**
     * Mod 10 algorithm
     * @param {string} user input string.
     * @return {boolean} The result whether a credit card number is present or not.
     */
    luhnChk: function (luhn) {
        // Mod 10 algorithm
        // https://gist.github.com/ShirtlessKirk/2134376

        var len = luhn.length,
            mul = 0,
            prodArr = [[0, 1, 2, 3, 4, 5, 6, 7, 8, 9], [0, 2, 4, 6, 8, 1, 3, 5, 7, 9]],
            sum = 0;

        while (len--) {
            sum += prodArr[mul][parseInt(luhn.charAt(len), 10)];
            mul ^= 1;
        }

        return sum % 10 === 0 && sum > 0;
    }
};

var siteNav = {
    openSub: function (e) {

        var subEl = e.currentTarget;

        $(subEl).parent().children().removeClass('open');

        if (e.type === 'mouseover') {
            if ($(subEl).is('.dropdown')) {
                $(subEl).addClass('open');
            }
        }
    },

    toggleSearchMenu: function () {
        $("#mobileSearchMenu").toggleClass("open");
    },

    getOffset: function () {
        var returnObj = {};
        returnObj.offset = {};

        returnObj.offset.top = 0;
        returnObj.offset.bottom = $('footer').outerHeight(true);

        return returnObj;
    },

    affixHeader: function () {
        var $header = $('.navigation-wrapper > header');
        $header.affix({
            offset: 0
        });
        var $header2 = $('.breadcrumb.container-fluid.hidden-xs');
        $header2.affix({
            offset: 163
        });
        $('.affixWithHeader').affix({
            offset: $header.outerHeight()
        });
    },

    affixReset: function () {
        var $header = $('.navigation-wrapper > header');
        $(window).off('.affix');
        $('.scrollHide').removeClass('scrollHide');
        $header.removeData('bs.affix').removeClass('affix affix-top affix-bottom');
        $header.affix({
            offset: 0
        });
        var $header2 = $('.breadcrumb.container-fluid.hidden-xs');
        $header2.removeData('bs.affix').removeClass('affix affix-top affix-bottom');
        $header2.affix({
            offset: 163
        });

    },
    getServiceDisruptions: function () {
        // poll the rest service for disruption, and insert into div is ready
/*        if (window.sessionStorage.getItem('display_serviceDisruptions') !== "false") {
            // cache copy
            if (!siteNav.displayServiceDisruptionsText(JSON.parse(window.sessionStorage.getItem('text_serviceDisruptions')))) {
                $.ajax({
                    url: "/app/api/centreDisruptions"
                }).done(function (data) {
                    $('.disruptionList').html("");
                    if (data.disruptions !== undefined && data.disruptions.length > 0) {
                        window.sessionStorage.setItem('text_serviceDisruptions', JSON.stringify(data.disruptions));
                        siteNav.displayServiceDisruptionsText(data.disruptions);
                    }
                });
            }
        }*/
    },
    displayServiceDisruptionsText: function (msg) {
        if (msg) {
            $('.service-disruption').show(100);
            for (var ind = 0; ind < msg.length; ind++) {
                if (ind < (msg.length - 1) || (msg.length === 1)) {
                    $('.disruptionList').append(msg[ind].name);
                    if (ind < msg.length - 1) {
                        $('.disruptionList').append(", ");
                    }
                } else {
                    if (msg.length > 1) {
                        $('.service-disruption .disruptionListAnd').removeClass('hide');
                        $('.disruptionList2').html(msg[ind].name);
                        $('.disruptionList2').removeClass('hide');
                    }
                }
            }
            if (msg.length > 1) {
                $('.disruptionListEndPlural').removeClass('hide');
            } else {
                $('.disruptionListEnd').removeClass('hide');
            }
        }
        return msg;
    },
    serviceDisruptions: function (evt) {
        window.sessionStorage.setItem('display_serviceDisruptions', false);
        $(evt.currentTarget).closest('.service-disruption').remove();
        return evt;
    },

    siteSearch: function (evt) {
        var searchInputEl = $(evt.currentTarget).find('input.form-control'),
            searchValue = searchInputEl.val();

        if (dtUtils.creditCardSubString(searchValue)) {
            searchInputEl.val('');
            return false;
        } else {
            return true;
        }
    }
};


siteNav.mobile = {

    lastScrollPos: $(window).scrollTop(),

    getSubNavItems: function () {
        return $('.navbar-nav .slide-menu');
    },

    toggleNav: function (e) {
        e.preventDefault();

        if (siteNav.mobile.isNavOpen()) {
            siteNav.mobile.getSubNavItems().removeClass('show-sub');
            siteNav.mobile.closeNav();
        } else {
            siteNav.mobile.openNav();
        }
    },

    closeSub: function (e) {
        e.preventDefault();
        $(".slide-menu").removeClass("show-sub");
    },

    toSubMenu: function (e) {
        e.preventDefault();

        if (!siteNav.mobile.isNavOpen()) {
            siteNav.mobile.toggleNav(e);
        }

        //data-target is an a, in an li, in a div we want to open.
        var target = $(e.currentTarget);
        var linkTarget = target.attr("data-target");

        var elem = $("[data-link='" + linkTarget + "']");

        $(".slide-menu").removeClass("show-sub");
        elem.parents(".slide-menu").addClass("show-sub");
    },

    // Callback function references the event target and adds the 'swipe' class to it
    swipeLeftHandler: function (e) {

        if (siteNav.mobile.isNavOpen()) {
            siteNav.mobile.closeNav();
        }
    },

    swipeRightHandler: function (e) {

        if (!siteNav.mobile.isNavOpen()) {
            siteNav.mobile.openNav();
        }
    },

    isNavOpen: function () {
        return $('#site-wrapper').hasClass('show-nav');
    },

    openNav: function () {
        $('#site-wrapper').addClass('show-nav');
        $('body').addClass('vscrollbar_hidden');
        $('#slide_menu_wrapper').removeClass('vscrollbar_hidden').addClass('vscrollbar_shown');
        $('#site-menu').removeClass('side-menu-hidden');
    },

    closeNav: function () {
        $('#site-wrapper').removeClass('show-nav');
        $('body').removeClass('vscrollbar_hidden');
        $('#slide_menu_wrapper').removeClass('vscrollbar_shown').addClass('vscrollbar_hidden');
        $('#site-menu').addClass('side-menu-hidden');
    },

    scrollWindow: function (evt) {

        var $el = $(evt.currentTarget);

        if ($el.outerWidth() < 768) {
            var currentScrollPos = $el.scrollTop(),
                $header = $('.navigation-wrapper > header'),
                $headerChildContainer = $header.find('.header.container');

            if ((currentScrollPos > siteNav.mobile.lastScrollPos) && (currentScrollPos > $headerChildContainer.outerHeight(true))) {
                if ($header.hasClass('affix')) {
                    $header.addClass('scrollHide');
                    $('.affixWithHeader').addClass('scrollHide');
                }
            } else {
                $header.removeClass('scrollHide');
                $('.affixWithHeader').removeClass('scrollHide');
            }
            siteNav.mobile.lastScrollPos = currentScrollPos;
        }

    }
};

function faqToggle(e) {
    if ($('.faq_questions').is(":hidden")) { /* mobile only */
        e.preventDefault();
        $('.faq_nav_row').toggleClass('open');
    }
}

function resizeWindow(e) {
    // side nav tasks
    if ($(window).width() >= 768) {
        siteNav.mobile.closeNav();
    }
    siteNav.mobile.closeSub(e);

    //Reset affix offset
    siteNav.affixReset();
}

function loadMoreResults(e) {
    e.preventDefault();
    var index = $("body").data("resultIndex");
    if (index === undefined) {
        index = 12;
    }
    $.getJSON("/app/search/query/loadMore/?query=" + $("input[name=userQuery]").val() + "&start=" + index +
        "&language=" + $("input[name=language]").val(), function (data) {
        var items = [];
        var col = 0;
        var results = 0;
        $.each(data.searchResults, function (key, val) {
            results++;
            $("body").data("resultIndex", ++index);
            col++;
            if (col === 1) {
                items.push("<div class='row'>");
            }
            items.push("<div class='col-md-6 col-sm-6 search-result'><h3><a class='result-link' href='" +
                val.url + "'>" + val.title + "</a></h3><p><em>" + val.url + "</em></p><p>" + val.description + "</p></div>");
            if (col === 2 || key === data.searchResults.length) {
                items.push("</div>");
                col = 0;
            }
        });
        if (results < 12) {
            $('.load-more-results').hide();
        }
        $("<div/>", {
            "class": "load-more",
            html: items.join("")
        }).appendTo(".search-result-col");
    });
}

function setNGBreadcrumb(val) {
    if (val !== undefined && val !== "") {
        $(".angularBreadcrumb").removeClass('hide');
        $(".angularBreadcrumbItem").html("<span>" + val + "</span>");
    } else {
        $(".angularBreadcrumb").addClass('hide');
    }
}

function disableBookingContainer() {
    $('#booking-container').addClass("disabled");
}

function enableBookingContainer() {
    $('#booking-container').removeClass("disabled");
}
'use strict';

(function($) {

  /**
   * Set all elements within the collection to have the same height.
   */
  $.fn.equalHeight = function(){
    var heights = [];
    $.each(this, function(i, element){
      var $element = $(element),
          element_height;
      // Should we include the elements padding in it's height?
      var includePadding = ($element.css('box-sizing') === 'border-box') || ($element.css('-moz-box-sizing') === 'border-box');
      if (includePadding) {
        element_height = $element.innerHeight();
      } else {
        element_height = $element.height();
      }
      heights.push(element_height);
    });
    this.height(Math.max.apply(window, heights));
    return this;
  };

  /**
   * Create a grid of equal height elements.
   */
  $.fn.equalHeightGrid = function(columns){
    var $tiles = this;
    $tiles.css('height', 'auto');
    for (var i = 0; i < $tiles.length; i++) {
      if (i % columns === 0) {
        var row = $($tiles[i]);
        for(var n = 1;n < columns;n++){
          row = row.add($tiles[i + n]);
        }
        row.equalHeight();
      }
    }
    return this;
  };

  /**
   * Detect how many columns there are in a given layout.
   */
  $.fn.detectGridColumns = function() {
    var offset = 0, cols = 0;
    this.each(function(i, elem) {
      var elem_offset = $(elem).offset().top;
      if (offset === 0 || elem_offset === offset) {
        cols++;
        offset = elem_offset;
      } else {
        return false;
      }
    });
    return cols;
  };

  /**
   * Ensure equal heights now, on ready, load and resize.
   */
  $.fn.responsiveEqualHeightGrid = function() {
    var _this = this;
    function syncHeights() {
      var cols = _this.detectGridColumns();
      _this.equalHeightGrid(cols);  
    }
    $(window).bind('resize load', syncHeights);
    syncHeights();
    return this;
  };

})(jQuery);

"use strict";

angular.module('DriveTest.services', ['pascalprecht.translate']);

'use strict';

//Block Booking App
angular.module('dtBlockBooking', ['DriveTest.AutoSubmit', 'DriveTest.SchoolBusinessNumber', 'DriveTest.filters', 'DriveTest.services', 'DriveTest.directives', 'pascalprecht.translate', 'google-maps', 'ngResource', 'ui.router', 'ui.bootstrap', 'ngRoute', 'DriveTest.findACentre', 'DriveTest.CompareValidation','DriveTest.DateValidation', 'ngAnimate', 'ngMessages', 'ngSanitize', 'DriveTest.ViewTitle', 'ngCookies', 'DriveTest.ButtonSubmit', 'ui.bootstrap.datetimepicker'])
    .config(['$urlRouterProvider', '$stateProvider', '$provide', '$httpProvider', '$translateProvider',
        function($urlRouterProvider, $stateProvider, $provide, $httpProvider, $translateProvider) {

            $httpProvider.interceptors.push(['$q','$location', '$log',function($q,$location, $log){
                return {

                    responseError: function(rejection){
                        $log.debug('interceptor rejection ---> ', rejection);
                        $log.debug('$location.path() --------> ', $location.path());

                        //place verify-drive/validate-driver-email 401 fail case here
                        if($location.path() !== '/login' && $location.path() !== '/signup') {
                            if (rejection.status === 401) {
                                document.cookie = 'ssid' +'=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
                                $location.path('quit-booking');
                            }
                        }

                        return $q.reject(rejection);
                    }
                };
            }]);

            //if (document.URL.match(/nobackend/)) {
            //    $provide.decorator('$httpBackend', angular.mock.e2e.$httpBackendDecorator);
            //}

            $translateProvider.useLoader('translateLoader', {
                // if you have some custom properties, go for it!
                prefix: '/application/error_codes',
                suffix: '.html'
            });
            $translateProvider.preferredLanguage('en_CA');
            $translateProvider.registerAvailableLanguageKeys(['en_CA', 'fr_CA'], {
                'en': 'en_CA',
                'en_US': 'en_CA',
                'fr': 'fr_CA',
                'fr_US': 'fr_CA'
            });
            //$translateProvider.useLocalStorage();

            $urlRouterProvider.otherwise('/signup');

            $stateProvider
                .state('school', {
                    abstract: true,
                    resolve: {
                        // A string value resolves to a service
                        locale: 'locale',

                        // A function value resolves to the return
                        // value of the function
                        localeStr: function(locale) {
                            return locale.getLocaleStr();
                        }
                    },
                    templateUrl: '/ng/dt-block'
                })
                .state('school.registration', {
                    url: '/signup',
                    allowAnonymous: true,
                    loginState: 'school.dashboard',
                    views: {
                        'blockbooking-general': {
                            templateUrl: '/ng/dt-block-signup',
                            controller: 'blockSignupController'
                        }
                    }
                })
                .state('school.login', {
                    url: '/login',
                    allowAnonymous: true,
                    loginState: 'school.dashboard',
                    views: {
                        'blockbooking-general': {
                            templateUrl: '/ng/dt-block-login',
                            controller: 'blockLoginController'
                        }
                    }
                })
                .state('school.dashboard', {
                    url: '/dashboard',
                    views: {
                        'blockbooking-general': {
                            templateUrl: '/ng/dt-block-dashboard',
                            controller: 'blockDashboardController'
                            //templateUrl: '/ng/dt-block-summary',
                            //controller: 'blockSummaryController'
                        }
                    }
                })
                .state('school.addbooking', {
                    url: '/addbooking',
                    views: {
                        'blockbooking-general': {
                            templateUrl: '/ng/dt-block-addstudent',
                            controller: 'blockAddStudentController'
                        }
                    }
                })
                .state('school.location', {
                    url: '/location',
                    views: {
                        'booking-location': {
                            templateUrl: '/ng/dt-block-find',
                            controller: 'blockLocationController'
                        }
                    }
                })
                .state('school.timeslot', {
                    url: '/timeslot',
                    views: {
                        'blockbooking-general': {
                            templateUrl: '/ng/dt-block-timeslot',
                            controller: 'blockTimeslotController'
                        }
                    }
                })
                .state('school.summary', {
                    url: '/summary',
                    views: {
                        'blockbooking-general': {
                            templateUrl: '/ng/dt-block-summary',
                            controller: 'blockSummaryController'
                        }
                    }
                })
                .state('school.paymentSubmit', {
                    url: '/school-payment-submit',
                    allowAnonymous: true,
                    views: {
                        'blockbooking-general': {
                            templateUrl: '/ng/dt-booking-submit',
                            controller: 'blockPaymentSubmitController'
                        }
                    }
                })
                .state('school.finalize', {
                    url: '/finalizing-booking',
                    views: {
                        'blockbooking-general': {
                            templateUrl: '/ng/dt-booking-finalize',
                            controller: 'finalizeBlockBookingController'
                        }
                    }
                })
                .state('school.verifyEmail', {
                    url: '/verify',
                    allowAnonymous: true,
                    views: {
                        'blockbooking-general': {
                            templateUrl: '/ng/dt-reg-verifyemail',
                            controller: 'blockVerifyEmailController'
                        }
                    }
                })
                .state('school.error', {
                    url: '/error',
                    allowAnonymous: true,
                    views: {
                        'blockbooking-general': {
                            templateUrl: '/ng/dt-booking-error'
                        }
                    }
                })
                .state('school.forgotPassword', {
                    url: '/forgot',
                    allowAnonymous: true,
                    views: {
                        'blockbooking-general': {
                            templateUrl: '/ng/dt-block-forgot-login',
                            controller: 'blockForgotPasswordController'
                        }
                    }
                })
                .state('school.resetPassword', {
                    url: '/reset?token',
                    allowAnonymous: true,
                    views: {
                        'blockbooking-general': {
                            templateUrl: '/ng/dt-block-reset-login',
                            controller: 'blockResetPasswordController'
                        }
                    }
                })
                .state('school.quit', {
                    url: '/quit-booking',
                    allowAnonymous: true,
                    views: {
                        'blockbooking-general': {
                            template: '',
                            controller: 'blockQuitController'
                        }
                    }
                });
        }
    ])
    .run([
        '$rootScope',
        '$state',
        '$location',
        '$stateParams',
        '$log',
        'stepState',
        'blockSignupService',
        '$cookies',
        'appointmentService',
        function(
            $rootScope,
            $state,
            $location,
            $stateParams,
            $log,
            stepState,
            blockSignupService,
            $cookies,
            appointmentService
        ) {
            $rootScope.$on("$stateChangeStart", function(event, toState, toParams, fromState, fromParams){
                $('.page-loading').show();

                appointmentService.getStatusToken(
                    function (result) {
                        $log.debug("SUCCESS Login Status: " + result.statusToken);
                        $log.debug("Login status defined: " + angular.isDefined(result.statusToken));

                        if((!angular.isDefined(toState.allowAnonymous) || !toState.allowAnonymous) && !(blockSignupService.isSet() || angular.isDefined(result.statusToken))) {
                            event.preventDefault();
                            stepState.changeState('school.quit');
                        } else if ((!angular.isDefined(toState.allowAnonymous) || !toState.allowAnonymous)
                            && (!blockSignupService.isSet() && angular.isDefined(result.statusToken) && toState.name !== 'school.dashboard' && toState.name !== 'school.addbooking')) {
                            event.preventDefault();
                            stepState.changeState('school.dashboard');
                        } else if(angular.isDefined(toState.allowAnonymous) && toState.allowAnonymous && angular.isDefined(toState.loginState)
                            && (blockSignupService.isSet() || angular.isDefined(result.statusToken))) {
                            event.preventDefault();
                            stepState.changeState(toState.loginState);
                        }
                    },
                    function (result) {
                        $log.debug("ERROR Login Status");
                    }
                );
            });

            $rootScope.$on('$viewContentLoaded', function(event, toState, toParams, fromState, fromParams){
                $('.page-loading').hide();
            });
        }
    ]);
/**
 * Created by at_user on 19/01/2015.
 */

'use strict';

angular.module('DriveTest.directives', [])
    .directive('ngNotACreditCard',['$window', function($window){
        return {
            restrict: 'AC',
            require: 'ngModel',
            link: function (scope, elem, attrs, ctrl) {

                ctrl.$validators.isCCNumber = function(value){
                    var isValid = true;

                    if(value !== undefined){
                        isValid = ($window.dtUtils.creditCardSubString(value) !== true);
                    }

                    ctrl.$setValidity('iscreditcard', isValid);

                    return value;
                };

            }
        };
    }]);
'use strict';

angular.module('DriveTest.CompareValidation',[])
    .directive('dtMatch', ['$parse', function ($parse) {
        return {
            restrict: 'A',
            require: 'ngModel',
            link: function (scope, elem, attrs, ctrl) {
                if (!ctrl) {
                    return;
                }
                var parsedFunction = $parse(attrs.dtMatch);
                var validator = function (value) {
                    var compareValue = parsedFunction(scope);
                    var result = value === compareValue;
                    ctrl.$setValidity('match', result);
                    return value;
                };

                ctrl.$parsers.unshift(validator);
                ctrl.$formatters.push(validator);
                scope.$watch(attrs.dtMatch, function () {
                    validator(ctrl.$viewValue);
                });
            }
        };
    }]);
'use strict';

angular.module('DriveTest.DateValidation', [])
    .directive('dtValidDate', ['$parse', function ($parse) {
        return {
            priority: 112,
            restrict: 'A',
            require: 'ngModel',
            link: function (scope, elem, attrs, ctrl) {
                if (!ctrl) {
                    return;
                }
                var dateFormat = attrs.dtValidDate.toUpperCase();

                var yearFormat = "YYYY";
                var monthFormat = "MM";
                var dayFormat = "DD";

                var yearStartIndex = dateFormat.indexOf(yearFormat);
                var yearEndIndex = yearStartIndex + yearFormat.length;
                var monthStartIndex = dateFormat.indexOf(monthFormat);
                var monthEndIndex = monthStartIndex + monthFormat.length;
                var dayStartIndex = dateFormat.indexOf(dayFormat);
                var dayEndIndex = dayStartIndex + dayFormat.length;

                var isWithinRange = function(currentIndex, startIndex, endIndex){
                    return (startIndex <= currentIndex && endIndex > currentIndex);
                };

                var validator = function (value) {
                    var isDateValid = false;
                    if(typeof(value) !== 'undefined' && value.length === dateFormat.length){
                        var year = parseInt(value.substring(yearStartIndex, yearEndIndex));
                        var month = parseInt(value.substring(monthStartIndex, monthEndIndex), 10) - 1;
                        var day = parseInt(value.substring(dayStartIndex, dayEndIndex));
                        if(!isNaN(year) && !isNaN(month) && !isNaN(day)) {
                            var expiryDate = new Date(year, month, day);
                            isDateValid = (year === expiryDate.getFullYear() && month === expiryDate.getMonth() && day === expiryDate.getDate());
                        }

                        if(isDateValid) {
                            for (var currentIndex = 0; currentIndex < dateFormat.length; currentIndex++) {
                                if (isWithinRange(currentIndex, yearStartIndex, yearEndIndex) || isWithinRange(currentIndex, monthStartIndex, monthEndIndex) || isWithinRange(currentIndex, dayStartIndex, dayEndIndex)) {
                                    continue;
                                }
                                var expectedSeparator = dateFormat.charAt(currentIndex);
                                var actualSeparator = value.length >= currentIndex ? value.charAt(currentIndex) : "";
                                if (expectedSeparator !== actualSeparator) {
                                    isDateValid = false;
                                }
                            }
                        }
                    }
                    ctrl.$setValidity('validdate', isDateValid);
                    return value;
                };

                ctrl.$parsers.push(validator);
                ctrl.$formatters.push(validator);

            }
        };
    }]);

'use strict';

angular.module('DriveTest.ViewTitle', ['ui.router'])
    .directive('dtViewTitle', ['$document','$state', function ($document, $state) {
        return {
            priority: 112,
            restrict: 'A',
            link: function (scope, elem, attrs, ctrl) {
                $document[0].title = attrs.dtViewTitle;
            }
        };
    }]);

'use strict';

angular.module('DriveTest.ButtonSubmit', [])
    .directive('btnSubmit', ['$log',function($log){
	return {

		scope: {
            blocked: '=',
            loading: '=',
            inactive: '='

        },
		restrict: 'AE',

        templateUrl: '/ng/dtd-booking-btn-submit',
		replace: true,
		transclude: false,

		link: function($scope, iElm, iAttrs, LocationController) {
			$log.debug('--- btnSubmit LiNK ---' + $scope.blocked+":"+$scope.loading+":"+$scope.disabled);

		}
	};
}]);

/**
 * Created by at_user on 19/01/2015.
 */

'use strict';

angular.module('DriveTest.Email', [])
    .directive('dtEmail', function(){
        return {
            restrict: 'AC',
            require: 'ngModel',
            link: function (scope, elem, attrs, ctrl) {

                ctrl.$validators.validateEmail = function(value) {

                    if (value !== undefined) {
                        var atIndex = value.indexOf('@');
                        var containsAt = atIndex != -1;
                        ctrl.$setValidity('doesNotContainAtSign', containsAt);

                        var domain = atIndex == -1 ? "" : value.substring(atIndex + 1);
                        var domainIsNotEmpty = domain.length > 0;
                        ctrl.$setValidity('domainIsEmpty', domainIsNotEmpty);

                        var containsDotInDomain = domainIsNotEmpty && domain.indexOf('.') != -1;
                        ctrl.$setValidity('domainDoesNotContainPeriod', containsDotInDomain);

                        var validDomainSymbolsRegex = /^([a-zA-Z0-9-.])*$/;
                        var containsValidSymbolsInDomain = domainIsNotEmpty && domain.match(validDomainSymbolsRegex) != null;
                        ctrl.$setValidity('domainContainsInvalidSymbol', containsValidSymbolsInDomain);

                        var startsWithAlphaNumeric = /^([a-zA-Z0-9]).*$/;
                        var domainStartsWithAlphaNumeric = domainIsNotEmpty && domain.match(startsWithAlphaNumeric) != null;
                        ctrl.$setValidity('domainDoesNotStartWithAlphaNumeric', domainStartsWithAlphaNumeric);

                        var domainRegex = /^([a-zA-Z0-9]+(-[a-zA-Z0-9]+)*\.)+[a-zA-Z]{1,}$/;
                        var dotsDashFollowedByAlphaNumeric = domainIsNotEmpty && domain.match(domainRegex) != null;
                        ctrl.$setValidity('domainPeriodOrDashIsNotFollowedByAlphaNumeric', dotsDashFollowedByAlphaNumeric);
                    }

                    return value;
                };


            }
        };
    });
/**
 * Created by XuTong Zhu on 25/04/2018.
 */

'use strict';

angular.module('DriveTest.Focus', [])
    .directive('dtFocusOn', function(){
        return {
            restrict: 'AC',
            require: 'ngModel',
            link: function (scope, elem, attrs) {
                scope.$on(attrs.dtFocusOn, function(e) {
                    elem[0].focus();
                });
            }
        };
    });
/**
 * Created by at_user on 19/01/2015.
 */

'use strict';

angular.module('DriveTest.LicenceExpiry', [])
    .directive('dtLicenceExpiry',['$window', function($window){
        return {
            restrict: 'AC',
            require: 'ngModel',
            link: function (scope, elem, attrs, ctrl) {

                elem.on('change blur', function() {
                    var singleSegmentFormat = /[0-9]{8}/;
                    var value = elem[0].value;
                    if (value.search(singleSegmentFormat) != -1) {
                        value = value.substring(0,4) + "/" + value.substring(4, 6) + "/" + value.substring(6, 8);
                        elem[0].value = value;
                        ctrl.$setViewValue(value);
                    }
                });

                ctrl.$validators.validateLicenceExpiry = function(value) {

                    if (value === undefined || !value) {
                        return value;
                    }

                    var segments = value.split('/');

                    var firstSegment = segments[0];
                    var firstSegmentFormat = /^[0-9]*$/;
                    ctrl.$setValidity('yearTooShort', firstSegment !== undefined && firstSegment.length >= 4);
                    ctrl.$setValidity('yearTooLong', firstSegment !== undefined && firstSegment.length <= 4);
                    ctrl.$setValidity('yearInvalidChar', firstSegment !== undefined && firstSegment.search(firstSegmentFormat) != -1);

                    var secondSegment = segments[1];
                    var secondSegmentFormat = /^[0-9]*$/;
                    ctrl.$setValidity('monthTooShort', secondSegment !== undefined && secondSegment.length >= 2);
                    ctrl.$setValidity('monthTooLong', secondSegment !== undefined && secondSegment.length <= 2);
                    ctrl.$setValidity('monthInvalidChar', secondSegment !== undefined && secondSegment.search(secondSegmentFormat) != -1);
                    var monthNumber = parseInt(secondSegment);
                    ctrl.$setValidity('monthInvalidRange', monthNumber >= 1 && monthNumber <= 12 );

                    var thirdSegment = segments[2];
                    ctrl.$setValidity('dayTooShort', thirdSegment !== undefined && thirdSegment.length >= 2);
                    ctrl.$setValidity('dayTooLong', thirdSegment !== undefined && thirdSegment.length <= 2);
                    ctrl.$setValidity('dayInvalidChar', thirdSegment !== undefined && thirdSegment.search(secondSegmentFormat) != -1);
                    var dayNumber = parseInt(thirdSegment);
                    ctrl.$setValidity('dayInvalidRange', dayNumber >= 1 && dayNumber <= 31 );

                    return value;
                };


            }
        };
    }]);
/**
 * Created by at_user on 19/01/2015.
 */

'use strict';

angular.module('DriveTest.LicenceNumber', [])
    .directive('dtLicenceNumber', function(){
        return {
            restrict: 'AC',
            require: 'ngModel',
            link: function (scope, elem, attrs, ctrl) {

                elem.on('change blur', function() {
                    var singleSegmentFormat = /[a-zA-Z][0-9]{14}/;
                    var value = elem[0].value;
                    if (value.search(singleSegmentFormat) != -1) {
                        value = value.substring(0,5) + "-" + value.substring(5, 10) + "-" + value.substring(10, 15);
                        elem[0].value = value;
                        ctrl.$setViewValue(value);
                    }
                });

                ctrl.$validators.validateLicenceNumber = function(value) {

                    if (value === undefined || !value) {
                        return value;
                    }

                    var startsWithAlpha = value && value.charAt(0).search(/[a-zA-Z]/) == 0;
                    ctrl.$setValidity('doesNotStartWithAlpha', startsWithAlpha);

                    var segments = value.split('-');
                    var validNumberOfSegments = segments.length == 3;
                    ctrl.$setValidity('invalidNumberOfSegments', validNumberOfSegments);

                    var firstSegment = segments[0];
                    var firstSegmentFormat = /^[a-zA-Z][0-9]{4}$/;
                    var validFirstSegment = firstSegment !== undefined && firstSegment.search(firstSegmentFormat) != -1;
                    ctrl.$setValidity('invalidFirstSegment', validFirstSegment);

                    var secondSegment = segments[1];
                    var secondSegmentFormat = /^[0-9]{5}$/;
                    var validSecondSegment = secondSegment !== undefined && secondSegment.search(secondSegmentFormat) != -1;
                    ctrl.$setValidity('invalidSecondSegment', validSecondSegment);

                    var thirdSegment = segments[2];
                    var validThirdSegment = thirdSegment !== undefined && thirdSegment.search(secondSegmentFormat) != -1;
                    ctrl.$setValidity('invalidThirdSegment', validThirdSegment);

                    return value;
                };


            }
        };
    });
/**
 * Created by Hisham on 2014-09-12.
 */
'use strict';

angular.module('DriveTest.AutoSubmit', [])
    .directive('autoSubmitForm', ['$timeout', function($timeout) {
        return {
            replace: true,
            scope: {},
            template: '<form name="dynamicForm" action="{{formData.url}}" method="POST">' +
                            '<input type="hidden" name="shopId" value="{{formData.shopId}}" />' +
                            '<input type="hidden" name="encodedMessage" value="{{formData.encodedMessage}}" />' +
                            '<input type="hidden" name="signature" value="{{formData.signature}}" />' +
                      '</form>',
            link: function($scope, element, $attrs) {
                $scope.$on($attrs.event, function(event, data) {

                    $scope.formData = data;

                    $timeout(function() {
                        $("form[name='dynamicForm']").submit();
                    }, 250);
                });
            }
        };
    }]);

'use strict';

var dtFiltersModule = angular.module('DriveTest.filters', ['DriveTest.services']);
dtFiltersModule.filter('dtCurrency', ['$filter', 'locale', function(filter, locale) {
    var currencyFilter = filter('currency');
    return function(amount) {
        if (locale.isFrench() && amount !== undefined) {
            return currencyFilter(amount, '')
                .replace('.', ';')
                .replace(/\,/g, ',')
                .replace(';', ',') + ' $';
        } else {
            return currencyFilter(amount, '$');
        }
    };
}]);
'use strict';
//A cache service for DriveTest.
//Inject object is dtCacheServices and dtCacheServicesKeys
//Add key to dtCacheServicesKeys to cache new data with the new key
//Note that Angular cache only exists in one web session. Page reloading will vanish the cache.
var dtCacheServicesModule = angular.module('DriveTest.services');
dtCacheServicesModule.value('dtCacheServicesKeys',
    {
        TILE_SERVER_URL: 'tileServerURL'
    }
);

dtCacheServicesModule.service('dtCacheServices',[
    '$cacheFactory',
    function ($cacheFactory) {
        return $cacheFactory('dtCacheServicesFactory');
    }
]);



'use strict';

/* exported error */

var dtServicesModule = angular.module('DriveTest.services');

dtServicesModule
    .service('errorService', ['$log', '$window', 'stepState', '$translate', 'HTTP_STATUS_CODES',
        function($log, $window, stepState, $translate, HTTP_STATUS_CODES) {

            var _this = this;

            this.clean = function($scope) {
                $log.debug("cleaning errors");

                var errMsg = { hasError: false, blocking: false, msg: undefined, status: undefined };
                return errMsg;
            };

            this.setBlocked = function($scope) {
                $log.debug("setting blocked");

                var errMsg = $scope.errMsg;
                errMsg.blocking = true;
                return errMsg;
            };

            this.isBlocked = function($scope) {
                $log.debug("checking blocked");

                var errMsg = $scope.errMsg;
                return errMsg.blocking;
            };

            this.getMessage = function(langKey){
                var genericErr = $translate.instant('booking.error'),
                    key = (angular.isDefined(langKey) && langKey !== null) ? $translate.instant(langKey) : genericErr,
                    errorMsg = (key !== langKey) ? key : genericErr;

                return errorMsg;
            };

            this.errorHandler = function($scope, callback, errorKey) {

                if (!angular.isDefined($scope.requestStatus)){
                    $scope.requestStatus = {};
                }

                var returnFunc = function(response) {

                    $log.debug('errorHandler response = ', response);

                    if(!response.data && response.headersLength === 0) {
                        showNetworkError($scope);
                    } else {

                        var completeHandling = callback ? callback(response) : false;

                        if(!completeHandling) {
                            _this.showServerError($scope, response, errorKey);
                        }
                    }
                };

                return returnFunc;
            };

            this.showServerError = function($scope, errorData, errorKey) {

                errorKey = (errorKey === undefined) ? '' : errorKey;
                var status = errorData.status,
                    errObj = {hasError : true, msg : '', status : 0};

                if(angular.isDefined(status) && HTTP_STATUS_CODES.hasOwnProperty(status)) {

                    errObj.msg = _this.getMessage(errorKey);
                    errObj.status = status;

                    if((status === 404 || /^[5]\d{2}/.test(status))){
                        stepState.goToGeneralErrorState();
                        return;
                    }

                } else {
                    stepState.goToGeneralErrorState();
                    return;
                }

                $scope.errMsg = errObj;
            };

            function showNetworkError($scope) {
                stepState.goToGeneralErrorState();
            }
        }
    ]);

/**
 * Created by at_user on 25/01/2015.
 */

angular.module('DriveTest.services')
    .constant('HTTP_STATUS_CODES', {
        '100': 'Continue',
        '101': 'Switching Protocols',
        '102': 'Processing',
        '200': 'OK',
        '201': 'Created',
        '202': 'Accepted',
        '203': 'Non-Authoritative Information',
        '204': 'No Content',
        '205': 'Reset Content',
        '206': 'Partial Content',
        '207': 'Multi-Status',
        '208': 'Already Reported',
        '226': 'IM Use',
        '300': 'Multiple Choices',
        '301': 'Moved Permanently',
        '302': 'Moved Temporarily',
        '303': 'See Other',
        '304': 'Not Modified',
        '305': 'Use Proxy',
        '307': 'Temporary Redirect',
        '400': 'Bad Request',
        '401': 'Unauthorized',
        '402': 'Payment Required',
        '403': 'Forbidden',
        '404': 'Not Found',
        '405': 'Method Not Allowed',
        '406': 'Not Acceptable',
        '407': 'Proxy Authentication Required',
        '408': 'Request Time-out',
        '409': 'Conflict',
        '410': 'Gone',
        '411': 'Length Required',
        '412': 'Precondition Failed',
        '413': 'Request Entity Too Large',
        '414': 'Request-URI Too Large',
        '415': 'Unsupported Media Type',
        '416': 'Requested range not satisfiable',
        '417': 'Expectation Failed',
        '419': 'Insufficient Space On Resource',
        '420': 'Method Failure',
        '421': 'Destination Locked',
        '422': 'Unprocessable Entity',
        '423': 'Locked',
        '424': 'Failed Dependency',
        '426': 'Upgrade Required',
        '500': 'Internal Server Error',
        '501': 'Not Implemented',
        '502': 'Bad Gateway',
        '503': 'Service Unavailable',
        '504': 'Gateway Time-out',
        '505': 'HTTP Version not supported',
        '506': 'Variant Also Negotiates',
        '507': 'Insufficient Storage',
        '508': 'Loop Detected',
        '510': 'Not Extended'
    });
'use strict';

var dtServicesModule = angular.module('DriveTest.services');

dtServicesModule.service('locale', ['$translate',
    function($translate) {

        var _cachedLocale;
        this.isFrench = function() {
            if (_cachedLocale === undefined) {
                _cachedLocale = this.getLocaleStr();
            }
            return _cachedLocale === 'fr_CA';
        };
        this.isEnglish = function() {
            if (_cachedLocale === undefined) {
                _cachedLocale = this.getLocaleStr();
            }
            return _cachedLocale === 'en_CA';
        };
        this.setLocale = function(inLocale) {
            _cachedLocale = inLocale;
            $translate.use(_cachedLocale);
        };
        this.getLocaleStr = function() {
            var loc = _cachedLocale;
            if (loc === undefined) {
                loc = angular.element('body').hasClass('en_CA') ? 'en_CA' : '';
                loc = angular.element('body').hasClass('fr_CA') ? 'fr_CA' : loc;
                _cachedLocale = loc;
                $translate.use(_cachedLocale);
            }
            return loc;
        };
    }
]);
'use strict';

var dtcLocationsModule = angular.module('DriveTest.services');

dtcLocationsModule
    .factory('dtcLocations', ['$window', '$resource', '$cacheFactory',
        function($window, $resource, $cacheFactory) {
            //Consume DTC Location service

            return $resource($window.dtServiceEndpoint + '/location', {}, {
                'query': {
                    method: 'GET',
                    isArray: false,
                    cache: $cacheFactory
                }
            });
        }
    ]).factory('tileServerURLCache', function($cacheFactory) {

            return $cacheFactory('tileServerURLCache');

    }).factory('dotcmsTileServer', ['$window', '$resource', '$cacheFactory', function($window, $resource, $cacheFactory) {
            return $resource($window.dtServiceEndpoint + '/tileserver', {}, {
                'query': {
                    method: 'GET',
                    isArray: false,
                    cache: $cacheFactory
                }
            });
        }
    ])
    .factory('dotcmsLocations', ['$window', '$resource', '$cacheFactory', function($window, $resource) {
        return $resource($window.dtServiceEndpoint + '/drivetestcentres', {}, {
            'query': {
                method: 'GET',
                isArray: false,
                params: {}
            }
        });
    }
    ])
    .factory('dtcBlockLocations', ['$window', '$resource', '$cacheFactory',
        function($window, $resource, $cacheFactory) {
            //Consume DTC Location service

            return $resource($window.dtServiceEndpoint + '/blockBooking/location', {}, {
                'query': {
                    method: 'GET',
                    isArray: false,
                    cache: $cacheFactory
                }
            });
        }
    ])
    .service('dtcBlockLocationData', ['$log', '$filter', 'dtcBlockLocations', 'dtcFilter',
        function($log, $filter, dtcBlockLocations, dtcFilter) {
            //Retieve DTC Location information

            var dtcData = {};

            function setData(data) {
                dtcData = data;
            }

            this.getBlockItems = function(locationId, customerIdToClasses) {
                var classes = customerIdToClasses.map(function(obj) { return obj.licenceClass});
                var location = $filter('filter')(dtcData, {id: parseInt(locationId)},true)[0];
                var services = $filter('filter')(location.services, function(value, index) {
                    return classes.indexOf(value.licenceClass) != -1;
                }, true);

                var blockItems = [];
                for (var i = 0; i < customerIdToClasses.length; i++) {
                    var currServiceId;
                    for (var j = 0; j < services.length; j++) {
                        if (services[j].licenceClass === customerIdToClasses[i].licenceClass) {
                            currServiceId =services[j].serviceId;
                        }
                    }
                    blockItems[i] = {customerId: customerIdToClasses[i].customerId, serviceId: currServiceId};
                    currServiceId = null;
                }
                return blockItems;
            };

            this.getData = function(showFrench, callback, errorCallback) {
                $log.debug("dtcBlockLocationData.getData: querying REST API");

                dtcBlockLocations.query({french: showFrench})
                    .$promise.then(function(results) {
                        var vettedResults = [];

                        angular.forEach(results.driveTestCentres, function(value, key) {
                            if (!angular.isUndefined(value.latitude) || !angular.isUndefined(value.longitude)) {
                                // process location
                                vettedResults.push(value);
                            }
                        });

                        setData(vettedResults);

                        if (callback) {
                            callback(dtcData);
                        } else {
                            return dtcData;
                        }
                    }, function(error) {
                        errorCallback(error);
                    });

            };
        }
    ])
    .service('dtcLocationData', ['$log', '$cacheFactory','$filter', 'dtcLocations', 'dotcmsTileServer', 'dotcmsLocations', 'dtcFilter', '$locale', '$injector', 'tileServerURLCache',
        function($log,$cacheFactory, $filter, dtcLocations, dotcmsTileServer, dotcmsLocations, dtcFilter, $locale, $injector, tileServerURLCache) {

            var bookingTimeService, userService;
            
            if ($injector.has('bookingTimeService') && $injector.has('userService')) {
                bookingTimeService = $injector.get('bookingTimeService');
                userService = $injector.get('userService');
            }

            //Retieve DTC Location information

            var dtcData = {};

            var dotcmsData = {};

            function setData(data) {
                dtcData = data;
            }

            function setDotcmsData(data) {
                dotcmsData = data;
            }

            this.getServiceId = function(locationId,licenceClass, endorsement, motorcycleType) {
                var licence = licenceClass;
                var bzfzClasses = ["BZ","CZ","DZ","EZ","FZ"];
                var bfClasses = ["B","C","D","E","F"];
                if(bzfzClasses.indexOf(licence) != -1) {
                    licence = "BZ-FZ";
                } else if (bfClasses.indexOf(licence) != -1) {
                    licence = "B-F";
                }
                var location = $filter('filter')(dtcData, {id: parseInt(locationId)},true)[0];
                return $filter('filter')(location.services, {licenceClass: licence}, true)[0].serviceId;

            };

            this.setUserTimeZone = function(locationId) {
                if (angular.isUndefined(dtcData.driveTestCentres)) {
                    dtcLocations.query()
                        .$promise.then(function(results) {
                            var timeZone;

                            angular.forEach(results.driveTestCentres, function(value, key) {
                                if (value.id == locationId){
                                    timeZone = value.timezone;
                                }
                            });
                            var user =  userService.getCurrentUser();
                            user.timezone = timeZone;
                            user.momentTimezone = bookingTimeService.getMomentTimezone(user.timezone);
                            userService.setUser(user);

                        }, function(error) {
                            errorCallback(error);
                        });

                } else if (angular.isDefined(dtcData.driveTestCentres)) {
                    $log.debug("dtcLocationData.getData: using cached data");
                } else {
                    errorCallback('something wasn\'t accounted for');
                }
            };

            function dayHours(day, start, end) {
                this.dayOFWeek = day;
                this.startTime = start;
                this.endTime = end;
            }

            function convertToArray (input) {

                var results = [];
                var hours = [];
                var count = 5;

                hours[0] = input.monday;
                hours[1] = input.tuesday;
                hours[2] = input.wednesday;
                hours[3] = input.thursday;
                hours[4] = input.friday;
                if(angular.isDefined(input.saturday) && input.saturday.length > 0) {
                    hours[5] = input.saturday;
                    count = 6;
                }

                for(var i = 0; i < count; i++) {
                    results[i] = new dayHours(i, hours[i].split(',')[0], hours[i].split(',')[1]);
                }

                return results;
            }

            this.getTileServerURL = function () {
                var rt = tileServerURLCache.get('tileServerURL');
                return  tileServerURLCache.get('tileServerURL');
            };
            this.setTileServerURL = function(tileServerURL){
                //locationCacheData.tileServerURL = tileServerURL;
                tileServerURLCache.put('tileServerURL', tileServerURL);
            };
            this.getData = function(licenceClass, endorsement, userInfo, callback, errorCallback) {

                var filter = licenceClass;
                if (endorsement != undefined) {
                    filter = licenceClass + endorsement;
                } else if (userInfo != undefined && userInfo.component != undefined) {
                    filter = userInfo.component + licenceClass;
                }

                var bfClasses = ["B","C","D","E","F"];
                if(bfClasses.indexOf(licenceClass) != -1) {
                    if(endorsement === "Z") {
                        filter = "BZ-FZ";
                    } else {
                        filter = "B-F";
                    }
                }

                if (angular.isUndefined(dtcData.driveTestCentres)) {
                    $log.debug("dtcLocationData.getData: querying REST API");
                    dtcLocations.query()
                        .$promise.then(function(results) {
                            var vettedResults = [];

                            angular.forEach(results.driveTestCentres, function(value, key) {
                                if (!angular.isUndefined(value.latitude) || !angular.isUndefined(value.longitude)) {

                                    //convert open hours to dates so we can use date format in the view
                                    for (var j = 0 ; j < value.locationHours.length; j++) {
                                        var locHours = value.locationHours[j];
                                        var locStartDate = new Date();
                                        locStartDate.setHours(locHours.startTime.substring(0, locHours.startTime.indexOf(':')));
                                        locStartDate.setMinutes(locHours.startTime.substring(locHours.startTime.indexOf(':') + 1));
                                        locHours.startDate = locStartDate;

                                        var locEndDate = new Date();
                                        locEndDate.setHours(locHours.endTime.substring(0, locHours.endTime.indexOf(':')));
                                        locEndDate.setMinutes(locHours.endTime.substring(locHours.endTime.indexOf(':') + 1));
                                        locHours.endDate= locEndDate;
                                    }
                                }

                                if(value.city === "Etobicoke") {
                                    var mfIndex = value.licenceTestTypes.indexOf("FM");
                                    if(mfIndex >= -1) {
                                        value.licenceTestTypes.splice(mfIndex, 1);
                                    }

                                    var m2FIndex = value.licenceTestTypes.indexOf("FM2");
                                    if(m2FIndex >= -1) {
                                        value.licenceTestTypes.splice(m2FIndex, 1);
                                    }
                                }
                                if(value.city === "Timmins") {
                                    var mLIndex = value.licenceTestTypes.indexOf("LM");
                                    if(mLIndex >= -1) {
                                        value.licenceTestTypes.splice(mLIndex, 1);
                                    }
                                    var mPIndex = value.licenceTestTypes.indexOf("PM");
                                    if(mPIndex >= -1) {
                                        value.licenceTestTypes.splice(mPIndex, 1);
                                    }
                                }

                                // process location
                                vettedResults.push(value);
                            });

                            setData(vettedResults);

                            if (!filter) {
                                if (callback) {
                                    callback(dtcData);
                                } else {
                                    return dtcData;
                                }
                            } else {
                                if (callback) {
                                    callback(dtcFilter.byTestTypes(filter, dtcData));
                                } else {
                                    return dtcFilter.byTestTypes(filter, dtcData);
                                }
                            }

                        }, function(error) {
                            errorCallback(error);
                        });

                } else if (angular.isDefined(dtcData.driveTestCentres)) {
                    $log.debug("dtcLocationData.getData: using cached data");
                    if (!filter) {
                        if (callback) {
                            callback(dtcData);
                        } else {
                            return dtcData;
                        }
                    } else {
                        if (callback) {
                            callback(dtcFilter.byTestTypes(filter, dtcData));
                        } else {
                            return dtcFilter.byTestTypes(filter, dtcData);
                        }
                    }

                } else {
                    errorCallback('something wasn\'t accounted for');
                }

            };

            this.getTileServer = function (callback, errorCallback) {

                dotcmsTileServer.query().$promise.then(
                    function(results) {
                        callback(results);
                    }, function(error) {
                        errorCallback(error);
                    }
                );
            };

            this.getDotcmsData = function (callback, errorCallback) {
                $log.debug("dtcLocationData.getDotcmsData: querying REST API");

                var lang = "1";
                if($locale.id === 'fr-ca') {
                    lang = "2";
                }

                dotcmsLocations.query({
                    language_id: lang
                }).$promise.then(
                    function(results) {
                        var vettedResults = [];

                        angular.forEach(results.driveTestCentres, function(value, key) {
                            if (!angular.isUndefined(value.latitude) || !angular.isUndefined(value.longitude)) {

                                //convert open hours to dates so we can use date format in the view
                                if (!angular.isDefined(value.locationhours.monday)) {
                                    $log.debug("Location hours not defined. Displaying notes field for " + value.name);
                                    value.travelpoint = true;

                                } else {
                                    value.locationhours = convertToArray(value.locationhours);

                                    for (var j = 0; j < value.locationhours.length; j++) {
                                        var locHours = value.locationhours[j];

                                        if ($locale.id === 'fr-ca') {
                                            locHours.startTime = locHours.startTime.split(":").join("h");
                                            locHours.startTime = locHours.startTime.split("00").join("");
                                            locHours.startDate = locHours.startTime;

                                            locHours.endTime = locHours.endTime.split(":").join("h");
                                            locHours.endTime = locHours.endTime.split("00").join("");
                                            locHours.endDate = locHours.endTime;

                                        } else {
                                            var locStartDate = new Date();
                                            locStartDate.setHours(locHours.startTime.substring(0, locHours.startTime.indexOf(':')));
                                            locStartDate.setMinutes(locHours.startTime.substring(locHours.startTime.indexOf(':') + 1));
                                            locHours.startDate = locStartDate;

                                            var locEndDate = new Date();
                                            locEndDate.setHours(locHours.endTime.substring(0, locHours.endTime.indexOf(':')));
                                            locEndDate.setMinutes(locHours.endTime.substring(locHours.endTime.indexOf(':') + 1));
                                            locHours.endDate = locEndDate;
                                        }
                                    }
                                }
                                value.address = value.address.replace(/Ã©/g,'\u00e9');
                                value.address = value.address.replace(/Ã»/g,'\u00fb');
                                value.address = value.address.replace(/à¨/g,'\u00e8');
                                value.address = value.address.replace(/Ã/g,'\u00e0');
                                if (value.travelpoint) {
                                    value.notes = value.notes.replace(/Ã©/g,'\u00e9');
                                    value.notes = value.notes.replace(/Ã»/g,'\u00fb');
                                    value.notes = value.notes.replace(/à¨/g,'\u00e8');
                                    value.notes = value.notes.replace(/Ã/g,'\u00e0');
                                }

                                //convert licence classes to lists
                                if (value.publiclicenceclasses) {
                                    if(value.publiclicenceclasses.slice(-1) == ",") {
                                        value.publiclicenceclasses = value.publiclicenceclasses.slice(0, value.publiclicenceclasses.length-1);
                                    }
                                    value.publiclicenceclasses = value.publiclicenceclasses.split(',');
                                }

                                if(value.commerciallicenceclasses) {
                                    if(value.commerciallicenceclasses.slice(-1) == ",") {
                                        value.commerciallicenceclasses = value.commerciallicenceclasses.slice(0, value.commerciallicenceclasses.length-1);
                                    }
                                    value.commerciallicenceclasses = value.commerciallicenceclasses.split(',');
                                }

                                // process location
                                vettedResults.push(value);
                            }
                        });

                        setDotcmsData(vettedResults);

                        callback(dotcmsData);

                    }, function(error) {
                        errorCallback(error);
                    }
                );
            };
        }
    ])
    .service('dtcFilter', [function() {
            //Filter DTC Results

            this.byTestTypes = function(filterStr, filterObj) {
                var toBeFiltered,
                    typeResults = [];

                if (angular.isDefined(filterObj.driveTestCentres)) {
                    toBeFiltered = filterObj.driveTestCentres;
                } else {
                    toBeFiltered = filterObj;
                }

                angular.forEach(toBeFiltered, function(parentValue, parentKey) {
                    if (angular.isDefined(parentValue.licenceTestTypes) && parentValue.licenceTestTypes.length > 0) {
                        angular.forEach(parentValue.licenceTestTypes, function(childValue, childKey) {
                            if (childValue === filterStr) {
                                typeResults.push(parentValue);
                            }
                        });
                    }
                });

                return typeResults;
            };

        }
    ])
    .filter('weekday', ['$locale', function ($locale) {
        //weekday starts Monday (0)
        var format = $locale.DATETIME_FORMATS;
        return function (date) {

            if(date !== 6) {
                date ++;
            } else {
                date = 0;
            }

            return format.DAY[date];
        };
    }]);
'use strict';

var dtServicesModule = angular.module('DriveTest.services');

dtServicesModule.factory('translateLoader', ['$http', '$q', 'locale',
    function($http, $q, locale) {
        return function(options) {
            var deferred = $q.defer();
            // do something with $http, $q and key to load localization files
            $http(angular.extend({
                url: [
                    options.prefix,
                    options.suffix
                ].join(''),
                method: 'GET',
                params: ''
            }, options.$http)).success(function(data) {
                // groom data for locale
                var dataObj = {};
                angular.forEach(data, function(value, key) {
                    if (key === locale.getLocaleStr()) {
                        dataObj = value;
                    }
                });
                deferred.resolve(dataObj);
            }).error(function(data) {
                deferred.reject(options.key);
            });
            return deferred.promise;
        };
    }
]);

'use strict';

//Booking App
angular.module('dtBooking', ['DriveTest.AutoSubmit', 'DriveTest.filters', 'DriveTest.services', 'DriveTest.LicenceExpiry', 'DriveTest.Email', 'DriveTest.LicenceNumber', 'DriveTest.Focus', 'DriveTest.directives', 'pascalprecht.translate', 'leaflet-directive', 'ngResource', 'ui.router', 'ui.bootstrap', 'ngRoute', 'DriveTest.findACentre', 'DriveTest.CompareValidation','DriveTest.DateValidation', 'ngAnimate', 'ngSanitize', 'DriveTest.ViewTitle', 'ngCookies', 'ngMessages', 'DriveTest.ButtonSubmit', 'vcRecaptcha'])
    .config(['$urlRouterProvider', '$stateProvider', '$provide', '$httpProvider', '$translateProvider',
        function($urlRouterProvider, $stateProvider, $provide, $httpProvider, $translateProvider) {

            $httpProvider.defaults.useXDomain = true;
            $httpProvider.defaults.withCredentials = true;
            delete $httpProvider.defaults.headers.common['X-Requested-With'];

            $httpProvider.interceptors.push(['$q','$location', '$cookies','$log',function($q,$location,$cookies,$log){
                return {
                    request: function(config){
                        $log.debug('interceptor config ---> ', config);

                        return config;
                    },

                    response: function(result){
                        $log.debug('interceptor result ---> ', result);

                        return result;
                    },

                    responseError: function(rejection){
                        //$log.debug('interceptor rejection ---> ', rejection);
                        //$log.debug('$location.path() --------> ', $location.path());

                        //place verify-drive/validate-driver-email 401 fail case here
                        if($location.path() !== '/validate-driver-email' && $location.path() !== '/verify-driver') {
                            if (rejection.status === 401) {
                                $location.path('quit-booking')
                            }
                        }

                        return $q.reject(rejection);
                    }
                };
            }]);

            $httpProvider.interceptors.push(['user', '$log', '$window', function (user, $log, $window) {
                return {
                    response: function (result) {
                        if (result.config.url.indexOf($window.dtServiceContext) >= 0) {
                            var btn = result.headers('transactionNumber');
                            user.btn = btn ? btn.substring(0, 8).toUpperCase() : '';
                            $log.debug('interceptor btn ---> ', btn, result.config.url);
                        }
                        return result;
                    }
                }
            }]);

            //if (document.URL.match(/nobackend/)) {
            //$provide.decorator('$httpBackend', angular.mock.e2e.$httpBackendDecorator);
            //}

            $translateProvider.useLoader('translateLoader', {
                // if you have some custom properties, go for it!
                prefix: '/application/error_codes',
                suffix: '.html'
            });
            $translateProvider.preferredLanguage('en_CA');
            $translateProvider.registerAvailableLanguageKeys(['en_CA', 'fr_CA'], {
                'en': 'en_CA',
                'en_US': 'en_CA',
                'fr': 'fr_CA',
                'fr_US': 'fr_CA'
            });
            //$translateProvider.useLocalStorage();

            $urlRouterProvider.otherwise('/validate-driver-email');

            var retrieveState = {
                sessionState: ['$q', 'userService', '$log', '$cookies', function ($q, userService, $log, $cookies) {

                    $log.debug('++ Token '+  $cookies.get('token'));

                    if (!userService.isSet()) {
                        var deferred = $q.defer();
                        userService.getDriverInformation(function (result) {
                            var user = userService.getCurrentUser();
                            user.customerId = result.customerId;
                            user.firstName = result.firstName;
                            user.lastName = result.lastName;
                            user.email = result.email;
                            userService.setUser(user);
                            deferred.resolve(true);
                        }, function (result) {
                            deferred.resolve(false);
                        });
                        return deferred.promise;
                    }
                    return true;
                }]
            };

            var bookingUserCheck = {
                sessionState: ['$q', 'userService', 'appointmentService', '$log', '$cookies', 'stepState',
                    function ($q, userService, appointmentService, $log, $cookies, stepState) {
                        $log.debug('-- Token '+  $cookies.get('token'));

                        appointmentService.getStatusToken(
                            userService.getCurrentUser().licenceNumber,
                            function (result) {
                                $log.debug("SUCCESS Login Status: " + result.statusToken);
                                $log.debug("Login status defined: " + angular.isDefined(result.statusToken));

                                if (!userService.isSet() || !angular.isDefined(result.statusToken)) {
                                    stepState.changeState('booking.licence');
                                    return false;
                                }
                                return true;

                            },
                            function (result) {
                                $log.debug("ERROR Login Status");
                                return false;
                            }
                        );
                    }]
            };

            var validateBooking = {
                sessionState: ['$q', 'userService', '$log', 'stepState', function ($q, userService, $log, stepState) {
                    $log.debug('-- Booking in progress: ' + userService.getCurrentUser().bookingInProgress);

                    if (!userService.getCurrentUser().bookingInProgress) {
                        stepState.changeState('booking.dashboard');
                        return false;
                    }

                    return true;
                }]
            };

            $stateProvider
                .state('booking', {
                    abstract: true,
                    resolve: {
                        // A string value resolves to a service
                        locale: 'locale',

                        // A function value resolves to the return
                        // value of the function
                        localeStr: function(locale) {
                            return locale.getLocaleStr();
                        }
                    },
                    templateUrl: '/ng/dt-booking'
                })
                .state('booking.registration', {
                    url: '/validate-driver-email',
                    allowAnonymous: true,
                    loginState: 'booking.licence',
                    views: {
                        'booking-registration': {
                            templateUrl: '/ng/dt-booking-reg',
                            controller: 'registrationController'
                        }
                    }
                })
                .state('booking.login', {
                    url: '/verify-driver',
                    allowAnonymous: true,
                    loginState: 'booking.dashboard',
                    views: {
                        'booking-login': {
                            templateUrl: '/ng/dt-booking-login',
                            controller: 'loginController'
                        }
                    }
                })
                .state('booking.verify', {
                    url: '/verify',
                    allowAnonymous: true,
                    views: {
                        'booking-verify': {
                            templateUrl: '/ng/dt-reg-verifyemail',
                            controller: 'verifyEmailController'
                        }
                    }
                })
                .state('booking.licence', {
                    url: '/licence',
                    resolve: retrieveState,
                    views: {
                        'booking-licence': {
                            templateUrl: '/ng/dt-booking-lic',
                            controller: 'licenceController'
                        }
                    }
                })
                .state('booking.location', {
                    url: '/location',
                    resolve: bookingUserCheck,
                    views: {
                        'booking-location': {
                            templateUrl: '/ng/dt-booking-find',
                            controller: 'locationController'
                        }
                    }
                })
                .state('booking.calendar', {
                    url: '/schedule',
                    resolve: bookingUserCheck,
                    views: {

                        'booking-location': {
                            templateUrl: '/ng/dt-booking-find',
                            controller: 'locationController'
                        },
                        'booking-calendar': {
                            templateUrl: '/ng/dt-booking-calendar',
                            controller: 'calendarController'
                        }
                    }
                })
                .state('booking.timeslot', {
                    url: '/timeslot',
                    resolve: bookingUserCheck,
                    views: {
                        'booking-location': {
                            templateUrl: '/ng/dt-booking-find',
                            controller: 'locationController'
                        },
                        'booking-calendar': {
                            templateUrl: '/ng/dt-booking-calendar',
                            controller: 'calendarController'
                        },
                        'booking-timeslot': {
                            templateUrl: '/ng/dt-booking-timeslot',
                            controller: 'timeslotController'
                        }
                    }
                })
                .state('booking.payment', {
                    url: '/payment',
                    resolve: bookingUserCheck,
                    views: {
                        'booking-payment': {
                            templateUrl: '/ng/dt-booking-payment',
                            controller: 'paymentController'
                        }
                    }
                })
                .state('booking.finalize', {
                    url: '/finalizing-booking',
                    resolve: validateBooking,
                    views: {
                        'booking-finalize': {
                            templateUrl: '/ng/dt-booking-finalize',
                            controller: 'finalizeBookingController'
                        }
                    }
                })
                .state('booking.complete', {
                    url: '/complete',
                    views: {
                        'booking-complete': {
                            templateUrl: '/ng/dt-booking-complete',
                            controller: 'completeController'
                        }
                    }
                })
                .state('booking.error', {
                    url: '/error',
                    allowAnonymous: true,
                    views: {
                        'booking-error': {
                            templateUrl: '/ng/dt-booking-error',
                            controller: 'errorController'
                        }
                    }
                })
                .state('booking.cancel', {
                    url: '/payment-cancel',
                    allowAnonymous: true,
                    views: {
                        'booking-cancel': {
                            templateUrl: '/ng/dt-booking-cancel',
                            controller: 'bookingCancel'
                        }
                    }
                })
                .state('booking.success', {
                    url: '/payment-success',
                    allowAnonymous: true,
                    views: {
                        'booking-cancel': {
                            templateUrl: '/ng/dt-booking-success',
                            controller: 'bookingSuccess'
                        }
                    }
                })
                .state('booking.submit', {
                    url: '/payment-submit',
                    allowAnonymous: true,
                    views: {
                        'booking-cancel': {
                            templateUrl: '/ng/dt-booking-submit',
                            controller: 'paymentSubmitController'
                        }
                    }
                })
                .state('booking.dashboard', {
                    url: '/dashboard',
                    resolve: retrieveState,
                    views: {
                        'booking-dashboard': {
                            templateUrl: '/ng/dt-appointments',
                            controller: 'dashboardController'
                        }
                    }
                })
                .state('booking.quit', {
                    url: '/quit-booking',
                    allowAnonymous: true,
                    views: {
                        'booking-quit': {
                            template: '',
                            controller: 'quitController'
                        }
                    }
                });
        }
    ])
    .run([
        '$rootScope',
        '$state',
        '$location',
        '$stateParams',
        '$log',
        'stepState',
        'userService',
        '$cookies',
        '$http',
        'appointmentService',
        function(
            $rootScope,
            $state,
            $location,
            $stateParams,
            $log,
            stepState,
            userService,
            $cookies,
            $http,
            appointmentService
        ) {
            $rootScope.$on("$stateChangeStart", function(event, toState, toParams, fromState, fromParams){
                $('.page-loading').show();

                $log.debug('++++ Token '+  $cookies.get('token'));

                appointmentService.getStatusToken(
                    userService.getCurrentUser().licenceNumber,
                    function (result) {
                        $log.debug("SUCCESS Login Status: " + result.statusToken);
                        $log.debug("Login status defined: " + angular.isDefined(result.statusToken));

                        var toStateProtected = !angular.isDefined(toState.allowAnonymous) || !toState.allowAnonymous;
                        if(toStateProtected && !(userService.isSet() || angular.isDefined(result.statusToken))){
                            event.preventDefault();
                            stepState.changeState('booking.quit');
                        } else if(toStateProtected && (!userService.isSet() && angular.isDefined(result.statusToken) && toState.name !== 'booking.dashboard' && toState.name !== 'booking.licence')){
                            event.preventDefault();
                            stepState.changeState('booking.dashboard');
                        } else if(!toStateProtected && angular.isDefined(toState.loginState) && (userService.isSet() || angular.isDefined(result.statusToken))) {
                            event.preventDefault();
                            if (fromState !== undefined && fromState.name === toState.loginState) {
                                stepState.changeState(toState.loginState, undefined, { reload: true });
                                $('.page-loading').hide();
                            } else {
                                stepState.changeState(toState.loginState);
                            }
                        }

                    },
                    function (result) {
                        $log.debug("ERROR Login Status");
                    }
                );

            });

            $rootScope.$on('$viewContentLoaded', function(event, toState, toParams, fromState, fromParams){
                $('.page-loading').hide();
            });
        }
    ]);
'use strict';

//Contact Us App
angular.module('dtContactUs', ['DriveTest.Email', 'ngResource', 'ui.router', 'ui.bootstrap', 'ngRoute','ngSanitize', 'ContactUsForm.Firstname', 'ContactUsForm.Lastname', 'ContactUsForm.Comments', 'ContactUsForm.Focus', 'ngCookies', 'ngMessages', 'vcRecaptcha'])
    .config(['$urlRouterProvider', '$stateProvider', '$provide', '$httpProvider',
        function($urlRouterProvider, $stateProvider, $provide, $httpProvider) {
            $urlRouterProvider.otherwise("/contact-us");
            $stateProvider
                .state('contact-us', {
                    url: '/contact-us',
                    templateUrl: '/ng/dt-contact-us',
                    controller: 'contactUsController'
                })
                .state('contact-us-success', {
                    url: '/contact-us-success',
                    templateUrl: '/ng/dt-contact-us-success',
                    controller: 'contactUsSuccessController',
                    params: {
                        'success': 'true'
                    }
                });
        }
    ]);

/**
 * Created by Hisham on 2014-09-02.
 */
'use strict';
angular.module('dtMaps', [ 'leaflet-directive', 'ui.router', 'ngResource', 'DriveTest.findACentre'])
    .config(['$urlRouterProvider', '$stateProvider', '$provide', '$httpProvider', function($urlRouterProvider, $stateProvider, $provide, $httpProvider) {

        $urlRouterProvider.otherwise('/locations');

        $stateProvider
            .state('locations', {
                templateUrl: '/ng/dt-locations',
                url: '/locations',
                controller: 'mapController'
            });
    }]);

(function(angular){
    'use strict';

    angular
        .module('dtBlockBooking')
        .controller('blockAddStudentController', ['$scope', 'blockSignupService', 'errorService', 'studentListService' , 'stepState', 'studentBlock', blockAddStudentController]);

    function blockAddStudentController($scope, blockSignupService, errorService, studentListService, stepState, studentBlock){
        $scope.studentBlock = studentBlock.getCurrent();
        $scope.student = {
            licenceNumber: '',
            licenceExpiry: '',
            licenceType: '',
            testLanguage: 'en_ca'
        };

        $scope.school = blockSignupService.getBlockSignupData().drivingSchool;
        $scope.loading = false;
        $scope.errMsg = {"hasError": false, "msg":""};

        var resetStudent = angular.copy($scope.student);

        $scope.licenceClass = [
            {value: 'A', class: 'A', endorsement: null},
            {value: 'B', class: 'B', endorsement: null},
            {value: 'C', class: 'C', endorsement: null},
            {value: 'D', class: 'D', endorsement: null},
            {value: 'E', class: 'E', endorsement: null},
            {value: 'F', class: 'F', endorsement: null},
            {value: 'Z', class: 'Z', endorsement: null},
            {value: 'AZ', class: 'A', endorsement: 'Z'},
            {value: 'BZ', class: 'B', endorsement: 'Z'},
            {value: 'CZ', class: 'C', endorsement: 'Z'},
            {value: 'DZ', class: 'D', endorsement: 'Z'},
            {value: 'EZ', class: 'E', endorsement: 'Z'}
        ];

        $scope.addStudentFormSubmit = function(student){

            if($scope.errMsg.hasError){
                $scope.errMsg = {hasError:false};
                $scope.errMsg.msg = '';
            }
            var newStudent = angular.copy(student);
            $scope.loading = true;

            studentListService.checkStudentEligibility(newStudent, function(response){

                $scope.loading = false;

                // add the student to the full list, with additional data

                /* {
                *      "eligible": true,
                *      "from": "2015-04-12 EDT",
                *      "to": "2016-02-15 EST",
                *      "customerId": "S0000000032"
                *      "firstName": "Cox",
                *      "lastName": "Sinthia"
                    } */

                newStudent.from = response.from;
                newStudent.to = response.to;
                newStudent.customerId = response.customerId;
                newStudent.firstName = response.firstName;
                newStudent.lastName = response.lastName;
                newStudent.licenceTypeValue = newStudent.licenceType.value;
                newStudent.licenceClass = newStudent.licenceType.class;
                newStudent.endorsement = newStudent.licenceType.endorsement;
                delete newStudent.checkSameLicense;
                delete newStudent.licenceType;

                if(!studentBlock.getCurrent(newStudent.customerId) && studentBlock.hasIntersectingEligibleDates(newStudent)) {
                    studentBlock.addStudent(newStudent, newStudent.customerId);

                    //reset form
                    $scope.student = angular.copy(resetStudent);
                    $scope.addStudentForm.$setPristine();

                } else if (!studentBlock.hasIntersectingEligibleDates(newStudent)) {
                    $scope.errMsg = {hasError:true};
                    $scope.errMsg.msg = errorService.getMessage('blockbooking.app.message.eligibility.no.intersecting.dates');
                } else {
                    $scope.errMsg = {hasError:true};
                    $scope.errMsg.msg = errorService.getMessage('blockbooking.app.message.eligibility.used.licence');
                }

            }, function(errorResponse){

                $scope.loading = false;
                var messageKey, returnVal = false;

                switch(errorResponse.status){

                    /* 1. Cancellation History (403) - the user has cancelled 3 appointments within the last 6 months <br>
                    * 2. Duplicate Licence (412) - the user has already booked the licence class specified in the eligibilityCheck parameter <br>
                    * 3. MTO Driver Eligibility (405) - the user is ineligible due to their mto driver record's status <br>
                    */

                    case 403:
                        messageKey = 'blockbooking.app.message.eligibility.ineligible.3.6.check';
                        returnVal = true;
                        break;
                    case 412:
                        messageKey = 'blockbooking.app.message.eligibility.ineligible.licence';
                        returnVal = true;
                        break;
                    case 405:
                        messageKey = 'blockbooking.app.message.eligibility.ineligible.record';
                        returnVal = true;
                        break;
                    case 417:
                        messageKey = 'blockbooking.app.message.eligibility.ineligible.number';
                        returnVal = true;
                        break;
                    default:
                        $scope.otherError = true;
                }

                $scope.errMsg = {hasError:true};
                $scope.errMsg.msg = errorService.getMessage(messageKey);

                return returnVal;
            });
        };

        $scope.listSubmit = function(){
            if ($scope.canContinue()) {
                stepState.changeState('school.location');
            }
        };

        $scope.canContinue = function(){
            return (studentBlock.size()>0);
        };

        $scope.removeStudent = function(evt, customerId){
            evt.preventDefault();

            return studentBlock.removeStudent(customerId);
        };

    }

}(window.angular));
/**
 * Created by at_user on 09/04/2015.
 */

(function(angular){
    'use strict';

    angular
        .module('dtBlockBooking')
        .controller('blockForgotPasswordController', ['$scope', 'loginService', 'errorService', 'stepState', blockForgotPasswordController]);

    function blockForgotPasswordController($scope, loginService, errorService){

        $scope.emailRecovered = false;
        $scope.loading = false;

        $scope.submitEmail = function(schoolEmail){

            $scope.loading = true;

            loginService.recover(schoolEmail, function(response){
                    $scope.loading = false;
                    $scope.emailRecovered = true;
                },
                errorService.errorHandler($scope, function(response){

                    $scope.loading = false;
                    var messageKey, returnVal = false;

                    switch(response.status){
                        case 400:
                            messageKey = 'booking.errorcode.email_invalid';
                            returnVal = true;
                            break;
                        case 404:
                            messageKey = 'blockbooking.app.message.account.not.found';
                            returnVal = true;
                            break;
                        case 424:
                            messageKey = 'blockbooking.app.message.account.pending.review';
                            returnVal = true;
                            break;
                    }

                    $scope.errMsg = {hasError:true};
                    $scope.errMsg.msg = errorService.getMessage(messageKey);

                    return returnVal;
                })
            );
        };
    }

}(window.angular));
'use strict';

angular.module('dtBlockBooking')
    .controller('blockLocationController', ['$scope', '$rootScope', '$window', '$log','$state', 'stepState', 'dtcBlockLocationData', 'errorService', 'studentBlock', 'blockSignupData', function($scope, $rootScope, $window, $log,$state,stepState,dtcBlockLocationData, errorService,studentBlock, blockSignupData) {
    $scope.locationInfo = {};
    $log.log('blockLocationController $state ---> ', $state);

    $scope.master = {};

    $scope.frenchOnly = studentBlock.frenchRequested();
    $scope.openSaturday = false;
    $scope.hideFrenchOption = true;

    $scope.broadcastSaturdayOrFrench = function(openSaturday) {
        $scope.openSaturday = openSaturday;
        $rootScope.$broadcast("refresh-map-saturday-french", openSaturday, $scope.showFrench);

        $log.debug('dtBooking.blockLocationController - broadcastSaturdayOrFrench ---> ', $scope);
    };

    $scope.loading = false;
    $scope.errMsg = errorService.clean($scope);

    $scope.errorCallback = function(error) {
        $log.error('loaded with error ---> ', error);
        stepState.goToGeneralErrorState();
    };

    $scope.previousStep = function() {
        stepState.changeState('school.addbooking');
        return;
    };

    $scope.mapCallback = function(selectedEl, location) {
        $log.log('Booking Page --> callback to maps directive', selectedEl, location);

        blockSignupData.booking.locationName = selectedEl.attr("title");
        blockSignupData.booking.locationId = selectedEl.attr("id");
        blockSignupData.booking.serviceIds = dtcBlockLocationData.getBlockItems(blockSignupData.booking.locationId, studentBlock.getCustomerIdsToLicenceClasses());
        blockSignupData.booking.locationTimeZone = location.timezone;

        if (stepState.currentStepName() === 'school.timeslot') {
            stepState.changeState('school.timeslot', true);
        }

    };

    $scope.locationSubmit = function() {
        stepState.changeState('school.timeslot');
    };

}]);
'use strict';

angular.module('dtBlockBooking')
    .controller('blockLoginController', ['$log', '$scope', 'loginService', 'blockSignupData', 'errorService', 'stepState',
        function($log, $scope, loginService, blockSignupData, errorService, stepState){
            $scope.school = {};

            $scope.schoolLoginSubmit = function(school){

                loginService.login(school,

                    function(result) {
                        //success
                        blockSignupData.drivingSchool = result.data.drivingSchool;
                        stepState.changeState('school.dashboard');

                    }, function(errorResponse){
                        var messageKey, returnVal = false;

                        switch(errorResponse.status){
                            case 401:
                                messageKey = 'blockbooking.app.message.invalid.credentials';
                                returnVal = true;
                                break;
                            case 403:
                                messageKey = 'blockbooking.app.message.account.suspended';
                                returnVal = true;
                                break;
                            case 406:
                                messageKey = 'blockbooking.app.message.account.rejected';
                                $scope.expiredTokenError = true;
                                returnVal = true;
                                $scope.sameLicenceError = true;
                                break;
                            case 417:
                                messageKey = 'blockbooking.app.message.account.pending.verification';
                                returnVal = true;
                                break;
                            case 424:
                                messageKey = 'blockbooking.app.message.account.pending.review';
                                returnVal = true;
                                break;
                            default:
                                $scope.otherError = true;
                        }

                        $scope.errMsg = {hasError:true};
                        $scope.errMsg.msg = errorService.getMessage(messageKey);

                        return returnVal;
                    });
            };
        }
    ]);
'use strict';

var dtBookingApp = angular.module('dtBlockBooking');

dtBookingApp.controller('blockPaymentSubmitController', ['$window', '$scope', '$rootScope', '$log', '$state', '$http', '$sce', 'stepState', 'studentBlock', 'blockSignupData', 'locale',
    function($window, $scope, $rootScope, $log, $state, $http, $sce, stepState, locale) {

        var blockSignupData = window.opener.$windowScope.blockSignupData;
        var studentBlock = window.opener.$windowScope.studentBlock;
        var bookingFeeRequest = {
            tests: [],
            guid: blockSignupData.booking.blockGuid
        };

        var block = studentBlock.getCurrent();
        var key;
        for (key in block) {
            if (block.hasOwnProperty(key)) {
                bookingFeeRequest.tests.push(
                    {
                        licenceNumber: block[key].licenceNumber,
                        licenceClass: block[key].licenceClass,
                        endorsement: block[key].endorsement
                    }
                );
            }
        }

        //TODO: Eligibility check????
        $http({
            url : $window.dtServiceEndpoint + "/blockBooking/payment",
            method : "POST",
            data: bookingFeeRequest,
            dataType:"json",
            headers: {
                "Content-Type" : "application/json; charset=utf-8",
                'Accept-Language': locale.getLocaleStr(),
                "Accept" : "application/json"}

        }).success(function (data, status, headers, config) {
            if (data.status && data.status.errorCode) {
                stepState.goToGeneralErrorState();

            } else {
                data.url = $sce.trustAsResourceUrl(data.url);
                $rootScope.$broadcast('gateway.redirect', data);
            }

        }).error(function (data, status, headers, config) {
            // The service call failed
        });
    }
]);
'use strict';

angular.
    module('dtBlockBooking')
    .controller('blockQuitController', ['$scope','blockSignupService','$window', 'stepState', function($scope,blockSignupService,$window, stepState){

    function cleanSession(){
        document.cookie = 'ssid' +'=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
        blockSignupService.clearBlockSignupData();
        $window.location.href = '/book-a-road-test/';
    }
    blockSignupService.removeToken(cleanSession, cleanSession);
}]);

/**
 * Created by at_user on 13/04/2015.
 */

//blockResetPasswordController

(function(angular){
    'use strict';

    angular
        .module('dtBlockBooking')
        .controller('blockResetPasswordController', ['$scope', '$log', 'stepState', 'loginService', '$stateParams', 'errorService', blockResetPasswordController]);

    function blockResetPasswordController($scope, $log, stepState, loginService, $stateParams, errorService){

        var token = $stateParams.token;
        $scope.loading = false;
        $scope.resetSuccess = false;
        $scope.passwordRegex = /^[\S]*$/;

        if(angular.isUndefined(token)){
            stepState.changeState('school.login');
        }

        $scope.submitReset = function(schoolReset){

            $scope.loading = true;

            loginService.reset({
                     "password": schoolReset.password,
                     "emailToken": token
                },
                function(response){
                    $scope.loading = false;
                    $scope.resetSuccess = response.data.success;
                },
                errorService.errorHandler($scope, function(response){
                    $scope.loading = false;
                    var messageKey, returnVal = false;

                    switch(response.status){
                        case 400:
                        case 404:
                            messageKey = 'booking.app.verifyemail.otherError';
                            returnVal = true;
                            break;
                        case 409:
                            messageKey = 'booking.app.message.email.token.expired';
                            returnVal = true;
                            break;
                        case 410:
                            messageKey = 'booking.errorcode.invalid_email_token';
                            returnVal = true;
                            break;
                    }

                    $scope.errMsg = {hasError:true};
                    $scope.errMsg.msg = errorService.getMessage(messageKey);

                    return returnVal;
                })
            );
        };

    }
}(window.angular));
'use strict';

var dtBlockBookingApp = angular.module('dtBlockBooking');

dtBlockBookingApp.controller('blockSignupController', ['$scope', '$log', '$state', '$routeParams', 'locale', 'blockSignupService', 'errorService', 'stepState',
    function($scope, $log, $state ,$routeParams, locale, blockSignupService, errorService, stepState) {
        $scope.school = { };
        $scope.captcha_response = null;
        $log.log('blockBookingController $state ---> ', $state);

        $scope.master = {};
        //$scope.completed = false;

        $scope.loading = false;
        $scope.errMsg = errorService.clean($scope);
        $scope.provinceList = [
            {value: 'AB'},
            {value: 'BC'},
            {value: 'MB'},
            {value: 'NB'},
            {value: 'NL'},
            {value: 'NS'},
            {value: 'NT'},
            {value: 'NU'},
            {value: 'ON'},
            {value: 'PE'},
            {value: 'QC'},
            {value: 'SK'},
            {value: 'YT'}
        ];

        $scope.postalCodeValidator = /^[ABCEGHJKLMNPRSTVXY]\d[ABCEGHJKLMNPRSTVWXYZ]([ -])?\d[ABCEGHJKLMNPRSTVWXYZ]\d/i;
        $scope.phoneNumberRegex = /^(?:(?:\+?1\s*(?:[.-]\s*)?)?(?:\(\s*([2-9]1[02-9]|[2-9][02-8]1|[2-9][02-8][02-9])\s*\)|([2-9]1[02-9]|[2-9][02-8]1|[2-9][02-8][02-9]))\s*(?:[.-]\s*)?)?([2-9]1[02-9]|[2-9][02-9]1|[2-9][02-9]{2})\s*(?:[.-]\s*)?([0-9]{4})(?:\s*(?:#|x\.?|ext\.?|extension)\s*(\d+))?$/;
        $scope.passwordRegex = /^[\S]*$/;

        $scope.removeDuplicateError = function() {
            $scope.schoolForm.username.$setValidity('duplicate',true);
        };


        $scope.schoolSubmit = function() {

            //if ($scope.schoolForm.$valid) {

            $scope.loading = true;

            var registrationData = blockSignupService.getBlockSignupData();
            registrationData.username = $scope.school.username;
            registrationData.password = $scope.school.password;
            registrationData.passwordConfirm = $scope.school.passwordConfirm;
            registrationData.consentToBookTestsOnBehalfOfStudents = $scope.school.consent;
            registrationData.captchaResponse = $scope.captcha_response;

            var drivingSchool = registrationData.drivingSchool;
            drivingSchool.id = $scope.id;
            drivingSchool.schoolName = $scope.school.schoolName;
            drivingSchool.businessNumber = $scope.school.businessNumber;
            //drivingSchool.address = (angular.isDefined($scope.school.schoolAddress2) ? $scope.school.schoolAddress2 + ' ' : '') +$scope.school.address;
            drivingSchool.address = $scope.school.address;
            drivingSchool.province = $scope.school.province.value;
            drivingSchool.postalCode = $scope.school.postalCode;
            drivingSchool.city = $scope.school.city;
            drivingSchool.contactPerson = $scope.school.contactPerson;
            drivingSchool.contactEmail = $scope.school.contactEmail;
            drivingSchool.contactPhoneNumber = $scope.school.contactPhoneNumber;

            blockSignupService.register(function(result){

                result = result.data;
                $scope.loading = false;

                if(result.success && angular.isDefined(result.drivingSchoolId)) {
                    // Go to verify email
                    //$scope.completed = true;
                    $log.debug('Registration completed ---> ', result);
                    stepState.changeState('school.verifyEmail');

                } else {
                    if(angular.isDefined(result.validationErrors)) {
                        angular.forEach(result.validationErrors, function (value, key) {
                            if (key.contains(".")) {
                                var errorCodes = key.split('.');
                                $scope.schoolForm[errorCodes[0]].$setValidity(errorCodes[1], false);
                                //$scope.errMsg.msg += errorService.getMessage(value) + "<br />" ;
                                //$scope.errMsg.hasError = true;
                            }
                        });
                    } else {
                        errorService.showServerError($scope, result);
                    }
                }
            }, errorService.errorHandler($scope, function(response){
                $scope.loading = false;
                var messageKey, returnVal = false;

                //account, a HTTP 417 Expectati
                //moved or disabled, a HTTP 403
                //a HTTP 406 Not Acceptable res
                //P 424 Failed Dependency respon

                switch(response.status){
                    case 400:
                        messageKey = 'blockbooking.app.field.missing';
                        returnVal = true;
                        break;
                    case 406:
                        messageKey= 'blockbooking.app.school.email.exists';
                        returnVal = true;
                        break;
                    case 409:
                        messageKey = 'blockbooking.app.school.name.exists';
                        returnVal = true;
                        break;
                    case 412:
                        messageKey = 'booking.app.registration.captcha.error.invalid';
                        returnVal = true;
                        break;
                }

                $scope.errMsg = {hasError:true};
                $scope.errMsg.msg = errorService.getMessage(messageKey);
                return returnVal;
            }));


            //}
        };
        $scope.setResponse = function(response) {
            $scope.captcha_response = response;
        };
        $scope.resetServerSideValidation = function(element, error){
            if(angular.isDefined($scope.schoolForm[element])) {
                $scope.schoolForm[element].$setValidity(error, true);
            }
        };


    }
]);

'use strict';

angular.module('dtBlockBooking')
    .controller('blockSummaryController', ['$scope', '$rootScope', '$window', '$log', '$state', 'stepState', 'errorService', 'studentBlock', 'blockSignupData',
        function ($scope, $rootScope, $window, $log, $state, stepState, errorService, studentBlock, blockSignupData) {
            $log.log('blockPaymentController $state ---> ', $state);
            stepState.checkStepState('school.summary');
            $scope.loading = false;
            $scope.errMsg = errorService.clean($scope);
            $scope.studentBlock = studentBlock;

            // TEST DATA
            //studentBlock.addStudent(
            //    {
            //        "firstName": "John",
            //        "lastName": "Smith",
            //        "licenceNumber": "licenceNumber1",
            //        "licenceClass": "A",
            //        "endorsement": "Z"
            //    },
            //    "s1");
            //
            //studentBlock.addStudent(
            //    {
            //        "firstName": "Jane",
            //        "lastName": "Smith",
            //        "licenceNumber": "licenceNumber2",
            //        "licenceClass": "C",
            //        "endorsement": null
            //    },
            //    "s2");
            //
            //// TEST DATA
            //blockSignupData.booking = {
            //    locationName: "Bancroft",
            //    locationId: 1,
            //    serviceIds: [1, 2, 3],
            //    blockGuid: "12345",
            //    timeslot: new Date().getTime(),
            //    fees: [
            //        {
            //            "licenceNumber": "licenceNumber1",
            //            "fees": {
            //                "roadTestFee": {"feeAmount": 10.00},
            //                "outOfOrderFee": {"feeAmount": 1.00},
            //                "providerRoadTestFee": {"feeAmount": 8.00},
            //                "providerOutOfOrderFee": {"feeAmount": .80}
            //            }
            //        },
            //        {
            //            "licenceNumber": "licenceNumber2",
            //            "fees": {
            //                "roadTestFee": {"feeAmount": 20.00},
            //                "outOfOrderFee": {"feeAmount": 2.00},
            //                "providerRoadTestFee": {"feeAmount": 16.00},
            //                "providerOutOfOrderFee": {"feeAmount": 1.60}
            //            }
            //        }
            //    ],
            //    details: {
            //        "guid": "12345",
            //        "publicUserId": "UserID",
            //        "appointments": [
            //            {"userId": "s1", "time": new Date().getTime() + 1000000},
            //            {"userId": "s2", "time": new Date().getTime() + 5000000}
            //        ]
            //    },
            //    feeTotal: 100.00
            //};

            $scope.blockSignupData = blockSignupData;
            $scope.booking = blockSignupData.booking;
            $scope.students = studentBlock.getCurrent();
            $scope.feeTotal = blockSignupData.booking.feeTotal;

            var appointments = blockSignupData.booking.details.appointments;

            if (appointments !== undefined) {

                angular.forEach(appointments, function (appointment) {
                    var student = studentBlock.getCurrent(appointment.userId);
                    student.appointmentTime = appointment.time;
                });
            }

            var fees = blockSignupData.booking.fees;

            if (fees !== undefined) {

                angular.forEach(fees, function (fee) {
                    studentBlock.getStudent(fee.licenceNumber).fees = fee.fees;
                });
            }

            $scope.fees = fees;

            $scope.errorCallback = function (error) {
                $log.error('loaded with error ---> ', error);
                stepState.goToGeneralErrorState();
            };

            $scope.previousStep = function () {
                stepState.changeState('school.timeslot');
            };

            $scope.paymentCancelled = function () {
                $log.log('Payment Cancelled');
            };

            $scope.editBooking = function(){
                stepState.changeState("school.location");
            };

            $scope.payNow = function(){
                $window.open('#/school-payment-submit',
                    '_blank', 'height=' + $window.innerHeight + ',width=1005');
            };


            $scope.paymentSuccessful = function(confirmationNumber, id) {
                $log.log('Payment Successful '+confirmationNumber);
                alert("Payment Successful: confirmationNumber=" + confirmationNumber + " id=" + id);
                stepState.changeState("school.finalize");
            };

            $scope.paymentError = function(response) {
                $log.log('Payment Error');

                errorService.errorHandler($scope, function(response){
                    var messageKey = 'booking.error', returnVal = false;

                    if(response.status === 400 || response.status === 500 ) {
                        messageKey = 'booking.app.message.incomplete.payment';
                        returnVal = true;
                    }

                    $scope.errMsg = {hasError:true};
                    $scope.errMsg.msg = errorService.getMessage(messageKey);

                    return returnVal;
                });
            };

            window.$windowScope = $scope;
        }
    ]
);

'use strict';

angular.module('dtBlockBooking')
    .controller('blockTimeslotController', ['$scope', '$rootScope', '$window', '$log','$state', 'stepState', 'errorService', 'blockTimeslotService', 'studentBlock', 'blockSignupData', 'blockSummaryService', 'locale',
        function($scope, $rootScope, $window, $log, $state, stepState, errorService, blockTimeslotService, studentBlock, blockSignupData, blockSummaryService, locale) {
            $scope.timeSlotInfo = {};
            $log.log('blockTimeslotController $state ---> ', $state);

            $scope.master = {};

            stepState.checkStepState('school.timeslot');

            $scope.loading = false;
            $scope.errMsg = errorService.clean($scope);
            $scope.selectedTimeslot = undefined;

            var eligibleDates = studentBlock.getEligibleDates();
            $scope.data= {};
            $scope.data.endDate= moment(eligibleDates.to, 'YYYY-MM-DD');

            $scope.dtcTimezone = blockSignupData.booking.locationTimeZone;
            var dtcMomentTimezone = blockTimeslotService.getMomentTimezone(blockSignupData.booking.locationTimeZone);

            var tomorrow = blockTimeslotService.getTomorrowsDate(blockSignupData.booking.locationTimeZone);
            $scope.data.startDate=moment(eligibleDates.from, 'YYYY-MM-DD').isBefore(tomorrow, 'day') ? tomorrow : moment(eligibleDates.from, 'YYYY-MM-DD');


            $scope.data.calendarWidgetStartDate=$scope.data.startDate.format('YYYY-MM-DD');

            $scope.inputOnTimeSet = function (newDate) {
                $scope.timeslots = undefined;
                $scope.sameDayAvailable = true;
                $scope.loading = true;

                var reqDate = moment(newDate).format('YYYY/MM/DD');
                var reqLastEligibleDate = moment($scope.data.endDate, 'YYYY-MM-DD').format('YYYY/MM/DD');
                var services = blockSignupData.booking.serviceIds.map(function (obj) { return obj.serviceId; });
                var theRequest = {startDate:reqDate, lastEligibleTestDateAllowed:reqLastEligibleDate, serviceIds:services};

                blockTimeslotService.getTimeslots(theRequest, function(response) {
                    $scope.loading = false;
                    $scope.timeslots = response.availableBlocks;
                    for (var i = 0 ; i < $scope.timeslots.length; i++) {
                        var temp = $scope.timeslots[i];
                        $scope.timeslots[i] = blockTimeslotService.getFormattedDate(temp.timeslot, dtcMomentTimezone);
                    }
                    $scope.sameDayAvailable = response.sameDayAvailable;
                    $scope.selectedTimeslot = undefined;

                    if($scope.errMsg.hasError){
                        $scope.errMsg = {hasError:false};
                        $scope.errMsg.msg = '';
                    }

                }, function(errorResponse) {

                    $scope.loading = false;
                    var messageKey, returnVal = false;

                    switch(errorResponse.status){

                        case 410:
                            messageKey = 'blockbooking.app.message.timeslot.0returned';
                            returnVal = true;
                            break;
                        default:
                            $scope.otherError = true;
                    }

                    $scope.errMsg = {hasError:true};
                    $scope.errMsg.msg = errorService.getMessage(messageKey);

                    return returnVal;
                });
            };

            $scope.beforeRender = function ($view, $dates, $leftDate, $upDate, $rightDate) {

                for (var i = 0 ; i < $dates.length; i++) {
                    var afterStart = moment.utc($dates[i]['utcDateValue']).isAfter($scope.data.startDate, $view);
                    var sameStart = moment.utc($dates[i]['utcDateValue']).isSame($scope.data.startDate, $view);
                    var beforeEnd = moment.utc($dates[i]['utcDateValue']).isBefore($scope.data.endDate, $view);
                    var sameEnd = moment.utc($dates[i]['utcDateValue']).isSame($scope.data.endDate, $view);

                    if ((afterStart || sameStart) && (beforeEnd || sameEnd)) {
                        $dates[i].selectable = true;
                    } else {
                        $dates[i].selectable = false;
                    }
                }
            };

            $scope.configDatePicker= function configDatePicker() {
                return {
                    startView: 'day',
                    minView: 'day',
                    dropdownSelector: '#dropdown6'
                };
            };

            $scope.errorCallback = function(error) {
                $log.error('loaded with error ---> ', error);
                stepState.goToGeneralErrorState();
            };

            $scope.previousStep = function() {
                stepState.changeState('school.location');
            };

            $scope.timeslotSubmit = function() {
                blockTimeslotService.holdBlock(
                    $scope.selectedTimeslot,
                    blockSignupData.booking.serviceIds,
                    function(response) {
                        blockSignupData.booking.blockGuid = response.blockGuid;
                        blockSignupData.booking.timeslot = $scope.selectedTimeslot;
                        _getSummary();
                    },
                    function(errorResponse) {
                        $scope.loading = false;
                        var messageKey, returnVal = false;

                        switch(errorResponse.status){
                            //TODO error check
                            case 410:
                                messageKey = 'blockbooking.app.message.timeslot.0returned';
                                returnVal = true;
                                break;
                            default:
                                $scope.otherError = true;
                        }

                        $scope.errMsg = {hasError:true};
                        $scope.errMsg.msg = errorService.getMessage(messageKey);

                        return returnVal;
                    });
            };

            $scope.canContinue = function() {
                return ($scope.selectedTimeslot!=undefined)
            };

            $scope.noTimeslots = function () {
                return $scope.timeslots === undefined || $scope.timeslots.length === 0;
            };

            //load the first available timeslots
            $scope.inputOnTimeSet($scope.data.calendarWidgetStartDate);

            var _getSummary = function() {
                var studentTests = [];

                angular.forEach(studentBlock.getCurrent(), function(student) {

                    studentTests.push(
                        {"licenceNumber": student.licenceNumber,
                        "licenceClass": student.licenceClass,
                        "endorsement": student.endorsement}
                    );
                });

                var booking = {
                    "guid": blockSignupData.booking.blockGuid,
                    "tests": studentTests
                };

                blockSummaryService.getPaymentSummary(
                    booking,
                    function(response) {
                        blockSignupData.booking.details = response.blockDetails;
                        blockSignupData.booking.fees = response.bookingFees;
                        blockSignupData.booking.feeTotal = response.total;
                        stepState.changeState('school.summary');
                    },
                    errorService.errorHandler($scope, function() {
                        $scope.loading = false;
                        errorService.setBlocked($scope);
                    }));
            };
        }]);

/**
 * Created by at_user on 18/03/2015.
 */

'use strict';

angular.module('dtBlockBooking')
    .controller('blockVerifyEmailController', ['$log', '$scope', '$location', 'blockSignupService', 'blockSignupData', 'loginService', 'stepState',
        function($log, $scope, $location, blockSignupService, blockSignupData, loginService, stepState){

            var email_query = ($location.search()).email || undefined,
                email_blockService = blockSignupService.getBlockSignupData().drivingSchool.contactEmail || undefined;

            $scope.currentModule = 'dtBlockBooking';
            $scope.token = ($location.search()).token;
            $scope.email = email_query || email_blockService;

            //$scope.loading = true;
            $scope.isEmailTokenVerified = false;
            $scope.emailVerificationInProgress = true;
            $scope.serviceError = false;

            $scope.validateEmailToken = function(){
                if(angular.isDefined($scope.token) && angular.isDefined($scope.email)) {
                    $log.debug('SUCCESS - angular.isDefined($scope.token) && angular.isDefined($scope.email)');

                    loginService.verifyEmailToken($scope.token, $scope.email, function(result){
                        $log.debug('SUCCESS - loginService.verifyEmailToken');

                        blockSignupData.drivingSchool.id = result.drivingSchoolId;
                        $scope.emailVerificationInProgress = false;
                        $scope.isEmailTokenVerified = true;

                    },function(errorResponse){
                        $log.debug('FAIL - loginService.verifyEmailToken', errorResponse);

                        $scope.emailVerificationInProgress = false;
                        $scope.serviceError = true;

                        switch(errorResponse.status){
                            case 410:
                                $scope.isExpired = true;
                                break;
                            case 412:
                                $scope.hasBeenVerified = true;
                                break;
                            case 400:
                            case 500:
                                stepState.changeState('school.error');
                                break;
                            default:
                                $scope.hasFailed = true;
                        }

                    });

                } else {
                    $log.debug('FAIL - angular.isDefined($scope.token) && angular.isDefined($scope.email)');

                    $scope.isEmailTokenVerified = true;
                    $scope.emailVerificationInProgress = false;
                }
            };
        }
    ]);
'use strict';

angular.module('dtBlockBooking')
    .controller('blockDashboardController', ['$scope', '$window',  '$log', 'locale', 'stepState',
        function($scope, $window, $log, locale, stepState) {
            $scope.submitDisabled = false;
            $scope.loading = false;

            $scope.addBooking = function() {
                stepState.changeState('school.addbooking');
            };

            $scope.printSummary = function(appointment) {
                angular.element(".appointment_wrapper").not(".appointment_wrapper#"+appointment.appointmentId).addClass("hidden-print");
                $window.print();
                angular.element(".appointment_wrapper").removeClass("hidden-print");

            };

        }
    ]);

/**
 * Created by Architech on 2014-10-20.
 */
'use strict';

angular.module('dtBlockBooking')
    .controller('finalizeBlockBookingController', ['$scope', '$log', 'stepState',
        function($scope, $log, stepState) {
            $scope.loadingUI = true;

            stepState.checkStepState('school.finalize');
        }
    ]);
/**
 * Created by at_user on 19/01/2015.
 */

'use strict';

angular.module('DriveTest.SchoolBusinessNumber', [])
    .directive('dtSchoolBusinessNumber', function(){
        return {
            restrict: 'AC',
            require: 'ngModel',
            link: function (scope, elem, attrs, ctrl) {

                ctrl.$validators.validateLicenceNumber = function(value) {

                    if (value === undefined || !value) {
                        return value;
                    }

                    var atMinimumLength = value.length >= 9;
                    ctrl.$setValidity('tooShort', atMinimumLength);

                    var atMaximumLength = value.length <= 9;
                    ctrl.$setValidity('tooLong', atMaximumLength);

                    var format = /^[0-9]{9}$/;
                    var matchesFormat = value.search(format) != -1;
                    ctrl.$setValidity('invalidChar', matchesFormat);

                    return value;
                };


            }
        };
    });
'use strict';

var dtBlockBookingApp = angular.module('dtBlockBooking');

dtBlockBookingApp.value('blockSignupData', {
    username: undefined,
    password: undefined,
    passwordConfirm: undefined,
    consentToBookTestsOnBehalfOfStudents: undefined,
    captchaResponse: undefined,

    drivingSchool: {
        id: undefined,
        schoolName: undefined,
        address: undefined,
        city: undefined,
        province: undefined,
        postalCode: undefined,
        billingAddressSameAsPrimary: undefined,
        businessNumber: undefined,
        contactPerson: undefined,
        contactEmail: undefined,
        contactPhoneNumber: undefined
    },

    booking: {
        locationName: undefined,
        locationId: undefined,
        serviceIds: undefined,
        blockGuid: undefined,
        fees: undefined,
        details: undefined
    }
});

dtBlockBookingApp.service('blockSignupService', ['blockSignupData', '$window', '$http', '$log', '$resource',
    function (blockSignupData, $window, $http, $log, $resource) {

        this.refreshDrivingSchool = function () {
            this.getSchool(
                function (result) {
                    blockSignupData.drivingSchool.id = result.drivingSchool.id;
                    blockSignupData.drivingSchool.schoolName = result.drivingSchool.schoolName;
                    blockSignupData.drivingSchool.businessNumber = result.drivingSchool.businessNumber;
                    blockSignupData.drivingSchool.address = result.drivingSchool.address;
                    blockSignupData.drivingSchool.province = result.drivingSchool.province.value;
                    blockSignupData.drivingSchool.postalCode = result.drivingSchool.postalCode;
                    blockSignupData.drivingSchool.city = result.drivingSchool.city;
                    blockSignupData.drivingSchool.contactPerson = result.drivingSchool.contactPerson;
                    blockSignupData.drivingSchool.contactEmail = result.drivingSchool.contactEmail;
                    blockSignupData.drivingSchool.contactPhoneNumber = result.drivingSchool.contactPhoneNumber;
                },
                function (result) {
                    log.error("Unable to load school information: " + result);
                    stepState.goToGeneralErrorState();
                });
        };

        this.getBlockSignupData = function () {

            if (!this.isSet()) {
                this.refreshDrivingSchool();
            }

            return blockSignupData;
        };

        this.clearBlockSignupData = function () {

            blockSignupData.username = undefined;
            blockSignupData.password = undefined;
            blockSignupData.passwordConfirm = undefined;
            blockSignupData.consentToBookTestsOnBehalfOfStudents = undefined;

            blockSignupData.drivingSchool.id = undefined;
            blockSignupData.drivingSchool.schoolName = undefined;
            blockSignupData.drivingSchool.businessNumber = undefined;
            blockSignupData.drivingSchool.address = undefined;
            blockSignupData.drivingSchool.province = undefined;
            blockSignupData.drivingSchool.postalCode = undefined;
            blockSignupData.drivingSchool.city = undefined;
            blockSignupData.drivingSchool.contactPerson = undefined;
            blockSignupData.drivingSchool.contactEmail = undefined;
            blockSignupData.drivingSchool.contactPhoneNumber = undefined;

            blockSignupData.booking.locationName = undefined;
            blockSignupData.booking.locationId = undefined;
            blockSignupData.booking.serviceIds = undefined;
            blockSignupData.booking.blockGuid = undefined;
        };

        this.setBlockSignupData = function (data) {
            blockSignupData.username = data.username;
            blockSignupData.password = data.password;
            blockSignupData.passwordConfirm = data.passwordConfirm;
            blockSignupData.consentToBookTestsOnBehalfOfStudents = data.consent;

            if (data.drivingSchool === undefined) {
                this.refreshDrivingSchool();
            }

            return blockSignupData;
        };

        this.register = function (successCallback, errorCallback) {
            return _schoolInformation.register(blockSignupData).$promise.then(successCallback, errorCallback);
        };

        this.register = function (successCallback, errorCallback) {
            return $http.post($window.dtServiceEndpoint + "/school", blockSignupData).then(successCallback, errorCallback);
        };

        this.removeToken = function (successCallback, errorCallback) {
            return $http.delete($window.dtServiceEndpoint + "/school/token").then(successCallback, errorCallback);
        };

        this.getSchool = function (success, errorHandler) {
            _schoolInformation.getSchool().$promise.then(success, errorHandler);
        };

        this.isSet = function () {
            $log.debug('checking is set ---> ', blockSignupData);
            //return blockSignupData.drivingSchool.schoolName !== undefined;
            return blockSignupData.drivingSchool.id !== undefined;
        };

        var _schoolInformation = $resource($window.dtServiceEndpoint + "/school/id", {}, {
                getSchool: {method: 'GET'}
            }
        );
    }]);

(function (angular) {
    'use strict';

    angular
        .module('dtBlockBooking')
        .service('blockSummaryService', ['$window', '$resource', 'blockSignupData', blockSummaryService]);

    function blockSummaryService($window, $resource, blockSignupData) {
        this.getPaymentSummary = function (blockFeesRequest, success, errorHandler) {
            return _bookingFees.check(blockFeesRequest).$promise.then(success, errorHandler)
        };

        var _bookingFees = $resource($window.dtServiceEndpoint + "/blockBooking/fees",
            {}, {
                check: {method: 'POST', params: {}}
            }
        );
    }
}(window.angular));


(function(angular) {
    'use strict';

    angular
        .module('dtBlockBooking')
        .service('blockTimeslotService', ['$window', '$log', '$resource', 'stepState','locale', blockTimeslotService]);

    function blockTimeslotService($window, $log, $resource, stepState, locale) {

        this.getTimeslots = function(data, successCallback, errorCallback) {
            return _timeslots.check(data).$promise.then(successCallback, errorCallback);
        };

        var _timeslots = $resource($window.dtServiceEndpoint + "/blockBooking/availableBlocks/",
            { }, {
                check: {method: 'POST', params:{ }  }
            }
        );

        this.holdBlock = function(localStartTime, blockItems, successCallback, errorCallback) {
            var data = {localStartTime: localStartTime,
                        blockItems: blockItems};
            return _holdblock.check(data).$promise.then(successCallback, errorCallback);
        };

        var _holdblock = $resource($window.dtServiceEndpoint + "/blockBooking/hold/",
            { }, {
                check: {method: 'POST', params:{ }  }
            }
        );

        this.getMomentTimezone = function(theTimezone) {
            var timeZone;
            //get timezone offset (daylight savings aware)
            if (theTimezone === 'EDT' || theTimezone === 'EST') {
                timeZone = "America/Toronto";
            } else if (theTimezone === 'CDT' || theTimezone === 'CST') {
                timeZone = "America/Winnipeg";
            } else {
                throw 'invalid timezone';
            }
            return timeZone;
        };

        this.getFormattedDate = function(timestamp, dtcMomentTimezone) {
        if (!locale.isFrench()) {
                return {timestamp: timestamp, display: moment.tz(timestamp, dtcMomentTimezone).format("dddd, MMMM Do YYYY, h:mm A")};
            } else {
                var frenchMoment = moment();
                frenchMoment.locale('fr-ca');
                frenchMoment = frenchMoment.tz(timestamp, dtcMomentTimezone);
                return {timestamp: timestamp, display: frenchMoment.format("dddd, MMMM Do YYYY, H[h]mm")};
            }
        };

        var getTimezoneOffset = function(theTimezone) {
            var timeZone;
            //get timezone offset (daylight savings aware)
            if (theTimezone === 'EDT' || theTimezone === 'EST') {
                timeZone = moment().tz("America/Toronto").format('Z');
            } else if (theTimezone === 'CDT' || theTimezone === 'CST') {
                timeZone = moment().tz("America/Winnipeg").format('Z');
            } else {
                throw 'invalid timezone';
            }
            return timeZone;
        };

        this.getTomorrowsDate = function(theTimezone) {
            var timeZone = getTimezoneOffset(theTimezone);

            var tomorrow = moment().utcOffset(timeZone).add(1, 'days');
            return tomorrow;
        };
    }
}(window.angular));


'use strict';

angular.module('dtBlockBooking')
    .service('loginService', ['$http', '$window',
        function($http, $window){

            this.login = function(school, successCallback, errorCallback){
                return $http.post($window.dtServiceEndpoint + "/school/token", school).then(successCallback, errorCallback);
            };

            this.verifyEmailToken = function (emailToken, email, successCallback, errorCallback) {
                var emailTokenData = {
                    email: email,
                    emailToken: emailToken
                };

                return $http.post($window.dtServiceEndpoint + "/school/emailToken", emailTokenData).then(successCallback, errorCallback);
            };

            this.recover = function (email, successCallback, errorCallback){

                return $http.post($window.dtServiceEndpoint + "/school/recover", email).then(successCallback, errorCallback);
            };

            this.reset = function (token, successCallback, errorCallback){

                return $http.post($window.dtServiceEndpoint + "/school/reset", token).then(successCallback, errorCallback);
            };

        }
    ]);
'use strict';

var dtBlockBookingApp = angular.module('dtBlockBooking');

dtBlockBookingApp.constant('States', {
        names: [
        'school.registration',
        'school.login',
        'school.forgotPassword',
        'school.dashboard',
        'school.addbooking',
        'school.location',
        'school.timeslot',
        'school.summary',
        'school.error',
        'school.quit'
    ], toTop: [
        true,
        true,
        true,
        true,
        true,
        true,
		true,
		true
    ]
});

dtBlockBookingApp.service('stepState', ['$rootScope','$state','$log', '$window', 'States', '$translate',
    function($rootScope, $state, $log, $window, States, $translate){

        var _currentStep = States.names.indexOf('school.registration'),
            _lastStep = States.names.indexOf('school.registration'),
            _this = this;

        this._currentInternalState = function() {
            return $state.current.name;
        };

        this.lastCompleteStepName = function() {
            return States.names[_lastStep];
        };

        this.currentStepName = function() {
            return States.names[_currentStep];
        };

        this.previousStepName = function() {
            return States.names[_currentStep>1?_currentStep-1:_currentStep];
        };

        this.goToPreviousState = function() {
            _lastStep = _currentStep;
            _currentStep = _currentStep - 1;
            $state.go(States.names[_currentStep]);
        };

        this.goToNextState = function() {
            _lastStep = _currentStep;
            _currentStep = _currentStep + 1;
            $state.go(States.names[_currentStep]);
        };

        this.goToGeneralErrorState = function () {
            $state.go('school.error');
        };

        this.checkStepState = function (state) {
            $log.log('checking state ---> ', $state);
            $log.log('checking _currentStep ---> ', _currentStep);
            this.updateHeader();

            var isInvalidSession = States.names.indexOf(state)>_currentStep;
            if (isInvalidSession) {
                $state.go(States.names[_currentStep]);
            }
            return isInvalidSession;
        };

        this.changeState = function (state,scroll) {
            if (scroll===undefined) {
                scroll = false;
            }
            _lastStep = _currentStep;
            _currentStep = States.names.indexOf(state);
            this.updateHeader();
            $state.go(state);

            $window.smoothScrollLink("#" + state.replace('.', '-'), scroll, States.toTop[_currentStep]);
        };

        this.updateHeader = function() {

            var stateArr = ['school.login','school.registration'],
                elQuitBtn = $('.quit-link'),
                elDashboardBtn = $('.dashboard-link');

            if(!$window._.contains(stateArr, _this._currentInternalState())){
                elQuitBtn.show();
            } else {
                elQuitBtn.hide();
            }

            if(!$window._.contains(stateArr, _this._currentInternalState()) && _this._currentInternalState() !== 'school.dashboard'){
                elDashboardBtn.show();
            } else {
                elDashboardBtn.hide();
            }

            $window.setNGBreadcrumb($translate.instant("ng.text." + this._currentInternalState()));
        };

}]);

(function(angular){
    'use strict';

    angular
        .module('dtBlockBooking')
        .service('studentBlock', [studentBlock]);


    function studentBlock(){
        var _studentBlock = {};

        this.getCurrent = function(studentId){
            var student = studentId || undefined;

            if(student !== undefined){
                return _studentBlock[student];
            } else {
                return _studentBlock;
            }
        };

        this.getStudent = function(licenceNumber) {
            var found = null;

            angular.forEach(_studentBlock, function (student) {

                if (student.licenceNumber === licenceNumber) {
                    found = student;
                }
            });

            return found;
        };

        this.addStudent = function(student, studentId){
            _studentBlock[studentId] = student;
        };

        this.removeStudent = function(studentId){
            return delete _studentBlock[studentId];
        };

        this.size = function(){
                var s = 0;
                var key;
                for (key in _studentBlock) {
                    if (_studentBlock.hasOwnProperty(key)) {
                        s++;
                    }
                }
                return s;
        };

        this.frenchRequested = function() {
            var key;
            for (key in _studentBlock) {
                if (_studentBlock.hasOwnProperty(key) && _studentBlock[key].testLanguage === "fr_ca") {
                    return true;
                }
            }
            return false;
        };

        this.getLicenceClasses = function() {
            var classes = {};
            var key;
            for (key in _studentBlock) {
                if (_studentBlock.hasOwnProperty(key)) {
                    classes[_studentBlock[key].licenceTypeValue] = _studentBlock[key].licenceTypeValue;
                }
            }
            return Object.keys(classes);
        };

        this.getCustomerIdsToLicenceClasses = function() {
            var classes = [];
            var key;
            var i = 0;
            for (key in _studentBlock) {
                if (_studentBlock.hasOwnProperty(key)) {
                    classes[i] ={customerId: _studentBlock[key].customerId, licenceClass:_studentBlock[key].licenceTypeValue};
                    i++;
                }
            }
            return classes;
        };

        this.getEligibleDates = function() {
            var key;
            var lastFrom = undefined;
            var firstTo = undefined;
            for (key in _studentBlock) {
                if ((_studentBlock.hasOwnProperty(key) && lastFrom === undefined) ||
                (_studentBlock[key].from > lastFrom)) {
                    lastFrom = _studentBlock[key].from;
                }
                if ((_studentBlock.hasOwnProperty(key) && firstTo === undefined) ||
                (_studentBlock[key].to < firstTo)) {
                    firstTo = _studentBlock[key].to;
                }
            }
            if (lastFrom > firstTo) {
                return undefined;
            }
            return {'from': moment(lastFrom, 'YYYY/MM/DD HH:mm Z').format('YYYY-MM-DD'), 'to': moment(firstTo, 'YYYY/MM/DD HH:mm Z').format('YYYY-MM-DD')};
        };

        this.hasIntersectingEligibleDates = function(theStudent) {
            var currDates = this.getEligibleDates();
            if (this.size() === 0) {
                return true;
            } else if (currDates === undefined) {
                return false;
            }

            var before = moment(theStudent.end, 'YYYY/MM/DD HH:mm Z').isBefore(moment(currDates.from, 'YYYY-MM-DD'));
            var after = moment(theStudent.from, 'YYYY/MM/DD HH:mm Z').isSame(moment(currDates.to, 'YYYY-MM-DD'));
            //student eligible range ends before intersection or student eligible range starts after intersection
            if (before || after) {
                return false;
            } else {
                return true;
            }
        }
    }
}(window.angular));


(function(angular) {
    'use strict';

    angular
        .module('dtBlockBooking')
        .service('studentListService', ['$log', '$resource', 'stepState', studentListService]);

    function studentListService($log, $resource, stepState) {

        var _studentEligibility = $resource('/app/api/blockBooking/eligibilityCheck',
            { }, {
                check: {method: 'POST', params:{  }  }
            }
        );

        //$http.post('/app/api/school/emailToken', emailTokenData).then(successCallback, errorCallback);

        this.checkStudentEligibility = function(params, successCallback, errorCallback) {

            /*{
             *      "licenceNumber": "A0871-04507-75515",
             *      "licenceExpiry": "2078/03/12",
             *      "licenceClass": "D",
             *      "endorsement": "Z",
             *      "checkSameLicense": true
             * }*/


            var data = angular.copy(params);
            data.checkSameLicense = true;
            data.endorsement = data.licenceType.endorsement;
            data.licenceClass = data.licenceType.class;
            delete data.licenceType;

            return _studentEligibility.check(data).$promise.then(successCallback, errorCallback);
        };
    }

}(window.angular));


/**
 * Created by Hisham on 2014-09-16.
 */
'use strict';

var dtBookingApp = angular.module('dtBooking');

dtBookingApp.controller('bookingCancel', ['$scope', '$log', '$state', function($scope, $log, $state) {

        $log.log('locationController bookingCancel $state ---> ', $state);

        $scope.parentWindow = window.opener.$windowScope;
        $scope.parentWindow.paymentCancelled();

        window.close();

    }
]);
/**
 * Created by Hisham on 2014-09-16.
 */
'use strict';

var dtBookingApp = angular.module('dtBooking');

dtBookingApp.controller('bookingSuccess', ['$scope', '$log', '$state', function($scope, $log, $state) {

        $log.log('locationController bookingSuccess $state ---> ', $state);

        $.urlParam = function(name){
            var results = new RegExp('[?&amp;]' + name + '=([^&amp;#]*)').exec(window.location.href);
            return results[1] || 0;
        };



        $scope.parentWindow = window.opener.$windowScope;
        $scope.parentWindow.paymentSuccessful(""+$.urlParam('confirmationNumber'), ""+$.urlParam('id'));

        window.close();

    }
]);
'use strict';

var dtBookingApp = angular.module('dtBooking');

dtBookingApp
    .constant('_', window._)
    .controller('calendarController', ['$scope', '$window', '$filter', '$log', '$state', 'stepState', 'userService', 'scheduleAvailability', 'bookingTimeService', 'appointmentService', 'errorService', 'locale',
    function($scope, $window, $filter, $log, $state, stepState, userService, scheduleAvailability, bookingTimeService, appointmentService, errorService, locale) {
        $scope.driverInfo = {};
        $log.log('calendarController $state ---> ', $state);

        var currentDate = new Date();
        var eligibilityRange;
        $scope.today = currentDate.getDay();
        $scope.currentDate = currentDate;

        stepState.checkStepState('booking.calendar');
        if (!userService.isSet()) {
            stepState.changeState('booking.registration');
            return;
        }

        $scope.loading = false;
        $scope.errMsg = errorService.clean($scope);

        $scope.$on('user:updated', function (event, user) {
            // rebuild calendar based on new location data in user
            if (user.locationId !== undefined && stepState.currentStepName() !== ' booking.calendar' && stepState.currentStepName() !== 'booking.timeslot') {
                setUpTheCalendarForUser(user);
            }
        });

        setUpTheCalendarForUser(userService.getCurrentUser());



        function setUpTheCalendarForUser(user) {
            $scope.driver = user;
            $scope.appointments = $scope.driver.appointments;
            eligibilityRange = getDisplayDateFromDriver($scope.driver);
            var tempDate = new Date(eligibilityRange.start);
            tempDate.setDate(1);
            if (user.selectedDate == undefined || ($scope.appointments.dirtyBit == undefined && $scope.appointments.selectedDateFlag  == undefined)) {
                //When appointment is not loaded or it is a new location(in which case  user.selectedDate == undefined)
                $scope.calculatedDate = (eligibilityRange.start >= currentDate) ? tempDate : new Date((currentDate).setDate(1));
            } else {
                //When appointment is loaded the selectedDate was of the local time zone, need to convert it the timezone of the test centre.
                var filterDate = "" + user.selectedDate;
                //Date format is "2017-12-08"
                var arr = filterDate.split("-");
                var year = arr[0];
                var month = arr[1];
                tempDate.setFullYear(year);
                tempDate.setMonth(month -1);
                tempDate.setDate(1);

                $scope.calculatedDate = tempDate;
            }
            getCalendarWithAvailability(eligibilityRange);
        }

        $scope.previousStep = function () {
            stepState.changeState('booking.location', false);
        };

        $scope.goToPreviousMonth = function (evt) {
            if (evt !== undefined) {
                evt.preventDefault();
            }

            if ($scope.previousMonthAvailable && !$scope.loading) {
                $scope.calculatedDate = advanceMonth($scope.calculatedDate, -1);
                if($scope.appointments.selectedDateFlag == undefined) {
                    //When user has not selected a date in the calendar, need to keep the appointment as the original state(not loaded)
                    $scope.appointments.dirtyBit = undefined;
                }
                getCalendarWithAvailability(eligibilityRange);
            }
        };

        $scope.goToNextMonth = function (evt) {
            if (evt !== undefined) {
                evt.preventDefault();
            }

            if ($scope.nextMonthAvailable && !$scope.loading) {
                $scope.calculatedDate = advanceMonth($scope.calculatedDate, +1);
                if($scope.appointments.selectedDateFlag == undefined) {
                    //When user has not selected a date in the calendar, need to keep the appointment as the original state(not loaded)
                    $scope.appointments.dirtyBit = undefined;
                }
                getCalendarWithAvailability(eligibilityRange);
            }
        };

        $scope.selectDate = function (evt, dateIn) {
            if (evt !== undefined) {
                evt.preventDefault();
            }

            $log.debug('dateIn ---> ', dateIn);
            $scope.userSelectedDate = dateIn;
            $scope.driver.selectedDate = $scope.formatDateToString(dateIn.year, dateIn.month, dateIn.day);

            $scope.appointments.selectedDateFlag = 1;

            if (dateIn.appointments.length>0) {
                $scope.sameDayWarning = true;
            } else {
                $scope.sameDayWarning = false;
            }

            userService.setSelectedDate($scope.driver.selectedDate);

        };

        $scope.driverSubmit = function () {
            $scope.errMsg = errorService.clean($scope);

            if (!errorService.isBlocked($scope) && $scope.driver.selectedDate !== undefined) {
                $scope.loading = true;
                disableBookingContainer();
                bookingTimeService.getTimeSlots($scope.driver.serviceId, $scope.driver.selectedDate, function (response) {
                    $scope.loadingData = false;
                    $scope.loading = false;
                    enableBookingContainer();
                    var user = userService.getCurrentUser();
                    user.timeslots = response.availableBookingTimes;

                    if (bookingTimeService.hasHeldAppointmentDate(user.selectedDate)){
                        bookingTimeService.addHeldAppointmentToCurrentUser(user);
                    }

                    userService.setUser(user);
                    stepState.changeState('booking.timeslot', true);

                }, errorService.errorHandler($scope));
            }
        };

        $scope.formatDateToString = function (y, m, d) {
            var month = "" + (m + 1);
            if (month.length === 1) {
                month = "0" + month;
            }
            var day = "" + d;
            if (day.length === 1) {
                day = "0" + day;
            }
            return  y + "-" + month + "-" + day;
        };

        function getCalendarWithAvailability(eligibilityRange) {
            $scope.loading = true;
            scheduleAvailability.getAvailability($scope.driver.serviceId, $scope.calculatedDate, function (response) {
                $scope.loading = false;

                $log.log('$promise getAvailability complete ---> ');
                if (response.availableBookingDates === undefined) {
                    errorService.showServerError($scope,response,'booking.app.timeslots.unavailable');
                    return;
                }

                $scope.displayedDate = $scope.calculatedDate;
                $scope.monthDisplayed = $scope.calculatedDate.getMonth();
                $scope.yearDisplayed = $scope.calculatedDate.getFullYear();

                $scope.dateMatrix = getDateMatrixForMonth($scope.monthDisplayed, $scope.yearDisplayed, response.availableBookingDates, eligibilityRange);
                $scope.previousMonthAvailable = isEligibleInMonth(new Date(new Date($scope.calculatedDate).setMonth($scope.calculatedDate.getMonth()-1)),eligibilityRange);
                $scope.nextMonthAvailable = isEligibleInMonth(new Date(new Date($scope.calculatedDate).setMonth($scope.calculatedDate.getMonth()+1)),eligibilityRange);

                if ($scope.driver.selectedDate !== undefined) {
                    var selectedDate;
                    //$scope.driver.selectedDate is at the local time the very first time loaded
                    var filterDate = "" + $scope.driver.selectedDate;
                    //Date format is "2017-12-08"
                    var arr = filterDate.split("-");
                    var year = arr[0];
                    var month = arr[1];
                    var day = arr[2];

                    selectedDate = new Date();
                    selectedDate.setFullYear(year);
                    //Javascript data month starts at 0
                    selectedDate.setMonth(month - 1);
                    selectedDate.setDate(day);
                    //driver has the timezone of the test centre
                    var timezone = $scope.driver.timezone;
                    //When customer has booked appointment, the selectedDate is set to the local timezone.
                    //Need to convert it to the date at the timezone of the test centre
                    var appt = $scope.appointments;
                    if ($scope.appointments != undefined && $scope.appointments.length > 0 && $scope.appointments.dirtyBit == undefined ) {
                            selectedDate = createDateByTimezone(selectedDate, "yyyy-MM-dd", timezone);
                            //set the dirtyBit of the appointments to indicate the appointment is loaded
                            $scope.appointments.dirtyBit = 1;
                    }
                   
                    for (var i = 0; i < $scope.dateMatrix.length; i++) {
                        for (var j = 0; j < $scope.dateMatrix[i].length; j++) {
                            if ($scope.dateMatrix[i][j].day === (selectedDate.getDate()) && $scope.dateMatrix[i][j].month === selectedDate.getMonth() && $scope.dateMatrix[i][j].year === selectedDate.getFullYear()) {
                                $scope.userSelectedDate = $scope.dateMatrix[i][j];
                            }
                        }
                    }
                }

                $log.debug("Setting calendar to display for " + $scope.displayedDate.toDateString());

            }, errorService.errorHandler($scope, function (response) {
                $scope.loading = false;
            }));
        }


        function getDateMatrixForMonth(month, year, availabilityInfo, eligibilityRange) {
            var currentDate = new Date();
            var dateMatrix = [];
            var startDate = new Date(year, month);
            var dateToSet = new Date(year, month);
            var daysToSkip = startDate.getDay();
            var daysUntilToday = $filter('date')(currentDate, 'yyyyMM') === $filter('date')(dateToSet, 'yyyyMM') ? currentDate.getDate() - 1 : 0;

            for (var i = 0; dateToSet.getMonth() === startDate.getMonth(); i++) {
                dateMatrix[i] = new Array(7);

                for (var j = 0; j < 7; j++) {
                    if ((i === 0 && j < daysToSkip) || dateToSet.getMonth() !== startDate.getMonth()) {
                        dateMatrix[i][j] = { day: -1, month: -1, year: -1, disabled: true, availability: 4 };
                    } else {
                        if (dateToSet.getDate() < availabilityInfo.length) {
                        }

                        var density = ((dateToSet.getDate()-1) < availabilityInfo.length) ? availabilityInfo[dateToSet.getDate()-1].density : 4;
                        var dateDisabled = (daysUntilToday-- > 0) || (density > 2) || (dateToSet > eligibilityRange.end) || (dateToSet < eligibilityRange.start);

                        dateMatrix[i][j] = {
                            day: dateToSet.getDate(),
                            month: dateToSet.getMonth(),
                            year: dateToSet.getFullYear(),
                            disabled: dateDisabled,
                            availability: density,
                            appointments: []
                        };

                        var fixedMonth;
                        var fixedDay;
                        var checkdate;

                        //checkdate and apptDate are matched to attach the appointments to the data matrix
                        //however, their values various depending on the local timezone and the booking flow(new or reschedule)
                        //the following logic are to make sure the appointments are attached to the right data the in the availability calendar(data matrix)
                        if(!angular.isUndefined($scope.appointments)) {

                            for (var a = 0; a < $scope.appointments.length; a ++) {
                                var timezone = $scope.driver.timezone;
                                var apptDate = $scope.appointments[a].date;

                                var apptDateStr = apptDate;
                                var flowName = stepState.getFlow();

                                if(flowName == "booking.new.appointment" ){
                                    if(isTimezoneDifferent(timezone)) {
                                        //When the timezone is changed in the new booking flow, the appointment date(apptDate) is of local timezone.
                                        //Need to convert it to the timezone of test center.

                                        fixedMonth = (dateMatrix[i][j].month * 1) + 1;
                                        fixedMonth = (fixedMonth < 10 ? "0" : "") + fixedMonth;
                                        fixedDay = (dateMatrix[i][j].day < 10 ? "0" : "") + dateMatrix[i][j].day;
                                        checkdate = "" + dateMatrix[i][j].year + fixedMonth + fixedDay;

                                        var y = apptDate.substr(0, 4);
                                        var m = ((apptDate.substr(4, 2) * 1) + 1 ) + "";
                                        var d = apptDate.substr(6, 2);

                                        var fixedm = (m < 10 ? "0" : "") + m;
                                        var fixedd = (d < 10 ? "0" : "") + d;

                                        apptDate = "" + y + fixedm + fixedd;

                                        apptDate = createDateByTimezone(apptDate, "yyyy-MM-dd", timezone);
                                        var year = apptDate.getFullYear();
                                        var month = (apptDate.getMonth() * 1) + 1;
                                        month = (month < 10 ? "0" : "") + month;
                                        var day = (apptDate.getDate() < 10 ? "0" : "") + apptDate.getDate();
                                        apptDateStr = "" + year + month + day;
                                    }else{
                                        fixedMonth = (dateMatrix[i][j].month < 10 ? "0" : "") + dateMatrix[i][j].month;
                                        fixedDay = (dateMatrix[i][j].day < 10 ? "0" : "") + dateMatrix[i][j].day;
                                        checkdate = "" + dateMatrix[i][j].year + fixedMonth + fixedDay;
                                    }
                                }else{
                                    //flowName == "booking.reschedule.appointment"
                                    //dateMatrix[i][j].month start at 0 but apptDate.month starts at 1, need to adjust.
                                    //in this reschedule flow the appDate is always of the timezone of test center, no need to convert
                                    fixedMonth = (dateMatrix[i][j].month * 1) + 1;
                                    fixedMonth = (fixedMonth < 10 ? "0" : "") + fixedMonth;
                                    fixedDay = (dateMatrix[i][j].day < 10 ? "0" : "") + dateMatrix[i][j].day;
                                    checkdate = "" + dateMatrix[i][j].year + fixedMonth + fixedDay;
                                }

                                if (apptDateStr === checkdate && dateMatrix[i][j].appointments.length < 3) {
                                    (dateMatrix[i][j]).appointments.push ({ id: $scope.appointments[a].appointmentId,  str: $scope.appointments[a].appString} );
                                }
                            }
                        }

                        if (currentDate.toDateString() === dateToSet.toDateString()) {
                            (dateMatrix[i][j]).today = true;
                        }

                        dateToSet.setDate(dateToSet.getDate() + 1);
                    }
                }
            }
            return dateMatrix;
        }

        function isTimezoneDifferent(timezonestr){
            var datestr = Date().toString();
            var str = datestr.match(/\(([A-Za-z\s].*)\)/)[1];
            var arr = str.split(" ");
            var acronym = '';
            for(var a = 0; a < arr.length; a++){
                acronym += arr[a].charAt(0);
            }
            return timezonestr != acronym;
        }

        function createDateByTimezone(datein, formatstr, timezonestr){
            if(datein.length < 9){
                //datein is of formate 20180101
                datein = datein.substr(0, 4) + "-" + datein.substr(4, 2) + "-" + datein.substr(6, 2);
            }
            var filterDate = "" + $filter('date')(datein, formatstr, timezonestr);
            //Date format is "2017-12-08"
            var arr = filterDate.split("-");
            var year = arr[0];
            var month = arr[1];
            var day = arr[2];
            var apptDate = new Date();
            apptDate.setFullYear(year);
            apptDate.setMonth(month -1);
            apptDate.setDate(day);
            return apptDate;
        }

        function getDisplayDateFromDriver(driverObj) {
            return { start: new Date(Date.parse(driverObj.eligibilityRange.from)), end: new Date(Date.parse(driverObj.eligibilityRange.to)) };
        }

        function isEligibleInMonth(date, eligibilityRange) {

            var startDate = $filter('date')(eligibilityRange.start, 'yyyyMM');
            var endDate = $filter('date')(eligibilityRange.end, 'yyyyMM');
            var testDate = $filter('date')(date, 'yyyyMM');
            var currentDate = $filter('date')(new Date(), 'yyyyMM');

            $log.debug("checking " + testDate + " in " + startDate + " to " + endDate + "(" + $filter('date')(date, 'yyyy-MM-dd') + ") = " + (testDate >= startDate && testDate <= endDate));

            return (testDate >= startDate && testDate <= endDate && currentDate <= testDate);
        }

        function advanceMonth(dateIn, monthOffset) {
            $scope.calculatedDate = new Date(new Date(dateIn).setMonth(dateIn.getMonth() + monthOffset));
            return $scope.calculatedDate;
        }
    }
]);

'use strict';

var dtBookingApp = angular.module('dtBooking');

dtBookingApp.controller('completeController', ['$scope', '$window', '$log', '$state', 'stepState', 'userService', 'bookingTimeService', 'finalizeBooking',
    function($scope, $window, $log, $state, stepState, userService, bookingTimeService, finalizeBooking) {
        $scope.driverInfo = {};
        $log.log('$state ---> ', $state);

        stepState.checkStepState('booking.complete');

        $scope.driver = userService.getCurrentUser();

        $scope.timeForDisplay = bookingTimeService.getFormattedTime($scope.driver.timeslot, $scope.driver.momentTimezone);

        $window.smoothScrollLink("#" + $state.current.name.replace('.', '-'));

        $scope.$watch(function() {
            return $scope.driver;
        }, function(newValue, oldValue) {
            if (newValue) {
                $scope.imgSrc = $scope.driver.barcode;
            }
        });
    }
]);
'use strict';

angular.module('dtBooking')
    .controller('dashboardController', ['$scope', '$window',  '$log', 'locale', 'appointmentService', 'bookingTimeService', '$modal', 'userService', 'stepState', 'errorService', 'registrationSvc','$filter',
        function($scope, $window, $log, locale, appointmentService, bookingTimeService, $modal, userService, stepState, errorService, registrationSvc, $filter) {

            $scope.driver = userService.getCurrentUser();
            stepState.updateHeader();

            function clearSession() {
                document.cookie = 'token' +'=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
                userService.setUser({});
            }

            if ($scope.driver.lastName === undefined && $scope.driver.email === undefined) {
                //$(".quit-link").hide();
                registrationSvc.removeToken($scope.driver.licenceNumber, clearSession, clearSession);
                stepState.changeState('booking.login');
            }

            if ($scope.driver.licenceNumber === undefined) {
                stepState.changeState('booking.login');
            }

            $scope.submitDisabled = false;
            $scope.loading = false;
            $scope.appointmentDeleted = false;
            if($scope.driver.reschedule) {
                $scope.driver.reschedule = false;
                $scope.rescheduleUser = {};
            }

            var getPending = function() {
                $scope.appointments = [];
                $scope.loading = true;
                appointmentService.pendingAppointments(
                    function(result) {
                        $scope.loading = false;
                        $log.log("pendingAppointments successful: " + result);
                        $scope.submitDisabled = false;
                        $scope.appointmentDeleted = false;

                        var thisdate;
                        if (angular.isDefined(result.pendingAppointments)) {
                            for (var a = 0; a < result.pendingAppointments.length; a++) {
                                thisdate = new Date(result.pendingAppointments[a].time);
                                //thisdate is of the local timezone, need to convert it to be of the timezone of the test centre.
                                var formatstr = "yyyy-MM-dd";
                                var timezonestr = result.pendingAppointments[a].centre.timezone;
                                var filterDate = "" + $filter('date')( thisdate, formatstr, timezonestr);
                                //Date format is "2017-12-08"
                                var arr = filterDate.split("-");
                                var year = arr[0];
                                var month = arr[1];
                                var day = arr[2];
                                var apptDate = new Date();
                                apptDate.setYear(year);
                                apptDate.setMonth(month -1);
                                apptDate.setDate(day);
                                thisdate = apptDate;

                                var min = thisdate.getMinutes() > 0 ? ":" + thisdate.getMinutes() : "";
                                var fixedMonth = (thisdate.getMonth() < 10 ? "0" : "") + (thisdate.getMonth() + 1);
                                var fixedDay = (thisdate.getDate() < 10 ? "0" : "") + thisdate.getDate();

                                result.pendingAppointments[a].date = "" + thisdate.getFullYear() + fixedMonth + fixedDay;

                                var momentTimezone = bookingTimeService.getMomentTimezone(result.pendingAppointments[a].centre.timezone);
                                result.pendingAppointments[a].timeForDisplay = bookingTimeService.getFormattedTime(result.pendingAppointments[a].time, momentTimezone);

                                result.pendingAppointments[a].appString = result.pendingAppointments[a].displayLicenceClass + " <span>@ " + result.pendingAppointments[a].timeForDisplay.display + "</span>";

                                result.pendingAppointments[a].timezone =  result.pendingAppointments[a].centre.timezone;
                            }

                            $scope.appointments = result.pendingAppointments;

                            $log.debug('dt-dashboardController -> driver.timezone = ' +  $scope.driver.timezone);
                            var user = userService.getCurrentUser();
                            user.appointments = result.pendingAppointments;
                            userService.setUser(user);
                        }


                    }, errorService.errorHandler($scope, function(response){
                        $log.debug("pendingAppointments failed: " + response);

                        $scope.loading = false;
                        $scope.submitDisabled = false;
                        $scope.appointmentDeleted = false;
                        var messageKey = "booking.error";

                        if(response.status === 400 || response.status === 500 ) {
                            messageKey = "booking.error.retrieving.appointments";
                        }

                        $scope.errMsg = {hasError:true};
                        $scope.errMsg.msg = errorService.getMessage(messageKey);

                        return true;
                    })
                );
            };

            getPending();

            $scope.startBookingFlow = function() {
                stepState.setFlow("booking.new.appointment");
                userService.clearCurrentUserAppointment();
                stepState.changeState("booking.licence");
            };

            $scope.quit = function() {
                stepState.changeState('booking.quit');
            };

            $scope.confirmDelete = function(appointment) {
                $modal.open({
                    windowClass: "confirmDeleteModal",
                    templateUrl: '/ng/dt-appointments-deletemodal',
                    controller: 'confirmDeleteController',
                    resolve: {
                        appointment: function() {
                            return appointment;
                        },
                        driver: function() {
                            return $scope.driver;
                        }
                    }
                }).result.then(function() {
                    $scope.appointmentDeleted = true;
                    getPending();
                }, function() {
                    $scope.appointmentDeleted = false;
                });
            };

            $scope.reschedule = function(appointment) {
                stepState.setFlow("booking.reschedule.appointment");
                $modal.open({
                    windowClass: "rescheduleModal",
                    templateUrl: '/ng/dt-appointments-reschedulemodal',
                    controller: 'RescheduleController',
                    resolve: {
                        appointment: function() {
                            return appointment;
                        },
                        driver: function() {
                            return $scope.driver;
                        }
                    }
                }).result.then(function() {
                    //$scope.appointmentDeleted = true;
                    $scope.appointments = [];
                }, function() {
                    //$scope.appointmentDeleted = false;
                    $log.debug('reschedule closed ---> ');
                });
            };

            $scope.printSummary = function(appointment) {

                angular.element(".appointment_wrapper").not(".appointment_wrapper#"+appointment.appointmentId).addClass("hidden-print");
                $window.print();
                angular.element(".appointment_wrapper").removeClass("hidden-print");

            };

        }
    ])
    .controller('confirmDeleteController', ["$scope",  "$log", "$modalInstance", "appointment", "appointmentService", "errorService", "driver",
        function($scope,  $log, $modalInstance, appointment, appointmentService, errorService, driver) {

            $scope.appointment = appointment;

            $scope.loading = true;

            $scope.cancelError = false;

            if(appointment.timezone == undefined && driver.timezone != undefined){
                appointment.timezone = driver.timezone;
            }
            $log.debug('dashnoardController,confirmDeleteController appointment.timezone =' + appointment.timezone);

            $scope.submitDisabled = false;
            appointmentService.with48hours(appointment.appointmentId, function(response) {
                $scope.within48hours = response.within48Hours;
                $scope.loading = false;
            }, errorService.errorHandler($scope, function(response) {
                $scope.loading = false;
            }));

            $scope.cancelTest = function() {
                $scope.loading = true;
                $scope.submitDisabled = true;

                appointmentService.cancelAppointment(appointment.appointmentId, function(result) {
                    if (result.success && (result.status !== 202 && result.status === 200)) {
                        $modalInstance.close();
                    } else {
                        $scope.cancelError = false;
                        $scope.loading = false;

                        errorService.showServerError($scope, result, 'booking.error.cancel.48.hours');
                    }
                }, errorService.errorHandler($scope, function (response) {

                    $scope.cancelError = true;
                    $scope.loading = false;
                    var messageKey, returnVal = false;

                    switch(response.status){
                        case 403:
                            messageKey = 'booking.app.message.appointment.3.6.check';
                            returnVal = true;
                            break;
                        case 409:
                            messageKey = 'booking.app.message.timeslot.conflict';
                            returnVal = true;
                            break;
                        case 420:
                            messageKey = 'booking.error.cancel';
                            returnVal = true;
                            break;
                        case 503:
                            messageKey = 'booking.error.cancel.48.hours';
                            returnVal = true;
                            break;
                    }

                    $scope.errMsg = {hasError:true};
                    $scope.errMsg.msg = errorService.getMessage(messageKey);
                    return returnVal;
                }));
            };

            $scope.closeDialog = function() {
                if ($scope.cancelError) {
                    $modalInstance.dismiss('cancel');
                } else {
                    $modalInstance.close();
                }
            };

        }
    ])
    .controller('RescheduleController', ["$scope", "$modalInstance", "appointment", "appointmentService", "bookingTimeService", "errorService", "userService", "driver", "$log", "stepState", "$filter",
        function($scope, $modalInstance, appointment, appointmentService, bookingTimeService, errorService, userService, driver, $log, stepState, $filter) {

            $scope.appointment = appointment;
            $scope.loading = true;
            $scope.submitDisabled = false;
            $scope.rescheduleUser = userService.getCurrentUser();

            if(appointment.timezone == undefined && driver.timezone != undefined){
                appointment.timezone = driver.timezone;
            }
            $log.debug('dashnoardController,RescheduleController appointment.timezone =' + appointment.timezone);

            userService.getRescheduleEligibility($scope.appointment, function(response){
                $log.debug('USERSERVICE ----- getRescheduleEligibility -------> ', response);

                $scope.loading = false;

                if (response.cancelEligible && response.licenceEligible) {

                    $scope.within48hours = response.within48HoursEligible;

                    $scope.rescheduleUser.eligible = response.licenceEligible;
                    $scope.rescheduleUser.eligibilityRange = {
                        from: response.from,
                        to: response.to
                    };
                    $scope.rescheduleUser.emailToken = driver.customerId;
                    $scope.rescheduleUser.token = driver.customerId;
                    //var licence = _parseLicence(appointment.licenseClass);
                    $scope.rescheduleUser.licenceClass = appointment.licenceClass;
                    $scope.rescheduleUser.endorsement = appointment.endorsement;
                    $scope.rescheduleUser.component = appointment.motorcycleType;
                    $scope.rescheduleUser.componentDisplay = appointment.motorcycleTypeDisplay !== undefined ? '(' + appointment.motorcycleTypeDisplay + ')' : '';

                    $scope.rescheduleUser.locationId = appointment.centre.id;
                    $scope.rescheduleUser.locationName = appointment.centre.city;
                    $scope.rescheduleUser.serviceId = appointment.serviceId;
                    $scope.rescheduleUser.dtcAddress1 = appointment.centre.address1;
                    $scope.rescheduleUser.dtcAddress2 = appointment.centre.address2;

                    $scope.rescheduleUser.timeslot = appointment.time;
                    $scope.rescheduleUser.selectedDate = $filter('date')(appointment.time, 'yyyy-MM-dd');

                    $scope.rescheduleUser.reschedule = true;

                    $scope.rescheduleUser.previousHold = {};
                    $scope.rescheduleUser.previousHold.location = appointment.centre.city;
                    var momentTimezone = bookingTimeService.getMomentTimezone(appointment.centre.timezone);
                    $scope.rescheduleUser.previousHold.selectedDate = bookingTimeService.getFormattedTime(appointment.time, momentTimezone);


                    $scope.rescheduleUser.previousHold.confirmed = {};
                    $scope.rescheduleUser.previousHold.confirmed.serviceId = appointment.serviceId;
                    $scope.rescheduleUser.previousHold.confirmed.guid = appointment.appointmentId;
                    $scope.rescheduleUser.previousHold.confirmed.timeslot = appointment.time;

                    // clear held appointment after a booking confirmation
                    $scope.rescheduleUser.heldAppointment = undefined;
                    $scope.rescheduleUser.bookingInProgress = true;

                    userService.setUser($scope.rescheduleUser);

                } else {
                    $scope.submitDisabled = true;
                    $scope.cancelEligible = response.cancelEligible;
                    $scope.within48hours = response.within48HoursEligible;

                    errorService.showServerError($scope, response, 'booking.error');
                }
            }, errorService.errorHandler($scope, function(response) {

                $scope.loading = false;
                $scope.submitDisabled = true;
                var messageKey, returnVal = false;

                switch(response.status){
                    case 403:
                        messageKey = 'booking.app.message.eligibility.ineligible.3.6.check';
                        returnVal = true;
                        break;
                    case 405:
                        messageKey = 'booking.app.message.eligibility.ineligible.road';
                        returnVal = true;
                        break;
                    case 412:
                        messageKey = 'booking.app.message.eligibility.ineligible.licence';
                        returnVal = true;
                        break;
                    case 500:
                        messageKey = 'booking.error';
                        returnVal = true;
                        break;
                }

                $scope.errMsg = {hasError:true};
                $scope.errMsg.msg = errorService.getMessage(messageKey);
                return returnVal;

            }));

            $scope.rescheduleTest = function(evt) {
                evt.preventDefault();

                $scope.loading = true;
                $scope.submitDisabled = true;

                $modalInstance.close();
                stepState.changeState('booking.location');
            };

            //var _parseLicence = function(licence) {
            //    var suffix = licence.slice(-1);
            //    if (suffix === "Z" && licence.length > 1) {
            //        return {
            //            licenceClass: licence.slice(0, -1),
            //            endorsement: suffix
            //        };
            //    } else if (suffix === "F" || suffix === "L" || suffix === "P") {
            //        return {
            //            licenceClass: licence.slice(0, -1),
            //            component: suffix
            //        };
            //    } else {
            //        return {
            //            licenceClass: licence
            //        };
            //    }
            //};

            $scope.closeDialog = function() {
                $modalInstance.dismiss('cancel');
            };

        }
    ]);

'use strict';

var dtBookingApp = angular.module('dtBooking');

dtBookingApp.controller('errorController', ['$scope','$log','$state', 'userService', 'stepState', '$window', function($scope,$log,$state,userService, stepState, $window) {
    $log.log('errorController $state ---> ', $state);

    $window.smoothScrollLink("#"+$state.current.name.replace('.','-'));

    $scope.retry = function() {
        stepState.changeState(stepState.lastCompleteStepName());
    };
}]);

'use strict';

angular.module('dtBooking')
    .controller('finalizeBookingController', ['$scope', '$log', 'stepState', 'userService', '$timeout', 'finalizeBooking', 'errorService', 'verifyPayment',
        function($scope, $log, stepState, userService, $timeout, finalizeBooking, errorService, verifyPayment) {
            $scope.driver = userService.getCurrentUser();
            $scope.loadingUI = true;

            stepState.checkStepState('booking.finalize');

            if (!userService.isSet()
                && userService.isEligible()) {

                stepState.changeState('booking.login');
                return;
            }

            var _onBookingSuccessful = function(response){
                if (response.barcode) {
                    var user = userService.getCurrentUser();
                    user.barcode = response.barcode;
                    user.completionID = response.displayId;
                    user.paymentRequiredOnDayOfTest = response.paymentRequiredOnDayOfTest;
                    user.payment = {};
                    user.payment.confirmationNumber = response.settlementConfirmationNumber === undefined ||
                        response.settlementConfirmationNumber === null ? '' : response.settlementConfirmationNumber;
                    user.bookingInProgress = false;

                    userService.setUser(user);

                    $timeout(function() {
                        if (stepState.currentStepName() === 'booking.finalize') {
                            stepState.changeState('booking.complete');
                        }
                    }, 5000);
                }
            };

            var _onBookingFailure = function(response){
                $scope.loadingUI = false;
                var messageKey, returnVal = false;

                switch(response.status) {
                    case 402:
                        messageKey = 'booking.app.message.incomplete.payment';
                        returnVal = true;
                        break;
                    case 403:
                        messageKey = 'booking.app.message.eligibility.ineligible.3.6.check';
                        returnVal = true;
                        break;
                    case 412:
                        messageKey = 'booking.app.message.complete.booking.ineligible.licence';
                        returnVal = true;
                        $scope.sameLicenceError = true;
                        break;
                    case 429:
                        messageKey = 'booking.app.message.complete.booking.multiple.attempt';
                        returnVal = true;
                        break;
                    case 500:
                        messageKey = 'booking.error.incomplete.booking';
                        returnVal = true;
                        break;
                }

                $scope.errMsg = {hasError:true};
                $scope.errMsg.msg = errorService.getMessage(messageKey);
                $scope.errMsg.errorCode = response.status;
                return returnVal;
            };

            var performFinalizeBooking = function(){
                if ($scope.driver.bookingInProgress) {
                    var payment = $scope.driver.payment;
                    if ($scope.driver.reschedule) {
                        finalizeBooking.completeReschedule(
                            $scope.driver.previousHold.confirmed.guid,
                            $scope.driver.heldAppointment.guid,
                            payment ? payment.confirmationNumber : "",
                            payment && payment.id ? $scope.driver.payment.id : "" + Date.now(),
                            $scope.driver.licenceClass,
                            $scope.driver.endorsement,
                            _onBookingSuccessful,
                            errorService.errorHandler($scope, _onBookingFailure));
                    } else {
                        finalizeBooking.completeBooking($scope.driver.heldAppointment.guid,
                            payment ? payment.confirmationNumber : "",
                            payment && payment.id ? $scope.driver.payment.id : "" + Date.now(),
                            $scope.driver.licenceClass,
                            $scope.driver.endorsement,
                            _onBookingSuccessful,
                            errorService.errorHandler($scope, _onBookingFailure));
                    }

                }
            };

            performFinalizeBooking();

        }
    ]);
'use strict';

var dtBookingApp = angular.module('dtBooking');

dtBookingApp.controller('licenceController', ['$scope', '$log', '$state', '$modal', 'userService', 'bookingTimeService', 'stepState', '$translate', 'licenceClassService', 'appointmentService', 'locale', 'errorService', '$filter', 'registrationSvc',
    function ($scope, $log, $state, $modal, userService, bookingTimeService, stepState, $translate, licenceClassService, appointmentService, locale, errorService, $filter, registrationSvc) {
        $scope.driverInfo = {};
        $scope.classesList = undefined;
        $scope.profileIneligible = false;

        $log.log('licenceController $state ---> ', $state);

        $scope.master = {};

        function clearSession() {
            document.cookie = 'token' +'=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
            userService.setUser({});
        }

        stepState.checkStepState("booking.licence");
        if (!userService.isSet()) {
            registrationSvc.removeToken(userService.getCurrentUser().licenceNumber, clearSession, clearSession);
            stepState.changeState('booking.registration');
            return;
        }
        $scope.driver = userService.getCurrentUser();

        $scope.sameLicenceError = false;
        $scope.loading = false;
        enableBookingContainer();
        $scope.loadingData = true;
        $scope.errMsg = errorService.clean($scope);
        if ($scope.driver.reschedule) {
            $scope.driver.reschedule = false;
            $scope.rescheduleUser.previousHold = {};
            $scope.rescheduleUser.previousHold.confirmed = {};
        }

        appointmentService.pendingAppointments(
            function (result) {
                $log.log("pendingAppointments successful: " + result);

                // add date field for better indexing
                var thisdate;
                if (angular.isDefined(result.pendingAppointments)) {
                    for (var a = 0; a < result.pendingAppointments.length; a++) {
                        thisdate = new Date(result.pendingAppointments[a].time);
                        var min = thisdate.getMinutes() > 0 ? ":" + thisdate.getMinutes() : "";
                        var fixedMonth = (thisdate.getMonth() < 10 ? "0" : "") + thisdate.getMonth();
                        var fixedDay = (thisdate.getDate() < 10 ? "0" : "") + thisdate.getDate();

                        result.pendingAppointments[a].date = "" + thisdate.getFullYear() + fixedMonth + fixedDay;

                        var momentTimezone = bookingTimeService.getMomentTimezone(result.pendingAppointments[a].centre.timezone);
                        result.pendingAppointments[a].timeForDisplay = bookingTimeService.getFormattedTime(result.pendingAppointments[a].time, momentTimezone);

                        result.pendingAppointments[a].appString = result.pendingAppointments[a].displayLicenceClass + " <span>@ " + result.pendingAppointments[a].timeForDisplay.display + "</span>";
                    }

                    $scope.appointments = result.pendingAppointments;

                    var user = userService.getCurrentUser();
                    user.appointments = result.pendingAppointments;

                    userService.setUser(user);
                }
            },
            function (reason) {
                $log.log("pendingAppointments failed: " + reason);
                $scope.submitDisabled = false;
                $scope.appointmentDeleted = false;
            },
            function (result) {

                function stopLoading() {
                    $scope.loading = false;
                    enableBookingContainer();
                    $scope.loadingData = false;
                }

                if ($scope.driver.licenceClasses !== undefined) {
                    $scope.classesList = $scope.driver.licenceClasses;
                    stopLoading();

                } else {
                    licenceClassService.loadLicenceClasses(function (response) {
                        stopLoading();
                        if (response.primaryDriverLicenceClasses.length > 0 || response.secondaryDriverLicenceClasses.length > 0) {
                            $scope.classesList = {
                                primaryDriverLicenceClasses: response.primaryDriverLicenceClasses,
                                secondaryDriverLicenceClasses: response.secondaryDriverLicenceClasses
                            };

                            $scope.driver.licenceClasses = $scope.classesList;

                        } else {
                            errorService.showServerError($scope, response);
                        }

                    }, errorService.errorHandler($scope, function (response) {
                        $scope.loading = false;
                        enableBookingContainer();
                        $scope.loadingData = false;

                        var messageKey, returnVal = false;

                        switch (response.status) {
                            case 417:
                                $scope.profileIneligible = true;
                                messageKey = 'booking.app.message.ineligible.status.flags';
                                returnVal = true;
                                break;
                        }

                        $scope.errMsg = {hasError: true};
                        $scope.errMsg.msg = errorService.getMessage(messageKey);

                        return returnVal;
                    }));
                }
            }
        );

        $scope.isLicenceClassEndorsable = function (licenceClass, endorsement) {
            if (licenceClass === undefined || endorsement === undefined) {
                return false;
            }

            var secondaryLicences = $scope.classesList !== undefined ? $scope.classesList.secondaryDriverLicenceClasses : undefined;
            if (secondaryLicences !== undefined) {
                for (var i = 0; i < secondaryLicences.length; i++) {
                    if (licenceClass === secondaryLicences[i].licenceClass) {
                        return secondaryLicences[i].endorsementAvailable;
                    }
                }
            } else {
                return false;
            }
        };

        $scope.onLicenceChange = function (thisLicence, licence) {
            $scope.errMsg = errorService.clean($scope);
            $scope.sameLicenceError = false;
        };

        $scope.isLicenceClassWithComponents = function (licenceClass, component) {
            if (licenceClass === undefined || component === undefined) {
                return false;
            }

            var primaryDriverLicenceClasses = $scope.classesList.primaryDriverLicenceClasses;
            if (primaryDriverLicenceClasses !== undefined) {
                var componentList = $filter('filter')(primaryDriverLicenceClasses, {licenceClass: licenceClass}, true);
                return (angular.isDefined(componentList) && componentList.length > 0);
            } else {
                return false;
            }
        };

        $scope.driverSubmit = function () {
            $log.debug('dt-licenceController.js $scope.driver ---> ', $scope.driver);
            $scope.driver.endorsement = $scope.isLicenceClassEndorsable($scope.driver.licenceClass, $scope.driver.endorsement) ?
                $scope.driver.endorsement : undefined;

            $scope.errMsg = errorService.clean($scope);

            var user = userService.getCurrentUser();

            $scope.driver.component = $scope.isLicenceClassWithComponents($scope.driver.licenceClass, $scope.driver.component)
                ? $scope.driver.component
                : undefined;

            var licenceClass = $filter('filter')($scope.classesList.primaryDriverLicenceClasses, {licenceClass: $scope.driver.licenceClass});

            var typeDisplay = '';
            //need to handle the error for licence 'G'
            if($scope.driver.licenceClass != 'G' && $scope.driver.licenceClass != 'G2'){
                typeDisplay = $scope.driver.component !== undefined ? licenceClass[0].componentDisplay[$scope.driver.component] : '';
            }

            $scope.driver.componentDisplay = typeDisplay !== undefined && typeDisplay !== '' ? '(' + typeDisplay + ')' : '';

            if($scope.driver.licenceClass != 'G' && $scope.driver.licenceClass != 'G2'){
                user.component = $scope.driver.component;
            }else{
                user.component = undefined;
            }
       
            user.checkSameLicense = true;
            user.bookingInProgress = true;
            user.locationName = undefined;
            user.locationId = undefined;
            user.serviceId = undefined;

            userService.setUser(user);

            $log.debug("Appointments = " + $scope.appointments);

            if (angular.isDefined($scope.appointments) && $scope.appointments.length > 0) {

                //do check
                var matchingLicence = $filter('filter')($scope.appointments, {licenseClass: $scope.driver.licenceClass}, true);
                var matchingEndorsement = $filter('filter')($scope.appointments, {licenseClass: $scope.driver.licenceClass + "Z"}, true);

                if ((matchingLicence !== undefined && matchingLicence.length > 0) ||
                    (matchingEndorsement !== undefined && matchingEndorsement.length > 0)) {
                    var errorMessage = $translate.instant("booking.errorcode.licence.blocked1") +
                        $scope.driver.licenceClass +
                        $translate.instant("booking.errorcode.licence.blocked2");

                    $scope.errMsg = {hasError: true, blocking: false, msg: errorMessage};
                    $scope.sameLicenceError = true;
                    $scope.driver.licenceClass = undefined;
                    return;
                }
            }

            // block continue
            if (!errorService.isBlocked($scope) && $scope.driver.licenceClass !== undefined) {
                $scope.loading = true;
                disableBookingContainer();

                userService.getEligibility(function (response) {

                    $scope.driver.eligible = response.eligible;
                    $scope.driver.eligibilityRange = {from: response.from, to: response.to};
                    $scope.sameLicenceError = false;

                    if (response.eligible) {

                        var user = userService.getCurrentUser();

                        user.eligible = response.eligible;
                        user.eligibilityRange = {from: response.from, to: response.to};

                        userService.setUser(user);

                        if (response.warnings.length === 0) {
                            $scope.loading = false;
                            enableBookingContainer();
                            stepState.goToNextState();
                        } else {
                            $modal.open({
                                windowClass: "eligibilityWarningModal",
                                templateUrl: '/ng/dt-eligibility-warning-modal?code=' + response.warnings.slice(0, 1),
                                controller: 'eligibilityWarningModalController'
                            }).result.then(function() {
                                $scope.loading = false;
                                enableBookingContainer();
                                stepState.goToNextState();
                            }, function() {
                                $scope.loading = false;
                                enableBookingContainer();
                            });
                        }
                    } else {
                        errorService.showServerError($scope, response, 'booking.error');
                        $scope.loading = false;
                        enableBookingContainer();
                    }

                }, errorService.errorHandler($scope, function (response) {

                    $scope.loading = false;
                    enableBookingContainer();
                    var messageKey, returnVal = false;

                    switch (response.status) {
                        case 403:
                            messageKey = 'booking.app.message.eligibility.ineligible.3.6.check';
                            returnVal = true;
                            break;
                        case 405:
                            messageKey = 'booking.app.message.eligibility.ineligible.road';
                            returnVal = true;
                            break;
                        case 412:
                            messageKey = 'booking.app.message.eligibility.ineligible.licence';
                            returnVal = true;
                            $scope.sameLicenceError = true;
                            break;
                        case 500:
                            messageKey = 'booking.error';
                            returnVal = true;
                            break;
                    }

                    $scope.errMsg = {hasError: true};
                    $scope.errMsg.msg = errorService.getMessage(messageKey);

                    return returnVal;

                }));
            }
        };

    }])
    .controller('eligibilityWarningModalController', ["$scope", "$modalInstance",
        function($scope, $modalInstance) {

            $scope.continueBooking = function() {
                $modalInstance.close();
            };

            $scope.close = function() {
                $modalInstance.dismiss('eligibilityWarningModal');
            };

        }
    ]);

'use strict';

var dtBookingApp = angular.module('dtBooking');

dtBookingApp.controller('locationController', ['$scope', '$rootScope', '$window', '$log', '$state', 'userService', 'bookingTimeService', 'stepState', 'dtcLocationData', 'errorService', 'leafletData', 'leafletBoundsHelpers','dtCacheServices', 'dtCacheServicesKeys',
    function ($scope, $rootScope, $window, $log, $state, userService, bookingTimeService, stepState, dtcLocationData, errorService, leafletData, leafletBoundsHelpers, dtCacheServices, dtCacheServicesKeys) {

        var southWestLat = 41.631538128024786;
        var southWestLng = -97.19560503959656;
        var northEastLat = 57.2;
        var northEastLng = -73.16411197185516;
        var maxbounds = leafletBoundsHelpers.createBoundsFromArray([
            [ southWestLat, southWestLng],
            [ northEastLat, northEastLng]
        ]);

        angular.extend($scope, {
            maxbounds: maxbounds,
            center: {
                lat: 45.653226,
                lng: -84.7494816,
                zoom: 5
            },
            tiles: {
                url: " "
            },
            defaults: {
                maxZoom: 16
            }
        });

        //Use local tile server in OpenStreetMap
        //Server URL is configurable in booking.properties. Use service call to get the URL
        var host = null;
        var getTileServerDataCallback = function (result) {
            host = result.MapTileServerHost;
            dtCacheServices.put(dtCacheServicesKeys.TILE_SERVER_URL, host);
            $scope.tiles = {
                url: host + "/{z}/{x}/{y}.png"
            };
        };

        var getTileServerDataErrorCallback = function (error) {
            $log.error('dtcLocationData.getTileServer ---> ', error);
        };
        var url =  dtCacheServices.get(dtCacheServicesKeys.TILE_SERVER_URL);//dtcLocationData.getTileServerURL();
        if( url === undefined) {
            dtcLocationData.getTileServer(getTileServerDataCallback, getTileServerDataErrorCallback);
        }else{
            $scope.tiles = {
                url: url + "/{z}/{x}/{y}.png"
            };
        }

        $scope.showFrench = false;
        $scope.openSaturday = false;

        $scope.broadcastSaturdayOrFrench = function (openSaturday, showFrench) {
            $scope.showFrench = showFrench;
            $scope.openSaturday = openSaturday;
            $scope.driver.locationName = undefined; //disables continue button
            $rootScope.$broadcast("refresh-map-saturday-french", openSaturday, showFrench, true);
        };

        stepState.checkStepState('booking.location');
        if (!userService.isSet() && userService.isEligible()) {
            stepState.changeState('booking.registration');
            return;
        }
        $scope.driver = userService.getCurrentUser();

        if ($scope.driver.locationId !== undefined) {
            $scope.userSelectedLocationId = $scope.driver.locationId;
        }

        $scope.loading = false;
        $scope.errMsg = errorService.clean($scope);

        $scope.errorCallback = function (error) {
            $log.error('loaded with error ---> ', error);
            stepState.goToGeneralErrorState();
        };

        $scope.previousStep = function () {
            stepState.changeState('booking.licence');
        };

        //This function is called by the directive 'ngDtcList' in dtc-map.js
        $scope.mapCallback = function (location) {
            $log.log('Booking Page --> callback to maps directive', location);

            $scope.driver.locationName = location.name;
            $scope.driver.locationId = location.id;

            var licenceClass = $scope.driver.licenceClass;

            if ($scope.driver.endorsement !== undefined) {
                licenceClass = $scope.driver.licenceClass + $scope.driver.endorsement;
            }

            $scope.driver.serviceId = dtcLocationData.getServiceId($scope.driver.locationId, licenceClass, $scope.driver.endorsement, $scope.driver.component);

            // save user if page isn't current
            var user = userService.getCurrentUser();

            user.locationName = $scope.driver.locationName;
            user.locationId = $scope.driver.locationId;
            user.dtcAddress1 = location.address1;
            user.frenchTestRequested = $scope.showFrench;
            user.timezone = location.timezone;
            user.momentTimezone = bookingTimeService.getMomentTimezone(user.timezone);
            user.dtcAddress2 = location.address2;
            user.serviceId = $scope.driver.serviceId;
            user.selectedDate = undefined;
            user.selectedTime = undefined;

            userService.setUser(user);

            if (stepState.currentStepName() === 'booking.timeslot') {
                stepState.changeState('booking.calendar', true);
            }
        };

        $scope.locationSubmit = function () {
            // licenceClass
            var user = userService.getCurrentUser();
            if (user.timezone == undefined) {
                dtcLocationData.setUserTimeZone($scope.driver.locationId)
                stepState.changeState('booking.calendar');
            }
            else {
                stepState.changeState('booking.calendar');
            }
        };
    }]);
'use strict';

var dtBookingApp = angular.module('dtBooking');

dtBookingApp.controller('loginController', ['$scope','$log','$state', '$window','userService', 'stepState', '$routeParams', 'locale','loginSvc', 'errorService','$location', 'vcRecaptchaService',
function($scope,$log,$state,$window,userService,stepState,$routeParams,locale, loginSvc, errorService, $location, recaptcha) {
    $scope.driverInfo = {};
    $scope.locale = locale;
    $scope.captcha_response = null;
    $log.log('loginController $state ---> ', $state);

    $scope.master = {};
    $scope.loading = false;
    $scope.isEmailTokenVerified = false;
    $scope.emailVerificationInProgress = true;
    $log.debug('token= '+($location.search()).token);
    $log.debug('$email= ' + ($location.search()).email);

    $scope.token = ($location.search()).token;
    $scope.email = ($location.search()).email;

    if ($window.location.href !== $location.absUrl() && $window.location.href.indexOf("#/verify-driver") !== -1 && $scope.token === undefined) {
        var url = $window.location.href;
        var params = url.substr(url.indexOf("token"));
        $scope.email = params.substr(params.indexOf("&email") + 7);
        $scope.token = decodeURIComponent(params.substr(6, params.indexOf("&email") - 6));
    }

    stepState.checkStepState('booking.login');
    $scope.driver = userService.getCurrentUser();

    $scope.errMsg = errorService.clean($scope);

    $scope.setValidationFocus = function () {
        if(angular.isUndefined($scope.driverInfo.licenceNumber) || !$scope.driverInfo.licenceNumber.$valid) {
            $scope.$broadcast('focusLicenceNumber');
            return false;
        } else if(angular.isUndefined($scope.driverInfo.licenceExpiryDate) || !$scope.driverInfo.licenceExpiryDate.$valid) {
            $scope.$broadcast('focusLicenceExpiry');
            return false;
        } else {
            return true;
        }
    };

    var _onLoginCompleted = function(result) {
        result = result.data;
        $scope.loading = false;
        if(result.registered) {
            stepState.changeState('booking.verify');
        } else if(result.authenticated) {
            var user =  userService.getCurrentUser();
            user.customerId = result.customerId;
            user.firstName = result.firstName;
            user.lastName = result.lastName;
            user.email = result.email;

            userService.setUser(user);
            if($scope.email !== undefined) {
                stepState.changeState('booking.licence');
            } else {
                stepState.changeState('booking.dashboard');
            }
        } else {
            if (!$scope.email) {
                recaptcha.reload();
            }
            if(angular.isDefined(result.validationErrors)) {
                angular.forEach(result.validationErrors, function (value, key) {
                    if (key.contains(".")) {
                        var errorCodes = key.split('.');
                        $scope.driverInfo[errorCodes[0]].$setValidity(errorCodes[1], false);
                    }
                });
            } else {
                errorService.showServerError($scope, result);
            }
        }
    };

    var _onLoginError = errorService.errorHandler($scope, function(response){

        $scope.loading = false;

        var messageKey, returnVal = false;

        switch(response.status){
            case 401:
                messageKey = 'booking.app.message.invalid.credentials';
                returnVal = true;
                break;
            case 409:
                messageKey = 'booking.app.message.invalid.email';
                returnVal = true;
                break;
            case 410:
                messageKey = 'booking.app.message.email.token.expired';
                returnVal = true;
                $scope.sameLicenceError = true;
                break;
            case 415:
                messageKey = 'booking.app.message.unsupported';
                returnVal = true;
                break;
            case 417:
                messageKey = 'booking.app.message.email.unverified';
                returnVal = true;
                break;
        }

        $scope.errMsg = {hasError:true};
        $scope.errMsg.msg = errorService.getMessage(messageKey);

        //if we are not in verify email flow reload captcha
        if ($scope.email === undefined) {
            recaptcha.reload($scope.recaptchaId);
            $scope.captcha_response = null;
            $scope.driverInfo.form_submitted = false;
        }

        return returnVal;

    });

    $scope.setRecaptchaId = function(widgetId) {
        $scope.recaptchaId = widgetId;
    };

    $scope.driverSubmit = function(){
        $window.focusOnSubmit($window, $('#booking-login'));

        var user =  userService.getCurrentUser();
        user.licenceExpiry = $scope.driver.licenceExpiry;
        user.licenceNumber = $scope.driver.licenceNumber.toUpperCase();
        user.email = $scope.email;
        user.emailConfirm = $scope.emailConfirm;
        user.token = $scope.token;

        userService.setUser(user);

        $scope.invalid_license_info = false;
        $scope.loading = true;

        if ($scope.token) {
            user.emailToken = $scope.token;
            loginSvc.verify(user, _onLoginCompleted, _onLoginError);
        } else {
            loginSvc.login(user, $scope.captcha_response, _onLoginCompleted, _onLoginError);
        }
    };

    $scope.setResponse = function(response) {
        $scope.captcha_response = response;
    };

    //Not required
    $scope.resetServerSideValidation = function(element, error){
        $log.debug('resetServerSideValidation - element --------------> ', element);
        $log.debug('resetServerSideValidation - error --------------> ', error);
        $log.debug('resetServerSideValidation - $scope.driverInfo[element] --------------> ', $scope.driverInfo[element]);
        if(angular.isDefined($scope.driverInfo[element])) {
            $scope.driverInfo[element].$setValidity(error, true);
        }
    };

    $scope.clearFieldValidation = function (element) {
      if(angular.isDefined($scope.driverInfo[element]) && $scope.driverInfo.form_submitted) {
          $scope.driverInfo[element].$setPristine();
          $scope.driverInfo[element].$error = {};
      }
    };

    $scope.validateEmailToken = function(){
        if(angular.isDefined($scope.token)) {
            $scope.driver.licenceNumber = '';
            $scope.driver.licenceExpiry = '';
            loginSvc.verifyEmailToken($scope.token, $scope.email, function (result) {
                result = result.data;
                $scope.emailVerificationInProgress = false;
                if (result.activated) {
                    $scope.isEmailTokenVerified = true;
                } else {
                    $scope.otherError = true;
                }
            }, errorService.errorHandler($scope, function(response){
                $scope.emailVerificationInProgress = false;
                var messageKey, returnVal = false;

                switch(response.status){
                    case 401:
                        messageKey = 'booking.app.message.invalid.credentials';
                        returnVal = true;
                        break;
                    case 409:
                        messageKey = 'booking.app.message.invalid.email';
                        returnVal = true;
                        break;
                    case 410:
                        messageKey = 'booking.app.message.email.token.expired';
                        $scope.expiredTokenError = true;
                        returnVal = true;
                        $scope.sameLicenceError = true;
                        break;
                    case 417:
                        messageKey = 'booking.app.message.email.unverified';
                        returnVal = true;
                        break;
                    default:
                        $scope.otherError = true;
                }

                $scope.errMsg = {hasError:true};
                $scope.errMsg.msg = errorService.getMessage(messageKey);

                return returnVal;
            }));
        } else {
            $scope.isEmailTokenVerified = true;
            $scope.emailVerificationInProgress = false;
        }
    };

    $scope.registrationStep = function(){
        stepState.changeState('booking.registration');
    };

}]);

'use strict';

var dtBookingApp = angular.module('dtBooking');

dtBookingApp.controller('paymentController', ['$filter', '$scope', '$rootScope', '$log', '$state', '$sce', 'stepState', 'userService', 'bookingTimeService', 'locale', 'bookingFees', 'errorService', '$window', '$cookies', 'appointmentService',
    function($filter, $scope, $rootScope, $log, $state, $sce, stepState, userService, bookingTimeService, locale, bookingFees, errorService, $window, $cookies, appointmentService) {
        $scope.driverInfo = {};
        $scope.locale = locale;
        $log.log('$state ---> ', $state);

        $scope.master = {};

        stepState.checkStepState('booking.payment');

        $scope.driver = userService.getCurrentUser();

        $scope.timeForDisplay = bookingTimeService.getFormattedTime($scope.driver.timeslot, $scope.driver.momentTimezone);

        $scope.TermsAccepted = false;

        $scope.driverSubmit = function() {
            if ($scope.TermsAccepted) {
                appointmentService.getStatusToken(
                    userService.getCurrentUser().licenceNumber,
                    function (result) {
                        $log.debug("SUCCESS Login Status: " + result.statusToken);
                        $log.debug("Login status defined: " + angular.isDefined(result.statusToken));

                        if (!angular.isDefined(result.statusToken) || !userService.isSet()) {
                            stepState.changeState('booking.quit');

                        } else if ($scope.driver.fees.total <= 0) {
                            stepState.changeState('booking.finalize');

                        } else {
                            openPaymentPopup();
                        }
                    },
                    function (result) {
                        $log.debug("ERROR Login Status");
                        errorService.showServerError($scope, 'ERROR Login Status');
                    }
                );
            }

            function openPaymentPopup() {
                var endorsement = $scope.driver.endorsement === undefined || $scope.driver.endorsement === null ? "" : $scope.driver.endorsement;
                var component = $scope.driver.component === undefined || $scope.driver.component === null ? "" : $scope.driver.component;
                var isAppointmentIdNotDefined = $scope.driver.heldAppointment.guid === undefined || $scope.driver.heldAppointment.guid === '';
                var isLicenceClassNotDefined = $scope.driver.licenceClass === undefined || $scope.driver.licenceClass === '';
                var isRescheduleNotBoolean = typeof $scope.driver.reschedule !== 'boolean';
                var invalidUrlParams = isAppointmentIdNotDefined || isLicenceClassNotDefined || isRescheduleNotBoolean;

                if (invalidUrlParams) {
                    errorService.showServerError($scope, 'ERROR Invalid URL Parameters');
                } else {
                    if (typeof $scope.paymentPopup == 'undefined' || $scope.paymentPopup === null || $scope.paymentPopup.closed) {
                        $scope.paymentPopup = $window.open(getPaymentGatewayUrl(endorsement, component), '_blank', getWindowParams());

                    } else {
                        $scope.paymentPopup.focus();
                    }
                }
            }

            function getPaymentGatewayUrl(endorsement, component) {
                var location = window.location.href;
                var query = location.substring(location.lastIndexOf('?')+1);
                var langId = '1';
                if(query.contains('language_id')) {
                    var indexOfLangId = query.indexOf('language_id') + 'language_id='.length;
                    langId = query.substring(indexOfLangId, indexOfLangId+1);
                }
                return '#/payment-submit' + '?guid=' +
                    $scope.driver.heldAppointment.guid +
                    '&licenceClass=' + $scope.driver.licenceClass +
                    '&endorsement=' + endorsement +
                    '&component=' + component +
                    '&reschedule=' + $scope.driver.reschedule +
                    '&language_id=' + langId +
                    ($scope.driver.previousHold === undefined ? '' : '&existingAppt=' + $scope.driver.previousHold.confirmed.guid);
            }

            function getWindowParams() {
                return 'height=' + $window.innerHeight + ',width=1005';
            }
        };

        $scope.agreeWithTermsAndConditions = function(checkStatus) {
            if (checkStatus) {
                $scope.TermsAccepted = true;
            } else {
                $scope.TermsAccepted = false;
            }
        };

        $scope.popupTermsAndConditions = function() {
            if (typeof $scope.termsAndConditionsPopup == 'undefined' || $scope.termsAndConditionsPopup === null || $scope.termsAndConditionsPopup.closed) {
                $scope.termsAndConditionsPopup = $window.open('/home/terms.html', '_blank', 'height=' + $window.innerHeight + ',width=1005');

            } else {
                $scope.termsAndConditionsPopup.focus();
            }
        };
        
        $scope.editBooking = function() {
            stepState.changeState('booking.timeslot');
        };

        $scope.paymentCancelled = function() {
            $log.log('Payment Cancelled');
        };

        $scope.paymentEligibilityError = function(response) {
            $log.log('Payment Eligibility Error');

            errorService.errorHandler($scope, function(response){
                var messageKey = 'booking.error', returnVal = false;

                if(response.status === 400 || response.status === 500 ) {
                    messageKey = 'booking.app.message.incomplete.payment';
                    returnVal = true;
                }

                $scope.errMsg = {hasError:true};
                $scope.errMsg.msg = errorService.getMessage(messageKey);

                return returnVal;
            });
        };

        $scope.paymentError = function(response) {
            $log.log('Payment Error');
            errorService.showServerError($scope, response);
        }

        $scope.paymentSuccessful = function(confirmationNumber, id) {
            $log.log('Payment Successful '+confirmationNumber);
            $scope.driver.payment = {};
            $scope.driver.payment.confirmationNumber = confirmationNumber;
            $scope.driver.payment.id = id;

            stepState.changeState('booking.finalize');
        };

        window.$windowScope = $scope;

    }]);

/**
 * Created by Hisham on 2014-10-06.
 */
'use strict';

var dtBookingApp = angular.module('dtBooking');

dtBookingApp.controller('paymentSubmitController', ['$scope', '$rootScope', '$log', '$sce', 'stepState', 'userService', 'prepareBookingPaymentService', 'locale',
    function($scope, $rootScope, $log, $sce, stepState, userService, prepareBookingPaymentService, locale) {

        var QueryString = function () {
            // This function is anonymous, is executed immediately and
            // the return value is assigned to QueryString!
            var query_string = {};
            var location = window.location.href;
            var query = location.substring(location.lastIndexOf('?')+1);
            var vars = query.split("&");
            for (var i=0;i<vars.length;i++) {
                var pair = vars[i].split("=");
                // If first entry with this name
                if (typeof query_string[pair[0]] === "undefined") {
                    query_string[pair[0]] = pair[1];
                    // If second entry with this name
                } else if (typeof query_string[pair[0]] === "string") {
                    var arr = [ query_string[pair[0]], pair[1] ];
                    query_string[pair[0]] = arr;
                    // If third or later entry with this name
                } else {
                    query_string[pair[0]].push(pair[1]);
                }
            }
            return query_string;
        } ();

        var bookingFeeRequest;

        if (QueryString.endorsement !== "undefined" &&
            QueryString.endorsement !== "") {

            bookingFeeRequest = {
                licenceClass: QueryString.licenceClass,
                endorsement: QueryString.endorsement,
                reschedule: QueryString.reschedule,
                existingAppointmentGuid: QueryString.existingAppt
            };

        } else {
            bookingFeeRequest = {
                licenceClass: QueryString.licenceClass,
                reschedule: QueryString.reschedule,
                existingAppointmentGuid: QueryString.existingAppt
            };
        }

        var user = userService.getCurrentUser();

        user.licenceClass = QueryString.licenceClass === "" ? undefined : QueryString.licenceClass;
        user.endorsement = QueryString.endorsement === "" ? undefined : QueryString.endorsement;
        user.component = QueryString.component === "" ? undefined : QueryString.component;
        user.checkSameLicense = false;

        userService.setUser(user);

        userService.getEligibility(function(response) {

            if (response.eligible) {
                prepareBookingPaymentService.preparePayment(QueryString.guid, bookingFeeRequest,
                function (data, status, headers, config) {
                    if (data.status && data.status.errorCode) {
                        stepState.goToGeneralErrorState();

                    } else {
                        data.url = $sce.trustAsResourceUrl(data.url);
                        $rootScope.$broadcast('gateway.redirect', data);
                    }
                },
                function (data, status, headers, config) {
                    $scope.parentWindow = window.opener.$windowScope;
                    $scope.parentWindow.paymentError({status: status});
                    window.close();
                });
            } else {
                $scope.parentWindow = window.opener.$windowScope;
                $scope.parentWindow.paymentEligibilityError(response);

                window.close();
            }

        });


    }
]);
'use strict';

angular.module('dtBooking').controller('quitController', ['$scope','userService','registrationSvc','$window', function($scope,userService,registrationSvc,$window){
    function cleanSession(){
        document.cookie = 'sid' +'=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
        document.cookie = 'token' +'=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
        userService.setUser({});
        $window.location.href = '/book-a-road-test';
    }
    registrationSvc.removeToken(userService.getCurrentUser().licenceNumber, cleanSession, cleanSession);
}]);

'use strict';

var dtBookingApp = angular.module('dtBooking');

dtBookingApp.controller('registrationController', ['$scope','$log','$state', '$window', 'userService', 'stepState', '$routeParams', 'locale','registrationSvc', 'errorService', 'vcRecaptchaService',
function($scope,$log,$state,$window,userService,stepState,$routeParams,locale, registrationSvc, errorService, recaptcha, $locale) {
    $scope.driverInfo = {};
    $scope.locale = locale;
    $scope.captcha_response = null;
    $log.log('registrationController $state ---> ', $state);

    $scope.master = {};
    $scope.loading = false;

    $log.debug('action = '+$routeParams.action);
    $log.debug('$routeParams = ', $routeParams);

    stepState.checkStepState('booking.registration');
    $scope.driver = userService.getCurrentUser();

    $scope.errMsg = errorService.clean($scope);

    function warnsPopupBlocker() {
        $scope.errMsg = {hasError:true};
        $scope.errMsg.msg = errorService.getMessage('booking.app.message.popup.blocked');
    }

    function checkPopup() {
        var popup = $window.open(null, '_blank', 'height=100, width=100, top=' + ($window.screenY) + ', left=' + ($window.screenX));
        if(!popup)
            warnsPopupBlocker();
        else {
            popup.onload = function() {
              setTimeout(function() {
                  if(popup.screenX === 0)
                      warnsPopupBlocker();
              }, 0);
            };
            popup.close();
        }
    }

    checkPopup();

    $scope.setValidationFocus = function () {
        if(angular.isUndefined($scope.driverInfo.emailAddress) || !$scope.driverInfo.emailAddress.$valid) {
            $scope.$broadcast('focusEmail');
            return false;
        } else if(angular.isUndefined($scope.driverInfo.confirmEmailAddress) || !$scope.driverInfo.confirmEmailAddress.$valid) {
            $scope.$broadcast('focusEmailConfirm');
            return false;
        } else if(angular.isUndefined($scope.driverInfo.licenceNumber) || !$scope.driverInfo.licenceNumber.$valid) {
            $scope.$broadcast('focusLicenceNumber');
            return false;
        } else if(angular.isUndefined($scope.driverInfo.licenceExpiryDate) || !$scope.driverInfo.licenceExpiryDate.$valid) {
            $scope.$broadcast('focusLicenceExpiry');
            return false;
        } else {
            return true;
        }
    };

    $scope.driverSubmit = function(){
        $window.focusOnSubmit($window, $('#booking-registration'));

        var user =  userService.getCurrentUser();
        user.licenceExpiry = $scope.driver.licenceExpiry;
        user.licenceNumber = $scope.driver.licenceNumber.toUpperCase();
        user.email = $scope.driver.emailAddress;
        user.emailConfirm = $scope.driver.confirmEmailAddress;

        userService.setUser(user);

        $scope.invalid_license_info = false;
        $scope.loading = true;
        registrationSvc.register(user, $scope.captcha_response, function(result){
            result = result.data;
            $scope.loading = false;
            if(result.registered) {
                stepState.changeState('booking.verify');
            } else if(result.authenticated) {
                var user =  userService.getCurrentUser();
                user.customerId = result.customerId;
                user.firstName = result.firstName;
                user.lastName = result.lastName;
                user.email = result.email;
                userService.setUser(user);
                stepState.changeState('booking.licence');
            } else {
                if(angular.isDefined(result.validationErrors)) {
                    angular.forEach(result.validationErrors, function (value, key) {
                        if (key.contains(".")) {
                            var errorCodes = key.split('.');
                            $scope.driverInfo[errorCodes[0]].$setValidity(errorCodes[1], false);
                        }
                    });
                } else {
                    errorService.showServerError($scope, result);
                }
            }
        }, errorService.errorHandler($scope, function(response){
            $scope.loading = false;
            var messageKey, returnVal = false;
            recaptcha.reload();
            $scope.captcha_response = null;
            $scope.driverInfo.form_submitted = false;

            switch(response.status){
                case 401:
                    messageKey = 'booking.app.message.invalid.credentials';
                    returnVal = true;
                    break;
                case 409:
                    messageKey = 'booking.app.message.invalid.email';
                    returnVal = true;
                    break;
                case 410:
                    messageKey = 'booking.app.message.email.token.expired';
                    returnVal = true;
                    break;
                case 415:
                    messageKey = 'booking.app.message.unsupported';
                    returnVal = true;
                    break;
                case 417:
                    messageKey = 'booking.app.message.email.unverified';
                    returnVal = true;
                    break;
            }

            $scope.errMsg = {hasError:true};
            $scope.errMsg.msg = errorService.getMessage(messageKey);

            return returnVal;
        }));
    };

    $scope.setResponse = function(response) {
        $scope.captcha_response = response;
    };

    $scope.resetServerSideValidation = function(element, error){
        if(angular.isDefined($scope.driverInfo[element])) {
            $scope.driverInfo[element].$setValidity(error, true);
        }
    };

    $scope.clearFieldValidation = function (element) {
        if(angular.isDefined($scope.driverInfo[element]) && $scope.driverInfo.form_submitted) {
            $scope.driverInfo[element].$setPristine();
            $scope.driverInfo[element].$error = {};
        }
    };

    $scope.editBooking = function(){
        stepState.changeState("booking.login");
    };

}]);

'use strict';

var dtBookingApp = angular.module('dtBooking');

dtBookingApp
    .constant('_', window._)
    .controller('timeslotController', ['$scope', '$filter','$window', '$log','$state', 'userService', 'stepState', 'bookingTimeService', 'appointmentService','locale', 'bookingFees', 'errorService', 'appointmentCancelService',
    function($scope, $filter, $window, $log, $state, userService, stepState, bookingTimeService, appointmentService, locale, bookingFees, errorService, appointmentCancelService) {
        $scope.locale = locale;
        $scope.locationInfo = {};
        $log.log('timeslotController $state ---> ', $state);

        $scope.master = {};
        $scope.loadingData = true;
        $scope.timeslotSubmitted = false;

        stepState.checkStepState('booking.timeslot');
        if (!userService.isSet()) {
            stepState.changeState('booking.registration');
            return;
        }

        $scope.driver = userService.getCurrentUser();

        $scope.loading = false;
        enableBookingContainer();
        $scope.errMsg = errorService.clean($scope);

        $scope.noTimeslots = function () {
            return $scope.timeslots.length == 0;
        };

        var _refreshTimeSlots = function(user, onSuccess) {
            $scope.loadingData = true;
            $scope.errMsg = errorService.clean($scope);
            $scope.driver.timeslot = undefined;
            disableBookingContainer();
            // if the selected date is changed, reset all of this
            bookingTimeService.getTimeSlots(user.serviceId, user.selectedDate, function(response) {
                enableBookingContainer();
                $scope.loadingData = false;
                user.timeslots = response.availableBookingTimes;

                if (bookingTimeService.hasHeldAppointmentDate($scope.driver.selectedDate)) {
                    bookingTimeService.addHeldAppointmentToCurrentUser();
                }

                userService.setUser(user);
                stepState.changeState('booking.timeslot', true);
                $scope.timeslots = user.timeslots;

                $log.log('$scope.driver.selectedDate ---> ', $scope.driver.selectedDate);

                if (onSuccess !== undefined) {
                    onSuccess($scope.timeslots);
                }

            }, errorService.errorHandler($scope));
        };

        $scope.$on('user:setSelectedDate', function(event,user) {
            $scope.timeslotSubmitted = false;
            _refreshTimeSlots(user);
        });

        $scope.$on('user:updated', function(event,user) {
            // if location is changed, nuke this panel
            //$scope.timeslots = undefined;
            //$scope.driver.timeslot = undefined;
        });

        if ($scope.driver.selectedDate !== undefined) {
            if($scope.driver.timeslots === null || angular.isUndefined($scope.driver.timeslots)){
                _refreshTimeSlots($scope.driver);
            } else {
                $scope.loadingData = false;
                $scope.timeslots = $scope.driver.timeslots;
            }
        }

        $scope.previousStep = function() {
            stepState.changeState('booking.location', false);
            $scope.driver.timeslot = undefined;
        };

        $scope.backToCalendar = function() {
            stepState.changeState('booking.calendar', false);
            $scope.driver.timeslot = undefined;
        };

        var _calculateFeesAndGoToBookingPayment = function() {
            var user = userService.getCurrentUser();
            bookingFees.calculate(user, function(fees) {
                user.fees = fees;
                userService.setUser(user);
                stepState.changeState('booking.payment');
                enableBookingContainer();
            }, errorService.errorHandler($scope, function (response) {
                $scope.loading = false;
                enableBookingContainer();
                errorService.setBlocked($scope);
            }));
        };

        $scope.driverSubmit = function() {
            $scope.errMsg = errorService.clean($scope);
            $scope.loading = true;
            disableBookingContainer();
            $scope.timeslotSubmitted = true;
            userService.getCurrentUser();

            function setError(messageKey) {
                $scope.errMsg = {
                    msg: errorService.getMessage(messageKey),
                    hasError: true
                };
            }

            var cancelExistingAppointment = userService.appointmentHoldExists();

            function clearHeldAppointment() {

                if (cancelExistingAppointment) {
                    userService.clearHeldAppointment();
                }
            }

            if (! userService.isAppointmentSameAsCurrentlyHeld()) {
                //Doesn't have held appointment

                if (cancelExistingAppointment) {
                    //Held appointment exists

                    appointmentCancelService.cancelHeldAppointment($scope.driver.heldAppointment.guid,
                        function(response) {
                            $log.log('Successfully cancelled appointment hold - guid=' + $scope.driver.heldAppointment.guid);
                        },
                        function() {
                            $log.log('Error trying to cancel appointment hold - guid=' + $scope.driver.heldAppointment.guid);
                        });
                }

                var user = userService.getCurrentUser();
                var licenceEndorsement;

                if (user.component !== undefined) {
                    licenceEndorsement = user.component;
                } else {
                    licenceEndorsement = $scope.driver.endorsement;
                }

                appointmentService.holdAppointment($scope.driver.serviceId, $scope.driver.timeslot,
                                                    $scope.driver.licenceClass, licenceEndorsement, $scope.driver.frenchTestRequested, function (response) {
                        if (response.success === undefined || response.success === false) {
                            $scope.loading = false;
                            enableBookingContainer();
                            errorService.showServerError($scope, response);
                            return;
                        }
                        var holdSuccess = response.success;

                        if (holdSuccess) {
                            userService.createAppointmentHold(response.guid);
                            _calculateFeesAndGoToBookingPayment();

                        } else {
                            clearHeldAppointment();

                            $scope.loading = false;
                            enableBookingContainer();
                            errorService.showServerError($scope, response);
                        }
                    }, errorService.errorHandler($scope, function (response) {

                        clearHeldAppointment();

                        $scope.loading = false;
                        enableBookingContainer();
                        errorService.setBlocked($scope);

                        var returnVal = false;

                        switch(response.status){
                            case 403:
                                setError('booking.app.message.appointment.3.6.check');
                                returnVal = true;
                                break;
                            case 409:
                                _refreshTimeSlots($scope.driver, function(timeslots) {
                                    setError('booking.app.message.timeslot.conflict');
                                });
                                returnVal = true;
                                break;
                            case 420:
                                setError('booking.error.cancel');
                                returnVal = true;
                                break;
                            case 503:
                                setError('booking.error.cancel.48.hours');
                                returnVal = true;
                                break;
                            default:
                                setError();
                        }

                        return returnVal;
                    }));
            } else {
                _calculateFeesAndGoToBookingPayment();
            }
        };
}
]);

'use strict';

var dtBookingApp = angular.module('dtBooking');

dtBookingApp.controller('verifyEmailController', ['$scope','$log','$state', '$routeParams', 'userService', function($scope,$log,$state,$routeParams, userService) {
    $log.debug('verifyEmailController $state ---> ', $state);

    $scope.email = userService.getCurrentUser().email;
}]);

'use strict';

angular.module('dtBooking')
    .directive('dtMotorcyclespeeds', ['$log','$translate',function($log,$translate){

		return {
			controller: 'ComponentController',

			scope: false,
			restrict: 'AE',
			templateUrl: '/ng/dt-motorcyclespeeds',
			replace: true,
			transclude: false,

			link: function(scope, element , iAttrs, componentController) {
				$log.debug('--- motorcyclespeeds LiNK ---');
				componentController.init(element);
			}
		};
	}])
	.controller('ComponentController', function($scope, $attrs, $filter, $log) {

		var self = this;

		this.init = function( element, componentClass ) {
			self.$element = element;
		};

		function watchComponent(value) {
			if (angular.isDefined($scope.classesList)) {
				if ($scope.driver !== undefined) {
					value = $scope.driver.licenceClass;
				}
				var componentList = $filter('filter')($scope.classesList.primaryDriverLicenceClasses, {licenceClass: value}, true);

				if (angular.isDefined(componentList)) {
					if ((componentList.length > 0) && angular.isDefined(componentList[0].components)) {
						$scope.components = componentList[0].components;
						if ($scope.driver.component === undefined) {
							$scope.driver.component = "F";
						}

					} else {
						$scope.components = undefined;
					}
				}
			}
		}

		$scope.$watch('driver.licenceClass', watchComponent);
		$scope.$watch('classesList', watchComponent);

	});
'use strict';

angular.module('dtBooking')
    .service('appointmentService', ['$log', '$window', '$resource', '$http', 'userService', function($log, $window, $resource, $http, userService) {
        var _within48hours = $resource($window.dtServiceEndpoint + '/appointment/within48hours/:guid',
            { }, {
                get: {method: 'GET', params:{  }  }
            }
        );

        var _pendingAppointments = $resource($window.dtServiceEndpoint + '/appointment/pending',
            { }, {
                get: {
                    method: 'GET',
                    params: {  },
                    cache : false
                }
            }
        );

        var _cancelAppointment = $resource($window.dtServiceEndpoint + '/appointment/cancel/:guid',
            { }, {
                cancel: {method: 'DELETE', params:{  }  }
            }
        );

        var _appointmentHold = $resource($window.dtServiceEndpoint + '/booking/hold', {}, {
            holdAppointment: {
                method: 'POST',
                params: {}
            }
        });

        this.with48hours = function(appointmentId, success, errorHandler) {
            $log.log('appointmentService.cancelAppointment');
            return _within48hours.get({guid: appointmentId}).$promise.then(success, errorHandler);
        };

        this.pendingAppointments = function(success, errorHandler, finalfn) {
            $log.log('appointmentService.pendingAppointments');
            if (finalfn === undefined) {
                finalfn = function() {};
            }
            return _pendingAppointments.get({iecachebuster: new Date().getTime()}).$promise.then(function(result) {
                userService.setUserAppointments(result.pendingAppointments);
                success(result);
            }, errorHandler)
                .finally(finalfn);
        };

        this.cancelAppointment = function(appointmentId, success, errorHandler) {
            $log.log('appointmentService.cancelAppointment');
            return _cancelAppointment.cancel({guid: appointmentId}).$promise.then(success, errorHandler);
        };

        this.holdAppointment = function(serviceId, time, licenceClass, endorsement, frenchTestRequested, success, errorHandler) {
            $log.log('holdAppointment start --->');
            return _appointmentHold.holdAppointment({
                serviceId: serviceId,
                time: time,
                licenceClass: licenceClass,
                endorsement: endorsement,
                frenchTest: frenchTestRequested
            }).$promise.then(success, errorHandler);
        };

        var _statusToken = $resource($window.dtServiceEndpoint + '/status', {}, {
                getToken: {
                    method: 'POST',
                    params: {}
                }
        });

        this.getStatusToken = function(licenceNumber, success, errorHandler) {
            if (!angular.isDefined(licenceNumber)) {
                errorHandler(licenceNumber);
            } else {
                return _statusToken.getToken({
                    licenceNumber: licenceNumber
                }).$promise.then(success, errorHandler);
            }
        }
    }]);
'use strict';

var dtBookingApp = angular.module('dtBooking');

dtBookingApp.factory('createBooking', ['$resource',
    function ($resource) {
        return $resource('/app/booking/booking/createBooking');
    }
]);

dtBookingApp.service('licenceClassService', ['$resource', '$log', '$window',
    function ($resource, $log, $window) {

        var _licenceClasses = $resource($window.dtServiceEndpoint + '/licenceClass', {}, {
            check: {
                method: 'GET',
                params: {}
            }
        });

        this.loadLicenceClasses = function (success, errorHandler) {
            $log.log('loadLicenceClasses start ---> ');
            return _licenceClasses.check({iecachebuster: new Date().getTime()}).$promise.then(success, errorHandler);
        };
    }
]);

dtBookingApp.service('bookingTimeService', ['$resource', '$log', '$filter', '$window', 'locale', 'userService',
    function ($resource, $log, $filter, $window, locale, userService) {
        var _timeSlots = $resource($window.dtServiceEndpoint + '/booking', {}, {
            check: {
                method: 'GET',
                params: {}
            }
        });

        this.hasHeldAppointmentDate = function (selectedCalendarDate) {
            var user = userService.getCurrentUser();
            var heldAppointment = user.heldAppointment;

            if ((selectedCalendarDate !== null || angular.isDefined(selectedCalendarDate))
                && (heldAppointment && heldAppointment.timeslot)
                && (heldAppointment.serviceId === user.serviceId)) {
                var retValue = this.dateMatch(selectedCalendarDate, heldAppointment, user.timezone);
                return retValue;
            }
        };

        this.dateMatch = function (selectedCalendarDate, heldAppointment, testCenterTimezone) {
            var dateArr = selectedCalendarDate.split('-');
            var retValue = !!(dateArr[0] === $filter('date')(heldAppointment.timeslot, 'yyyy', testCenterTimezone) && (dateArr[1] === $filter('date')(heldAppointment.timeslot, 'MM', testCenterTimezone)) && (dateArr[2] === $filter('date')(heldAppointment.timeslot, 'dd', testCenterTimezone)));
            return retValue;
        };

        this.addHeldAppointmentToCurrentUser = function () {
            var user = userService.getCurrentUser();

            if(!user.timeslot && user.heldAppointment){
                user.timeslot = user.heldAppointment.timeslot;
            }
            user.timeslots.push(this.getFormattedTime(user.timeslot, user.momentTimezone));
            //Add the held appointment to the available timeslots and sort the list
            user.selectedTime = user.timeslot = user.heldAppointment.timeslot;
            user.timeslots = $filter('orderBy')(user.timeslots, 'timestamp',false);
            user.timeslots = $window._.uniq(user.timeslots, 'timestamp');
        };

        var _getFormattedTime = function (timestamp, momentTimezone) {
            if (!locale.isFrench()) {
                return {timestamp: timestamp, display: moment.tz(timestamp, momentTimezone).format("h:mm A")};
            } else {
                var frenchMoment = moment.tz(timestamp, momentTimezone);
                frenchMoment.locale('fr-ca');
                return {timestamp: timestamp, display: frenchMoment.format("H[h]mm")};
            }
        };

        this.getTimeSlots = function (serviceId, date, success, errorHandler) {
            $log.log('getTimeSlots start ---> ');
            return _timeSlots.check({
                is: serviceId,
                date: date
            }).$promise.then(function (results) {
                    var user = userService.getCurrentUser();
                    for (var i = 0; i < results.availableBookingTimes.length; i++) {
                        var temp = results.availableBookingTimes[i];
                        results.availableBookingTimes[i] = _getFormattedTime(temp.timeslot, user.momentTimezone);
                    }
                    success(results);
                }, errorHandler);
        };

        this.getFormattedTime = function (timestamp, momentTimezone) {
            return _getFormattedTime(timestamp, momentTimezone);
        };

        this.getMomentTimezone = function (theTimezone) {
            var timeZone;
            //get timezone offset (daylight savings aware)
            if (theTimezone === 'EDT' || theTimezone === 'EST') {
                timeZone = "America/Toronto";
            } else if (theTimezone === 'CDT' || theTimezone === 'CST') {
                timeZone = "America/Winnipeg";
            } else {
                throw 'invalid timezone';
            }
            return timeZone;
        }
    }


])
;

dtBookingApp.service('scheduleAvailability', ['$resource', '$log', '$window',
    function ($resource, $log, $window) {
        var _availabilityInfo = $resource($window.dtServiceEndpoint + '/booking/:serviceId/', {
            locationId: '@serviceId'
        }, {
            get: {
                method: 'GET',
                params: {}
            }
        });

        this.getAvailability = function (serviceId, date, success, errorHandler) {
            $log.log('getAvailability start ---> ' + serviceId + ":" + date);
            return _availabilityInfo.get({
                serviceId: serviceId,
                year: date.getFullYear(),
                month: date.getMonth() + 1
            }).$promise.then(success, errorHandler);
        };
    }
]);

dtBookingApp.service('appointmentCancelService', ['$resource', '$log', '$window',
    function ($resource, $log, $window) {
        var _appointmentCancel = $resource($window.dtServiceEndpoint + '/booking/hold/:guid', {}, {
            cancelHeldAppointment: {
                method: 'DELETE',
                params: {}
            }
        });

        this.cancelHeldAppointment = function (appointmentGuid, success, errorHandler) {
            $log.log('cancelHeldAppointment start --->');
            return _appointmentCancel.cancelHeldAppointment({
                guid: appointmentGuid
            }).$promise.then(success, errorHandler);
        };
    }
]);

dtBookingApp.service('finalizeBooking', ['$resource', '$window', 'locale',
    function ($resource, $window, locale) {

        var _finalizeBooking = $resource($window.dtServiceEndpoint + '/booking/complete', {}, {
            query: {
                method: 'POST',
                params: {},
                headers: { 'Accept-Language': locale.getLocaleStr() }
            }
        });

        var _completeReschedule = $resource($window.dtServiceEndpoint + '/booking/reschedule', {}, {
            query: {
                method: 'POST',
                params: {},
                headers: { 'Accept-Language': locale.getLocaleStr() }
            }
        });

        this.completeBooking = function (heldGuid, confirmNum, timestamp, licenceClass, endorsement, success, errorHandler) {
            return _finalizeBooking.query({
                appointmentHoldGuid: heldGuid,
                confirmationNumber: confirmNum,
                timestamp: timestamp,
                licenceClass: licenceClass,
                endorsement: endorsement
            }).$promise.then(success, errorHandler);
        };

        this.completeReschedule = function (existingAppointmentGuid, heldAppointmentGuid, confirmNum, timestamp, licenceClass, endorsement, success, errorHandler) {
            return _completeReschedule.query({
                existingAppointmentGuid: existingAppointmentGuid,
                booking: {
                    appointmentHoldGuid: heldAppointmentGuid,
                    confirmationNumber: confirmNum,
                    timestamp: timestamp,
                    licenceClass: licenceClass,
                    endorsement: endorsement
                }
            }).$promise.then(success, errorHandler);
        };
    }
]);

dtBookingApp.service('bookingFees', ['$resource', '$window', function ($resource, $window) {
    var _bookingFees = $resource($window.dtServiceEndpoint + '/booking/fees', {}, {
        calculate: {
            method: 'POST',
            params: {}
        }
    });

    this.calculate = function (user, success, errorHandler) {
        return _bookingFees.calculate({
            licenceClass: user.licenceClass,
            endorsement: user.endorsement,
            reschedule: user.reschedule,
            motorcycleType: user.component,
            existingAppointmentGuid: ((user.previousHold === undefined ? null : user.previousHold.confirmed.guid))
        }).$promise.then(success, errorHandler);
    };
}]);

dtBookingApp.service('prepareBookingPaymentService', ['$resource', '$window', 'locale', function ($resource, $window, locale) {
    var _preparePayment = $resource($window.dtServiceEndpoint + '/booking/payment/:guid',
        {}, { post: {method: 'POST', params: {}, headers: { 'Accept-Language': locale.getLocaleStr() }}}
    );

    this.preparePayment = function (appointmentGuid, bookingFeeRequestData, success, errorHandler) {
        bookingFeeRequestData.returnHost = window.location.protocol + "//" + window.location.host;
        return _preparePayment.post({guid: appointmentGuid}, bookingFeeRequestData).$promise.then(success, errorHandler);
    };
}]);

dtBookingApp.service('verifyPayment', ['$resource', '$window', function ($resource, $window) {
    var _verifyPayment = $resource($window.dtServiceEndpoint + '/booking/payment/verification', {}, {
        verify: {
            method: 'POST',
            params: {}
        }
    });

    this.verify = function (guid, id, confirmationNumber, success, errorHandler) {
        return _verifyPayment.verify({
            guid: guid,
            timestamp: id,
            confirmationNumber: confirmationNumber
        }).$promise.then(success, errorHandler);
    };
}]);

'use strict';

var dtBookingApp = angular.module('dtBooking');
dtBookingApp.service('loginSvc', ['$window', '$http', function ($window, $http) {
    this.login = function (user, captcha, successCallback, errorCallback) {
        var registrationData = {
            licenceNumber: user.licenceNumber,
            licenceExpiry: user.licenceExpiry,
            captchaResponse: captcha
        };
        return $http.post($window.dtServiceEndpoint + "/driver/token", registrationData).then(successCallback, errorCallback);
    };

    this.verify = function (user, successCallback, errorCallback) {
        var loginData = {
            licenceNumber: user.licenceNumber,
            licenceExpiry: user.licenceExpiry,
            emailToken: user.emailToken
        };
        return $http.post($window.dtServiceEndpoint + "/driver/email/token", loginData).then(successCallback, errorCallback);
    };

    //Not in use
    this.activate = function (user, successCallback, errorCallback) {
        var data = {
            email: user.email,
            emailToken: user.emailToken
        };
        return $http.post($window.dtServiceEndpoint + "/driver/emailToken", data).then(successCallback, errorCallback);
    };

    this.verifyEmailToken = function (emailToken, email, successCallback, errorCallback) {
        var emailTokenData = {
            email: email,
            emailToken: emailToken
        };
        return $http.post($window.dtServiceEndpoint + "/driver/emailToken", emailTokenData).then(successCallback, errorCallback);
    };
}]);
'use strict';

var dtBookingApp = angular.module('dtBooking');
dtBookingApp.service('registrationSvc', ['$window', '$http', '$resource', function ($window, $http, $resource) {
    this.register = function (user, captcha, successCallback, errorCallback) {
        var registrationData = {
            licenceNumber: user.licenceNumber,
            licenceExpiry: user.licenceExpiry,
            email: user.email,
            emailConfirm: user.emailConfirm,
            captchaResponse: captcha
        };

        return $http.post($window.dtServiceEndpoint + '/driver/email', registrationData).then(successCallback, errorCallback);
    };

    this.removeToken = function(licenceNumber, successCallback, errorCallback) {
        if(!angular.isDefined(licenceNumber)) {
            errorCallback(licenceNumber);
        } else {
            return $http({
                method: 'DELETE',
                url: $window.dtServiceEndpoint + '/driver/token',
                data: { licence: licenceNumber },
                headers: {'Content-Type': 'application/json;charset=utf-8'}
            }).then(successCallback, errorCallback);
        }
    };
}]);
'use strict';

var dtBookingApp = angular.module('dtBooking');

dtBookingApp.constant('States', {
        names: [
        'booking.login',
        'booking.dashboard',
        'booking.registration',
        'booking.verify',
        'booking.licence',
        'booking.location',
        'booking.calendar',
        'booking.timeslot',
        'booking.payment',
        'booking.finalize',
        'booking.complete',
        'booking.finish',
        'booking.error',
        'booking.cancel',
        'booking.success'
    ], toTop: [
        0,
        0,
        0,
        0,
        0,
        0,
        50,
        -145,
        0,
        0,
        0,
        0,
        0,
        0,
        0
    ]
});

dtBookingApp.service('stepState', ['$rootScope','$state','$log', '$window', 'userService', 'States', '$translate',
    function($rootScope, $state, $log, $window, userService, States, $translate){

    var flow;
    this.getFlow = function(){
        return  flow;
    }


    this.setFlow = function(flowin){
        flow = flowin;
    }
    var _currentStep = States.names.indexOf('booking.registration'),
        _lastStep = States.names.indexOf('booking.registration'),
        _this = this;

    this._currentInternalState = function() {
        return $state.current.name;
    };

    this.lastCompleteStepName = function() {
        return States.names[_lastStep];
    };

    this.currentStepName = function() {
        return States.names[_currentStep];
    };

    this.previousStepName = function() {
        return States.names[_currentStep>1?_currentStep-1:_currentStep];
    };
    this.previousPreviousStepName = function() {
            return States.names[_currentStep>2?_currentStep-2:(_currentStep>1?_currentStep-1:_currentStep)];
    };

    this.goToPreviousState = function() {
        _lastStep = _currentStep;
        _currentStep = _currentStep - 1;
        $state.go(States.names[_currentStep]);
    };

    this.goToNextState = function() {
        _lastStep = _currentStep;
        _currentStep = _currentStep + 1;
        $state.go(States.names[_currentStep]);
    };

    this.goToGeneralErrorState = function () {
        $state.go('booking.error');
    };

    this.checkStepState = function (state) {
        $log.log('checking state ---> ', $state);
        $log.log('checking _currentStep ---> ', _currentStep);
        this.updateHeader();

        var isInvalidSession = States.names.indexOf(state)>_currentStep;
        if (isInvalidSession) {
            $state.go(States.names[_currentStep]);
        }
        return isInvalidSession;
    };

    this.changeState = function (state,scroll,options) {
        if (scroll===undefined) {
            scroll = false;
        }
        _lastStep = _currentStep;
        _currentStep = States.names.indexOf(state);
        this.updateHeader();
        $state.go(state,undefined,options);

        $window.smoothScrollLink("#" + state.replace('.', '-'), scroll, States.toTop[_currentStep]);
    };

   this.updateHeader = function() {

       var stateArr = ['booking.login','booking.registration'],
           elQuitBtn = $('.quit-link'),
           elDashboardBtn = $('.dashboard-link');

       if(!$window._.contains(stateArr, _this._currentInternalState())){
           elQuitBtn.show();
       } else {
           elQuitBtn.hide();
       }

       if(!$window._.contains(stateArr, _this._currentInternalState()) && _this._currentInternalState() !== 'booking.dashboard'){
           elDashboardBtn.show();
       } else {
           elDashboardBtn.hide();
       }

       $translate("ng.text." + this._currentInternalState()).then(function (translatedBreadcrumb) {
           $window.setNGBreadcrumb(translatedBreadcrumb);
       });
    };

}]);
'use strict';

var dtBookingApp = angular.module('dtBooking');

dtBookingApp.value('user', {
        eligibilityRange: { "from": undefined, "to": undefined },
        firstName: undefined,
        lastName: undefined,
        licenceClass: undefined,
        licenceExpiry: undefined,
        licenceNumber: undefined,
        locationName: undefined,
        timezone: undefined,
        selectedDate: undefined,
        locationId: undefined,
        serviceId: undefined,
        eligible: false,
        email: undefined,
        emailConfirm: undefined,
        emailToken: undefined,
        endorsement: undefined,
        component: undefined,
        heldAppointment: undefined,
        reschedule: false,
        fees: undefined,
        rescheduleUser: undefined,
        appointments: undefined,
        checkSameLicense: true,
        bookingInProgress: false
});

dtBookingApp.service('userService', ['$window', '$resource', '$log', 'user', '$rootScope', 'registrationSvc', '$filter',
    function($window, $resource, $log, user, $rootScope, registrationSvc, $filter) {

    var _this = this;

    var _eligibility = $resource($window.dtServiceEndpoint + '/eligibilityCheck', {}, {
        check: {
            method: 'POST',
            params:{  }
        }
    });

    var _eligibleReschedule = $resource($window.dtServiceEndpoint + '/eligibilityCheck/reschedule', {}, {
        check: {
            method: 'POST',
            params: {}
        }
    });

    this.setUser = function (userIn) {
        user = userIn;
        $log.debug('storing user ---> ', user);
        $rootScope.$broadcast('user:updated',user);
    };

    this.getCurrentUser = function () {
        $log.debug('getting user ---> ', user);
        return user;
    };

    this.setSelectedDate = function(date) {
        user.selectedDate = date;
        user.timeslot = undefined;
        $rootScope.$broadcast('user:setSelectedDate',user);
    };

    this.isSet = function () {
        $log.debug('checking is set ---> ', user);
        return user.lastName !== undefined;
    };

    this.isEligible = function () {
        $log.debug('checking is set ---> ', user);
        return user.eligible;
    };

    this.getEligibility = function (success, errorHandler) {
        $log.debug('getEligibility start ---> ', user);
        _eligibility.check({licenceClass:user.licenceClass, endorsement: user.endorsement, motorcycleType: user.component, checkSameLicense: user.checkSameLicense}).$promise.then(success, errorHandler);
    };

    this.getRescheduleEligibility = function(appointment, success, errorHandler) {
        _eligibleReschedule.check({
            guid: appointment.appointmentId,
            licenceClass: appointment.licenceClass,
            endorsement: appointment.endorsement,
            motorcycleType: appointment.motorcycleType
        }).$promise.then(success, errorHandler);
    };

    this.createAppointmentHold = function(appointmentGuid) {
        $log.log("Holding appointment: " + appointmentGuid);
        user.heldAppointment = {
            guid: appointmentGuid,
            serviceId: user.serviceId,
            timeslot: user.timeslot
        };
    };

    this.clearHeldAppointment = function() {
        $log.log("Clearing held appointment: " + user.heldAppointment);
        user.heldAppointment = undefined;
    };

    this.appointmentHoldExists = function() {
        return user.heldAppointment !== undefined && user.heldAppointment.guid !== undefined;
    };

    this.isAppointmentSameAsCurrentlyHeld = function() {
        var heldAppointment = user.heldAppointment;
        if(heldAppointment !== undefined) {
            return heldAppointment.serviceId === user.serviceId && heldAppointment.timeslot === user.timeslot;
        }
        return false;
    };

    var _driverInformation =  $resource($window.dtServiceEndpoint + '/customer',
        {  }, {
            getName: {method: 'GET', params:{  }  }
        }
    );

    this.getCachedUserAppointments = function() {
        return user.appointments;
    };

    this.setUserAppointments = function(appointments) {
        user.appointments = appointments;
    };

    this.getDriverInformation = function(success, errorHandler) {
        _driverInformation.getName().$promise.then(success, errorHandler);
    };

    this.clearCurrentUserAppointment = function() {
        user.eligibilityRange.from = undefined;
        user.eligibilityRange.to = undefined;
        user.licenceClass = undefined;
        user.locationName = undefined;
        user.selectedDate = undefined;
        user.locationId = undefined;
        user.serviceId = undefined;
        user.eligible = false;
        user.endorsement = undefined;
        user.component = undefined;
        user.heldAppointment = undefined;
        user.reschedule = false;
        user.fees = undefined;
        user.rescheduleUser = undefined;
        user.appointments = undefined;
        user.checkSameLicense = true;
        user.previousHold = undefined;
    };
}]);
    'use strict';

var dtContactUsApp = angular.module('dtContactUs');

dtContactUsApp.controller('contactUsController', ['$scope', '$state', '$http', '$window', 'vcRecaptchaService', '$log',
function($scope, $state, $http, $window, recaptcha, $log) {
    $scope.contactUsForm = {};
    $scope.showCaptchaFailMsg = false;
    $scope.loading = false;
    $scope.captcha_response = null;

    $scope.setValidationFocus = function () {
        if(angular.isUndefined($scope.contactUsForm.firstName) || !$scope.contactUsForm.firstName.$valid) {
            $scope.$broadcast('focusFirstName');
            return false;
        } else if(angular.isUndefined($scope.contactUsForm.lastName) || !$scope.contactUsForm.lastName.$valid) {
            $scope.$broadcast('focusLastName');
            return false;
        } else if(angular.isUndefined($scope.contactUsForm.emailAddress) || !$scope.contactUsForm.emailAddress.$valid) {
            $scope.$broadcast('focusEmail');
            return false;
        } else if(angular.isUndefined($scope.contactUsForm.drivetestLocation) || !$scope.contactUsForm.drivetestLocation.$valid) {
            $scope.$broadcast('focusLocation');
            return false;
        } else if(angular.isUndefined($scope.contactUsForm.commentTopicsSelect) || !$scope.contactUsForm.commentTopicsSelect.$valid) {
            $scope.$broadcast('focusTopic');
            return false;
        } else if(angular.isUndefined($scope.contactUsForm.comments) || !$scope.contactUsForm.comments.$valid) {
            $scope.$broadcast('focusComments');
            return false;
        } else {
            return true;
        }
    };

    $scope.clearFieldValidation = function (element) {
        if(angular.isDefined($scope.contactUsForm[element]) && $scope.contactUsForm.form_submitted) {
            $scope.contactUsForm[element].$setPristine();
            $scope.contactUsForm[element].$error = {};
        }
    };

    $scope.contactUsFormSubmit = function(){
        $scope.loading = true;
        var fullName = $scope.contactUs.firstName.concat(" ", $scope.contactUs.lastName);
        var reply = true;
        if ($scope.contactUs.replyNeeded == undefined) {
            reply=false;
        }
        var data = JSON.stringify({
                email: $scope.contactUs.emailAddress,
                name: fullName,
                reply: reply,
                topic: $scope.data.commentTopicsSelect,
                body: $scope.contactUs.comments,
                drivetestCentre: $scope.data.commentDrivetestLocation,
                captchaResponse: $scope.captcha_response
        });
        $http.post($window.dtServiceEndpoint + '/contactus', data).success(function(result, status) {
            if (result.statusCode == '200') {
                $scope.showCaptchaFailMsg = false;
                $state.go('contact-us-success',{'success':'true'});
            } else if (result.statusCode == '500') {
                $scope.showCaptchaFailMsg = false;
                $state.go('contact-us-success',{'success':'false'});
            } else if (result.statusCode == '600') {
                $scope.showCaptchaFailMsg = true;
            } else {
                $scope.showCaptchaFailMsg = false;
                $state.go('contact-us-success',{'success':'false'});
            }
        }).error(function(){
            $scope.showCaptchaFailMsg = false;
            $scope.captcha_response = null;
            $scope.contactUsForm.form_submitted = false;
            $state.go('contact-us-success',{'success':'false'});
        });
    };

    $scope.setResponse = function(response) {
        $scope.captcha_response = response;
    };

}]);
'use strict';

var dtContactUsApp = angular.module('dtContactUs');

dtContactUsApp.controller('contactUsSuccessController', ['$scope', '$state', '$stateParams', '$http',
function($scope,$state,$stateParams,$http) {
    $scope.emailResultSuccess = true;
    if($stateParams.success == 'true') {
        $scope.resultMessage = 'Your message was sent successfully!';
        $scope.alertClass = 'alert-success';
    }else if($stateParams.success == 'false') {
        $scope.emailResultSuccess = false;
        $scope.resultMessage = 'Your message was not sent successfully, try again later';
        $scope.alertClass = 'alert-warning';
    }

    $scope.returnToForm = function(){
        $state.go('contact-us');
    };
}]);
'use strict';

angular.module('ContactUsForm.Comments', [])
    .directive('dtContactusComments', function() {
        return {
            estrict: 'AC',
            require: 'ngModel',
            link: function(scope, element, attr, ctrl) {
                ctrl.$validators.validateComments = function(value) {
                    if (value !== undefined) {
                        var ValidityCheck = /^[-'a-zA-ZÀ-ÖØ-öø-ſ ,.()?0-9\r\n]+$/.test(value);
                        ctrl.$setValidity('isNotValidComment', ValidityCheck);

                        var commentsOverCharacterLimit = true;
                        if (value.length>500) {
                            commentsOverCharacterLimit = false;
                        }
                        ctrl.$setValidity('commentTooLong', commentsOverCharacterLimit);
                    }

                    return value;
                };
            }
        };
    });
'use strict';

angular.module('ContactUsForm.Firstname', [])
    .directive('dtContactusFirstname', function() {
        return {
            restrict: 'AC',
            require: 'ngModel',
            link: function(scope, element, attr, ctrl) {
                ctrl.$validators.validateFirstName = function(value) {
                    if (value !== undefined) {
                        var ValidityCheck = /^[-'a-zA-ZÀ-ÖØ-öø-ſ ]+$/.test(value);
                        ctrl.$setValidity('isNotValidName', ValidityCheck);

                        if(value.length > 100) {
                            ctrl.$setValidity('nameTooLong', false);
                        } else {
                            ctrl.$setValidity('nameTooLong', true);
                        }
                    }

                    return value;
                };
            }
        };
    });
/**
 * Created by XuTong Zhu on 10/05/2018.
 */

'use strict';

angular.module('ContactUsForm.Focus', [])
    .directive('dtContactusFocus', function(){
        return {
            restrict: 'AC',
            require: 'ngModel',
            link: function (scope, elem, attrs) {
                scope.$on(attrs.dtContactusFocus, function(e) {
                    elem[0].focus();
                });
            }
        };
    });
'use strict';

angular.module('ContactUsForm.Lastname', [])
    .directive('dtContactusLastname', function() {
        return {
            restrict: 'AC',
            require: 'ngModel',
            link: function(scope, element, attr, ctrl) {
                ctrl.$validators.validateLastName = function(value) {
                    if (value !== undefined) {
                        var ValidityCheck = /^[-'a-zA-ZÀ-ÖØ-öø-ſ ]+$/.test(value);
                        ctrl.$setValidity('isNotValidName', ValidityCheck);

                        if(value.length > 100) {
                            ctrl.$setValidity('nameTooLong', false);
                        } else {
                            ctrl.$setValidity('nameTooLong', true);
                        }
                    }

                    return value;
                };
            }
        };
    });
/**
 * Created by Hisham on 2014-09-02.
 */
'use strict';

angular.module('dtMaps')
    .controller('mapController', ['$scope', '$log', '$rootScope', "leafletData", "leafletBoundsHelpers", "dtcLocationData",  function ($scope, $log, $rootScope, leafletData, leafletBoundsHelpers, dtcLocationData) {

        $scope.showFrench = false;
        $scope.openSaturday = false;
        $scope.findDTC = true;

        var southWestLat = 41.631538128024786;
        var southWestLng = -97.19560503959656;
        var northEastLat = 57.2;
        var northEastLng = -73.16411197185516;
        var maxbounds = leafletBoundsHelpers.createBoundsFromArray([
            [ southWestLat, southWestLng],
            [ northEastLat, northEastLng]
        ]);
        angular.extend($scope, {
            maxbounds: maxbounds,
            center: {
                lat: 45.653226,
                lng: -84.7494816,
                zoom: 5
            },
            tiles: {
                url: " "
            },
            defaults: {
                maxZoom: 16
            }
        });

        //Use local tile server in OpenStreetMap
        //Server URL is configurable in booking.properties. Use service call to get the URL
        var host = null;
        var getTileServerDataCallback = function (result) {
            host = result.MapTileServerHost;
            $scope.tiles = {
                url: host + "/{z}/{x}/{y}.png"
            };
        };

        var getTileServerDataErrorCallback = function (error) {
            $log.error('dtcLocationData.getTileServer ---> ', error);
        };

        dtcLocationData.getTileServer(getTileServerDataCallback, getTileServerDataErrorCallback);

        $scope.broadcastSaturdayOrFrench = function (openSaturday, showFrench) {
            $scope.showFrench = showFrench;
            $scope.openSaturday = openSaturday;
            $rootScope.$broadcast("refresh-map-saturday-french", openSaturday, showFrench);

            $log.debug('dtMaps.mapController - broadcastSaturdayOrFrench ---> ', $scope);

        };

        $scope.errorCallback = function (error) {
            $log.error('loaded with error ---> ', error);
        };

        //This function is called by directive 'ngDotcmsList' in dtc-map.jz
        $scope.mapCallback = function (location) {
            $log.log('Find a Center --> callback to maps directive');
            $scope.detailView = location;
            $scope.detailView.visible = true;
        };

    }]);
'use strict';

//This file provides the angular directives for the location controller(dt-locationController.js) and mapController(dt-c-map.js)
//There are five directives, the parent directive 'ngDtcLocations' and four child directives 'ngDtcMap', 'ngDtcList', 'ngDotcmsMap', and 'ngDotcmsList'.
//The communication between the directives uses $scope.$broadcast
//Directives 'ngDtcMap' and 'ngDtcList' are to implement the function for the locations and the map in the booking flow.
//Directives 'ngDotcmsMap' and 'ngDotcmsList' are to implement the functions for the location and the map in the 'Find Test Center' page.

angular.module('DriveTest.findACentre', ['leaflet-directive', 'DriveTest.services'])
    .constant('_', window._)
    .directive('ngDtcLocations', ['$log', '$rootScope', 'dtcLocations', 'dtcLocationData', 'dtcBlockLocationData', '$locale',
        function ($log, $rootScope, dtcLocations, dtcLocationData, dtcBlockLocationData, $locale) {

            function LocationController($scope, $element, $attrs) {
                this.markerOfInterest = function (location) {
                    $scope.$broadcast('mapZoom', location);
                };

                this.resetMap = function () {
                    //Ontario central point
                    var mapDefault = {
                        lat: 45.653226,
                        lng: -84.7494816,
                        zoom: 5
                    };
                    //Broadcast the resetMap action/event to all the child module 'ngDtcMap', 'ngDtcList', 'ngDotcmsMap', and 'ngDotcmsList'
                    $scope.$broadcast('resetMap', mapDefault);
                };

                this.getLanguageCode = function () {
                    return $locale.id;
                };

                this.convert24 = function (hour) {
                    var minutes = hour, ampm;

                    if ($locale.id === 'fr-ca') {
                        if (hour.match(':')) {
                            minutes = hour.split(':')[1];
                            hour = hour.split(':')[0] + 'h' + minutes;

                            return hour;
                        } else {
                            return hour;
                        }
                    } else {
                        if (hour.match($locale.DATETIME_FORMATS.AMPMS[1]) || hour.match($locale.DATETIME_FORMATS.AMPMS[0])) {
                            return hour;
                        } else {
                            if (hour > 12) {
                                ampm = $locale.DATETIME_FORMATS.AMPMS[0];
                            } else {
                                ampm = $locale.DATETIME_FORMATS.AMPMS[1];
                            }

                            hour = ((parseInt(hour, 10) + 11) % 12 + 1);
                            hour += ':' + minutes.split(':')[1] + ampm;

                            return hour;
                        }
                    }
                };
                this.locationScope = $scope;
            }
            return {
                scope: {
                    mapCallback: '=',
                    errorCallback: '=',
                    userInfo: '=ngUserInfo',
                    licenceType: '=ngFilterClassType',
                    endorsement: '=ngFilterEndorsement',
                    blockBook: '=ngBlockBook',
                    componentDisplay: '=ngComponentDisplay',
                    pullFrenchOnly: '=',
                    dotcms: '=ngDotcms'
                },
                controller: LocationController,
                restrict: 'AE',
                templateUrl: function (elem, attr) {
                    return attr.ngBlockBook !== undefined ?
                        '/ng/template-block-location-transclude' : '/ng/template-location-transclude';
                },
                replace: false,
                transclude: true,
                link: function ($scope, iElm, iAttrs, controller) {

                    $scope.loading = true;
                    $scope.currentLocale = controller.getLanguageCode();
                    $scope.broadcastSaturdayOrFrench = function (openSaturday, showFrench) {
                        controller.locationScope.showFrench = showFrench;
                        controller.locationScope.openSaturday = openSaturday;
                        controller.locationScope.$broadcast("refresh-map-saturday-french", openSaturday, showFrench);
                    };

                    var getDataCallback = function (result) {
                        $scope.dtcResultCount = result.length;
                        $scope.dtcResults = result;
                        $scope.loading = false;

                    };

                    var getDotcmsDataCallback = function (result) {
                        $scope.dotcmsResults = result;
                        $scope.loading = false;
                    };

                    if ($scope.blockBook) {
                        dtcBlockLocationData.getData($scope.pullFrenchOnly, getDataCallback, $scope.errorCallback);
                    } else {
                        if ($scope.dotcms) {
                            dtcLocationData.getDotcmsData(getDotcmsDataCallback, $scope.errorCallback);
                        } else {
                            dtcLocationData.getData($scope.licenceType, $scope.endorsement, $scope.userInfo, getDataCallback, $scope.errorCallback);
                        }
                    }
                }
            };
        }
    ])
    //'ngDtcMap' directive is used for the map functions in the booking flow
    .directive('ngDtcMap', ['$log', 'dtcLocations', 'dtcLocationData', '$rootScope', 'leafletMarkerEvents',
        function ($log, dtcLocations, dtcLocationData, $rootScope, $leafletMarkerEvents) {
            function mapController($scope, $element, $attrs) {
                $scope.center = {
                    lat: 45.653226,
                    lng: -84.7494816,
                    zoom: 5
                };
                $scope.leafletMarkers = [];

                $scope.icons = {
                    icon1: "/img/site/Map-Pin-Default@1x.png",
                    icon2: "/img/site/Map-Pin-hover@1x.png",
                    icon3: "/img/site/Map-Pin-ServiceDisruption@1x.png",
                    icon4: "/img/site/Map-Pin-ServiceDisruption-selectedhover@1x.png"
                };
            }
            //////////////////////////// ngDtcMap
            return {
                controller: mapController,
                require: '^ngDtcLocations',
                restrict: 'AE',
                templateUrl: '/ng/dtc-map',
                transclude: true,
                link: function ($scope, iElm, iAttrs, LocationController) {
                    var licenceType = LocationController.locationScope.licenceType;
                    if (LocationController.locationScope.endorsement !== undefined) {
                        licenceType = LocationController.locationScope.licenceType + LocationController.locationScope.endorsement;
                    }

                    $scope.$watch(function () {
                        return LocationController.locationScope.dtcResults;
                    }, function (newVal, oldVal) {
                        ///////////////////////// Leaflet event-capture framework start ////////////////////////////
                        //Capture the events on the markers
                        $scope.events = {
                            markers: {
                                enable: $leafletMarkerEvents.getAvailableEvents(),
                            }
                        };
                        var markerEvents = $leafletMarkerEvents.getAvailableEvents();
                        for (var k in markerEvents) {
                            var eventName = 'leafletDirectiveMarker.driveTestLeafletMap.' + markerEvents[k];
                            $scope.$on(eventName, function (event, args) {
                                var leafEvent = args.leafletEvent;
                                if (event.name.indexOf('.click') !== -1) {
                                    $scope.selectedMarkerLat = '' + leafEvent.latlng.lat;
                                    $scope.selectedMarkerLng = '' + leafEvent.latlng.lng;
                                    for (var i = 0; i < $scope.dtcLocationList.length; i++) {
                                        var location = $scope.dtcLocationList[i];
                                        location.latitude = '' + Number(location.latitude);
                                        location.longitude = '' + Number(location.longitude);
                                        if (location.latitude == $scope.selectedMarkerLat && location.longitude == $scope.selectedMarkerLng) {
                                            LocationController.markerOfInterest(location);

                                            var container = $(".dtc_listings");
                                            $("#" + location.id).click();
                                            container.scrollTop(
                                                $("#" + location.id).offset().top - container.offset().top +
                                                container.scrollTop() - 250
                                            );
                                            //////////////////////////// ngDtcMap
                                            if (angular.isDefined(LocationController.locationScope.mapCallback)) {
                                                $scope.userSelectedLocation = location;
                                                $scope.$parent.$parent.userSelectedLocationId = $scope.userSelectedLocationId = location.id;
                                                //test selectedEI can be removed
                                                LocationController.locationScope.mapCallback(location);
                                            }

                                        }
                                    }
                                }
                                ///////// ngDtcMap
                                if (event.name.indexOf('.mouseover') !== -1) {

                                    for (var i = 0; i < $scope.leafletMarkers.length; i++) {
                                        var marker = $scope.leafletMarkers[i];
                                        if (marker.lat == leafEvent.latlng.lat && marker.lng == leafEvent.latlng.lng) {
                                            //enable or disable the message pop-up
                                            //marker.focus = true;
                                            marker.icon = {
                                                iconUrl: $scope.icons.icon2
                                            };
                                        }
                                    }

                                }

                                if (event.name.indexOf('.mouseout') !== -1) {

                                    for (var i = 0; i < $scope.leafletMarkers.length; i++) {
                                        var marker = $scope.leafletMarkers[i];
                                        if (marker.lat == leafEvent.latlng.lat && marker.lng == leafEvent.latlng.lng) {
                                            //enable or disable the message pop-up
                                            //marker.focus = false;
                                            marker.icon = {
                                                iconUrl: $scope.icons.icon1
                                            };
                                        }
                                    }
                                }
                            });
                        }
                        ///////////////////////// Leaflet event-capture framework end ////////////////////////////
                        if (!angular.isUndefined(newVal)  &&  (angular.isUndefined($scope.leafletMarkers) ||  $scope.leafletMarkers.length == 0)) {
                            $scope.leafletMarkers = [];
                            angular.forEach(newVal, function (value, key) {
                                if(value.latitude != undefined && value.latitude != "" && !isNaN(value.latitude) &&
                                    value.longitude != undefined && value.longitude != "" && !isNaN(value.longitude)){
                                    $scope.leafletMarkers.push({
                                        lat: Number(value.latitude),
                                        lng: Number(value.longitude),
                                        //message: value.name,
                                        icon: {
                                            iconSize: [26,32],
                                            iconAnchor: [12,36],
                                            iconUrl: $scope.icons.icon1
                                        }
                                    });
                                }else{
                                    $log.debug('dtc-map.js  ngDtcMap found invalid lat/lng data->', value)
                                }
                            });
                        }
                        $log.debug('dtc-map.js  ngDtcMap  markers ->',  $scope.leafletMarkers);
                    }, true);
                    ///////// ngDtcMap
                    LocationController.locationScope.$on('mapZoom', function (event, data) {
                        var location = data;
                        var lat = location.latitude;
                        var lng = location.longitude;
                        if(lat == undefined || lat == "" || isNaN(lat) || lng == undefined || lng == "" || isNaN(lng)){
                            $log.debug('dtc-map.js  ngDtcMap mapZoom found invalid lat/lng data->', data);
                            return false;
                        }
                        $scope.center = {
                            lat: Number(lat),
                            lng: Number(lng),
                            zoom: 15
                        };

                        var oldMarkers = $scope.leafletMarkers;
                        $scope.leafletMarkers = [];
                        for (var i = 0; i < oldMarkers.length; i++) {
                            var marker = oldMarkers[i];
                            if (marker.lat == lat && marker.lng == lng) {
                                $scope.leafletMarkers.push({
                                    lat: Number(marker.lat),
                                    lng: Number(marker.lng),
                                    //message: marker.message,
                                    icon: {
                                        iconSize: [26,32],
                                        iconAnchor: [12,36],
                                        iconUrl: $scope.icons.icon2
                                    }
                                });
                                ///////// ngDtcMap
                            } else {
                                $scope.leafletMarkers.push({
                                    lat: Number(marker.lat),
                                    lng: Number(marker.lng),
                                    //message: marker.message,
                                    icon: {
                                        iconSize: [26,32],
                                        iconAnchor: [12,36],
                                        iconUrl: $scope.icons.icon1
                                    }
                                });
                            }
                        }
                        $log.debug('dtc-map.js  ngDtcMap mapZoom markers ->',  $scope.leafletMarkers);
                    });

                    LocationController.locationScope.$on('resetMap', function (event, data) {
                        $scope.center = data;
                    });

                    ///////// ngDtcMap
                    $scope.$watch(function () {
                        return $scope.initialized;
                    }, function (newVal, oldVal) {
                    }, true);

                    ////////////////////////////// ngDtcMap
                    LocationController.locationScope.$on('refresh-map-saturday-french',
                        function (event, openSaturday, showFrench, clearSelectedLocation) {

                            //de-select the location
                            $scope.userSelectedLocationId = undefined;
                            $scope.center = {
                                lat: 45.653226,
                                lng: -84.7494816,
                                zoom: 5
                            };
                            $scope.leafletMarkers = [];

                            angular.forEach($scope.dtcLocationList, function (value, key) {

                                if ((!showFrench || value.frenchAvailable) && (!openSaturday || value.openSaturday)) {
                                    if(value.latitude != undefined && value.latitude != "" && !isNaN(value.latitude) &&
                                        value.longitude != undefined && value.longitude != "" && !isNaN(value.longitude)){
                                        $scope.leafletMarkers.push({
                                            lat: Number(value.latitude),
                                            lng: Number(value.longitude),
                                            //message: value.name,
                                            icon: {
                                                iconSize: [26, 32],
                                                iconAnchor: [12, 36],
                                                iconUrl: $scope.icons.icon1
                                            }
                                        });
                                    }else{
                                        $log.debug('dtc-map.js  ngDtcMap  found invalid lat/lng data->', value);
                                    }
                                }
                            });
                            ////////////////////////////// ngDtcMap
                        }
                    );
                }
            };
        }
    ])
    //'ngDtcList' directive is used for the location functions in the booking flow
    .directive('ngDtcList', ['$log', 'dtcLocationData', '$window',
        function ($log, dtcLocationData, $window) {
            return {
                require: '^ngDtcLocations',
                restrict: 'AE',
                templateUrl: '/ng/dtc-list',
                transclude: true,
                link: function ($scope, iElm, iAttrs, LocationController) {
                    $scope.openDetails = {};
                    $scope.userSelectedLocation = {};
                    $scope.userSelectedLocationId = $scope.$parent.$parent.userSelectedLocationId || {};
                    $scope.locationScope = LocationController.locationScope;

                    $scope.$watch(function () {
                        return LocationController.locationScope.dtcResults;
                    }, function (newVal) {
                        if (!angular.isUndefined(newVal) || newVal !== null) {
                            $scope.dtcLocationList = newVal;
                            $log.debug(' $scope.dtcLocationList', $scope.dtcLocationList);
                        }
                    }, true);
                    ///////////////////////////////ngDtcList
                    $scope.detailRequest = function (evt, location) {
                        evt.preventDefault();
                        LocationController.markerOfInterest(location);
                        if (angular.isDefined(LocationController.locationScope.mapCallback)) {
                            $scope.userSelectedLocation = location;
                            $scope.$parent.$parent.userSelectedLocationId = $scope.userSelectedLocationId = location.id;
                            LocationController.locationScope.mapCallback(location);
                        }
                    };
                }
            };
        }
    ])
    //'ngDotcmsMap' directive is used for the map functions in the 'Find Test Center' page
    .directive('ngDotcmsMap', ['$log', 'dtcLocations', 'dtcLocationData', '$rootScope', 'leafletMarkerEvents',
        function ($log, $dtcLocations, $dtcLocationData, $rootScope, $leafletMarkerEvents) {
            // Runs during compile
            function mapController($scope, $element, $attrs) {

                $scope.initialized = {};

                $scope.center = {
                    lat: 45.653226,
                    lng: -84.7494816,
                    zoom: 5
                };
                $scope.leafletMarkers = [];
                $scope.icons = {
                    icon1: "/img/site/Map-Pin-Default@1x.png",
                    icon2: "/img/site/Map-Pin-hover@1x.png",
                    icon3: "/img/site/Map-Pin-ServiceDisruption@1x.png",
                    icon4: "/img/site/Map-Pin-ServiceDisruption-selectedhover@1x.png"
                };
            }
            return {
                ////////////////////////////////  'ngDotcmsMap'
                controller: mapController,
                require: '^ngDtcLocations',
                restrict: 'AE',
                templateUrl: '/ng/dtc-map',
                transclude: true,
                link: function ($scope, iElm, iAttrs, LocationController) {

                    $scope.selectedMarkerId = -100;
                    $scope.listenerStore = [];

                    $scope.$watch(function () {
                        return LocationController.locationScope.dotcmsResults;
                    }, function (newVal, oldVal) {

                        ///////////////////////// Leaflet event-capture framework start ////////////////////////////
                        //Capture the events on the markers
                        $scope.events = {
                            markers: {
                                enable: $leafletMarkerEvents.getAvailableEvents(),
                            }
                        };
                        ////////////////////////////////  'ngDotcmsMap'
                        var markerEvents = $leafletMarkerEvents.getAvailableEvents();
                        for (var k in markerEvents) {
                            var eventName = 'leafletDirectiveMarker.driveTestLeafletMap.' + markerEvents[k];
                            $scope.$on(eventName, function (event, args) {
                                var leafEvent = args.leafletEvent;
                                if (event.name.indexOf('.click') !== -1) {
                                    $scope.selectedMarkerLat = '' + leafEvent.latlng.lat;
                                    $scope.selectedMarkerLng = '' + leafEvent.latlng.lng;
                                    for (var i = 0; i < $scope.dtcLocationList.length; i++) {
                                        var location = $scope.dtcLocationList[i];

                                        location.latitude = '' + Number(location.latitude);
                                        location.longitude = '' + Number(location.longitude);
                                        if (location.latitude == $scope.selectedMarkerLat && location.longitude == $scope.selectedMarkerLng) {
                                            LocationController.markerOfInterest(location);

                                            //The following two lines do a 'click' action the location item in the location list
                                            var container = $(".dtc_listings");
                                            $("#" + location.id).click();
                                            //////////////////////////// 'ngDotcmsMap'
                                            if (angular.isDefined(LocationController.locationScope.mapCallback)) {
                                                $scope.userSelectedLocation = location;
                                                $scope.$parent.$parent.userSelectedLocationId = $scope.userSelectedLocationId = location.id;
                                                LocationController.locationScope.mapCallback(location);
                                            }
                                        }
                                    }
                                }
                                ///////// ngDtcMap
                                if (event.name.indexOf('.mouseover') !== -1) {

                                    for (var i = 0; i < $scope.leafletMarkers.length; i++) {
                                        var marker = $scope.leafletMarkers[i];
                                        if (marker.lat == leafEvent.latlng.lat && marker.lng == leafEvent.latlng.lng) {
                                            //enable or disable the message pop-up
                                            //marker.focus = true;
                                            marker.icon = {
                                                iconUrl: $scope.icons.icon2
                                            };
                                        }
                                    }
                                }

                                if (event.name.indexOf('.mouseout') !== -1) {

                                    for (var i = 0; i < $scope.leafletMarkers.length; i++) {
                                        var marker = $scope.leafletMarkers[i];
                                        if (marker.lat == leafEvent.latlng.lat && marker.lng == leafEvent.latlng.lng) {
                                            //enable or disable the message pop-up
                                            //marker.focus = false;
                                            marker.icon = {
                                                iconUrl: $scope.icons.icon1
                                            };
                                        }
                                    }
                                }
                            });
                        }
                        ///////////////////////// Leaflet event-capture framework end ////////////////////////////

                        if (!angular.isUndefined(newVal) &&  (angular.isUndefined($scope.leafletMarkers) ||  $scope.leafletMarkers.length == 0)) {
                            $scope.leafletMarkers = [];
                            angular.forEach(newVal, function (value, key) {
                                if(value.latitude != undefined && value.latitude != "" && !isNaN(value.latitude) &&
                                    value.longitude != undefined && value.longitude != "" && !isNaN(value.longitude)){
                                    $scope.leafletMarkers.push({
                                        lat: Number(value.latitude),
                                        lng: Number(value.longitude),
                                        //message: value.name,
                                        icon: {
                                            iconSize: [26, 32],
                                            iconAnchor: [12, 36],
                                            iconUrl: $scope.icons.icon1
                                        }
                                    });
                                }else{
                                    $log.debug('dtc-map.js  ngDotcmsMap  found invalid lat/lng data->', value);
                                }
                            });
                        }
                        $log.debug('dtc-map.js  ngDotcmsMap  markers ->', $scope.leafletMarkers);
                    }, true);

                    ////////////////////////////////  'ngDotcmsMap'
                    LocationController.locationScope.$on('mapZoom', function (event, data) {
                        var location = data;
                        var lat = location.latitude,
                            lng = location.longitude;

                        if(lat == undefined || lat == "" || isNaN(lat) || lng == undefined || lng == "" || isNaN(lng)){
                            $log.debug('dtc-map.js  ngDotcmsMap mapZoom found invalid lat/lng data->', data);
                            return false;
                        }
                        $scope.center = {
                            lat: Number(lat),
                            lng: Number(lng),
                            zoom: 15
                        };

                        var oldMarkers = $scope.leafletMarkers;
                        $scope.leafletMarkers = [];
                        for (var i = 0; i < oldMarkers.length; i++) {
                            var marker = oldMarkers[i];
                            if (marker.lat == lat && marker.lng == lng) {
                                $scope.leafletMarkers.push({
                                    lat: Number(marker.lat),
                                    lng: Number(marker.lng),
                                    //message: marker.message,
                                    icon: {
                                        iconSize: [26,32],
                                        iconAnchor: [12,36],
                                        iconUrl: $scope.icons.icon2
                                    }
                                });
                                ///////// ngDtcMap
                            } else {
                                $scope.leafletMarkers.push({
                                    lat: Number(marker.lat),
                                    lng: Number(marker.lng),
                                    //message: marker.message,
                                    icon: {
                                        iconSize: [26,32],
                                        iconAnchor: [12,36],
                                        iconUrl: $scope.icons.icon1
                                    }
                                });
                            }
                        }

                    });
                    ////////////////////////////////  'ngDotcmsMap'
                    LocationController.locationScope.$on('resetMap', function (event, data) {
                        $scope.center = {
                            lat: 45.653226,
                            lng: -84.7494816,
                            zoom: 5
                        };
                    });

                    $scope.$watch(function () {
                        return $scope.initialized;
                    }, function (newVal, oldVal) {
                    }, true);

                    ////////////////////////////////  'ngDotcmsMap'
                    LocationController.locationScope.$on('refresh-map-saturday-french',
                        function (event, openSaturday, showFrench, clearSelectedLocation) {

                            if (LocationController.locationScope.userInfo && clearSelectedLocation) {
                                LocationController.locationScope.userInfo.location = undefined;
                                $scope.userSelectedLocationId = undefined;
                            }
                            ////////////////////////////////  'ngDotcmsMap'
                            $scope.center = {
                                lat: 45.653226,
                                lng: -84.7494816,
                                zoom: 5
                            };
                            $scope.leafletMarkers = [];

                            angular.forEach($scope.dtcLocationList, function (value, key) {
                                if ((!showFrench || value.frenchavailable) && (!openSaturday || value.openSaturday)) {
                                    if(value.latitude != undefined && value.latitude != "" && !isNaN(value.latitude) &&
                                        value.longitude != undefined && value.longitude != "" && !isNaN(value.longitude)){
                                        $scope.leafletMarkers.push({
                                            lat: Number(value.latitude),
                                            lng: Number(value.longitude),
                                            //message: value.name,
                                            icon: {
                                                iconSize: [26, 32],
                                                iconAnchor: [12, 36],
                                                iconUrl: $scope.icons.icon1
                                            }
                                        });
                                    }else{
                                        $log.debug('dtc-map.js  ngDtcMap  found invalid lat/lng data->', value);
                                    }
                                }
                                ////////////////////////////// ngDtcMap test end
                            });
                            $log.debug('dtc-map.js  ngDotcmsMap  markers ->', $scope.leafletMarkers);
                        }
                    );
                }
            };
        }
    ])
    //'ngDotcmsList' directive is used for the location functions in the 'Find Test Center' page
    .directive('ngDotcmsList', ['$log', 'dtcLocationData', '$window',
        function ($log, dtcLocationData, $window) {
            return {
                require: '^ngDtcLocations',
                restrict: 'AE',
                templateUrl: '/ng/dotcms-list',
                transclude: true,
                link: function ($scope, iElm, iAttrs, LocationController) {
                    $scope.openDetails = {};
                    $scope.userSelectedLocation = {};
                    $scope.locationScope = LocationController.locationScope;
                    var mapLinkElement = document.getElementById("googlemapid");
                    var mapLinkElementText = mapLinkElement.innerHTML
                    if(mapLinkElementText.indexOf('Recevez des instructions pour vous rendre') != -1){
                        $scope.language_str = '?hl=fr';
                    }else{
                        $scope.language_str = '';
                    }

                    $scope.$watch(function () {
                        return LocationController.locationScope.dotcmsResults;
                    }, function (newVal) {
                        if (!angular.isUndefined(newVal) || newVal !== null) {
                            $scope.dtcLocationList = newVal;
                        }
                    }, true);

                    $scope.detailRequest = function (evt, location) {
                        evt.preventDefault();
                        LocationController.markerOfInterest(location);
                        if (angular.isDefined(LocationController.locationScope.mapCallback)) {
                            $scope.userSelectedLocation = location;
                            angular.forEach(location.locationhours, function (value, key) {
                                value.startTime = LocationController.convert24(value.startTime);
                                value.endTime = LocationController.convert24(value.endTime);
                            });
                            LocationController.locationScope.mapCallback(location);
                        }
                    };

                    $scope.closeDetails = function (evt) {
                        evt.preventDefault();
                        $scope.detailView.visible = false;
                        LocationController.locationScope.$broadcast('unselect-marker');
                        LocationController.resetMap();
                    };
                }
            };
        }
    ]);
"use strict";

angular.module('dtBooking')
    .run(['$httpBackend', '$log',
        function($httpBackend, $log) {

        		var eligibilityCheck = {
                    "eligible": true,
                    "from": "2014-10-01",
                    "to": "2022-12-31",
                },

                licenceClasses = {
                    "primaryDriverLicenceClasses": ["G2", "G", "M2", "M"],
                    "secondaryDriverLicenceClasses": ["D"]
                },

                locationData = {"driveTestCentres":[{"id":19,"name":"Aurora","licenceTestTypes":["G2","G","LM2","LM","M2","M"],"locationHours":[{"dayOFWeek":0,"startTime":"17:00","endTime":"17:00"},{"dayOFWeek":1,"startTime":"17:00","endTime":"17:00"},{"dayOFWeek":2,"startTime":"17:00","endTime":"17:00"},{"dayOFWeek":3,"startTime":"17:00","endTime":"17:00"},{"dayOFWeek":4,"startTime":"17:00","endTime":"17:00"},{"dayOFWeek":5,"startTime":"16:00","endTime":"16:00"}],"latitude":"43.9823244","longitude":"-79.4648307","address1":"1 Henderson Drive","address2":"Unit 4","city":"Aurora","province":"Ontario","postalCode":"L4G 4J7","services":[{"ID":9,"Name":"Aurora Services","LocationID":19,"MinuteDuration":15,"Active":true,"VariableDuration":false,"AllowPublic":true,"SelectionMethod":0,"DayReminder":false,"WeekReminder":false,"UsesWeeklySchedule":true,"TakesPayment":false}]},{"id":5,"name":"Bancroft","licenceTestTypes":["G2","G","LM2","LM","M2","M","A","B","C","D","E","F","Z"],"locationHours":[{"dayOFWeek":0,"startTime":"17:00","endTime":"17:00"},{"dayOFWeek":1,"startTime":"17:00","endTime":"17:00"},{"dayOFWeek":2,"startTime":"17:00","endTime":"17:00"},{"dayOFWeek":3,"startTime":"17:00","endTime":"17:00"},{"dayOFWeek":4,"startTime":"17:00","endTime":"17:00"}],"latitude":"45.0625872","longitude":"-77.8566692","address1":"142 Hastings St.N.","address2":"Unit 2","city":"Bancroft","province":"Ontario","postalCode":"K0L 1C0","services":[{"ID":10,"Name":"Bancroft Services","LocationID":5,"MinuteDuration":0,"Active":true,"VariableDuration":false,"AllowPublic":true,"SelectionMethod":0,"DayReminder":false,"WeekReminder":false,"UsesWeeklySchedule":true,"TakesPayment":false}]},{"id":42,"name":"Barrie","licenceTestTypes":["G2","G","LM2","LM","M2","M","A","B","C","D","E","F","Z","PerfTestServiceGroup 2509152031"],"locationHours":[{"dayOFWeek":0,"startTime":"17:00","endTime":"17:00"},{"dayOFWeek":1,"startTime":"17:00","endTime":"17:00"},{"dayOFWeek":2,"startTime":"17:00","endTime":"17:00"},{"dayOFWeek":3,"startTime":"17:00","endTime":"17:00"},{"dayOFWeek":4,"startTime":"17:00","endTime":"17:00"}],"latitude":"44.3302802","longitude":"-79.6886708","address1":"520 Bryne Drive","address2":"Unit 7","city":"Barrie","province":"Ontario","postalCode":"L4N 9P6","services":[{"ID":11,"Name":"Barrie Services","LocationID":42,"MinuteDuration":0,"Active":true,"VariableDuration":false,"AllowPublic":true,"SelectionMethod":0,"DayReminder":false,"WeekReminder":false,"UsesWeeklySchedule":true,"TakesPayment":false}]},{"id":6,"name":"Belleville","licenceTestTypes":["G2","G","LM2","LM","M2","M","A","B","C","D","E","F","Z","PerfTestServiceGroup 2409161036"],"locationHours":[{"dayOFWeek":0,"startTime":"17:00","endTime":"17:00"},{"dayOFWeek":1,"startTime":"17:00","endTime":"17:00"},{"dayOFWeek":2,"startTime":"17:00","endTime":"17:00"},{"dayOFWeek":3,"startTime":"17:00","endTime":"17:00"},{"dayOFWeek":4,"startTime":"17:00","endTime":"17:00"}],"latitude":"44.1852801","longitude":"-77.3689885","address1":"345 College Street E.","address2":"Unit 12, R.R. #6","city":"Belleville","province":"Ontario","postalCode":"K8N 5S7","services":[{"ID":12,"Name":"Belleville Services","LocationID":6,"MinuteDuration":0,"Active":true,"VariableDuration":false,"AllowPublic":true,"SelectionMethod":0,"DayReminder":false,"WeekReminder":false,"UsesWeeklySchedule":true,"TakesPayment":false}]},{"id":23,"name":"Brampton","licenceTestTypes":["G2","G","LM2","LM","M2","M","A","B","C","D","E","F","Z"],"locationHours":[{"dayOFWeek":0,"startTime":"17:00","endTime":"17:00"},{"dayOFWeek":1,"startTime":"17:00","endTime":"17:00"},{"dayOFWeek":2,"startTime":"17:00","endTime":"17:00"},{"dayOFWeek":3,"startTime":"17:00","endTime":"17:00"},{"dayOFWeek":4,"startTime":"17:00","endTime":"17:00"}],"latitude":"43.6784563","longitude":"-79.7145361","address1":"59 First Gulf Blvd","address2":"Unit #9","city":"Brampton","province":"Ontario","postalCode":"L6W 4P9","services":[{"ID":13,"Name":"Brampton Services","LocationID":23,"MinuteDuration":0,"Active":true,"VariableDuration":false,"AllowPublic":true,"SelectionMethod":0,"DayReminder":false,"WeekReminder":false,"UsesWeeklySchedule":true,"TakesPayment":false}]}]},

                bookingDates = {"availableBookingDates":[{"day":1,"code":4,"description":"Unavailable"},{"day":2,"code":4,"description":"Unavailable"},{"day":3,"code":4,"description":"Unavailable"},{"day":4,"code":4,"description":"Unavailable"},{"day":5,"code":4,"description":"Unavailable"},{"day":6,"code":1,"description":"Some"},{"day":7,"code":1,"description":"Some"},{"day":8,"code":1,"description":"Some"},{"day":9,"code":1,"description":"Some"},{"day":10,"code":1,"description":"Some"},{"day":11,"code":4,"description":"Unavailable"},{"day":12,"code":4,"description":"Unavailable"},{"day":13,"code":1,"description":"Some"},{"day":14,"code":1,"description":"Some"},{"day":15,"code":1,"description":"Some"},{"day":16,"code":2,"description":"Most"},{"day":17,"code":1,"description":"Some"},{"day":18,"code":4,"description":"Unavailable"},{"day":19,"code":4,"description":"Unavailable"},{"day":20,"code":1,"description":"Some"},{"day":21,"code":1,"description":"Some"},{"day":22,"code":1,"description":"Some"},{"day":23,"code":1,"description":"Some"},{"day":24,"code":0,"description":"Open"},{"day":25,"code":4,"description":"Unavailable"},{"day":26,"code":4,"description":"Unavailable"},{"day":27,"code":0,"description":"Open"},{"day":28,"code":1,"description":"Some"},{"day":29,"code":1,"description":"Some"},{"day":30,"code":1,"description":"Some"},{"day":31,"code":0,"description":"Open"}]},

                bookingTimes = {"availableBookingTimes":[{"timeslot":1418929200000},{"timeslot":1418931000000},{"timeslot":1418932800000},{"timeslot":1418934600000},{"timeslot":1418936400000},{"timeslot":1418938200000},{"timeslot":1418940000000},{"timeslot":1418941800000},{"timeslot":1418943600000},{"timeslot":1418945400000},{"timeslot":1418947200000},{"timeslot":1418949000000},{"timeslot":1418950800000},{"timeslot":1418952600000},{"timeslot":1418954400000}]},

                bookingHold = {
                    "success":true
                };

            if ($httpBackend && $httpBackend.whenGET) {

                $httpBackend.whenGET(/ng/).passThrough();

                $httpBackend.whenGET(/app\/api\/licenceClasses/).respond(function() {
                    $log.debug('/app/api/licenceClass/ -- >', licenceClasses);

                    return [200, licenceClasses];
                });

                $httpBackend.whenPOST(/app\/api\/eligibilityCheck/).respond(function() {
                    $log.debug('/app/api/eligibilityCheck/ -- >', eligibilityCheck);

                    return [200, eligibilityCheck];
                });

                $httpBackend.whenGET(/app\/api\/locations/).respond(function() {
                    $log.debug('/app/api/location/ -- >', locationData);

                    return [200, locationData];
                });

                $httpBackend.whenGET(/app\/api\/bookings/).respond(function() {
                    $log.debug('/app/api/booking/ -- >', bookingDates);

                    return [200, bookingDates];
                });

                $httpBackend.whenGET(/app\/api\/bookings\?date=/).respond(function() {
                    $log.debug('/app/api/booking?date= -- >', bookingTimes);

                    return [200, bookingTimes];
                });

                $httpBackend.whenPOST(/app\/api\/bookings\/hold/).respond(function() {
                    $log.debug('/app/api/booking/hold/ -- >', bookingHold);

                    return [200, bookingHold];
                });

            }
        }
    ]);