!function(f) {
    function $() {
        return f(window, document, navigator, undefined, setTimeout, encodeURIComponent, decodeURIComponent)
    }
    if ('object' == typeof exports) module.exports = $()
    else if ('function' == typeof define && define.amd) define($)
    else window.$ = $()
}(function(window, document, navigator, undefined, setTimeout, encodeURIComponent, decodeURIComponent) {

var $ = function(selector, context) {
    return new $.fn.init(selector, context)
}

// var
var rtrim = /^[\s\uFEFF\xA0]+|[\s\uFEFF\xA0]+$/g
    , rvalidtokens = /(,)|(\[|{)|(}|])|"(?:[^"\\\r\n]|\\["\\\/bfnrt]|\\u[\da-fA-F]{4})*"\s*:?|true|false|null|-?(?!0\d)\d+(?:\.\d+|)(?:[eE][+-]?\d+|)/g
    , slice = [].slice
    , getExpando = function() {
        return ('miniJQ' + Math.random()).replace(/\W/g, '')
    }
    , classes = 'boolean number string function array date regexp object error'.split(' ')
    , isLetter = function(str) {
        if (str && 'string' == typeof str) {
            str = str.toUpperCase()
            for (var i = 0; i < str.length; i++) {
                var code = str.charCodeAt(i)
                if (code < 65 || code > 90) {
                    return false
                }
            }
            return true
        }
        return false
    }
    , eventNS = 'events' // $.data(eventNS) event namespace
    , cors = 'withCredentials'

// core
$.fn = $.prototype = {
    version: '0.0.1',
    constructor: $,
    length: 0,
    expando: getExpando(),
    jquery: true,
    toArray: function() {
        return $.toArray(this)
    },
    get: function(i) {
        return null == i ? (num < 0 ? this[i + this.length] : this[i]) : this.toArray()
    },
    find: function(str) {
        return $($.find(str, this))
    },
    each: function(fn) {
        return $.each(this, fn)
    },
    map: function(fn) {
        return $.map(this, fn)
    },
    filter: function(fn) {
        return $.filter(this, fn)
    },
    pushStack: function(elems) {
        var ret = $.merge($(), elems)
        ret.prevObject = this
        ret.context = this.context
        return ret
    },
    end: function() {
        return this.prevObject || $()
    },
    eq: function(i) {
        var len = this.length
        i = +i
        if (i < 0) {
            i = len + i
        }
        return this.pushStack([this[i]] || [])
    },
    first: function() {
        return this.eq(0)
    },
    last: function() {
        return this.eq(-1)
    },
    slice: function() {
        return this.pushStack(slice.apply(this, arguments))
    }
}

$.extend = $.fn.extend = function() {
    var len = arguments.length
    var i = 1, obj
    var target = arguments[0] || {}
    if (1 == len) {
        target = this
        i--
    }
    for (; i < len; i++) {
        obj = arguments[i]
        if (obj && 'object' == typeof obj) {
            for (var k in obj) {
                target[k] = obj[k]
            }
        }
    }
    return target
}

$.extend({
    expando: getExpando(),
    toArray: function(arr, ret) {
        ret = ret || []
        var len = arr.length
        for (var i = 0; i < len; i++) {
            ret.push(arr[i])
        }
        return ret
    },
    isLetter: isLetter,
    noop: function() {},
    each: function(arr, fn) {
        var len = arr ? arr.length : 0
        if (len && len > 0 && 'function' == typeof fn) {
            for (var i = 0; i < len; i++) {
                if (false === fn.call(arr[i], i, arr[i], arr)) break
            }
        }
        return arr
    },
    filter: function(arr, fn) {
        var ret = []
        $.each(arr, function(k, v) {
            if (fn.call(this, k, v, arr)) {
                ret.push(v)
            }
        })
        return ret
    },
    map: function(arr, fn) {
        var ret = []
        $.each(arr, function(k, v) {
            ret[k] = fn.call(this, k, v, arr)
        })
        return ret
    },
    trim: function(str) {
        var rtrim = /^[\s\uFEFF\xA0]+|[\s\uFEFF\xA0]+$/g
        return null == str ? '' : (str + '').replace(rtrim, '')
    },
    proxy: function(fn, context) {
        var tmp
        if ('string' == typeof context) {
            tmp = fn
            fn = fn[context]
            context = tmp
        }
        var args = slice.call(arguments, 2)
        return function() {
            return fn.apply(context || this, args.concat(slice.call(arguments)))
        }
    },
    indexOf: function(arr, val, from) {
        from = from || 0
        if ('string' == typeof arr) {
            var index = arr.substr(from).indexOf(val)
            if (-1 == index) return -1
            return from + index
        }
        var i = from - 1
        var len = arr.length
        while (++i < len) {
            if (arr[i] === val) {
                return i
            }
        }
        return -1
    },
    now: function() {
        return +new Date()
    },
    merge: function(a, b) {
        var len = +b.length
        , j = 0
        , i = a.length
        for (; j < len; j++) {
            a[i++] = b[j]
        }
        a.length = i
        return a
    },
    find: function(str, context, ret) {
        ret = ret || []
        if (context && context.length) {
            for (var i = 0, node; node = context[i++];) {
                ret.push.apply(ret, $.findOne(str, node))
            }
        } else {
            ret.push.apply(ret, $.findOne(str, context))
        }
        return ret
    },
    findOne: function(str, dom) {
        dom = dom || document
        var nodes, ret = []
        if ('#' == str.charAt(0)) {
            var id = str.substr(1)
            if (id && dom.getElementById) {
                // $('#id')
                var el = dom.getElementById(id)
                if (el) {
                    return [el]
                }
            }
        }
        if (isLetter(str)) {
            // $('tag')
            nodes = dom.getElementsByTagName(str)
        } else if ($.customFind) {
            nodes = $.customFind(str, dom)
        } else if (dom.querySelectorAll) {
            nodes = dom.querySelectorAll(str)
        }
        // cannot use slice.call because dom object throw error in ie8-
        var len = nodes ? nodes.length : 0
        var ret = []
        for (var i = 0; i < len; i++) {
            ret[i] = nodes[i]
        }
        return ret
    },
    error: function(msg) {
        throw new Error(msg)
    },
    parseJSON: function(str) {
        if (window.JSON && JSON.parse) {
            return JSON.parse(str + '')
        }

        str = $.trim(str + '')
        var depth, requireNonComma

        return str && !$.trim(str.replace(rvalidtokens, function(token, comma, open, close) {
            if (requireNonComma && comma) depth = 0
            if (depth = 0) return token
            requireNonComma = open || comma
            depth += !close - !open
            return ''
        })) ? (Function('return ' + str))() : $.error("Invalid JSON: " + data)
    },
    parseXML: function(data) {
        var xml, tmp
        if (!data || 'string' != typeof data) return null
        try {
            if (window.DOMParser) {
                tmp = new DOMParser()
                xml = tmp.parseFromString(data, 'text/xml')
            } else {
                xml = new ActiveXObject('Microsoft.XMLDOM')
                xml.async = 'false'
                xml.loadXML(data)
            }
        } catch (e) {
            xml = undefined
        }
        if (!xml || !xml.documentElement || xml.getElementsByTagName('parsererror').length) {
            $.error('Invalid XML: ' + data)
        }
        return xml
    },
    getClassName: function(val) {
        var ret = {}.toString.call(val).split(' ')[1]
        return ret.substr(0, ret.length - 1).toLowerCase()
    },
    type: function(val) {
        var tp = typeof val
        if (tp == 'object' || tp == 'function') {
            var className = $.getClassName(val)
            if (-1 == $.indexOf(classes, className)) {
                return 'object'
            }
            return className
        }
        return tp
    },
    isArray: function(arr) {
        return 'array' == $.type(arr)
    },
    isFunction: function(func) {
        return 'function' == $.type(func)
    }
})

var root

$.fn.init = function (selector, context) {
    if (!selector) return this // $(null)
    var len = selector.length
    if ($.isFunction(selector)) {
        // TODO ready
        // $(callback)
        return selector()
    } else if (selector.nodeType) {
        // $(dom)
        this.context = this[0] = selector
        this.length = 1
    } else if ('string' == typeof selector) {
        if (-1 != selector.indexOf('<')) {
            // $('<div>') parse html
            var div = document.createElement('div')
            div.innerHTML = selector
            return $(div.childNodes)
        } else {
            context = context || root
            if (context.nodeType) context = $(context) // convert dom to $(dom)
            return context.find(selector)
        }
    } else if (len) {
        // $([dom1, dom2])
        // should after string, because string, function has length too
        for (var i = 0; i < len; i++) {
            if (selector[i] && selector[i].nodeType) {
                this[this.length] = selector[i]
                this.length++
            }
        }
    }
    return this
}

// the order...
$.fn.init.prototype = $.fn
root = $(document)

// access
$.access = function(elems, fn, key, val, isChain) {
    var i = 0
    if (key && 'object' === typeof key) {
        // set multi k, v
        for (i in key) {
            $.access(elems, fn, i, key[i], true)
        }
    } else if (undefined === val) {
        // get value
        var ret
        if (elems[0]) { // TODO text, html should be ''
            ret = fn(elems[0], key)
        }
        if (!isChain) {
            return ret
        }
    } else {
        // set one k, v
        for (i = 0; i < elems.length; i++) {
            fn(elems[i], key, val)
        }
    }
    return elems
}

// access
$.fn.extend({
    text: function(val) {
        return $.access(this, $.text, null, val)
    },
    html: function(val) {
        return $.access(this, $.html, null, val)
    },
    attr: function(key, val) {
        return $.access(this, $.attr, key, val)
    },
    prop: function(key, val) {
        return $.access(this, $.prop, key, val)
    },
    css: function(key, val) {
        return $.access(this, $.css, key, val)
    },
    data: function(key, val) {
        return $.access(this, $.data, key, val)
    },
    _data: function(key, val) {
        return $.access(this, $.data, key, val)
    },
    removeData: function(key) {
        return $.access(this, $.removeData, key, undefined, true)
    }
})

// manipulation
$.fn.extend({
    domManip: function(args, fn) {
        return this.each(function() {
            var node = $.buildFragment(args)
            // always build to one node(fragment)
            fn.call(this, node)
        })
    },
    remove: function() {
        return this.domManip(arguments, function() {
            if (this.parentNode) {
                this.parentNode.removeChild(this)
            }
        })
    },
    before: function() {
        return this.domManip(arguments, function(elem) {
            if (elem.parentNode) {
                elem.parentNode.insertBefore(elem, this)
            }
        })
    },
    after: function() {
        return this.domManip(arguments, function(elem) {
            if (elem.parentNode) {
                elem.parentNode.insertBefore(elem, this.nextSibling)
            }
        })
    },
    append: function() {
        return this.domManip(arguments, function(elem) {
            this.appendChild(elem)
        })
    }
})

function Data() {
    this.expando = getExpando()
    this.cache = []
}

Data.prototype = {
    get: function(owner, key) {
        var ret = this.cache[owner[this.expando]] || {}
        if (key) {
            return ret[key]
        }
        return ret
    },
    set: function(owner, key, val) {
        var expando = this.expando
        var cache = this.cache
        if (owner[expando] >= 0) {
            cache[owner[expando]][key] = val
        } else {
            var len = cache.length
            owner[expando] = len
            cache[len] = {}
            cache[len][key] = val
        }
    },
    remove: function(owner, key) {
        var expando = this.expando
        var cache = this.cache
        var len = owner[expando]
        if (len >= 0) {
            if (undefined === key) {
                cache[len] = {}
            } else {
                delete cache[len][key]
            }
        }
    }
}

var data_user = new Data

var data_priv = new Data

$.extend({
    attr: function(elem, key, val) {
        if (undefined === val) {
            return elem.getAttribute(key)
        } else if (null === val) {
            return elem.removeAttribute(key)
        }
        elem.setAttribute(key, '' + val)
    },
    text: function(elem, key, val) {
        if (undefined !== val) return elem.textContent = '' + val
        var nodeType = elem.nodeType
        if (3 == nodeType || 4 == nodeType) {
            return elem.nodeValue
        }
        if ('string' == typeof elem.textContent) {
            return elem.textContent
        }
        var ret = ''
        for (elem = elem.firstChild; elem; elem = elem.nextSibling) {
            ret += $.text(elem)
        }
        return ret
    },
    html: function(elem, key, val) {
        if (undefined === val) {
            return elem.innerHTML
        }
        elem.innerHTML = '' + val
    },
    prop: function(elem, key, val) {
        if (undefined === val) {
            return elem[key]
        }
        elem[key] = val
    },
    css: function(elem, key, val) {
        var style = elem.style || {}
        if (undefined === val) {
            var ret = style[key]
            if (ret) return ret
            if (window.getComputedStyle) {
                return getComputedStyle(elem, null)[key]
            }
        } else {
            style[key] = val
        }
    },
    data: function(elem, key, val) {
        if (undefined !== val) {
            // set val
            data_user.set(elem, key, val)
        } else {
            if (key && 'object' == typeof key) {
                // set multi
                for (var k in key) {
                    $.data(elem, k, key[k])
                }
            } else {
                // get
                return data_user.get(elem, key)
            }
        }
    },
    _data: function(elem, key, val) {
        if (undefined !== val) {
            // set val
            data_priv.set(elem, key, val)
        } else {
            if (key && 'object' == typeof key) {
                // set multi
                for (var k in key) {
                    $._data(elem, k, key[k])
                }
            } else {
                // get
                return data_priv.get(elem, key)
            }
        }
    },
    removeData: function(elem, key) {
        data_user.remove(elem, key)
    },
    buildFragment: function(elems, context) {
        context = context || document
        var fragment = context.createDocumentFragment()
        for (var i = 0, elem; elem = elems[i++];) {
            var nodes = []
            if ('string' == typeof elem) {
                if (elem.indexOf('<') == -1) {
                    nodes.push(context.createTextNode(elem))
                } else {
                    var div = document.createElement('div')
                    div.innerHTML = elem
                    $.toArray(div.childNodes, nodes)
                }
            } else if ('object' == typeof elem) {
                if (elem.nodeType) {
                    nodes.push(elem)
                } else {
                    $.toArray(elem, nodes)
                }

            }
        }
        for (var i = 0, node; node = nodes[i++];) {
            fragment.appendChild(node)
        }
        return fragment
    }
})


// events
$.Event = function(src, props) {
    if (!(this instanceof $.Event)) {
        return new $.Event(src, props)
    }
    src = src || {}
    if ('string' == typeof src) {
        src = {
            type: src
        }
    }
    this.originalEvent = src
    this.type = src.type
    this.target = src.target || src.srcElement
    if (props) {
        $.extend(this, props)
    }
}

function miniHandler(ev) {
    ev = $.Event(ev)
    var elem = this
    if (!elem.nodeType)
        elem = ev.target
    $.trigger(elem, ev)
}

$.extend({
    on: function(elem, type, handler, data, selector) {
        var events = $._data(elem, eventNS)
        var arr = type.split('.')
        type = arr[0]
        var namespace = arr[1]
        if (!events) {
            // set data priv
            events = {}
            $._data(elem, eventNS, events)
        }
        if (!events[type]) {
            events[type] = []
            if (elem.addEventListener) {
                elem.addEventListener(type, miniHandler, false)
            } else if (elem.attachEvent) {
                elem.attachEvent('on' + type, miniHandler)
            }
        }
        var typeEvents = events[type]
        var ev = {
            handler: handler,
            namespace: namespace,
            selector: selector,
            type: type
        }
        typeEvents.push(ev)
    },
    trigger: function(elem, ev) {
        var events = $._data(elem, eventNS)
        if ('string' == typeof ev) {
            // self trigger, ev = type
            ev = $.Event(ev, {
                target: elem
            })
        }
        var typeEvents = events[ev.type]
        if (typeEvents) {
            typeEvents = typeEvents.slice() // avoid self remove
            var len = typeEvents.length
            for (var i = 0; i < len; i++) {
                var obj = typeEvents[i]
                var ret = obj.handler.call(elem, ev)
                if (false === ret) {
                    ev.preventDefault()
                    ev.stopPropagation()
                }
            }
        }
    },
    off: function(elem, ev, handler) {
        var events = $._data(elem, eventNS)
        if (!events) return

        ev = ev || ''
        if (!ev || '.' == ev.charAt(0)) {
            // namespace
            for (var key in events) {
                $.off(elem, key + ev, handler)
            }
            return
        }

        var arr = ev.split('.')
        var type = arr[0] // always have
        var namespace = arr[1]
        var typeEvents = events[type]
        if (typeEvents) {
            for (var i = typeEvents.length - 1; i >= 0; i--) {
                var x = typeEvents[i]
                var shouldRemove = true
                if (namespace && namespace != x.namespace) {
                    shouldRemove = false
                }
                if (handler && handler != x.handler) {
                    shouldRemove = false
                }
                if (shouldRemove) {
                    typeEvents.splice(i, 1)
                }
            }
        }
    }
})

$.fn.extend({
    eventHandler: function(events, handler, fn) {
        if (!events) {
            return this.each(function() {
                fn(this)
            })
        }
        events = events.split(' ')
        return this.each(function() {
            for (var i = 0; i < events.length; i++) {
                fn(this, events[i], handler)
            }
        })
    },
    on: function(events, handler) {
        return this.eventHandler(events, handler, $.on)
    },
    off: function(events, handler) {
        return this.eventHandler(events, handler, $.off)
    },
    trigger: function(events, handler) {
        return this.eventHandler(events, handler, $.trigger)
    }
})

// TODO ajax, jsonp, evalUrl
var ajaxSetting = {
    xhr: function() {
        try {
            if (window.XMLHttpRequest) {
                return new XMLHttpRequest()
            }
            return new ActiveXObject('Microsoft.XMLHTTP')
        } catch (e) {}
    },
    jsonp: 'callback'
}

$.ajaxSetting = ajaxSetting

$.request = request

$.getScript = getScript

var support = {}

$.support = support

var querystring = {
    parse: function(qs, sep, eq) {
        if ('string' != typeof qs) return {}
        sep = sep || '&'
        eq = eq || '='
        var ret = {}
        qs = qs.split(sep)
        for (var i = 0; i < qs.length; i++) {
            var arr = qs[i].split(eq)
            if (2 == arr.length) {
                var k = arr[0]
                var v = arr[1] || ''
                if (k) {
                    try {
                        k = decodeURIComponent(k)
                        v = decodeURIComponent(v)
                        ret[k] = v
                    } catch (e) {}
                }
            }
        }
        return ret
    },
    stringify: function(obj, sep, eq) {
        if (obj && 'object' == typeof obj) {
            sep = sep || '&'
            eq = eq || '='
            var ret = []
            for (var k in obj) {
                var v = obj[k]
                if (v || '' === v || 0 === v) { // be careful with 0, ''
                    ret.push(encodeURIComponent(k) + eq + encodeURIComponent(v))
                }
            }
            return ret.join(sep)
        }
        return ''
    }
}

$.querystring = querystring

function getScript(url, opt, cb) {
    var head = $('head')[0]
    var script = document.createElement('script')

    function finish(err) {
        if (script) {
            if (cb) {
                // fake xhr
                var xhr = {
                    status: 200
                }
                if (err) {
                    xhr = {
                        status: 0
                    }
                }
                cb(err, xhr)
            }
            script.onload = script.onerror = script.onreadystatechange = null
            head.removeChild(script)
            script = null
        }
    }

    function send() {
        script.async = true
        script.src = url
        script.onload = script.onerror = script.onreadystatechange = function(ev) {
            var state = script.readyState
            if (ev && 'error' == ev.type) { // old browser has no onerror
                finish('error')
            } else if (!state || /loaded|complete/.test(state)) { // new browser has no state
                finish(null)
            }
        }
        head.appendChild(script)
    }

    return {
        abort: function() {
            cb = null
            finish()
        },
        send: send
    }
}

function getXhr(url, opt, cb) {

    var xhr = ajaxSetting.xhr()

    function send() {
        if (!xhr) return
        var type = opt.type || 'GET'

        url = bindQuery2url(url, opt.data)

        xhr.open(type, url, !cb.async)

        if (cors in xhr) {
            support.cors = true
            xhr[cors] = true // should after open
        }

        var headers = opt.headers
        if (headers) {
            for (var key in headers) {
                xhr.setRequestHeader(key, headers[key])
            }
        }

        xhr.send(opt.data || null)

        var handler = function() {
            if (handler &&  4 === xhr.readyState) {
                handler = undefined
                cb && cb(null, xhr, xhr.responseText)
            }
        }

        if (false === opt.async) {
            handler()
        } else if (4 === xhr.readyState) {
            setTimeout(handler)
        } else {
            xhr.onreadystatechange = handler
        }
    }

    function abort() {
        if (!xhr) return
        cb = null
        xhr.abort()
    }

    return {
        send: send,
        abort: abort
    }

}

function bindQuery2url(url, query) {
    if (-1 == url.indexOf('?')) {
        url += '?'
    }
    var last = url.charAt(url.length - 1)
    if ('&' != last && '?' != last) {
        url += '&'
    }
    return url + $.querystring.stringify(query)
}

function request(url, opt, cb) {
    // cb can only be called once
    if (url && 'object' == typeof url) {
        return $.ajax(url.url, url, cb)
    }

    if ('function' == typeof opt) {
        cb = opt
        opt = {}
    }

    var hasCalled = false
    var callback = function(err, res, body) {
        if (hasCalled) return
        cb = cb || $.noop
        hasCalled = true
        cb(err, res, body)
    }

    var dataType = opt.dataType || 'text' // html, script, jsonp, text

    var req
    var isJsonp = false
    if ('jsonp' == dataType) {
        isJsonp = true
        var jsonp = opt.jsonp || ajaxSetting.jsonp
        var jsonpCallback = opt.jsonpCallback || $.expando + '_' + $.now()
        var keyTmpl = jsonp + '=?'
        var query = $.extend({}, opt.data)
        if (url.indexOf(keyTmpl) != -1) {
            url.replace(keyTmpl, jsonp + '=' + jsonpCallback)
        } else {
            query[jsonp] = jsonpCallback
        }
        if (!opt.cache) {
            query._ = $.now()
        }
        url = bindQuery2url(url, query)

        dataType = 'script'
        window[jsonpCallback] = function(ret) { // only get first one
            callback(null, {
                status: 200
            }, ret)
            window[jsonpCallback] = null
        }
    }
    if ('script' == dataType) {
        req = getScript(url, opt, isJsonp ? null : callback)
    } else if (/html|text/.test(dataType)) {
        req = getXhr(url, opt, callback)
    }
    req.send()

    if (opt.timeout) {
        setTimeout(function() {
            req.abort() // should never call callback, because user know he abort it
            callback('timeout', {
                status: 0,
                readyState: 0,
                statusText: 'timeout'
            })
            if (isJsonp) {
                window[jsonpCallback] = $.noop
            }
        }, opt.timeout)
    }
}

$.ajax = function(url, opt) {
    // TODO fuck the status, statusText, even for jsonp
    var ret = {}
    request(url, opt, function(err, xhr, body) {
        xhr = xhr || {}
        var jqxhr = {
            status: xhr.status,
            statusText: xhr.statusText,
            readyState: xhr.readyState
        }
        $.extend(ret, jqxhr)
        // success, timeout, error
        var resText = 'success'
        if (err || 200 != ret.status) {
            resText = 'error'
            if ('timeout' == err) {
                resText = 'timeout'
            }
            opt.error && opt.error(ret, resText, xhr.statusText)
        } else {
            opt.success && opt.success(body || xhr.responseText, resText, ret)
        }
        opt.complete && opt.complete(ret, resText)
    })
    return ret
}

$.each(['get', 'post'], function(i, method) {
    $[method] = function(url, data, callback, dataType) {
        if ('function' == typeof data) {
            dataType = callback
            callback = data
            data = undefined
        }
        $.ajax(url, {
            type: method,
            dataType: dataType,
            data: data,
            success: callback
        })
    }
})

// show hide toggle
var display = 'display'
// init iframe for dirty thing
/*
var iframe = $("<iframe frameborder='0' width='0' height='0' />")
$(document.documentElement).append(iframe)
var iframeDoc
try {
    iframeDoc = (iframe[0].contentWindow || iframe[0].contentDocument).document
    iframeDoc.write()
    iframeDoc.close()
} catch (e) {}
var defaultDisplayCache = {}
function getDefaultDisplay(tag, doc) {
    if (!doc) return
    if (!defaultDisplayCache[tag]) {
        var el = doc.createElement(tag)
        $(doc.body).append(el)
        defaultDisplayCache[tag] = $.css(el, display) || 'block'
    }
    return defaultDisplayCache[tag]
}*/

$.each('show hide toggle'.split(' '), function(i, key) {
    $.fn[key] = function() {
        return this.each(function(i, el) {
            var act = key
            var old = $.css(el, display)
            var isHiden = 'none' == old
            if ('toggle' == key) {
                act = 'hide'
                if (isHiden) {
                    act = 'show'
                }
            }
            if ('show' == act && isHiden) {
                var css = $._data(el, display) || 'block'
                $.css(el, display, css)
            } else if ('hide' == act && !isHiden) {
                $._data(el, display, old)
                $.css(el, display, 'none')
            }
        })
    }
})

return $

})
