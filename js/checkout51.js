/**
 * @see https://gist.github.com/cms/1147214
 */
if (typeof Function.prototype.bind != 'function') {
  Function.prototype.bind = (function () {
    var slice = Array.prototype.slice;
    return function (thisArg) {
      var target = this, boundArgs = slice.call(arguments, 1);

      if (typeof target != 'function') throw new TypeError();

      function bound() {
	var args = boundArgs.concat(slice.call(arguments));
	target.apply(this instanceof bound ? this : thisArg, args);
      }

      bound.prototype = (function F(proto) {
          proto && (F.prototype = proto);
          if (!(this instanceof F)) return new F;          
	})(target.prototype);

      return bound;
    };
  })();
}

$(document).ready(function() {
	var template = $('#offer-template')[0].innerHTML;
	var offers = C51.data.batch.offers;
	var i = 0;
	// @TODO rather than rolling my own templating, I should have implemented
	// handlebars or something... so, do that. - CAL
	function renderOffer(offer, init) {
		var element = document.createElement('div');
		var elements, i, data, dataAttribute, elementAttribute;
		element.innerHTML = template;
		$(element).addClass('offer');
		if (init) $(element).addClass('init');
		elements = $('[data-bind], [data-bind-src]', element);
		for (i = 0; i < elements.length; i++) {
			data = $(elements[i]).data();
			for (dataAttribute in data) {
				if (!dataAttribute.indexOf('bind') == 0) continue;
				if (dataAttribute != 'bind') {
					elementAttribute = dataAttribute.replace('bind', '').toLowerCase();
					$(elements[i]).attr(elementAttribute, offer[data[dataAttribute]]);
				} else {
				  if (data[dataAttribute] == 'cash_back' && $(elements[i]).attr('lang') == 'fr') {
				    $(elements[i]).append(offer[data[dataAttribute]].replace(/\./g, ",").replace(/(\d)(?=(\d\d\d)+(?!\d))/g, "$1 "));
				  } else{
				    $(elements[i]).append(offer[data[dataAttribute]]);
				  };
				}
			}
		}
		$('.ticker-inner').prepend(element);
		$('.ticker-inner').find('.offer:gt(6)').remove();
	}

	function getRandomOffer() {
		var i = Math.floor(Math.random() * offers.length);
		var offer = offers[i];

		// Exclude certain offers. This introduces the possibility
		// of an infinate loop that will cause a max call stack error
		// if all offers fail. Shouldn't ever happen in production.
		if (offer.cash_back == '0.00') return getRandomOffer();
		if (offer.remaining == 0) return getRandomOffer();
		if (offer.custom_content) return getRandomOffer();

		return offer;
	}

	// Fill in initial offers.
	for (i = 0; i < 6; i++) renderOffer(getRandomOffer(), true);

	setTimeout(function addOffer() {
		var timeout = Math.floor(Math.random() * 3000) + 1000;
		renderOffer(getRandomOffer())
		setTimeout(addOffer, timeout);
	}, 1000);
});

// Google analytics tracking.
$(document).ready(function() {
	$('.index .button.facebook').click(function() {
		_gaq.push(['_trackEvent', 'CTA_Buttons', 'click', 'Facebook Signup']);
	});
});

// .snap-to-footer control.
$(document).on('scroll touchmove', function(event) {
	var remainingScrollDistance = $(document).height() - window.innerHeight - $(document).scrollTop();
	var footerHeight = $('footer').height();
	var snap = $('.snap-to-footer');
	
	if (remainingScrollDistance < footerHeight) {
		$(snap).css('bottom', footerHeight);
		$(snap).css('position', 'absolute');
	} else {
		$(snap).removeAttr('style');
	}
});

$(document).ready(function() {$(document).scroll()});
$(window).resize(function() {$(document).scroll()});

C51 = window.C51 || {};
C51.data = window.C51.data || {};
C51.protocol = C51.protocol || {};
C51.protocol.regex = /c51:\/\/|\//g;
C51.protocol.listeners = {};

/**
 * Function registers a listener for the specified uri.
 * @param uri as String
 * @param callback as Function
 * @return Number as callback identifier.
 */
C51.protocol.addListener = function C51ProtocolAddListener(uri, callback) {
	C51.protocol.listeners[uri] = C51.protocol.listeners[uri] || [];
	C51.protocol.listeners[uri].push(callback);
	return C51.protocol.listeners[uri].length - 1;
}

/**
 * Function removes a listener for the specified uri.
 * @param uri as String
 * @param i as Number representing callback identifier
 * @return Boolean success
 */
C51.protocol.removeListener = function C51ProtocolRemoveListener(uri, i) {
	var success;

	C51.protocol.listeners[uri] = C51.protocol.listeners[uri] || [];

	if (success = C51.protocol.listeners[uri][i]) {
		C51.protocol.listeners[uri][i] = undefined;
	}

	return success;
}

/**
 * Function triggers all registered listeners for specified uri.
 * @param uri
 * @param thisArg
 * @param event
 */
C51.protocol.trigger = function C51ProtocolTrigger(uri, thisArg, event) {
	var i, arr;
	C51.protocol.listeners[uri] = C51.protocol.listeners[uri] || [];
	arr = C51.protocol.listeners[uri];
	for (i = 0; i < arr.length; i++) {
		if (typeof arr[i] == 'function') arr[i].call(thisArg, event);
	}
}

// Add listener for "back" URI.
C51.protocol.addListener('back', function() {
	if(window.location.pathname=="/account/offers") {
		//close all offer
 		$('.container .close').click();
	} else if(window.location.pathname.indexOf("/account/notifications/") > -1) {
		if (!window.location.origin) {
  			window.location.origin = window.location.protocol + "//" + window.location.hostname;
		}
		window.location.href = window.location.origin + '/account/notifications';
	} else {
		if (!window.location.origin) {
  			window.location.origin = window.location.protocol + "//" + window.location.hostname;
		}
		window.location.href = window.location.origin + '/offers';
	}
});
C51.protocol.addListener('offerlist', function() {
	if (!window.location.origin) {
  		window.location.origin = window.location.protocol + "//" + window.location.hostname;
	}
	window.location.href = window.location.origin + '/offers';
});
C51.protocol.addListener('myaccount', function() {
	if (!window.location.origin) {
  		window.location.origin = window.location.protocol + "//" + window.location.hostname;
	}
	window.location.href = window.location.origin + '/account/profile';
});
C51.protocol.addListener('offerdetails', function() {

	// requires offer_id
	if($(this).attr('c51action')) {
		var vars = deparam($(this).attr('c51action'));
	} else {
		var vars = deparam($(this).attr('href'));
	}
	
	if(typeof vars.offer_id === "undefined" || vars.offer_id === "") {
		location.hash = '#modal/whoops_later';
		return false;
	} else if(window.location.pathname=="/account/offers") {
		//open offer
		
		if($('li.offer[data-id="' + vars.offer_id + '"] .title').length<1) {
 			location.hash = '#modal/whoops_notavailable';
 		} else {
 			$('li.offer[data-id="' + vars.offer_id + '"] .title').click();
 		}
		
	} else {
		if(window.location.pathname.indexOf("/account/notifications/") > -1) {
			if (typeof batch === 'undefined') {
    			// variable is undefined
			} else {
				var found = 0;
				jQuery.each(batch.offers, function() {
					if($(this)[0].offer_id == vars.offer_id) {
						found = 1;
					}
				});
				if(found == 0) {
					location.hash = '#modal/whoops_notavailable';
					return false;
				}
			}
		}
		if (!window.location.origin) {
  			window.location.origin = window.location.protocol + "//" + window.location.hostname;
		}
		window.location.href = window.location.origin + '/account/offers?offer='+vars.offer_id;
	}
});
C51.protocol.addListener('updatedofferlist', function() {
	// go to offerlist

	if (!window.location.origin) {
		window.location.origin = window.location.protocol + "//" + window.location.hostname;
	}
	window.location.href = window.location.origin + '/offers';
	
});
C51.protocol.addListener('unlockoffer', function() {
	if($(this).attr('c51action')) {
		var vars = deparam($(this).attr('c51action'));
	} else {
		var vars = deparam($(this).attr('href'));
	}
	if(typeof vars.offer_id === "undefined" || typeof vars.offer_token === "undefined") {
		return false;
	} else {
		//unlock offer
		$.ajax({
			type: "POST",
			url: C51.data.API_URL + "/v1/unlockOffer",
			data: $.extend(vars, C51.data.authParams)
		}).done(function( msg ) {
			try {
				var result = JSON.parse(msg);
				if(result.status == "success") {
				/*if (!window.location.origin) {
  					window.location.origin = window.location.protocol + "//" + window.location.hostname;
				}
				window.location.href = window.location.origin + '/offers';*/
				//if (window.location.match(/\/notifications\/\d+/)) {
				  window.location.reload();
				//}
				}
				else {
				  location.hash = '#modal/whoops';
				}
			} catch(e) {
				location.hash = '#modal/whoops';
			}
		});
	}
});
C51.protocol.addListener('crmsubscribe', function() {
	var vars = deparam($(this).attr('c51action'));
	$.get('/account/ajax/crm_subscribe', vars, function(response) {
		if (response.subscription_id) {
			window.location.assign('/account/offers');
		} else {
			alert('There was an error joining the CRM program. Please email help@checkout51.com.');
		}
	});
});

$(document).ready(function() {
	var selector, iframes;

	// If there is a c51action_v2 link, make it c51action.
	$('a[c51action_v2]').each(function() {
		$(this).attr('c51action', $(this).attr('c51action_v2'));
	});

	// Attribute `href` starts with `c51://`.
	selector = 'a[c51action^="c51://"]';

	$(document.body).on('click', selector, {}, listener);

	selector = 'a[href^="c51://"]';
	
	

	$(document.body).on('click', selector, {}, listener);

	// Look for new iFrames, bind to them.
	setInterval(function() {
		var new_iframes = $('iframe:not([src])').not(iframes);

		new_iframes.each(function() {
			// If there is a c51action_v2 link, make it c51action.
			$('a[c51action_v2]', $(this).contents()).each(function() {
				$(this).attr('c51action', $(this).attr('c51action_v2'));
			});
		});

		selector = 'a[c51action^="c51://"]';
		$(new_iframes.contents()).on('click', selector, {}, listener);
		selector = 'a[href^="c51://"]';
		$(new_iframes.contents()).on('click', selector, {}, listener);
		iframes = $('iframe');
	}, 1000);

	function listener(event) {
		if(this.hasAttribute('c51action')) {
			
			var behaviour = this.getAttribute('c51action').replace(C51.protocol.regex, '').split('?')[0];
		} else {
			var behaviour = this.getAttribute('href').replace(C51.protocol.regex, '').split('?')[0];
		}
		
		C51.protocol.trigger(behaviour, this, event);
		event.preventDefault();
		event.stopPropagation();
	}
});


deparam = function (querystring) {
  // remove any preceding url and split
  querystring = querystring.substring(querystring.indexOf('?')+1).split('&');
  var params = {}, pair, d = decodeURIComponent, i;
  // march and parse
  for (i = querystring.length; i > 0;) {
    pair = querystring[--i].split('=');
    params[d(pair[0])] = d(pair[1]);
  }

  return params;
};

window.C51 = window.C51 || {};
C51.data = C51.data || {};
$(document).ready(function() {
	var notification_center = $('.notification-center');
	if (!notification_center.length) return;
	var notification_center_drop = $('header .notification-center').parent();
	var notification_center_hover = false;
	var expected_notificatsions_per_page = 15;
	var timestamp_notification_first, timestamp_notification_last;
	var uuid = Cookies.get('c51_uuid');
	var user_cookie = JSON.parse(Cookies.get(C51.data.WEB_USER_COOKIE_NAME));
	C51.data.authParams = {
		'user_id': user_cookie.user_id,
		'uuid': uuid,
		'authtoken': user_cookie.app_token,
		'platform': 'web',
		'token_version': 2
	}

	// If we end up using the Cookies JS library outside of this file, we should
	// move the defaults to a more sane location.
	Cookies.defaults = {
		domain: location.host,
		path: '/',
		expires: moment().add('hours', 12)._d
	};

	// Just personal preference. Using mustache/handlebars style instead of ruby.
	// If we end up using Underscore for templating outside of this file, we
	// should move the template settings to a more sane location.
	_.templateSettings = {
		evaluate: /\{\{#(.+?)\}\}/g,
		interpolate: /\{\{[^#\{](.+?)[^\}]\}\}/g,
		escape: /\{\{\{(.+?)\}\}\}/g
	}

	/**
	 * Compile templates.
	 */
	$(notification_center).each(function() {
		var tpl_source = $(' + .template', this).html();
		$(this).data('template', _.template(tpl_source));
	});

	// Initial notifications.
	getNotifications();

	/**
	 * Prevent the window from scrolling while the mouse is over the
	 * notifications dropdown. This avoids odd behaviour when trying to scroll
	 * throguh notifications in the dropdown.
	 */
	$(notification_center_drop).on('mouseover mouseout', function(e) {
		if (e.type == 'mouseover') {
			$('body').addClass('no-scroll');
		} else {
			$('body').removeClass('no-scroll');
		}
	});

	/**
	 * Clicking bell when notifications is open takes you
	 * to `/account/notifications`.
	 */
	$('header .notifications .drop-control').click(function() {
		var isOpen = !$('+ .drop', this).hasClass('hide');
		//if (isOpen) {
		//	location.assign('/account/notifications');
		//}
	});

	/**
	 * Mark notifications as read and hide unread notifications count.
	 */
	$('header .notifications.drop-wrap').click(function() {
		var params = C51.data.authParams,
		    drop = $('.drop', this)[0];

		// Kill the notification count.
		C51.data.notifications_unread_count = 0;
		renderNotifications();

		// Tell the server we've read everything.
		if (!C51.data.notifications) return;
		params.last_notif_read_ts = C51.data.notifications[0].timestamp;
		$.get(C51.data.API_URL + '/v1/markNotifsRead', params);
		$('body').removeClass('no-scroll');
	});

	/**
	 * When the user clicks a notification, remove it from our list of unread 
	 * notifications.
	 */
	$(notification_center).on('click', '.unread.notification', {}, function(e) {
		var notification = getNotificationFromElement(this);
		_.defer(markNotificationsRead.bind(this, [notification]));
	});

	/**
	 * Load older notifications when the user approaches the bottom of the
	 * notifications dropdown.
	 */
	$(notification_center).on('scroll', function(e) {
		var scroll_distance = e.target.clientHeight + e.target.scrollTop;
		if (scroll_distance >= e.target.scrollHeight) {
			getNotifications({
				before_notif_ts: timestamp_notification_first
			});
		}
	});

	/**
	 * Poll the server for new notifications periodically. 
	 */
	setTimeout(function poll() {
		var is_open = !$('header .notifications .drop').hasClass('hide');

		// Update in 30 seconds.
		setTimeout(poll, 30000);

		// Don't update the dropdown while the user is viewing it. 
		if (is_open) return;

		getNotifications({
			after_notif_ts: timestamp_notification_last
		});
	}, 30000);

	/**
	 * Load new notifications when the user clicks the "Load more" link at the
	 * bottom of the `/account/notifications` page.
	 */
	$('.get-more-notifications').click(function() {
		getNotifications({
			before_notif_ts: timestamp_notification_first
		});
		return false;
	});

	/**
	 * Google Analytics.
	 */
	$('header .notifications .drop-control').click(function() {
		var unread, open = !$(' + .drop', this).hasClass('hide');
		if (open) {
			_gaq.push(['_trackEvent', 'Notifications', 'click', 'NOTIF_ICON_2']);
		} else {
			unread = JSON.parse(Cookies.get('c51_notifications_unread') || '[]');
			_gaq.push(['_trackEvent', 'Notifications', 'click', 'NOTIF_ICON_1', unread.length]);
		}
	});
	$(notification_center).on('click', '.notification a', function(e) {
		var notification = getNotificationFromElement($(this).closest('.notification'));
		var inside_dropdown = $(this).closest('header').length ? true : false;
		if (inside_dropdown) {
			_gaq.push(['_trackEvent', 'Notifications', 'click', 'NOTIF_DROPDOWN_' + notification.template_id]);
		} else {
			_gaq.push(['_trackEvent', 'Notifications', 'click', 'NOTIF_LIST_' + notification.template_id]);
		}
	});
	$('header .notifications .see-all').click(function(e) {
		_gaq.push(['_trackEvent', 'Notifications', 'click', 'NOTIF_SEE_ALL_NOTIF']);
	});

	/**
	 * Function removes notifications from local Cookie that manages read/unread
	 * notifications states. Function also posts to the API to set latest
	 * notification in set as read.
	 * @param notifications as Array.
	 */
	function markNotificationsRead(notifications) {
		var unread, without_args, last_notification_read;

		// Set cookie.
		unread = JSON.parse(Cookies.get('c51_notifications_unread') || '[]');
		without_args = _(notifications).map(function(notification) { return notification.notif_id });
		without_args.unshift(unread);
		unread = _.without.apply(unread, without_args);
		Cookies.set('c51_notifications_unread', JSON.stringify(unread));

		// Mark notifications as read, and re-render.
		_(notifications).each(function(notification) {
			notification.unread = 0;
			notification.just_read = true;
		});
		renderNotifications();

		// Get the last notification read timestamp (from our group).
		last_notification_read = _(notifications).reduce(function(memo, value) {
			return memo.timestamp > value.timestamp ? memo : value;
		});

		params = C51.data.authParams;
		params.last_notif_read_ts = last_notification_read.timestamp;

		$.get(C51.data.API_URL + '/v1/markNotifsRead', params)
	}

	/**
	 * Function requests new notifications from server and handles business logic
	 * for their display.
	 * @param params as Object representing additional paramaters for
	 *   `/v1/getNotifs`.
	 * @param callback as Function accepting paramater `response`.
	 */
	function getNotifications(params, callback) {
		var params = _(params || {}).extend(C51.data.authParams);

		$.get(C51.data.API_URL + '/v1/getNotifs', params, function(response) {
			//this needs to exist
			if(!response.notifs)
				return;

			var unread = JSON.parse(Cookies.get('c51_notifications_unread') || '[]');
			var first, last;

			C51.data.notifications_unread_count = parseInt(response.unread_notif_count);

			$(notification_center).removeClass('hide');

			if (response.notifs.length < expected_notificatsions_per_page) {
				$('.main .get-more-notifications').addClass('hide');
			} else {
				$('.main .get-more-notifications.hide').removeClass('hide');
			}

			if (!response.notifs.length) return;
			$('.notifications .see-all').removeClass('hide');

			// In the case that the count of all of the users' notifications is
			// equal to a multiple of our page length, the user would see a 
			// `Load More` button, that when clicked, didn't load any more. By
			// removing one notification from a complete page, we ensure that if the
			// `Load More button is visible, it will function as the user expects.
			// The notification that is removed from each page will be loaded as the
			// first result in the next page. 
			if (response.notifs.length == expected_notificatsions_per_page) {
				response.notifs.pop();
			}

			// Notifications returned from server in reverse chronological order.
			last = parseInt(_(response.notifs).first().timestamp);
			first = parseInt(_(response.notifs).last().timestamp);
			if (first < (timestamp_notification_first || Infinity)) {
				timestamp_notification_first = first;
			}
			if (last > (timestamp_notification_last || -Infinity)) {
				timestamp_notification_last = last;
			}

			// Sanitize.
			_(response.notifs).each(function(notification) {
				notification.just_read = false;
				if (!isNotificationRead(notification)) {
					// The server says the notification is unread. This notification is
					// probably new, and hasn't been entered into our Cookie yet. We
					// need to store it in our cookie because the server only stores the
					// latest read cookie, not whether each one has been read.
					if (unread.indexOf(notification.notif_id) == -1) {
						unread.push(notification.notif_id);
					}
				}
			});

			// Probably not the fastest way to sort this, but there's not that much
			// data here. I'm benchmarking at ~1ms with 100+ notifications. - cal
			C51.data.notifications = _.chain(C51.data.notifications || [])
				.union(response.notifs)
				.uniq(function(notification) {
					return notification.notif_id;
				})
				.sortBy(function(notification) {
					return -notification.timestamp;
				})
				.value();

			Cookies.set('c51_notifications_unread', JSON.stringify(unread));

			renderNotifications();

			if (callback) callback.call(this, response);
		});
	}

	/**
	 * Function controls display of notifications and related elements (such
	 * as notifications count and see all button).
	 */
	function renderNotifications() {
		var unread = JSON.parse(Cookies.get('c51_notifications_unread') || '[]');
		var unread_count = C51.data.notifications_unread_count < 99 ? C51.data.notifications_unread_count : '99+';
		var header_notification_center = $('header nav .notification-center');
		var element_count, element_seeAll;

		if (!C51.data.notifications) return;

		$(notification_center).each(function() {
			var notifications = '';
			var center = this;
			_.each(C51.data.notifications, function(notification) {
				notifications += $(center).data('template')(notification);
			});
			$(this).html(notifications);
		});

		_.each(C51.data.notifications, function(notification) {
			notification.just_read = false;
		});

		element_count = $('.notification-count');
		element_seeAll = $('.notifications .see-all');

		if (C51.data.notifications_unread_count) {
			$(element_count).html(unread_count);
			$(element_count).removeClass('dismissed');
			$(element_count).removeClass('hide');
		} else {
			$(element_count).addClass('dismissed');
		}
	}

	/**
	 * Function checks local Cookie to determin if a given notification should
	 * be considered read.
	 * @param notification
	 */
	function isNotificationRead(notification) {
		var unread = JSON.parse(Cookies.get('c51_notifications_unread') || '[]');
		var local_unread = _(unread).contains(notification.notif_id);
		notification.unread = parseInt(notification.unread);

		if (local_unread) {
			// Browser thinks the notification is unread, it's probably right.
			notification.unread = true;
		}

		return !notification.unread;
	}

	/**
	 * Function finds matches an HTML element to it's associated notification
	 * based on data attribute `notification-id`.
	 * @param element as HTMLElement.
	 * @return notification.
	 */
	function getNotificationFromElement(element) {
		return _(C51.data.notifications).findWhere({
			notif_id: $(element).data('notificationId').toString()
		});
	}
});

// Special cases for header dropdown.
$(document).ready(function() {
	// Open or close dropdown on hover.
  $('header .drop-wrap').on('mouseenter mouseleave', function(event) {
		var dropIsOpen = !$('.drop', this).hasClass('hide');
		var shouldOpenDrop = event.type == 'mouseover' && !dropIsOpen;
		var shouldCloseDrop = event.type == 'mouseout' && dropIsOpen;
		if (shouldOpenDrop || shouldCloseDrop) $(this).click();
  });
});

// Country toggle.
$(document).ready(function() {
	var mql;
	var country_toggle = $('header .locale-toggle .locale-country');
	// Assign `mql` in supporting browsers (IE9+, FF, Chrome, Safari).
	if (window.matchMedia) mql = window.matchMedia('(max-width: 768px)');
	
	if (mql && mql.matches) {
	  $('header .locale-toggle .drop').width($('header nav').width());
	  $('header .locale-toggle .drop-curtain').width($('header nav').width());
	}

	/**
	 * Manages switching of country cookie and refreshes browser on change.
	 */
	$('.country a', country_toggle).on('click', function(e) {
		// Set the country cookie to it's new value.
		Cookies.set(C51.data.country_cookie_key, $(this).data('country'), {
			expires: moment().add('days', 30)._d,
			path: '/'
		});

		// Reload the browser, without the `country` paramater (allowing the users'
		// selected country to take effect).
		location.search = location.search.replace(/[?&]country=[^&;]*/,'');

		e.stopPropagation();
	});

	/**
	 * Stops clicks on the country toggle from propagating to the document,
	 * which would close them instantaneously.
	 */
	$('header nav').click(function(e) {
		e.stopPropagation();
	});

	if (!mql || !mql.matches) {
    $('header nav .locale-toggle').hoverIntent(function(){
  	  $('header nav .locale-toggle .drop-control').click();
  	});
 }
});

// Control stylized dropdowns.
$(document).ready(function() {
	$('html').click(function() {
		$('.drop').addClass('hide');
	});
	$('.drop-wrap').on('click', function(event) {
		var drop = $('.drop', this);
		var transitionInProgress = $(drop).hasClass('transition-in-progress');
		if (transitionInProgress) return false;
		$('.drop').not(drop).addClass('hide');
		$(drop).toggleClass('hide');
		$(drop).addClass('transition-in-progress');
		setTimeout(function() {$(drop).removeClass('transition-in-progress')}, 500);
		event.stopPropagation();
		
	});
});

<script type="application/javascript">
	C51 = window.C51 || {};
	C51.data = C51.data || {};
	C51.data.pingdom_id = '5166dc01e6e53d7711000000';

	var _prum={id: C51.data.pingdom_id};
	var PRUM_EPISODES=PRUM_EPISODES||{};
	PRUM_EPISODES.q=[];
	PRUM_EPISODES.mark=function(b,a){PRUM_EPISODES.q.push(["mark",b,a||new Date().getTime()])};
	PRUM_EPISODES.measure=function(b,a,b){PRUM_EPISODES.q.push(["measure",b,a,b||new Date().getTime()])};
	PRUM_EPISODES.done=function(a){PRUM_EPISODES.q.push(["done",a])};
	PRUM_EPISODES.mark("firstbyte");
	(function(){
		var b=document.getElementsByTagName("script")[0];
		var a=document.createElement("script");
		a.type="text/javascript";
		a.async=true;
		a.charset="UTF-8";
		a.src="//rum-static.pingdom.net/prum.min.js";
		b.parentNode.insertBefore(a,b);
	})();

</script>

<script type="text/javascript">
	C51 = window.C51 || {};
	C51.data = C51.data || {};
	C51.data.country_cookie_key = 'c51_production_country';
</script>
