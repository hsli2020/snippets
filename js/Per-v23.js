/* Per.js Community */
/* Version: 2.3 */
/* (c) 2018 Skyogo Studio */
/* Released under the Apache License Version 2.0 */

(function(window,undefined){
    var allVersion = 2.3;
    console.info("Welcome running Per.js "+allVersion+" Community!\nVisit ours website: http://www.skyogo.com\nto download our other projects, or check the update for Per.js!");
    var pluginArr = new Array();
    var usedPluginArr = ["Per","Per.do","Per.version","Per.joinModule","Per.use","Per.getAllModuleName","Per.getAllModuleVersion","Per.isThisModuleUsed","Per.getAllModuleNameAndVersion"];
    var dataElArr = new Array();
    //初始化方法
    var isThisModuleUsed = function(moduleName){
        if(typeof moduleName == "string"){
            for(var i=0;i<usedPluginArr.length;i++){
                if(usedPluginArr[i] == moduleName){
                    return true;
                }
            }
            return false;
        }
    }
    var getAllModuleNameAndVersion = function(){
        var allModuleArr = new Array();
        for(var i=0;i<pluginArr.length;i++){
            if(i%3 == 0){
                allModuleArr[allModuleArr.length] = pluginArr[i];
                allModuleArr[allModuleArr.length] = pluginArr[i+1];
            }
        }
        return allModuleArr;
    }
    var joinModule = function(moduleName,moduleVersion,moduleFunction){
        //Per.js官方规定，moduleName里面的结构路径分隔符应为"."，例如dom.ajax，则为dom包下面的ajax包。
        //结构路径可为多重包，例如：dom.do.test
        if(typeof moduleName == "string"&&typeof moduleFunction == "function"&&typeof moduleVersion == "number"){
            var allModule = getAllModuleNameAndVersion();
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
    var use = function(moduleName){
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
        }else if(typeof moduleName == "string"){
            if(moduleName == "all"){
                moduleName = getAllModuleNameAndVersion();
                for(var a=0;a<moduleName.length;a++){
                    if(a%2 == 0){
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
            }else{
                if(!isThisModuleUsed(moduleName)){
                    for(var i=0;i<pluginArr.length;i++){
                        if(i%3 == 0&&pluginArr[i] == moduleName){
                            pluginArr[i+2]();
                            usedPluginArr[usedPluginArr.length] = moduleName;
                        }
                    }
                }else{
                    console.error("Per.js please do not reuse the module! at module "+moduleName);
                }
            }
        }
    }
    window.Per = function(el){
        this.el = el;
        this.version = allVersion;
        this.do = function(Obj,isReactive){
            if(isReactive == null||isReactive == undefined||isReactive == ""){
                isReactive = false;
            }
            if(typeof Obj == "object"&&typeof isReactive == "boolean"){
                if(this.el!=null&&this.el!=undefined&&this.el!=""){
                    Obj.el = this.el;
                }
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
                                    for(var i=0,len = Element.length;i<len;i++){
                                        Element[i].innerHTML = Obj.html;
                                    }
                                },
                                get: function(){
                                    return Obj.html;
                                }
                            })
                        }
                        for(var i=0,len = Element.length;i<len;i++){
                            Element[i].innerHTML = Obj.html;
                        }
                    }
                    if(typeof Obj.text=="string"){
                        if(isReactive){
                            Object.defineProperty(this.do,"text",{
                                set: function(newVal){
                                    Obj.text = newVal;
                                    for(var i=0,len = Element.length;i<len;i++){
                                        Element[i].innerText = Obj.text;
                                    }
                                },
                                get: function(){
                                    return Obj.text;
                                }
                            })
                        }
                        for(var i=0,len = Element.length;i<len;i++){
                            Element[i].innerText = Obj.text;
                        }
                    }
                    if(typeof Obj.val=="string"){
                        if(isReactive){
                            Object.defineProperty(this.do,"val",{
                                set: function(newVal){
                                    Obj.val = newVal;
                                    for(var i=0,len = Element.length;i<len;i++){
                                        Element[i].value = Obj.val;
                                    }
                                },
                                get: function(){
                                    return Obj.val;
                                }
                            })
                        }
                        for(var i=0,len = Element.length;i<len;i++){
                            Element[i].value = Obj.val;
                        }
                    }
                    if(typeof Obj.css=="object"){
                        if(isReactive){
                            Object.defineProperty(this.do,"css",{
                                set: function(newVal){
                                    Obj.css = newVal;
                                    var cssObjArr = getObjKeyAndVal(Obj.css);
                                    for(var a=0,len = cssObjArr.length;a<len;a++){
                                        if(a%2 == 0){
                                            for(var i=0,len2 = Element.length;i<len2;i++){
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
                        for(var a=0,len = cssObjArr.length;a<len;a++){
                            if(a%2 == 0){
                                for(var i=0,len2 = Element.length;i<len2;i++){
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
                                    for(var a=0,len = attrObjArr.length;a<len;a++){
                                        if(a%2 == 0){
                                            for(var i=0,len2 = Element.length;i<len2;i++){
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
                        for(var a=0,len = attrObjArr.length;a<len;a++){
                            if(a%2 == 0){
                                for(var i=0,len2 = Element.length;i<len2;i++){
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
                                    for(var i=0,len = Element.length;i<len;i++){
                                        if(Element[i].getAttribute("p-for-in") != null){
                                            var singleLi = Element[i].innerHTML;
                                            var obj = new Object();
                                            obj[Element[i].getAttribute("p-for-in")] = Obj.for[0];
                                            var nowElInner = dataReplace(Obj.el,singleLi,obj);
                                            for(var a=1;a<Obj.for.length;a++){
                                                nowElInner += singleLi;
                                                obj[Element[i].getAttribute("p-for-in")] = Obj.for[a];
                                                var nowElInner = dataReplace(Obj.el,nowElInner,obj);
                                            }
                                            Element[i].innerHTML = nowElInner;
                                        }else{
                                            var forInnerHTML = "";
                                            for(var a=0,len2 = Obj.for.length;a<len2-1;a++){
                                                if(a==len2-2){
                                                    forInnerHTML += "<li>"+Obj.for[a]+"</li>";
                                                    break;
                                                }
                                                forInnerHTML += "<li>"+Obj.for[a]+"</li>"+Obj.for[len2-1];
                                            }
                                            Element[i].innerHTML = forInnerHTML;
                                        }
                                    }
                                },
                                get: function(){
                                    return Obj.for;
                                }
                            })
                        }
                        //数组最后一位为遍历中间插入的内容，如果不需要可以留空
                        //如果有p-for-in参数则执行in var
                        for(var i=0,len = Element.length;i<len;i++){
                            if(Element[i].getAttribute("p-for-in") != null){
                                var singleLi = Element[i].innerHTML;
                                var obj = new Object();
                                obj[Element[i].getAttribute("p-for-in")] = Obj.for[0];
                                var nowElInner = dataReplace(Obj.el,singleLi,obj);
                                for(var a=1;a<Obj.for.length;a++){
                                    nowElInner += singleLi;
                                    obj[Element[i].getAttribute("p-for-in")] = Obj.for[a];
                                    var nowElInner = dataReplace(Obj.el,nowElInner,obj);
                                }
                                Element[i].innerHTML = nowElInner;
                            }else{
                                var forInnerHTML = "";
                                for(var a=0,len2 = Obj.for.length;a<len2-1;a++){
                                    if(a==len2-2){
                                        forInnerHTML += "<li>"+Obj.for[a]+"</li>";
                                        break;
                                    }
                                    forInnerHTML += "<li>"+Obj.for[a]+"</li>"+Obj.for[len2-1];
                                }
                                Element[i].innerHTML = forInnerHTML;
                            }
                        }
                    }
                    if(typeof Obj.import=="string"){
                        if(isReactive){
                            Object.defineProperty(this.do,"importCallback",{
                                set: function(newVal){
                                    Obj.importCallback = newVal;
                                },
                                get: function(){
                                    return Obj.importCallback;
                                }
                            });
                            Object.defineProperty(this.do,"import",{
                                set: function(newVal){
                                    Obj.import = newVal;
                                    if(!isThisModuleUsed("Per.ajax")){
                                        use(["Per.ajax"]);
                                    }
                                    Per().ajax("GET",Obj.import,"",true,function(html){
                                        for(var i=0;i<Element.length;i++){
                                            Element[i].innerHTML = html;
                                        }
                                        if(typeof Obj.importCallback == "string"){
                                            var MethodArr = getObjKeyAndVal(Obj.method);
                                            for(var a=0;a<MethodArr.length;a++){
                                                if(MethodArr[a] == Obj.importCallback){
                                                    for(var i=0;i<Element.length;i++){
                                                        MethodArr[a+1]();
                                                    }
                                                    break;
                                                }
                                            }
                                        }else if(typeof Obj.importCallback == "function"){
                                            Obj.importCallback();
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
                        Per().ajax("GET",Obj.import,"",true,function(html){
                            for(var i=0;i<Element.length;i++){
                                Element[i].innerHTML = html;
                            }
                            if(typeof Obj.importCallback == "string"){
                                var MethodArr = getObjKeyAndVal(Obj.method);
                                for(var a=0;a<MethodArr.length;a++){
                                    if(MethodArr[a] == Obj.importCallback){
                                        for(var i=0;i<Element.length;i++){
                                            MethodArr[a+1]();
                                        }
                                        break;
                                    }
                                }
                            }else if(typeof Obj.importCallback == "function"){
                                Obj.importCallback();
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
                                    if(Obj.bindType.substr(0,3) == "in "){
                                        var bindElement = document.querySelector(Obj.el);
                                        var dataObj = new Object();
                                        dataObj[Obj.bindType.substr(3,Obj.bindType.length-3)] = bindElement.value;
                                        bindElement.el = Obj.el;
                                        bindElement.dataName = Obj.bindType.substr(3,Obj.bindType.length-3);
                                        bindElement.per = Per(Obj.bind);
                                        bindElement.per.do({
                                            data: dataObj
                                        },true);
                                        bindElement.oninput = function(){
                                            var dataObj = new Object();
                                            dataObj[this.dataName] = document.querySelector(this.el).value;
                                            this.per.do.data = dataObj;
                                        }
                                    }else if(Obj.bindType == "html"){
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
                        if(Obj.bindType.substr(0,3) == "in "){
                            var bindElement = document.querySelector(Obj.el);
                            var dataObj = new Object();
                            dataObj[Obj.bindType.substr(3,Obj.bindType.length-3)] = bindElement.value;
                            bindElement.el = Obj.el;
                            bindElement.dataName = Obj.bindType.substr(3,Obj.bindType.length-3);
                            bindElement.per = Per(Obj.bind);
                            bindElement.per.do({
                                data: dataObj
                            },true);
                            bindElement.oninput = function(){
                                var dataObj = new Object();
                                dataObj[this.dataName] = document.querySelector(this.el).value;
                                this.per.do.data = dataObj;
                            }
                        }else if(Obj.bindType == "html"){
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
                    if(typeof Obj.if == "object"){
                        if(isReactive){
                            Object.defineProperty(this.do,"if",{
                                set: function(newVal){
                                    Obj.if = newVal;
                                    if(Obj.if.if.con){
                                        for(var i=0;i<Element.length;i++){
                                            Element[i].innerHTML = Obj.if.if.template;
                                        }
                                    }else if(typeof Obj.if.else == "object"){
                                        var arr = getObjKeyAndVal(Obj.if.else);
                                        for(var a=0;a<arr.length;a++){
                                            if(a%2 == 0&&arr[a] == "con"){
                                                if(arr[a+1]&&arr[a+2] == "template"){
                                                    for(var i=0;i<Element.length;i++){
                                                        Element[i].innerHTML = arr[a+3];
                                                    }
                                                    break;
                                                }
                                            }else if(arr[a] == "else"){
                                                for(var i=0;i<Element.length;i++){
                                                    Element[i].innerHTML = arr[a+1];
                                                }
                                            }
                                        }
                                    }
                                },
                                get: function(){
                                    return Obj.if;
                                }
                            })
                        }
                        if(Obj.if.if.con){
                            for(var i=0;i<Element.length;i++){
                                Element[i].innerHTML = Obj.if.if.template;
                            }
                        }else if(typeof Obj.if.else == "object"){
                            var arr = getObjKeyAndVal(Obj.if.else);
                            for(var a=0;a<arr.length;a++){
                                if(a%2 == 0&&arr[a] == "con"){
                                    if(arr[a+1]&&arr[a+2] == "template"){
                                        for(var i=0;i<Element.length;i++){
                                            Element[i].innerHTML = arr[a+3];
                                        }
                                        break;
                                    }
                                }else if(arr[a] == "else"){
                                    for(var i=0;i<Element.length;i++){
                                        Element[i].innerHTML = arr[a+1];
                                    }
                                }
                            }
                        }
                    }
                    if(typeof Obj.data == "object"){//只支持单元素选择
                        if(isReactive){
                            Object.defineProperty(this.do,"data",{
                                set: function(newVal){
                                    for(var i in newVal){
                                        Obj.data[i] = newVal[i];
                                    }
                                    for(var i = 0,len = dataElArr.length;i<len;i++){
                                        if(i%2 == 0 && dataElArr[i] == Obj.el){
                                            var dataArr = getObjKeyAndVal(Obj.data);
                                            var html = dataElArr[i+1];
                                            for(var a=0;a<dataArr.length;a++){
                                                if(a%2 == 0){
                                                    var o = 0;
                                                    while(html.indexOf("\{\{"+dataArr[a]+".",o)!=-1||html.indexOf("\{\{"+dataArr[a]+"\}\}",o)!=-1){
                                                        var reg = new RegExp(dataArr[a]);
                                                        var splitOr = html.substr(html.indexOf("\{\{",o)+2,html.indexOf("\}\}",o)-html.indexOf("\{\{",o)-2);
                                                        if(typeof dataArr[a+1] == "string"){
                                                             if(dataArr[a] == splitOr.substr(0,dataArr[a].length)){
                                                                var val = splitOr.replace(reg,"\""+dataArr[a+1]+"\"");
                                                                var reg = new RegExp("``"+dataArr[a]+"``","g");
                                                                val = val.replace(reg,"\""+dataArr[a+1]+"\"");
                                                            }else{
                                                                var val = "";
                                                            }
                                                        }else if(typeof dataArr[a+1] == "object"&&Array.isArray(dataArr[a+1]) == false){
                                                            if(dataArr[a] == splitOr.substr(0,dataArr[a].length)){
                                                                var val = splitOr.replace(reg,JSON.stringify(dataArr[a+1]));
                                                            }else{
                                                                var val = "";
                                                            }
                                                        }else{
                                                            if(dataArr[a] == splitOr.substr(0,dataArr[a].length)){
                                                                var val = splitOr.replace(reg,dataArr[a+1]);
                                                            }else{
                                                                var val = "";
                                                            }
                                                        }
                                                        if(val != ""){
                                                            var returnVal = new Function("return "+val)();
                                                            if(typeof returnVal == "object"&&Array.isArray(returnVal) == false){
                                                                returnVal = JSON.stringify(returnVal);
                                                            }
                                                            var splitOr2 = html.substr(html.indexOf("\{\{",o)+2,html.indexOf("\}\}",o)-html.indexOf("\{\{",o)-2);
                                                            splitOr2 = splitOr2.replace(/\(/g,"\\(").replace(/\)/g,"\\)").replace(/\./g,"\\.").replace(/\,/g,"\\,");
                                                            var reg = new RegExp("\{\{"+splitOr2+"\}\}","g");
                                                            html = html.replace(reg,returnVal);
                                                            o += returnVal.length;
                                                        }else{
                                                            o = html.indexOf("\}\}",o)+2;
                                                        }
                                                    }
                                                   if(document.querySelector(Obj.el).getAttribute("p-html") == null){
                                                        html = html.replace(/</g,"&lt;").replace(/>/g,"&gt;");
                                                    }
                                                    document.querySelector(Obj.el).innerHTML = html.toString();
                                                }
                                            }
                                        }
                                    }
                                },
                                get: function(){
                                    return Obj.data;
                                }
                            })
                        }
                        var dataElement = document.querySelector(Obj.el);
                        var dataArr = getObjKeyAndVal(Obj.data);
                        var html = dataElement.innerHTML;
                        //html标签解析替换
                        if(dataElement.getAttribute("p-html") == null){
                            html = html.replace(/</g,"&lt;").replace(/>/g,"&gt;");
                        }
                        dataElArr[dataElArr.length] = Obj.el;
                        dataElArr[dataElArr.length] = html;
                        for(var i = 0,len = dataElArr.length;i<len;i++){
                            if(i%2 == 0 && dataElArr[i] == Obj.el){
                                var dataArr = getObjKeyAndVal(Obj.data);
                                var html = dataElArr[i+1];
                                for(var a=0;a<dataArr.length;a++){
                                    if(a%2 == 0){
                                        var o = 0;
                                        while(html.indexOf("\{\{"+dataArr[a]+".",o)!=-1||html.indexOf("\{\{"+dataArr[a]+"\}\}",o)!=-1){
                                            var reg = new RegExp(dataArr[a]);
                                            var splitOr = html.substr(html.indexOf("\{\{",o)+2,html.indexOf("\}\}",o)-html.indexOf("\{\{",o)-2);
                                            if(typeof dataArr[a+1] == "string"){
                                                 if(dataArr[a] == splitOr.substr(0,dataArr[a].length)){
                                                    var val = splitOr.replace(reg,"\""+dataArr[a+1]+"\"");
                                                    var reg = new RegExp("``"+dataArr[a]+"``","g");
                                                    val = val.replace(reg,"\""+dataArr[a+1]+"\"");
                                                }else{
                                                    var val = "";
                                                }
                                            }else if(typeof dataArr[a+1] == "object"&&Array.isArray(dataArr[a+1]) == false){
                                                if(dataArr[a] == splitOr.substr(0,dataArr[a].length)){
                                                    var val = splitOr.replace(reg,JSON.stringify(dataArr[a+1]));
                                                }else{
                                                    var val = "";
                                                }
                                            }else{
                                                if(dataArr[a] == splitOr.substr(0,dataArr[a].length)){
                                                    var val = splitOr.replace(reg,dataArr[a+1]);
                                                }else{
                                                    var val = "";
                                                }
                                            }
                                            if(val != ""){
                                                var returnVal = new Function("return "+val)();
                                                if(typeof returnVal == "object"&&Array.isArray(returnVal) == false){
                                                    returnVal = JSON.stringify(returnVal);
                                                }
                                                var splitOr2 = html.substr(html.indexOf("\{\{",o)+2,html.indexOf("\}\}",o)-html.indexOf("\{\{",o)-2);
                                                splitOr2 = splitOr2.replace(/\(/g,"\\(").replace(/\)/g,"\\)").replace(/\./g,"\\.").replace(/\,/g,"\\,");
                                                var reg = new RegExp("\{\{"+splitOr2+"\}\}","g");
                                                html = html.replace(reg,returnVal);
                                                o += returnVal.length;
                                            }else{
                                                o = html.indexOf("\}\}",o)+2;
                                            }
                                        }
                                       if(document.querySelector(Obj.el).getAttribute("p-html") == null){
                                            html = html.replace(/</g,"&lt;").replace(/>/g,"&gt;");
                                        }
                                        document.querySelector(Obj.el).innerHTML = html.toString();
                                    }
                                }
                            }
                        }
                        dataElement.innerHTML = html;
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
                    console.error("Per.js: para el cannot be null!");
                }
            }
        }
        this.joinModule = joinModule;
        this.use = use;
        this.getAllModuleNameAndVersion = getAllModuleNameAndVersion;
        this.isThisModuleUsed = isThisModuleUsed;
        if(this == window){
            return new Per(el);
        }
    }
    
    /* 以下代码为了无括号构造 */
    Per.isThisModuleUsed = isThisModuleUsed;
    Per.getAllModuleNameAndVersion = getAllModuleNameAndVersion;
    Per.joinModule = joinModule;
    Per.use = use;
    
    /* 以下是内置模块 */
    var per = Per();
    per.joinModule("Per.component",allVersion, function(){
        var componentArr = new Array();
        window.Per.prototype.component = {
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
        Per.component = window.Per.prototype.component;
    });
    per.joinModule("Per.ajax",allVersion, function(){
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
        Per.ajax = window.Per.prototype.ajax;
    });
    per.joinModule("Per.page",allVersion,function(){
        var perPageModulePageArr = new Array();
        window.Per.prototype.page = {
            create: {
                page: function(pageGroupName, pageEl){
                    if(typeof pageEl == "string"){
                        if(perPageModulePageArr[pageGroupName] == undefined||perPageModulePageArr[pageGroupName] == null){
                            console.error("Per.js: pageGroup need create frist! at pageGroup: "+pageGroupName);
                        }else{
                            perPageModulePageArr[pageGroupName][perPageModulePageArr[pageGroupName].length] = pageEl;
                        }
                    }else if(Array.isArray(pageEl)){
                        for(var i=0,len = pageEl.length;i<len;i++){
                            if(perPageModulePageArr[pageGroupName] == undefined||perPageModulePageArr[pageGroupName] == null){
                                console.error("Per.js: pageGroup need create frist! at pageGroup: "+pageGroupName);
                            }else{
                                perPageModulePageArr[pageGroupName][perPageModulePageArr[pageGroupName].length] = pageEl[i];
                            }
                        }
                    }else{
                        console.error("Per.js: unknow pageEl's type.");
                    }
                },
                pageGroup: function(pageGroupName){
                    if(typeof pageGroupName == "string"){
                        if(perPageModulePageArr[pageGroupName] == undefined||perPageModulePageArr[pageGroupName] == null){
                            perPageModulePageArr[pageGroupName] = new Array();
                        }else{
                            console.error("Per.js: this pageGroup has already been create! at pageGroup: "+pageGroupName);
                        }
                    }else if(Array.isArray(pageGroupName)){
                        for(var i=0,len = pageGroupName.length;i<len;i++){
                            if(perPageModulePageArr[pageGroupName[i]] == undefined||perPageModulePageArr[pageGroupName[i]] == null){
                                perPageModulePageArr[pageGroupName[i]] = new Array();
                            }else{
                                console.error("Per.js: this pageGroup has already been create! at pageGroup: "+pageGroupName[i]);
                            }
                        }
                    }else{
                        console.error("Per.js: unknow pageGroupName's type.");
                    }
                }
            },
            to: function(pageGroupName, pageNumber){
                if(perPageModulePageArr[pageGroupName] == undefined||perPageModulePageArr[pageGroupName] == null){
                    console.error("Per.js: pageGroup need create frist! at pageGroup: "+pageGroupName);
                }else{
                    var arr = perPageModulePageArr[pageGroupName];
                    for(var i=0;i<arr.length;i++){
                        var elArr = document.querySelectorAll(arr[i]);
                        for(var a=0;a<elArr.length;a++){
                            elArr[a].style.display = "none";
                        }
                    }
                    var elArr = document.querySelectorAll(arr[pageNumber-1]);
                    for(var a=0;a<elArr.length;a++){
                        elArr[a].style.display = "";
                    }
                }
            },
            remove: {
                page: function(pageGroupName, pageEl){
                    if(typeof pageEl == "string"){
                        if(perPageModulePageArr[pageGroupName] == undefined||perPageModulePageArr[pageGroupName] == null){
                            console.error("Per.js: you need create this pageGroup before remove it! at pageGroup: "+pageGroupName);
                        }else{
                            var arr = perPageModulePageArr[pageGroupName];
                            for(var i=0;i<arr.length;i++){
                                if(arr[i] == pageEl){
                                    arr.splice(i,1);
                                }
                            }
                            perPageModulePageArr[pageGroupName] = arr;
                        }
                    }else if(Array.isArray(pageEl)){
                        for(var i=0,len = pageEl.length;i<len;i++){
                            if(perPageModulePageArr[pageGroupName] == undefined||perPageModulePageArr[pageGroupName] == null){
                                console.error("Per.js: you need create this pageGroup before remove it! at pageGroup: "+pageGroupName);
                            }else{
                                var arr = perPageModulePageArr[pageGroupName];
                                for(var i=0;i<arr.length;i++){
                                    if(arr[i] == pageEl[i]){
                                        arr.splice(i,1);
                                    }
                                }
                                perPageModulePageArr[pageGroupName] = arr;
                            }
                        }
                    }else{
                        console.error("Per.js: unknow pageEl's type.");
                    }
                },
                pageGroup: function(pageGroupName){
                    if(typeof pageGroupName == "string"){
                        if(perPageModulePageArr[pageGroupName] == undefined||perPageModulePageArr[pageGroupName] == null){
                            console.error("Per.js: you need create this pageGroup before remove it! at pageGroup: "+pageGroupName);
                        }else{
                            perPageModulePageArr[pageGroupName] = undefined;
                        }
                    }else if(Array.isArray(pageGroupName)){
                        for(var i=0,len = pageGroupName.length;i<len;i++){
                            if(perPageModulePageArr[pageGroupName[i]] == undefined||perPageModulePageArr[pageGroupName[i]] == null){
                                console.error("Per.js: you need create this pageGroup before remove it! at pageGroup: "+pageGroupName[i]);
                            }else{
                                perPageModulePageArr[pageGroupName[i]] = undefined;
                            }
                        }
                    }else{
                        console.error("Per.js: unknow pageGroupName's type.");
                    }
                }
            },
            get: {
                pageGroup: function(pageGroupName){
                    if(perPageModulePageArr[pageGroupName] == undefined||perPageModulePageArr[pageGroupName] == null){
                        console.error("Per.js: you need create this pageGroup before get it! at pageGroup: "+pageGroupName);
                    }else{
                        return perPageModulePageArr[pageGroupName];
                    }
                }
            }
        }
        Per.page = window.Per.prototype.page;
    });
    per.joinModule("Per.check",allVersion,function(){
        window.Per.prototype.check = {
            mail: function(mailText){
                var reg = /^([0-9A-Za-z\-_\.]+)@([0-9a-z]+\.[a-z]{2,3}(\.[a-z]{2})?)$/g;
                return reg.test(mailText);
            },
            html: function(text){
                var reg = new RegExp(/\<|\>|\\/g);
                return reg.test(text);
            },
            URL: function(text){
                var RegUrl = new RegExp();
                RegUrl.compile("^[A-Za-z]+://[A-Za-z0-9-_]+\\.[A-Za-z0-9-_%&\?\/.=]+$");
                return RegUrl.test(text);
            }
        }
        Per.check = window.Per.prototype.check;
    });
    per.joinModule("Per.get",allVersion,function(){
        var perGetFun = function(el){
            if(el == undefined||el == null||el == ""){
                console.error("Per.js: you need set el attr, before use Per.get module!");
            }else{
                return {
                    css: function(cssName){
                        var elArr = document.querySelectorAll(el);
                        var returnText = "";
                        for(var i=0;i<elArr.length;i++){
                            returnText += elArr[i].style[cssName];
                        }
                        return returnText;
                    },
                    attr: function(attrName){
                        var elArr = document.querySelectorAll(el);
                        var returnText = "";
                        for(var i=0;i<elArr.length;i++){
                            returnText += elArr[i].getAttribute(attrName);
                        }
                        return returnText;
                    },
                    height: function(){
                        var elArr = document.querySelectorAll(el);
                        var returnText = "";
                        for(var i=0;i<elArr.length;i++){
                            returnText += elArr[i].getAttribute("height");
                        }
                        return returnText;
                    },
                    width: function(){
                        var elArr = document.querySelectorAll(el);
                        var returnText = "";
                        for(var i=0;i<elArr.length;i++){
                            returnText += elArr[i].getAttribute("width");
                        }
                        return returnText;
                    },
                    html: function(){
                        var elArr = document.querySelectorAll(el);
                        var returnText = "";
                        for(var i=0;i<elArr.length;i++){
                            returnText += elArr[i].innerHTML;
                        }
                        return returnText;
                    },
                    text: function(){
                        var elArr = document.querySelectorAll(el);
                        var returnText = "";
                        for(var i=0;i<elArr.length;i++){
                            returnText += elArr[i].innerText;
                        }
                        return returnText;
                    },
                    val: function(){
                        var elArr = document.querySelectorAll(el);
                        var returnText = "";
                        for(var i=0;i<elArr.length;i++){
                            returnText += elArr[i].value;
                        }
                        return returnText;
                    },
                    class: function(){
                        var elArr = document.querySelectorAll(el);
                        var returnText = "";
                        for(var i=0;i<elArr.length;i++){
                            returnText += elArr[i].getAttribute("class");
                        }
                        return returnText;
                    },
                    parent: function(){
                        //返回数组或字符串形式
                        var elArr = document.querySelectorAll(el);
                        if(elArr.length == 1){
                            var returnText = "";
                            for(var i=0;i<elArr.length;i++){
                                returnText += elArr[i].parentNode;
                            }
                            return returnText;
                        }else{
                            var returnArr = new Array();
                            for(var i=0;i<elArr.length;i++){
                                returnArr[returnArr.length] = elArr[i].parentNode;
                            }
                            return returnArr;
                        }
                    },
                    children: function(){
                        //返回2维数组形式，第一维为父元素，第二维为子节点
                        var elArr = document.querySelectorAll(el);
                        var returnArr = new Array();
                        for(var i=0;i<elArr.length;i++){
                            returnArr[returnArr.length] = elArr[i].childNodes;
                        }
                        return returnArr;
                    }
                }
            }
        }
        window.Per.prototype.get = perGetFun;
        window.Per.prototype.$ = perGetFun;
    });
    per.joinModule("Per.lazyLoad",allVersion,function(){
        var lazyLoadListenerRepeatTime = 25;//ms
        var lazyLoadRange = 100;//px
        var lazyLoadList = new Array();
        window.Per.prototype.lazyLoad = {
            setLazyLoadListenerRepeatTime: function(num){
                if(typeof num == "number"){
                    lazyLoadListenerRepeatTime = num;
                }else{
                    console.error("Per.js: function setLazyLoadListenerRepeatTime's para num's type should be number!");
                }
            },
            setLazyLoadRange: function(num){
                if(typeof num == "number"){
                    lazyLoadRange = num;
                }else{
                    console.error("Per.js: function setLazyLoadRange's para num's type should be number!");
                }
            },
            setLazyLoad: function(el, url){
                lazyLoadList[lazyLoadList.length] = el;
                lazyLoadList[lazyLoadList.length] = url;
            },
            clearLazyLoadTimer: function(){
                clearInterval(lazyLoadTimer);
            }
        }
        var lazyLoadTimer = setInterval(function(){//懒加载监听器
            if(lazyLoadList.length != 0){
                for(var i=0;i<lazyLoadList.length;i++){
                    if(i%2 == 0){
                        var Element = document.querySelector(lazyLoadList[i]);
                        var h = window.screen.availHeight;
                        if(Element.getBoundingClientRect().top-(h+lazyLoadRange) <= 0){
                            Element.setAttribute("src",lazyLoadList[i+1]);
                            lazyLoadList.splice(0,2);
                            i-=2;
                        }
                    }
                }
            }
        },lazyLoadListenerRepeatTime);
        Per.lazyLoad = window.Per.prototype.lazyLoad;
    });
    per.joinModule("Per.each",allVersion, function(){
        //fun参数类型需为function，且必须要有2个参数，第一个用来获取i的值，用来获取当前arr的值
        window.Per.prototype.each = function(arr,fun){
            if(typeof fun == "function"&&Array.isArray(arr)){
                for(var i=0,len = arr.length;i<len;i++){
                    fun(i,arr[i]);
                }
            }
        }
        Per.each = window.Per.prototype.each;
    });
    per.joinModule("Per.browser",allVersion, function(){
        window.Per.prototype.browser = {
            type: function(){
                var userAgent = navigator.userAgent;
                var isOpera = userAgent.indexOf("Opera") > -1;
                if(isOpera) {
                    return "Opera";
                }else if(userAgent.indexOf("Firefox") > -1) {
                    return "Firefox";
                }else if(userAgent.indexOf("Chrome") > -1){
                    return "Chrome";
                }else if(userAgent.indexOf("Safari") > -1) {
                    return "Safari";
                }else if(userAgent.indexOf("compatible") > -1 && userAgent.indexOf("MSIE") > -1 && !isOpera) {
                    return "IE";
                }else if (userAgent.indexOf("Trident") > -1) {
                    return "Edge";
                }
            },
            isPC: function(){
                if(/Android|webOS|iPhone|iPod|BlackBerry/i.test(navigator.userAgent)) {
                    return false;
                } else {
                    return true;
                }
            },
            OSType: function(){
                var sUserAgent = navigator.userAgent;
                var isWin = (navigator.platform == "Win32") || (navigator.platform == "Windows");
                var isMac = (navigator.platform == "Mac68K") || (navigator.platform == "MacPPC") || (navigator.platform == "Macintosh") || (navigator.platform == "MacIntel");
                if (isMac) return "MacOS";
                var isUnix = (navigator.platform == "X11") && !isWin && !isMac;
                if (isUnix) return "Unix";
                var isLinux = (String(navigator.platform).indexOf("Linux") > -1);
                if (isLinux) return "Linux";
                if (isWin) {
                    var isWin2K = sUserAgent.indexOf("Windows NT 5.0") > -1 || sUserAgent.indexOf("Windows 2000") > -1;
                    if (isWin2K) return "Windows2000";
                    var isWinXP = sUserAgent.indexOf("Windows NT 5.1") > -1 || sUserAgent.indexOf("Windows XP") > -1;
                    if (isWinXP) return "WindowsXP";
                    var isWin2003 = sUserAgent.indexOf("Windows NT 5.2") > -1 || sUserAgent.indexOf("Windows 2003") > -1;
                    if (isWin2003) return "Windows2003";
                    var isWinVista= sUserAgent.indexOf("Windows NT 6.0") > -1 || sUserAgent.indexOf("Windows Vista") > -1;
                    if (isWinVista) return "Windows Vista";
                    var isWin7 = sUserAgent.indexOf("Windows NT 6.1") > -1 || sUserAgent.indexOf("Windows 7") > -1;
                    if (isWin7) return "Windows7";
                }
                return undefined;
            }
        }
        Per.browser = window.Per.prototype.browser;
    });
    var getObjKeyAndVal = function(obj){
        var arr = new Array(); 
        for(var i in obj){
            arr[arr.length] = i;
            arr[arr.length] = obj[i];
        }
        return arr;
    }
    var dataReplace = function(el,html,data){
        var dataElement = document.querySelector(el);
        var dataArr = getObjKeyAndVal(data);
        var html = html;
        //html标签解析替换
        if(dataElement.getAttribute("p-html") == null){
            html = html.replace(/</g,"&lt;").replace(/>/g,"&gt;");
        }
        dataElArr[dataElArr.length] = el;
        dataElArr[dataElArr.length] = html;
        for(var i = 0,len = dataElArr.length;i<len;i++){
            if(i%2 == 0 && dataElArr[i] == el){
                var dataArr = getObjKeyAndVal(data);
                for(var a=0;a<dataArr.length;a++){
                    if(a%2 == 0){
                        var o = 0;
                        while(html.indexOf("\{\{"+dataArr[a]+".",o)!=-1||html.indexOf("\{\{"+dataArr[a]+"\}\}",o)!=-1){
                            var reg = new RegExp(dataArr[a]);
                            var splitOr = html.substr(html.indexOf("\{\{",o)+2,html.indexOf("\}\}",o)-html.indexOf("\{\{",o)-2);
                            if(typeof dataArr[a+1] == "string"){
                                 if(dataArr[a] == splitOr.substr(0,dataArr[a].length)){
                                    var val = splitOr.replace(reg,"\""+dataArr[a+1]+"\"");
                                    var reg = new RegExp("``"+dataArr[a]+"``","g");
                                    val = val.replace(reg,"\""+dataArr[a+1]+"\"");
                                }else{
                                    var val = "";
                                }
                            }else if(typeof dataArr[a+1] == "object"&&Array.isArray(dataArr[a+1]) == false){
                                if(dataArr[a] == splitOr.substr(0,dataArr[a].length)){
                                    var val = splitOr.replace(reg,JSON.stringify(dataArr[a+1]));
                                }else{
                                    var val = "";
                                }
                            }else{
                                if(dataArr[a] == splitOr.substr(0,dataArr[a].length)){
                                    var val = splitOr.replace(reg,dataArr[a+1]);
                                }else{
                                    var val = "";
                                }
                            }
                            if(val != ""){
                                var returnVal = new Function("return "+val)();
                                if(typeof returnVal == "object"&&Array.isArray(returnVal) == false){
                                    returnVal = JSON.stringify(returnVal);
                                }
                                var splitOr2 = html.substr(html.indexOf("\{\{",o)+2,html.indexOf("\}\}",o)-html.indexOf("\{\{",o)-2);
                                splitOr2 = splitOr2.replace(/\(/g,"\\(").replace(/\)/g,"\\)").replace(/\./g,"\\.").replace(/\,/g,"\\,");
                                var reg = new RegExp("\{\{"+splitOr2+"\}\}","g");
                                html = html.replace(reg,returnVal);
                                o += returnVal.length;
                            }else{
                                o = html.indexOf("\}\}",o)+2;
                            }
                        }
                       if(document.querySelector(el).getAttribute("p-html") == null){
                            html = html.replace(/</g,"&lt;").replace(/>/g,"&gt;");
                        }
                        document.querySelector(el).innerHTML = html.toString();
                    }
                }
            }
        }
        return html;
    }
})(window);