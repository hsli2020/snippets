function isMobile() {
    return window.screen.availWidth<=768 || 
        /Android|iPhone|iPod/i.test(navigator.userAgent)
}
function isApp() {
    return navigator.userAgent.toLowerCase().indexOf("educationapp")>-1
}

document.documentElement.className+=isMobile()?" mobile":" pc",
document.documentElement.className+=isApp()?" app":" ",

function() {
    var e,n;
    n=isMobile() ? [ "http://ke.qq.com/mobilev2/h5components/header/header.css" ]
                 : [ {com:"r_header",url:"http://9.url.cn/edu/cms/edu/shengkao/component/r_header/css/header.3689247.css" },
                     {com:"r_panel",url:"http://9.url.cn/edu/cms/edu/shengkao/component/r_panel/css/responsive.e147748.css" },
                     {com:"r_teacher",url:"http://9.url.cn/edu/cms/edu/shengkao/component/r_teacher/css/responsive.6e31002.css" }
                   ],
    setTimeout(function(){
        for(var o,t=document.getElementsByTagName("head")[0],r=0,s=n.length;s>r;r++)
            e=document.createElement("link"),
            e.rel="stylesheet",
            "string"==typeof(o=n[r])?e.href=o:g_act.pageCom[o.com]&&(e.href=o.url),
            t.appendChild(e)
        })
}();

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
