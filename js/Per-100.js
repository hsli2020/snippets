/* Per.js Community */
/* Version: 1.0 */
/* (c) 2018 Skyogo Studio */
/* Released under the Apache License Version 2.0 */

(function(window,undefined){
    var allVersion = 1.0;
    var pluginArr = new Array();
    var usedPluginArr = ["Per","Per.do","Per.version","Per.joinModule","Per.use","Per.getAllModuleName","Per.getAllModuleVersion","Per.isThisModuleUsed"];
    window.Per = function(){
        this.version = allVersion;
        this.do = function(Obj,isReactive){
            if(isReactive == null||isReactive == undefined||isReactive == ""){
                isReactive = false;
            }
            if(typeof Obj == "object"&&typeof isReactive == "boolean"){
                if(typeof Obj.el=="string"&&Obj.el!=""){
                    var Element = document.querySelectorAll(Obj.el);
                    if(typeof Obj.copy=="string"){
                        if(isReactive){
                            Object.defineProperty(this.do,"copy",{
                                set: function(newVal){
                                    Obj.copy = newVal;
                                    document.querySelector(Obj.el).innerHTML =   document.querySelector(Obj.copy).innerHTML;
                                },
                                get: function(){
                                    return Obj.copy;
                                }
                            })
                        }
                        document.querySelector(Obj.el).innerHTML =   document.querySelector(Obj.copy).innerHTML;
                    }
                    if(typeof Obj.html=="string"){
                        if(isReactive){
                            Object.defineProperty(this.do,"html",{
                                set: function(newVal){
                                    Obj.html = newVal;
                                    for(var i=0;i<Element.length;i++){
                                        Element[i].innerHTML = Obj.html;
                                    }
                                },
                                get: function(){
                                    return Obj.html;
                                }
                            })
                        }
                        for(var i=0;i<Element.length;i++){
                            Element[i].innerHTML = Obj.html;
                        }
                    }
                    if(typeof Obj.text=="string"){
                        if(isReactive){
                            Object.defineProperty(this.do,"text",{
                                set: function(newVal){
                                    Obj.text = newVal;
                                    for(var i=0;i<Element.length;i++){
                                        Element[i].innerText = Obj.text;
                                    }
                                },
                                get: function(){
                                    return Obj.text;
                                }
                            })
                        }
                        for(var i=0;i<Element.length;i++){
                            Element[i].innerText = Obj.text;
                        }
                    }
                    if(typeof Obj.val=="string"){
                        if(isReactive){
                            Object.defineProperty(this.do,"val",{
                                set: function(newVal){
                                    Obj.val = newVal;
                                    for(var i=0;i<Element.length;i++){
                                        Element[i].value = Obj.val;
                                    }
                                },
                                get: function(){
                                    return Obj.val;
                                }
                            })
                        }
                        for(var i=0;i<Element.length;i++){
                            Element[i].value = Obj.val;
                        }
                    }
                    if(typeof Obj.css=="object"){
                        if(isReactive){
                            Object.defineProperty(this.do,"css",{
                                set: function(newVal){
                                    Obj.css = newVal;
                                    var cssObjArr = getObjKeyAndVal(Obj.css);
                                    for(var a=0;a<cssObjArr.length;a++){
                                        if(a%2 == 0){
                                            for(var i=0;i<Element.length;i++){
                                                Element[i].style[cssObjArr[a]] = cssObjArr[++a];
                                            }
                                        }
                                    }
                                },
                                get: function(){
                                    return Obj.css;
                                }
                            })
                        }
                        var cssObjArr = getObjKeyAndVal(Obj.css);
                        for(var a=0;a<cssObjArr.length;a++){
                            if(a%2 == 0){
                                for(var i=0;i<Element.length;i++){
                                    Element[i].style[cssObjArr[a]] = cssObjArr[++a];
                                }
                            }
                        }
                    }
                    if(typeof Obj.attr=="object"){
                        if(isReactive){
                            Object.defineProperty(this.do,"attr",{
                                set: function(newVal){
                                    Obj.attr = newVal;
                                    var attrObjArr = getObjKeyAndVal(Obj.attr);
                                    for(var a=0;a<attrObjArr.length;a++){
                                        if(a%2 == 0){
                                            for(var i=0;i<Element.length;i++){
                                                Element[i].setAttribute(attrObjArr[a],attrObjArr[++a]);
                                            }
                                        }
                                    }
                                },
                                get: function(){
                                    return Obj.attr;
                                }
                            })
                        }
                        var attrObjArr = getObjKeyAndVal(Obj.attr);
                        for(var a=0;a<attrObjArr.length;a++){
                            if(a%2 == 0){
                                for(var i=0;i<Element.length;i++){
                                    Element[i].setAttribute(attrObjArr[a],attrObjArr[++a]);
                                }
                            }
                        }
                    }
                    if(Array.isArray(Obj.for)){
                        if(isReactive){
                            Object.defineProperty(this.do,"for",{
                                set: function(newVal){
                                    Obj.for = newVal;
                                    for(var i=0;i<Element.length;i++){
                                        var forInnerHTML = "";
                                        for(var a=0;a<Obj.for.length-1;a++){
                                            if(a==Obj.for.length-2){
                                                forInnerHTML += "<li>"+Obj.for[a]+"</li>";
                                                break;
                                            }
                                            forInnerHTML += "<li>"+Obj.for[a]+"</li>"+Obj.for[Obj.for.length-1];
                                        }
                                        Element[i].innerHTML += forInnerHTML;
                                    }
                                },
                                get: function(){
                                    return Obj.for;
                                }
                            })
                        }
                        //数组最后一位为遍历中间插入的内容，如果不需要可以留空
                        for(var i=0;i<Element.length;i++){
                            var forInnerHTML = "";
                            for(var a=0;a<Obj.for.length-1;a++){
                                if(a==Obj.for.length-2){
                                    forInnerHTML += "<li>"+Obj.for[a]+"</li>";
                                    break;
                                }
                                forInnerHTML += "<li>"+Obj.for[a]+"</li>"+Obj.for[Obj.for.length-1];
                            }
                            Element[i].innerHTML += forInnerHTML;
                        }
                    }
                    if(typeof Obj.import=="string"){
                        if(isReactive){
                            Object.defineProperty(this.do,"import",{
                                set: function(newVal){
                                    Obj.import = newVal;
                                    if(!isThisModuleUsed("Per.ajax")){
                                        use(["Per.ajax"]);
                                    }
                                    Per().ajax("POST",Obj.import,"",true,function(html){
                                        for(var i=0;i<Element.length;i++){
                                            Element[i].innerHTML = html;
                                        }
                                    });
                                },
                                get: function(){
                                    return Obj.import;
                                }
                            })
                        }
                        if(!isThisModuleUsed("Per.ajax")){
                            use(["Per.ajax"]);
                        }
                        Per().ajax("POST",Obj.import,"",true,function(html){
                            for(var i=0;i<Element.length;i++){
                                Element[i].innerHTML = html;
                            }
                        });
                    }
                    //当以function类型初始化时，你之后重新更改值将不能为method，反之亦然
                    //method属性不支持响应式
                    if(typeof Obj.click=="function"){
                        if(isReactive){
                            Object.defineProperty(this.do,"click",{
                                set: function(newVal){
                                    Obj.click = newVal;
                                    for(var i=0;i<Element.length;i++){
                                        Element[i].onclick = Obj.click;
                                    }
                                },
                                get: function(){
                                    return Obj.click;
                                }
                            })
                        }
                        for(var i=0;i<Element.length;i++){
                            Element[i].onclick = Obj.click;
                        }
                    }else if(typeof Obj.click=="string"){
                        if(isReactive){
                            Object.defineProperty(this.do,"click",{
                                set: function(newVal){
                                    Obj.click = newVal;
                                    var MethodArr = getObjKeyAndVal(Obj.method);
                                    for(var a=0;a<MethodArr.length;a++){
                                        if(MethodArr[a] == Obj.click){
                                            for(var i=0;i<Element.length;i++){
                                                Element[i].onclick = MethodArr[a+1];
                                            }
                                            break;
                                        }
                                    }
                                },
                                get: function(){
                                    return Obj.click;
                                }
                            })
                        }
                        var MethodArr = getObjKeyAndVal(Obj.method);
                        for(var a=0;a<MethodArr.length;a++){
                            if(MethodArr[a] == Obj.click){
                                for(var i=0;i<Element.length;i++){
                                    Element[i].onclick = MethodArr[a+1];
                                }
                                break;
                            }
                        }
                    }
                    if(typeof Obj.mousemove=="function"){
                        if(isReactive){
                            Object.defineProperty(this.do,"mousemove",{
                                set: function(newVal){
                                    Obj.mousemove = newVal;
                                    for(var i=0;i<Element.length;i++){
                                        Element[i].onmousemove = Obj.mousemove;
                                    }
                                },
                                get: function(){
                                    return Obj.mousemove;
                                }
                            })
                        }
                        for(var i=0;i<Element.length;i++){
                            Element[i].onmousemove = Obj.mousemove;
                        }
                    }else if(typeof Obj.mousemove=="string"){
                        if(isReactive){
                            Object.defineProperty(this.do,"mousemove",{
                                set: function(newVal){
                                    Obj.mousemove = newVal;
                                    var MethodArr = getObjKeyAndVal(Obj.method);
                                    for(var a=0;a<MethodArr.length;a++){
                                        if(MethodArr[a] == Obj.mousemove){
                                            for(var i=0;i<Element.length;i++){
                                                Element[i].onmousemove = MethodArr[a+1];
                                            }
                                            break;
                                        }
                                    }
                                },
                                get: function(){
                                    return Obj.mousemove;
                                }
                            })
                        }
                        var MethodArr = getObjKeyAndVal(Obj.method);
                        for(var a=0;a<MethodArr.length;a++){
                            if(MethodArr[a] == Obj.mousemove){
                                for(var i=0;i<Element.length;i++){
                                    Element[i].onmousemove = MethodArr[a+1];
                                }
                                break;
                            }
                        }
                    }
                    if(typeof Obj.mousedown=="function"){
                        if(isReactive){
                            Object.defineProperty(this.do,"mousedown",{
                                set: function(newVal){
                                    Obj.mousedown = newVal;
                                    for(var i=0;i<Element.length;i++){
                                        Element[i].onmousedown = Obj.mousedown;
                                    }
                                },
                                get: function(){
                                    return Obj.mousedown;
                                }
                            })
                        }
                        for(var i=0;i<Element.length;i++){
                            Element[i].onmousedown = Obj.mousedown;
                        }
                    }else if(typeof Obj.mousedown=="string"){
                        if(isReactive){
                            Object.defineProperty(this.do,"mousedown",{
                                set: function(newVal){
                                    Obj.mousedown = newVal;
                                    var MethodArr = getObjKeyAndVal(Obj.method);
                                    for(var a=0;a<MethodArr.length;a++){
                                        if(MethodArr[a] == Obj.mousedown){
                                            for(var i=0;i<Element.length;i++){
                                                Element[i].onmousedown = MethodArr[a+1];
                                            }
                                            break;
                                        }
                                    }
                                },
                                get: function(){
                                    return Obj.mousedown;
                                }
                            })
                        }
                        var MethodArr = getObjKeyAndVal(Obj.method);
                        for(var a=0;a<MethodArr.length;a++){
                            if(MethodArr[a] == Obj.mousedown){
                                for(var i=0;i<Element.length;i++){
                                    Element[i].onmousedown = MethodArr[a+1];
                                }
                                break;
                            }
                        }
                    }
                    if(typeof Obj.mouseover=="function"){
                        if(isReactive){
                            Object.defineProperty(this.do,"mouseover",{
                                set: function(newVal){
                                    Obj.mouseover = newVal;
                                    for(var i=0;i<Element.length;i++){
                                        Element[i].onmouseover = Obj.mouseover;
                                    }
                                },
                                get: function(){
                                    return Obj.mouseover;
                                }
                            })
                        }
                        for(var i=0;i<Element.length;i++){
                            Element[i].onmouseover = Obj.mouseover;
                        }
                    }else if(typeof Obj.mouseover=="string"){
                        if(isReactive){
                            Object.defineProperty(this.do,"mouseover",{
                                set: function(newVal){
                                    Obj.mouseover = newVal;
                                    var MethodArr = getObjKeyAndVal(Obj.method);
                                    for(var a=0;a<MethodArr.length;a++){
                                        if(MethodArr[a] == Obj.mouseover){
                                            for(var i=0;i<Element.length;i++){
                                                Element[i].onmouseover = MethodArr[a+1];
                                            }
                                            break;
                                        }
                                    }
                                },
                                get: function(){
                                    return Obj.mouseover;
                                }
                            })
                        }
                        var MethodArr = getObjKeyAndVal(Obj.method);
                        for(var a=0;a<MethodArr.length;a++){
                            if(MethodArr[a] == Obj.mouseover){
                                for(var i=0;i<Element.length;i++){
                                    Element[i].onmouseover = MethodArr[a+1];
                                }
                                break;
                            }
                        }
                    }
                    if(typeof Obj.mouseout=="function"){
                        if(isReactive){
                            Object.defineProperty(this.do,"mouseout",{
                                set: function(newVal){
                                    Obj.mouseout = newVal;
                                    for(var i=0;i<Element.length;i++){
                                        Element[i].onmouseout = Obj.mouseout;
                                    }
                                },
                                get: function(){
                                    return Obj.mouseout;
                                }
                            })
                        }
                        for(var i=0;i<Element.length;i++){
                            Element[i].onmouseout = Obj.mouseout;
                        }
                    }else if(typeof Obj.mouseout=="string"){
                        if(isReactive){
                            Object.defineProperty(this.do,"mouseout",{
                                set: function(newVal){
                                    Obj.mouseout = newVal;
                                    var MethodArr = getObjKeyAndVal(Obj.method);
                                    for(var a=0;a<MethodArr.length;a++){
                                        if(MethodArr[a] == Obj.mouseout){
                                            for(var i=0;i<Element.length;i++){
                                                Element[i].onmouseout = MethodArr[a+1];
                                            }
                                            break;
                                        }
                                    }
                                },
                                get: function(){
                                    return Obj.mouseout;
                                }
                            })
                        }
                        var MethodArr = getObjKeyAndVal(Obj.method);
                        for(var a=0;a<MethodArr.length;a++){
                            if(MethodArr[a] == Obj.mouseout){
                                for(var i=0;i<Element.length;i++){
                                    Element[i].onmouseout = MethodArr[a+1];
                                }
                                break;
                            }
                        }
                    }
                    if(typeof Obj.mouseup=="function"){
                        if(isReactive){
                            Object.defineProperty(this.do,"mouseup",{
                                set: function(newVal){
                                    Obj.mouseup = newVal;
                                    for(var i=0;i<Element.length;i++){
                                        Element[i].onmouseup = Obj.mouseup;
                                    }
                                },
                                get: function(){
                                    return Obj.mouseup;
                                }
                            })
                        }
                        for(var i=0;i<Element.length;i++){
                            Element[i].onmouseup = Obj.mouseup;
                        }
                    }else if(typeof Obj.mouseup=="string"){
                        if(isReactive){
                            Object.defineProperty(this.do,"mouseup",{
                                set: function(newVal){
                                    Obj.mouseup = newVal;
                                    var MethodArr = getObjKeyAndVal(Obj.method);
                                    for(var a=0;a<MethodArr.length;a++){
                                        if(MethodArr[a] == Obj.mouseup){
                                            for(var i=0;i<Element.length;i++){
                                                Element[i].onmouseup = MethodArr[a+1];
                                            }
                                            break;
                                        }
                                    }
                                },
                                get: function(){
                                    return Obj.mouseup;
                                }
                            })
                        }
                        var MethodArr = getObjKeyAndVal(Obj.method);
                        for(var a=0;a<MethodArr.length;a++){
                            if(MethodArr[a] == Obj.mouseup){
                                for(var i=0;i<Element.length;i++){
                                    Element[i].onmouseup = MethodArr[a+1];
                                }
                                break;
                            }
                        }
                    }
                    if(typeof Obj.class == "string"){
                        if(isReactive){
                            Object.defineProperty(this.do,"class",{
                                set: function(newVal){
                                    Obj.class = newVal;
                                    for(var i=0;i<Element.length;i++){
                                        Element[i].setAttribute("class",Obj.class);
                                    }
                                },
                                get: function(){
                                    return Obj.class;
                                }
                            })
                        }
                        for(var i=0;i<Element.length;i++){
                            Element[i].setAttribute("class",Obj.class);
                        }
                    }
                    if(typeof Obj.name == "string"){
                        if(isReactive){
                            Object.defineProperty(this.do,"name",{
                                set: function(newVal){
                                    Obj.name = newVal;
                                    for(var i=0;i<Element.length;i++){
                                        Element[i].setAttribute("name",Obj.name);
                                    }
                                },
                                get: function(){
                                    return Obj.name;
                                }
                            })
                        }
                        for(var i=0;i<Element.length;i++){
                            Element[i].setAttribute("name",Obj.name);
                        }
                    }
                    //bindType用来设置绑定的值，例如innerHTML，innerText，value等，必须写
                    if(typeof Obj.bind == "string"&&typeof Obj.bindType == "string"){
                        if(isReactive){
                            Object.defineProperty(this.do,"bindType",{
                                set: function(newVal){
                                    Obj.bindType = newVal;
                                },
                                get: function(){
                                    return Obj.bindType;
                                }
                            });
                            Object.defineProperty(this.do,"bind",{
                                set: function(newVal){
                                    Obj.bind = newVal;
                                    if(Obj.bindType == "html"){
                                        for(var i=0;i<Element.length;i++){
                                            var bindHTMLListener = Element[i];
                                            bindHTMLListener.paraEl = i;
                                            bindHTMLListener.targetEl = document.querySelectorAll(Obj.bind);
                                            bindHTMLListener.oninput = function(){
                                                var elementList = this.targetEl;
                                                for(var a=0;a<elementList.length;a++){
                                                    elementList[a].innerHTML = Element[this.paraEl].value;
                                                }
                                            }
                                        }
                                    }else if(Obj.bindType == "text"){
                                        for(var i=0;i<Element.length;i++){
                                            var bindHTMLListener = Element[i];
                                            bindHTMLListener.paraEl = i;
                                            bindHTMLListener.targetEl = document.querySelectorAll(Obj.bind);
                                            bindHTMLListener.oninput = function(){
                                                var elementList = this.targetEl;
                                                for(var a=0;a<elementList.length;a++){
                                                    elementList[a].innerText = Element[this.paraEl].value;
                                                }
                                            }
                                        }
                                    }else if(Obj.bindType == "value"){
                                        for(var i=0;i<Element.length;i++){
                                            var bindHTMLListener = Element[i];
                                            bindHTMLListener.paraEl = i;
                                            bindHTMLListener.targetEl = document.querySelectorAll(Obj.bind);
                                            bindHTMLListener.oninput = function(){
                                                var elementList = this.targetEl;
                                                for(var a=0;a<elementList.length;a++){
                                                    elementList[a].value = Element[this.paraEl].value;
                                                }
                                            }
                                        }
                                    }else{
                                        console.error("Per.js: unknow bindType!");
                                    }
                                },
                                get: function(){
                                    return Obj.bind;
                                }
                            })
                        }
                        if(Obj.bindType == "html"){
                            for(var i=0;i<Element.length;i++){
                                var bindHTMLListener = Element[i];
                                bindHTMLListener.paraEl = i;
                                bindHTMLListener.targetEl = document.querySelectorAll(Obj.bind);
                                bindHTMLListener.oninput = function(){
                                    var elementList = this.targetEl;
                                    for(var a=0;a<elementList.length;a++){
                                        elementList[a].innerHTML = Element[this.paraEl].value;
                                    }
                                }
                            }
                        }else if(Obj.bindType == "text"){
                            for(var i=0;i<Element.length;i++){
                                var bindHTMLListener = Element[i];
                                bindHTMLListener.paraEl = i;
                                bindHTMLListener.targetEl = document.querySelectorAll(Obj.bind);
                                bindHTMLListener.oninput = function(){
                                    var elementList = this.targetEl;
                                    for(var a=0;a<elementList.length;a++){
                                        elementList[a].innerText = Element[this.paraEl].value;
                                    }
                                }
                            }
                        }else if(Obj.bindType == "value"){
                            for(var i=0;i<Element.length;i++){
                                var bindHTMLListener = Element[i];
                                bindHTMLListener.paraEl = i;
                                bindHTMLListener.targetEl = document.querySelectorAll(Obj.bind);
                                bindHTMLListener.oninput = function(){
                                    var elementList = this.targetEl;
                                    for(var a=0;a<elementList.length;a++){
                                        elementList[a].value = Element[this.paraEl].value;
                                    }
                                }
                            }
                        }else{
                            console.error("Per.js: unknow bindType!");
                        }
                    }
                    //callback永远放到最后执行
                    if(typeof Obj.callback == "function"){
                        Obj.callback();
                    }else if(typeof Obj.callback == "string"){
                        var callbackMethodArr = getObjKeyAndVal(Obj.method);
                        for(var i=0;i<callbackMethodArr.length;i++){
                            if(callbackMethodArr[i] == Obj.callback){
                                callbackMethodArr[i+1]();
                                break;
                            }
                        }
                    }
                }else{
                    //el参数无值
                    console.error("Per.js: para el cannot be null!");
                }
            }
        }
        this.joinModule = function(moduleName,moduleVersion,moduleFunction){
            //Per.js官方规定，moduleName里面的结构路径分隔符应为"."，例如dom.ajax，则为dom包下面的ajax包。
            //结构路径可为多重包，例如：dom.do.test
            if(typeof moduleName == "string"&&typeof moduleFunction == "function"&&typeof moduleVersion == "number"){
                var allModule = getAllModuleName();
                for(var i=0;i<allModule.length;i++){
                    if(moduleName == allModule[i]){
                        console.error("Per.js: please do not rejoin the module!");
                        return;
                    }
                }
                pluginArr[pluginArr.length] = moduleName;
                pluginArr[pluginArr.length] = moduleVersion;
                pluginArr[pluginArr.length] = moduleFunction;
            }
        }
        this.use = function(moduleName){
            if(Array.isArray(moduleName)){
                for(var a=0;a<moduleName.length;a++){
                    if(isThisModuleUsed(moduleName[a])){
                        console.error("Per.js: please do not reuse the module! at module "+moduleName[a]);
                    }else{
                        for(var i=0;i<pluginArr.length;i++){
                            if(i%3 == 0&&pluginArr[i] == moduleName[a]){
                                pluginArr[i+2]();
                                usedPluginArr[usedPluginArr.length] = pluginArr[i];
                                break;
                            }
                        }
                    }
                }
            }else if(typeof moduleName == "string"&&moduleName == "all"){
                moduleName = getAllModuleName();
                for(var a=0;a<moduleName.length;a++){
                    if(isThisModuleUsed(moduleName[a])){
                        console.error("Per.js: please do not reuse the module! at module "+moduleName[a]);
                    }else{
                        for(var i=0;i<pluginArr.length;i++){
                            if(i%3 == 0&&pluginArr[i] == moduleName[a]){
                                pluginArr[i+2]();
                                usedPluginArr[usedPluginArr.length] = pluginArr[i];
                                break;
                            }
                        }
                    }
                }
            }
        }
        this.getAllModuleName = function(){
            var allModuleArr = new Array();
            for(var i=0;i<pluginArr.length;i++){
                if(i%3 == 0){
                    allModuleArr[allModuleArr.length] = pluginArr[i];
                }
            }
            return allModuleArr;
        }
        this.getAllModuleVersion = function(){
            var allModuleArr = new Array();
            for(var i=0;i<pluginArr.length;i++){
                if(i%3 == 0){
                    allModuleArr[allModuleArr.length] = pluginArr[i+1];
                }
            }
            return allModuleArr;
        }
        this.isThisModuleUsed = function(moduleName){
            if(typeof moduleName == "string"){
                for(var i=0;i<usedPluginArr.length;i++){
                    if(usedPluginArr[i] == moduleName){
                        return true;
                    }
                }
                return false;
            }
        }
        if(this == window){
            return new Per();
        }
    }
    Per().joinModule("Per.component",allVersion, function(){
        var componentArr = new Array();
        window.Per.prototype.component = function(){
            return {
                set: function(componentName, template){
                    componentArr[componentArr.length] = componentName;
                    componentArr[componentArr.length] = template;
                },
                apply: function(componentName){
                    for(var i=0;i<componentArr.length;i++){
                        if(componentArr[i] == componentName && i%2 == 0){
                            var componentElementList = document.querySelectorAll(componentName);
                            for(var a=0;a<componentElementList.length;a++){
                                componentElementList[a].innerHTML = componentArr[i+1];
                            }
                            break;
                        }
                    }
                },
                //componentArr结构如下：componentName template
                getAllComponent: function(){
                    return componentArr;
                }
            }
        }
    });
    Per().joinModule("Per.ajax",allVersion, function(){
        window.Per.prototype.ajax = function(type,url,msg,async,callback){
            //callback参数必须有一个值来接受信息
            //当async为true时，并且请求出现异常时，系统将会自动向callback方法里传入异常状态码
            if(type.toUpperCase() == "GET"){
                //GET请求msg参数无效，需要提交数据的话请放置在url中
                if(async||async==undefined||async==""||async==null){
                    if(callback!=""&&callback!=null&&callback!=undefined){
                        var xmlhttp;
                        xmlhttp=new XMLHttpRequest();
                        xmlhttp.onreadystatechange=function(){
                            if (xmlhttp.readyState==4 && xmlhttp.status==200){
                                callback(xmlhttp.responseText);
                            }else if(xmlhttp.readyState==4 && xmlhttp.status!=200){
                                callback(xmlhttp.status);
                            }
                        }
                        xmlhttp.open("GET",url,true);
                        xmlhttp.send();
                    }else{
                        console.error("Per.js: unknow function ajax's callback value!");
                    }
                }else if(!async){
                    if(callback!=""&&callback!=null&&callback!=undefined){
                        var xmlhttp;
                        xmlhttp=new XMLHttpRequest();
                        xmlhttp.open("GET",url,false);
                        xmlhttp.send();
                        callback(xmlhttp.responseText);
                    }else{
                        console.error("Per.js: unknow function ajax's callback value!");
                    }
                }else{
                    console.error("Per.js: unknow function ajax's async value!");
                }
            }else if(type.toUpperCase() == "POST"){
                //POST请求msg参数必须填写
                if(async||async==undefined||async==""||async==null){
                    if(callback!=""&&callback!=null&&callback!=undefined){
                        var xmlhttp;
                        xmlhttp=new XMLHttpRequest();
                        xmlhttp.onreadystatechange=function(){
                            if (xmlhttp.readyState==4 && xmlhttp.status==200){
                                callback(xmlhttp.responseText);
                            }else if(xmlhttp.readyState==4 && xmlhttp.status!=200){
                                callback(xmlhttp.status);
                            }
                        }
                        xmlhttp.open("POST",url,true);
                        xmlhttp.setRequestHeader("Content-type","application/x-www-form-urlencoded");
                        xmlhttp.send(msg);
                    }else{
                        console.error("Per.js: unknow function ajax's callback value!");
                    }
                }else if(!async){
                    if(callback!=""&&callback!=null&&callback!=undefined){
                        var xmlhttp;
                        xmlhttp=new XMLHttpRequest();
                        xmlhttp.open("POST",url,false);
                        xmlhttp.setRequestHeader("Content-type","application/x-www-form-urlencoded");
                        xmlhttp.send(msg);
                        callback(xmlhttp.responseText);
                    }else{
                        console.error("Per.js: unknow function ajax's callback value!");
                    }
                }else{
                    console.error("Per.js: unknow function ajax's async value!");
                }
            }else{
                console.error("Per.js: unknow function ajax's type value!")
            }
        }        
    });
    var getObjKeyAndVal = function(obj){
        var arr = new Array(); 
        for(var i in obj){
            arr[arr.length] = i;
            arr[arr.length] = obj[i];
        }
        return arr;
    }
})(window);