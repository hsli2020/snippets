// ==UserScript==
// @name         Fuck X
// @namespace    http://tampermonkey.net/
// @version      0.1
// @description  Save your eyes while using X (formerly Twitter)
// @author       Cyandev <unixzii@gmail.com>
// @match        https://twitter.com/*
// @grant        none
// ==/UserScript==

(function() {
    'use strict';

    const mo = new MutationObserver(handleDOMMutations);
    let pendingMods = [];

    function addMod(selector, fn) {
        pendingMods.push({ selector, fn });
    }

    function runPendingMods() {
        pendingMods = pendingMods.filter(m => {
            const el = document.querySelector(m.selector);
            if (!el) {
                // The element is not ready, keep it in the queue.
                return true;
            }
            console.debug(`running mod for \`${m.selector}\``);
            try {
                m.fn(el);
            } catch {
                // An error occurred, maybe we should retry later.
                return true;
            }
            return false;
        });

        return pendingMods.length > 0;
    }

    function handleDOMMutations(records) {
        if (!records.some(r => !!r.addedNodes.length)) {
            return;
        }
        if (!runPendingMods()) {
            this.disconnect();
        }
    }

    addMod('header[role="banner"] h1>a>div>svg', (el) => {
        const replacement = '<g><path d="M23.643 ... 2.323-2.41z"/></g>';
        el.innerHTML = replacement;
    });
    addMod('link[rel="shortcut icon"]', (el) => {
        const origFaviconUri = '//abs.twimg.com/favicons/twitter.ico';
        el.href = origFaviconUri;
    });

    mo.observe(document, {childList: true, subtree: true });
})();


// ==UserScript==
// @name         Make bluebird great again
// @namespace    TypeNANA
// @version      0.6
// @description  Twitter's bluebird is better than 𝕏
// @author       You
// @match        *://*.twitter.com/*
// @match        *://twitter.com/*
// @icon         https://www.google.com/s2/favicons?sz=64&domain=twitter.com
// @require      http://code.jquery.com/jquery-3.3.1.min.js
// @grant        none
// @license      MIT
// ==/UserScript==
 
(function() {
    waitForKeyElements('a[aria-label="Twitter"]', makeBlueBirdGreatAgain);
    waitForKeyElements('div[data-testid="TopNavBar"]', makeBlueBirdGreatAgain);

    changeFavIcon();
 
    function changeFavIcon(){
        let v = $('link[rel="shortcut icon"]');
        if (v == null || v.length == 0) return;
        v[0].setAttribute("href", 'https://abs.twimg.com/favicons/twitter.ico');
    }
 
    function makeBlueBirdGreatAgain(v){
        if (v == null || v.length == 0) return;
        v[0].getElementsByTagName("path")[0].setAttribute("d", 'M23.643 ... 2.323-2.41z')
    }
 
    function waitForKeyElements(selectorTxt, actionFunction, bWaitOnce, iframeSelector) {
        var targetNodes, btargetsFound;
 
        if (typeof iframeSelector == "undefined") {
            targetNodes = $(selectorTxt);
        } else {
            targetNodes = $(iframeSelector).contents().find(selectorTxt);
        }
 
        if (targetNodes && targetNodes.length > 0) {
            btargetsFound = true;
            targetNodes.each(function () {
                var jThis = $(this);
                var alreadyFound = jThis.data('alreadyFound') || false;
 
                if (!alreadyFound) {
                    var cancelFound = actionFunction(jThis);
                    if (cancelFound) {
                        btargetsFound = false;
                    } else {
                        jThis.data('alreadyFound', true);
                    }
                }
            });
        } else {
            btargetsFound = false;
        }
 
        var controlObj = waitForKeyElements.controlObj || {};
        var controlKey = selectorTxt.replace(/[^\w]/g, "_");
        var timeControl = controlObj[controlKey];
 
        if (btargetsFound && bWaitOnce && timeControl) {
            clearInterval(timeControl);
            delete controlObj[controlKey]
        } else {
            if (!timeControl) {
                timeControl = setInterval(function () {
                    waitForKeyElements(selectorTxt, actionFunction, bWaitOnce, iframeSelector);
                }, 300);
                controlObj[controlKey] = timeControl;
            }
        }
        waitForKeyElements.controlObj = controlObj;
    }
})();

https://greasyfork.org/en/scripts/471742-fix-x-button/code

// ==UserScript==
// @name         Fix X button
// @namespace    https://twitter.com/
// @version      1.0
// @description  Fixes the faulty new X button
// @author       You
// @match        https://twitter.com/
// @match        https://*.twitter.com/
// @match        https://twitter.com/*
// @match        https://*.twitter.com/*
// @icon         https://www.google.com/s2/favicons?sz=64&domain=twitter.com
// @grant        window.close
// @license      JSON
// ==/UserScript==
 
(function() {
    'use strict';
 
    const closeTab = () => {
        // in case window.close doesn't work, we need to have a fallback
        document.head.innerHTML = "<style>body { color: orange; background: black }</style>"
        document.body.innerHTML = "It's now safe to turn off your computer."
 
        // close the tab
        window.close()
    }
 
    const fixSvgElement = (v) => {
        if(v.innerHTML.includes("M18.244 ... 4.126H5.117z")) { // filter to the X button
            v.style.cursor = "pointer";
            v.onclick = closeTab
        }
    }
    const observer = new MutationObserver(mutations => {
        Array.from(document.getElementsByTagName("svg")).forEach(fixSvgElement);
    });
 
    observer.observe(document, {
        childList: true,
        subtree: true
    });
 
    Array.from(document.getElementsByTagName("svg")).forEach(fixSvgElement);
})();

