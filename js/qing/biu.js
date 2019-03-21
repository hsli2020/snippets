/**
 * @author      : 马蹄疾
 * @date        : 2017-12-17
 * @version     : v1.0
 * @description : 一个轻巧的mvvm框架BiuJS
 * @repository  : https://github.com/veedrin/biu
 * @license     : MIT
 */

(function(root) {

    // .和[]正则表达式，识别splitChain方法中的多级属性
    let regChain = /[\[\]\.'"]/;

    // 十六进制颜色正则表达式，识别hexToRgb方法中的十六进制颜色
    let regHex = /^#([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$/;

    // 空白符正则表达式，识别节点中的空白节点
    let regBlank = /^\s+$/;

    // 胡子模板正则表达式，识别文本节点的胡子模板
    let regMustache = /\{\{(.*?)\}\}/g;

    // for循环正则表达式，识别$for指令的(item, index)
    let regIterate = /^\(([\w\,\s]+)\)/;

    // 方法调用正则表达式，识别$click中的方法
    let regCall = /\((.*)\)$/;

    // 方法参数正则表达式，识别$click中方法的字符串参数
    let regQuote = /[\'\"](.+)[\'\"]/;

    /////////////// *Utility* ///////////////

    /**
     * 将嵌套属性拆分成数组
     * @param {string} 指令和模板表达式
     */
    function splitChain(exp) {
        let arr = exp.split(regChain);
        if (arr.length === 1) {
            return arr;
        }
        let chain = [];
        for (let i = 0, len = arr.length; i < len; i++) {
            arr[i] && chain.push(arr[i]);
        }
        return chain;
    }

    /**
     * 循环获取嵌套属性的值
     * @param {string} 指令和模板表达式
     * @param {object} Biu的实例
     */
    function execChain(exp, vm) {
        let chain = splitChain(exp);
        let obj = vm.$data;
        for (let i = 0; i < chain.length; i++) {
            if (obj == null) {
                return undefined;
            }
            obj = obj[chain[i]];
        }
        return obj;
    }

    /**
     * 递归获取嵌套属性的值并赋值，在$model中用到
     * @param {string} 指令和模板表达式
     * @param {object} Biu的实例
     * @param {any} 用户输入的新值
     */
    function assignChain(exp, vm, newValue) {
        let chain = splitChain(exp);
        let len = chain.length;
        let obj = vm.$data;
        for (let i = 0; i < len - 1; i++) {
            if (obj == null) {
                return undefined;
            }
            obj = obj[chain[i]];
        }
        obj[chain[len - 1]] = newValue;
    }

    /**
     * hex颜色转成rgb颜色
     * @param {string} hex颜色值
     */
    function hexToRgb(hex) {
        hex = hex.trim();
        if (hex && regHex.test(hex)) {
            if (hex.length === 4) {
                let temp = '#';
                for (let i = 1; i < 4; i += 1) {
                    let num = hex.slice(i, i + 1);
                    temp += num.concat(num);
                }
                hex = temp;
            }
            let rgbArr = [];
            for (let i = 1; i < 7; i += 2) {
                rgbArr.push(parseInt(`0x${hex.slice(i, i + 2)}`));
            }
            return `rgb(${rgbArr.join(', ')})`;
        } else {
            return hex;
        }
    }

    /**
     * 获得交集，在$style中用到
     * @param {array} Biu绑定在DOM中的样式数组
     * @param {array} 用户写入DOM中的样式数组
     */
    function getCoverSet(biuSet, userSet) {
        let coverSet = [];
        let flip = {};
        biuSet.forEach((value) => {
            value = value.trim();
            let prop = value.split(':')[0];
            flip[prop] = true;
        });
        userSet.forEach((value) => {
            let prop = value.split(':')[0];
            if (flip[value]) {
                coverSet.push(value);
            }
        });
        return coverSet;
    }

    /**
     * 获得补集，在$style中用到
     * @param {array} Biu绑定在DOM中的样式数组
     * @param {array} 页面获取的混合样式数组
     */
    function getComplementSet(subSet, completeSet) {
        let complementSet = [];
        let flip = {};
        subSet.forEach((value) => {
            value = value.trim();
            flip[value] = true;
        });
        completeSet.forEach((value) => {
            value = value.trim();
            if (!flip[value]) {
                complementSet.push(value);
            }
        });
        return complementSet;
    }

    /////////////// *Observer* ///////////////

    /**
     * 监听者主体方法
     */
    function Observer(data) {
        this.observe(data);
    }

    /**
     * 劫持数据的getter和setter
     */
    Observer.prototype.observe = function(data) {
        if (!data || typeof data !== 'object') {
            return;
        }
        let self = this;

        Object.keys(data).forEach((key) => {
            let value = data[key];
            let dep = new Dep;

            Object.defineProperty(data, key, {
                enumerable: true,
                configurable: true,
                get: function() {
                    // 触发getter添加订阅
                    Dep.target && dep.addSub(Dep.target);
                    return value;
                },
                set: function(newValue) {
                    if (value === newValue) {
                        return;
                    }
                    // 更新value的值，方便下次比较
                    value = newValue;
                    // 触发setter通知订阅者
                    dep.notify(newValue);
                }
            });

            if (!Array.isArray(value)) {
                self.observe(value);
            } else {
                self.observeArray(value, dep);
            }
        });
    };

    /**
     * 劫持操作数组的方法
     */
    Observer.prototype.observeArray = function(arr, dep) {
        let self = this;
        let arrayProto = Array.prototype;
        // 把数组原型上的方法嫁接到空对象上，使之能被Object.defineProperty接受
        let arrayObject = Object.create(arrayProto);
        let methods = ['push', 'pop', 'unshift', 'shift', 'slice', 'splice', 'concat'];

        methods.forEach((method) => {
            let origin = arrayProto[method];

            // 改造操作数组的方法，添加通知功能
            Object.defineProperty(arrayObject, method, {
                enumerable: true,
                writable: true,
                configurable: true,
                value: function() {
                    let args = Array.from(arguments);
                    let result = origin.apply(this, args);
                    let inserted;
                    switch (method) {
                        case 'push':
                        case 'unshift':
                            inserted = args;
                            break;
                        case 'splice':
                            inserted = args.slice(2);
                            break;
                    }
                    // 增量监听
                    if (inserted && inserted.length) {
                        self.observe(this);
                    }
                    // 通知订阅者
                    dep.notify(Array.from(this));
                    return result;
                }
            });
        });
        // 把改造过的方法挂载到arr的原型上
        Object.setPrototypeOf(arr, arrayObject);
        arr.forEach((value) => {
            this.observe(value);
        });
    };

    /////////////// *Dependency* ///////////////

    /**
     * 订阅者主体方法
     */
    function Dep() {
        this.subs = [];
    }

    // 缓存订阅项
    Dep.target = null;

    /**
     * 添加订阅项
     */
    Dep.prototype.addSub = function(sub) {
        this.subs.push(sub);
    };

    /**
     * 通知订阅项触发更新
     */
    Dep.prototype.notify = function(newValue) {
        for (let i = 0; i < this.subs.length; i++) {
            this.subs[i].update(newValue);
        }
    };

    /////////////// *Compiler* ///////////////

    /**
     * 编译者主体方法
     */
    function Compiler(mount, vm) {
        this.vm = vm;
        if (mount) {
            let fragment = this.eleToFragment(mount);
            this.compile(fragment);
            mount.appendChild(fragment);
        }
    }

    /**
     * mount子元素用文档片段包裹
     */
    Compiler.prototype.eleToFragment = function(ele) {
        let fragment = document.createDocumentFragment();
        let child;
        while(child = ele.firstChild) {
            fragment.appendChild(child);
        }
        return fragment;
    };

    /**
     * 元素节点和文本节点分开编译
     * 目前指令和模板表达式均不支持视图内计算
     * 已经实现：<div>{{a.b.c}}</div>或者<div>{{a['b']['c']}}</div>
     * 尚未实现：<div>{{a + b + c}}</div>
     */
    Compiler.prototype.compile = function(ele) {
        let self = this;
        if (ele.childNodes && ele.childNodes.length) {
            Array.from(ele.childNodes).forEach((child) => {
                // 忽略空白文本节点
                if (child.nodeType === 3 && !regBlank.test(child.textContent)) {
                    self.compileText(child);
                } else if (child.nodeType === 1) {
                    self.compileElement(child);
                }
                // 防止编译已经被移除的$for元素，否则找不到(item, index)表达式实际的值
                if (child.parentNode) {
                    self.compile(child);
                }
            });
        }
    };

    /**
     * 编译文本节点
     */
    Compiler.prototype.compileText = function(ele) {
        let content = ele.textContent.trim();
        let fragment = document.createDocumentFragment();
        let i = 0;
        let match;
        let text;

        while (match = regMustache.exec(content)) {
            // 文本
            if (i < match.index) {
                text = content.slice(i, match.index);
                let element = document.createTextNode(text);
                fragment.appendChild(element);
            }

            i = regMustache.lastIndex;
            // 变量
            let exp = match[1];
            let element = document.createTextNode('');
            let result = execChain(exp, this.vm);
            if (result === undefined) {
                console.error(`[Biu warn]无法找到胡子模板中表达式${exp}的值`);
                return;
            }
            element.textContent = result;
            fragment.appendChild(element);

            new Watcher(exp, this.vm, (newValue) => {
                element.textContent = newValue;
            });
        }
        // 文本
        if (i < content.length) {
            text = content.slice(i);
            let element = document.createTextNode(text);
            fragment.appendChild(element);
        }
        ele.parentNode.replaceChild(fragment, ele);
    };

    /**
     * 编译元素节点
     */
    Compiler.prototype.compileElement = function(ele) {
        let self = this;
        Array.from(ele.attributes).forEach((attr) => {
            let name = attr.name;
            let type;
            if (name.indexOf('$') === 0) {
                type = name;
            } else {
                return;
            }
            let exp = attr.value;
            let handler = self[type];
            if (handler) {
                handler.call(self, ele, exp, self.vm);
            }
            ele.removeAttribute(name);
        });
    };

    /**
     * $for指令
     * @用法：<div $for="item in list"></div>
     * @用法：<div $for="(item, index) in list"></div>
     * item和index写法可以自定义
     */
    Compiler.prototype.$for = function(ele, exp, vm) {
        let self = this;
        let split = exp.split('in');
        let expItem = split[0].trim();
        let expList = split[1].trim();
        let expIndex;

        if (regIterate.test(expItem)) {
            let match = regIterate.exec(expItem)[1];
            split = match.split(',');
            expItem = split[0].trim();
            expIndex = split[1].trim();
        }

        let result = execChain(expList, vm);
        if (result === undefined) {
            console.error(`[Biu warn]无法找到$for中表达式${expList}的值`);
            return;
        }

        let list = result;
        let divWrap = document.createElement('div');
        let parentNode = ele.parentNode;
        parentNode.insertBefore(divWrap, ele);
        parentNode.removeChild(ele);
        // 克隆一个副本，防止被编译，破坏模板结构
        ele.removeAttribute('$for');
        let cloneOrigin = ele.cloneNode(true);

        Array.isArray(list) && list.forEach((item, index) => {
            let cloneNode = cloneOrigin.cloneNode(true);
            self.compile$for(cloneNode, expItem, item, expIndex, index);
            divWrap.appendChild(cloneNode);
        });

        new Watcher(exp, vm, (newValue) => {
            Array.from(divWrap.childNodes).forEach((child) => {
                divWrap.removeChild(child);
            });

            newValue.forEach((item, index) => {
                let cloneNode = cloneOrigin.cloneNode(true);
                self.compile$for(cloneNode, expItem, item, expIndex, index);
                divWrap.appendChild(cloneNode);
            });
        });
    };

    /**
     * 编译$for指令
     * @param {node} $for指令所在的元素
     * @param {string} (item, index)中的item表达式
     * @param {any} (item, index)中的item表达式的实际的值
     * @param {string} (item, index)中的index表达式
     * @param {number} (item, index)中的index表达式的实际的值
     */
    Compiler.prototype.compile$for = function(ele, expItem, item, expIndex, index) {
        let self = this;
        let regItem = new RegExp(`{{${expItem}}}`, 'g');
        let regIndex = new RegExp(`{{${expIndex}}}`, 'g');

        function recursion(ele) {
            if (!ele.childNodes || !ele.childNodes.length) {
                return;
            }

            Array.from(ele.childNodes).forEach((child) => {
                if (child.nodeType === 3 && !regBlank.test(child.textContent)) {
                    let content = child.textContent.trim();
                    let str = self.replace$for(content, item, regItem);

                    // 如果有索引表达式
                    if (expIndex) {
                        str = self.replace$for(str, index, regIndex);
                    }
                    child.textContent = str;
                } else if (child.nodeType === 1) {
                    // $click如果传入了expItem或者expIndex参数，把它转成实际的值
                    Array.from(child.attributes).forEach((attr) => {
                        if (attr.name === '$click') {
                            let exp = attr.value;
                            let match = regCall.exec(exp);

                            if (match) {
                                let funcName = exp.slice(0, exp.indexOf('('));
                                if (match[1] === expItem) {
                                    attr.value = `${funcName}('${item}')`;
                                } else if (match[1] === expIndex) {
                                    attr.value = `${funcName}('${index}')`;
                                }
                            }
                        }
                    });
                }
                recursion(child);
            });
        }
        recursion(ele);
        // 编译其他指令和模板
        this.compile(ele);
    };

    /**
     * 替换$for指令所在元素的子元素的{{item}}或者{{index}}模板为实际的值
     * @param {string} $for指令所在元素的子元素的文本
     * @param {any} $for指令所在元素的子元素的{{item}}或者{{index}}模板实际的值
     * @param {regexp} 识别{{item}}或者{{index}}模板的正则表达式
     */
    Compiler.prototype.replace$for = function(content, value, reg) {
        let i = 0;
        let match;
        let text;
        let temp = '';

        while (match = reg.exec(content)) {
            if (i < match.index) {
                text = content.slice(i, match.index);
                temp += text;
            }
            i = reg.lastIndex;
            temp += value;
        }
        if (i < content.length) {
            text = content.slice(i);
            temp += text;
        }
        return temp;
    };

    /**
     * $model指令
     * @用法：<input type="checkbox" $model="exp">
     * @用法：<input type="radio" $model="exp" value="{{other exp}}">
     * @用法：<input type="file" $model="exp">
     * @用法：<input type="text" $model="exp">
     * 只能用于输入框，数据双向绑定
     * radio输入框的$model指令和value属性配合使用，获取选中的值
     * 建议多个checkbox输入框与同一个数组的元素绑定
     */
    Compiler.prototype.$model = function(ele, exp, vm) {
        if (ele.tagName.toLowerCase() !== 'input') {
            return;
        }

        let result = execChain(exp, vm);
        if (result === undefined) {
            console.error(`[Biu warn]无法找到$model中表达式${exp}的值`);
            return;
        }

        let type = ele.type;
        switch (type) {
            case 'checkbox':
                ele.checked = result;

                new Watcher(exp, vm, (newValue) => {
                    ele.checked = newValue;
                });

                ele.addEventListener('change', function(event) {
                    let newValue = event.target.checked;
                    assignChain(exp, vm, newValue);
                });
                break;

            case 'radio':
                let modelValue = result;
                let value = ele.value;
                let radioValue = this.compileInputValue(exp, value);
                modelValue == radioValue ? ele.checked = true : ele.checked = false;

                new Watcher(exp, vm, (newValue) => {
                    newValue == radioValue ? ele.checked = true : ele.checked = false;
                });

                ele.addEventListener('change', function(event) {
                    let newValue = event.target.value;
                    assignChain(exp, vm, newValue);
                });
                break;

            case 'file':
                ele.addEventListener('change', function(event) {
                    let newValue = event.target.files;
                    assignChain(exp, vm, newValue);
                });
                break;

            default:
                ele.value = result;

                new Watcher(exp, vm, (newValue) => {
                    // input事件和Watcher循环触发，会导致中文输入法不可用，加一个flag判断
                    if (!ele.isInputting) {
                        ele.value = newValue;
                    }
                    ele.isInputting = false;
                });

                ele.addEventListener('input', function(event) {
                    ele.isInputting = true;
                    let newValue = event.target.value;
                    assignChain(exp, vm, newValue);
                });
        }
    };

    /**
     * 编译输入框value属性的表达式
     * @param {string} value属性的表达式
     * @param {any} value的值
     */
    Compiler.prototype.compileInputValue = function(exp, value) {
        if (value.match(regMustache)) {
            let match = regMustache.exec(value)[1];

            let result = execChain(match, this.vm);
            if (result === undefined) {
                console.error(`[Biu warn]无法找到value属性中表达式${exp}的值`);
                return;
            }

            value = result;
        }
        return value;
    };

    /**
     * $show指令
     * @用法：<div $show="exp"></div>
     * @用法：<div $show="!exp"></div>
     */
    Compiler.prototype.$show = function(ele, exp, vm) {
        let exclamation;
        if (/^\!/.test(exp)) {
            exclamation = true;
            exp = exp.slice(1);
            let result = execChain(exp, vm);
            if (result === undefined) {
                console.error(`[Biu warn]无法找到$show中表达式${exp}的值`);
                return;
            }
            ele.style.display = result ? 'none' : 'initial';
        } else {
            exclamation = false;
            let result = execChain(exp, vm);
            if (result === undefined) {
                console.error(`[Biu warn]无法找到$show中表达式${exp}的值`);
                return;
            }
            ele.style.display = result ? 'initial' : 'none';
        }

        new Watcher(exp, vm, (newValue) => {
            if (exclamation) {
                ele.style.display = newValue ? 'none' : 'initial';
            } else {
                ele.style.display = newValue ? 'initial' : 'none';
            }
        });
    };

    /**
     * $if指令
     * @用法：<div $if="exp"></div>
     * @用法：<div $if="!exp"></div>
     */
    Compiler.prototype.$if = function(ele, exp, vm) {
        let exclamation;
        let refNode = document.createTextNode('');
        ele.parentNode.insertBefore(refNode, ele);

        if (/^\!/.test(exp)) {
            exclamation = true;
            exp = exp.slice(1);
            let result = execChain(exp, vm);
            if (result === undefined) {
                console.error(`[Biu warn]无法找到$if中表达式${exp}的值`);
                return;
            }
            if (result) {
                // 元素被移除之前，先编译子元素
                this.compile(ele);
                ele.parentNode.removeChild(ele);
            }
        } else {
            exclamation = false;
            let result = execChain(exp, vm);
            if (result === undefined) {
                console.error(`[Biu warn]无法找到$if中表达式${exp}的值`);
                return;
            }
            if (!result) {
                // 元素被移除之前，先编译子元素
                this.compile(ele);
                ele.parentNode.removeChild(ele);
            }
        }

        new Watcher(exp, vm, (newValue) => {
            if (exclamation) {
                if (newValue) {
                    refNode.parentNode.removeChild(ele);
                } else {
                    refNode.parentNode.insertBefore(ele, refNode);
                }
            } else {
                if (newValue) {
                    refNode.parentNode.insertBefore(ele, refNode);
                } else {
                    refNode.parentNode.removeChild(ele);
                }
            }
        });
    };

    /**
     * $class指令
     * @用法：<div class="box" $class="exp"></div>
     * exp的数据类型是数组，不会覆盖用户写入DOM的class
     */
    Compiler.prototype.$class = function(ele, exp, vm) {
        let result = execChain(exp, vm);
        if (result === undefined) {
            console.error(`[Biu warn]无法找到$class中表达式${exp}的值`);
            return;
        }

        let classArr = result;
        let classList = ele.classList;
        Array.isArray(classArr) && classArr.forEach((value) => {
            classList.add(value);
        });
        let oldValue = [].slice.call(classArr);

        new Watcher(exp, vm, (newValue) => {
            oldValue.forEach((value) => {
                classList.remove(value);
            });
            newValue.forEach((value) => {
                classList.add(value);
            });
            oldValue = [].slice.call(newValue);
        });
    };

    /**
     * $style指令
     * @用法：<div style="background: #f00;" $style="exp"></div>
     * exp的数据类型是数组，不会覆盖用户写入DOM的style
     */
    Compiler.prototype.$style = function(ele, exp, vm) {
        let self = this;
        let cssText = ele.style.cssText;
        let userStyle = cssText.split(';');
        userStyle.pop();

        let result = execChain(exp, vm);
        if (result === undefined) {
            console.error(`[Biu warn]无法找到$style中表达式${exp}的值`);
            return;
        }

        let biuStyle = result;
        let coverStyle = getCoverSet(biuStyle, userStyle);
        if (Array.isArray(biuStyle)) {
            ele.style.cssText = cssText + self.arrayToStyle(biuStyle);
        } else {
            return;
        }

        new Watcher(exp, vm, (newValue) => {
            let mixStyle = ele.style.cssText.split(';');
            mixStyle.pop();
            biuStyle = self.switchStyleColor(biuStyle);
            userStyle = getComplementSet(biuStyle, mixStyle);
            userStyle = [].concat(userStyle, coverStyle);
            // newValue应该放最后
            mixStyle = [].concat(userStyle, newValue);
            ele.style.cssText = self.arrayToStyle(mixStyle);

            // 新值赋值给旧值，方便下次比较
            biuStyle = [].slice.call(newValue);
            // 重新计算被覆盖的样式
            coverStyle = getCoverSet(biuStyle, userStyle);
        });
    };

    /**
     * 切换hex颜色为rgb颜色，在$style中用到
     * @param {array} 以样式为元素的数组
     */
     Compiler.prototype.switchStyleColor = function(styleArr) {
        let newStyleArr = styleArr.map((styleItem) => {
            styleItem = styleItem.trim();

            if (styleItem.indexOf('#') > -1) {
                let arr = styleItem.split(':');
                let temp = arr[1];
                temp = hexToRgb(temp);
                return `${arr[0]}: ${temp}`;
            } else {
                return styleItem;
            }
        });
        return newStyleArr;
     };

    /**
     * 样式数组转成样式字符串，在$style中用到
     * @param {array} 以样式为元素的数组
     */
     Compiler.prototype.arrayToStyle = function(styleArr) {
        let temp = '';
        styleArr.forEach((value) => {
            if (value.indexOf(';') < 0) {
                temp += `${value};`;
            } else {
                temp += value;
            }
        });
        return temp;
     };

    /**
     * $text指令
     * @用法：<div $text="exp"></div>
     * exp的数据类型是字符串，会覆盖子文本节点
     */
    Compiler.prototype.$text = function(ele, exp, vm) {
        let result = execChain(exp, vm);
        if (result === undefined) {
            console.error(`[Biu warn]无法找到$text中表达式${exp}的值`);
            return;
        }
        ele.textContent = result;

        new Watcher(exp, vm, (newValue) => {
            ele.textContent = newValue;
        });
    };

    /**
     * $html指令
     * @用法：<div $html="exp"></div>
     * exp的数据类型是node，会覆盖子元素节点
     */
    Compiler.prototype.$html = function(ele, exp, vm) {
        let result = execChain(exp, vm);
        if (result === undefined) {
            console.error(`[Biu warn]无法找到$html中表达式${exp}的值`);
            return;
        }
        ele.innerHTML = result;

        new Watcher(exp, vm, (newValue) => {
            ele.innerHTML = newValue;
        });
    };

    /**
     * $click指令
     * @用法：<div $click="func(exp)"></div>
     * @用法：<div $click="func('str')"></div>
     * exp可以是$for指令里的item或者index，或其他数据
     */
    Compiler.prototype.$click = function(ele, exp, vm) {
        let expFunc;
        if (regCall.test(exp)) {
            expFunc = exp.slice(0, exp.indexOf('('));
        } else {
            console.error('[Biu warn]$click指令中的表达式必须加圆括号');
            return;
        }
        // action没有嵌套，直接用[]操作符
        let func = vm.$action[expFunc];
        let args = regCall.exec(exp)[1];
        let match = regQuote.exec(args);

        // 判断参数是否为字符串
        if (match) {
            args = match[1];
        } else {
            if (args) {
                let result = execChain(args, vm);
                if (result === undefined) {
                    console.error(`[Biu warn]无法找到$click中表达式${args}的值`);
                    return;
                }
                args = result;
            }
        }

        ele.addEventListener('click', () => {
            func.call(ele, args);
        });
    };

    /////////////// *Watcher* ///////////////

    /**
     * 观察者主体方法
     */
    function Watcher(exp, vm, cb) {
        this.exp = exp;
        this.vm = vm;
        this.update = cb;
        this.trigger();
    }

    /**
     * 触发订阅
     */
    Watcher.prototype.trigger = function() {
        // 缓存订阅项
        Dep.target = this;
        // 触发getter，添加订阅
        execChain(this.exp, this.vm);
        // 清空订阅项缓存
        Dep.target = null;
    };

    /////////////// *core* ///////////////

    /**
     * Biu主体方法
     */
    function Biu(options) {
        this.checkBiuOptions(options);
        let mountId = options.mount.slice(1);
        this.$mount = document.getElementById(mountId);
        this.$data = options.data;
        this.$action = options.action;
        new Observer(this.$data);
        new Compiler(this.$mount, this);
    }

    /**
     * 校验主体函数的参数
     * @param {object} 传入主体函数的参数
     */
    Biu.prototype.checkBiuOptions = function(options) {
        if (Object.prototype.toString.call(options) !== '[object Object]') {
            throw new Error('[Biu error]Biu的参数不是对象');
        }

        let mount = options.mount;
        if (!mount) {
            throw new Error('[Biu error]Biu找不到挂载点');
        }
        if (mount.slice(0, 1) !== '#') {
            throw new Error('[Biu error]Biu的挂载点选择器必须是id');
        }

        let data = options.data;
        if (!data) {
            throw new Error('[Biu error]Biu找不到数据模型');
        }

        let action = options.action;
        if (action) {
            for (let key in action) {
                if (typeof action[key] !== 'function') {
                    throw new Error('[Biu error]Biu的方法集合的属性必须是function');
                }
            }
        }
    };

    // 挂载到全局变量上
    root.Biu = Biu;

})(window);
