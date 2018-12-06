/*
 * @file H5转盘抽奖插件
 * @author 梁亚军 1148063373@qq.com 2018.11.15
 * luckDraw.js 文件参数说明
 *
 * 插件传入参数说明：
 * pointerId 指针对象的id
 * turntableId 转盘对象的id
 * rotateId 当前旋转对象的id，可以是指针也可以是转盘
 * luckNumber 指针最后停止的位置（对应奖品的位置）
 * typeNumber 转盘奖品个数的设置
 * circleNumber 旋转对象的旋转圈数，决定转速
 * time 旋转的时长，决定转速‘
 * success 旋转停止后进行的操作回调函数
 * 
 * 插件全局参数说明：
 * defaultOPts 插件基础参数存储对象
 * flag 判断当前状态是否在旋转，true在旋转，false停止旋转
 * pointer 指针对象，没有用，可以不传
 * turntable 转盘对象，没有用，可以不传
 * rotate 旋转对象，必传
 * _resetVariable 重置全局变量方法
 * _keepTwoDecimalFull 保留两位小数方法
 * _getRandomNumber 获取两个数之间的随机数方法
 * _getVendorPrefix 获取当前浏览器前缀方法
 * _turn 旋转的兼容性操作方法
 * _running 开始旋转操作方法
 * _init 启动抽奖方法，注意此方法可以重置中奖奖品等参数
 */

(function (win) {
    function LuckDraw(opts) {
        this.defaultOPts = {
            pointerId: 'luckPointer',
            turntableId: 'luckTurntable',
            rotateId: 'luckTurntable',
            activeClass: 'rui-active',
            type: 'turntable',
            time: 2000,
            luckNumber: 4,
            typeNumber: 6,
            circleNumber: 10,
            success: function () { }
        };
        // 扩展defaultOPts对象
        for (var key in opts) {
            this.defaultOPts[key] = opts[key];
        }
        // 公用参数
        this.flag = true;
    }
    // 重置全局变量
    LuckDraw.prototype._resetVariable = function () {
        this.flag = true;
        this.index = 0;
        this.timer = null;
        this.endCircle = 1;
        this.span = 5;
        this.startTime = 20;
        this.degAverage = this._keepTwoDecimalFull(360 / this.defaultOPts.typeNumber);
        this.minDeg = this.degAverage * this.defaultOPts.luckNumber + this.degAverage / 4 - this.degAverage / 2;
        this.maxDeg = this.degAverage * (this.defaultOPts.luckNumber + 1) - this.degAverage / 4 - this.degAverage / 2;
        this.luckDeg = this._getRandomNumber(this.minDeg, this.maxDeg);
        this.totalLuckDeg = 360 * this.defaultOPts.circleNumber + this.luckDeg;
        this.totalLuckNumber = this.defaultOPts.typeNumber * this.defaultOPts.circleNumber + (this.defaultOPts.luckNumber + 1);
        this.stepTime = Math.floor(this.defaultOPts.time / this.totalLuckNumber);
        this.rotate = this._getElem(this.defaultOPts.rotateId);
    }
    // 下取整保留两位小数的方法
    LuckDraw.prototype._keepTwoDecimalFull = function (num) {
        return Math.floor(num * 100) / 100;
    }
    // 获取两个数之间的随机数
    LuckDraw.prototype._getRandomNumber = function (min, max) {
        return Math.floor(Math.random() * (max - min) + min);
    }
    // 获取浏览器前缀
    LuckDraw.prototype._getVendorPrefix = function () {
        var body, i, style, transition, vendor;
        body = document.body || document.documentElement;
        style = body.style;
        transition = "transition";
        vendor = ["Moz", "Webkit", "Khtml", "O", "ms"];
        transition = transition.charAt(0).toUpperCase() + transition.substr(1);
        i = 0;
        while (i < vendor.length) {
            if (typeof style[vendor[i] + transition] === "string") {
                return vendor[i];
            }
            i++;
        }
        return false;
    }
    // 旋转操作
    LuckDraw.prototype._turn = function (time, deg) {
        var vendor = this._getVendorPrefix();
        
        /* 兼容性处理 */
        this.rotate[0].style[vendor + 'TransitionTimingFunction'] = 'cubic-bezier(.53,.3,.24,1.01)';
        this.rotate[0].style[vendor + 'TransitionDuration'] = time + 'ms';
        this.rotate[0].style[vendor + 'TransitionProperty'] = 'all';
        this.rotate[0].style[vendor + 'Transform'] = 'translate(-50%,-50%) rotate(' + deg + 'deg)';
    }
    // 改变角标
    LuckDraw.prototype._changePrize = function(){
        var _this = this;
        var resetIndex = function(){
            for (var j = 0, len = _this.rotate.length ; j < len ; j++){
                _this.rotate[j].className = _this.defaultOPts.rotateId.split('.')[1];
            }
            var index = _this.index % _this.defaultOPts.typeNumber;
            _this.rotate[index].className = _this.defaultOPts.rotateId.split('.')[1] + ' ' + _this.defaultOPts.activeClass;
            _this.index++;
            
            if (_this.index < _this.totalLuckNumber - _this.defaultOPts.typeNumber * _this.endCircle) {
                _this.startTime += 1;
                _this.timer = setTimeout(resetIndex, _this.startTime);
            } else if (_this.index < _this.totalLuckNumber && _this.index >= _this.totalLuckNumber - _this.defaultOPts.typeNumber * _this.endCircle) {
                _this.startTime += 50;
                _this.timer = setTimeout(resetIndex, _this.startTime);
            }else{
                clearTimeout(_this.timer);
                setTimeout(function(){
                    _this._resetVariable();
                    _this.defaultOPts.success(_this.defaultOPts);
                }, _this.startTime);
            }
        }
        _this.timer = setTimeout(resetIndex, _this.startTime);
    }
    // 指针旋转函数
    LuckDraw.prototype._running = function () {
        var _this = this;
        var newTime = _this.defaultOPts.time + 600;
        setTimeout(function () {
            switch (_this.defaultOPts.type) {
                case 'turntable':
                    _this._turn(_this.defaultOPts.time, _this.totalLuckDeg);
                    break;
                case 'prize':
                    _this._changePrize();
                    break;
                default:
                    _this._turn(_this.defaultOPts.time, _this.totalLuckDeg);
                    break;
            }
        }, 100);
        if (_this.defaultOPts.type === 'turntable'){
            // 停止旋转后操作
            setTimeout(function () {
                _this._turn(0, _this.luckDeg);
                _this.defaultOPts.success(_this.defaultOPts);
                _this._resetVariable();
            }, newTime);
        }
    }
    // 参数测试，抛出异常
    LuckDraw.prototype._paramTesting = function(){
        if (!this.rotate[0] || this.rotate[0] === null) {
            throw this.defaultOPts.rotateId + ' is not find!';
        }
        if (this.defaultOPts.luckNumber >= this.defaultOPts.typeNumber) {
            throw '中奖坐标大于了奖品个数，请重新设置！';
        }
        if (this.defaultOPts.time < 1000) {
            throw '抽奖时间必须大于1000ms！';
        }
    }
    // 获取元素
    LuckDraw.prototype._getElem = function (str, parEle) {
        str = str.split(" ");
        var par = [];
        parEle = parEle || document;
        var retn = [parEle];
        for (var i in str) { if (str[i].length != 0) par.push(str[i]) } //去掉重复空格
        for (var i in par) {
            if (retn.length == 0) return false;
            var _retn = [];
            for (var r in retn) {
                if (par[i][0] == "#") _retn.push(document.getElementById(par[i].replace("#", "")));
                else if (par[i][0] == ".") {
                    var tag = retn[r].getElementsByTagName('*');
                    for (var j = 0; j < tag.length; j++) {
                        var cln = tag[j].className;
                        if (cln && cln.search(new RegExp("\\b" + par[i].replace(".", "") + "\\b")) != -1) { _retn.push(tag[j]); }
                    }
                }
                else { var tag = retn[r].getElementsByTagName(par[i]); for (var j = 0; j < tag.length; j++) { _retn.push(tag[j]) } }
            }
            retn = _retn;
        }

        return retn.length == 0 || retn[0] == parEle ? false : retn;
    }
    // 控制是否点击
    LuckDraw.prototype._init = function (opts) {
        if (this.flag) {
            // 扩展defaultOPts对象
            for (var key in opts) {
                this.defaultOPts[key] = opts[key];
            }
            this._resetVariable();
            this._paramTesting();
            this.flag = !this.flag;
            this._running();
        } else {
            console.log('正在抽奖，请稍后点击！');
        }
    }
    win.LuckDraw = LuckDraw;
})(window)