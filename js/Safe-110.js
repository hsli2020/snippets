/* Safe.js Powered By Skyogo */
/* Version: 1.1.0 */

window.isDebug = false;
window.pluginArr = new Array();
console.log("Safe.js ©Skyogo工作室版权所有");
//Safe.js版本
function safeVersion(){
    return "1.1.0";
}
//Safe.js配置
function safeConfig(thing){
    for(var i in thing){
        if(i==="debug"){
            isDebug = thing[i];
        }
    }
}
//Safe.js初始化
function safeInit(thing){
    if(isDebug){
        console.time("Safe执行时间：");
    }
    if(pluginArr.length > 0){
        for(var i=0;i<pluginArr.length;i++){
            try{
                eval(pluginArr[i]+"Do(thing)");
            }catch(err){
                if(isDebug){
                    console.error("Safe: Cannot use plugin! AT plugin name: "+pluginArr[i]);
                }
                break;
            }
        }
    }
    if(typeof thing.el==="string"){
        var safeThingElArr = thing.el.toString().split("&");
        for(var u = 0;u<safeThingElArr.length;u++){
            thing.el = safeThingElArr[u];
            if(thing.copy!=null&&thing.copy!=undefined&&thing.copy!=""){
                document.querySelector(thing.el).innerHTML = document.querySelector(thing.copy).innerHTML;
            }
            if(thing.var!=null&&thing.var!=""&&thing.var!=undefined){
                for(var a = 0;a<document.querySelectorAll(thing.el).length;a++){
                    document.querySelectorAll(thing.el)[a].innerHTML = safeChange(document.querySelectorAll(thing.el)[a].innerHTML,thing.var).do();
                }
            }
            if(thing.css!=null&&thing.css!=undefined&&thing.css!=""){
                for(var a = 0;a<document.querySelectorAll(thing.el).length;a++){
                    for(var i in thing.css){
                        document.querySelectorAll(thing.el)[a].style[i] = thing.css[i];
                    }
                }
            }
            if(thing.display!=null&&thing.display!=undefined&&thing.display!=""){
                for(var a = 0;a<document.querySelectorAll(thing.el).length;a++){
                    document.querySelectorAll(thing.el)[a].style.display = thing.display;
                }
            }
            if(document.querySelector(thing.el).innerHTML.indexOf("-import")!=-1){
                var safeElInner = document.querySelector(thing.el).innerHTML;
                var xmlhttp;
                if (window.XMLHttpRequest){
                    // code for IE7+, Firefox, Chrome, Opera, Safari
                    xmlhttp=new XMLHttpRequest();
                }else{
                    // code for IE6, IE5
                    xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
                }
                for(var i = 0;i<safeElInner.length;i++){
                    if(safeElInner.indexOf("-import",i)!=-1){
                        i=safeElInner.indexOf("-import",i)+8;
                        var safeImportStringFindVal =   safeElInner.substr(i,safeElInner.indexOf(")",i)-i);
                        xmlhttp.onreadystatechange=function(){
                            if (xmlhttp.readyState==4 && xmlhttp.status==200){
                                document.querySelector(thing.el).innerHTML = xmlhttp.responseText;
                            }
                        }
                        eval("xmlhttp.open('POST',"+safeImportStringFindVal+",true)");
                        xmlhttp.send();
                    }else{
                        break;
                    }
                }
            }
            if(thing.import!=null&&thing.import!=undefined&&thing.import!=""){
                var xmlhttp;
                if (window.XMLHttpRequest){
                    xmlhttp=new XMLHttpRequest();
                }else{
                    xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
                }
                xmlhttp.onreadystatechange=function(){
                    if (xmlhttp.readyState==4 && xmlhttp.status==200){
                        document.querySelector(thing.el).innerHTML = xmlhttp.responseText;
                    }
                }
                xmlhttp.open('POST',thing.import,true);
                xmlhttp.send();
            }
            if(thing.attr!=null&&thing.attr!=undefined&&thing.attr!=""){
                for(var a = 0;a<document.querySelectorAll(thing.el).length;a++){
                    for(var i in thing.attr){
                        document.querySelectorAll(thing.el)[a].setAttribute(i,thing.attr[i])
                    }
                }
            }
            if(thing.click!=null&&thing.click!=undefined&&thing.click!=""){
                for(var a = 0;a<document.querySelectorAll(thing.el).length;a++){
                    if(thing.method!=null&&thing.method!=undefined&&thing.method!=""){
                        eval("document.querySelectorAll(thing.el)["+a+"].onclick =    "+safeChange(thing.click.toString(),thing.method).do());
                    }else{
                        eval("document.querySelectorAll(thing.el)["+a+"].onclick = "+thing.click);
                    }
                }
            }
            if(thing.mousemove!=null&&thing.mousemove!=undefined&&thing.mousemove!=""){
                for(var a = 0;a<document.querySelectorAll(thing.el).length;a++){
                    if(thing.method!=null&&thing.method!=undefined&&thing.method!=""){
                        eval("document.querySelectorAll(thing.el)["+a+"].onmousemove =        "+safeChange(thing.mousemove.toString(),thing.method).do());
                    }else{
                        eval("document.querySelectorAll(thing.el)["+a+"].onmousemove =  "+thing.mousemove);
                    }
                }
            }
            if(thing.mousedown!=null&&thing.mousedown!=undefined&&thing.mousedown!=""){
                for(var a = 0;a<document.querySelectorAll(thing.el).length;a++){
                    if(thing.method!=null&&thing.method!=undefined&&thing.method!=""){
                        eval("document.querySelectorAll(thing.el)["+a+"].onmousedown = "+safeChange(thing.mousedown.toString(),thing.method).do());
                    }else{
                        eval("document.querySelectorAll(thing.el)["+a+"].onmousedown = "+thing.mousedown);
                    }
                }
            }
            if(thing.mouseover!=null&&thing.mouseover!=undefined&&thing.mouseover!=""){
                for(var a = 0;a<document.querySelectorAll(thing.el).length;a++){
                    if(thing.method!=null&&thing.method!=undefined&&thing.method!=""){
                        eval("document.querySelectorAll(thing.el)["+a+"].onmouseover = "+safeChange(thing.mouseover.toString(),thing.method).do());
                    }else{
                        eval("document.querySelectorAll(thing.el)["+a+"].onmouseover =  "+thing.mouseover);
                    }
                }
            }
            if(thing.mouseout!=null&&thing.mouseout!=undefined&&thing.mouseout!=""){
                for(var a = 0;a<document.querySelectorAll(thing.el).length;a++){
                    if(thing.method!=null&&thing.method!=undefined&&thing.method!=""){
                        eval("document.querySelectorAll(thing.el)["+a+"].onmouseout = "+safeChange(thing.mouseout.toString(),thing.method).do());
                    }else{
                        eval("document.querySelectorAll(thing.el)["+a+"].onmouseout = "+thing.mouseout);
                    }
                }
            }
            if(thing.cut!=null&&thing.cut!=undefined&&thing.cut!=""){
                document.querySelector(thing.el).innerHTML = document.querySelector(thing.cut).innerHTML;
                document.querySelector(thing.cut).innerHTML = "";
            }
    }
    /* 需要thing.el的内容在这个大括号前面写 */    
    }
    /* CallBack永远放到最后 */
    if(thing.callback!=null&&thing.callback!=undefined&&thing.callback!=""){
        if(thing.method!=null&&thing.method!=undefined&&thing.method!=""){
            eval("("+safeChange(thing.callback.toString(),thing.method).do()+")()");
        }else{
            eval("("+thing.callback+")()");
        }
    }
    if(isDebug){
        console.timeEnd("Safe执行时间：");
    }
}
//Safe.js模板方法
function safeChange(content,varList){
    var safeChangeDivisionSign = "";
    this.content=content;
    this.varList=varList;
    this.do=function(){
        for(var i in this.varList){
            if(i === "DivisionSign"){
                safeChangeDivisionSign = this.varList[i];
            }
            try{
                var safeChangeString = this.varList[i].join(safeChangeDivisionSign);
                this.content=this.content.replace(new RegExp("-##"+i+"-##","g"),safeChangeString);
            }catch(err){}
            this.content=this.content.replace(new RegExp("##"+i+"##","g"),this.varList[i].toString());
            try{
                var safeChangeString = this.varList[i].join(safeChangeDivisionSign);
                this.content=this.content.replace(new RegExp("-#"+i+"-#","g"),safeChangeString.replace(/&/g,"&amp;").replace(/</g,"&lt;").replace(/\'/g,"&apos;").replace(/\"/g,"&quot;").replace(/>/g,"&gt;"));
            }catch(err){}
            this.content=this.content.replace(new RegExp("#"+i+"#","g"),this.varList[i].toString().replace(/&/g,"&amp;").replace(/</g,"&lt;").replace(/\'/g,"&apos;").replace(/\"/g,"&quot;").replace(/>/g,"&gt;"));
        }
        return this.content;
    }
    if(this===window){
        return new safeChange(content,varList);
    }
}
//Safe.js创建组件方法
function safeComponent(name,html){
    if(name!=null&&name!=undefined&&name!=""&&html!=null&&html!=undefined&&html!=""){
        var safeComponentElArr = document.getElementsByTagName(name);
        for(var i=0;i<safeComponentElArr.length;i++){
            safeComponentElArr[i].innerHTML = html;
        }
    }else{
        if(isDebug){
            console.error("Safe: name or template is null.");
        }
    }
}
//安装插件方法
function safeUse(pluginName){
    pluginArr[pluginArr.length] = pluginName;
}
//获取所有插件方法
function safeGetAllPlugin(){
    return pluginArr;
}