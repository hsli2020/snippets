/**
 * http://4dd.jp/ 
 * @author http://nki2.jp/
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

jQuery.extend({	

	url: function(arg, url) {
		var _ls = url || window.location.toString();

		if(_ls.substring(0,2) === '//') _ls = 'http:' + _ls;
		else if(_ls.split('://').length === 1) _ls = 'http://' + _ls;

		url = _ls.split('/');
		var _l = {auth:''}, host = url[2].split('@');

		if(host.length === 1) host = host[0].split(':');
		else{ _l.auth = host[0]; host = host[1].split(':'); }

		_l.protocol=url[0], _l.hostname=host[0], _l.port=(host[1]||'80'), _l.pathname='/' + url.slice(3, url.length).join('/').split('?')[0].split('#')[0];
		var _p = _l.pathname;
		if(_p.split('.').length === 1 && _p[_p.length-1] !== '/') _p += '/';
		var _h = _l.hostname, _hs = _h.split('.'), _ps = _p.split('/');

		if(!arg) return _ls;
		else if(arg === 'hostname') return _h;
		else if(arg === 'domain') return _hs.slice(-2).join('.');
		else if(arg === 'tld') return _hs.slice(-1).join('.');
		else if(arg === 'sub') return _hs.slice(0, _hs.length - 2).join('.');
		else if(arg === 'port') return _l.port || '80';
		else if(arg === 'protocol') return _l.protocol.split(':')[0];
		else if(arg === 'auth') return _l.auth;
		else if(arg === 'user') return _l.auth.split(':')[0];
		else if(arg === 'pass') return _l.auth.split(':')[1] || '';
		else if(arg === 'path') return _p;
		else if(arg[0] === '.')
		{
			arg = arg.substring(1);
			if($.isNumeric(arg)) {arg = parseInt(arg); return _hs[arg < 0 ? _hs.length + arg : arg-1] || ''; }
		}
		else if($.isNumeric(arg)){ arg = parseInt(arg); return _ps[arg < 0 ? _ps.length - 1 + arg : arg] || ''; }
		else if(arg === 'file') return _ps.slice(-1)[0];
		else if(arg === 'filename') return _ps.slice(-1)[0].split('.')[0];
		else if(arg === 'fileext') return _ps.slice(-1)[0].split('.')[1] || '';
		else if(arg[0] === '?' || arg[0] === '#')
		{
			var params = _ls, param = null;

			if(arg[0] === '?') params = (params.split('?')[1] || '').split('#')[0];
			else if(arg[0] === '#') params = (params.split('#')[1] || '');

			if(!arg[1]) return params;

			arg = arg.substring(1);
			params = params.split('&');

			for(var i=0,ii=params.length; i<ii; i++)
			{
				param = params[i].split('=');
				if(param[0] === arg) return param[1];
			}
		}

		return '';
	}
});

var okb = okb || {};

okb.EventDispatcher = Class.extend({

	__construct:function(){
		var me = this;
		me.listeners = {};
	},

	destroy:function(){
		var me = this;
		delete me.listeners;
	},

	bind:function(type, callback, args) {
		var me = this;
		if(!me.listeners[type]) me.listeners[type] = [];
		me.listeners[type].push({callback:callback, args:args});
	},
	unbind:function(type, callback) {
		var me = this;
		if(!me.listeners[type]) return;
		var i,len = me.listeners[type].length;
		var newArray = [];
		for(i=0; i<len; i++) {
			var listener = me.listeners[type][i];
			if(listener.callback == callback) {
			} else {
				newArray.push(listener);
			}
		}
		this.listeners[type] = newArray;
	},
	trigger:function(type, args) {
		var me = this;
		var i,len;
		var event = {
			type:type,
			target:me
		};
		args = args || [];
		args = [event].concat(args);
		if(!me.listeners[type]) return;
		len = me.listeners[type].length;
		for(i=0; i<len; i++) {
			var listener = this.listeners[type][i];
			if(listener && listener.callback) {
				listener.args = args.concat(listener.args);
				listener.callback.apply(null, listener.args);
			}
		}
	}
});

okb.Ctrl = okb.EventDispatcher.extend({

	EV_DOMREADY:"evDomReady",
	EV_WINDOW_LOADED:"evWindowLoaded",
	EV_RESIZE:"evResize",
	EV_SCROLL:"evScroll",
	EV_HASHCHANGE:"evHashChanged",

	__construct:function(){
		this.__super.__construct.apply(this, arguments)
		var me = this;

		//ユーザーエージェントを取得
		var ua = navigator.userAgent;
		if(ua.indexOf("iPhone")>=0 ||
			ua.indexOf("iPad")>=0 ||
			ua.indexOf("iPod")>=0 ||
			ua.indexOf("Android")>=0 ||
			ua.indexOf("BlackBerry")>=0 ||
			ua.indexOf("IEMobile")>=0) {
			me.sp = true;
		}
		me.touchDevice = (document.ontouchstart !== undefined);
		if(ua.indexOf("Safari")>=0) me.safari = true;
		if(ua.indexOf("Chrome")>=0) me.chrome = true;
		if(ua.indexOf("Firefox")>=0) me.ff = true;
		if(ua.indexOf("Opera")>=0) me.opera = true;
		if(ua.indexOf("MSIE")>=0) me.ie = true;
		if(ua.indexOf("MSIE 6")>=0) me.ie6 = true;
		if(ua.indexOf("MSIE 7")>=0) me.ie7 = true;
		if(ua.indexOf("MSIE 8")>=0) me.ie8 = true;
		if(ua.indexOf("MSIE 9")>=0) me.ie9 = true;
		if(ua.indexOf("iPhone")>=0) me.iPhone = true;
		if(ua.indexOf("iPhone OS 4")) me.iPhone4 = true;
		if(ua.indexOf("iPhone OS 5")) me.iPhone5 = true;
		if(ua.indexOf("iPhone OS 6")) me.iPhone6 = true;
		if(ua.indexOf("iPad")>=0) me.iPad = true;
		if(ua.indexOf("Android")>=0) me.android = true;
		if(me.ie6 || me.ie7 || me.ie8) me.ie678 = true;
		if(me.ie6 || me.touchDevice) me.noFixed = true;
		if(me.iPhone5 || me.iPhone6) me.noFixed = false;
		if(me.iPhone && me.iPad && (ua.indexOf("OS 5")>=0 || ua.indexOf("OS 6")>=0)) {
			me.ios5 = true;
			me.noFixed = false;
		}
	},

	domReady:function(){
		var me = this;

		//cast
		me.window = window;
		me.document = document;
		me.location = location;
		me.$window = $(window);
		me.$document = $(document);
//		me.$html_body = $($.browser.opera ? document.compatMode == 'BackCompat' ? 'body' : 'html' :'html,body');
		me.$html_body = $("html,body");
		me.$body = $("body");
		me.$html = $("html");


		//htmlにクラスを追加
		if (me.ie) me.$html.addClass("ie")
		if (!me.ie) me.$html.addClass("notIE")
		if (me.ie678) me.$html.addClass("ie678")
		if (!me.chrome) me.$html.addClass("notChrome")
		if (me.noFixed) me.$html.addClass("noFixed")
		if (me.ff) me.$html.addClass("ff")

		if(_ctrl.touchDevice) $('html').addClass('touchDevice');
		else $('html').addClass('notTouchDevice');

		//スクロールイベント
		me.scrollTop = 0;
		me.scrolled = function (e) {
			me.preScrollTop = me.scrollTop;
			me.scrollTop = me.$window.scrollTop();
			me._getSize();
			me.trigger(me.EV_SCROLL);
		}
		me.$window.bind("scroll", me.scrolled);

		//ブラウザのリサイズイベント
		me.innerWidth = 0;
		me.innerHeight = 0;
		me._getSize = function(){
			me.windowW = me.$window.width();
			me.windowH = me.$window.height();
			me.bodyW = me.$body.width();
			me.bodyH = me.$body.height();
			me.scrollW = document.documentElement.scrollWidth || document.body.scrollWidth;
			me.scrollH = document.documentElement.scrollHeight || document.body.scrollHeight;
			me.clientW = document.documentElement.clientWidth || document.body.clientWidth;
			me.clientH = document.documentElement.clientHeight || document.body.clientHeight;
			var innerWidth = window.innerWidth || me.windowW;
			var innerHeight = window.innerHeight || me.windowH;
			me.stageW = Math.max(me.clientW, me.scrollW);
			me.stageH = Math.max(me.clientH, me.scrollH);

			//スマホではヘッダーもinnerHeightに含まれちゃうので、innerHeightを変動させないように
			me.pre_ration = me.ratio;
			me.ratio = innerWidth/innerHeight;
			var changeRotate = ( me.pre_ration<1&&me.ratio>1 || me.pre_ration>1&&me.ratio<1 );
			if(changeRotate) {
				me.innerWidth = 0;
				me.innerHeight = 0;
			}
			if(me.touchDevice) {
				me.innerWidth = Math.max( me.innerWidth, innerWidth );
				me.innerHeight = Math.max( me.innerHeight, innerHeight );
			} else {
				me.innerWidth = innerWidth;
				me.innerHeight = innerHeight;
			}
		}
		me.resized = function (e) {
			me._getSize();
			me.trigger(me.EV_RESIZE);

			me.scrolled(null);
		}
		me.$window.bind("resize", me.resized);
		me.$window.trigger("resize");

		//ドキュメント要素のサイズ変更を検出（ブラウザサイズの変更でなく）
		me.preScrollW = 0;
		me.preScrollH = 0;
		me.resizeWatch = function(){
			var scrollW = document.documentElement.scrollWidth || document.body.scrollWidth;
			var scrollH = document.documentElement.scrollHeight || document.body.scrollHeight;
			if(scrollW != me.preScrollW || scrollH != me.preScrollH) {
				me.resized(null)
			}
			me.preScrollW = scrollW;
			me.preScrollH = scrollH;
			setTimeout(me.resizeWatch, 500)
		}
		me.resizeWatch();

		//ハッシュの変更イベント
		me.hashChanged = function(e){
			var hash = me.location.hash || "";
			if(hash.indexOf("#/")>=0) hash = hash.substr(2);
			if(hash.indexOf("#")>=0) hash = hash.substr(1);
			me.hash = hash;

			//shadowbox
			if(me.shadowboxEnabledDeepLink) me._switchShadowbox();

			me.trigger(me.EV_HASHCHANGE);
		}
		me.$window.hashchange(me.hashChanged)
		me.hashChanged(null);

		me.trigger(me.EV_DOMREADY);
	},
	windowLoaded:function(){
		var me = this;

		//ステージサイズの再取得
		me.resized(null);

		me.trigger(me.EV_WINDOW_LOADED);
	},

	/* locationの変更 */
	changeLoc:function(href){
		var me = this;
		me.location.href = href;
	},
	replaceLoc:function(href){
		var me = this;
		me.location.replace(href);
	},

	/* クッキー操作 getter/setter */
	cookie: function(cookieName, value, expires){
		var me = this;
		if(value===undefined || value===null)  return $.cookie(cookieName);
		if(expires===undefined) expires = 14;
		var option = { expires: expires, path: '/' };
		if(expires==0 || expires==null) delete option["expires"];
		$.cookie(cookieName, value, option);
	},

	/* shadowboxの初期化 */
	setUpShadowbox:function(){
		var me = this;
		me.shadowboxReady = true;

		me.shadowbox = new okb.ui.Shadowbox();
		$("a.okb-shadowbox").each(function(index){
			var $btn = $(this);
			var href = $btn.attr("href");
			var rel = $btn.attr("rel");
			$btn.click(function(e){
				e.preventDefault();
				if(href.substr(0,1)=="#") me.changeLoc(href);
				else _ctrl.shadowbox.open( href, rel );
			})
		})
	},

	/*
	hashIdxを渡すとディープリンクが有効化される
	@hashIdx => {href:"URL", rel:"オプション"}
	 */
	enableShadowBoxDeepLink:function(hashIdx) {
		var me = this;

		me.shadowboxEnabledDeepLink = true;
		me.shadowboxHashIdx = hashIdx;

		me.shadowbox.bind(me.shadowbox.EV_CLOSE, function(){
			me.changeLoc("#/");
		})

		setTimeout(function(){
			me._switchShadowbox();
		}, 0)
	},
	updateShadowBoxDeepLink:function(hashIdx) {
		var me = this;
		me.shadowboxHashIdx = hashIdx;
	},
	_switchShadowbox:function(){
		var me = this;
		var hash = me.hash;
		if(hash.substr(-1)=="/") hash = hash.substr(0, hash.length-1);//最後がスラッシュでもOKなように
		var obj = me.shadowboxHashIdx[hash];
		//close
		if(!obj) {
			me.shadowbox.close();
		}
		//open
		else {
			me.shadowbox.open(obj.href, obj.rel)
		}
	}
});

okb.Comm = Class.extend({

	__construct:function(){
		var me = this;
	},

	load:function(apiObj, postData, successFunc, errorFunc) {
		var me = this;

		//cencel
		if(me._$currentXHR) me.cancel();

		//送るデータを整形
		if(!postData || postData == "") postData = {};
		if(typeof postData == "string") {
			var str = postData;
			postData = {};
			var strArr = str.split("&");
			var i, len = strArr.length;
			for(i=0; i<len; i++) {
				var combiArr = strArr[i].split("=");
				postData[combiArr[0]] = combiArr[1];
			}
		}

		//load
		me._$currentXHR = $.ajax({
			type: apiObj.method,
			dataType: "text",
			cache: false,
			url: apiObj.path,
			data: postData,
			success:function(data) {
				try {
					me.data = $.evalJSON(data);
				} catch(e) {
					trace("api parse error ::: "+apiObj.path)
					trace(e)
					trace("data:"+data)
					me.data = {};
					if(errorFunc) errorFunc();
					return;
				}
				if(successFunc) successFunc();
			},
			error:function(XMLHttpRequest, textStatus, errorThrown) {
				if(errorFunc) errorFunc();
			}
		});
	},

	cancel:function() {
		var me = this;
		if(me._$currentXHR) {
			me._$currentXHR.abort();
			me._$currentXHR = null;
			me.data = null;
		}
	}
});

var _ctrl = new okb.Ctrl();
$(function(){
	_ctrl.domReady();
})
$(window).load(function(){
	_ctrl.windowLoaded();
})

okb.ui = {};

okb.ui.Shadowbox = okb.EventDispatcher.extend({
	EV_OPEN:"evOpen",
	EV_CLOSE:"evClose",
	EV_SHOW_NEXT:"evShowNext",

	__construct:function(){
		this.__super.__construct.apply(this, arguments);
		var me = this;

		//const
		me.TYPE_IMAGE = "image";
		me.TYPE_SWF = "swf";
		me.TYPE_IFRAME = "iframe";
		me.DEF_PADDING = 10;
		me.DEF_STAGE_PADDING = 20;

		//cast
		me.$ = $('<div id="okb-shadowbox-wrapper">'+
				'   <div class="bg"></div>'+
				'	<div class="base"><div class="baseInner"><div class="baseInner2"></div></div></div>'+
				'	<div class="container">'+
				'		<div class="holder">'+
				'			<div class="inner">inner</div>'+
				'		</div>'+
				'		<div class="cover"></div>'+
				'		<a href="#" class="close">close</a>'+
				'       <p class="loading"></p>'+
				'	</div>'+
				'</div>',+
				'').appendTo(_ctrl.$body);

		me.$bg = $(".bg", me.$);
		me.$base = $(".base", me.$);
		me.$baseInner = $(".baseInner", me.$);
		me.$container = $(".container", me.$);
		me.$holder = $(".holder", me.$);
		me.$inner = $(".inner", me.$);
		me.$cover = $(".cover", me.$);
		me.$close = $(".close", me.$);
		me.$loading = $(".loading", me.$);

		//base
		me.c_base = new okb.ui.Cast( me.$base );
		me.c_base.castHide();

		//holder
		me.$holder.css("visibility", "hidden");

		//inner
		me.$inner.html("");

		//cover
		me.c_cover = new okb.ui.Cast( me.$cover );
		me.c_cover.castHide();

		//bg
		me.c_bg = new okb.ui.Cast( me.$bg );
		me.c_bg.castHide();

		//close
		me.c_close = new okb.ui.Cast( me.$close );
		me.c_close.castHide();
		me.c_close.$.click(function(){
			me.close();
			return false;
		})

		//loading
		me.c_loading = new okb.ui.Cast( me.$loading );
		me.c_loading.castHide();

		//resize
		_ctrl.bind(_ctrl.EV_RESIZE, function(e){
			me._resized.apply(me);
		});
		me._resized();

		//クリックで閉じる
		me.$close.remove();
		me.$.click(function(e){
			e.preventDefault();
			me.close();
		})

		me.option = {};
		me.padding = me.DEF_PADDING;
		me.marginTop = 0;
	},

	domReady:function(){
		var me = this;
		me._resized();
	},

	_getSizeObj:function(){
		var me = this;

		var sizeObj = {
			"width": 0+"px",
			"height": 0+"px",
			"margin-left": 0+"px",
			"margin-top": 0+"px"
		}
		if(me.isOpen) {
			me.padding = Number( me.option["padding"] || me.DEF_PADDING );
			me.contentW = me.scaledW = Number(me.option["width"]);
			me.contentH = me.scaledH = Number(me.option["height"]);
			me.marginTop = Number( me.option["margin-top"] || 0 );
			var w = me.contentW + me.padding*2;
			var h = me.contentH + me.padding*2;
			var sw = _ctrl.clientW-me.DEF_STAGE_PADDING*2;
			var sh = _ctrl.clientH-me.DEF_STAGE_PADDING*2;
			if(me.currentType == me.TYPE_IMAGE){
				var scale = (w>sw | h>sh )? Math.min(sw/w, sh/h): 1;
				w = Math.round(w*scale);
				h = Math.round(h*scale);
				me.scaledW = w-me.padding*2;
				me.scaledH = h-me.padding*2;
			}
			var w2 = Math.round(w*0.5);
			var h2 = Math.round(h*0.5);
			sizeObj = {
				"width": w+"px",
				"height": h+"px",
				"margin-left": -w2+"px",
				"margin-top": (-h2-Math.round(me.marginTop*0.5))+"px"
			}
		}
		return sizeObj;
	},

	_resized:function(withBaseAnime){
		var me = this;

		if(me.isOpen){
			var sizeObj = me._getSizeObj();

			//container
			me.$container.css(sizeObj)

			//base
			if(withBaseAnime) {
				sizeObj.avoidTransforms = true;
				me.$baseInner.stop().animate(sizeObj, 400, "easeInOutQuart");
			} else {
				me.$baseInner.stop().css(sizeObj, 400);
			}

			//image
			if(me.currentType == me.TYPE_IMAGE) {
				me.$img.css({
					"width": me.scaledW+"px",
					"height":me.scaledH+"px"
				})
			}
			else if(me.currentType == me.TYPE_IFRAME) {
				me.$iframe.css({
					"width": me.contentW+"px",
					"height": (me.contentH-me.marginTop)+"px"
				})
			}

			//holder
			me.marginTop = Number( me.option["margin-top"] || 0 );
			me.$holder.css({
				"left": me.padding+"px",
				"top": (me.padding+me.marginTop)+"px"
			})
		}

		var p = $("#page")[0];
		var pageH = _ctrl.stageH;
		if(p) pageH = Math.max( p.scrollHeight, pageH );

		//bg
		me.$bg.css("height", pageH+"px");

		//this
		me.$.css("height", pageH+"px");
	},

	openSerial:function(contentArr, serialCompCallback){
		var me = this;

		me.serialCompCallback = serialCompCallback;

		if(me.hashNextSerial) {
			me.contentArr = me.contentArr.concat(contentArr);
		}

		me.contentArr = contentArr;

		if(!me.serialReady) {
			me.serialReady = true;
			me.bind(me.EV_SHOW_NEXT, function(){
				me._serialLoop();
			})
		}

		me._serialLoop();
	},

	_serialLoop:function(){
		var me = this;

		var contentObj = me.contentArr.shift();
		if(me.contentArr.length>0) me.hashNextSerial = true;
		else me.hashNextSerial = false;

		me.open( contentObj.href, contentObj.rel, callback );
	},

	open:function(contentStr, rel, callback){
		var me = this;

		me.callback = callback;

		//resize
		me._resized();

		if(me.isOpen) {
			me._removeContent();
//			return;
		}
		me.isOpen = true;
		me.trigger(me.EV_OPEN);

		if(me.delayID) clearTimeout(me.delayID);

		//loading
		me.c_loading.castShow(100, 33);

		//this
		me.$.css("visibility", "visible");

		//bg
		me.c_bg.castShow( (!_ctrl.ie6? 100: 0) );
		if(_ctrl.ie6) me.$bg.fixPng();

		//fixed効かない場合の位置取り
		if(_ctrl.noFixed){
			var ty = Math.round( _ctrl.scrollTop + (_ctrl.clientH*0.5) );
			me.$base.css("top", ty+"px");
			me.$container.css("top", ty+"px");
		}

		me.currentType = "";

		//option
		me.option = {};
		rel = rel || "";
		rel = rel.split(" ").join("");
		var relArr = rel.split(",");
		var i,len = relArr.length;
		for(i=0; i<len; i++){
			var opArr = relArr[i].split(":");
			me.option[opArr[0]] = opArr[1];
		}

		//クラス
		if(me.option["class"]) me.$.addClass(me.option["class"]);
		else me.$.removeAttr("class");


		if(me.delayID) clearTimeout(me.delayID);
		me.delayID = setTimeout(function(){

			me._addContent(contentStr);

		}, 100)
	},

	_addContent:function(contentStr){
		var me = this;

		if(contentStr.match(/.jpg|.gif|.png/i)) me.currentType = me.TYPE_IMAGE;
		else if(contentStr.match(/.swf/)) me.currentType = me.TYPE_SWF;
		else me.currentType = me.TYPE_IFRAME;

		if(me.option["type"]) me.currentType = me.option["type"];

		//image
		if(me.currentType == me.TYPE_IMAGE) {
			me.img = new Image();
			me.img.src = contentStr;
			me.$img = $(me.img).appendTo(me.$inner);
			me.imgLoaded = false;
			if(me.delayID) clearTimeout(me.delayID);
			me.delayID = setTimeout(function(){
				me.$img.trigger("imgLoad")
			}, 10000)
			me.$img.bind("imgLoad", function(){
				if(me.imgLoaded) return;
				me.imgLoaded = true;
				me.$img.unbind();

				var imgSize = okb.util.getImageSize(me.img);
				me.option["width"] = imgSize["width"];
				me.option["height"] = imgSize["height"];
				me._getSizeObj();
				me._resized(true);

				me._open(contentStr);
			})
		}
		//swf
		else if(me.currentType == me.TYPE_SWF) {
			me._open(contentStr);
		}
		//iframe
		else {
			me.$iframe = $('<iframe src="javascript:false;" width="800" height="800" frameborder="0" scrolling="no" allowtransparency="true"></iframe>').appendTo(me.$inner);
			me.$iframe.load(function(){
				me.$iframe.unbind();
				if(me.option["bodyStyle"]) me.$iframe.contents().find("body").attr("style", me.option["bodyStyle"])
				me._open(contentStr);
			});
			me.$iframe[0].contentWindow.location.replace(contentStr);
		}
	},

	_open:function(contentStr){
		var me = this;

		//loading
		me.c_loading.castHide(100, 0);

		if(me.delayID) clearTimeout(me.delayID);
		me.delayID = setTimeout(function(){

			//resize
			me._resized(true);

			//base
			var time = _ctrl.ie678? 0: 400;
			me.c_base.castShow(time, 0)

			//base表示後に
			if(me.delayID) clearTimeout(me.delayID);
			me.delayID = setTimeout(function(){

				//content
				me.$holder.css("visibility", "visible");

				//SWF
				if(me.currentType == me.TYPE_SWF) {
					me.$inner.html('<div id="ex_flash_okb"></div>')
					var flashvars = {
					};
					var params = {
						menu: "false",
						allowscriptaccess: "always",
						wmode: (_ctrl.os == "Windows XP")? null: "transparent",
						base:"."
					};
					var attributes = {
						id: "ex_flash_okb",
						name: "ex_flash_okb"
					};
					swfobject.embedSWF(contentStr, "ex_flash_okb", me.contentW, (me.contentH-me.marginTop), "10.0.0", "/common/js_libs/expressInstall.swf", flashvars, params, attributes);
				}


				//cover
				if(me.currentType != me.TYPE_SWF) {
					me.$inner.stop().fadeTo(0,1).css("filter", "none");
					me.c_cover.castShow();
					me.c_cover.castHide(  me.option["fade"]=="none"? 0: 300 );
				} else {
//					me.$inner.fadeTo(0,0).stop().fadeTo(500, 1);
//					me.delayID = setTimeout(function(){
						me.$inner.stop().fadeTo(0,1).css("filter", "none");
//					}, 500);
				}

				//close
				if(!me.option["noClose"]) {
					me.c_close.castShow( 300, 300 );
				}


				//念のためbg表示
				if(!_ctrl.ie678) {
					if(me.delayID) clearTimeout(me.delayID);
					me.delayID = setTimeout(function(){
						me.$bg.stop().fadeTo(0,1);
					}, 100);
				}

			}, 300);

		}, 0)
	},


	close:function(){
		var me = this;

		if(!me.isOpen) return;
		me.isOpen = false;
		me.trigger(me.EV_CLOSE);

		//close
		me.c_close.castHide(50, 0);

		//cover
		if(me.marginTop==0) {
			me.c_cover.castShow(100, 0);
		} else {
			me.$inner.stop().fadeTo(100,0);
		}

		if(me.delayID) clearTimeout(me.delayID);
		me.delayID = setTimeout(function(){
			me.c_cover.castHide();

			//content (remove)
			me._removeContent();

			//base
			var time = _ctrl.ie678? 0: 100;
			me.c_base.castHide(time);
			me.$baseInner.stop().animate({
				"width": 0+"px",
				"height": 0+"px",
				"margin-left": 0+"px",
				"margin-top": 0+"px",
				avoidTransforms: true
			}, 100, "easeOutQuart");

			if(me.hashNextSerial) {
				setTimeout(function(){
					me.trigger(me.EV_SHOW_NEXT);
				}, 100)
			} else {
				//bg
				me.c_bg.castHide( (!_ctrl.ie6? 100: 0), 100);

				//this
				if(me.delayID) clearTimeout(me.delayID);
				me.delayID = setTimeout(function(){
					me.$.css("visibility", "hidden");

					//callback
					if(me.callback) {
						me.callback();
						me.callback = null;
					}
					if(me.serialCompCallback) {
						me.serialCompCallback();
						me.serialCompCallback = null;
					}
				}, 200)
			}
		}, 100)
	},

	_removeContent:function(){
		var me = this;

		me.$inner.html("");
		me.$holder.css("visibility", "hidden");
		//image
		if(me.currentType == me.TYPE_IMAGE) {
			me.$img.unbind().remove();
			me.$img = me.img = null;
		}
		//SWF
		if(me.currentType == me.TYPE_SWF) {
			swfobject.removeSWF("ex_flash_okb");
		}
		//iframe
		if(me.currentType == me.TYPE_IFRAME) {
			me.$iframe.unbind().remove();
			me.$iframe = null;
		}
	}
});

_ctrl.bind(_ctrl.EV_DOMREADY, function(e){

	var isSupportPlaceHolder = ("placeholder" in document.createElement('input') );

	$(".imgBtn").each(function(index){
		var $this = $(this);

		new okb.ui.ImgBtn( $this, {
			isRev: $this.hasClass("rev")
		} );
	});

	$(".alpBtn").each(function(index){
		new okb.ui.AlphaBtn( $(this) );
	});

	//scrollTopBtn
	$(".scrollTopBtn").each(function(index){
		var $this = $(this);
		$this.click(function(e){
			if($this.attr("href") == "#") e.preventDefault();

//			_ctrl.$html_body.scrollTop(0)
			_ctrl.$html_body.stop().animate({
				"scrollTop":0,
				avoidTransforms: true
			}, 500)
		})
	})
})

okb.form = {};

okb.form.FormButton = Class.extend({

	__construct: function ($me) {
		var me = this;

		me.$ = $me;
		me.$input = $("input", me.$);

		//tabindexを引き継ぐ
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

// var me = this;
// var my = this;

dd.SceneAbout = dd.SceneBase.extend({

	__construct: function ($me, pathName) {
		this.__super.__construct.apply(this, arguments)
		var me = this;

		//main cover
		me.$mainBlock = me.$.find(".block-main");

		//map
		me._showMap();
	},

	_showMap:function(){
		var me = this;

		/*  map
		--------------------------------------------------*/
		lat = 35.661823;
		lng = 139.705538;
		me.zoom = 17;
		var zoom = me.zoom;

		/* js map */
		me.myLatlng = new google.maps.LatLng(lat, lng);
		var mapOptions = {
			center: me.myLatlng,
			zoom: zoom,
			disableDefaultUI: true,
			zoomControl: false,
			scrollwheel: false,
			mapTypeId: google.maps.MapTypeId.ROADMAP
		};
		me.map = new google.maps.Map(document.getElementById("map"), mapOptions);

		//pin
		if(!_ctrl.ie) {
			var pinSrc = location.origin + $("#pin img").attr("src");
			var markerPin = new google.maps.Marker({
				position: me.myLatlng,
				map: me.map,
				icon: new google.maps.MarkerImage(
					pinSrc,                     // url
					new google.maps.Size(200,50), // size
					new google.maps.Point(0,0),  // origin
					new google.maps.Point(100,50) // anchor
				)
			});
		} else {
			var marker = new google.maps.Marker({
				position: me.myLatlng,
				map: me.map
			});
		}

	},

	show:function(withAnime){
		var me = this;
		this.__super.show.apply(this, arguments)
	},

	doResize: function(){
		var me = this;
		this.__super.doResize.apply(this, arguments);

		//resize map
		if(me.resizeID) clearTimeout(me.resizeID);
		me.resizeID = setTimeout(function(){

			//move
			me.map.panTo( me.myLatlng );

			//zoom
			var zoom = me.zoom;
			if(_ctrl.scrollW<900) zoom -= 1;
			if(_ctrl.scrollW<600) zoom-= 2;
			me.map.setZoom(zoom);

		}, 200);

		//main cover
//		var w = _ctrl.innerWidth;
//		var h = _ctrl.innerHeight;
//		var picH = h*0.75;
//		var minRatio = 500/1600;
//		var ratio = picH/w;
//		if(ratio<minRatio) picH = w * minRatio;
//		me.$mainBlock.find(".cover").css("height", picH);
	},

	doScroll: function(){
		var me = this;
		this.__super.doScroll.apply(this, arguments)
	},

	destroy:function(){
		var me = this;
		this.__super.destroy.apply(this, arguments)
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
		if(validityms != undefined) arr.push("expires=" + (new Date((new Date()).getTime() + validityms).toUTCString()));
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


