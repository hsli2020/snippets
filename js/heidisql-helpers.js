function maillink(prefix, domainname, tld, subject, mcontent)
{
	var mehl = prefix+"@"+domainname+"."+tld;
	if(mcontent == '')
		mcontent = mehl;
	if(subject != '')
		subject = '?subject='+encodeURIComponent(subject);
	return '<a hr'+'ef="mai'+'lto:'+mehl+subject+'">'+mcontent+'</a>';
}

function formatText(btn, strBefore, strAfter, defaultText)
{
    var $box = $('#post_message');
    if($box.selection() == '') {
        $box.selection('replace', {text: defaultText});
    }
	$box.selection('insert', {text: strBefore, mode: 'before'});
	$box.selection('insert', {text: strAfter, mode: 'after'});
	$box.keyup();
}

function forumQuote(msgid, username, msgtext)
{
	var messagecontainer = document.getElementById('msg'+msgid);
	if(messagecontainer.textContent)
		message = messagecontainer.textContent;
	else
		message = messagecontainer.innerText;
	var quote = '> '+msgtext;
	quote = quote.replace(/(\r\n|\r|\n)/g, '$&> ');
	var postBox = $('#post_message');
    postBox.val(postBox.val() + quote);
    window.scrollTo(0, postBox.offset().top);
    postBox.keyup();
}

function addImage(btn)
{
	var default_url = 'http://www.example.com/sample';
	var url = prompt('URL to image file:', default_url);
	if(url != default_url)
		formatText(btn, '![', ']('+url+')', 'image description');
	$('#post_message').keyup();
}

function addWebLink(btn)
{
	var default_url = 'http://www.example.com/sample';
	var url = prompt('Web address:', default_url);
	if(url != default_url)
		formatText(btn, '[', ']('+url+')', 'link text');
	$('#post_message').keyup();
}

function addList(btn, usenumbers)
{
	var seltext = $.trim($('#post_message').selection());
	var lines = seltext.split("\n");
	var newtext = '';
	for(var i=0; i<lines.length; i++)
	{
		if(usenumbers) newtext += (i+1)+'. '; else newtext += '* ';
		newtext += lines[i] + "\n";
	}
	$('#post_message').selection('replace', {text: newtext});
	$('#post_message').keyup();
}

function addCode(btn)
{
    var $selText = $('#post_message').selection();
    var $hasLineBreaks = /[\r\n]/.test($selText);
    if($hasLineBreaks) {
        formatText(btn, '```\r\n', '\r\n```');
    }
    else {
        formatText(btn, '`', '`', 'SELECT');
    }
}

/**
 * Creates a debounced function that delays invoking `func` until after `wait`
 * milliseconds have elapsed since the last time the debounced function was
 * invoked. The debounced function comes with a `cancel` method to cancel
 * delayed `func` invocations and a `flush` method to immediately invoke them.
 * Provide `options` to indicate whether `func` should be invoked on the
 * leading and/or trailing edge of the `wait` timeout. The `func` is invoked
 * with the last arguments provided to the debounced function. Subsequent
 * calls to the debounced function return the result of the last `func`
 * invocation.
 *
 * **Note:** If `leading` and `trailing` options are `true`, `func` is
 * invoked on the trailing edge of the timeout only if the debounced function
 * is invoked more than once during the `wait` timeout.
 *
 * If `wait` is `0` and `leading` is `false`, `func` invocation is deferred
 * until to the next tick, similar to `setTimeout` with a timeout of `0`.
 *
 * See [David Corbacho's article](https://css-tricks.com/debouncing-throttling-explained-examples/)
 * for details over the differences between `_.debounce` and `_.throttle`.
 *
 * @static
 * @memberOf _
 * @since 0.1.0
 * @category Function
 * @param {Function} func The function to debounce.
 * @param {number} [wait=0] The number of milliseconds to delay.
 * @param {Object} [options={}] The options object.
 * @param {boolean} [options.leading=false]
 *  Specify invoking on the leading edge of the timeout.
 * @param {number} [options.maxWait]
 *  The maximum time `func` is allowed to be delayed before it's invoked.
 * @param {boolean} [options.trailing=true]
 *  Specify invoking on the trailing edge of the timeout.
 * @returns {Function} Returns the new debounced function.
 * @example
 *
 * // Avoid costly calculations while the window size is in flux.
 * jQuery(window).on('resize', _.debounce(calculateLayout, 150));
 *
 * // Invoke `sendMail` when clicked, debouncing subsequent calls.
 * jQuery(element).on('click', _.debounce(sendMail, 300, {
     *   'leading': true,
     *   'trailing': false
     * }));
 *
 * // Ensure `batchLog` is invoked once after 1 second of debounced calls.
 * var debounced = _.debounce(batchLog, 250, { 'maxWait': 1000 });
 * var source = new EventSource('/stream');
 * jQuery(source).on('message', debounced);
 *
 * // Cancel the trailing debounced invocation.
 * jQuery(window).on('popstate', debounced.cancel);
 */
function debounce(func, wait, options) {
    var lastArgs,
        lastThis,
        maxWait,
        result,
        timerId,
        lastCallTime,
        lastInvokeTime = 0,
        leading = false,
        maxing = false,
        trailing = true;

    if (typeof func != 'function') {
        throw new TypeError('error');
    }
    wait = wait || 0;
	leading = !!options.leading;
	maxing = 'maxWait' in options;
	maxWait = maxing ? Math.max(options.maxWait || 0, wait) : maxWait;
	trailing = 'trailing' in options ? !!options.trailing : trailing;

    function invokeFunc(time) {
        var args = lastArgs,
            thisArg = lastThis;

        lastArgs = lastThis = undefined;
        lastInvokeTime = time;
        result = func.apply(thisArg, args);
        return result;
    }

    function leadingEdge(time) {
        // Reset any `maxWait` timer.
        lastInvokeTime = time;
        // Start the timer for the trailing edge.
        timerId = setTimeout(timerExpired, wait);
        // Invoke the leading edge.
        return leading ? invokeFunc(time) : result;
    }

    function remainingWait(time) {
        var timeSinceLastCall = time - lastCallTime,
            timeSinceLastInvoke = time - lastInvokeTime,
            result = wait - timeSinceLastCall;

        return maxing ? Math.min(result, maxWait - timeSinceLastInvoke) : result;
    }

    function shouldInvoke(time) {
        var timeSinceLastCall = time - lastCallTime,
            timeSinceLastInvoke = time - lastInvokeTime;

        // Either this is the first call, activity has stopped and we're at the
        // trailing edge, the system time has gone backwards and we're treating
        // it as the trailing edge, or we've hit the `maxWait` limit.
        return (lastCallTime === undefined || (timeSinceLastCall >= wait) ||
            (timeSinceLastCall < 0) || (maxing && timeSinceLastInvoke >= maxWait));
    }

    function timerExpired() {
        var time = jQuery.now();
        if (shouldInvoke(time)) {
            return trailingEdge(time);
        }
        // Restart the timer.
        timerId = setTimeout(timerExpired, remainingWait(time));
    }

    function trailingEdge(time) {
        timerId = undefined;

        // Only invoke if we have `lastArgs` which means `func` has been
        // debounced at least once.
        if (trailing && lastArgs) {
            return invokeFunc(time);
        }
        lastArgs = lastThis = undefined;
        return result;
    }

    function cancel() {
        if (timerId !== undefined) {
            clearTimeout(timerId);
        }
        lastInvokeTime = 0;
        lastArgs = lastCallTime = lastThis = timerId = undefined;
    }

    function flush() {
        return timerId === undefined ? result : trailingEdge(jQuery.now());
    }

    function debounced() {
        var time = jQuery.now(),
            isInvoking = shouldInvoke(time);

        lastArgs = arguments;
        lastThis = this;
        lastCallTime = time;

        if (isInvoking) {
            if (timerId === undefined) {
                return leadingEdge(lastCallTime);
            }
            if (maxing) {
                // Handle invocations in a tight loop.
                timerId = setTimeout(timerExpired, wait);
                return invokeFunc(lastCallTime);
            }
        }
        if (timerId === undefined) {
            timerId = setTimeout(timerExpired, wait);
        }
        return result;
    }
    debounced.cancel = cancel;
    debounced.flush = flush;
    return debounced;
}

function scrollHandler() {

    // Lazy load images or contents of elements with .lazy-load class
    $.each($('.lazy-load'), function() {
        var $elem = $(this);

        var $doLoad = $elem.is(':visible') && isInViewport(this);
		if(!$doLoad) {
			return true; // false would stop the .each loop!
		}

		if($elem.is('img')) {
			// Move data-src to src attribute
			if($elem.data('src')) {
				$elem.attr('src', $elem.data('src'));
				$elem.removeAttr('data-src');
			}
		}


		// Performance: ensure we don't process this item on further scroll events again and again
		$elem.removeClass('lazy-load');
    });

}

/**
 * Check if DOM element is in browsers view port
 * @param el The element
 * @return boolean
 */
function isInViewport(el) {
    var rect = el.getBoundingClientRect();

    return (
        rect.bottom >= 0 &&
        rect.right >= 0 &&

        rect.top <= (
            window.innerHeight ||
            document.documentElement.clientHeight) &&

        rect.left <= (
            window.innerWidth ||
            document.documentElement.clientWidth)
    );
}

function postPreview()
{
	$.ajax({
		type: "POST",
		url: '/post_message_preview.php',
		data: $('#postform').serialize(), // 'post_message='+$('#post_message').val(),
		success: function(result) {
			$('#post_message_preview').toggle(result!='').html(result);
		},
		dataType: 'html'
		});
}

function handlePaste(e) {
    clipboardData = event.clipboardData || window.clipboardData || event.originalEvent.clipboardData;
    if (clipboardData == null) {
        return;
    }
    for (var i = 0 ; i < clipboardData.items.length ; i++) {
        var item = clipboardData.items[i];
        //console.log("Item type: " + item.type);
        if (item.type.indexOf("image") != -1) {
            uploadFile(item.getAsFile());
        } else {
            //console.log("Discarding non-image paste data");
        }
    }
}

function uploadFile(file) {
    var xhr = new XMLHttpRequest();

    /*xhr.upload.onprogress = function(e) {
        var percentComplete = (e.loaded / e.total) * 100;
        console.log("Uploaded: " + percentComplete + "%");
    };*/

    xhr.onload = function() {
        if (xhr.status == 200) {
            // Sucess! Upload completed
            var imageUrl = JSON.parse(xhr.responseText);
            //console.log(imageUrl);
            formatText(null, '![', ']('+imageUrl+')', 'Description');

        } else {
            alert("Image upload failed with status "+xhr.status+". Be sure you are logged in, and probably slow down your uploads.");
        }
    };

    xhr.onerror = function() {
        alert("Error! Upload failed. Can not connect to server.");
    };

    xhr.open("POST", "/post_message_upload.php", true);
    xhr.setRequestHeader("Content-Type", file.type);
    xhr.send(file);
}


function isMobile() {
    return window.matchMedia && window.matchMedia("(max-width: 1100px)").matches;
}

function fixedElementHeight() {
    var fixedEl = $('#header');
    if(fixedEl.length > 0) {
        if(fixedEl.css('position') == 'fixed')
            return fixedEl.outerHeight() + 20;
        else // not fixed on small screens
            return 0;
    }
    else
        return 140;
}

/*
$(window).on('hashchange', function() {
    goToAnchor();
});

$(window).on('load', function() {
    setTimeout(function() {
        goToAnchor();
    },1);
});
*/

function goToAnchor() {
    //console.log('goToAnchor fired');
    var $anchor = $(':target');

    if(!isMobile() && $anchor.length > 0) {
        window.scrollTo(0, $anchor.offset().top - fixedElementHeight());
    }
}


