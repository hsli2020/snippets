function addCSSRule(selector, rule) {
    var isMsie = (/msie/.test(navigator.userAgent.toLowerCase()));

    var styleTag = document.createElement("style");
    styleTag.setAttribute("type", "text/css");
    styleTag.setAttribute("id", "hideResults");

    if (!isMsie) {
        styleTag.appendChild(document.createTextNode(selector + " {" + rule + "}"));
    }

    document.getElementsByTagName("head")[0].appendChild(styleTag);

    if (isMsie && document.styleSheets && document.styleSheets.length > 0) {
        var lastStyleTag = document.styleSheets[document.styleSheets.length - 1];
        var selectors = selector.split(',');
        for(var i = 0; i < selectors.length; i++){
            lastStyleTag.addRule(selectors[i], rule);
        }
    }
};
