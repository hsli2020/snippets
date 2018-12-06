/**
 * Creates a script tag in the document head, the old fashioned way.
 * @param options
 * @private
 */
function addScriptTag(options) {

    var script = document.createElement('script');

    // Set the script properties.
    script.async = options.async || false;
    script.src = options.src;

    // Only actually do this if we are NOT in debug mode.
    if (!this._debug) {

        // Append the script element to the end of the document head.
        document.head.appendChild(script);
    }
}

// Add the script tag.
addScriptTag({
    async: true,
    src: '//www.googletagmanager.com/gtm.js?id=' + trackingCode
});


