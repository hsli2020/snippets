/*
 * liquidjs@8.0.2, https://github.com/harttle/liquidjs
 * (c) 2016-2019 harttle
 * Released under the MIT License.
 */
(function (global, factory) {
    typeof exports === 'object' && typeof module !== 'undefined' ? module.exports = factory() :
    typeof define === 'function' && define.amd ? define(factory) :
    (global = global || self, global.Liquid = factory());
}(this, function () { 'use strict';

    /*! *****************************************************************************
    Copyright (c) Microsoft Corporation. All rights reserved.
    Licensed under the Apache License, Version 2.0 (the "License"); you may not use
    this file except in compliance with the License. You may obtain a copy of the
    License at http://www.apache.org/licenses/LICENSE-2.0

    THIS CODE IS PROVIDED ON AN *AS IS* BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
    KIND, EITHER EXPRESS OR IMPLIED, INCLUDING WITHOUT LIMITATION ANY IMPLIED
    WARRANTIES OR CONDITIONS OF TITLE, FITNESS FOR A PARTICULAR PURPOSE,
    MERCHANTABLITY OR NON-INFRINGEMENT.

    See the Apache Version 2.0 License for specific language governing permissions
    and limitations under the License.
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
        return new (P || (P = Promise))(function (resolve, reject) {
            function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
            function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
            function step(result) { result.done ? resolve(result.value) : new P(function (resolve) { resolve(result.value); }).then(fulfilled, rejected); }
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

    var toStr = Object.prototype.toString;
    /*
     * Checks if value is classified as a String primitive or object.
     * @param {any} value The value to check.
     * @return {Boolean} Returns true if value is a string, else false.
     */
    function isString(value) {
        return toStr.call(value) === '[object String]';
    }
    function isFunction(value) {
        return typeof value === 'function';
    }
    function stringify(value) {
        if (isNil(value))
            return '';
        value = toLiquid(value);
        return String(value);
    }
    function toLiquid(value) {
        if (isFunction(value.toLiquid))
            return toLiquid(value.toLiquid());
        return value;
    }
    function isNil(value) {
        return value === null || value === undefined;
    }
    function isArray(value) {
        // be compatible with IE 8
        return toStr.call(value) === '[object Array]';
    }
    /*
     * Iterates over own enumerable string keyed properties of an object and invokes iteratee for each property.
     * The iteratee is invoked with three arguments: (value, key, object).
     * Iteratee functions may exit iteration early by explicitly returning false.
     * @param {Object} object The object to iterate over.
     * @param {Function} iteratee The function invoked per iteration.
     * @return {Object} Returns object.
     */
    function forOwn(object, iteratee) {
        object = object || {};
        for (var k in object) {
            if (object.hasOwnProperty(k)) {
                if (iteratee(object[k], k, object) === false)
                    break;
            }
        }
        return object;
    }
    function last(arr) {
        return arr[arr.length - 1];
    }
    /*
     * Checks if value is the language type of Object.
     * (e.g. arrays, functions, objects, regexes, new Number(0), and new String(''))
     * @param {any} value The value to check.
     * @return {Boolean} Returns true if value is an object, else false.
     */
    function isObject(value) {
        var type = typeof value;
        return value !== null && (type === 'object' || type === 'function');
    }
    function range(start, stop, step) {
        if (step === void 0) { step = 1; }
        var arr = [];
        for (var i = start; i < stop; i += step) {
            arr.push(i);
        }
        return arr;
    }
    function padStart(str, length, ch) {
        if (ch === void 0) { ch = ' '; }
        str = String(str);
        var n = length - str.length;
        while (n-- > 0)
            str = ch + str;
        return str;
    }

    var Drop = /** @class */ (function () {
        function Drop() {
        }
        Drop.prototype.valueOf = function () {
            return undefined;
        };
        Drop.prototype.liquidMethodMissing = function (key) {
            return undefined;
        };
        return Drop;
    }());

    var LiquidError = /** @class */ (function (_super) {
        __extends(LiquidError, _super);
        function LiquidError(err, token) {
            var _this = _super.call(this, err.message) || this;
            _this.originalError = err;
            _this.token = token;
            return _this;
        }
        LiquidError.prototype.update = function () {
            var err = this.originalError;
            var context = mkContext(this.token);
            this.message = mkMessage(err.message, this.token);
            this.stack = this.message + '\n' + context +
                '\n' + this.stack + '\nFrom ' + err.stack;
        };
        return LiquidError;
    }(Error));
    var TokenizationError = /** @class */ (function (_super) {
        __extends(TokenizationError, _super);
        function TokenizationError(message, token) {
            var _this = _super.call(this, new Error(message), token) || this;
            _this.name = 'TokenizationError';
            _super.prototype.update.call(_this);
            return _this;
        }
        return TokenizationError;
    }(LiquidError));
    var ParseError = /** @class */ (function (_super) {
        __extends(ParseError, _super);
        function ParseError(err, token) {
            var _this = _super.call(this, err, token) || this;
            _this.name = 'ParseError';
            _this.message = err.message;
            _super.prototype.update.call(_this);
            return _this;
        }
        return ParseError;
    }(LiquidError));
    var RenderError = /** @class */ (function (_super) {
        __extends(RenderError, _super);
        function RenderError(err, tpl) {
            var _this = _super.call(this, err, tpl.token) || this;
            _this.name = 'RenderError';
            _this.message = err.message;
            _super.prototype.update.call(_this);
            return _this;
        }
        return RenderError;
    }(LiquidError));
    var RenderBreakError = /** @class */ (function (_super) {
        __extends(RenderBreakError, _super);
        function RenderBreakError(message) {
            var _this = _super.call(this, message) || this;
            _this.resolvedHTML = '';
            _this.name = 'RenderBreakError';
            _this.message = message + '';
            return _this;
        }
        return RenderBreakError;
    }(Error));
    var AssertionError = /** @class */ (function (_super) {
        __extends(AssertionError, _super);
        function AssertionError(message) {
            var _this = _super.call(this, message) || this;
            _this.name = 'AssertionError';
            _this.message = message + '';
            return _this;
        }
        return AssertionError;
    }(Error));
    function mkContext(token) {
        var lines = token.input.split('\n');
        var begin = Math.max(token.line - 2, 1);
        var end = Math.min(token.line + 3, lines.length);
        var context = range(begin, end + 1)
            .map(function (lineNumber) {
            var indicator = (lineNumber === token.line) ? '>> ' : '   ';
            var num = padStart(String(lineNumber), String(end).length);
            var text = lines[lineNumber - 1];
            return "" + indicator + num + "| " + text;
        })
            .join('\n');
        return context;
    }
    function mkMessage(msg, token) {
        if (token.file)
            msg += ", file:" + token.file;
        msg += ", line:" + token.line + ", col:" + token.col;
        return msg;
    }

    function assert (predicate, message) {
        if (!predicate) {
            message = message || "expect " + predicate + " to be true";
            throw new AssertionError(message);
        }
    }

    var defaultOptions = {
        root: ['.'],
        cache: false,
        extname: '',
        dynamicPartials: true,
        trimTagRight: false,
        trimTagLeft: false,
        trimOutputRight: false,
        trimOutputLeft: false,
        greedy: true,
        tagDelimiterLeft: '{%',
        tagDelimiterRight: '%}',
        outputDelimiterLeft: '{{',
        outputDelimiterRight: '}}',
        strictFilters: false,
        strictVariables: false
    };
    function normalize(options) {
        options = options || {};
        if (options.hasOwnProperty('root')) {
            options.root = normalizeStringArray(options.root);
        }
        return options;
    }
    function applyDefault(options) {
        return __assign({}, defaultOptions, options);
    }
    function normalizeStringArray(value) {
        if (isArray(value))
            return value;
        if (isString(value))
            return [value];
        return [];
    }

    var BlockMode;
    (function (BlockMode) {
        /* store rendered html into blocks */
        BlockMode[BlockMode["OUTPUT"] = 0] = "OUTPUT";
        /* output rendered html directly */
        BlockMode[BlockMode["STORE"] = 1] = "STORE";
    })(BlockMode || (BlockMode = {}));
    var BlockMode$1 = BlockMode;

    var Context = /** @class */ (function () {
        function Context(ctx, opts) {
            if (ctx === void 0) { ctx = {}; }
            this.scopes = [{}];
            this.blocks = {};
            this.groups = {};
            this.blockMode = BlockMode$1.OUTPUT;
            this.opts = applyDefault(opts);
            this.environments = ctx;
        }
        Context.prototype.getAll = function () {
            return [this.environments].concat(this.scopes).reduce(function (ctx, val) { return __assign(ctx, val); }, {});
        };
        Context.prototype.get = function (path) {
            return __awaiter(this, void 0, void 0, function () {
                var paths, ctx, _i, paths_1, path_1;
                return __generator(this, function (_a) {
                    switch (_a.label) {
                        case 0: return [4 /*yield*/, this.parseProp(path)];
                        case 1:
                            paths = _a.sent();
                            ctx = this.findScope(paths[0]) || this.environments;
                            for (_i = 0, paths_1 = paths; _i < paths_1.length; _i++) {
                                path_1 = paths_1[_i];
                                ctx = this.readProperty(ctx, path_1);
                                if (isNil(ctx) && this.opts.strictVariables) {
                                    throw new TypeError("undefined variable: " + path_1);
                                }
                            }
                            return [2 /*return*/, ctx];
                    }
                });
            });
        };
        Context.prototype.push = function (ctx) {
            return this.scopes.push(ctx);
        };
        Context.prototype.pop = function (ctx) {
            if (!arguments.length) {
                return this.scopes.pop();
            }
            var i = this.scopes.findIndex(function (scope) { return scope === ctx; });
            if (i === -1) {
                throw new TypeError('scope not found, cannot pop');
            }
            return this.scopes.splice(i, 1)[0];
        };
        Context.prototype.findScope = function (key) {
            for (var i = this.scopes.length - 1; i >= 0; i--) {
                var candidate = this.scopes[i];
                if (key in candidate) {
                    return candidate;
                }
            }
            return null;
        };
        Context.prototype.readProperty = function (obj, key) {
            if (isNil(obj))
                return obj;
            obj = toLiquid(obj);
            if (obj instanceof Drop) {
                if (isFunction(obj[key]))
                    return obj[key]();
                if (obj.hasOwnProperty(key))
                    return obj[key];
                return obj.liquidMethodMissing(key);
            }
            return key === 'size' ? readSize(obj) : obj[key];
        };
        /*
         * Parse property access sequence from access string
         * @example
         * accessSeq("foo.bar")         // ['foo', 'bar']
         * accessSeq("foo['bar']")      // ['foo', 'bar']
         * accessSeq("foo['b]r']")      // ['foo', 'b]r']
         * accessSeq("foo[bar.coo]")    // ['foo', 'bar'], for bar.coo == 'bar'
         */
        Context.prototype.parseProp = function (str) {
            return __awaiter(this, void 0, void 0, function () {
                function push() {
                    if (name.length)
                        seq.push(name);
                    name = '';
                }
                var seq, name, j, i, _a, delemiter, _b;
                return __generator(this, function (_c) {
                    switch (_c.label) {
                        case 0:
                            str = String(str);
                            seq = [];
                            name = '';
                            i = 0;
                            _c.label = 1;
                        case 1:
                            if (!(i < str.length)) return [3 /*break*/, 10];
                            _a = str[i];
                            switch (_a) {
                                case '[': return [3 /*break*/, 2];
                                case '.': return [3 /*break*/, 7];
                            }
                            return [3 /*break*/, 8];
                        case 2:
                            push();
                            delemiter = str[i + 1];
                            if (!/['"]/.test(delemiter)) return [3 /*break*/, 3];
                            j = str.indexOf(delemiter, i + 2);
                            assert(j !== -1, "unbalanced " + delemiter + ": " + str);
                            name = str.slice(i + 2, j);
                            push();
                            i = j + 2;
                            return [3 /*break*/, 6];
                        case 3:
                            j = matchRightBracket(str, i + 1);
                            assert(j !== -1, "unbalanced []: " + str);
                            name = str.slice(i + 1, j);
                            if (!!/^[+-]?\d+$/.test(name)) return [3 /*break*/, 5];
                            _b = String;
                            return [4 /*yield*/, this.get(name)];
                        case 4:
                            name = _b.apply(void 0, [_c.sent()]);
                            _c.label = 5;
                        case 5:
                            push();
                            i = j + 1;
                            _c.label = 6;
                        case 6: return [3 /*break*/, 9];
                        case 7:
                            push();
                            i++;
                            return [3 /*break*/, 9];
                        case 8:
                            name += str[i++];
                            _c.label = 9;
                        case 9: return [3 /*break*/, 1];
                        case 10:
                            push();
                            if (!seq.length) {
                                throw new TypeError("invalid path:\"" + str + "\"");
                            }
                            return [2 /*return*/, seq];
                    }
                });
            });
        };
        return Context;
    }());
    function readSize(obj) {
        if (!isNil(obj['size']))
            return obj['size'];
        if (isArray(obj) || isString(obj))
            return obj.length;
        return obj['size'];
    }
    function matchRightBracket(str, begin) {
        var stack = 1; // count of '[' - count of ']'
        for (var i = begin; i < str.length; i++) {
            if (str[i] === '[') {
                stack++;
            }
            if (str[i] === ']') {
                stack--;
                if (stack === 0) {
                    return i;
                }
            }
        }
        return -1;
    }



    var Types = /*#__PURE__*/Object.freeze({
        ParseError: ParseError,
        TokenizationError: TokenizationError,
        RenderBreakError: RenderBreakError,
        AssertionError: AssertionError,
        Drop: Drop
    });

    function domResolve(root, path) {
        var base = document.createElement('base');
        base.href = root;
        var head = document.getElementsByTagName('head')[0];
        head.insertBefore(base, head.firstChild);
        var a = document.createElement('a');
        a.href = path;
        var resolved = a.href;
        head.removeChild(base);
        return resolved;
    }
    function resolve(root, filepath, ext) {
        if (root.length && last(root) !== '/')
            root += '/';
        var url = domResolve(root, filepath);
        return url.replace(/^(\w+:\/\/[^/]+)(\/[^?]+)/, function (str, origin, path) {
            var last$$1 = path.split('/').pop();
            if (/\.\w+$/.test(last$$1))
                return str;
            return origin + path + ext;
        });
    }
    function readFile(url) {
        return __awaiter(this, void 0, void 0, function () {
            return __generator(this, function (_a) {
                return [2 /*return*/, new Promise(function (resolve, reject) {
                        var xhr = new XMLHttpRequest();
                        xhr.onload = function () {
                            if (xhr.status >= 200 && xhr.status < 300) {
                                resolve(xhr.responseText);
                            }
                            else {
                                reject(new Error(xhr.statusText));
                            }
                        };
                        xhr.onerror = function () {
                            reject(new Error('An error occurred whilst receiving the response.'));
                        };
                        xhr.open('GET', url);
                        xhr.send();
                    })];
            });
        });
    }
    function exists() {
        return __awaiter(this, void 0, void 0, function () {
            return __generator(this, function (_a) {
                return [2 /*return*/, true];
            });
        });
    }
    var fs = { readFile: readFile, resolve: resolve, exists: exists };

    var Token = /** @class */ (function () {
        function Token(raw, input, line, col, file) {
            this.trimLeft = false;
            this.trimRight = false;
            this.type = 'notset';
            this.col = col;
            this.line = line;
            this.raw = raw;
            this.value = raw;
            this.input = input;
            this.file = file;
        }
        return Token;
    }());

    var DelimitedToken = /** @class */ (function (_super) {
        __extends(DelimitedToken, _super);
        function DelimitedToken(raw, value, input, line, pos, trimLeft, trimRight, file) {
            var _this = _super.call(this, raw, input, line, pos, file) || this;
            var tl = value[0] === '-';
            var tr = last(value) === '-';
            _this.value = value
                .slice(tl ? 1 : 0, tr ? -1 : value.length)
                .trim();
            _this.trimLeft = tl || trimLeft;
            _this.trimRight = tr || trimRight;
            return _this;
        }
        return DelimitedToken;
    }(Token));

    // quote related
    var singleQuoted = /'[^']*'/;
    var doubleQuoted = /"[^"]*"/;
    var quoted = new RegExp(singleQuoted.source + "|" + doubleQuoted.source);
    var quoteBalanced = new RegExp("(?:" + quoted.source + "|[^'\"])*");
    // basic types
    var number = /[+-]?(?:\d+\.?\d*|\.?\d+)/;
    var bool = /true|false/;
    // property access
    var identifier = /[\w-]+[?]?/;
    var subscript = new RegExp("\\[(?:" + quoted.source + "|[\\w-\\.]+)\\]");
    var literal = new RegExp("(?:" + quoted.source + "|" + bool.source + "|" + number.source + ")");
    var variable = new RegExp(identifier.source + "(?:\\." + identifier.source + "|" + subscript.source + ")*");
    // range related
    var rangeLimit = new RegExp("(?:" + variable.source + "|" + number.source + ")");
    var range$1 = new RegExp("\\(" + rangeLimit.source + "\\.\\." + rangeLimit.source + "\\)");
    var rangeCapture = new RegExp("\\((" + rangeLimit.source + ")\\.\\.(" + rangeLimit.source + ")\\)");
    var value = new RegExp("(?:" + variable.source + "|" + literal.source + "|" + range$1.source + ")");
    // hash related
    var hash = new RegExp("(?:" + identifier.source + ")\\s*:\\s*(?:" + value.source + ")");
    var hashCapture = new RegExp("(" + identifier.source + ")\\s*:\\s*(" + value.source + ")", 'g');
    // full match
    var tagLine = new RegExp("^\\s*(" + identifier.source + ")\\s*([\\s\\S]*?)\\s*$");
    var quotedLine = new RegExp("^" + quoted.source + "$");
    var rangeLine = new RegExp("^" + rangeCapture.source + "$");
    var operators = [
        /\s+or\s+/,
        /\s+and\s+/,
        /==|!=|<=|>=|<|>|\s+contains\s+/
    ];

    var TagToken = /** @class */ (function (_super) {
        __extends(TagToken, _super);
        function TagToken(raw, value$$1, input, line, pos, options, file) {
            var _this = _super.call(this, raw, value$$1, input, line, pos, options.trimTagLeft, options.trimTagRight, file) || this;
            _this.type = 'tag';
            var match = _this.value.match(tagLine);
            if (!match) {
                throw new TokenizationError("illegal tag syntax", _this);
            }
            _this.name = match[1];
            _this.args = match[2];
            return _this;
        }
        TagToken.is = function (token) {
            return token.type === 'tag';
        };
        return TagToken;
    }(DelimitedToken));

    var HTMLToken = /** @class */ (function (_super) {
        __extends(HTMLToken, _super);
        function HTMLToken(str, input, line, col, file) {
            var _this = _super.call(this, str, input, line, col, file) || this;
            _this.type = 'html';
            _this.value = str;
            return _this;
        }
        HTMLToken.is = function (token) {
            return token.type === 'html';
        };
        return HTMLToken;
    }(Token));

    function whiteSpaceCtrl(tokens, options) {
        options = __assign({ greedy: true }, options);
        var inRaw = false;
        for (var i = 0; i < tokens.length; i++) {
            var token = tokens[i];
            if (!inRaw && token.trimLeft) {
                trimLeft(tokens[i - 1], options.greedy);
            }
            if (TagToken.is(token)) {
                if (token.name === 'raw')
                    inRaw = true;
                else if (token.name === 'endraw')
                    inRaw = false;
            }
            if (!inRaw && token.trimRight) {
                trimRight(tokens[i + 1], options.greedy);
            }
        }
    }
    function trimLeft(token, greedy) {
        if (!token || !HTMLToken.is(token))
            return;
        var rLeft = greedy ? /\s+$/g : /[\t\r ]*$/g;
        token.value = token.value.replace(rLeft, '');
    }
    function trimRight(token, greedy) {
        if (!token || !HTMLToken.is(token))
            return;
        var rRight = greedy ? /^\s+/g : /^[\t\r ]*\n?/g;
        token.value = token.value.replace(rRight, '');
    }

    var OutputToken = /** @class */ (function (_super) {
        __extends(OutputToken, _super);
        function OutputToken(raw, value, input, line, pos, options, file) {
            var _this = _super.call(this, raw, value, input, line, pos, options.trimOutputLeft, options.trimOutputRight, file) || this;
            _this.type = 'output';
            return _this;
        }
        OutputToken.is = function (token) {
            return token.type === 'output';
        };
        return OutputToken;
    }(DelimitedToken));

    var ParseState;
    (function (ParseState) {
        ParseState[ParseState["HTML"] = 0] = "HTML";
        ParseState[ParseState["OUTPUT"] = 1] = "OUTPUT";
        ParseState[ParseState["TAG"] = 2] = "TAG";
    })(ParseState || (ParseState = {}));
    var Tokenizer = /** @class */ (function () {
        function Tokenizer(options) {
            this.options = applyDefault(options);
        }
        Tokenizer.prototype.tokenize = function (input, file) {
            var tokens = [];
            var _a = this.options, tagDelimiterLeft = _a.tagDelimiterLeft, tagDelimiterRight = _a.tagDelimiterRight, outputDelimiterLeft = _a.outputDelimiterLeft, outputDelimiterRight = _a.outputDelimiterRight;
            var p = 0;
            var curLine = 1;
            var state = ParseState.HTML;
            var buffer = '';
            var lineBegin = 0;
            var line = 1;
            var col = 1;
            while (p < input.length) {
                if (input[p] === '\n') {
                    curLine++;
                    lineBegin = p + 1;
                }
                if (state === ParseState.HTML) {
                    if (input.substr(p, outputDelimiterLeft.length) === outputDelimiterLeft) {
                        if (buffer)
                            tokens.push(new HTMLToken(buffer, input, line, col, file));
                        buffer = outputDelimiterLeft;
                        line = curLine;
                        col = p - lineBegin + 1;
                        p += outputDelimiterLeft.length;
                        state = ParseState.OUTPUT;
                        continue;
                    }
                    else if (input.substr(p, tagDelimiterLeft.length) === tagDelimiterLeft) {
                        if (buffer)
                            tokens.push(new HTMLToken(buffer, input, line, col, file));
                        buffer = tagDelimiterLeft;
                        line = curLine;
                        col = p - lineBegin + 1;
                        p += tagDelimiterLeft.length;
                        state = ParseState.TAG;
                        continue;
                    }
                }
                else if (state === ParseState.OUTPUT &&
                    input.substr(p, outputDelimiterRight.length) === outputDelimiterRight) {
                    buffer += outputDelimiterRight;
                    tokens.push(new OutputToken(buffer, buffer.slice(outputDelimiterLeft.length, -outputDelimiterRight.length), input, line, col, this.options, file));
                    p += outputDelimiterRight.length;
                    buffer = '';
                    line = curLine;
                    col = p - lineBegin + 1;
                    state = ParseState.HTML;
                    continue;
                }
                else if (input.substr(p, tagDelimiterRight.length) === tagDelimiterRight) {
                    buffer += tagDelimiterRight;
                    tokens.push(new TagToken(buffer, buffer.slice(tagDelimiterLeft.length, -tagDelimiterRight.length), input, line, col, this.options, file));
                    p += tagDelimiterRight.length;
                    buffer = '';
                    line = curLine;
                    col = p - lineBegin + 1;
                    state = ParseState.HTML;
                    continue;
                }
                buffer += input[p++];
            }
            if (state !== ParseState.HTML) {
                var t = state === ParseState.OUTPUT ? 'output' : 'tag';
                var str = buffer.length > 16 ? buffer.slice(0, 13) + '...' : buffer;
                throw new TokenizationError(t + " \"" + str + "\" not closed", new Token(buffer, input, line, col, file));
            }
            if (buffer)
                tokens.push(new HTMLToken(buffer, input, line, col, file));
            whiteSpaceCtrl(tokens, this.options);
            return tokens;
        };
        return Tokenizer;
    }());

    var Render = /** @class */ (function () {
        function Render() {
        }
        Render.prototype.renderTemplates = function (templates, ctx) {
            return __awaiter(this, void 0, void 0, function () {
                var html, _i, templates_1, tpl, _a, e_1;
                return __generator(this, function (_b) {
                    switch (_b.label) {
                        case 0:
                            assert(ctx, 'unable to evalTemplates: context undefined');
                            html = '';
                            _i = 0, templates_1 = templates;
                            _b.label = 1;
                        case 1:
                            if (!(_i < templates_1.length)) return [3 /*break*/, 6];
                            tpl = templates_1[_i];
                            _b.label = 2;
                        case 2:
                            _b.trys.push([2, 4, , 5]);
                            _a = html;
                            return [4 /*yield*/, tpl.render(ctx)];
                        case 3:
                            html = _a + _b.sent();
                            return [3 /*break*/, 5];
                        case 4:
                            e_1 = _b.sent();
                            if (e_1.name === 'RenderBreakError') {
                                e_1.resolvedHTML = html;
                                throw e_1;
                            }
                            throw e_1.name === 'RenderError' ? e_1 : new RenderError(e_1, tpl);
                        case 5:
                            _i++;
                            return [3 /*break*/, 1];
                        case 6: return [2 /*return*/, html];
                    }
                });
            });
        };
        return Render;
    }());

    function isComparable(arg) {
        return arg && isFunction(arg.equals);
    }

    var EmptyDrop = /** @class */ (function (_super) {
        __extends(EmptyDrop, _super);
        function EmptyDrop() {
            return _super !== null && _super.apply(this, arguments) || this;
        }
        EmptyDrop.prototype.equals = function (value) {
            if (isString(value) || isArray(value))
                return value.length === 0;
            if (isObject(value))
                return Object.keys(value).length === 0;
            return false;
        };
        EmptyDrop.prototype.gt = function () {
            return false;
        };
        EmptyDrop.prototype.geq = function () {
            return false;
        };
        EmptyDrop.prototype.lt = function () {
            return false;
        };
        EmptyDrop.prototype.leq = function () {
            return false;
        };
        EmptyDrop.prototype.valueOf = function () {
            return '';
        };
        return EmptyDrop;
    }(Drop));

    var BlankDrop = /** @class */ (function (_super) {
        __extends(BlankDrop, _super);
        function BlankDrop() {
            return _super !== null && _super.apply(this, arguments) || this;
        }
        BlankDrop.prototype.equals = function (value) {
            if (value === false)
                return true;
            if (isNil(value instanceof Drop ? value.valueOf() : value))
                return true;
            if (isString(value))
                return /^\s*$/.test(value);
            return _super.prototype.equals.call(this, value);
        };
        return BlankDrop;
    }(EmptyDrop));

    var NullDrop = /** @class */ (function (_super) {
        __extends(NullDrop, _super);
        function NullDrop() {
            return _super !== null && _super.apply(this, arguments) || this;
        }
        NullDrop.prototype.equals = function (value) {
            return isNil(value instanceof Drop ? value.valueOf() : value) || value instanceof BlankDrop;
        };
        NullDrop.prototype.gt = function () {
            return false;
        };
        NullDrop.prototype.geq = function () {
            return false;
        };
        NullDrop.prototype.lt = function () {
            return false;
        };
        NullDrop.prototype.leq = function () {
            return false;
        };
        NullDrop.prototype.valueOf = function () {
            return null;
        };
        return NullDrop;
    }(Drop));

    var binaryOperators = {
        '==': function (l, r) {
            if (isComparable(l))
                return l.equals(r);
            if (isComparable(r))
                return r.equals(l);
            return l === r;
        },
        '!=': function (l, r) {
            if (isComparable(l))
                return !l.equals(r);
            if (isComparable(r))
                return !r.equals(l);
            return l !== r;
        },
        '>': function (l, r) {
            if (isComparable(l))
                return l.gt(r);
            if (isComparable(r))
                return r.lt(l);
            return l > r;
        },
        '<': function (l, r) {
            if (isComparable(l))
                return l.lt(r);
            if (isComparable(r))
                return r.gt(l);
            return l < r;
        },
        '>=': function (l, r) {
            if (isComparable(l))
                return l.geq(r);
            if (isComparable(r))
                return r.leq(l);
            return l >= r;
        },
        '<=': function (l, r) {
            if (isComparable(l))
                return l.leq(r);
            if (isComparable(r))
                return r.geq(l);
            return l <= r;
        },
        'contains': function (l, r) {
            return l && isFunction(l.indexOf) ? l.indexOf(r) > -1 : false;
        },
        'and': function (l, r) { return isTruthy(l) && isTruthy(r); },
        'or': function (l, r) { return isTruthy(l) || isTruthy(r); }
    };
    function parseExp(exp, ctx) {
        return __awaiter(this, void 0, void 0, function () {
            var operatorREs, match, i, operatorRE, expRE, l, op, r, low, high;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0:
                        assert(ctx, 'unable to parseExp: scope undefined');
                        operatorREs = operators;
                        i = 0;
                        _a.label = 1;
                    case 1:
                        if (!(i < operatorREs.length)) return [3 /*break*/, 5];
                        operatorRE = operatorREs[i];
                        expRE = new RegExp("^(" + quoteBalanced.source + ")(" + operatorRE.source + ")(" + quoteBalanced.source + ")$");
                        if (!(match = exp.match(expRE))) return [3 /*break*/, 4];
                        return [4 /*yield*/, parseExp(match[1], ctx)];
                    case 2:
                        l = _a.sent();
                        op = binaryOperators[match[2].trim()];
                        return [4 /*yield*/, parseExp(match[3], ctx)];
                    case 3:
                        r = _a.sent();
                        return [2 /*return*/, op(l, r)];
                    case 4:
                        i++;
                        return [3 /*break*/, 1];
                    case 5:
                        if (!(match = exp.match(rangeLine))) return [3 /*break*/, 8];
                        return [4 /*yield*/, evalValue(match[1], ctx)];
                    case 6:
                        low = _a.sent();
                        return [4 /*yield*/, evalValue(match[2], ctx)];
                    case 7:
                        high = _a.sent();
                        return [2 /*return*/, range(+low, +high + 1)];
                    case 8: return [2 /*return*/, parseValue(exp, ctx)];
                }
            });
        });
    }
    function evalExp(str, ctx) {
        return __awaiter(this, void 0, void 0, function () {
            var value$$1;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, parseExp(str, ctx)];
                    case 1:
                        value$$1 = _a.sent();
                        return [2 /*return*/, value$$1 instanceof Drop ? value$$1.valueOf() : value$$1];
                }
            });
        });
    }
    function parseValue(str, ctx) {
        return __awaiter(this, void 0, void 0, function () {
            return __generator(this, function (_a) {
                if (!str)
                    return [2 /*return*/, null];
                str = str.trim();
                if (str === 'true')
                    return [2 /*return*/, true];
                if (str === 'false')
                    return [2 /*return*/, false];
                if (str === 'nil' || str === 'null')
                    return [2 /*return*/, new NullDrop()];
                if (str === 'empty')
                    return [2 /*return*/, new EmptyDrop()];
                if (str === 'blank')
                    return [2 /*return*/, new BlankDrop()];
                if (!isNaN(Number(str)))
                    return [2 /*return*/, Number(str)];
                if ((str[0] === '"' || str[0] === "'") && str[0] === last(str))
                    return [2 /*return*/, str.slice(1, -1)];
                return [2 /*return*/, ctx.get(str)];
            });
        });
    }
    function evalValue(str, ctx) {
        return __awaiter(this, void 0, void 0, function () {
            var value$$1;
            return __generator(this, function (_a) {
                switch (_a.label) {
                    case 0: return [4 /*yield*/, parseValue(str, ctx)];
                    case 1:
                        value$$1 = _a.sent();
                        return [2 /*return*/, value$$1 instanceof Drop ? value$$1.valueOf() : value$$1];
                }
            });
        });
    }
    function isTruthy(val) {
        return !isFalsy(val);
    }
    function isFalsy(val) {
        return val === false || undefined === val || val === null;
    }

    /**
     * Key-Value Pairs Representing Tag Arguments
     * Example:
     *    For the markup `{% include 'head.html' foo='bar' %}`,
     *    hash['foo'] === 'bar'
     */
    var Hash = /** @class */ (function () {
        function Hash() {
        }
        Hash.create = function (markup, ctx) {
            return __awaiter(this, void 0, void 0, function () {
                var instance, match, k, v, _a, _b;
                return __generator(this, function (_c) {
                    switch (_c.label) {
                        case 0:
                            instance = new Hash();
                            hashCapture.lastIndex = 0;
                            _c.label = 1;
                        case 1:
                            if (!(match = hashCapture.exec(markup))) return [3 /*break*/, 3];
                            k = match[1];
                            v = match[2];
                            _a = instance;
                            _b = k;
                            return [4 /*yield*/, evalValue(v, ctx)];
                        case 2:
                            _a[_b] = _c.sent();
                            return [3 /*break*/, 1];
                        case 3: return [2 /*return*/, instance];
                    }
                });
            });
        };
        return Hash;
    }());

    var Template = /** @class */ (function () {
        function Template(token) {
            this.token = token;
        }
        return Template;
    }());

    var Tag = /** @class */ (function (_super) {
        __extends(Tag, _super);
        function Tag(token, tokens, liquid) {
            var _this = _super.call(this, token) || this;
            _this.name = token.name;
            var impl = Tag.impls[token.name];
            assert(impl, "tag " + token.name + " not found");
            _this.impl = Object.create(impl);
            _this.impl.liquid = liquid;
            if (_this.impl.parse) {
                _this.impl.parse(token, tokens);
            }
            return _this;
        }
        Tag.prototype.render = function (ctx) {
            return __awaiter(this, void 0, void 0, function () {
                var hash, impl, _a, _b;
                return __generator(this, function (_c) {
                    switch (_c.label) {
                        case 0: return [4 /*yield*/, Hash.create(this.token.args, ctx)];
                        case 1:
                            hash = _c.sent();
                            impl = this.impl;
                            if (!isFunction(impl.render)) return [3 /*break*/, 3];
                            _b = stringify;
                            return [4 /*yield*/, impl.render(ctx, hash)];
                        case 2:
                            _a = _b.apply(void 0, [_c.sent()]);
                            return [3 /*break*/, 4];
                        case 3:
                            _a = '';
                            _c.label = 4;
                        case 4: return [2 /*return*/, _a];
                    }
                });
            });
        };
        Tag.register = function (name, tag) {
            Tag.impls[name] = tag;
        };
        Tag.clear = function () {
            Tag.impls = {};
        };
        Tag.impls = {};
        return Tag;
    }(Template));

    var Filter = /** @class */ (function () {
        function Filter(name, args, strictFilters) {
            var impl = Filter.impls[name];
            if (!impl && strictFilters)
                throw new TypeError("undefined filter: " + name);
            this.name = name;
            this.impl = impl || (function (x) { return x; });
            this.args = args;
        }
        Filter.prototype.render = function (value, ctx) {
            return __awaiter(this, void 0, void 0, function () {
                var argv, _i, _a, arg, _b, _c, _d, _e, _f;
                return __generator(this, function (_g) {
                    switch (_g.label) {
                        case 0:
                            argv = [];
                            _i = 0, _a = this.args;
                            _g.label = 1;
                        case 1:
                            if (!(_i < _a.length)) return [3 /*break*/, 6];
                            arg = _a[_i];
                            if (!isArray(arg)) return [3 /*break*/, 3];
                            _c = (_b = argv).push;
                            _d = [arg[0]];
                            return [4 /*yield*/, evalValue(arg[1], ctx)];
                        case 2:
                            _c.apply(_b, [_d.concat([_g.sent()])]);
                            return [3 /*break*/, 5];
                        case 3:
                            _f = (_e = argv).push;
                            return [4 /*yield*/, evalValue(arg, ctx)];
                        case 4:
                            _f.apply(_e, [_g.sent()]);
                            _g.label = 5;
                        case 5:
                            _i++;
                            return [3 /*break*/, 1];
                        case 6: return [2 /*return*/, this.impl.apply(null, [value].concat(argv))];
                    }
                });
            });
        };
        Filter.register = function (name, filter) {
            Filter.impls[name] = filter;
        };
        Filter.clear = function () {
            Filter.impls = {};
        };
        Filter.impls = {};
        return Filter;
    }());

    var ParseStream = /** @class */ (function () {
        function ParseStream(tokens, parseToken) {
            this.handlers = {};
            this.stopRequested = false;
            this.tokens = tokens;
            this.parseToken = parseToken;
        }
        ParseStream.prototype.on = function (name, cb) {
            this.handlers[name] = cb;
            return this;
        };
        ParseStream.prototype.trigger = function (event, arg) {
            var h = this.handlers[event];
            return h ? (h(arg), true) : false;
        };
        ParseStream.prototype.start = function () {
            this.trigger('start');
            var token;
            while (!this.stopRequested && (token = this.tokens.shift())) {
                if (this.trigger('token', token))
                    continue;
                if (TagToken.is(token) && this.trigger("tag:" + token.name, token)) {
                    continue;
                }
                var template = this.parseToken(token, this.tokens);
                this.trigger('template', template);
            }
            if (!this.stopRequested)
                this.trigger('end');
            return this;
        };
        ParseStream.prototype.stop = function () {
            this.stopRequested = true;
            return this;
        };
        return ParseStream;
    }());

    var Value = /** @class */ (function () {
        /**
         * @param str value string, like: "i have a dream | truncate: 3
         */
        function Value(str, strictFilters) {
            this.filters = [];
            var tokens = Value.tokenize(str);
            this.strictFilters = strictFilters;
            this.initial = tokens[0];
            this.parseFilters(tokens, 1);
        }
        Value.prototype.parseFilters = function (tokens, begin) {
            var i = begin;
            while (i < tokens.length) {
                if (tokens[i] !== '|') {
                    i++;
                    continue;
                }
                var j = ++i;
                while (i < tokens.length && tokens[i] !== '|')
                    i++;
                this.parseFilter(tokens, j, i);
            }
        };
        Value.prototype.parseFilter = function (tokens, begin, end) {
            var name = tokens[begin];
            var args = [];
            var argName, argValue;
            for (var i = begin + 1; i < end + 1; i++) {
                if (i === end || tokens[i] === ',') {
                    if (argName || argValue) {
                        args.push(argName ? [argName, argValue] : argValue);
                    }
                    argValue = argName = undefined;
                }
                else if (tokens[i] === ':') {
                    argName = argValue;
                    argValue = undefined;
                }
                else if (argValue === undefined) {
                    argValue = tokens[i];
                }
            }
            this.filters.push(new Filter(name, args, this.strictFilters));
        };
        Value.prototype.value = function (ctx) {
            return __awaiter(this, void 0, void 0, function () {
                var val, _i, _a, filter;
                return __generator(this, function (_b) {
                    switch (_b.label) {
                        case 0: return [4 /*yield*/, evalExp(this.initial, ctx)];
                        case 1:
                            val = _b.sent();
                            _i = 0, _a = this.filters;
                            _b.label = 2;
                        case 2:
                            if (!(_i < _a.length)) return [3 /*break*/, 5];
                            filter = _a[_i];
                            return [4 /*yield*/, filter.render(val, ctx)];
                        case 3:
                            val = _b.sent();
                            _b.label = 4;
                        case 4:
                            _i++;
                            return [3 /*break*/, 2];
                        case 5: return [2 /*return*/, val];
                    }
                });
            });
        };
        Value.tokenize = function (str) {
            var tokens = [];
            var i = 0;
            while (i < str.length) {
                var ch = str[i];
                if (ch === '"' || ch === "'") {
                    var j = i;
                    for (i += 2; i < str.length && str[i - 1] !== ch; ++i)
                        ;
                    tokens.push(str.slice(j, i));
                }
                else if (/\s/.test(ch)) {
                    i++;
                }
                else if (/[|,:]/.test(ch)) {
                    tokens.push(str[i++]);
                }
                else {
                    var j = i++;
                    for (; i < str.length && !/[|,:\s]/.test(str[i]); ++i)
                        ;
                    tokens.push(str.slice(j, i));
                }
            }
            return tokens;
        };
        return Value;
    }());

    var Output = /** @class */ (function (_super) {
        __extends(Output, _super);
        function Output(token, strictFilters) {
            var _this = _super.call(this, token) || this;
            _this.value = new Value(token.value, strictFilters);
            return _this;
        }
        Output.prototype.render = function (ctx) {
            return __awaiter(this, void 0, void 0, function () {
                var html;
                return __generator(this, function (_a) {
                    switch (_a.label) {
                        case 0: return [4 /*yield*/, this.value.value(ctx)];
                        case 1:
                            html = _a.sent();
                            return [2 /*return*/, stringify(html)];
                    }
                });
            });
        };
        return Output;
    }(Template));

    var default_1 = /** @class */ (function (_super) {
        __extends(default_1, _super);
        function default_1(token) {
            var _this = _super.call(this, token) || this;
            _this.str = token.value;
            return _this;
        }
        default_1.prototype.render = function () {
            return __awaiter(this, void 0, void 0, function () {
                return __generator(this, function (_a) {
                    return [2 /*return*/, this.str];
                });
            });
        };
        return default_1;
    }(Template));

    var Parser = /** @class */ (function () {
        function Parser(liquid) {
            this.liquid = liquid;
        }
        Parser.prototype.parse = function (tokens) {
            var token;
            var templates = [];
            while ((token = tokens.shift())) {
                templates.push(this.parseToken(token, tokens));
            }
            return templates;
        };
        Parser.prototype.parseToken = function (token, remainTokens) {
            try {
                if (TagToken.is(token)) {
                    return new Tag(token, remainTokens, this.liquid);
                }
                if (OutputToken.is(token)) {
                    return new Output(token, this.liquid.options.strictFilters);
                }
                return new default_1(token);
            }
            catch (e) {
                throw new ParseError(e, token);
            }
        };
        Parser.prototype.parseStream = function (tokens) {
            var _this = this;
            return new ParseStream(tokens, function (token, tokens) { return _this.parseToken(token, tokens); });
        };
        return Parser;
    }());

    var re = new RegExp("(" + identifier.source + ")\\s*=([^]*)");
    var assign = {
        parse: function (token) {
            var match = token.args.match(re);
            assert(match, "illegal token " + token.raw);
            this.key = match[1];
            this.value = match[2];
        },
        render: function (ctx) {
            return __awaiter(this, void 0, void 0, function () {
                var _a, _b;
                return __generator(this, function (_c) {
                    switch (_c.label) {
                        case 0:
                            _a = ctx.scopes[0];
                            _b = this.key;
                            return [4 /*yield*/, this.liquid.evalValue(this.value, ctx)];
                        case 1:
                            _a[_b] = _c.sent();
                            return [2 /*return*/];
                    }
                });
            });
        }
    };

    var ForloopDrop = /** @class */ (function (_super) {
        __extends(ForloopDrop, _super);
        function ForloopDrop(length) {
            var _this = _super.call(this) || this;
            _this.i = 0;
            _this.length = length;
            return _this;
        }
        ForloopDrop.prototype.next = function () {
            this.i++;
        };
        ForloopDrop.prototype.index0 = function () {
            return this.i;
        };
        ForloopDrop.prototype.index = function () {
            return this.i + 1;
        };
        ForloopDrop.prototype.first = function () {
            return this.i === 0;
        };
        ForloopDrop.prototype.last = function () {
            return this.i === this.length - 1;
        };
        ForloopDrop.prototype.rindex = function () {
            return this.length - this.i;
        };
        ForloopDrop.prototype.rindex0 = function () {
            return this.length - this.i - 1;
        };
        ForloopDrop.prototype.valueOf = function () {
            return JSON.stringify(this);
        };
        return ForloopDrop;
    }(Drop));

    var re$1 = new RegExp("^(" + identifier.source + ")\\s+in\\s+" +
        ("(" + value.source + ")") +
        ("(?:\\s+" + hash.source + ")*") +
        "(?:\\s+(reversed))?" +
        ("(?:\\s+" + hash.source + ")*$"));
    var For = {
        type: 'block',
        parse: function (tagToken, remainTokens) {
            var _this = this;
            var match = re$1.exec(tagToken.args);
            assert(match, "illegal tag: " + tagToken.raw);
            this.variable = match[1];
            this.collection = match[2];
            this.reversed = !!match[3];
            this.templates = [];
            this.elseTemplates = [];
            var p;
            var stream = this.liquid.parser.parseStream(remainTokens)
                .on('start', function () { return (p = _this.templates); })
                .on('tag:else', function () { return (p = _this.elseTemplates); })
                .on('tag:endfor', function () { return stream.stop(); })
                .on('template', function (tpl) { return p.push(tpl); })
                .on('end', function () {
                throw new Error("tag " + tagToken.raw + " not closed");
            });
            stream.start();
        },
        render: function (ctx, hash$$1) {
            return __awaiter(this, void 0, void 0, function () {
                var collection, offset, limit, context, html, _i, collection_1, item, _a, e_1;
                return __generator(this, function (_b) {
                    switch (_b.label) {
                        case 0: return [4 /*yield*/, evalExp(this.collection, ctx)];
                        case 1:
                            collection = _b.sent();
                            if (!isArray(collection)) {
                                if (isString(collection) && collection.length > 0) {
                                    collection = [collection];
                                }
                                else if (isObject(collection)) {
                                    collection = Object.keys(collection).map(function (key) { return [key, collection[key]]; });
                                }
                            }
                            if (!isArray(collection) || !collection.length) {
                                return [2 /*return*/, this.liquid.renderer.renderTemplates(this.elseTemplates, ctx)];
                            }
                            offset = hash$$1.offset || 0;
                            limit = (hash$$1.limit === undefined) ? collection.length : hash$$1.limit;
                            collection = collection.slice(offset, offset + limit);
                            if (this.reversed)
                                collection.reverse();
                            context = { forloop: new ForloopDrop(collection.length) };
                            ctx.push(context);
                            html = '';
                            _i = 0, collection_1 = collection;
                            _b.label = 2;
                        case 2:
                            if (!(_i < collection_1.length)) return [3 /*break*/, 8];
                            item = collection_1[_i];
                            context[this.variable] = item;
                            _b.label = 3;
                        case 3:
                            _b.trys.push([3, 5, , 6]);
                            _a = html;
                            return [4 /*yield*/, this.liquid.renderer.renderTemplates(this.templates, ctx)];
                        case 4:
                            html = _a + _b.sent();
                            return [3 /*break*/, 6];
                        case 5:
                            e_1 = _b.sent();
                            if (e_1.name === 'RenderBreakError') {
                                html += e_1.resolvedHTML;
                                if (e_1.message === 'break')
                                    return [3 /*break*/, 8];
                            }
                            else
                                throw e_1;
                            return [3 /*break*/, 6];
                        case 6:
                            context.forloop.next();
                            _b.label = 7;
                        case 7:
                            _i++;
                            return [3 /*break*/, 2];
                        case 8:
                            ctx.pop();
                            return [2 /*return*/, html];
                    }
                });
            });
        }
    };

    var re$2 = new RegExp("(" + identifier.source + ")");
    var capture = {
        parse: function (tagToken, remainTokens) {
            var _this = this;
            var match = tagToken.args.match(re$2);
            assert(match, tagToken.args + " not valid identifier");
            this.variable = match[1];
            this.templates = [];
            var stream = this.liquid.parser.parseStream(remainTokens);
            stream.on('tag:endcapture', function () { return stream.stop(); })
                .on('template', function (tpl) { return _this.templates.push(tpl); })
                .on('end', function () {
                throw new Error("tag " + tagToken.raw + " not closed");
            });
            stream.start();
        },
        render: function (ctx) {
            return __awaiter(this, void 0, void 0, function () {
                var html;
                return __generator(this, function (_a) {
                    switch (_a.label) {
                        case 0: return [4 /*yield*/, this.liquid.renderer.renderTemplates(this.templates, ctx)];
                        case 1:
                            html = _a.sent();
                            ctx.scopes[0][this.variable] = html;
                            return [2 /*return*/];
                    }
                });
            });
        }
    };

    var Case = {
        parse: function (tagToken, remainTokens) {
            var _this = this;
            this.cond = tagToken.args;
            this.cases = [];
            this.elseTemplates = [];
            var p = [];
            var stream = this.liquid.parser.parseStream(remainTokens)
                .on('tag:when', function (token) {
                _this.cases.push({
                    val: token.args,
                    templates: p = []
                });
            })
                .on('tag:else', function () { return (p = _this.elseTemplates); })
                .on('tag:endcase', function () { return stream.stop(); })
                .on('template', function (tpl) { return p.push(tpl); })
                .on('end', function () {
                throw new Error("tag " + tagToken.raw + " not closed");
            });
            stream.start();
        },
        render: function (ctx) {
            return __awaiter(this, void 0, void 0, function () {
                var i, branch, val, cond;
                return __generator(this, function (_a) {
                    switch (_a.label) {
                        case 0:
                            i = 0;
                            _a.label = 1;
                        case 1:
                            if (!(i < this.cases.length)) return [3 /*break*/, 5];
                            branch = this.cases[i];
                            return [4 /*yield*/, evalExp(branch.val, ctx)];
                        case 2:
                            val = _a.sent();
                            return [4 /*yield*/, evalExp(this.cond, ctx)];
                        case 3:
                            cond = _a.sent();
                            if (val === cond) {
                                return [2 /*return*/, this.liquid.renderer.renderTemplates(branch.templates, ctx)];
                            }
                            _a.label = 4;
                        case 4:
                            i++;
                            return [3 /*break*/, 1];
                        case 5: return [2 /*return*/, this.liquid.renderer.renderTemplates(this.elseTemplates, ctx)];
                    }
                });
            });
        }
    };

    var comment = {
        parse: function (tagToken, remainTokens) {
            var stream = this.liquid.parser.parseStream(remainTokens);
            stream
                .on('token', function (token) {
                if (token.name === 'endcomment')
                    stream.stop();
            })
                .on('end', function () {
                throw new Error("tag " + tagToken.raw + " not closed");
            });
            stream.start();
        }
    };

    var staticFileRE = /[^\s,]+/;
    var withRE = new RegExp("with\\s+(" + value.source + ")");
    var include = {
        parse: function (token) {
            var match = staticFileRE.exec(token.args);
            if (match) {
                this.staticValue = match[0];
            }
            match = value.exec(token.args);
            if (match) {
                this.value = match[0];
            }
            match = withRE.exec(token.args);
            if (match) {
                this.with = match[1];
            }
        },
        render: function (ctx, hash$$1) {
            return __awaiter(this, void 0, void 0, function () {
                var filepath, template, originBlocks, originBlockMode, _a, _b, templates, html;
                return __generator(this, function (_c) {
                    switch (_c.label) {
                        case 0:
                            if (!ctx.opts.dynamicPartials) return [3 /*break*/, 5];
                            if (!quotedLine.exec(this.value)) return [3 /*break*/, 2];
                            template = this.value.slice(1, -1);
                            return [4 /*yield*/, this.liquid.parseAndRender(template, ctx.getAll(), ctx.opts)];
                        case 1:
                            filepath = _c.sent();
                            return [3 /*break*/, 4];
                        case 2: return [4 /*yield*/, evalValue(this.value, ctx)];
                        case 3:
                            filepath = _c.sent();
                            _c.label = 4;
                        case 4: return [3 /*break*/, 6];
                        case 5:
                            filepath = this.staticValue;
                            _c.label = 6;
                        case 6:
                            assert(filepath, "cannot include with empty filename");
                            originBlocks = ctx.blocks;
                            originBlockMode = ctx.blockMode;
                            ctx.blocks = {};
                            ctx.blockMode = BlockMode$1.OUTPUT;
                            if (!this.with) return [3 /*break*/, 8];
                            _a = hash$$1;
                            _b = filepath;
                            return [4 /*yield*/, evalValue(this.with, ctx)];
                        case 7:
                            _a[_b] = _c.sent();
                            _c.label = 8;
                        case 8: return [4 /*yield*/, this.liquid.getTemplate(filepath, ctx.opts)];
                        case 9:
                            templates = _c.sent();
                            ctx.push(hash$$1);
                            return [4 /*yield*/, this.liquid.renderer.renderTemplates(templates, ctx)];
                        case 10:
                            html = _c.sent();
                            ctx.pop(hash$$1);
                            ctx.blocks = originBlocks;
                            ctx.blockMode = originBlockMode;
                            return [2 /*return*/, html];
                    }
                });
            });
        }
    };

    var decrement = {
        parse: function (token) {
            var match = token.args.match(identifier);
            assert(match, "illegal identifier " + token.args);
            this.variable = match[0];
        },
        render: function (context) {
            var scope = context.environments;
            if (typeof scope[this.variable] !== 'number') {
                scope[this.variable] = 0;
            }
            return --scope[this.variable];
        }
    };

    var groupRE = new RegExp("^(?:(" + value.source + ")\\s*:\\s*)?(.*)$");
    var candidatesRE = new RegExp(value.source, 'g');
    var cycle = {
        parse: function (tagToken) {
            var match = groupRE.exec(tagToken.args);
            assert(match, "illegal tag: " + tagToken.raw);
            this.group = match[1] || '';
            var candidates = match[2];
            this.candidates = [];
            while ((match = candidatesRE.exec(candidates))) {
                this.candidates.push(match[0]);
            }
            assert(this.candidates.length, "empty candidates: " + tagToken.raw);
        },
        render: function (ctx) {
            return __awaiter(this, void 0, void 0, function () {
                var group, fingerprint, groups, idx, candidate;
                return __generator(this, function (_a) {
                    switch (_a.label) {
                        case 0: return [4 /*yield*/, evalValue(this.group, ctx)];
                        case 1:
                            group = _a.sent();
                            fingerprint = "cycle:" + group + ":" + this.candidates.join(',');
                            groups = ctx.groups;
                            idx = groups[fingerprint];
                            if (idx === undefined) {
                                idx = groups[fingerprint] = 0;
                            }
                            candidate = this.candidates[idx];
                            idx = (idx + 1) % this.candidates.length;
                            groups[fingerprint] = idx;
                            return [2 /*return*/, evalValue(candidate, ctx)];
                    }
                });
            });
        }
    };

    var If = {
        parse: function (tagToken, remainTokens) {
            var _this = this;
            this.branches = [];
            this.elseTemplates = [];
            var p;
            var stream = this.liquid.parser.parseStream(remainTokens)
                .on('start', function () { return _this.branches.push({
                cond: tagToken.args,
                templates: (p = [])
            }); })
                .on('tag:elsif', function (token) {
                _this.branches.push({
                    cond: token.args,
                    templates: p = []
                });
            })
                .on('tag:else', function () { return (p = _this.elseTemplates); })
                .on('tag:endif', function () { return stream.stop(); })
                .on('template', function (tpl) { return p.push(tpl); })
                .on('end', function () {
                throw new Error("tag " + tagToken.raw + " not closed");
            });
            stream.start();
        },
        render: function (ctx) {
            return __awaiter(this, void 0, void 0, function () {
                var _i, _a, branch, cond;
                return __generator(this, function (_b) {
                    switch (_b.label) {
                        case 0:
                            _i = 0, _a = this.branches;
                            _b.label = 1;
                        case 1:
                            if (!(_i < _a.length)) return [3 /*break*/, 4];
                            branch = _a[_i];
                            return [4 /*yield*/, evalExp(branch.cond, ctx)];
                        case 2:
                            cond = _b.sent();
                            if (isTruthy(cond)) {
                                return [2 /*return*/, this.liquid.renderer.renderTemplates(branch.templates, ctx)];
                            }
                            _b.label = 3;
                        case 3:
                            _i++;
                            return [3 /*break*/, 1];
                        case 4: return [2 /*return*/, this.liquid.renderer.renderTemplates(this.elseTemplates, ctx)];
                    }
                });
            });
        }
    };

    var increment = {
        parse: function (token) {
            var match = token.args.match(identifier);
            assert(match, "illegal identifier " + token.args);
            this.variable = match[0];
        },
        render: function (context) {
            var scope = context.environments;
            if (typeof scope[this.variable] !== 'number') {
                scope[this.variable] = 0;
            }
            var val = scope[this.variable];
            scope[this.variable]++;
            return val;
        }
    };

    var staticFileRE$1 = /\S+/;
    var layout = {
        parse: function (token, remainTokens) {
            var match = staticFileRE$1.exec(token.args);
            if (match) {
                this.staticLayout = match[0];
            }
            match = value.exec(token.args);
            if (match) {
                this.layout = match[0];
            }
            this.tpls = this.liquid.parser.parse(remainTokens);
        },
        render: function (ctx, hash$$1) {
            return __awaiter(this, void 0, void 0, function () {
                var layout, _a, html, templates, partial;
                return __generator(this, function (_b) {
                    switch (_b.label) {
                        case 0:
                            if (!ctx.opts.dynamicPartials) return [3 /*break*/, 2];
                            return [4 /*yield*/, evalValue(this.layout, ctx)];
                        case 1:
                            _a = _b.sent();
                            return [3 /*break*/, 3];
                        case 2:
                            _a = this.staticLayout;
                            _b.label = 3;
                        case 3:
                            layout = _a;
                            assert(layout, "cannot apply layout with empty filename");
                            // render the remaining tokens immediately
                            ctx.blockMode = BlockMode$1.STORE;
                            return [4 /*yield*/, this.liquid.renderer.renderTemplates(this.tpls, ctx)];
                        case 4:
                            html = _b.sent();
                            if (ctx.blocks[''] === undefined) {
                                ctx.blocks[''] = html;
                            }
                            return [4 /*yield*/, this.liquid.getTemplate(layout, ctx.opts)];
                        case 5:
                            templates = _b.sent();
                            ctx.push(hash$$1);
                            ctx.blockMode = BlockMode$1.OUTPUT;
                            return [4 /*yield*/, this.liquid.renderer.renderTemplates(templates, ctx)];
                        case 6:
                            partial = _b.sent();
                            ctx.pop(hash$$1);
                            return [2 /*return*/, partial];
                    }
                });
            });
        }
    };

    var block = {
        parse: function (token, remainTokens) {
            var _this = this;
            var match = /\w+/.exec(token.args);
            this.block = match ? match[0] : '';
            this.tpls = [];
            var stream = this.liquid.parser.parseStream(remainTokens)
                .on('tag:endblock', function () { return stream.stop(); })
                .on('template', function (tpl) { return _this.tpls.push(tpl); })
                .on('end', function () {
                throw new Error("tag " + token.raw + " not closed");
            });
            stream.start();
        },
        render: function (ctx) {
            return __awaiter(this, void 0, void 0, function () {
                var childDefined, html, _a;
                return __generator(this, function (_b) {
                    switch (_b.label) {
                        case 0:
                            childDefined = ctx.blocks[this.block];
                            if (!(childDefined !== undefined)) return [3 /*break*/, 1];
                            _a = childDefined;
                            return [3 /*break*/, 3];
                        case 1: return [4 /*yield*/, this.liquid.renderer.renderTemplates(this.tpls, ctx)];
                        case 2:
                            _a = _b.sent();
                            _b.label = 3;
                        case 3:
                            html = _a;
                            if (ctx.blockMode === BlockMode$1.STORE) {
                                ctx.blocks[this.block] = html;
                                return [2 /*return*/, ''];
                            }
                            return [2 /*return*/, html];
                    }
                });
            });
        }
    };

    var raw = {
        parse: function (tagToken, remainTokens) {
            var _this = this;
            this.tokens = [];
            var stream = this.liquid.parser.parseStream(remainTokens);
            stream
                .on('token', function (token) {
                if (token.name === 'endraw')
                    stream.stop();
                else
                    _this.tokens.push(token);
            })
                .on('end', function () {
                throw new Error("tag " + tagToken.raw + " not closed");
            });
            stream.start();
        },
        render: function () {
            return this.tokens.map(function (token) { return token.raw; }).join('');
        }
    };

    var TablerowloopDrop = /** @class */ (function (_super) {
        __extends(TablerowloopDrop, _super);
        function TablerowloopDrop(length, cols) {
            var _this = _super.call(this, length) || this;
            _this.length = length;
            _this.cols = cols;
            return _this;
        }
        TablerowloopDrop.prototype.row = function () {
            return Math.floor(this.i / this.cols) + 1;
        };
        TablerowloopDrop.prototype.col0 = function () {
            return (this.i % this.cols);
        };
        TablerowloopDrop.prototype.col = function () {
            return this.col0() + 1;
        };
        TablerowloopDrop.prototype.col_first = function () {
            return this.col0() === 0;
        };
        TablerowloopDrop.prototype.col_last = function () {
            return this.col() === this.cols;
        };
        return TablerowloopDrop;
    }(ForloopDrop));

    var re$3 = new RegExp("^(" + identifier.source + ")\\s+in\\s+" +
        ("(" + value.source + ")") +
        ("(?:\\s+" + hash.source + ")*$"));
    var tablerow = {
        parse: function (tagToken, remainTokens) {
            var _this = this;
            var match = re$3.exec(tagToken.args);
            assert(match, "illegal tag: " + tagToken.raw);
            this.variable = match[1];
            this.collection = match[2];
            this.templates = [];
            var p;
            var stream = this.liquid.parser.parseStream(remainTokens)
                .on('start', function () { return (p = _this.templates); })
                .on('tag:endtablerow', function () { return stream.stop(); })
                .on('template', function (tpl) { return p.push(tpl); })
                .on('end', function () {
                throw new Error("tag " + tagToken.raw + " not closed");
            });
            stream.start();
        },
        render: function (ctx, hash$$1) {
            return __awaiter(this, void 0, void 0, function () {
                var collection, offset, limit, cols, tablerowloop, scope, html, idx, _a;
                return __generator(this, function (_b) {
                    switch (_b.label) {
                        case 0: return [4 /*yield*/, evalExp(this.collection, ctx)];
                        case 1:
                            collection = (_b.sent()) || [];
                            offset = hash$$1.offset || 0;
                            limit = (hash$$1.limit === undefined) ? collection.length : hash$$1.limit;
                            collection = collection.slice(offset, offset + limit);
                            cols = hash$$1.cols || collection.length;
                            tablerowloop = new TablerowloopDrop(collection.length, cols);
                            scope = { tablerowloop: tablerowloop };
                            ctx.push(scope);
                            html = '';
                            idx = 0;
                            _b.label = 2;
                        case 2:
                            if (!(idx < collection.length)) return [3 /*break*/, 5];
                            scope[this.variable] = collection[idx];
                            if (tablerowloop.col0() === 0) {
                                if (tablerowloop.row() !== 1)
                                    html += '</tr>';
                                html += "<tr class=\"row" + tablerowloop.row() + "\">";
                            }
                            html += "<td class=\"col" + tablerowloop.col() + "\">";
                            _a = html;
                            return [4 /*yield*/, this.liquid.renderer.renderTemplates(this.templates, ctx)];
                        case 3:
                            html = _a + _b.sent();
                            html += '</td>';
                            _b.label = 4;
                        case 4:
                            idx++, tablerowloop.next();
                            return [3 /*break*/, 2];
                        case 5:
                            if (collection.length)
                                html += '</tr>';
                            ctx.pop(scope);
                            return [2 /*return*/, html];
                    }
                });
            });
        }
    };

    var unless = {
        parse: function (tagToken, remainTokens) {
            var _this = this;
            this.templates = [];
            this.elseTemplates = [];
            var p;
            var stream = this.liquid.parser.parseStream(remainTokens)
                .on('start', function () {
                p = _this.templates;
                _this.cond = tagToken.args;
            })
                .on('tag:else', function () { return (p = _this.elseTemplates); })
                .on('tag:endunless', function () { return stream.stop(); })
                .on('template', function (tpl) { return p.push(tpl); })
                .on('end', function () {
                throw new Error("tag " + tagToken.raw + " not closed");
            });
            stream.start();
        },
        render: function (ctx) {
            return __awaiter(this, void 0, void 0, function () {
                var cond;
                return __generator(this, function (_a) {
                    switch (_a.label) {
                        case 0: return [4 /*yield*/, evalExp(this.cond, ctx)];
                        case 1:
                            cond = _a.sent();
                            return [2 /*return*/, isFalsy(cond)
                                    ? this.liquid.renderer.renderTemplates(this.templates, ctx)
                                    : this.liquid.renderer.renderTemplates(this.elseTemplates, ctx)];
                    }
                });
            });
        }
    };

    var Break = {
        render: function () {
            return __awaiter(this, void 0, void 0, function () {
                return __generator(this, function (_a) {
                    throw new RenderBreakError('break');
                });
            });
        }
    };

    var Continue = {
        render: function () {
            return __awaiter(this, void 0, void 0, function () {
                return __generator(this, function (_a) {
                    throw new RenderBreakError('continue');
                });
            });
        }
    };

    var tags = {
        assign: assign, 'for': For, capture: capture, 'case': Case, comment: comment, include: include, decrement: decrement, increment: increment, cycle: cycle, 'if': If, layout: layout, block: block, raw: raw, tablerow: tablerow, unless: unless, 'break': Break, 'continue': Continue
    };

    var escapeMap = {
        '&': '&amp;',
        '<': '&lt;',
        '>': '&gt;',
        '"': '&#34;',
        "'": '&#39;'
    };
    var unescapeMap = {
        '&amp;': '&',
        '&lt;': '<',
        '&gt;': '>',
        '&#34;': '"',
        '&#39;': "'"
    };
    function escape(str) {
        return String(str).replace(/&|<|>|"|'/g, function (m) { return escapeMap[m]; });
    }
    function unescape(str) {
        return String(str).replace(/&(amp|lt|gt|#34|#39);/g, function (m) { return unescapeMap[m]; });
    }
    var html = {
        'escape': escape,
        'escape_once': function (str) { return escape(unescape(str)); },
        'newline_to_br': function (v) { return v.replace(/\n/g, '<br />'); },
        'strip_html': function (v) { return v.replace(/<script.*?<\/script>|<!--.*?-->|<style.*?<\/style>|<.*?>/g, ''); }
    };

    var str = {
        'append': function (v, arg) { return v + arg; },
        'prepend': function (v, arg) { return arg + v; },
        'capitalize': function (str) { return String(str).charAt(0).toUpperCase() + str.slice(1); },
        'lstrip': function (v) { return String(v).replace(/^\s+/, ''); },
        'downcase': function (v) { return v.toLowerCase(); },
        'upcase': function (str) { return String(str).toUpperCase(); },
        'remove': function (v, arg) { return v.split(arg).join(''); },
        'remove_first': function (v, l) { return v.replace(l, ''); },
        'replace': function (v, pattern, replacement) {
            return String(v).split(pattern).join(replacement);
        },
        'replace_first': function (v, arg1, arg2) { return String(v).replace(arg1, arg2); },
        'rstrip': function (str) { return String(str).replace(/\s+$/, ''); },
        'split': function (v, arg) { return String(v).split(arg); },
        'strip': function (v) { return String(v).trim(); },
        'strip_newlines': function (v) { return String(v).replace(/\n/g, ''); },
        'truncate': function (v, l, o) {
            if (l === void 0) { l = 50; }
            if (o === void 0) { o = '...'; }
            v = String(v);
            if (v.length <= l)
                return v;
            return v.substr(0, l - o.length) + o;
        },
        'truncatewords': function (v, l, o) {
            if (l === void 0) { l = 15; }
            if (o === void 0) { o = '...'; }
            var arr = v.split(/\s+/);
            var ret = arr.slice(0, l).join(' ');
            if (arr.length >= l)
                ret += o;
            return ret;
        }
    };

    var math = {
        'abs': function (v) { return Math.abs(v); },
        'ceil': function (v) { return Math.ceil(v); },
        'divided_by': function (v, arg) { return v / arg; },
        'floor': function (v) { return Math.floor(v); },
        'minus': function (v, arg) { return v - arg; },
        'modulo': function (v, arg) { return v % arg; },
        'round': function (v, arg) {
            if (arg === void 0) { arg = 0; }
            var amp = Math.pow(10, arg);
            return Math.round(v * amp) / amp;
        },
        'plus': function (v, arg) { return Number(v) + Number(arg); },
        'times': function (v, arg) { return v * arg; }
    };

    var url = {
        'url_decode': function (x) { return x.split('+').map(decodeURIComponent).join(' '); },
        'url_encode': function (x) { return x.split(' ').map(encodeURIComponent).join('+'); }
    };

    var array = {
        'join': function (v, arg) { return v.join(arg === undefined ? ' ' : arg); },
        'last': function (v) { return last(v); },
        'first': function (v) { return v[0]; },
        'map': function (arr, arg) { return arr.map(function (v) { return v[arg]; }); },
        'reverse': function (v) { return v.reverse(); },
        'sort': function (v, arg) { return v.sort(arg); },
        'size': function (v) { return v.length; },
        'concat': function (v, arg) { return Array.prototype.concat.call(v, arg); },
        'slice': function (v, begin, length) {
            if (length === undefined)
                length = 1;
            return v.slice(begin, begin + length);
        },
        'uniq': function (arr) {
            var u = {};
            return (arr || []).filter(function (val) {
                if (u.hasOwnProperty(String(val)))
                    return false;
                u[String(val)] = true;
                return true;
            });
        }
    };

    var monthNames = [
        'January', 'February', 'March', 'April', 'May', 'June', 'July', 'August',
        'September', 'October', 'November', 'December'
    ];
    var dayNames = [
        'Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'
    ];
    var monthNamesShort = monthNames.map(abbr);
    var dayNamesShort = dayNames.map(abbr);
    var suffixes = {
        1: 'st',
        2: 'nd',
        3: 'rd',
        'default': 'th'
    };
    function abbr(str) {
        return str.slice(0, 3);
    }
    // prototype extensions
    var _date = {
        daysInMonth: function (d) {
            var feb = _date.isLeapYear(d) ? 29 : 28;
            return [31, feb, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
        },
        getDayOfYear: function (d) {
            var num = 0;
            for (var i = 0; i < d.getMonth(); ++i) {
                num += _date.daysInMonth(d)[i];
            }
            return num + d.getDate();
        },
        getWeekOfYear: function (d, startDay) {
            // Skip to startDay of this week
            var now = this.getDayOfYear(d) + (startDay - d.getDay());
            // Find the first startDay of the year
            var jan1 = new Date(d.getFullYear(), 0, 1);
            var then = (7 - jan1.getDay() + startDay);
            return padStart(String(Math.floor((now - then) / 7) + 1), 2, '0');
        },
        isLeapYear: function (d) {
            var year = d.getFullYear();
            return !!((year & 3) === 0 && (year % 100 || (year % 400 === 0 && year)));
        },
        getSuffix: function (d) {
            var str = d.getDate().toString();
            var index = parseInt(str.slice(-1));
            return suffixes[index] || suffixes['default'];
        },
        century: function (d) {
            return parseInt(d.getFullYear().toString().substring(0, 2), 10);
        }
    };
    var formatCodes = {
        a: function (d) {
            return dayNamesShort[d.getDay()];
        },
        A: function (d) {
            return dayNames[d.getDay()];
        },
        b: function (d) {
            return monthNamesShort[d.getMonth()];
        },
        B: function (d) {
            return monthNames[d.getMonth()];
        },
        c: function (d) {
            return d.toLocaleString();
        },
        C: function (d) {
            return _date.century(d);
        },
        d: function (d) {
            return padStart(d.getDate(), 2, '0');
        },
        e: function (d) {
            return padStart(d.getDate(), 2);
        },
        H: function (d) {
            return padStart(d.getHours(), 2, '0');
        },
        I: function (d) {
            return padStart(String(d.getHours() % 12 || 12), 2, '0');
        },
        j: function (d) {
            return padStart(_date.getDayOfYear(d), 3, '0');
        },
        k: function (d) {
            return padStart(d.getHours(), 2);
        },
        l: function (d) {
            return padStart(String(d.getHours() % 12 || 12), 2);
        },
        L: function (d) {
            return padStart(d.getMilliseconds(), 3, '0');
        },
        m: function (d) {
            return padStart(d.getMonth() + 1, 2, '0');
        },
        M: function (d) {
            return padStart(d.getMinutes(), 2, '0');
        },
        p: function (d) {
            return (d.getHours() < 12 ? 'AM' : 'PM');
        },
        P: function (d) {
            return (d.getHours() < 12 ? 'am' : 'pm');
        },
        q: function (d) {
            return _date.getSuffix(d);
        },
        s: function (d) {
            return Math.round(d.valueOf() / 1000);
        },
        S: function (d) {
            return padStart(d.getSeconds(), 2, '0');
        },
        u: function (d) {
            return d.getDay() || 7;
        },
        U: function (d) {
            return _date.getWeekOfYear(d, 0);
        },
        w: function (d) {
            return d.getDay();
        },
        W: function (d) {
            return _date.getWeekOfYear(d, 1);
        },
        x: function (d) {
            return d.toLocaleDateString();
        },
        X: function (d) {
            return d.toLocaleTimeString();
        },
        y: function (d) {
            return d.getFullYear().toString().substring(2, 4);
        },
        Y: function (d) {
            return d.getFullYear();
        },
        z: function (d) {
            var tz = d.getTimezoneOffset() / 60 * 100;
            return (tz > 0 ? '-' : '+') + padStart(String(Math.abs(tz)), 4, '0');
        },
        '%': function () {
            return '%';
        }
    };
    formatCodes.h = formatCodes.b;
    formatCodes.N = formatCodes.L;
    function strftime (d, format) {
        var output = '';
        var remaining = format;
        while (true) {
            var r = /%./g;
            var results = r.exec(remaining);
            // No more format codes. Add the remaining text and return
            if (!results) {
                return output + remaining;
            }
            // Add the preceding text
            output += remaining.slice(0, r.lastIndex - 2);
            remaining = remaining.slice(r.lastIndex);
            // Add the format code
            var ch = results[0].charAt(1);
            var func = formatCodes[ch];
            output += func ? func(d) : '%' + ch;
        }
    }

    var date = {
        'date': function (v, arg) {
            var date = v;
            if (v === 'now') {
                date = new Date();
            }
            else if (isString(v)) {
                date = new Date(v);
            }
            return isValidDate(date) ? strftime(date, arg) : v;
        }
    };
    function isValidDate(date) {
        return date instanceof Date && !isNaN(date.getTime());
    }

    var obj = {
        'default': function (v, arg) { return isTruthy(v) ? v : arg; }
    };

    var builtinFilters = __assign({}, html, str, math, url, date, obj, array);

    var Liquid = /** @class */ (function () {
        function Liquid(opts) {
            if (opts === void 0) { opts = {}; }
            var _this = this;
            this.cache = {};
            this.options = applyDefault(normalize(opts));
            this.parser = new Parser(this);
            this.renderer = new Render();
            this.tokenizer = new Tokenizer(this.options);
            forOwn(tags, function (conf, name) { return _this.registerTag(name, conf); });
            forOwn(builtinFilters, function (handler, name) { return _this.registerFilter(name, handler); });
        }
        Liquid.prototype.parse = function (html, filepath) {
            var tokens = this.tokenizer.tokenize(html, filepath);
            return this.parser.parse(tokens);
        };
        Liquid.prototype.render = function (tpl, ctx, opts) {
            var options = __assign({}, this.options, normalize(opts));
            var scope = new Context(ctx, options);
            return this.renderer.renderTemplates(tpl, scope);
        };
        Liquid.prototype.parseAndRender = function (html, ctx, opts) {
            return __awaiter(this, void 0, void 0, function () {
                var tpl;
                return __generator(this, function (_a) {
                    switch (_a.label) {
                        case 0: return [4 /*yield*/, this.parse(html)];
                        case 1:
                            tpl = _a.sent();
                            return [2 /*return*/, this.render(tpl, ctx, opts)];
                    }
                });
            });
        };
        Liquid.prototype.getTemplate = function (file, opts) {
            return __awaiter(this, void 0, void 0, function () {
                var options, roots, paths, _i, paths_1, filepath, value, _a, err;
                var _this = this;
                return __generator(this, function (_b) {
                    switch (_b.label) {
                        case 0:
                            options = normalize(opts);
                            roots = options.root ? options.root.concat(this.options.root) : this.options.root;
                            paths = roots.map(function (root) { return fs.resolve(root, file, _this.options.extname); });
                            _i = 0, paths_1 = paths;
                            _b.label = 1;
                        case 1:
                            if (!(_i < paths_1.length)) return [3 /*break*/, 5];
                            filepath = paths_1[_i];
                            return [4 /*yield*/, fs.exists(filepath)];
                        case 2:
                            if (!(_b.sent()))
                                return [3 /*break*/, 4];
                            if (this.options.cache && this.cache[filepath])
                                return [2 /*return*/, this.cache[filepath]];
                            _a = this.parse;
                            return [4 /*yield*/, fs.readFile(filepath)];
                        case 3:
                            value = _a.apply(this, [_b.sent(), filepath]);
                            if (this.options.cache)
                                this.cache[filepath] = value;
                            return [2 /*return*/, value];
                        case 4:
                            _i++;
                            return [3 /*break*/, 1];
                        case 5:
                            err = new Error('ENOENT');
                            err.message = "ENOENT: Failed to lookup \"" + file + "\" in \"" + roots + "\"";
                            err.code = 'ENOENT';
                            throw err;
                    }
                });
            });
        };
        Liquid.prototype.renderFile = function (file, ctx, opts) {
            return __awaiter(this, void 0, void 0, function () {
                var options, templates;
                return __generator(this, function (_a) {
                    switch (_a.label) {
                        case 0:
                            options = normalize(opts);
                            return [4 /*yield*/, this.getTemplate(file, options)];
                        case 1:
                            templates = _a.sent();
                            return [2 /*return*/, this.render(templates, ctx, opts)];
                    }
                });
            });
        };
        Liquid.prototype.evalValue = function (str, ctx) {
            return new Value(str, this.options.strictFilters).value(ctx);
        };
        Liquid.prototype.registerFilter = function (name, filter) {
            return Filter.register(name, filter);
        };
        Liquid.prototype.registerTag = function (name, tag) {
            return Tag.register(name, tag);
        };
        Liquid.prototype.plugin = function (plugin) {
            return plugin.call(this, Liquid);
        };
        Liquid.prototype.express = function () {
            var self = this;
            return function (filePath, ctx, cb) {
                var opts = { root: this.root };
                self.renderFile(filePath, ctx, opts).then(function (html) { return cb(null, html); }, cb);
            };
        };
        Liquid.default = Liquid;
        Liquid.isTruthy = isTruthy;
        Liquid.isFalsy = isFalsy;
        Liquid.evalExp = evalExp;
        Liquid.evalValue = evalValue;
        Liquid.Types = Types;
        return Liquid;
    }());

    return Liquid;

}));
//# sourceMappingURL=liquid.js.map
