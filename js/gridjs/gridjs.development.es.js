/*! *****************************************************************************
Copyright (c) Microsoft Corporation.

Permission to use, copy, modify, and/or distribute this software for any
purpose with or without fee is hereby granted.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR
OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
PERFORMANCE OF THIS SOFTWARE.
***************************************************************************** */
/* global Reflect, Promise */

var extendStatics = function(d, b) {
    extendStatics = Object.setPrototypeOf ||
        ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
        function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
    return extendStatics(d, b);
};

function __extends(d, b) {
    extendStatics(d, b);
    function __() { this.constructor = d; }
    d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
}

var __assign = function() {
    __assign = Object.assign || function __assign(t) {
        for (var s, i = 1, n = arguments.length; i < n; i++) {
            s = arguments[i];
            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p)) t[p] = s[p];
        }
        return t;
    };
    return __assign.apply(this, arguments);
};

function __awaiter(thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
}

function __generator(thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (_) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
}

function __spreadArrays() {
    for (var s = 0, i = 0, il = arguments.length; i < il; i++) s += arguments[i].length;
    for (var r = Array(s), k = 0, i = 0; i < il; i++)
        for (var a = arguments[i], j = 0, jl = a.length; j < jl; j++, k++)
            r[k] = a[j];
    return r;
}

function generateID() {
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
        var r = (Math.random() * 16) | 0, v = c == 'x' ? r : (r & 0x3) | 0x8;
        return v.toString(16);
    });
}

var Base = /** @class */ (function () {
    function Base(id) {
        this._id = id || generateID();
    }
    Object.defineProperty(Base.prototype, "id", {
        get: function () {
            return this._id;
        },
        enumerable: false,
        configurable: true
    });
    return Base;
}());

var Cell = /** @class */ (function (_super) {
    __extends(Cell, _super);
    function Cell(data) {
        var _this = _super.call(this) || this;
        _this.data = data;
        return _this;
    }
    Object.defineProperty(Cell.prototype, "data", {
        get: function () {
            return this._data;
        },
        set: function (data) {
            this._data = data;
        },
        enumerable: false,
        configurable: true
    });
    return Cell;
}(Base));

var Row = /** @class */ (function (_super) {
    __extends(Row, _super);
    function Row(cells) {
        var _this = _super.call(this) || this;
        _this.cells = cells || [];
        return _this;
    }
    Object.defineProperty(Row.prototype, "cells", {
        get: function () {
            return this._cells;
        },
        set: function (cells) {
            this._cells = cells;
        },
        enumerable: false,
        configurable: true
    });
    /**
     * Creates a new Row from an array of Cell(s)
     * This method generates a new ID for the Row and all nested elements
     *
     * @param cells
     * @returns Row
     */
    Row.fromCells = function (cells) {
        return new Row(cells.map(function (cell) { return new Cell(cell.data); }));
    };
    Object.defineProperty(Row.prototype, "length", {
        get: function () {
            return this.cells.length;
        },
        enumerable: false,
        configurable: true
    });
    return Row;
}(Base));

function oneDtoTwoD(data) {
    if (data[0] && !(data[0] instanceof Array)) {
        return [data];
    }
    return data;
}

var Tabular = /** @class */ (function (_super) {
    __extends(Tabular, _super);
    function Tabular(rows) {
        var _this = _super.call(this) || this;
        if (rows instanceof Array) {
            _this.rows = rows;
        }
        else if (rows instanceof Row) {
            _this.rows = [rows];
        }
        else {
            _this.rows = [];
        }
        return _this;
    }
    Object.defineProperty(Tabular.prototype, "rows", {
        get: function () {
            return this._rows;
        },
        set: function (rows) {
            this._rows = rows;
        },
        enumerable: false,
        configurable: true
    });
    Object.defineProperty(Tabular.prototype, "length", {
        get: function () {
            return this._length || this.rows.length;
        },
        // we want to sent the length when storage is ServerStorage
        set: function (len) {
            this._length = len;
        },
        enumerable: false,
        configurable: true
    });
    /**
     * Creates a new Tabular from an array of Row(s)
     * This method generates a new ID for the Tabular and all nested elements
     *
     * @param rows
     * @returns Tabular
     */
    Tabular.fromRows = function (rows) {
        return new Tabular(rows.map(function (row) { return Row.fromCells(row.cells); }));
    };
    /**
     * Creates a new Tabular from a 2D array
     * This method generates a new ID for the Tabular and all nested elements
     *
     * @param data
     * @returns Tabular
     */
    Tabular.fromArray = function (data) {
        data = oneDtoTwoD(data);
        return new Tabular(data.map(function (row) { return new Row(row.map(function (cell) { return new Cell(cell); })); }));
    };
    Tabular.fromStorageResponse = function (storageResponse) {
        var tabular = Tabular.fromArray(storageResponse.data);
        // for server-side storage
        tabular.length = storageResponse.total;
        return tabular;
    };
    return Tabular;
}(Base));

function width(width, containerWidth) {
    if (typeof width == 'string') {
        if (width.indexOf('%') > -1) {
            return (containerWidth / 100) * parseInt(width, 10);
        }
        else {
            return parseInt(width, 10);
        }
    }
    return width;
}
function px(width) {
    if (!width)
        return '';
    return Math.floor(width) + "px";
}
/**
 * Accepts a ShadowTable and tries to find the clientWidth
 * that is already rendered on the web browser
 *
 * @param shadowTable
 * @param columnIndex
 */
function getWidth(shadowTable, columnIndex) {
    if (!shadowTable) {
        return null;
    }
    var tds = shadowTable.querySelectorAll('tr:first-child > td');
    if (tds && tds[columnIndex]) {
        return tds[columnIndex].clientWidth;
    }
    return null;
}

var n,u,i,t,o,r,f,e={},c=[],a=/acit|ex(?:s|g|n|p|$)|rph|grid|ows|mnc|ntw|ine[ch]|zoo|^ord|itera/i;function s(n,l){for(var u in l)n[u]=l[u];return n}function v(n){var l=n.parentNode;l&&l.removeChild(n);}function h(n,l,u){var i,t=arguments,o={};for(i in l)"key"!==i&&"ref"!==i&&(o[i]=l[i]);if(arguments.length>3)for(u=[u],i=3;i<arguments.length;i++)u.push(t[i]);if(null!=u&&(o.children=u),"function"==typeof n&&null!=n.defaultProps)for(i in n.defaultProps)void 0===o[i]&&(o[i]=n.defaultProps[i]);return y(n,o,l&&l.key,l&&l.ref,null)}function y(l,u,i,t,o){var r={type:l,props:u,key:i,ref:t,__k:null,__:null,__b:0,__e:null,__d:void 0,__c:null,constructor:void 0,__v:o};return null==o&&(r.__v=r),n.vnode&&n.vnode(r),r}function p(){return {}}function d(n){return n.children}function m(n,l){this.props=n,this.context=l;}function w(n,l){if(null==l)return n.__?w(n.__,n.__.__k.indexOf(n)+1):null;for(var u;l<n.__k.length;l++)if(null!=(u=n.__k[l])&&null!=u.__e)return u.__e;return "function"==typeof n.type?w(n):null}function k(n){var l,u;if(null!=(n=n.__)&&null!=n.__c){for(n.__e=n.__c.base=null,l=0;l<n.__k.length;l++)if(null!=(u=n.__k[l])&&null!=u.__e){n.__e=n.__c.base=u.__e;break}return k(n)}}function g(l){(!l.__d&&(l.__d=!0)&&u.push(l)&&!i++||o!==n.debounceRendering)&&((o=n.debounceRendering)||t)(_);}function _(){for(var n;i=u.length;)n=u.sort(function(n,l){return n.__v.__b-l.__v.__b}),u=[],n.some(function(n){var l,u,i,t,o,r,f;n.__d&&(r=(o=(l=n).__v).__e,(f=l.__P)&&(u=[],(i=s({},o)).__v=i,t=z(f,o,i,l.__n,void 0!==f.ownerSVGElement,null,u,null==r?w(o):r),T(u,o),t!=r&&k(o)));});}function b(n,l,u,i,t,o,r,f,a,s){var h,p,m,k,g,_,b,x,A,P=i&&i.__k||c,C=P.length;for(a==e&&(a=null!=r?r[0]:C?w(i,0):null),u.__k=[],h=0;h<l.length;h++)if(null!=(k=u.__k[h]=null==(k=l[h])||"boolean"==typeof k?null:"string"==typeof k||"number"==typeof k?y(null,k,null,null,k):Array.isArray(k)?y(d,{children:k},null,null,null):null!=k.__e||null!=k.__c?y(k.type,k.props,k.key,null,k.__v):k)){if(k.__=u,k.__b=u.__b+1,null===(m=P[h])||m&&k.key==m.key&&k.type===m.type)P[h]=void 0;else for(p=0;p<C;p++){if((m=P[p])&&k.key==m.key&&k.type===m.type){P[p]=void 0;break}m=null;}if(g=z(n,k,m=m||e,t,o,r,f,a,s),(p=k.ref)&&m.ref!=p&&(x||(x=[]),m.ref&&x.push(m.ref,null,k),x.push(p,k.__c||g,k)),null!=g){if(null==b&&(b=g),A=void 0,void 0!==k.__d)A=k.__d,k.__d=void 0;else if(r==m||g!=a||null==g.parentNode){n:if(null==a||a.parentNode!==n)n.appendChild(g),A=null;else {for(_=a,p=0;(_=_.nextSibling)&&p<C;p+=2)if(_==g)break n;n.insertBefore(g,a),A=a;}"option"==u.type&&(n.value="");}a=void 0!==A?A:g.nextSibling,"function"==typeof u.type&&(u.__d=a);}else a&&m.__e==a&&a.parentNode!=n&&(a=w(m));}if(u.__e=b,null!=r&&"function"!=typeof u.type)for(h=r.length;h--;)null!=r[h]&&v(r[h]);for(h=C;h--;)null!=P[h]&&D(P[h],P[h]);if(x)for(h=0;h<x.length;h++)j(x[h],x[++h],x[++h]);}function A(n,l,u,i,t){var o;for(o in u)"children"===o||"key"===o||o in l||C(n,o,null,u[o],i);for(o in l)t&&"function"!=typeof l[o]||"children"===o||"key"===o||"value"===o||"checked"===o||u[o]===l[o]||C(n,o,l[o],u[o],i);}function P(n,l,u){"-"===l[0]?n.setProperty(l,u):n[l]="number"==typeof u&&!1===a.test(l)?u+"px":null==u?"":u;}function C(n,l,u,i,t){var o,r,f,e,c;if(t?"className"===l&&(l="class"):"class"===l&&(l="className"),"style"===l)if(o=n.style,"string"==typeof u)o.cssText=u;else {if("string"==typeof i&&(o.cssText="",i=null),i)for(e in i)u&&e in u||P(o,e,"");if(u)for(c in u)i&&u[c]===i[c]||P(o,c,u[c]);}else "o"===l[0]&&"n"===l[1]?(r=l!==(l=l.replace(/Capture$/,"")),f=l.toLowerCase(),l=(f in n?f:l).slice(2),u?(i||n.addEventListener(l,N,r),(n.l||(n.l={}))[l]=u):n.removeEventListener(l,N,r)):"list"!==l&&"tagName"!==l&&"form"!==l&&"type"!==l&&"size"!==l&&!t&&l in n?n[l]=null==u?"":u:"function"!=typeof u&&"dangerouslySetInnerHTML"!==l&&(l!==(l=l.replace(/^xlink:?/,""))?null==u||!1===u?n.removeAttributeNS("http://www.w3.org/1999/xlink",l.toLowerCase()):n.setAttributeNS("http://www.w3.org/1999/xlink",l.toLowerCase(),u):null==u||!1===u&&!/^ar/.test(l)?n.removeAttribute(l):n.setAttribute(l,u));}function N(l){this.l[l.type](n.event?n.event(l):l);}function z(l,u,i,t,o,r,f,e,c){var a,v,h,y,p,w,k,g,_,x,A,P=u.type;if(void 0!==u.constructor)return null;(a=n.__b)&&a(u);try{n:if("function"==typeof P){if(g=u.props,_=(a=P.contextType)&&t[a.__c],x=a?_?_.props.value:a.__:t,i.__c?k=(v=u.__c=i.__c).__=v.__E:("prototype"in P&&P.prototype.render?u.__c=v=new P(g,x):(u.__c=v=new m(g,x),v.constructor=P,v.render=E),_&&_.sub(v),v.props=g,v.state||(v.state={}),v.context=x,v.__n=t,h=v.__d=!0,v.__h=[]),null==v.__s&&(v.__s=v.state),null!=P.getDerivedStateFromProps&&(v.__s==v.state&&(v.__s=s({},v.__s)),s(v.__s,P.getDerivedStateFromProps(g,v.__s))),y=v.props,p=v.state,h)null==P.getDerivedStateFromProps&&null!=v.componentWillMount&&v.componentWillMount(),null!=v.componentDidMount&&v.__h.push(v.componentDidMount);else {if(null==P.getDerivedStateFromProps&&g!==y&&null!=v.componentWillReceiveProps&&v.componentWillReceiveProps(g,x),!v.__e&&null!=v.shouldComponentUpdate&&!1===v.shouldComponentUpdate(g,v.__s,x)||u.__v===i.__v){for(v.props=g,v.state=v.__s,u.__v!==i.__v&&(v.__d=!1),v.__v=u,u.__e=i.__e,u.__k=i.__k,v.__h.length&&f.push(v),a=0;a<u.__k.length;a++)u.__k[a]&&(u.__k[a].__=u);break n}null!=v.componentWillUpdate&&v.componentWillUpdate(g,v.__s,x),null!=v.componentDidUpdate&&v.__h.push(function(){v.componentDidUpdate(y,p,w);});}v.context=x,v.props=g,v.state=v.__s,(a=n.__r)&&a(u),v.__d=!1,v.__v=u,v.__P=l,a=v.render(v.props,v.state,v.context),null!=v.getChildContext&&(t=s(s({},t),v.getChildContext())),h||null==v.getSnapshotBeforeUpdate||(w=v.getSnapshotBeforeUpdate(y,p)),A=null!=a&&a.type==d&&null==a.key?a.props.children:a,b(l,Array.isArray(A)?A:[A],u,i,t,o,r,f,e,c),v.base=u.__e,v.__h.length&&f.push(v),k&&(v.__E=v.__=null),v.__e=!1;}else null==r&&u.__v===i.__v?(u.__k=i.__k,u.__e=i.__e):u.__e=$(i.__e,u,i,t,o,r,f,c);(a=n.diffed)&&a(u);}catch(l){u.__v=null,n.__e(l,u,i);}return u.__e}function T(l,u){n.__c&&n.__c(u,l),l.some(function(u){try{l=u.__h,u.__h=[],l.some(function(n){n.call(u);});}catch(l){n.__e(l,u.__v);}});}function $(n,l,u,i,t,o,r,f){var a,s,v,h,y,p=u.props,d=l.props;if(t="svg"===l.type||t,null!=o)for(a=0;a<o.length;a++)if(null!=(s=o[a])&&((null===l.type?3===s.nodeType:s.localName===l.type)||n==s)){n=s,o[a]=null;break}if(null==n){if(null===l.type)return document.createTextNode(d);n=t?document.createElementNS("http://www.w3.org/2000/svg",l.type):document.createElement(l.type,d.is&&{is:d.is}),o=null,f=!1;}if(null===l.type)p!==d&&n.data!=d&&(n.data=d);else {if(null!=o&&(o=c.slice.call(n.childNodes)),v=(p=u.props||e).dangerouslySetInnerHTML,h=d.dangerouslySetInnerHTML,!f){if(null!=o)for(p={},y=0;y<n.attributes.length;y++)p[n.attributes[y].name]=n.attributes[y].value;(h||v)&&(h&&v&&h.__html==v.__html||(n.innerHTML=h&&h.__html||""));}A(n,d,p,t,f),h?l.__k=[]:(a=l.props.children,b(n,Array.isArray(a)?a:[a],l,u,i,"foreignObject"!==l.type&&t,o,r,e,f)),f||("value"in d&&void 0!==(a=d.value)&&a!==n.value&&C(n,"value",a,p.value,!1),"checked"in d&&void 0!==(a=d.checked)&&a!==n.checked&&C(n,"checked",a,p.checked,!1));}return n}function j(l,u,i){try{"function"==typeof l?l(u):l.current=u;}catch(l){n.__e(l,i);}}function D(l,u,i){var t,o,r;if(n.unmount&&n.unmount(l),(t=l.ref)&&(t.current&&t.current!==l.__e||j(t,null,u)),i||"function"==typeof l.type||(i=null!=(o=l.__e)),l.__e=l.__d=void 0,null!=(t=l.__c)){if(t.componentWillUnmount)try{t.componentWillUnmount();}catch(l){n.__e(l,u);}t.base=t.__P=null;}if(t=l.__k)for(r=0;r<t.length;r++)t[r]&&D(t[r],u,i);null!=o&&v(o);}function E(n,l,u){return this.constructor(n,u)}function H(l,u,i){var t,o,f;n.__&&n.__(l,u),o=(t=i===r)?null:i&&i.__k||u.__k,l=h(d,null,[l]),f=[],z(u,(t?u:i||u).__k=l,o||e,e,void 0!==u.ownerSVGElement,i&&!t?[i]:o?null:u.childNodes.length?c.slice.call(u.childNodes):null,f,i||e,t),T(f,l);}n={__e:function(n,l){for(var u,i;l=l.__;)if((u=l.__c)&&!u.__)try{if(u.constructor&&null!=u.constructor.getDerivedStateFromError&&(i=!0,u.setState(u.constructor.getDerivedStateFromError(n))),null!=u.componentDidCatch&&(i=!0,u.componentDidCatch(n)),i)return g(u.__E=u)}catch(l){n=l;}throw n}},m.prototype.setState=function(n,l){var u;u=this.__s!==this.state?this.__s:this.__s=s({},this.state),"function"==typeof n&&(n=n(u,this.props)),n&&s(u,n),null!=n&&this.__v&&(l&&this.__h.push(l),g(this));},m.prototype.forceUpdate=function(n){this.__v&&(this.__e=!0,n&&this.__h.push(n),g(this));},m.prototype.render=d,u=[],i=0,t="function"==typeof Promise?Promise.prototype.then.bind(Promise.resolve()):setTimeout,r=e,f=0;

var BaseComponent = /** @class */ (function (_super) {
    __extends(BaseComponent, _super);
    function BaseComponent() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    return BaseComponent;
}(m));

var ShadowTable = /** @class */ (function (_super) {
    __extends(ShadowTable, _super);
    function ShadowTable() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    ShadowTable.prototype.resetStyle = function () {
        return { padding: 0, margin: 0, border: 'none' };
    };
    ShadowTable.prototype.head = function () {
        var _this = this;
        return (h("thead", { style: this.resetStyle() },
            h("tr", null, this.props.header.columns.map(function (col) {
                return h("th", { style: _this.resetStyle() }, col.name);
            }))));
    };
    ShadowTable.prototype.td = function (cell) {
        return h("td", { style: this.resetStyle() }, cell.data);
    };
    ShadowTable.prototype.tr = function (row) {
        var _this = this;
        return (h("tr", { style: this.resetStyle() }, row.cells.map(function (cell) {
            return _this.td(cell);
        })));
    };
    ShadowTable.prototype.body = function () {
        var _this = this;
        return (h("tbody", { style: this.resetStyle() }, this.props.data &&
            this.props.data.rows.map(function (row) {
                return _this.tr(row);
            })));
    };
    ShadowTable.prototype.render = function () {
        return (h("table", { style: __assign({ position: 'absolute', zIndex: '-2147483640', visibility: 'hidden', tableLayout: 'auto', width: 'auto' }, this.resetStyle()) },
            this.head(),
            this.body()));
    };
    return ShadowTable;
}(BaseComponent));

var Header = /** @class */ (function (_super) {
    __extends(Header, _super);
    function Header() {
        var _this = _super.call(this) || this;
        _this._columns = [];
        return _this;
    }
    Object.defineProperty(Header.prototype, "columns", {
        get: function () {
            return this._columns;
        },
        set: function (columns) {
            this._columns = columns;
        },
        enumerable: false,
        configurable: true
    });
    /**
     * Tries to automatically adjust the width of columns based on:
     *    - Header cell content
     *    - Cell content of the first row
     *    - Cell content of the last row
     * @param autoWidth
     * @param container
     * @param data
     */
    Header.prototype.adjustWidth = function (container, data, autoWidth) {
        if (autoWidth === void 0) { autoWidth = true; }
        if (!container) {
            // we can't calculate the width because the container
            // is unknown at this stage
            return this;
        }
        // pixels
        var containerWidth = container.clientWidth;
        // let's create a shadow table with the first 10 rows of the data
        // and let the browser to render the table with table-layout: auto
        // no padding, margin or border to get the minimum space required
        // to render columns. One the table is rendered and widths are known,
        // we unmount the shadow table from the DOM and set the correct width
        var shadowTable = p();
        if (data && data.length && autoWidth) {
            // render a ShadowTable with the first 10 rows
            var el = h(ShadowTable, {
                data: Tabular.fromRows(data.rows.slice(0, 10)),
                header: this,
            });
            el.ref = shadowTable;
            // TODO: we should NOT query the container here. use Refs instead
            H(el, container.querySelector('#gridjs-temp'));
        }
        for (var _i = 0, _a = this.columns; _i < _a.length; _i++) {
            var column = _a[_i];
            if (!column.width && autoWidth) {
                var i = this.columns.indexOf(column);
                // tries to find the corresponding cell from the ShadowTable and
                // set the correct width
                column.width = px(getWidth(shadowTable.current.base, i));
            }
            else {
                column.width = px(width(column.width, containerWidth));
            }
        }
        if (data && data.length && autoWidth) {
            // unmount the shadow table from temp
            H(null, container.querySelector('#gridjs-temp'));
        }
        return this;
    };
    Header.prototype.setSort = function (userConfig) {
        for (var _i = 0, _a = this.columns; _i < _a.length; _i++) {
            var column = _a[_i];
            // implicit userConfig.sort flag
            if (column.sort === undefined && userConfig.sort) {
                column.sort = {
                    enabled: true,
                };
            }
            // false, null, etc.
            if (!column.sort) {
                column.sort = {
                    enabled: false,
                };
            }
            else if (typeof column.sort === 'object') {
                column.sort = __assign({ enabled: true }, column.sort);
            }
        }
    };
    Header.fromUserConfig = function (userConfig) {
        // because we should be able to render a table without the header
        if (!userConfig.columns && !userConfig.from) {
            return null;
        }
        var header = new Header();
        if (userConfig.from) {
            header.columns = Header.fromHTMLTable(userConfig.from).columns;
        }
        else {
            header.columns = [];
            for (var _i = 0, _a = userConfig.columns; _i < _a.length; _i++) {
                var column = _a[_i];
                if (typeof column === 'string') {
                    header.columns.push({
                        name: column,
                    });
                }
                else if (typeof column === 'object') {
                    header.columns.push(column);
                }
            }
        }
        header.setSort(userConfig);
        return header;
    };
    Header.fromHTMLTable = function (element) {
        var header = new Header();
        var thead = element.querySelector('thead');
        var ths = thead.querySelectorAll('th');
        for (var _i = 0, _a = ths; _i < _a.length; _i++) {
            var th = _a[_i];
            header.columns.push({
                name: th.innerText,
                width: th.width,
            });
        }
        return header;
    };
    return Header;
}(Base));

var Config = /** @class */ (function () {
    function Config(userConfig) {
        // FIXME: not sure if this makes sense because Config is a subset of UserConfig
        var updatedConfig = __assign(__assign({}, Config.defaultConfig()), userConfig);
        Object.assign(this, updatedConfig);
    }
    Config.defaultConfig = function () {
        return {
            width: '100%',
            autoWidth: true,
        };
    };
    Config.fromUserConfig = function (userConfig) {
        var config = new Config(userConfig);
        if (!userConfig)
            return config;
        if (typeof config.sort === 'boolean' && config.sort) {
            config.sort = {
                multiColumn: true,
            };
        }
        config.header = Header.fromUserConfig(config);
        config.pagination = __assign({ enabled: userConfig.pagination === true ||
                userConfig.pagination instanceof Object }, userConfig.pagination);
        config.search = __assign({ enabled: userConfig.search === true || userConfig.search instanceof Object }, userConfig.search);
        return config;
    };
    return Config;
}());

/**
 * Base Storage class. All storage implementation must inherit this class
 */
var Storage = /** @class */ (function () {
    function Storage() {
    }
    return Storage;
}());

var MemoryStorage = /** @class */ (function (_super) {
    __extends(MemoryStorage, _super);
    function MemoryStorage(data) {
        var _this = _super.call(this) || this;
        _this.set(data);
        return _this;
    }
    MemoryStorage.prototype.get = function () {
        return __awaiter(this, void 0, void 0, function () {
            var data;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.data()];
                    case 1:
                        data = _a.sent();
                        return [2 /*return*/, {
                                data: data,
                                total: data.length,
                            }];
                }
            });
        });
    };
    MemoryStorage.prototype.set = function (data) {
        if (data instanceof Array) {
            this.data = function () { return data; };
        }
        else if (data instanceof Function) {
            this.data = data;
        }
        return this;
    };
    return MemoryStorage;
}(Storage));

/**
 * Centralized logging lib
 *
 * This class needs some improvements but so far it has been used to have a coherent way to log
 */
var Logger = /** @class */ (function () {
    function Logger() {
    }
    Logger.prototype.format = function (message, type) {
        return "[Grid.js] [" + type.toUpperCase() + "]: " + message;
    };
    Logger.prototype.error = function (message, throwException) {
        if (throwException === void 0) { throwException = false; }
        var msg = this.format(message, 'error');
        if (throwException) {
            throw Error(msg);
        }
        else {
            console.error(msg);
        }
    };
    Logger.prototype.warn = function (message) {
        console.warn(this.format(message, 'warn'));
    };
    Logger.prototype.info = function (message) {
        console.info(this.format(message, 'info'));
    };
    return Logger;
}());
var log = new Logger();

var ServerStorage = /** @class */ (function (_super) {
    __extends(ServerStorage, _super);
    function ServerStorage(options) {
        var _this = _super.call(this) || this;
        _this.options = options;
        return _this;
    }
    ServerStorage.prototype.get = function (options) {
        // this.options is the initial config object
        // options is the runtime config passed by the pipeline (e.g. search component)
        var opts = __assign(__assign({}, this.options), options);
        return fetch(opts.url, opts)
            .then(function (res) {
            if (res.ok) {
                return res.json();
            }
            else {
                log.error("Could not fetch data: " + res.status + " - " + res.statusText, true);
                return null;
            }
        })
            .then(function (res) {
            return {
                data: opts.then(res),
                total: typeof opts.total === 'function' ? opts.total(res) : undefined,
            };
        });
    };
    return ServerStorage;
}(Storage));

var HTMLContent = /** @class */ (function (_super) {
    __extends(HTMLContent, _super);
    function HTMLContent() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    HTMLContent.prototype.render = function () {
        return h(this.props.parentElement, {
            dangerouslySetInnerHTML: { __html: this.props.content },
        });
    };
    HTMLContent.defaultProps = {
        parentElement: 'span',
    };
    return HTMLContent;
}(BaseComponent));

function html(content, parentElement) {
    return h(HTMLContent, { content: content, parentElement: parentElement });
}

var StorageUtils = /** @class */ (function () {
    function StorageUtils() {
    }
    /**
     * Accepts the userConfig dict and tries to guess and return a Storage type
     *
     * @param userConfig
     */
    StorageUtils.createFromUserConfig = function (userConfig) {
        var storage = null;
        // `data` array is provided
        if (userConfig.data) {
            storage = new MemoryStorage(userConfig.data);
        }
        if (userConfig.from) {
            storage = new MemoryStorage(this.tableElementToArray(userConfig.from));
            // remove the source table element from the DOM
            userConfig.from.style.display = 'none';
        }
        if (userConfig.server) {
            storage = new ServerStorage(userConfig.server);
        }
        if (!storage) {
            log.error('Could not determine the storage type', true);
        }
        return storage;
    };
    /**
     * Accepts a HTML table element and converts it into a 2D array of data
     *
     * TODO: This function can be a step in the pipeline: Convert Table -> Load into a memory storage -> ...
     *
     * @param element
     */
    StorageUtils.tableElementToArray = function (element) {
        var arr = [];
        var tbody = element.querySelector('tbody');
        var rows = tbody.querySelectorAll('tr');
        for (var _i = 0, _a = rows; _i < _a.length; _i++) {
            var row = _a[_i];
            var cells = row.querySelectorAll('td');
            var parsedRow = [];
            for (var _b = 0, cells_1 = cells; _b < cells_1.length; _b++) {
                var cell = cells_1[_b];
                // try to capture a TD with single text element first
                if (cell.childNodes.length === 1 &&
                    cell.childNodes[0].nodeType === Node.TEXT_NODE) {
                    parsedRow.push(cell.innerText);
                }
                else {
                    parsedRow.push(html(cell.innerHTML));
                }
            }
            arr.push(parsedRow);
        }
        return arr;
    };
    return StorageUtils;
}());

function className() {
    var args = [];
    for (var _i = 0; _i < arguments.length; _i++) {
        args[_i] = arguments[_i];
    }
    var prefix = 'gridjs';
    return "" + prefix + args.reduce(function (prev, cur) { return prev + "-" + cur; }, '');
}
function classJoin() {
    var classNames = [];
    for (var _i = 0; _i < arguments.length; _i++) {
        classNames[_i] = arguments[_i];
    }
    return classNames
        .filter(function (x) { return x; })
        .reduce(function (className, prev) { return (className || '') + " " + prev; }, '')
        .trim();
}

// container status
var Status;
(function (Status) {
    Status[Status["Init"] = 0] = "Init";
    Status[Status["Loading"] = 1] = "Loading";
    Status[Status["Loaded"] = 2] = "Loaded";
    Status[Status["Rendered"] = 3] = "Rendered";
    Status[Status["Error"] = 4] = "Error";
})(Status || (Status = {}));

var TD = /** @class */ (function (_super) {
    __extends(TD, _super);
    function TD() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    TD.prototype.content = function () {
        if (this.props.column &&
            typeof this.props.column.formatter === 'function') {
            return this.props.column.formatter(this.props.cell.data, this.props.row, this.props.column);
        }
        return this.props.cell.data;
    };
    TD.prototype.render = function () {
        return (h("td", { colSpan: this.props.colSpan, className: classJoin(className('td'), this.props.className) }, this.content()));
    };
    return TD;
}(BaseComponent));

var TR = /** @class */ (function (_super) {
    __extends(TR, _super);
    function TR() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    TR.prototype.getColumn = function (cellIndex) {
        if (this.props.header) {
            return this.props.header.columns[cellIndex];
        }
        return null;
    };
    TR.prototype.render = function () {
        var _this = this;
        if (this.props.children) {
            return h("tr", { className: className('tr') }, this.props.children);
        }
        else {
            return (h("tr", { className: className('tr') }, this.props.row.cells.map(function (cell, i) {
                return (h(TD, { key: cell.id, cell: cell, row: _this.props.row, column: _this.getColumn(i) }));
            })));
        }
    };
    return TR;
}(BaseComponent));

var MessageRow = /** @class */ (function (_super) {
    __extends(MessageRow, _super);
    function MessageRow() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    MessageRow.prototype.render = function () {
        return (h(TR, null,
            h(TD, { colSpan: this.props.colSpan, cell: new Cell(this.props.message), className: classJoin(className('message'), this.props.className ? this.props.className : null) })));
    };
    return MessageRow;
}(BaseComponent));

var TBody = /** @class */ (function (_super) {
    __extends(TBody, _super);
    function TBody() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    TBody.prototype.headerLength = function () {
        if (this.props.header) {
            return this.props.header.columns.length;
        }
        return 0;
    };
    TBody.prototype.render = function () {
        var _this = this;
        return (h("tbody", { className: className('tbody') },
            this.props.data &&
                this.props.data.rows.map(function (row) {
                    return h(TR, { key: row.id, row: row, header: _this.props.header });
                }),
            this.props.status === Status.Loading &&
                (!this.props.data || this.props.data.length === 0) && (h(MessageRow, { message: "Loading...", colSpan: this.headerLength(), className: className('loading') })),
            this.props.status === Status.Loaded &&
                this.props.data &&
                this.props.data.length === 0 && (h(MessageRow, { message: "No matching records found", colSpan: this.headerLength(), className: className('notfound') })),
            this.props.status === Status.Error && (h(MessageRow, { message: "An error happened while fetching the data.", colSpan: this.headerLength(), className: className('error') }))));
    };
    return TBody;
}(BaseComponent));

var ProcessorType;
(function (ProcessorType) {
    ProcessorType[ProcessorType["Initiator"] = 0] = "Initiator";
    ProcessorType[ProcessorType["ServerFilter"] = 1] = "ServerFilter";
    ProcessorType[ProcessorType["ServerSort"] = 2] = "ServerSort";
    ProcessorType[ProcessorType["ServerLimit"] = 3] = "ServerLimit";
    ProcessorType[ProcessorType["Extractor"] = 4] = "Extractor";
    ProcessorType[ProcessorType["Transformer"] = 5] = "Transformer";
    ProcessorType[ProcessorType["Filter"] = 6] = "Filter";
    ProcessorType[ProcessorType["Sort"] = 7] = "Sort";
    ProcessorType[ProcessorType["Limit"] = 8] = "Limit";
})(ProcessorType || (ProcessorType = {}));
var PipelineProcessor = /** @class */ (function () {
    function PipelineProcessor(props) {
        this.propsUpdatedCallback = new Set();
        this.beforeProcessCallback = new Set();
        this.afterProcessCallback = new Set();
        this._props = {};
        this.id = generateID();
        if (props)
            this.setProps(props);
    }
    /**
     * process is used to call beforeProcess and afterProcess callbacks
     * This function is just a wrapper that calls _process()
     *
     * @param args
     */
    PipelineProcessor.prototype.process = function () {
        var args = [];
        for (var _i = 0; _i < arguments.length; _i++) {
            args[_i] = arguments[_i];
        }
        if (this.validateProps instanceof Function) {
            this.validateProps.apply(this, args);
        }
        this.trigger.apply(this, __spreadArrays([this.beforeProcessCallback], args));
        var result = this._process.apply(this, args);
        this.trigger.apply(this, __spreadArrays([this.afterProcessCallback], args));
        return result;
    };
    PipelineProcessor.prototype.trigger = function (fns) {
        var args = [];
        for (var _i = 1; _i < arguments.length; _i++) {
            args[_i - 1] = arguments[_i];
        }
        if (fns) {
            fns.forEach(function (fn) { return fn.apply(void 0, args); });
        }
    };
    PipelineProcessor.prototype.setProps = function (props) {
        Object.assign(this._props, props);
        this.trigger(this.propsUpdatedCallback, this);
        return this;
    };
    Object.defineProperty(PipelineProcessor.prototype, "props", {
        get: function () {
            return this._props;
        },
        enumerable: false,
        configurable: true
    });
    PipelineProcessor.prototype.propsUpdated = function (callback) {
        this.propsUpdatedCallback.add(callback);
        return this;
    };
    PipelineProcessor.prototype.beforeProcess = function (callback) {
        this.beforeProcessCallback.add(callback);
        return this;
    };
    PipelineProcessor.prototype.afterProcess = function (callback) {
        this.afterProcessCallback.add(callback);
        return this;
    };
    return PipelineProcessor;
}());

var NativeSort = /** @class */ (function (_super) {
    __extends(NativeSort, _super);
    function NativeSort() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    NativeSort.prototype.validateProps = function () {
        for (var _i = 0, _a = this.props.columns; _i < _a.length; _i++) {
            var condition = _a[_i];
            if (condition.direction === undefined) {
                condition.direction = 1;
            }
            if (condition.direction !== 1 && condition.direction !== -1) {
                log.error("Invalid sort direction " + condition.direction);
            }
        }
    };
    Object.defineProperty(NativeSort.prototype, "type", {
        get: function () {
            return ProcessorType.Sort;
        },
        enumerable: false,
        configurable: true
    });
    NativeSort.prototype.compare = function (cellA, cellB) {
        if (cellA > cellB) {
            return 1;
        }
        else if (cellA < cellB) {
            return -1;
        }
        return 0;
    };
    NativeSort.prototype.compareWrapper = function (a, b) {
        var cmp = 0;
        for (var _i = 0, _a = this.props.columns; _i < _a.length; _i++) {
            var column = _a[_i];
            if (cmp === 0) {
                var cellA = a.cells[column.index].data;
                var cellB = b.cells[column.index].data;
                if (typeof column.compare === 'function') {
                    cmp |= column.compare(cellA, cellB) * column.direction;
                }
                else {
                    cmp |= this.compare(cellA, cellB) * column.direction;
                }
            }
            else {
                break;
            }
        }
        return cmp;
    };
    NativeSort.prototype._process = function (data) {
        var sorted = __spreadArrays(data.rows);
        sorted.sort(this.compareWrapper.bind(this));
        return new Tabular(sorted);
    };
    return NativeSort;
}(PipelineProcessor));

var EventEmitter = /** @class */ (function () {
    function EventEmitter() {
    }
    // because we are using EventEmitter as a mixin and the
    // constructor won't be called by the applyMixins function
    // see src/base.ts and src/util/applyMixin.ts
    EventEmitter.prototype.init = function (event) {
        if (!this.callbacks) {
            this.callbacks = {};
        }
        if (event && !this.callbacks[event]) {
            this.callbacks[event] = [];
        }
    };
    EventEmitter.prototype.on = function (event, listener) {
        this.init(event);
        this.callbacks[event].push(listener);
        return this;
    };
    EventEmitter.prototype.off = function (event, listener) {
        var eventName = event;
        this.init();
        if (!this.callbacks[eventName] || this.callbacks[eventName].length === 0) {
            // there is no callbacks with this key
            return this;
        }
        this.callbacks[eventName] = this.callbacks[eventName].filter(function (value) { return value != listener; });
        return this;
    };
    EventEmitter.prototype.emit = function (event) {
        var args = [];
        for (var _i = 1; _i < arguments.length; _i++) {
            args[_i - 1] = arguments[_i];
        }
        var eventName = event;
        this.init(eventName);
        if (this.callbacks[eventName].length > 0) {
            this.callbacks[eventName].forEach(function (value) { return value.apply(void 0, args); });
            return true;
        }
        return false;
    };
    return EventEmitter;
}());

var BaseStore = /** @class */ (function (_super) {
    __extends(BaseStore, _super);
    function BaseStore(dispatcher) {
        var _this = _super.call(this) || this;
        _this.dispatcher = dispatcher;
        _this._state = _this.getInitialState();
        dispatcher.register(_this._handle.bind(_this));
        return _this;
    }
    BaseStore.prototype._handle = function (action) {
        this.handle(action.type, action.payload);
    };
    BaseStore.prototype.setState = function (newState) {
        var prevState = this._state;
        this._state = newState;
        this.emit('updated', newState, prevState);
    };
    Object.defineProperty(BaseStore.prototype, "state", {
        get: function () {
            return this._state;
        },
        enumerable: false,
        configurable: true
    });
    return BaseStore;
}(EventEmitter));

var SortStore = /** @class */ (function (_super) {
    __extends(SortStore, _super);
    function SortStore() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    SortStore.prototype.getInitialState = function () {
        return [];
    };
    SortStore.prototype.handle = function (type, payload) {
        if (type === 'SORT_COLUMN') {
            var index = payload.index, direction = payload.direction, multi = payload.multi, compare = payload.compare;
            this.sortColumn(index, direction, multi, compare);
        }
        else if (type === 'SORT_COLUMN_TOGGLE') {
            var index = payload.index, multi = payload.multi, compare = payload.compare;
            this.sortToggle(index, multi, compare);
        }
    };
    SortStore.prototype.sortToggle = function (index, multi, compare) {
        var columns = __spreadArrays(this.state);
        var column = columns.find(function (x) { return x.index === index; });
        if (!column) {
            this.sortColumn(index, 1, multi, compare);
        }
        else {
            this.sortColumn(index, column.direction === 1 ? -1 : 1, multi, compare);
        }
    };
    SortStore.prototype.sortColumn = function (index, direction, multi, compare) {
        var columns = __spreadArrays(this.state);
        var count = columns.length;
        var column = columns.find(function (x) { return x.index === index; });
        var exists = column !== undefined;
        var add = false;
        var reset = false;
        var remove = false;
        var update = false;
        if (!exists) {
            // the column has not been sorted
            if (count === 0) {
                // the first column to be sorted
                add = true;
            }
            else if (count > 0 && !multi) {
                // remove the previously sorted column
                // and sort the current column
                add = true;
                reset = true;
            }
            else if (count > 0 && multi) {
                // multi-sorting
                // sort this column as well
                add = true;
            }
        }
        else {
            // the column has been sorted before
            if (!multi) {
                // single column sorting
                if (count === 1) {
                    update = true;
                }
                else if (count > 1) {
                    // this situation happens when we have already entered
                    // multi-sorting mode but then user tries to sort a single column
                    reset = true;
                    add = true;
                }
            }
            else {
                // multi sorting
                if (column.direction === -1) {
                    // remove the current column from the
                    // sorted columns array
                    remove = true;
                }
                else {
                    update = true;
                }
            }
        }
        if (reset) {
            // resetting the sorted columns
            columns = [];
        }
        if (add) {
            columns.push({
                index: index,
                direction: direction,
                compare: compare,
            });
        }
        else if (update) {
            var index_1 = columns.indexOf(column);
            columns[index_1].direction = direction;
        }
        else if (remove) {
            var index_2 = columns.indexOf(column);
            columns.splice(index_2, 1);
        }
        this.setState(columns);
    };
    return SortStore;
}(BaseStore));

var BaseActions = /** @class */ (function () {
    function BaseActions(dispatcher) {
        this.dispatcher = dispatcher;
    }
    BaseActions.prototype.dispatch = function (type, payload) {
        this.dispatcher.dispatch({
            type: type,
            payload: payload,
        });
    };
    return BaseActions;
}());

var SortActions = /** @class */ (function (_super) {
    __extends(SortActions, _super);
    function SortActions() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    SortActions.prototype.sortColumn = function (index, direction, multi, compare) {
        this.dispatch('SORT_COLUMN', {
            index: index,
            direction: direction,
            multi: multi,
            compare: compare,
        });
    };
    SortActions.prototype.sortToggle = function (index, multi, compare) {
        this.dispatch('SORT_COLUMN_TOGGLE', {
            index: index,
            multi: multi,
            compare: compare,
        });
    };
    return SortActions;
}(BaseActions));

var ServerSort = /** @class */ (function (_super) {
    __extends(ServerSort, _super);
    function ServerSort() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    Object.defineProperty(ServerSort.prototype, "type", {
        get: function () {
            return ProcessorType.ServerSort;
        },
        enumerable: false,
        configurable: true
    });
    ServerSort.prototype._process = function (options) {
        var updates = {};
        if (this.props.url) {
            updates['url'] = this.props.url(options.url, this.props.columns);
        }
        if (this.props.body) {
            updates['body'] = this.props.body(options.body, this.props.columns);
        }
        return __assign(__assign({}, options), updates);
    };
    return ServerSort;
}(PipelineProcessor));

var Sort = /** @class */ (function (_super) {
    __extends(Sort, _super);
    function Sort(props) {
        var _this = _super.call(this, props) || this;
        _this.actions = new SortActions(props.dispatcher);
        _this.store = new SortStore(props.dispatcher);
        if (props.enabled) {
            _this.sortProcessor = _this.getOrCreateSortProcessor();
            _this.store.on('updated', _this.storeUpdated.bind(_this));
            _this.state = { direction: 0 };
        }
        return _this;
    }
    Sort.prototype.componentWillUnmount = function () {
        this.store.off('updated', this.storeUpdated.bind(this));
    };
    Sort.prototype.storeUpdated = function () {
        var _this = this;
        var currentColumn = this.store.state.find(function (x) { return x.index === _this.props.index; });
        if (!currentColumn) {
            this.setState({
                direction: 0,
            });
        }
        else {
            this.setState({
                direction: currentColumn.direction,
            });
        }
    };
    Sort.prototype.getOrCreateSortProcessor = function () {
        var _this = this;
        var processorType = ProcessorType.Sort;
        if (this.props.sort && typeof this.props.sort.server === 'object') {
            processorType = ProcessorType.ServerSort;
        }
        var processors = this.props.pipeline.getStepsByType(processorType);
        // my assumption is that we only have ONE sorting processor in the
        // entire pipeline and that's why I'm displaying a warning here
        if (processors.length > 1) {
            log.warn('There are more than sorting pipeline registered, selecting the first one');
        }
        var processor;
        // A sort process is already registered
        if (processors.length > 0) {
            processor = processors[0];
        }
        else {
            // let's create a new sort processor
            // this event listener is here because
            // we want to subscribe to the sort store only once
            this.store.on('updated', function (sortedColumns) {
                // updates the Sorting processor
                _this.sortProcessor.setProps({
                    columns: sortedColumns,
                });
            });
            if (processorType === ProcessorType.ServerSort) {
                processor = new ServerSort(__assign({ columns: this.store.state }, this.props.sort.server));
            }
            else {
                processor = new NativeSort({
                    columns: this.store.state,
                });
            }
            this.props.pipeline.register(processor);
        }
        return processor;
    };
    Sort.prototype.changeDirection = function (e) {
        e.preventDefault();
        e.stopPropagation();
        // to sort two or more columns at the same time
        this.actions.sortToggle(this.props.index, e.shiftKey === true && this.props.sort.multiColumn, this.props.compare);
    };
    Sort.prototype.render = function () {
        if (!this.props.enabled) {
            return null;
        }
        var direction = this.state.direction;
        var sortClassName = 'neutral';
        if (direction === 1) {
            sortClassName = 'asc';
        }
        else if (direction === -1) {
            sortClassName = 'desc';
        }
        return (h("button", { title: "Sort column " + (direction === 1 ? 'descending' : 'ascending'), className: classJoin(className('sort'), className('sort', sortClassName)), onClick: this.changeDirection.bind(this) }));
    };
    return Sort;
}(BaseComponent));

var TH = /** @class */ (function (_super) {
    __extends(TH, _super);
    function TH() {
        var _this = _super !== null && _super.apply(this, arguments) || this;
        _this.sortRef = p();
        return _this;
    }
    TH.prototype.isSortable = function () {
        return this.props.column.sort.enabled;
    };
    TH.prototype.onClick = function (e) {
        e.stopPropagation();
        if (this.isSortable()) {
            this.sortRef.current.changeDirection(e);
        }
    };
    TH.prototype.render = function () {
        var cls = classJoin(className('th'), this.isSortable() ? className('th', 'sort') : null);
        return (h("th", { className: cls, onClick: this.onClick.bind(this), style: { width: this.props.column.width } },
            this.props.column.name,
            this.isSortable() && (h(Sort, __assign({ ref: this.sortRef, dispatcher: this.props.dispatcher, pipeline: this.props.pipeline, index: this.props.index, sort: this.props.sort }, this.props.column.sort)))));
    };
    return TH;
}(BaseComponent));

var THead = /** @class */ (function (_super) {
    __extends(THead, _super);
    function THead() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    THead.prototype.render = function () {
        var _this = this;
        if (this.props.header) {
            return (h("thead", { key: this.props.header.id, className: className('thead') },
                h(TR, null, this.props.header.columns.map(function (col, i) {
                    return (h(TH, { dispatcher: _this.props.dispatcher, pipeline: _this.props.pipeline, sort: _this.props.sort, column: col, index: i }));
                }))));
        }
        return null;
    };
    return THead;
}(BaseComponent));

var Table = /** @class */ (function (_super) {
    __extends(Table, _super);
    function Table() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    Table.prototype.getStyle = function () {
        var style = {};
        if (this.props.width) {
            style['width'] = this.props.width;
        }
        return style;
    };
    Table.prototype.render = function () {
        return (h("table", { className: className('table'), style: this.getStyle() },
            h(THead, { pipeline: this.props.pipeline, header: this.props.header, dispatcher: this.props.dispatcher, sort: this.props.sort }),
            h(TBody, { data: this.props.data, status: this.props.status, header: this.props.header })));
    };
    return Table;
}(BaseComponent));

function search (keyword, tabular) {
    // escape special regex chars
    keyword = keyword.replace(/[-[\]{}()*+?.,\\^$|#\\s]/g, '\\$&');
    return new Tabular(tabular.rows.filter(function (row) {
        return row.cells.some(function (cell) {
            if (!cell || !cell.data) {
                return false;
            }
            var data = '';
            if (typeof cell.data === 'object') {
                // HTMLContent element
                var element = cell.data;
                if (element.props.content) {
                    // TODO: we should only search in the content of the element. props.content is the entire HTML element
                    data = element.props.content;
                }
            }
            else {
                // primitive types
                data = String(cell.data);
            }
            return new RegExp(keyword, 'gi').test(data);
        });
    }));
}

var GlobalSearchFilter = /** @class */ (function (_super) {
    __extends(GlobalSearchFilter, _super);
    function GlobalSearchFilter() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    Object.defineProperty(GlobalSearchFilter.prototype, "type", {
        get: function () {
            return ProcessorType.Filter;
        },
        enumerable: false,
        configurable: true
    });
    GlobalSearchFilter.prototype._process = function (data) {
        if (this.props.keyword) {
            return search(String(this.props.keyword).trim(), data);
        }
        return data;
    };
    return GlobalSearchFilter;
}(PipelineProcessor));

var SearchStore = /** @class */ (function (_super) {
    __extends(SearchStore, _super);
    function SearchStore() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    SearchStore.prototype.getInitialState = function () {
        return { keyword: null };
    };
    SearchStore.prototype.handle = function (type, payload) {
        if (type === 'SEARCH_KEYWORD') {
            var keyword = payload.keyword;
            this.search(keyword);
        }
    };
    SearchStore.prototype.search = function (keyword) {
        this.setState({ keyword: keyword });
    };
    return SearchStore;
}(BaseStore));

var SearchActions = /** @class */ (function (_super) {
    __extends(SearchActions, _super);
    function SearchActions() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    SearchActions.prototype.search = function (keyword) {
        this.dispatch('SEARCH_KEYWORD', {
            keyword: keyword,
        });
    };
    return SearchActions;
}(BaseActions));

var ServerGlobalSearchFilter = /** @class */ (function (_super) {
    __extends(ServerGlobalSearchFilter, _super);
    function ServerGlobalSearchFilter() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    Object.defineProperty(ServerGlobalSearchFilter.prototype, "type", {
        get: function () {
            return ProcessorType.ServerFilter;
        },
        enumerable: false,
        configurable: true
    });
    ServerGlobalSearchFilter.prototype._process = function (options) {
        if (!this.props.keyword)
            return options;
        var updates = {};
        if (this.props.url) {
            updates['url'] = this.props.url(options.url, this.props.keyword);
        }
        if (this.props.body) {
            updates['body'] = this.props.body(options.body, this.props.keyword);
        }
        return __assign(__assign({}, options), updates);
    };
    return ServerGlobalSearchFilter;
}(PipelineProcessor));

var debounce = function (func, waitFor) {
    var timeout;
    return function () {
        var args = [];
        for (var _i = 0; _i < arguments.length; _i++) {
            args[_i] = arguments[_i];
        }
        return new Promise(function (resolve) {
            if (timeout) {
                clearTimeout(timeout);
            }
            timeout = setTimeout(function () { return resolve(func.apply(void 0, args)); }, waitFor);
        });
    };
};

var Search = /** @class */ (function (_super) {
    __extends(Search, _super);
    function Search(props) {
        var _this = _super.call(this) || this;
        _this.actions = new SearchActions(props.dispatcher);
        _this.store = new SearchStore(props.dispatcher);
        var enabled = props.enabled, keyword = props.keyword;
        if (enabled) {
            // initial search
            _this.actions.search(keyword);
            _this.store.on('updated', _this.storeUpdated.bind(_this));
            var searchProcessor = void 0;
            if (props.server) {
                searchProcessor = new ServerGlobalSearchFilter({
                    keyword: props.keyword,
                    url: props.server.url,
                    body: props.server.body,
                });
            }
            else {
                searchProcessor = new GlobalSearchFilter({
                    keyword: props.keyword,
                });
            }
            _this.searchProcessor = searchProcessor;
            // adds a new processor to the pipeline
            props.pipeline.register(searchProcessor);
        }
        return _this;
    }
    Search.prototype.storeUpdated = function (state) {
        // updates the processor state
        this.searchProcessor.setProps({
            keyword: state.keyword,
        });
    };
    Search.prototype.onChange = function (event) {
        var keyword = event.target.value;
        this.actions.search(keyword);
    };
    Search.prototype.render = function () {
        if (!this.props.enabled)
            return null;
        var onInput = this.onChange.bind(this);
        // add debounce to input only if it's a server-side search
        if (this.searchProcessor instanceof ServerGlobalSearchFilter) {
            onInput = debounce(onInput, this.props.debounceTimeout);
        }
        return (h("div", { className: className('search') },
            h("input", { type: "search", placeholder: this.props.placeholder, onInput: onInput, className: classJoin(className('input'), className('search', 'input')), value: this.store.state.keyword })));
    };
    Search.defaultProps = {
        placeholder: 'Type a keyword...',
        debounceTimeout: 250,
    };
    return Search;
}(BaseComponent));

var t$1,u$1,r$1,i$1=0,o$1=[],c$1=n.__r,f$1=n.diffed,e$1=n.__c,a$1=n.unmount;function v$1(t,r){n.__h&&n.__h(u$1,t,i$1||r),i$1=0;var o=u$1.__H||(u$1.__H={__:[],__h:[]});return t>=o.__.length&&o.__.push({}),o.__[t]}function d$1(n){return i$1=5,h$1(function(){return {current:n}},[])}function h$1(n,u){var r=v$1(t$1++,7);return x(r.__H,u)?(r.__H=u,r.__h=n,r.__=n()):r.__}function _$1(){o$1.some(function(t){if(t.__P)try{t.__H.__h.forEach(g$1),t.__H.__h.forEach(q),t.__H.__h=[];}catch(u){return t.__H.__h=[],n.__e(u,t.__v),!0}}),o$1=[];}function g$1(n){"function"==typeof n.u&&n.u();}function q(n){n.u=n.__();}function x(n,t){return !n||t.some(function(t,u){return t!==n[u]})}n.__r=function(n){c$1&&c$1(n),t$1=0;var r=(u$1=n.__c).__H;r&&(r.__h.forEach(g$1),r.__h.forEach(q),r.__h=[]);},n.diffed=function(t){f$1&&f$1(t);var u=t.__c;u&&u.__H&&u.__H.__h.length&&(1!==o$1.push(u)&&r$1===n.requestAnimationFrame||((r$1=n.requestAnimationFrame)||function(n){var t,u=function(){clearTimeout(r),cancelAnimationFrame(t),setTimeout(n);},r=setTimeout(u,100);"undefined"!=typeof window&&(t=requestAnimationFrame(u));})(_$1));},n.__c=function(t,u){u.some(function(t){try{t.__h.forEach(g$1),t.__h=t.__h.filter(function(n){return !n.__||q(n)});}catch(r){u.some(function(n){n.__h&&(n.__h=[]);}),u=[],n.__e(r,t.__v);}}),e$1&&e$1(t,u);},n.unmount=function(t){a$1&&a$1(t);var u=t.__c;if(u&&u.__H)try{u.__H.__.forEach(g$1);}catch(t){n.__e(t,u.__v);}};

var HeaderContainer = /** @class */ (function (_super) {
    __extends(HeaderContainer, _super);
    function HeaderContainer(props) {
        var _this = _super.call(this, props) || this;
        _this.headerRef = d$1(null);
        _this.state = {
            isActive: true,
        };
        return _this;
    }
    HeaderContainer.prototype.componentDidMount = function () {
        if (this.headerRef.current.children.length === 0) {
            this.setState({
                isActive: false,
            });
        }
    };
    HeaderContainer.prototype.render = function () {
        if (this.state.isActive) {
            return (h("div", { ref: this.headerRef, className: className('head') },
                h(Search, __assign({ dispatcher: this.props.config.dispatcher, pipeline: this.props.config.pipeline }, this.props.config.search))));
        }
        return null;
    };
    return HeaderContainer;
}(BaseComponent));

var PaginationLimit = /** @class */ (function (_super) {
    __extends(PaginationLimit, _super);
    function PaginationLimit() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    PaginationLimit.prototype.validateProps = function () {
        if (isNaN(Number(this.props.limit)) || isNaN(Number(this.props.page))) {
            throw Error('Invalid parameters passed');
        }
    };
    Object.defineProperty(PaginationLimit.prototype, "type", {
        get: function () {
            return ProcessorType.Limit;
        },
        enumerable: false,
        configurable: true
    });
    PaginationLimit.prototype._process = function (data) {
        var page = this.props.page;
        var start = page * this.props.limit;
        var end = (page + 1) * this.props.limit;
        return new Tabular(data.rows.slice(start, end));
    };
    return PaginationLimit;
}(PipelineProcessor));

var ServerPaginationLimit = /** @class */ (function (_super) {
    __extends(ServerPaginationLimit, _super);
    function ServerPaginationLimit() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    Object.defineProperty(ServerPaginationLimit.prototype, "type", {
        get: function () {
            return ProcessorType.ServerLimit;
        },
        enumerable: false,
        configurable: true
    });
    ServerPaginationLimit.prototype._process = function (options) {
        var updates = {};
        if (this.props.url) {
            updates['url'] = this.props.url(options.url, this.props.page, this.props.limit);
        }
        if (this.props.body) {
            updates['body'] = this.props.body(options.body, this.props.page, this.props.limit);
        }
        return __assign(__assign({}, options), updates);
    };
    return ServerPaginationLimit;
}(PipelineProcessor));

var Pagination = /** @class */ (function (_super) {
    __extends(Pagination, _super);
    function Pagination(props) {
        var _this = _super.call(this) || this;
        _this.state = {
            limit: props.limit,
            page: props.page || 0,
            total: 0,
        };
        return _this;
    }
    Pagination.prototype.componentWillMount = function () {
        var _this = this;
        if (this.props.enabled) {
            var processor = void 0;
            if (this.props.server) {
                processor = new ServerPaginationLimit({
                    limit: this.state.limit,
                    page: this.state.page,
                    url: this.props.server.url,
                    body: this.props.server.body,
                });
                this.props.pipeline.afterProcess(function (result) {
                    _this.setTotal(result.length);
                });
            }
            else {
                processor = new PaginationLimit({
                    limit: this.state.limit,
                    page: this.state.page,
                });
                // Pagination (all Limit processors) is the last step in the pipeline
                // and we assume that at this stage, we have the rows that we care about.
                // Let's grab the rows before processing Pagination and set total number of rows
                processor.beforeProcess(function (tabular) { return __awaiter(_this, void 0, void 0, function () {
                    return __generator(this, function (_a) {
                        this.setTotal(tabular.length);
                        return [2 /*return*/];
                    });
                }); });
            }
            this.processor = processor;
            this.props.pipeline.register(processor);
        }
    };
    Pagination.prototype.componentDidMount = function () {
        var _this = this;
        this.props.pipeline.updated(function (processor) {
            // this is to ensure that the current page is set to 0
            // when a processor is updated for some reason
            if (processor !== _this.processor) {
                _this.setPage(0);
            }
        });
    };
    Object.defineProperty(Pagination.prototype, "pages", {
        get: function () {
            return Math.ceil(this.state.total / this.state.limit);
        },
        enumerable: false,
        configurable: true
    });
    Pagination.prototype.setPage = function (page) {
        if (page >= this.pages || page < 0 || page === this.state.page) {
            return null;
        }
        this.setState({
            page: page,
        });
        this.processor.setProps({
            page: page,
        });
    };
    Pagination.prototype.setTotal = function (totalRows) {
        // to set the correct total number of rows
        // when running in-memory operations
        this.setState({
            total: totalRows,
        });
    };
    Pagination.prototype.render = function () {
        var _this = this;
        if (!this.props.enabled)
            return null;
        // how many pagination buttons to render?
        var maxCount = Math.min(this.pages, this.props.buttonsCount);
        var pagePivot = Math.min(this.state.page, Math.floor(maxCount / 2));
        if (this.state.page + Math.floor(maxCount / 2) >= this.pages) {
            pagePivot = maxCount - (this.pages - this.state.page);
        }
        return (h("div", { className: className('pagination') },
            this.props.summary && this.state.total > 0 && (h("div", { className: className('summary'), title: "Page " + (this.state.page + 1) + " of " + this.pages },
                "Showing ",
                h("span", null, this.state.page * this.state.limit + 1),
                " to",
                ' ',
                h("span", null, Math.min((this.state.page + 1) * this.state.limit, this.state.total)),
                ' ',
                "of ",
                h("span", null, this.state.total),
                " results")),
            h("div", { className: className('pages') },
                this.props.prevButton && (h("button", { onClick: this.setPage.bind(this, this.state.page - 1) }, "Previous")),
                this.pages > maxCount && this.state.page - pagePivot > 0 && (h(d, null,
                    h("button", { onClick: this.setPage.bind(this, 0), title: "Page 1" }, "1"),
                    h("button", { className: className('spread') }, "..."))),
                Array.from(Array(maxCount).keys())
                    .map(function (i) { return _this.state.page + (i - pagePivot); })
                    .map(function (i) { return (h("button", { onClick: _this.setPage.bind(_this, i), className: _this.state.page === i ? className('currentPage') : null, title: "Page " + (i + 1) }, i + 1)); }),
                this.pages > maxCount &&
                    this.pages > this.state.page + pagePivot + 1 && (h(d, null,
                    h("button", { className: className('spread') }, "..."),
                    h("button", { onClick: this.setPage.bind(this, this.pages - 1), title: "Page " + this.pages }, this.pages))),
                this.props.nextButton && (h("button", { onClick: this.setPage.bind(this, this.state.page + 1) }, "Next")))));
    };
    Pagination.defaultProps = {
        summary: true,
        nextButton: true,
        prevButton: true,
        buttonsCount: 3,
        limit: 10,
    };
    return Pagination;
}(BaseComponent));

var FooterContainer = /** @class */ (function (_super) {
    __extends(FooterContainer, _super);
    function FooterContainer() {
        var _this = _super.call(this) || this;
        _this.footerRef = d$1(null);
        _this.state = {
            isActive: true,
        };
        return _this;
    }
    FooterContainer.prototype.componentDidMount = function () {
        if (this.footerRef.current.children.length === 0) {
            this.setState({
                isActive: false,
            });
        }
    };
    FooterContainer.prototype.render = function () {
        if (this.state.isActive) {
            return (h("div", { ref: this.footerRef, className: className('footer') },
                h(Pagination, __assign({ storage: this.props.config.storage, pipeline: this.props.config.pipeline }, this.props.config.pagination))));
        }
        return null;
    };
    return FooterContainer;
}(BaseComponent));

var Container = /** @class */ (function (_super) {
    __extends(Container, _super);
    function Container(props) {
        var _this = _super.call(this, props) || this;
        _this.state = {
            status: Status.Loading,
            header: props.header,
            data: null,
        };
        return _this;
    }
    Container.prototype.processPipeline = function () {
        return __awaiter(this, void 0, void 0, function () {
            var _a, _b, e_1;
            return __generator(this, function (_c) {
                switch (_c.label) {
                    case 0:
                        this.setState({
                            status: Status.Loading,
                        });
                        _c.label = 1;
                    case 1:
                        _c.trys.push([1, 3, , 4]);
                        _a = this.setState;
                        _b = {};
                        return [4 /*yield*/, this.props.pipeline.process()];
                    case 2:
                        _a.apply(this, [(_b.data = _c.sent(),
                                _b.status = Status.Loaded,
                                _b)]);
                        return [3 /*break*/, 4];
                    case 3:
                        e_1 = _c.sent();
                        log.error(e_1);
                        this.setState({
                            status: Status.Error,
                        });
                        return [3 /*break*/, 4];
                    case 4: return [2 /*return*/];
                }
            });
        });
    };
    Container.prototype.componentDidMount = function () {
        return __awaiter(this, void 0, void 0, function () {
            var config;
            var _this = this;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        config = this.props.config;
                        return [4 /*yield*/, this.processPipeline()];
                    case 1:
                        _a.sent();
                        if (config.header) {
                            // now that we have the data, let's adjust columns width
                            // note that we only calculate the columns width once
                            this.setState({
                                header: config.header.adjustWidth(config.container, this.state.data, config.autoWidth),
                            });
                        }
                        this.props.pipeline.updated(function () { return __awaiter(_this, void 0, void 0, function () {
                            return __generator(this, function (_a) {
                                switch (_a.label) {
                                    case 0: return [4 /*yield*/, this.processPipeline()];
                                    case 1:
                                        _a.sent();
                                        return [2 /*return*/];
                                }
                            });
                        }); });
                        return [2 /*return*/];
                }
            });
        });
    };
    Container.prototype.render = function () {
        return (h(d, null,
            h("div", { className: classJoin('gridjs', className('container'), this.state.status === Status.Loading ? className('loading') : null), style: { width: this.props.width } },
                this.state.status === Status.Loading && (h("div", { className: className('loading-bar') })),
                h(HeaderContainer, { config: this.props.config }),
                h("div", { className: className('wrapper'), style: { width: this.props.width } },
                    h(Table, { dispatcher: this.props.config.dispatcher, pipeline: this.props.pipeline, data: this.state.data, header: this.state.header, width: this.props.width, status: this.state.status, sort: this.props.config.sort })),
                h(FooterContainer, { config: this.props.config })),
            h("div", { id: "gridjs-temp", className: className('temp') })));
    };
    return Container;
}(BaseComponent));

var Pipeline = /** @class */ (function () {
    function Pipeline(steps) {
        var _this = this;
        // available steps for this pipeline
        this._steps = new Map();
        // used to cache the results of processors using their id field
        this.cache = new Map();
        // keeps the index of the last updated processor in the registered
        // processors list and will be used to invalidate the cache
        // -1 means all new processors should be processed
        this.lastProcessorIndexUpdated = -1;
        this.propsUpdatedCallback = new Set();
        this.afterRegisterCallback = new Set();
        this.updatedCallback = new Set();
        this.afterProcessCallback = new Set();
        if (steps) {
            steps.forEach(function (step) { return _this.register(step); });
        }
    }
    /**
     * Clears the `cache` array
     */
    Pipeline.prototype.clearCache = function () {
        this.cache = new Map();
    };
    /**
     * Registers a new processor
     *
     * @param processor
     * @param priority
     */
    Pipeline.prototype.register = function (processor, priority) {
        if (priority === void 0) { priority = null; }
        if (processor.type === null) {
            throw Error('Processor type is not defined');
        }
        // binding the propsUpdated callback to the Pipeline
        processor.propsUpdated(this.processorPropsUpdated.bind(this));
        this.addProcessorByPriority(processor, priority);
        this.afterRegistered(processor);
    };
    /**
     * Registers a new processor
     *
     * @param processor
     * @param priority
     */
    Pipeline.prototype.addProcessorByPriority = function (processor, priority) {
        var subSteps = this._steps.get(processor.type);
        if (!subSteps) {
            var newSubStep = [];
            this._steps.set(processor.type, newSubStep);
            subSteps = newSubStep;
        }
        if (priority === null || priority < 0) {
            subSteps.push(processor);
        }
        else {
            if (!subSteps[priority]) {
                // slot is empty
                subSteps[priority] = processor;
            }
            else {
                // slot is NOT empty
                var first = subSteps.slice(0, priority - 1);
                var second = subSteps.slice(priority + 1);
                this._steps.set(processor.type, first.concat(processor).concat(second));
            }
        }
    };
    Object.defineProperty(Pipeline.prototype, "steps", {
        /**
         * Flattens the _steps Map and returns a list of steps with their correct priorities
         */
        get: function () {
            var steps = [];
            for (var _i = 0, _a = this.getSortedProcessorTypes(); _i < _a.length; _i++) {
                var type = _a[_i];
                var subSteps = this._steps.get(type);
                if (subSteps && subSteps.length) {
                    steps = steps.concat(subSteps);
                }
            }
            // to remove any undefined elements
            return steps.filter(function (s) { return s; });
        },
        enumerable: false,
        configurable: true
    });
    /**
     * Accepts ProcessType and returns an array of the registered processes
     * with the give type
     *
     * @param type
     */
    Pipeline.prototype.getStepsByType = function (type) {
        return this.steps.filter(function (process) { return process.type === type; });
    };
    /**
     * Returns a list of ProcessorType according to their priority
     */
    Pipeline.prototype.getSortedProcessorTypes = function () {
        return Object.keys(ProcessorType)
            .filter(function (key) { return !isNaN(Number(key)); })
            .map(function (key) { return Number(key); });
    };
    /**
     * Runs all registered processors based on their correct priority
     * and returns the final output after running all steps
     *
     * @param data
     */
    Pipeline.prototype.process = function (data) {
        return __awaiter(this, void 0, void 0, function () {
            var lastProcessorIndexUpdated, steps, prev, _i, steps_1, processor, processorIndex;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        lastProcessorIndexUpdated = this.lastProcessorIndexUpdated;
                        steps = this.steps;
                        prev = data;
                        _i = 0, steps_1 = steps;
                        _a.label = 1;
                    case 1:
                        if (!(_i < steps_1.length)) return [3 /*break*/, 5];
                        processor = steps_1[_i];
                        processorIndex = this.findProcessorIndexByID(processor.id);
                        if (!(processorIndex >= lastProcessorIndexUpdated)) return [3 /*break*/, 3];
                        return [4 /*yield*/, processor.process(prev)];
                    case 2:
                        // we should execute process() here since the last
                        // updated processor was before "processor".
                        // This is to ensure that we always have correct and up to date
                        // data from processors and also to skip them when necessary
                        prev = _a.sent();
                        this.cache.set(processor.id, prev);
                        return [3 /*break*/, 4];
                    case 3:
                        // cached results already exist
                        prev = this.cache.get(processor.id);
                        _a.label = 4;
                    case 4:
                        _i++;
                        return [3 /*break*/, 1];
                    case 5:
                        // means the pipeline is up to date
                        this.lastProcessorIndexUpdated = steps.length;
                        // triggers the afterProcess callbacks with the results
                        this.trigger(this.afterProcessCallback, prev);
                        return [2 /*return*/, prev];
                }
            });
        });
    };
    /**
     * Returns the registered processor's index in _steps array
     *
     * @param processorID
     */
    Pipeline.prototype.findProcessorIndexByID = function (processorID) {
        return this.steps.findIndex(function (p) { return p.id == processorID; });
    };
    /**
     * Sets the last updates processors index locally
     * This is used to invalid or skip a processor in
     * the process() method
     */
    Pipeline.prototype.setLastProcessorIndex = function (processor) {
        var processorIndex = this.findProcessorIndexByID(processor.id);
        if (this.lastProcessorIndexUpdated > processorIndex) {
            this.lastProcessorIndexUpdated = processorIndex;
        }
    };
    Pipeline.prototype.trigger = function (fns) {
        var args = [];
        for (var _i = 1; _i < arguments.length; _i++) {
            args[_i - 1] = arguments[_i];
        }
        if (fns) {
            fns.forEach(function (fn) { return fn.apply(void 0, args); });
        }
    };
    Pipeline.prototype.processorPropsUpdated = function (processor) {
        this.setLastProcessorIndex(processor);
        this.trigger(this.propsUpdatedCallback);
        this.trigger(this.updatedCallback, processor);
    };
    Pipeline.prototype.afterRegistered = function (processor) {
        this.setLastProcessorIndex(processor);
        this.trigger(this.afterRegisterCallback);
        this.trigger(this.updatedCallback, processor);
    };
    /**
     * Triggers the callback when a registered
     * processor's property is updated
     *
     * @param fn
     */
    Pipeline.prototype.propsUpdated = function (fn) {
        this.propsUpdatedCallback.add(fn);
        return this;
    };
    /**
     * Triggers the callback function when a new
     * processor is registered successfully
     *
     * @param fn
     */
    Pipeline.prototype.afterRegister = function (fn) {
        this.afterRegisterCallback.add(fn);
        return this;
    };
    /**
     * Generic updated event. Triggers the callback function when the pipeline
     * is updated, including when a new processor is registered, a processor's props
     * get updated, etc.
     *
     * @param fn
     */
    Pipeline.prototype.updated = function (fn) {
        this.updatedCallback.add(fn);
        return this;
    };
    /**
     * Triggers the callback function when the pipeline
     * is fully processed, before returning the results
     *
     * @param fn
     */
    Pipeline.prototype.afterProcess = function (fn) {
        this.afterProcessCallback.add(fn);
        return this;
    };
    return Pipeline;
}());

var StorageExtractor = /** @class */ (function (_super) {
    __extends(StorageExtractor, _super);
    function StorageExtractor() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    Object.defineProperty(StorageExtractor.prototype, "type", {
        get: function () {
            return ProcessorType.Extractor;
        },
        enumerable: false,
        configurable: true
    });
    StorageExtractor.prototype._process = function (opts) {
        return __awaiter(this, void 0, void 0, function () {
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, this.props.storage.get(opts)];
                    case 1: return [2 /*return*/, _a.sent()];
                }
            });
        });
    };
    return StorageExtractor;
}(PipelineProcessor));

var ArrayToTabularTransformer = /** @class */ (function (_super) {
    __extends(ArrayToTabularTransformer, _super);
    function ArrayToTabularTransformer() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    Object.defineProperty(ArrayToTabularTransformer.prototype, "type", {
        get: function () {
            return ProcessorType.Transformer;
        },
        enumerable: false,
        configurable: true
    });
    ArrayToTabularTransformer.prototype._process = function (storageResponse) {
        return Tabular.fromStorageResponse(storageResponse);
    };
    return ArrayToTabularTransformer;
}(PipelineProcessor));

var ServerInitiator = /** @class */ (function (_super) {
    __extends(ServerInitiator, _super);
    function ServerInitiator() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    Object.defineProperty(ServerInitiator.prototype, "type", {
        get: function () {
            return ProcessorType.Initiator;
        },
        enumerable: false,
        configurable: true
    });
    ServerInitiator.prototype._process = function () {
        return {
            url: this.props.serverStorageOptions.url,
            method: this.props.serverStorageOptions.method,
        };
    };
    return ServerInitiator;
}(PipelineProcessor));

var PipelineUtils = /** @class */ (function () {
    function PipelineUtils() {
    }
    PipelineUtils.createFromConfig = function (config) {
        var pipeline = new Pipeline();
        if (config.storage instanceof ServerStorage) {
            pipeline.register(new ServerInitiator({
                serverStorageOptions: config.server,
            }));
        }
        pipeline.register(new StorageExtractor({ storage: config.storage }));
        pipeline.register(new ArrayToTabularTransformer());
        return pipeline;
    };
    return PipelineUtils;
}());

var _prefix = 'ID_';
/**
 * This class is mostly based on Flux's Dispatcher by Facebook
 * https://github.com/facebook/flux/blob/master/src/Dispatcher.js
 */
var Dispatcher = /** @class */ (function () {
    function Dispatcher() {
        this._callbacks = {};
        this._isDispatching = false;
        this._isHandled = {};
        this._isPending = {};
        this._lastID = 1;
    }
    /**
     * Registers a callback to be invoked with every dispatched payload. Returns
     * a token that can be used with `waitFor()`.
     */
    Dispatcher.prototype.register = function (callback) {
        var id = _prefix + this._lastID++;
        this._callbacks[id] = callback;
        return id;
    };
    /**
     * Removes a callback based on its token.
     */
    Dispatcher.prototype.unregister = function (id) {
        if (!this._callbacks[id]) {
            throw Error("Dispatcher.unregister(...): " + id + " does not map to a registered callback.");
        }
        delete this._callbacks[id];
    };
    /**
     * Waits for the callbacks specified to be invoked before continuing execution
     * of the current callback. This method should only be used by a callback in
     * response to a dispatched payload.
     */
    Dispatcher.prototype.waitFor = function (ids) {
        if (!this._isDispatching) {
            throw Error('Dispatcher.waitFor(...): Must be invoked while dispatching.');
        }
        for (var ii = 0; ii < ids.length; ii++) {
            var id = ids[ii];
            if (this._isPending[id]) {
                if (!this._isHandled[id]) {
                    throw Error("Dispatcher.waitFor(...): Circular dependency detected while ' +\n            'waiting for " + id + ".");
                }
                continue;
            }
            if (!this._callbacks[id]) {
                throw Error("Dispatcher.waitFor(...): " + id + " does not map to a registered callback.");
            }
            this._invokeCallback(id);
        }
    };
    /**
     * Dispatches a payload to all registered callbacks.
     */
    Dispatcher.prototype.dispatch = function (payload) {
        if (this._isDispatching) {
            throw Error('Dispatch.dispatch(...): Cannot dispatch in the middle of a dispatch.');
        }
        this._startDispatching(payload);
        try {
            for (var id in this._callbacks) {
                if (this._isPending[id]) {
                    continue;
                }
                this._invokeCallback(id);
            }
        }
        finally {
            this._stopDispatching();
        }
    };
    /**
     * Is this Dispatcher currently dispatching.
     */
    Dispatcher.prototype.isDispatching = function () {
        return this._isDispatching;
    };
    /**
     * Call the callback stored with the given id. Also do some internal
     * bookkeeping.
     *
     * @internal
     */
    Dispatcher.prototype._invokeCallback = function (id) {
        this._isPending[id] = true;
        this._callbacks[id](this._pendingPayload);
        this._isHandled[id] = true;
    };
    /**
     * Set up bookkeeping needed when dispatching.
     *
     * @internal
     */
    Dispatcher.prototype._startDispatching = function (payload) {
        for (var id in this._callbacks) {
            this._isPending[id] = false;
            this._isHandled[id] = false;
        }
        this._pendingPayload = payload;
        this._isDispatching = true;
    };
    /**
     * Clear bookkeeping used for dispatching.
     *
     * @internal
     */
    Dispatcher.prototype._stopDispatching = function () {
        delete this._pendingPayload;
        this._isDispatching = false;
    };
    return Dispatcher;
}());

var Grid = /** @class */ (function () {
    function Grid(userConfig) {
        this._userConfig = userConfig;
    }
    Grid.prototype.bootstrap = function () {
        this.setConfig(this._userConfig);
        this.setDispatcher(this._userConfig);
        this.setStorage(this._userConfig);
        this.setPipeline(this.config);
    };
    Object.defineProperty(Grid.prototype, "config", {
        get: function () {
            return this._config;
        },
        set: function (config) {
            this._config = config;
        },
        enumerable: false,
        configurable: true
    });
    Grid.prototype.setDispatcher = function (userConfig) {
        this.config.dispatcher = userConfig.dispatcher || new Dispatcher();
        return this;
    };
    Grid.prototype.updateConfig = function (userConfig) {
        this._userConfig = __assign(__assign({}, this._userConfig), userConfig);
        return this;
    };
    Grid.prototype.setConfig = function (userConfig) {
        // sets the current global config
        this.config = __assign(__assign({}, (this.config || {})), Config.fromUserConfig(userConfig));
        return this;
    };
    Grid.prototype.setStorage = function (userConfig) {
        this.config.storage = StorageUtils.createFromUserConfig(userConfig);
        return this;
    };
    Grid.prototype.setPipeline = function (config) {
        this.config.pipeline = PipelineUtils.createFromConfig(config);
        return this;
    };
    Grid.prototype.createElement = function () {
        return h(Container, {
            config: this.config,
            pipeline: this.config.pipeline,
            header: this.config.header,
            width: this.config.width,
        });
    };
    /**
     * Uses the existing container and tries to clear the cache
     * and re-render the existing Grid.js instance again. This is
     * useful when a new config is set/updated.
     *
     */
    Grid.prototype.forceRender = function () {
        if (!this.config.container) {
            log.error('Container is empty', true);
        }
        // re-creates essential components
        this.bootstrap();
        // clear the pipeline cache
        this.config.pipeline.clearCache();
        // TODO: not sure if it's a good idea to render a null element but I couldn't find a better way
        H(null, this.config.container);
        H(this.createElement(), this.config.container);
        return this;
    };
    /**
     * Mounts the Grid.js instance to the container
     * and renders the instance
     *
     * @param container
     */
    Grid.prototype.render = function (container) {
        this.bootstrap();
        if (!container) {
            log.error('Container element cannot be null', true);
        }
        if (container.childNodes.length > 0) {
            log.error("The container element " + container + " is not empty. Make sure the container is empty and call render() again");
            return this;
        }
        this.config.container = container;
        H(this.createElement(), container);
        return this;
    };
    return Grid;
}());

export { Grid, h, html };
//# sourceMappingURL=gridjs.development.es.js.map
