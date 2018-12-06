/**
 * mini $
 * 
 * @param  {string} selector 选择器
 * @return {HTMLElement} 返回匹配的元素
 */
function $(selector) {
	return document.querySelector(selector);
}

/**
 * 给一个element绑定一个针对event事件的响应，响应函数为listener，兼容IE
 * 
 * @param {HTMLELement} 
 * @param {event}
 * @param {Function}
 */
function addEvent(element, event, listener) {
    if (!!element.addEventListener){
        element.addEventListener(event,listener,!1);
    }else{
        element.attachEvent('on'+event,listener);
    }
}

/**
 * 移除element对象对于event事件发生时执行listener的响应
 * 
 * @param  {HTMLElement}
 * @param  {event}
 * @param  {Function}
 */
function removeEvent(element, event, listener) {
    if (!!element.removeEventListener){
        element.removeEventListener(event,listener,!1);
    }else{
        element.detachEvent('on'+event,listener);
    }
}

/**
 * 实现对click事件的绑定
 * 
 * @param {HTMLElement}
 * @param {Function}
 */
function addClickEvent(element, listener) {
    addEvent(element,"click",listener);
}

/**
 * 实现对于按Enter键时的事件绑定
 * 
 * @param {HTMLElement}
 * @param {Function}
 */
function addEnterEvent(element, listener) {
    addEvent(element,"keyup",function(event){
    	if (event.ctrlKey) {
    		listener();
    	};
    });
}

/**
 * 判断parent是否element的祖先节点
 * 
 * @param  {HTMLElement}
 * @param  {HTMLElement}
 * @return {Boolean}
 */
function isParent(element,parentName,stopElem){
	for (var node = element; node !== stopElem; node = node.parentNode) {
        if (node.nodeName.toLowerCase() === parentName) {
            return node;
        }
    }
    return false;
}

/**
 * 事件代理函数，实现对element里面所有tag的eventName事件进行响应
 * 
 * @param  {HTMLElement} 目标节点的祖先元素
 * @param  {HTMLElement} 目标节点
 * @param  {event} 监听事件
 * @param  {Function} 响应函数
 */
function delegateEvent(element, tag, eventName, listener) {
    addEvent(element, eventName, function(event){
    	var event = event || window.event;
    	var target = event.target || event.srcElement;
        var parent = isParent(target,tag,element);
    	if (!!parent) {
    		listener(parent);
    	};
    })
}

$.on = function(selector, event, listener) {
    addEvent($(selector), event, listener);
}

$.un = function(selector, event, listener) {
    removeEvent($(selector), event, listener);
}

$.click = function(selector, listener) {
    addClickEvent($(selector), listener);
}

$.enter = function(selector, listener) {
	addEnterEvent($(selector),listener);
}

$.delegate = function(selector, tag, event, listener) {
    delegateEvent($(selector),tag,event,listener);
}

/**
 * 以下是class的操作，增删改查
 * 
 * @author 是百度IFE的hushicai
 */

/**
* 判断是否有某个className
* @param {HTMLElement} element 元素
* @param {string} className className
* @return {boolean}
*/
function hasClass(element, className) {
    var classNames = element.className;
    if (!classNames) {
        return false;
    }
    classNames = classNames.split(/\s+/);
    for (var i = 0, len = classNames.length; i < len; i++) {
        if (classNames[i] === className) {
            return true;
        }
    }
    return false;
}

/**
* 添加className
*
* @param {HTMLElement} element 元素
* @param {string} className className
*/
function addClass(element, className) {
    if (!hasClass(element, className)) {
        element.className = element.className ?[element.className, className].join(' ') : className;
    }
}

/**
* 删除元素className
*
* @param {HTMLElement} element 元素
* @param {string} className className
*/
function removeClass(element, className) {
    if (className && hasClass(element, className)) {
        var classNames = element.className.split(/\s+/);
        for (var i = 0, len = classNames.length; i < len; i++) {
            if (classNames[i] === className) {
                classNames.splice(i, 1);
                break;
            }
        }
    element.className = classNames.join(' ');
    }
}