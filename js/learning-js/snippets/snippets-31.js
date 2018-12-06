var isMobile = {
    Android: function() {
        return navigator.userAgent.match(/Android/i);
    },
    BlackBerry: function() {
        return navigator.userAgent.match(/BlackBerry/i);
    },
    iOS: function() {
        return navigator.userAgent.match(/iPhone|iPad|iPod/i);
    },
    Opera: function() {
        return navigator.userAgent.match(/Opera Mini/i);
    },
    Windows: function() {
        return navigator.userAgent.match(/IEMobile/i);
    },
    any: function() {
        return (isMobile.Android() || isMobile.BlackBerry() || 
                isMobile.iOS() || isMobile.Opera() || isMobile.Windows());
    }
};

function getDevice(){        
    if ( $(window).width() > 1199 ) return 'desktop';
    if ( $(window).width() < 1200 ) return 'tablet';
    if ( $(window).width() < 768 ) return 'mobile';        
}

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

function isMobile() {
    return window.screen.availWidth<=768 || 
        /Android|iPhone|iPod/i.test(navigator.userAgent)
}
function isApp() {
    return navigator.userAgent.toLowerCase().indexOf("educationapp")>-1
}

document.documentElement.className+=isMobile()?" mobile":" pc",
document.documentElement.className+=isApp()?" app":" ",

!function(n){
    n.console=n.console || {
        log:function(){},
        debug:function(){},
        info:function(){},
        warn:function(){},
        error:function(){}
    },
    n.badjs=n.Badjs=function(){}
}(window);

/MSIE [6|7|8]/i.test(navigator.userAgent) && (
    Array.prototype.forEach || (Array.prototype.forEach=function(r,t) {
        var n,o;
        if (null==this) throw new TypeError(" this is null or not defined");
        var e=Object(this),
        i=e.length>>>0;
        if("[object Function]"!={}.toString.call(r))
            throw new TypeError(r+" is not a function");
        for(t&&(n=t),o=0;i>o;) {
            var a;
            o in e && (a=e[o], r.call(n,a,o,e)), o++
        }
    }),
    Array.prototype.indexOf || (Array.prototype.indexOf=function(r){
        if(null==this)return-1;
        for(var t=0,n=this.length;n>t;t++)
            if(this[t]==r)return t;
        return-1
    }),
    Array.isArray || (Array.isArray=function(r) { 
        return"[object Array]" === Object.prototype.toString.call(r)
    }),
    Date.now || (Date.now=function() {
        return(new Date).getTime()
    })
);

// Avoid `console` errors in browsers that lack a console.
(function() {
    var method;
    var noop = function noop() {};
    var methods = [
        'assert', 'clear', 'count', 'debug', 'dir', 'dirxml', 'error',
        'exception', 'group', 'groupCollapsed', 'groupEnd', 'info', 'log',
        'markTimeline', 'profile', 'profileEnd', 'table', 'time', 'timeEnd',
        'timeStamp', 'trace', 'warn'
    ];
    var length = methods.length;
    var console = (window.console = window.console || {});

    while (length--) {
        method = methods[length];

        // Only stub undefined methods.
        if (!console[method]) {
            console[method] = noop;
        }
    }
}());

var trace = function(a) {
	try {
		console.log(a);
	} catch(e) {}
};

if(!Array.indexOf){
	Array.prototype.indexOf = function(object){
		for(var i = 0; i < this.length; i++){
			if(this[i] == object){return i;break;}
		}
		return -1;
	};
}

/**
 * http://4dd.jp/ 
 * @author nki2 / http://nki2.jp/
 * @author okb
 * @revision 4
 */
function newClass(classObj, superClass) {
	if(!classObj) classObj = {};
	if(typeof classObj.__construct !== "function") classObj.__construct = function() {};
	var f = classObj.__construct;
	f.extend = function(classObj) { return newClass(classObj, this); }
	
	if(superClass) {
		for(var i in superClass.prototype) f.prototype[i] = superClass.prototype[i];
		classObj.__super = superClass.prototype;
	}
	
	for(var j in classObj) {
		if(superClass && typeof classObj[j] == "function") {
			f.prototype[j] = (function(func, superClass) {
				return function() {
					var tmpSuper = this.__super;
					this.__super = superClass.prototype;
					var result = func.apply(this, arguments);
					this.__super = tmpSuper;
					return result;
				};
			})(classObj[j], superClass);
		} else {
			f.prototype[j] = classObj[j];
		}
	}
	return f;
}

var Class = newClass({});

okb.form.TextFld = Class.extend({
    // ...
	__construct: function ($me) {
		this.__super.__construct.apply(this, arguments);
		var me = this;

		me.$ = $me;
		me.$doc = $(document);
		me.$input = $("input", me.$);
		me.$img = $('<img src="' + okb.form.EMPTY_SRC + '" class="space" />');
		me.$df = $('<span class="df"></span>')
		me.$ov = $('<span class="ov"></span>')
		me.$down = $('<span class="down"></span>')
		me.$label = $("p", me.$);

		me.name = me.$input.attr("name");

		me.$.prepend(me.$img)
		me.$.append(me.$df)
		me.$.append(me.$ov)
		me.$.append(me.$down)

		me.$ov.fadeTo(0, 0);
		me.$down.fadeTo(0, 0);

		me.$input.data("enhanced", me);
		me.$.data("enhanced", me);
	},
});

okb.form.CheckBox = okb.form.CheckRadioBase.extend({
	__construct: function ($me, option) {
		this.__super.__construct.apply(this, arguments);
	},

	_toggleCheck: function () {
		var me = this;
		me.$input.prop("checked", !(me.$input.prop("checked") == true));
		this.__super._toggleCheck.apply(this, arguments);
	}
});

okb.form.FormButton = Class.extend({
	__construct: function ($me) {
		var me = this;

		me.$ = $me;
		me.$input = $("input", me.$);

		//tabindex¤òÒý¤­¾@¤°
		me.$input.attr("tabindex", me.$.attr("tabindex"));
		me.$.removeAttr("tabindex");
	}
});

okb.form.Label = Class.extend({
	__construct: function ($me) {
		var me = this;

		me.$ = $me;
		me.$input = $("input", me.$);

		me.$.click(function (e) {
			me.$input.prop("checked", !me.$input.prop("checked"))
		})
		me.$input.click(function (e) {
			e.stopPropagation();
		})
	}
});

function loadScript(src, onload) {
    var e = document.createElement("script");
    if(onload) e.onload = onload;
    e.async = true;
    e.src = src;
    var h = document.getElementsByTagName("head")[0];
    h.appendChild(e);
}

function loadStyle(src, onload) {
    var e = document.createElement("link");
    if(onload) e.onload = onload;
    e.rel = "stylesheet";
    e.type = "text/css";
    e.href = src;
    var h = document.getElementsByTagName("head")[0];
    h.appendChild(e);
}

function getCookie(key) {
    var cookies = document.cookie;
    var idx = {};
    //cookies.split("; ").forEach(function(val) {
    //	var arr = val.split("=");
    //	idx[arr.shift()] = arr.join("=");
    //});
    var cookieArr = cookies.split("; ");
    for(var i = 0, len = cookieArr.length; i < len; i++) {
        var arr = cookieArr[i].split("=");
        idx[arr.shift()] = arr.join("=");
    }
    if(idx[key]) {
        return idx[key];
    } else {
        return null;
    }
}

function setCookie(key, value, path, validityms) {
    var arr = [(key + "=" + value)];
    if(path != undefined) arr.push("path=" + path);
    if(validityms != undefined) 
        arr.push("expires=" + (new Date((new Date()).getTime() + validityms).toUTCString()));
    document.cookie = arr.join("; ");
}
	
function genId(len) {
    len = len || 8;
    var pool = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
    var plen = pool.length;
    var str = "";
    for(var i = 0; i < len; i++) str += pool.charAt((Math.random() * plen * 1000) % plen);
    return str;
}

function genCid() { return genId(32); }
	
