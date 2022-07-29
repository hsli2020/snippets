// ---------------------------------------------------------
jquery UI

* 基本的鼠标互动：
拖拽(drag and dropping)、排序(sorting)、选择(selecting)、缩放(resizing)

* 各种互动效果：
手风琴式的折叠菜单(accordions)、日历(date pickers)、对话框(dialogs)、滑动条(sliders)、表格排序(table sorters)、页签(tabs)
放大镜效果(magnifier)、阴影效果(shadow)

每个 jQuery UI 组件提供一个可串联的标准 jQuery 方法，创建实例时，仅需在 jQuery 对象上调用组件方法。如：

$("#login-form").dialog();	// 创建对话框

* 组件方法可串联使用：

创建可拖动位置、可调整大小的对象
$("#id").draggable().resizable();

创建标签页，设置每5秒自动切换标签
$("#news-panel").tabs().tabs("rotate",5000);
// ---------------------------------------------------------
JavaScript Ext

Backbone.js
Knockout.js
jQuery tmpl模板引擎
jQueryCoreUISelect 
Sisyphus.js
TextExt
Validate.js
jQuery File Upload
Handsontables: Excel-Like Tables For The Web
Pivot.js
Date.js
RequireJS
Grunt.js
JSLint验证
UglifyJS代码压缩
qUnit单元测试
// ---------------------------------------------------------
<form>
  <input type="text" id="msginput" onKeyPress="return handlekeys(this, state, event);" />
</form>

function handlekeys(field, state, e)
{
    var keycode;

    if (window.event) {
        keycode = window.event.keyCode;
    } else if (e) {
        keycode = e.which;
    } else {
        return true;
    }

    if (keycode == 13) {
       state.handle('SEND');
       return false;
    } else {
       return true;
    }
}
// ---------------------------------------------------------
<a href="javascript:void(0)" onClick="return showHideBox();" id="hide_detls" style="display:block;"><@ Hide details @></a>
<a href="javascript:void(0)" onClick="return showHideBox();" id="show_detls" style="display:none;"><@ Show details @></a>
// ---------------------------------------------------------
var isIE  = (navigator.appVersion.indexOf("MSIE") != -1) ? true : false;
var isWin = (navigator.appVersion.toLowerCase().indexOf("win") != -1) ? true : false;
var isOpera = (navigator.userAgent.indexOf("Opera") != -1) ? true : false;
// ---------------------------------------------------------
<script>
function redirectUser(url) {
  var ac = '{:this.affiliateCode:}';
  var keywords = '{:this.keywords:}';
  window.location.href = url + '?ac=' + encodeURIComponent(ac) + '&keywords=' + encodeURIComponent(keywords);
}
</script>

<script>redirectUser('http://www.ashleyrnadison.com/');</script>

// load jquery
<script src="//ajax.googleapis.com/ajax/libs/jquery/1.7.1/jquery.min.js"></script>
<script>window.jQuery || document.write('<script src="{:pinf:static-url shared/javascript/jquery-1.7.1.min.js:}"><\/script>')</script>

@@
header('HTTP/1.1 503 Service Unavailable.', TRUE, '503');
// ---------------------------------------------------------
<script language="JavaScript">
setTimeout('initRedirect()',4000);
function initRedirect() {
  if(document.images) {
    top.location.replace('/app/public/login.p');
  } else {
    top.location.href = '/app/public/login.p';
  }
}
</script>
// ---------------------------------------------------------
iframe autoResize

<script language="JavaScript">
<!--
function autoResize(id){
    var newheight;
    var newwidth;

    newheight=document.getElementById(id).contentWindow.document.body.scrollHeight;
    newwidth=document.getElementById(id).contentWindow.document.body.scrollWidth;

    document.getElementById(id).height= (newheight) + "px";
    document.getElementById(id).width= (newwidth) + "px";
}
//-->
</script>

<iframe src="..." width="100%" height="200px" id="iframe1" marginheight="0" 
 frameborder="0" onLoad="autoResize('iframe1');"></iframe>
--------------------------------------------------------------------------------
iframe(嵌入式帧)自适应高度

填写的嵌入地址一定要和本页面在同一个站点上，否则会提示“拒绝访问！”。对跨域引用有权限问题，请查阅其他资料。

<iframe name="guestbook" src="http://www.site.org/index.asp" 
  scrolling=no width="100%" height="100%" 
  frameborder=no 
  onload="document.all['guestbook'].style.height=guestbook.document.body.scrollHeight">
</iframe>

不过，这建议将<iframe>代码转换成JS放入网页中。
// ---------------------------------------------------------
<div style="padding-left: 4px; padding-right: 4px;" id="debug-footer-summary"
     onMouseOver="this.style.cursor='pointer';"
     onMouseOut="this.style.cursor='auto';"
     onClick="toggleDebugFooterDetail()"></div>

<div id="debug-footer-detail-div" style="display: none; padding-top: 1px;"></div>

<script language="JavaScript">
function toggleDebugFooterDetail() {
  var obj = document.getElementById("debug-footer-detail-div");
  if(obj.style.display=='none') {
    obj.style.display = '';
  } else {
    obj.style.display = 'none';
  }
}
</script>


<a href="javascript:window.close()"
   onMouseOver="window.status='Click to close Window'; return true;"
   onMouseOut="window.status=''; return true;"><@ Close Window @>
</a>
// ---------------------------------------------------------
disable right click on page
<body onContextMenu="return false;">

other way
<script language="JavaScript" type="text/javascript">
  function rightclickerror(e)
  { 
    var ie = (document.all) ? true : false;
      
    if ((ie && event.button==2) || (!ie && e.which==3 )) {
      alert("This feature is disabled");
      return false;
    } 
  }
  if (document.layers) document.captureEvents(Event.MOUSEDOWN);
    document.onmousedown = rightclickerror;
</script>
// ---------------------------------------------------------
在表单中禁止“回车键”的js方法

$('form').keypress(function(e){
  if(e.which === 13){
    e.preventDefault;
    //或者
    return false;
  }
})
// ---------------------------------------------------------
滚动条到底部时自己加载新的内容

<!DOCTYPE html>
<html lang="en">
    <head>
      <meta charset="gb2312">
      <title>滚动条到底部时自己加载新的内容</title>
      <script type='text/javascript' src='js/jquery.js'></script>
        <script type="text/javascript">
          var page_num=2;
          $(document).ready(function(){
            $(window).scroll(function(){
              if($(document).scrollTop()>=$(document).height()-$(window).height()){
                var div1tem = $('#container').height()
                var str =''
                $.ajax({
                    type:"GET",
                    url:'ajaxdata.php',
                    dataType:'json',
                    beforeSend:function(){
                      $('.ajax_loading').show() //显示加载时候的提示
                    },
                    success:function(ret){
                     $(".after_div").before(ret) //将ajxa请求的数据追加到内容里面
                     $('.ajax_loading').hide() //请求成功,隐藏加载提示
                    }
                })
              }
            })
          })
        </script>
    </head>
    <body>
     <div>
        <div style='width:100%;height:1200px'>文章内容</div>
        <div class='after_div'></div>
        <div class='ajax_loading' style='background:#F0F0F0;height:60px;width:100%;text-align:center;line-height:60px;font-size:18px;display:none;position:fixed;bottom:0px'>
            <img src="img/loadinfo.net.gif">数据加载中
        </div>
     </div>
    </body>
</html>
// ---------------------------------------------------------
JavaScript getParameter

function getParameterByName(name) {
    name = name.replace(/[\[]/, "\\\[").replace(/[\]]/, "\\\]");
    var regex = new RegExp("[\\?&]" + name + "=([^&#]*)"),
        results = regex.exec(location.search);
    return results == null ? "" : decodeURIComponent(results[1].replace(/\+/g, " "));
}

(function($) {
    $.QueryString = (function(a) {
        if (a == "") return {};
        var b = {};
        for (var i = 0; i < a.length; ++i)
        {
            var p=a[i].split('=');
            if (p.length != 2) continue;
            b[p[0]] = decodeURIComponent(p[1].replace(/\+/g, " "));
        }
        return b;
    })(window.location.search.substr(1).split('&'))
})(jQuery);

function getParameterByName(name) {
    var match = RegExp('[?&]' + name + '=([^&]*)').exec(window.location.search);
    return match && decodeURIComponent(match[1].replace(/\+/g, ' '));
}
// ---------------------------------------------------------
function $(id) {
  return document.getElementById(id);
}

function $$(sel) {
  return document.querySelectorAll(sel);
}

primitive types are:
• string
• boolean
• number
• null
• undefined

The undefined type has only one value—the value undefined.
The null type has only one value—the value null. 

Anything between single or double quotes is a string value. There’s no difference between single and double quotes.

in JavaScript, there are no floats, ints, doubles, and so on; they are all numbers. 

You can use the operator typeof to find the type of value you’re working with

Remember that typeof is not a function, but an operator.

Note that the typeof operator always returns a string. For example, b has the value undefined, but typeof b returns the string "undefined". It’s common for newcomers to confuse the string "undefined" and the value undefined.  

var b;
b === undefined; // true
b === "undefined"; // false
typeof b === "undefined"; // true 
typeof b === undefined; // false

When you declare a variable but do not initialize it, it’s initialized with undefined. Also, when you have a function that doesn’t return a value explicitly, it returns undefined

Surprisingly, typeof returns "object" when used with a null value: 
    var a = null;
    typeof a; // "object"

Arrays in JavaScript are objects:
    var a = [1, "yes", false, null, undefined, [1, 2, 3]];
    typeof a; // "object"

using === is good practice. 

null == undefined; // true 
"" == 0; // true

null === undefined; // false 
0 === ""; // false

undefined == null; // true 
undefined == 0; // false
// ---------------------------------------------------------
// Define a function
function sum() {
    var res = 0;
    for (var i = 0; i < arguments.length; i++) {
        res += arguments[i];
    }
    return res; 
}

Number.MAX_VALUE; // 1.7976931348623157e+308 
Number.MIN_VALUE; // 5e-324 
Number.POSITIVE_INFINITY; // Infinity 
Number.NEGATIVE_INFINITY; // -Infinity
Number.NaN; // NaN

Math.E;
Math.LN2;
Math.LN10;
Math.LOG2E;
Math.LOG10E; // 0.4342944819032518 
Math.PI; // 3.141592653589793 
Math.SQRT1_2; // 0.7071067811865476 
Math.SQRT2; // 1.4142135623730951
// ---------------------------------------------------------
错误的方法，要引起注意，异步调用

function callAPI(data) {
  var url = 'https://www.ashleymadison.com/api/signup/jp';
  errors = {};

  $.ajax({
    type: 'POST',
    url:   url,
    data:  data,
    headers: { "Accept-Language": "ja_JP" }
  }).done(function(msg) {
    return true;
  }).fail(function(jqXHR) {
    var rsp = $.parseJSON(jqXHR.responseText);
    errors = rsp.errors;
    console.log(errors);
    return false;
  });
  return false;
}
// ---------------------------------------------------------
        $.ajax({
          type:  'POST',
          url:    'https://www-dev.ashleymadison.com/click' + location.search,
          dataType: 'jsonp', // 跨域, FireFox
          success: function(data, status) {
               console.log("Success:");
               console.log("Status: " + status);
               console.log("Data: ");
               console.log(data);
          },
          error: function(request, status, error) {
              console.log("Error:");
              console.log("Status: " + status);
              console.log("Request: ");
              console.log(request);
              console.log("Error: ");
              console.log(error);
          },
          complete: function(httpObj, textStatus){
              console.log("Complete:");
              console.log(httpObj);
              console.log(textStatus);
          },
          async: false // 同步
        });
// ---------------------------------------------------------
function getParameter(paramName) {
  var searchString = window.location.search.substring(1),
      i, val, params = searchString.split("&");

  for (i=0;i<params.length;i++) {
    val = params[i].split("=");
    if (val[0] == paramName) {
      return unescape(val[1]);
    }
  }
  return null;
}

同步ajax调用
var APIURL = 'https://www.ashleymadison.com/api/signup/jp';

function callAPI(data) {
  var rspTxt = $.ajax({
      type:  'POST',
      url:   APIURL,
      data:  data,
      async: false,
      headers: { "Accept-Language": "ja_JP" }
  }).responseText
  var rsp = $.parseJSON(rspTxt);
  return rsp.errors;
}
// ---------------------------------------------------------
var x = [];

// wrong
// -----
// for (var i = 0; i < 3; ++i)
//   x[i] = function () { return i; };

// wrong
// -----
// for (var i = 0; i < 3; ++i) {
//   var j = i;
//   x[j] = function () { return j; };
// }

// good
// -----
for (var i = 0; i < 3; ++i) {
  (function (new_i) {
    x[new_i] = function () { return new_i; };
  })(i);
}

console.log( x[0]() );  // What will these be?
console.log( x[1]() );
console.log( x[2]() );
// ---------------------------------------------------------
JavaScript benchmark: going through an array

  <script type="text/javascript">
    benchTest();

    function benchTest() {
      var arr = [];
      while (arr.length < 500000) {
        arr.push(arr.length);
      }

      console.time("test1");
      for(var i in arr) { }
      console.timeEnd("test1");

      console.time("test2");
      var length = arr.length;
      for(var i=0; i<length; i++) { }
      console.timeEnd("test2");

      console.time("test3");
      var length = arr.length;
      for(var i=length; i--;) { }
      console.timeEnd("test3");
 
      console.time("test4");
      arr.map(function() { });
      console.timeEnd("test4");
  
      console.time("test5");
      arr.filter(function() { });
      console.timeEnd("test5");
   
      console.time("test6");
      arr.forEach(function() { });
      console.timeEnd("test6");
    
      console.time("test7");
      arr.some(function() { });
      console.timeEnd("test7");

      console.time("test8");
      arr.every(function() { return true; });
      console.timeEnd("test8");
    }
  </script>
// ---------------------------------------------------------
Confirmation for Deleting & Uninstalling

<script type="text/javascript">
//-----------------------------------------
// Confirm Actions (delete, uninstall)
//-----------------------------------------
$(document).ready(function(){
    // Confirm Delete
    $('#form').submit(function(){
        if ($(this).attr('action').indexOf('delete',1) != -1) {
            if (!confirm('<?php echo $text_confirm; ?>')) {
                return false;
            }
        }
    });
    // Confirm Uninstall
    $('a').click(function(){
        if ($(this).attr('href') != null && $(this).attr('href').indexOf('uninstall', 1) != -1) {
            if (!confirm('<?php echo $text_confirm; ?>')) {
                return false;
            }
        }
    });
});
</script>
// ---------------------------------------------------------
var el = document.getElementById("ID3");
     
el.style.position = 'absolute';
el.style.left = event.clientX  + "px";
el.style.top = event.clientY + "px";
 
if (el.style.display=="block") {
     el.style.display = "none";
} else {
     el.style.display = "block";
}

$('[data-toggle="popover"]').popover('show');
// ---------------------------------------------------------
function pr(x) {
    if (typeof(x) == 'undefined') {
        console.log('');
    } else {
        console.log(x);
    }
}

function generateUUID() {
    var d = new Date().getTime();
    var uuid = 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
        var r = (d + Math.random()*16)%16 | 0;
        d = Math.floor(d/16);
        return (c=='x' ? r : (r&0x3|0x8)).toString(16);
    });
    return uuid;
};
//pr(generateUUID());

function guid() {
  function s4() {
    return Math.floor((1 + Math.random()) * 0x10000).toString(16).substring(1);
  }
  return s4() + s4() + '-' + s4() + '-' + s4() + '-' + s4() + '-' + s4() + s4() + s4();
}
//pr(guid());

typeof "John"                // Returns string
typeof 3.14                  // Returns number
typeof false                 // Returns boolean
typeof [1,2,3,4]             // Returns object
typeof {name:'John', age:34} // Returns object 
typeof null                  // object !!
typeof undefined             // undefined
// ---------------------------------------------------------
<body onload="window.print()">

function getQueryVariable(variable)
{
       var query = window.location.search.substring(1);
       var vars = query.split("&");
       for (var i=0;i<vars.length;i++) {
               var pair = vars[i].split("=");
               if(pair[0] == variable){return pair[1];}
       }
      return(false);
}
// ---------------------------------------------------------

// ---------------------------------------------------------

// ---------------------------------------------------------

// ---------------------------------------------------------

// ---------------------------------------------------------

// ---------------------------------------------------------

// ---------------------------------------------------------

// ---------------------------------------------------------

// ---------------------------------------------------------

// ---------------------------------------------------------
