// DOM UTILITIES

function elt(tag, content, className, style) {
    var e = document.createElement(tag);
    if (className) e.className = className;
    if (style) e.style.cssText = style;
    if (typeof content == "string") setTextContent(e, content);
    else if (content) for (var i = 0; i < content.length; ++i) e.appendChild(content[i]);
    return e;
}

function removeChildren(e) {
    for (var count = e.childNodes.length; count > 0; --count)
        e.removeChild(e.firstChild);
    return e;
}

function removeChildrenAndAdd(parent, e) {
    return removeChildren(parent).appendChild(e);
}

function setTextContent(e, str) {
    if (ie_lt9) {
        e.innerHTML = "";
        e.appendChild(document.createTextNode(str));
    } else e.textContent = str;
}

function getRect(node) {
    return node.getBoundingClientRect();
}

/* The right word count in respect for CJK. */
function wordCount(data) {
	var pattern = /[a-zA-Z0-9_\u0392-\u03c9]+|[\u4E00-\u9FFF\u3400-\u4dbf\uf900-\ufaff\u3040-\u309f\uac00-\ud7af]+/g;
	var m = data.match(pattern);
	var count = 0;
	if (typeof m == "string") {
		for (var i = 0; i < m.length; i++) {
			if (m[i].charCodeAt(0) >= 0x4E00) {
				if (typeof m[i] == "string") {
					count += m[i].length;
				} else {
					count += 1;
				}
			} else {
				count += 1;
			}
		}
	}
	return count;
}

function lineCount(data) {
	return data.split(/\r\n|\r|\n/).length;
}

/**
 * Toggle full screen of the editor.
 */
function toggleFullScreen(editor) {
	var el = editor.codemirror.getWrapperElement();

	// https://developer.mozilla.org/en-US/docs/DOM/Using_fullscreen_mode
	var doc = document;
	var isFull = doc.fullScreen || doc.mozFullScreen || doc.webkitFullScreen;
	var request = function() {
		if (el.requestFullScreen) {
			el.requestFullScreen();
		} else if (el.mozRequestFullScreen) {
			el.mozRequestFullScreen();
		} else if (el.webkitRequestFullScreen) {
			el.webkitRequestFullScreen(Element.ALLOW_KEYBOARD_INPUT);
		}
	};
	var cancel = function() {
		if (doc.cancelFullScreen) {
			doc.cancelFullScreen();
		} else if (doc.mozCancelFullScreen) {
			doc.mozCancelFullScreen();
		} else if (doc.webkitCancelFullScreen) {
			doc.webkitCancelFullScreen();
		}
	};
	if (!isFull) {
		request();
	} else if (cancel) {
		cancel();
	}
}

function createSep() {
	el = document.createElement('i');
	el.className = 'separator';
	el.innerHTML = '|';
	return el;
}

var isMac = /Mac/.test(navigator.platform);

/**
 * Fix shortcut. Mac use Command, others use Ctrl.
 */
function fixShortcut(name) {
	if (isMac) {
		name = name.replace('Ctrl', 'Cmd');
	} else {
		name = name.replace('Cmd', 'Ctrl');
	}
	return name;
}


/**
 * Create icon element for toolbar.
 */
function createIcon(name, options) {
	options = options || {};
	var el = document.createElement('a');

	var shortcut = options.shortcut || shortcuts[name];
	if (shortcut) {
		shortcut = fixShortcut(shortcut);
		el.title = shortcut;
		el.title = el.title.replace('Cmd', 'âŒ˜');
		if (isMac) {
			el.title = el.title.replace('Alt', 'âŒ¥');
		}
	}

	el.className = options.className || 'icon-' + name;
	return el;
}

var keyNames = {
    3: "Enter", 
    8: "Backspace", 
    9: "Tab", 
    13: "Enter", 
    16: "Shift", 
    17: "Ctrl", 
    18: "Alt",
    19: "Pause", 
    20: "CapsLock", 
    27: "Esc", 
    32: "Space", 
    33: "PageUp", 
    34: "PageDown", 
    35: "End",
    36: "Home", 
    37: "Left", 
    38: "Up", 
    39: "Right", 
    40: "Down", 
    44: "PrintScrn", 
    45: "Insert",
    46: "Delete", 
    59: ";", 
    91: "Mod", 
    92: "Mod", 
    93: "Mod", 
    109: "-", 
    107: "=", 
    127: "Delete",
    186: ";", 
    187: "=", 
    188: ",", 
    189: "-", 
    190: ".", 
    191: "/", 
    192: "`", 
    219: "[", 
    220: "\\",
    221: "]", 
    222: "'", 
    63276: "PageUp", 
    63277: "PageDown", 
    63275: "End", 
    63273: "Home",
    63234: "Left", 
    63232: "Up", 
    63235: "Right", 
    63233: "Down", 
    63302: "Insert", 
    63272: "Delete"
};

(function() {
    // Number keys
    for (var i = 0; i < 10; i++) keyNames[i + 48] = String(i);
    // Alphabetic keys
    for (var i = 65; i <= 90; i++) keyNames[i] = String.fromCharCode(i);
    // Function keys
    for (var i = 1; i <= 12; i++) keyNames[i + 111] = keyNames[i + 63235] = "F" + i;
})();

var hasSelection = window.getSelection ? function(te) {
    try { return te.selectionStart != te.selectionEnd; }
    catch(e) { return false; }
} : function(te) {
    try {var range = te.ownerDocument.selection.createRange();}
    catch(e) {}
    if (!range || range.parentElement() != te) return false;
    return range.compareEndPoints("StartToEnd", range) != 0;
};

var hasCopyEvent = (function() {
    var e = elt("div");
    if ("oncopy" in e) return true;
    e.setAttribute("oncopy", "return;");
    return typeof e.oncopy == 'function';
})();

// See if "".split is the broken IE version, if so, provide an
// alternative way to split lines.
var splitLines = "\n\nb".split(/\n/).length != 3 ? function(string) {
    var pos = 0, result = [], l = string.length;
    while (pos <= l) {
        var nl = string.indexOf("\n", pos);
        if (nl == -1) nl = string.length;
        var line = string.slice(pos, string.charAt(nl - 1) == "\r" ? nl - 1 : nl);
        var rt = line.indexOf("\r");
        if (rt != -1) {
            result.push(line.slice(0, rt));
            pos += rt + 1;
        } else {
            result.push(line);
            pos = nl + 1;
        }
    }
    return result;
} : function(string){return string.split(/\r\n?|\n/);};

var knownScrollbarWidth;
function scrollbarWidth(measure) {
    if (knownScrollbarWidth != null) return knownScrollbarWidth;
    var test = elt("div", null, null, "width: 50px; height: 50px; overflow-x: scroll");
    removeChildrenAndAdd(measure, test);
    if (test.offsetWidth)
        knownScrollbarWidth = test.offsetHeight - test.clientHeight;
    return knownScrollbarWidth || 0;
}

var zwspSupported;
function zeroWidthElement(measure) {
    if (zwspSupported == null) {
        var test = elt("span", "\u200b");
        removeChildrenAndAdd(measure, elt("span", [test, document.createTextNode("x")]));
        if (measure.firstChild.offsetHeight != 0)
            zwspSupported = test.offsetWidth <= 1 && test.offsetHeight > 2 && !ie_lt8;
    }
    if (zwspSupported) return elt("span", "\u200b");
    else return elt("span", "\u00a0", null, "display: inline-block; width: 1px; margin-right: -1px");
}

// For a reason I have yet to figure out, some browsers disallow
// word wrapping between certain characters *only* if a new inline
// element is started between them. This makes it hard to reliably
// measure the position of things, since that requires inserting an
// extra span. This terribly fragile set of tests matches the
// character combinations that suffer from this phenomenon on the
// various browsers.
function spanAffectsWrapping() { return false; }
if (gecko) // Only for "$'"
    spanAffectsWrapping = function(str, i) {
        return str.charCodeAt(i - 1) == 36 && str.charCodeAt(i) == 39;
    };
else if (safari && !/Version\/([6-9]|\d\d)\b/.test(navigator.userAgent))
    spanAffectsWrapping = function(str, i) {
        return /\-[^ \-?]|\?[^ !\'\"\),.\-\/:;\?\]\}]/.test(str.slice(i - 1, i + 1));
    };
else if (webkit && !/Chrome\/(?:29|[3-9]\d|\d\d\d)\./.test(navigator.userAgent))
    spanAffectsWrapping = function(str, i) {
        if (i > 1 && str.charCodeAt(i - 1) == 45) {
            if (/\w/.test(str.charAt(i - 2)) && /[^\-?\.]/.test(str.charAt(i))) return true;
            if (i > 2 && /[\d\.,]/.test(str.charAt(i - 2)) && /[\d\.,]/.test(str.charAt(i))) return false;
        }
        return /[~!#%&*)=+}\]|\"\.>,:;][({[<]|-[^\-?\.\u2010-\u201f\u2026]|\?[\w~`@#$%\^&*(_=+{[|><]|â€¦[\w~`@#$%\^&*(_=+{[><]/.test(str.slice(i - 1, i + 1));
    };

// Detect drag-and-drop
var dragAndDrop = function() {
    // There is *some* kind of drag-and-drop support in IE6-8, but I
    // couldn't get it to work yet.
    if (ie_lt9) return false;
    var div = elt('div');
    return "draggable" in div || "dragDrop" in div;
}();

function emptyArray(size) {
    for (var a = [], i = 0; i < size; ++i) a.push(undefined);
    return a;
}

function bind(f) {
    var args = Array.prototype.slice.call(arguments, 1);
    return function(){return f.apply(null, args);};
}

var nonASCIISingleCaseWordChar = /[\u3040-\u309f\u30a0-\u30ff\u3400-\u4db5\u4e00-\u9fcc\uac00-\ud7af]/;

function isWordChar(ch) {
    return /\w/.test(ch) || ch > "\x80" &&
        (ch.toUpperCase() != ch.toLowerCase() || nonASCIISingleCaseWordChar.test(ch));
}

function isEmpty(obj) {
    for (var n in obj) if (obj.hasOwnProperty(n) && obj[n]) return false;
    return true;
}

var isExtendingChar = /[\u0300-\u036F\u0483-\u0487\u0488-\u0489\u0591-\u05BD\u05BF\u05C1-\u05C2\u05C4-\u05C5\u05C7\u0610-\u061A\u064B-\u065F\u0670\u06D6-\u06DC\u06DF-\u06E4\u06E7-\u06E8\u06EA-\u06ED\uA66F\uA670-\uA672\uA674-\uA67D\uA69F\udc00-\udfff]/;

function createObj(base, props) {
    function Obj() {}
    Obj.prototype = base;
    var inst = new Obj();
    if (props) copyObj(props, inst);
    return inst;
}

function copyObj(obj, target) {
    if (!target) target = {};
    for (var prop in obj) if (obj.hasOwnProperty(prop)) target[prop] = obj[prop];
    return target;
}

var spaceStrs = [""];
function spaceStr(n) {
    while (spaceStrs.length <= n)
        spaceStrs.push(lst(spaceStrs) + " ");
    return spaceStrs[n];
}

function lst(arr) { return arr[arr.length-1]; }

function selectInput(node) {
    if (ios) { // Mobile Safari apparently has a bug where select() is broken.
        node.selectionStart = 0;
        node.selectionEnd = node.value.length;
    } else {
        // Suppress mysterious IE10 errors
        try { node.select(); }
        catch(_e) {}
    }
}

function indexOf(collection, elt) {
    if (collection.indexOf) return collection.indexOf(elt);
    for (var i = 0, e = collection.length; i < e; ++i)
        if (collection[i] == elt) return i;
    return -1;
}

// Counts the column offset in a string, taking tabs into account.
// Used mostly to find indentation.
function countColumn(string, end, tabSize, startIndex, startValue) {
    if (end == null) {
        end = string.search(/[^\s\u00a0]/);
        if (end == -1) end = string.length;
    }
    for (var i = startIndex || 0, n = startValue || 0; i < end; ++i) {
        if (string.charAt(i) == "\t") n += tabSize - (n % tabSize);
        else ++n;
    }
    return n;
}

function Delayed() {this.id = null;}
Delayed.prototype = {set: function(ms, f) {clearTimeout(this.id); this.id = setTimeout(f, ms);}};

function eventMixin(ctor) {
    ctor.prototype.on = function(type, f) {on(this, type, f);};
    ctor.prototype.off = function(type, f) {off(this, type, f);};
}

// EVENT HANDLING

function on(emitter, type, f) {
    if (emitter.addEventListener)
        emitter.addEventListener(type, f, false);
    else if (emitter.attachEvent)
        emitter.attachEvent("on" + type, f);
    else {
        var map = emitter._handlers || (emitter._handlers = {});
        var arr = map[type] || (map[type] = []);
        arr.push(f);
    }
}

function off(emitter, type, f) {
    if (emitter.removeEventListener)
        emitter.removeEventListener(type, f, false);
    else if (emitter.detachEvent)
        emitter.detachEvent("on" + type, f);
    else {
        var arr = emitter._handlers && emitter._handlers[type];
        if (!arr) return;
        for (var i = 0; i < arr.length; ++i)
            if (arr[i] == f) { arr.splice(i, 1); break; }
    }
}

function signal(emitter, type /*, values...*/) {
    var arr = emitter._handlers && emitter._handlers[type];
    if (!arr) return;
    var args = Array.prototype.slice.call(arguments, 2);
    for (var i = 0; i < arr.length; ++i) arr[i].apply(null, args);
}

var delayedCallbacks, delayedCallbackDepth = 0;
function signalLater(emitter, type /*, values...*/) {
    var arr = emitter._handlers && emitter._handlers[type];
    if (!arr) return;
    var args = Array.prototype.slice.call(arguments, 2);
    if (!delayedCallbacks) {
        ++delayedCallbackDepth;
        delayedCallbacks = [];
        setTimeout(fireDelayed, 0);
    }
    function bnd(f) {return function(){f.apply(null, args);};};
    for (var i = 0; i < arr.length; ++i)
        delayedCallbacks.push(bnd(arr[i]));
}

function signalDOMEvent(cm, e, override) {
    signal(cm, override || e.type, cm, e);
    return e_defaultPrevented(e) || e.codemirrorIgnore;
}

function fireDelayed() {
    --delayedCallbackDepth;
    var delayed = delayedCallbacks;
    delayedCallbacks = null;
    for (var i = 0; i < delayed.length; ++i) delayed[i]();
}

function hasHandler(emitter, type) {
    var arr = emitter._handlers && emitter._handlers[type];
    return arr && arr.length > 0;
}

function e_target(e) {return e.target || e.srcElement;}
function e_button(e) {
    var b = e.which;
    if (b == null) {
        if (e.button & 1) b = 1;
        else if (e.button & 2) b = 3;
        else if (e.button & 4) b = 2;
    }
    if (mac && e.ctrlKey && b == 1) b = 3;
    return b;
}

// STRING STREAM

// Fed to the mode parsers, provides helper functions to make
// parsers more succinct.

// The character stream used by a mode's parser.
function StringStream(string, tabSize) {
    this.pos = this.start = 0;
    this.string = string;
    this.tabSize = tabSize || 8;
    this.lastColumnPos = this.lastColumnValue = 0;
}

StringStream.prototype = {
    eol: function() {return this.pos >= this.string.length;},
    sol: function() {return this.pos == 0;},
    peek: function() {return this.string.charAt(this.pos) || undefined;},
    next: function() {
        if (this.pos < this.string.length)
            return this.string.charAt(this.pos++);
    },
    eat: function(match) {
        var ch = this.string.charAt(this.pos);
        if (typeof match == "string") var ok = ch == match;
        else var ok = ch && (match.test ? match.test(ch) : match(ch));
        if (ok) {++this.pos; return ch;}
    },
    eatWhile: function(match) {
        var start = this.pos;
        while (this.eat(match)){}
        return this.pos > start;
    },
    eatSpace: function() {
        var start = this.pos;
        while (/[\s\u00a0]/.test(this.string.charAt(this.pos))) ++this.pos;
        return this.pos > start;
    },
    skipToEnd: function() {this.pos = this.string.length;},
    skipTo: function(ch) {
        var found = this.string.indexOf(ch, this.pos);
        if (found > -1) {this.pos = found; return true;}
    },
    backUp: function(n) {this.pos -= n;},
    column: function() {
        if (this.lastColumnPos < this.start) {
            this.lastColumnValue = countColumn(this.string, this.start, this.tabSize, this.lastColumnPos, this.lastColumnValue);
            this.lastColumnPos = this.start;
        }
        return this.lastColumnValue;
    },
    indentation: function() {return countColumn(this.string, null, this.tabSize);},
    match: function(pattern, consume, caseInsensitive) {
        if (typeof pattern == "string") {
            var cased = function(str) {return caseInsensitive ? str.toLowerCase() : str;};
            var substr = this.string.substr(this.pos, pattern.length);
            if (cased(substr) == cased(pattern)) {
                if (consume !== false) this.pos += pattern.length;
                return true;
            }
        } else {
            var match = this.string.slice(this.pos).match(pattern);
            if (match && match.index > 0) return null;
            if (match && consume !== false) this.pos += match[0].length;
            return match;
        }
    },
    current: function(){return this.string.slice(this.start, this.pos);}
};

// STANDARD KEYMAPS

var keyMap = CodeMirror.keyMap = {};
keyMap.basic = {
    "Left": "goCharLeft", "Right": "goCharRight", "Up": "goLineUp", "Down": "goLineDown",
    "End": "goLineEnd", "Home": "goLineStartSmart", "PageUp": "goPageUp", "PageDown": "goPageDown",
    "Delete": "delCharAfter", "Backspace": "delCharBefore", "Tab": "defaultTab", "Shift-Tab": "indentAuto",
    "Enter": "newlineAndIndent", "Insert": "toggleOverwrite"
};
// Note that the save and find-related commands aren't defined by
// default. Unknown commands are simply ignored.
keyMap.pcDefault = {
    "Ctrl-A": "selectAll", "Ctrl-D": "deleteLine", "Ctrl-Z": "undo", "Shift-Ctrl-Z": "redo", "Ctrl-Y": "redo",
    "Ctrl-Home": "goDocStart", "Alt-Up": "goDocStart", "Ctrl-End": "goDocEnd", "Ctrl-Down": "goDocEnd",
    "Ctrl-Left": "goGroupLeft", "Ctrl-Right": "goGroupRight", "Alt-Left": "goLineStart", "Alt-Right": "goLineEnd",
    "Ctrl-Backspace": "delGroupBefore", "Ctrl-Delete": "delGroupAfter", "Ctrl-S": "save", "Ctrl-F": "find",
    "Ctrl-G": "findNext", "Shift-Ctrl-G": "findPrev", "Shift-Ctrl-F": "replace", "Shift-Ctrl-R": "replaceAll",
    "Ctrl-[": "indentLess", "Ctrl-]": "indentMore",
    fallthrough: "basic"
};
keyMap.macDefault = {
    "Cmd-A": "selectAll", "Cmd-D": "deleteLine", "Cmd-Z": "undo", "Shift-Cmd-Z": "redo", "Cmd-Y": "redo",
    "Cmd-Up": "goDocStart", "Cmd-End": "goDocEnd", "Cmd-Down": "goDocEnd", "Alt-Left": "goGroupLeft",
    "Alt-Right": "goGroupRight", "Cmd-Left": "goLineStart", "Cmd-Right": "goLineEnd", "Alt-Backspace": "delGroupBefore",
    "Ctrl-Alt-Backspace": "delGroupAfter", "Alt-Delete": "delGroupAfter", "Cmd-S": "save", "Cmd-F": "find",
    "Cmd-G": "findNext", "Shift-Cmd-G": "findPrev", "Cmd-Alt-F": "replace", "Shift-Cmd-Alt-F": "replaceAll",
    "Cmd-[": "indentLess", "Cmd-]": "indentMore", "Cmd-Backspace": "delLineLeft",
    fallthrough: ["basic", "emacsy"]
};
keyMap["default"] = mac ? keyMap.macDefault : keyMap.pcDefault;
keyMap.emacsy = {
    "Ctrl-F": "goCharRight", "Ctrl-B": "goCharLeft", "Ctrl-P": "goLineUp", "Ctrl-N": "goLineDown",
    "Alt-F": "goWordRight", "Alt-B": "goWordLeft", "Ctrl-A": "goLineStart", "Ctrl-E": "goLineEnd",
    "Ctrl-V": "goPageDown", "Shift-Ctrl-V": "goPageUp", "Ctrl-D": "delCharAfter", "Ctrl-H": "delCharBefore",
    "Alt-D": "delWordAfter", "Alt-Backspace": "delWordBefore", "Ctrl-K": "killLine", "Ctrl-T": "transposeChars"
};

// KEYMAP DISPATCH

function getKeyMap(val) {
    if (typeof val == "string") return keyMap[val];
    else return val;
}

function lookupKey(name, maps, handle) {
    function lookup(map) {
        map = getKeyMap(map);
        var found = map[name];
        if (found === false) return "stop";
        if (found != null && handle(found)) return true;
        if (map.nofallthrough) return "stop";

        var fallthrough = map.fallthrough;
        if (fallthrough == null) return false;
        if (Object.prototype.toString.call(fallthrough) != "[object Array]")
            return lookup(fallthrough);
        for (var i = 0, e = fallthrough.length; i < e; ++i) {
            var done = lookup(fallthrough[i]);
            if (done) return done;
        }
        return false;
    }

    for (var i = 0; i < maps.length; ++i) {
        var done = lookup(maps[i]);
        if (done) return done != "stop";
    }
}
function isModifierKey(event) {
    var name = keyNames[event.keyCode];
    return name == "Ctrl" || name == "Alt" || name == "Shift" || name == "Mod";
}
function keyName(event, noShift) {
    if (opera && event.keyCode == 34 && event["char"]) return false;
    var name = keyNames[event.keyCode];
    if (name == null || event.altGraphKey) return false;
    if (event.altKey) name = "Alt-" + name;
    if (flipCtrlCmd ? event.metaKey : event.ctrlKey) name = "Ctrl-" + name;
    if (flipCtrlCmd ? event.ctrlKey : event.metaKey) name = "Cmd-" + name;
    if (!noShift && event.shiftKey) name = "Shift-" + name;
    return name;
}

