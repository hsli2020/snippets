/* Version: 2.2.0 */
/* (c) 2018 Skyogo Studio */
/* Released under the Apache License Version 2.0 */

window.isDebug = false;
window.pluginArr = new Array();
console.log("Safe.js ©Skyogo工作室版权所有");
//Safe.js版本
function safeVersion(){
    return "2.2.0";
}
//Safe.js配置
function safeConfig(thing){
    for(var i in thing){
        if(i=="debug"){
            isDebug = thing[i];
        }
    }
}
//Safe.js初始化
function safeInit(thing,isReactive){
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
    if(typeof thing.el=="string"){
        var safeThingElArr = thing.el.toString().split("&");
        for(var u = 0;u<safeThingElArr.length;u++){
            thing.el = safeThingElArr[u];
            var safeElArr = document.querySelectorAll(thing.el);
            if(thing.copy!=null&&thing.copy!=undefined&&thing.copy!=""){
                if(isReactive){
                    Object.defineProperty(this,"copy",{
                        set: function(newVal){
                            thing.copy = newVal;
                            document.querySelector(thing.el).innerHTML = document.querySelector(thing.copy).innerHTML;
                        }
                    })
                }
                document.querySelector(thing.el).innerHTML = document.querySelector(thing.copy).innerHTML;
            }
            if(thing.var!=null&&thing.var!=""&&thing.var!=undefined){
                if(isReactive){
                    Object.defineProperty(this,"var",{
                        set: function(newVal){
                            thing.var = newVal;
                            for(var a = 0;a<safeElArr.length;a++){
                                safeElArr[a].innerHTML = safeChange(safeElArr[a].innerHTML,thing.var).do();
                            }
                        }
                    })
                }
                for(var a = 0;a<safeElArr.length;a++){
                    safeElArr[a].innerHTML = safeChange(safeElArr[a].innerHTML,thing.var).do();
                }
            }
            if(thing.css!=null&&thing.css!=undefined&&thing.css!=""){
                if(isReactive){
                    Object.defineProperty(this,"css",{
                        set: function(newVal){
                            thing.css = newVal;
                            for(var a = 0;a<safeElArr.length;a++){
                                for(var i in thing.css){
                                    safeElArr[a].style[i] = thing.css[i];
                                }
                            }
                        }
                    })
                }
                for(var a = 0;a<safeElArr.length;a++){
                    for(var i in thing.css){
                        safeElArr[a].style[i] = thing.css[i];
                    }
                }
            }
            if(thing.display!=null&&thing.display!=undefined&&thing.display!=""){
                if(isReactive){
                    Object.defineProperty(this,"display",{
                        set: function(newVal){
                            thing.display = newVal;
                            for(var a = 0;a<safeElArr.length;a++){
                                safeElArr[a].style.display = thing.display;
                            }
                        }
                    })
                }
                for(var a = 0;a<safeElArr.length;a++){
                    safeElArr[a].style.display = thing.display;
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
                if(isReactive){
                    Object.defineProperty(this,"attr",{
                        set: function(newVal){
                            thing.attr = newVal;
                            for(var a = 0;a<safeElArr.length;a++){
                                for(var i in thing.attr){
                                    safeElArr[a].setAttribute(i,thing.attr[i])
                                }
                            }
                        }
                    })
                }
                for(var a = 0;a<safeElArr.length;a++){
                    for(var i in thing.attr){
                        safeElArr[a].setAttribute(i,thing.attr[i])
                    }
                }
            }
            if(thing.click!=null&&thing.click!=undefined&&thing.click!=""){
                if(isReactive){
                    Object.defineProperty(this,"click",{
                        set: function(newVal){
                            thing.click = newVal;
                            for(var a = 0;a<safeElArr.length;a++){
                                if(thing.method!=null&&thing.method!=undefined&&thing.method!=""){
                                    eval("safeElArr["+a+"].onclick =    "+safeChange(thing.click.toString(),thing.method).do());
                                }else{
                                    safeElArr[a].onclick = thing.click;
                                }
                            }
                        }
                    })
                }
                for(var a = 0;a<safeElArr.length;a++){
                    if(thing.method!=null&&thing.method!=undefined&&thing.method!=""){
                        eval("safeElArr["+a+"].onclick =    "+safeChange(thing.click.toString(),thing.method).do());
                    }else{
                        safeElArr[a].onclick = thing.click;
                    }
                }
            }
            if(thing.mousemove!=null&&thing.mousemove!=undefined&&thing.mousemove!=""){
                if(isReactive){
                    Object.defineProperty(this,"mousemove",{
                        set: function(newVal){
                            thing.mousemove = newVal;
                            for(var a = 0;a<safeElArr.length;a++){
                                if(thing.method!=null&&thing.method!=undefined&&thing.method!=""){
                                    eval("safeElArr["+a+"].onmousemove =        "+safeChange(thing.mousemove.toString(),thing.method).do());
                                }else{
                                    safeElArr[a].onmousemove = thing.mousemove;
                                }
                            }
                        }
                    })
                }
                for(var a = 0;a<safeElArr.length;a++){
                    if(thing.method!=null&&thing.method!=undefined&&thing.method!=""){
                        eval("safeElArr["+a+"].onmousemove =        "+safeChange(thing.mousemove.toString(),thing.method).do());
                    }else{
                        safeElArr[a].onmousemove = thing.mousemove;
                    }
                }
            }
            if(thing.mousedown!=null&&thing.mousedown!=undefined&&thing.mousedown!=""){
                if(isReactive){
                    Object.defineProperty(this,"mousedown",{
                        set: function(newVal){
                            thing.mousedown = newVal;
                            for(var a = 0;a<safeElArr.length;a++){
                                if(thing.method!=null&&thing.method!=undefined&&thing.method!=""){
                                    eval("safeElArr["+a+"].onmousedown = "+safeChange(thing.mousedown.toString(),thing.method).do());
                                }else{
                                    safeElArr[a].onmousedown = thing.mousedown;
                                }
                            }
                        }
                    })
                }
                for(var a = 0;a<safeElArr.length;a++){
                    if(thing.method!=null&&thing.method!=undefined&&thing.method!=""){
                        eval("safeElArr["+a+"].onmousedown = "+safeChange(thing.mousedown.toString(),thing.method).do());
                    }else{
                        safeElArr[a].onmousedown = thing.mousedown;
                    }
                }
            }
            if(thing.mouseover!=null&&thing.mouseover!=undefined&&thing.mouseover!=""){
                if(isReactive){
                    Object.defineProperty(this,"mouseover",{
                        set: function(newVal){
                            thing.mouseover = newVal;
                            for(var a = 0;a<safeElArr.length;a++){
                                if(thing.method!=null&&thing.method!=undefined&&thing.method!=""){
                                    eval("safeElArr["+a+"].onmouseover = "+safeChange(thing.mouseover.toString(),thing.method).do());
                                }else{
                                    safeElArr[a].onmouseover = thing.mouseover;
                                }
                            }
                        }
                    })
                }
                for(var a = 0;a<safeElArr.length;a++){
                    if(thing.method!=null&&thing.method!=undefined&&thing.method!=""){
                        eval("safeElArr["+a+"].onmouseover = "+safeChange(thing.mouseover.toString(),thing.method).do());
                    }else{
                        safeElArr[a].onmouseover = thing.mouseover;
                    }
                }
            }
            if(thing.mouseout!=null&&thing.mouseout!=undefined&&thing.mouseout!=""){
                if(isReactive){
                    Object.defineProperty(this,"mouseout",{
                        set: function(newVal){
                            thing.mouseout = newVal;
                            for(var a = 0;a<safeElArr.length;a++){
                                if(thing.method!=null&&thing.method!=undefined&&thing.method!=""){
                                    eval("safeElArr["+a+"].onmouseout = "+safeChange(thing.mouseout.toString(),thing.method).do());
                                }else{
                                    safeElArr[a].onmouseout = thing.mouseout;
                                }
                            }
                        }
                    })
                }
                for(var a = 0;a<safeElArr.length;a++){
                    if(thing.method!=null&&thing.method!=undefined&&thing.method!=""){
                        eval("safeElArr["+a+"].onmouseout = "+safeChange(thing.mouseout.toString(),thing.method).do());
                    }else{
                        safeElArr[a].onmouseout = thing.mouseout;
                    }
                }
            }
            if(thing.html!=null&&thing.html!=undefined&&thing.html!=""){
                if(isReactive){
                    Object.defineProperty(this,"html",{
                        set: function(newVal){
                            thing.html = newVal;
                            for(var a = 0;a<safeElArr.length;a++){
                                safeElArr[a].innerHTML = thing.html;
                            }
                        }
                    })
                }
                for(var a = 0;a<safeElArr.length;a++){
                    safeElArr[a].innerHTML = thing.html;
                }
            }
            if(thing.text!=null&&thing.text!=undefined&&thing.text!=""){
                if(isReactive){
                    Object.defineProperty(this,"text",{
                        set: function(newVal){
                            thing.text = newVal;
                            for(var a = 0;a<safeElArr.length;a++){
                                safeElArr[a].innerText = thing.text;
                            }
                        }
                    })
                }
                for(var a = 0;a<safeElArr.length;a++){
                    safeElArr[a].innerText = thing.text;
                }
            }
            if(thing.class!=null&&thing.class!=undefined&&thing.class!=""){
                if(isReactive){
                    Object.defineProperty(this,"class",{
                        set: function(newVal){
                            thing.class = newVal;
                            for(var a = 0;a<safeElArr.length;a++){
                                safeElArr[a].setAttribute("class",thing.class);
                            }
                        }
                    })
                }
                for(var a = 0;a<safeElArr.length;a++){
                        safeElArr[a].setAttribute("class",thing.class);
                }
            }
            if(thing.name!=null&&thing.name!=undefined&&thing.name!=""){
                if(isReactive){
                    Object.defineProperty(this,"name",{
                        set: function(newVal){
                            thing.name = newVal;
                            for(var a = 0;a<safeElArr.length;a++){
                                safeElArr[a].setAttribute("name",thing.name);
                            }
                        }
                    })
                }
                for(var a = 0;a<safeElArr.length;a++){
                    safeElArr[a].setAttribute("name",thing.name);
                }
            }
        /* 需要thing.el的内容在这个大括号前面写 */    
    }
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
            if(i == "DivisionSign"){
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
    if(this==window){
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