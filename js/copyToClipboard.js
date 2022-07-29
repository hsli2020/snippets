function copyToClipboard(text) {
    var textarea = document.createElement("textarea");
    textarea.textContent = text;
    textarea.style.position = "fixed";
    document.body.appendChild(textarea);
    textarea.select();
    try {
        return document.execCommand("cut");
    } catch (ex) {
        console.warn("Copy to clipboard failed.", ex);
        return false;
    } finally {
        document.body.removeChild(textarea);
    }
}

function htmlDecode(input) {
    var doc = new DOMParser().parseFromString(input, "text/html");
    return doc.documentElement.textContent;
}

function disableSelection(e) {
    if (typeof e.onselectstart !== "undefined")
        e.onselectstart = function () {
            show_tooltip();
            return false
        };
    else if (typeof e.style.MozUserSelect !== "undefined")
        e.style.MozUserSelect = "none";
    else e.onmousedown = function () {
        show_tooltip();
        return false
    };
    e.style.cursor = "default"
}

function stopPrntScr() {
    var inpFld = document.createElement("input");
    inpFld.setAttribute("value", "Access Denied");
    inpFld.setAttribute("width", "0");
    inpFld.style.height = "0px";
    inpFld.style.width = "0px";
    inpFld.style.border = "0px";
    document.body.appendChild(inpFld);
    inpFld.select();
    document.execCommand("copy");
    inpFld.remove(inpFld);
}

