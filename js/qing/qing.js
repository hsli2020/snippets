/**
 * @author      : 马蹄疾
 * @date        : 2018-03-23
 * @version     : v1.0
 * @description : 一个UI组件库
 * @repository  : https://github.com/veedrin/qing
 * @license     : MIT
 */


////////// * DatePicker和TimePicker公共父类 * //////////

class TimeCommon {
    constructor() {}

    twoDigitsFormat(num) {
        if (String(num).length === 1 && num < 10) {
            num = `0${num}`;
        }
        return num;
    }
}

////////// * 日期选择器组件类 * //////////

class DatePicker extends TimeCommon {
    constructor({
        id = 'date-picker',
        yearRange = [1970, 2050],
        lang = 'zh',
        callback = () => {},
    } = {}) {
        // 继承父类的this对象
        super();
        this.$mount = document.querySelector(`#${id}`);
        // 年份选项范围
        this.yearRange = yearRange;
        this.lang = lang;
        this.callback = callback;
        // 用户选中的年月日，初始是当前年月日
        [this.Y, this.M, this.D] = this.nowDate();
        // 上一次选中的日期
        this.oldD = this.D;
        // 面板选择年份时的锚点年份
        this.anchor = this.Y;
        this.init();
    }

    init() {
        // 参数检查
        this.verifyOptions();
        this.render();
        this.$trough = this.$mount.querySelector('.trough');
        this.$view = this.$trough.querySelector('.view');
        this.$curtain = this.$mount.querySelector('.curtain');
        this.renderDay();
        this.$troughEvent();
    }

    render() {
        const tpl = `
            <div class="qing qing-date-picker">
                <div class="trough">
                    <div class="view">${this.dateFormat(this.Y, this.M, this.D)}</div>
                    <span class="arrow">^</span>
                </div>
                <div class="curtain"></div>
            </div>
        `;
        this.$mount.innerHTML = tpl;
    }

    renderDay() {
        const [left, right] = this.yearRange;
        let tpl = `
            <div class="bar">
                <div class="bar-item">
                    <span class="${this.Y > left ? 'angle year-prev' : 'angle disabled'}">◀</span>
                    <span class="year-pop"></span>
                    <span class="${this.Y < right ? 'angle year-next' : 'angle disabled'}">▶</span>
                </div>
                <div class="bar-item">
                    <span class="${this.Y === left && this.M === 1 ? 'angle disabled' : 'angle month-prev'}">◀</span>
                    <span class="month-pop"></span>
                    <span class="${this.Y === right && this.M === 12 ? 'angle disabled' : 'angle month-next'}">▶</span>
                </div>
            </div>
            <div class="week">
                <span class="week-item">一</span>
                <span class="week-item">二</span>
                <span class="week-item">三</span>
                <span class="week-item">四</span>
                <span class="week-item">五</span>
                <span class="week-item">六</span>
                <span class="week-item">日</span>
            </div>
        `;
        const daysCountThisMonth = this.daysCountThisMonth();
        const daysCountLastMonth = this.daysCountThisMonth(-1);
        // weekie指的是星期几（自创）
        const weekieFirstDay = this.weekieOfSomedayThisMonth(1);
        const weekieLastDay = this.weekieOfSomedayThisMonth(daysCountThisMonth);
        // 当前年月日
        const [Y, M, D] = this.nowDate();
        tpl += '<div class="board">';
        if (weekieFirstDay > 1) {
            for (let i = daysCountLastMonth - weekieFirstDay + 2; i <= daysCountLastMonth; i++) {
                tpl += `<span class="day-disable">${i}</span>`;
            }
        }
        for (let i = 1; i <= daysCountThisMonth; i++) {
            if (this.Y === Y && this.M === M && i === D && i !== this.D) {
                tpl += `<span class="day today">${i}</span>`;
            } else if (this.Y === Y && this.M === M && i === D && i === this.D) {
                tpl += `<span class="day today active">${i}</span>`;
            } else if (i === this.D) {
                tpl += `<span class="day active">${i}</span>`;
            } else {
                tpl += `<span class="day">${i}</span>`;
            }
        }
        if (weekieLastDay < 7) {
            for (let i = 1; i <= 7 - weekieLastDay; i++) {
                tpl += `<span class="day-disable">${i}</span>`;
            }
        }
        tpl += '</div>'
        tpl += `
            <div class="control">
                <button class="today">${this.lang === 'en' ? 'Today' : '今天'}</button>
                <button class="close">${this.lang === 'en' ? 'Close' : '关闭'}</button>
            </div>
        `;
        this.$curtain.innerHTML = tpl;
        this.$bar = this.$curtain.querySelector('.bar');
        this.$yearPop = this.$bar.querySelector('.year-pop');
        this.$monthPop = this.$bar.querySelector('.month-pop');
        this.$dayEvent();
        this.$barEvent();
        this.$controlEvent();
        // 中英文配置
        this.langConfig();
    }

    $troughEvent() {
        const self = this;
        const curtainCL = this.$curtain.classList;
        const arrowCL = this.$trough.querySelector('.arrow').classList;
        this.$trough.addEventListener('click', function(event) {
            event.stopPropagation();
            arrowCL.toggle('active');
            if (curtainCL.contains('active')) {
                curtainCL.remove('active');
            } else {
                curtainCL.add('active');
                // 打开curtain重新渲染
                self.renderDay();
            }
        });
        document.addEventListener('click', function() {
            if (curtainCL.contains('active')) {
                arrowCL.remove('active');
                curtainCL.remove('active');
            }
        });
    }

    $dayEvent() {
        const self = this;
        const $days = this.$curtain.querySelectorAll('.board .day');
        for (let i = 0; i < $days.length; i++) {
            const $day = $days[i];
            const CL = $day.classList;
            $day.addEventListener('click', function(event) {
                event.stopPropagation();
                $days[self.oldD - 1].classList.remove('active');
                // 切换class
                CL.toggle('active');
                self.D = Number.parseInt(this.innerHTML);
                // 重置oldD
                self.oldD = self.D;
                const format = self.dateFormat(self.Y, self.M, this.innerHTML);
                self.$view.innerHTML = format;
                // 回调
                self.callback(format);
                // 稍微延迟关闭curtain
                setTimeout(() => {
                    self.$trough.click();
                }, 100);
            });
        }
    }

    $barEvent() {
        const self = this;
        const $yearPrev = this.$bar.querySelector('.year-prev');
        const $yearNext = this.$bar.querySelector('.year-next');
        const $monthPrev = this.$bar.querySelector('.month-prev');
        const $monthNext = this.$bar.querySelector('.month-next');
        // 初始内容
        this.$yearPop.innerHTML = this.Y;
        this.$monthPop.innerHTML = this.twoDigitsFormat(this.M);
        // 置灰时获取不到元素
        if ($yearPrev) {
            $yearPrev.addEventListener('click', function(event) {
                event.stopPropagation();
                self.Y--;
                self.yearAndMonthChange('year');
            });
        }
        // 置灰时获取不到元素
        if ($yearNext) {
            $yearNext.addEventListener('click', function(event) {
                event.stopPropagation();
                self.Y++;
                self.yearAndMonthChange('year');
            });
        }
        this.$yearPop.addEventListener('click', function(event) {
            event.stopPropagation();
            self.renderYear();
        });
        // 置灰时获取不到元素
        if ($monthPrev) {
            $monthPrev.addEventListener('click', function(event) {
                event.stopPropagation();
                self.M--;
                if (self.M > 0) {
                    self.yearAndMonthChange('month');
                } else {
                    self.M = 12;
                    self.Y--;
                    self.yearAndMonthChange('both');
                }
            });
        }
        // 置灰时获取不到元素
        if ($monthNext) {
            $monthNext.addEventListener('click', function(event) {
                event.stopPropagation();
                self.M++;
                if (self.M < 13) {
                    self.yearAndMonthChange('month');
                } else {
                    self.M = 1;
                    self.Y++;
                    self.yearAndMonthChange('both');
                }
            });
        }
        this.$monthPop.addEventListener('click', function(event) {
            event.stopPropagation();
            self.renderMonth();
        });
    }

    renderYear() {
        const [left, right] = this.yearRange;
        const start = this.anchor - 4 > left ? this.anchor - 4 : left;
        const end = this.anchor + 7 < right ? this.anchor + 7 : right;
        const [Y, , ] = this.nowDate();
        let tpl = `<div class="title">${this.lang === 'en' ? 'Choose a Year' : '选择年份'}</div>`;
        if (this.anchor - 4 > left) {
            tpl += '<div class="prev">◀</div>';
        } else {
            tpl += '<div class="prev-disabled">◀</div>';
        }
        tpl += '<div class="year-wrap">';
        for (let i = start; i <= end; i++) {
            if (i !== Y) {
                tpl += `<span class="year">${i}</span>`;
            } else {
                tpl += `<span class="year thisyear">${i}</span>`;
            }
        }
        tpl += '</div>';
        if (this.anchor + 7 < right) {
            tpl += '<div class="next">▶</div>';
        } else {
            tpl += '<div class="next-disabled">▶</div>';
        }
        this.$curtain.innerHTML = tpl;
        this.$yearEvent();
        this.$yearEndEvent();
    }

    $yearEvent() {
        const self = this;
        const $years = this.$curtain.querySelectorAll('.year');
        for (let i = 0; i < $years.length; i++) {
            $years[i].addEventListener('click', function(event) {
                event.stopPropagation();
                self.Y = Number.parseInt(this.innerHTML);
                self.yearAndMonthChange('year');
            });
        }
    }

    $yearEndEvent() {
        const self = this;
        const $prev = this.$curtain.querySelector('.prev');
        const $next = this.$curtain.querySelector('.next');
        // 置灰时获取不到元素
        if ($prev) {
            $prev.addEventListener('click', function(event) {
                event.stopPropagation();
                self.anchor -= 12;
                self.renderYear();
            });
        }
        // 置灰时获取不到元素
        if ($next) {
            $next.addEventListener('click', function(event) {
                event.stopPropagation();
                self.anchor += 12;
                self.renderYear();
            });
        }
    }

    renderMonth() {
        const [, M, ] = this.nowDate();
        let tpl = `<div class="title">${this.lang === 'en' ? 'Choose a Month' : '选择月份'}</div>`;
        for (let i = 1; i <= 12; i++) {
            if (i !== M) {
                tpl += `<span class="month">${this.twoDigitsFormat(i)}</span>`;
            } else {
                tpl += `<span class="month thismonth">${this.twoDigitsFormat(i)}</span>`;
            }
        }
        this.$curtain.innerHTML = tpl;
        this.$monthEvent();
    }

    $monthEvent() {
        const self = this;
        const $months = this.$curtain.querySelectorAll('.month');
        for (let i = 0; i < $months.length; i++) {
            $months[i].addEventListener('click', function(event) {
                event.stopPropagation();
                self.M = Number.parseInt(this.innerHTML);
                self.yearAndMonthChange('month');
            });
        }
    }

    yearAndMonthChange(code) {
        // 当前年月日
        const [Y, M, D] = this.nowDate();
        if (this.Y !== Y || this.M !== M) {
            this.D = 1;
        } else {
            this.D = D;
        }
        // 重置oldD
        this.oldD = this.D;
        switch (code) {
            case 'year':
                this.$yearPop.innerHTML = this.Y;
                break;
            case 'month':
                this.$monthPop.innerHTML = this.twoDigitsFormat(this.M);
                break;
            case 'both':
                this.$yearPop.innerHTML = this.Y;
                this.$monthPop.innerHTML = this.twoDigitsFormat(this.M);
                break;
        }
        this.$view.innerHTML = this.dateFormat(this.Y, this.M, this.D);
        this.renderDay();
    }

    $controlEvent() {
        const self = this;
        const $today = this.$curtain.querySelector('.control .today');
        const $close = this.$curtain.querySelector('.control .close');
        $today.addEventListener('click', function(event) {
            event.stopPropagation();
            [self.Y, self.M, self.D] = self.nowDate();
            // 重置oldD
            self.oldD = self.D;
            self.$view.innerHTML = self.dateFormat(self.Y, self.M, self.D);
            self.renderDay();
        });
        $close.addEventListener('click', function(event) {
            self.$trough.click();
        });
    }

    langConfig() {
        if (this.lang === 'en') {
            const dict = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'];
            const $weeks = this.$curtain.querySelectorAll('.week-item');
            for (let i = 0; i < $weeks.length; i++) {
                $weeks[i].innerHTML = dict[i];
            }
        }
    }

    nowDate() {
        const date = new Date();
        return [date.getFullYear(), date.getMonth() + 1, date.getDate()];
    }

    daysCountThisMonth(num = 0) {
        return new Date(this.Y, this.M + num, 0).getDate();
    }

    weekieOfSomedayThisMonth(day) {
        let weekie = new Date(this.Y, this.M - 1, day).getDay();
        if (weekie === 0) {
            weekie = 7;
        }
        return weekie;
    }

    dateFormat(year, month, day) {
        month = this.twoDigitsFormat(month);
        day = this.twoDigitsFormat(day);
        return `${year}-${month}-${day}`;
    }

    verifyOptions() {
        // $mount
        if (!this.$mount) {
            throw new Error('[Qing error]: 日期选择器组件无法找到挂载点');
        }
        // yearRange
        if (Object.prototype.toString.call(this.yearRange) !== '[object Array]') {
            this.yearRange = [1970, 2050];
            console.warn('[Qing warn]: 日期选择器的yearRange必须是数组');
        } else if (typeof this.yearRange[0] !== 'number' || typeof this.yearRange[1] !== 'number') {
            this.yearRange = [1970, 2050];
            console.warn('[Qing warn]: 日期选择器的yearRange的年份必须是数字');
        } else if (this.yearRange[0] > this.Y || this.yearRange[1] < this.Y) {
            this.yearRange = [1970, 2050];
            console.warn('[Qing warn]: 日期选择器的yearRange的范围必须包含当年');
        }
        // callback
        if (typeof this.callback !== 'function') {
            this.callback = () => {};
            console.warn('[Qing warn]: 日期选择器的callback必须是函数');
        }
    }
}

////////// * 时间选择器组件类 * //////////

class TimePicker extends TimeCommon {
    constructor({
        id = 'time-picker',
        lang = 'zh',
        callback = () => {},
    } = {}) {
        // 继承父类的this对象
        super();
        this.$mount = document.querySelector(`#${id}`);
        this.lang = lang;
        this.callback = callback;
        // 用户选中的时间，初始是当前时间，秒数为0
        [this.H, this.M, ] = this.nowTime();
        this.S = 0;
        // 用户上一次选中的时间
        [this.oldH, this.oldM, this.oldS] = [this.H, this.M, this.S];
        // 上一次被禁用的按钮
        this.oldDisabled;
        this.init();
    }

    init() {
        // 参数检查
        this.verifyOptions();
        this.render();
        this.$trough = this.$mount.querySelector('.trough');
        this.$view = this.$trough.querySelector('.view');
        this.$curtain = this.$mount.querySelector('.curtain');
        this.$time = this.$curtain.querySelector('.time');
        this.$troughEvent();
        this.$selectEvent();
        this.$controlEvent();
        // 中英文配置
        this.langConfig();
    }

    render() {
        const tpl = `
            <div class="qing qing-time-picker">
                <div class="trough">
                    <div class="view">${this.timeFormat(this.H, this.M, this.S)}</div>
                    <span class="arrow">^</span>
                </div>
                <div class="curtain">
                    <div class="select">
                        <button class="select-item select-h">时钟</button>
                        <button class="select-item select-m">分钟</button>
                        <button class="select-item select-s">秒钟</button>
                    </div>
                    <div class="time"></div>
                    <div class="control">
                        <button class="now">${this.lang === 'en' ? 'Now' : '现在'}</button>
                        <button class="ok">${this.lang === 'en' ? 'OK' : '确定'}</button>
                    </div>
                </div>
            </div>
        `;
        this.$mount.innerHTML = tpl;
    }

    $troughEvent() {
        const self = this;
        const curtainCL = this.$curtain.classList;
        const arrowCL = this.$trough.querySelector('.arrow').classList;
        this.$trough.addEventListener('click', function(event) {
            event.stopPropagation();
            arrowCL.toggle('active');
            if (curtainCL.contains('active')) {
                curtainCL.remove('active');
            } else {
                curtainCL.add('active');
                // 打开curtain，回到时钟界面
                if (self.oldDisabled === self.$selectH) {
                    // 如果已经disabled，则去掉disabled才可以点击
                    self.oldDisabled.removeAttribute('disabled');
                    self.oldDisabled.click();
                } else {
                    self.$selectH.click();
                    self.oldDisabled = self.$selectH;
                }
            }
        });
        document.addEventListener('click', function() {
            if (curtainCL.contains('active')) {
                arrowCL.remove('active');
                curtainCL.remove('active');
            }
        });
    }

    $selectEvent() {
        const self = this;
        this.$selectH = this.$curtain.querySelector('.select-h');
        this.$selectM = this.$curtain.querySelector('.select-m');
        this.$selectS = this.$curtain.querySelector('.select-s');
        this.$selectH.addEventListener('click', function(event) {
            event.stopPropagation();
            self.ableAndDisableEvent(this);
            self.renderHour();
            self.$hourEvent();
        });
        this.$selectM.addEventListener('click', function(event) {
            event.stopPropagation();
            self.ableAndDisableEvent(this);
            self.renderMinute();
            self.$minuteEvent();
        });
        this.$selectS.addEventListener('click', function(event) {
            event.stopPropagation();
            self.ableAndDisableEvent(this);
            self.renderSecond();
            self.$secondEvent();
        });
        // 初始触发
        this.$selectH.click();
    }

    ableAndDisableEvent(ableNode) {
        if (this.oldDisabled) {
            this.oldDisabled.removeAttribute('disabled');
        }
        ableNode.setAttribute('disabled', '');
        this.oldDisabled = ableNode;
    }

    renderHour() {
        let tpl = '';
        for (let i = 0; i < 24; i++) {
            if (this.H !== i) {
                tpl += `<span class="hour">${this.twoDigitsFormat(i)}</span>`;
            } else {
                tpl += `<span class="hour active">${this.twoDigitsFormat(i)}</span>`;
            }
        }
        this.$time.innerHTML = tpl;
    }

    $hourEvent() {
        const self = this;
        const $hours = this.$time.querySelectorAll('.hour');
        for (let i = 0; i < $hours.length; i++) {
            const $hour = $hours[i];
            const CL = $hour.classList;
            $hour.addEventListener('click', function(event) {
                event.stopPropagation();
                $hours[self.oldH].classList.remove('active');
                CL.add('active');
                const hour = Number.parseInt(this.innerHTML);
                self.oldH = hour;
                self.timeChangeEvent(hour, 0, 0);
            });
        }
    }

    renderMinute() {
        let tpl = '';
        for (let i = 0; i < 60; i++) {
            if (this.M !== i) {
                tpl += `<span class="minute">${this.twoDigitsFormat(i)}</span>`;
            } else {
                tpl += `<span class="minute active">${this.twoDigitsFormat(i)}</span>`;
            }
        }
        this.$time.innerHTML = tpl;
    }

    $minuteEvent() {
        const self = this;
        const $minutes = this.$time.querySelectorAll('.minute');
        for (let i = 0; i < $minutes.length; i++) {
            const $minute = $minutes[i];
            const CL = $minute.classList;
            $minute.addEventListener('click', function(event) {
                event.stopPropagation();
                $minutes[self.oldM].classList.remove('active');
                CL.add('active');
                const minute = Number.parseInt(this.innerHTML);
                self.oldM = minute;
                self.timeChangeEvent(0, minute, 0);
            });
        }
    }

    renderSecond() {
        let tpl = '';
        for (let i = 0; i < 60; i++) {
            if (this.S !== i) {
                tpl += `<span class="second">${this.twoDigitsFormat(i)}</span>`;
            } else {
                tpl += `<span class="second active">${this.twoDigitsFormat(i)}</span>`;
            }
        }
        this.$time.innerHTML = tpl;
    }

    $secondEvent() {
        const self = this;
        const $seconds = this.$time.querySelectorAll('.second');
        for (let i = 0; i < $seconds.length; i++) {
            const $second = $seconds[i];
            const CL = $second.classList;
            $second.addEventListener('click', function(event) {
                event.stopPropagation();
                $seconds[self.oldS].classList.remove('active');
                CL.add('active');
                const second = Number.parseInt(this.innerHTML);
                self.oldS = second;
                self.timeChangeEvent(0, 0, second);
            });
        }
    }

    timeChangeEvent(hour, minute, second) {
        hour ? this.H = hour : '';
        minute ? this.M = minute : '';
        second ? this.S = second : '';
        this.$view.innerHTML = this.timeFormat(this.H, this.M, this.S);
    }

    $controlEvent() {
        const self = this;
        const $now = this.$curtain.querySelector('.control .now');
        const $ok = this.$curtain.querySelector('.control .ok');
        $now.addEventListener('click', function(event) {
            event.stopPropagation();
            [self.H, self.M, self.S] = self.nowTime();
            self.$view.innerHTML = self.timeFormat(self.H, self.M, self.S);
            // 缓存选择的时间
            [self.oldH, self.oldM, self.oldS] = [self.H, self.M, self.S];
            // 更新时间面板，回到时钟界面
            if (self.oldDisabled === self.$selectH) {
                // 如果已经disabled，则去掉disabled才可以点击
                self.oldDisabled.removeAttribute('disabled');
                self.oldDisabled.click();
            } else {
                self.$selectH.click();
                self.oldDisabled = self.$selectH;
            }
        });
        $ok.addEventListener('click', function(event) {
            event.stopPropagation();
            // 回调
            self.callback(self.timeFormat(self.H, self.M, self.S));
            // 稍微延迟关闭curtain
            setTimeout(() => {
                self.$trough.click();
            }, 100);
        });
    }

    langConfig() {
        if (this.lang === 'en') {
            [this.$selectH.innerHTML, this.$selectM.innerHTML, this.$selectS.innerHTML] = ['Hour', 'Minute', 'Second'];
        }
    }

    nowTime() {
        const date = new Date();
        return [date.getHours(), date.getMinutes(), date.getSeconds()];
    }

    timeFormat(hour = 0, minute = 0, second = 0) {
        hour = this.twoDigitsFormat(hour);
        minute = this.twoDigitsFormat(minute);
        second = this.twoDigitsFormat(second);
        return `${hour} : ${minute} : ${second}`;
    }

    verifyOptions() {
        // $mount
        if (!this.$mount) {
            throw new Error('[Qing error]: 时间选择器组件无法找到挂载点');
        }
        // callback
        if (typeof this.callback !== 'function') {
            this.callback = () => {};
            console.warn('[Qing warn]: 时间选择器的callback必须是函数');
        }
    }
}

////////// * 分页组件类 * //////////

class Paginator {
    constructor({
        id = 'paginator',
        pageCount = 1,
        showSizeChanger = false,
        pageSizeOptions = [10, 20, 30],
        pageSize = 10,
        total,
        showQuickJumper = false,
        lang = 'zh',
        callback = () => {},
    } = {}) {
        this.$mount = document.querySelector(`#${id}`);
        // 页数
        this.pageCount = pageCount;
        // 是否显示调整每页条数的下拉框
        this.showSizeChanger = showSizeChanger;
        // 每页条数选项
        this.pageSizeOptions = pageSizeOptions;
        // 每页条数
        this.pageSize = pageSize;
        // 总条数
        this.total = total ? total : this.pageCount * this.pageSize;
        // 是否显示快速跳转至某页
        this.showQuickJumper = showQuickJumper;
        this.lang = lang;
        this.callback = callback;
        // 分页数据模型
        this.model = [];
        // 当前页
        this.position = 1;
        this.init();
    }

    init() {
        // 参数检查
        this.verifyOptions();
        this.render();
        this.$prev = this.$mount.querySelector('.prev');
        this.$next = this.$mount.querySelector('.next');
        this.$bar = this.$mount.querySelector('.bar');
        this.actPerEvent();
        this.$endEvent();
        if (this.showSizeChanger) {
            this.$combobox = this.$mount.querySelector('.combobox');
            this.$panel = this.$combobox.querySelector('.panel');
            this.$arrow = this.$combobox.querySelector('.arrow');
            this.$comboboxEvent();
            this.$optionEvent();
        }
        if (this.showQuickJumper) {
            this.$jumpEvent();
        }
    }

    actPerEvent() {
        this.buildModel();
        this.renderBar();
        this.$pageEvent();
    }

    render() {
        let tpl = '<div class="qing qing-paginator">';
        tpl += `
            <div class="square end prev">﹤</div>
            <div class="bar"></div>
            <div class="square end next">﹥</div>
        `;
        if (this.showSizeChanger === true) {
            const perPage = this.lang === 'en' ? 'page' : '页';
            // 下拉列表的模板
            let optionTpl = '';
            for (const item of this.pageSizeOptions) {
                optionTpl += `<div class="option">${item}&nbsp;/&nbsp;${perPage}</div>`;
            }
            tpl += `
                <div class="combobox">
                    <div class="show">
                        <span class="size">${this.pageSize}</span>
                        <span>/&nbsp;${perPage}</span>
                    </div>
                    <span class="arrow">^</span>
                    <div class="panel">${optionTpl}</div>
                </div>
            `;
        }
        if (this.showQuickJumper === true) {
            tpl += `
                <div class="jumper">
                    <span class="goto">${this.lang === 'en' ? 'Goto' : '前往'}</span>
                    <input class="jump" type="text">
                </div>
            `;
        }
        tpl += '</div>';
        this.$mount.innerHTML = tpl;
    }

    buildModel() {
        // 每次重新初始化
        this.model = [];
        const c = this.pageCount, p = this.position;
        // 首页和尾页必须展示
        // 如果有省略号则首尾只展示一条，当前页前后各展示两条共五条，一边没有空间则叠加到另一边
        // 首尾页与当前页五条可以重合
        // 跨度大于等于两条才出现省略号，省略号用0表示
        if (c < 8) {
            for (let i = 1; i <= c; i++) {
                this.model.push(i);
            }
        } else {
            if (p < 4) {
                for (let i = 1; i <= 5; i++) {
                    this.model.push(i);
                }
                this.model.push(0, c);
            } else if (p < 6) {
                for (let i = 1; i <= p + 2; i++) {
                    this.model.push(i);
                }
                this.model.push(0, c);
            } else {
                if (p < c - 4) {
                    this.model.push(1, 0);
                    for (let i = p - 2; i <= p + 2; i++) {
                        this.model.push(i);
                    }
                    this.model.push(0, c);
                } else if (p < c - 1) {
                    this.model.push(1, 0);
                    for (let i = p - 2; i <= c; i++) {
                        this.model.push(i);
                    }
                } else {
                    this.model.push(1, 0);
                    for (let i = c - 4; i <= c; i++) {
                        this.model.push(i);
                    }
                }
            }
        }
    }

    renderBar() {
        let tpl = '';
        for (const item of this.model) {
            if (item > 0) {
                if (this.position !== item) {
                    tpl += `<div class="square page">${item}</div>`;
                } else {
                    tpl += `<div class="square page active">${item}</div>`;
                }
            } else {
                tpl += '<div class="square gap">···</div>';
            }
        }
        this.$bar.innerHTML = tpl;
        // 控制prev和next是否置灰
        // 如果只有一页，则this.pageCount === 1
        if (this.pageCount === 1) {
            this.$prev.classList.add('disabled');
            this.$next.classList.add('disabled');
            return;
        }
        if (this.position === 1) {
            this.$prev.classList.add('disabled');
            this.$next.classList.remove('disabled');
        } else if (this.position === this.pageCount) {
            this.$next.classList.add('disabled');
            this.$prev.classList.remove('disabled');
        } else {
            this.$prev.classList.remove('disabled');
            this.$next.classList.remove('disabled');
        }
    }

    $pageEvent() {
        const self = this;
        const $pages = this.$mount.querySelectorAll('.page');
        for (const $page of $pages) {
            $page.addEventListener('click', function(event) {
                event.stopPropagation();
                self.position = Number.parseInt(this.innerHTML);
                // 回调
                self.callback(self.position, self.pageSize);
                // 重新渲染
                self.actPerEvent();
            });
        }
    }

    $endEvent() {
        const self = this;
        this.$prev.addEventListener('click', function(event) {
            event.stopPropagation();
            if (self.position === 1) {
                return;
            }
            self.position--;
            // 回调
            self.callback(self.position, self.pageSize);
            // 重新渲染
            self.actPerEvent();
        });
        this.$next.addEventListener('click', function(event) {
            event.stopPropagation();
            if (self.position === self.pageCount) {
                return;
            }
            self.position++;
            // 回调
            self.callback(self.position, self.pageSize);
            // 重新渲染
            self.actPerEvent();
        });
    }

    $comboboxEvent() {
        const panelCL = this.$panel.classList;
        const arrowCL = this.$arrow.classList;
        this.$combobox.addEventListener('click', function(event) {
            event.stopPropagation();
            panelCL.toggle('active');
            arrowCL.toggle('active');
        });
        // 点击页面任何地方关闭combobox
        document.addEventListener('click', function() {
            panelCL.contains('active') ? panelCL.remove('active') : '';
            arrowCL.contains('active') ? arrowCL.remove('active') : '';
        });
    }

    $optionEvent() {
        const self = this;
        const $options = this.$panel.querySelectorAll('.option');
        const $size = this.$combobox.querySelector('.size');
        // 如果pageSize是pageSizeOptions当中的一项，该项所在的$option添加active
        let index = this.pageSizeOptions.indexOf(this.pageSize);
        if (index > -1) {
            $options[index].classList.add('active');
        } else if (index === -1) {
            index = 0;
        }
        const panelCL = this.$panel.classList;
        const arrowCL = this.$arrow.classList;
        for (let i = 0; i < $options.length; i++) {
            const $option = $options[i];
            const optionCL = $option.classList;
            $option.addEventListener('click', function(event) {
                event.stopPropagation();
                // 收起combobox和arrow；如果不阻止冒泡则可省略
                panelCL.remove('active');
                arrowCL.remove('active');
                // 处理$option的active
                optionCL.add('active');
                $options[index].classList.remove('active');
                index = i;
                // 选中的每页条数
                const num = Number.parseInt(this.innerHTML.split('/')[0].trim());
                self.pageSize = num;
                $size.innerHTML = num;
                // 如果每页显示条数变化导致总页数比当前页小，则当前页变成最后一页
                const pageCount = Math.ceil(self.total / self.pageSize);
                if (pageCount < self.position) {
                    self.position = pageCount;
                }
                self.pageCount = pageCount;
                // 回调
                self.callback(self.position, self.pageSize);
                // 重新渲染
                self.actPerEvent();
            });
        }
    }

    $jumpEvent() {
        const self = this;
        const $jump = this.$mount.querySelector('.jump');
        $jump.addEventListener('keyup', function(event) {
            event.stopPropagation();
            if (event.keyCode !== 13) {
                return;
            }
            const value = this.value;
            this.value = '';
            if (Number.isInteger(value) && value > 0) {
                // 如果value大于页数，则前往最后一页
                value <= self.pageCount ? self.position = value : self.position = self.pageCount;
                // 回调
                self.callback(self.position, self.pageSize);
                // 重新渲染
                self.actPerEvent();
            }
        });
    }

    verifyOptions() {
        // $mount
        if (!this.$mount) {
            throw new Error('[Qing error]: 分页组件无法找到挂载点');
        }
        // pageCount
        if (typeof this.pageCount !== 'number') {
            this.pageCount = 1;
            console.warn('[Qing warn]: 分页组件的pageCount必须是数字');
        }
        // pageSizeOptions
        if (Object.prototype.toString.call(this.pageSizeOptions) !== '[object Array]') {
            this.pageSizeOptions = [10, 20, 30];
            console.warn('[Qing warn]: 分页组件的pageSizeOptions必须是数组');
        } else {
            for (let i = 0; i < this.pageSizeOptions.length; i++) {
                if (typeof this.pageSizeOptions[i] !== 'number') {
                    this.pageSizeOptions = [10, 20, 30];
                    console.warn('[Qing warn]: 分页组件的pageSizeOptions数组的项必须是数字');
                    break;
                }
            }
        }
        // pageSize
        if (typeof this.pageSize !== 'number') {
            this.pageSize = 10;
            console.warn('[Qing warn]: 分页组件的pageSize必须是数字');
        }
        // total
        if (typeof this.total !== 'number') {
            this.total = 10;
            console.warn('[Qing warn]: 分页组件的total必须是数字');
        } else if (this.total > this.pageCount * this.pageSize) {
            this.total = this.pageCount * this.pageSize;
            console.warn('[Qing warn]: 分页组件的total不能超过pageCount和pageSize的乘积');
        }
        // callback
        if (typeof this.callback !== 'function') {
            this.callback = () => {};
            console.warn('[Qing warn]: 分页组件的callback必须是函数');
        }
    }
}

////////// * 树组件类 * //////////

class Tree {
    constructor({
        id = 'tree',
        data = [],
        checkable = true,
        indent = 40,
        expand = 'none',
        callback = () => {},
    } = {}) {
        this.$mount = document.querySelector(`#${id}`);
        this.data = data;
        // 是否显示checkbox
        this.checkable = checkable;
        // 每一级的缩进距离
        this.indent = indent;
        // 初始伸展方式
        this.expand = expand;
        this.callback = callback;
        // 与data结构相同的cb树
        this.cbTree = [];
        this.init();
    }

    init() {
        // 参数检查
        this.verifyOptions();
        this.render();
        if (this.checkable) {
            // $mount.firstElementChild是根元素
            this.buildCbTree(this.$mount.firstElementChild, {sub: this.cbTree});
            this.$checkboxEvent(this.cbTree);
        }
        this.$fruitEvent();
    }

    render() {
        const tpl = `
            <div class="qing qing-tree">
                ${this.renderTrunk(this.data)}
            </div>
        `;
        this.$mount.innerHTML = tpl;
    }

    renderTrunk(data) {
        const inner = data !== this.data;
        const marginLeft = `${inner ? `style="margin-left: ${this.indent}px;"` : ''}`;
        let tpl = '';
        for (let i = 0; i < data.length; i++) {
            const item = data[i];
            tpl += `<div class="trunk" ${marginLeft}>`;
            // 判断展开或闭合
            let arrowTpl, subTpl;
            if (this.expand === 'all') {
                arrowTpl = '<span class="arrow active"></span>';
                subTpl = '<div class="sub" style="height: auto;">';
            } else if (this.expand === 'first') {
                const boo = !inner && i === 0;
                arrowTpl = boo ? '<span class="arrow active"></span>' : '<span class="arrow"></span>';
                subTpl = `<div class="sub" ${boo ? 'style="height: auto;"' : ''}>`;
            } else {
                arrowTpl = '<span class="arrow"></span>';
                subTpl = '<div class="sub">';
            }
            tpl += `
                <div class="fruit">
                    ${item.sub ? arrowTpl : '<span class="blank"></span>'}
                    ${this.checkable ? '<span class="cb"></span>' : ''}
                    <span class="label">${item.label}</span>
                </div>
            `;
            if (item.sub) {
                tpl += subTpl;
                tpl += this.renderTrunk(item.sub);
                tpl += '</div>';
            }
            tpl += '</div>';
        }
        return tpl;
    }

    buildCbTree(fatherCb, item) {
        const childCbs = fatherCb.children;
        for (let i = 0; i < childCbs.length; i++) {
            const fruit = childCbs[i].firstElementChild;
            const sub = fruit.nextElementSibling;
            const cb = fruit.firstElementChild.nextElementSibling;
            const checked = cb.classList.contains('checked');
            const queue = item.queue ? item.queue + String(i) : String(i);
            let obj = {cb, checked, queue};
            if (sub) {
                obj.sub = [];
                this.buildCbTree(sub, obj);
            }
            item.sub.push(obj);
        }
    }

    $fruitEvent() {
        const self = this;
        const $fruits = this.$mount.querySelectorAll('.fruit');
        for (const $fruit of $fruits) {
            const $sub = $fruit.nextElementSibling;
            if (!$sub) {
                continue;
            }
            const CL = $fruit.querySelector('.arrow').classList;
            $fruit.addEventListener('click', function(event) {
                event.stopPropagation();
                // sub动画
                self.subHeightToggle($sub);
                // arrow动画
                CL.toggle('active');
            });
        }
    }

    subHeightToggle($sub) {
        let h = $sub.getBoundingClientRect().height;
        if (h > 0) {
            // 从auto变成具体的值
            $sub.style.height = `${h}px`;
            setTimeout(() => {
                $sub.style.height = '0px';
            }, 0);
        } else {
            $sub.style.height = 'auto';
            h = $sub.getBoundingClientRect().height;
            $sub.style.height = '0px';
            setTimeout(() => {
                $sub.style.height = `${h}px`;
            }, 0);
            // 动画完成变成auto
            setTimeout(() => {
                $sub.style.height = 'auto';
            }, 200);
        }
    }

    $checkboxEvent(arr) {
        const self = this;
        for (const item of arr) {
            const $cb = item.cb;
            const $sub = item.sub;
            const queue = item.queue;
            $cb.addEventListener('click', function(event) {
                event.stopPropagation();
                const checked = !this.classList.contains('checked');
                // cb事件
                checked ? self.$cbEvent(item, 'all') : self.$cbEvent(item, 'none');
                // cb子代事件
                checked ? self.childCbsEvent($sub, 'all') : self.childCbsEvent($sub, 'none');
                // cb父代事件
                self.fatherCbsEvent(queue);
                // 根据cbTree更新数据
                self.updateData(self.cbTree, self.data);
                // 触发回调
                self.callback(self.data);
            });
            if ($sub) {
                this.$checkboxEvent($sub);
            }
        }
    }

    $cbEvent(item, action) {
        const $cb = item.cb;
        const CL = $cb.classList;
        switch (action) {
            case 'all':
                item.checked = true;
                CL.contains('somechecked') ? CL.remove('somechecked') : '';
                CL.add('checked');
                break;
            case 'some':
                item.checked = false;
                CL.contains('checked') ? CL.remove('checked') : '';
                CL.add('somechecked');
                break;
            case 'none':
                item.checked = false;
                CL.contains('somechecked') ? CL.remove('somechecked') : '';
                CL.contains('checked') ? CL.remove('checked') : '';
                break;
        }
    }

    childCbsEvent($sub, action) {
        if (!$sub) {
            return;
        }
        const self = this;
        function recursive($sub) {
            for (const $item of $sub) {
                self.$cbEvent($item, action);
                if ($item.sub) {
                    recursive($item.sub);
                }
            }
        }
        recursive($sub);
    }

    fatherCbsEvent(queue) {
        const fatherItem = this.findFatherItem(queue);
        if (!fatherItem) {
            return;
        }
        const $siblingCbs = this.findSiblingCbs(fatherItem.sub);
        let allFlag = true;
        let noneFlag = true;
        for (const $item of $siblingCbs) {
            const cl = $item.classList;
            if (!cl.contains('checked')) {
                allFlag = false;
            }
            if (cl.contains('checked') || cl.contains('somechecked')) {
                noneFlag = false;
            }
            // flag全都已经变化，退出循环
            if (!allFlag && !noneFlag) {
                break;
            }
        }
        if (allFlag) {
            this.$cbEvent(fatherItem, 'all');
        } else if (noneFlag) {
            this.$cbEvent(fatherItem, 'none');
        } else {
            this.$cbEvent(fatherItem, 'some');
        }
        this.fatherCbsEvent(fatherItem.queue);
    }

    findFatherItem(queue) {
        const n = queue.length - 1;
        // 顶级item没有父item
        if (n === 0) {
            return;
        }
        let fatherItem = this.cbTree;
        for (let i = 0; i < n; i++) {
            const char = queue.charAt(i);
            if (i < n - 1) {
                fatherItem = fatherItem[char].sub;
            } else {
                fatherItem = fatherItem[char];
            }
        }
        return fatherItem;
    }

    findSiblingCbs($sub) {
        let $siblingCbs = [];
        for (const $item of $sub) {
            $siblingCbs.push($item.cb);
        }
        return $siblingCbs;
    }

    updateData(tree, data) {
        for (let i = 0; i < tree.length; i++) {
            const ti = tree[i];
            const di = data[i];
            if (ti.checked) {
                di.checked = true;
            } else {
                di.checked = false;
            }
            if (ti.sub) {
                this.updateData(ti.sub, di.sub);
            }
        }
    }

    verifyOptions() {
        // $mount
        if (!this.$mount) {
            throw new Error('[Qing error]: 树组件无法找到挂载点');
        }
        // data
        if (Object.prototype.toString.call(this.data) !== '[object Array]') {
            this.data = [];
            console.warn('[Qing warn]: 树组件的data必须是数组');
        }
        // indent
        if (typeof this.indent !== 'number') {
            this.indent = 40;
            console.warn('[Qing warn]: 树组件的indent必须是数字');
        }
        // callback
        if (typeof this.callback !== 'function') {
            this.callback = () => {};
            console.warn('[Qing warn]: 树组件的callback必须是函数');
        }
    }
}

////////// * 级联选择器组件类 * //////////

class Cascader {
    constructor({
        id = 'cascader',
        data = [],
        searchable = false,
        debounce = 300,
        trigger = 'click',
        seperator = ' / ',
        callback = () => {},
    } = {}) {
        this.$mount = document.querySelector(`#${id}`);
        this.data = data;
        // 是否提供搜索框
        this.searchable = searchable;
        // 输入防抖
        this.debounce = debounce;
        // 显示级联的触发方式
        this.trigger = trigger;
        // 分隔符
        this.seperator = seperator;
        this.callback = callback;
        // 缓存用户选择的路径
        this.path = [];
        // queue的hash属性名，queue标记级联的层级
        this.queue = Date.now();
        this.eventType = this.trigger === 'hover' ? 'mouseenter' : 'click';
        // 缓存所有的路径
        this.pathPool = [];
        // 是否被搜索结果填充
        this.withResult = false;
        this.init();
    }

    init() {
        // 参数检查
        this.verifyOptions();
        // 给data添加queue
        this.buildQueue(this.data);
        this.render();
        this.$trough = this.$mount.querySelector('.trough');
        this.$view = this.$trough.querySelector('.view');
        this.$curtain = this.$mount.querySelector('.curtain');
        // 第一个board以及子元素
        this.$board = this.$curtain.querySelector('.board');
        this.$trunks = this.$board.querySelectorAll('.trunk');
        this.$subBoard = this.$board.querySelector('.board');
        this.$troughEvent();
        this.$rowEvent(this.$board);
        // 渲染搜索框，则添加事件
        if (this.searchable === true) {
            this.$search = this.$curtain.querySelector('.search');
            this.iterateAllPath();
            this.$searchEvent();
        }
    }

    buildQueue(data, queue = '') {
        for (let i = 0; i < data.length; i++) {
            const item = data[i];
            const sub = item.sub;
            const newQueue = `${queue}${i}`;
            item[this.queue] = newQueue;
            if (sub) {
                this.buildQueue(sub, newQueue);
            }
        }
    }

    render() {
        const tpl = `
            <div class="qing qing-cascader">
                <div class="trough">
                    <div class="view"></div>
                    <span class="arrow">^</span>
                </div>
                <div class="curtain">
                    <div class="board">
                        ${this.renderCascade(this.data)}
                    </div>
                    ${this.searchable === true ? '<div class="control"><input class="search" type="text"></div>' : ''}
                </div>
            </div>
        `;
        this.$mount.innerHTML = tpl;
    }

    renderCascade(data) {
        let tpl = '';
        for (const item of data) {
            if (item.sub) {
                tpl += `
                    <div class="row trunk" data-v="${item[this.queue]}">
                        <span class="label">${item.label}</span>
                        <span class="arrow">﹥</span>
                    </div>
                `;
            } else {
                tpl += `
                    <div class="row leaf" data-v="${item[this.queue]}">
                        <span class="label">${item.label}</span>
                    </div>
                `;
            }
        }
        tpl += '<div class="board"></div>';
        return tpl;
    }

    $troughEvent() {
        const self = this;
        const curtainCL = this.$curtain.classList;
        const arrowCL = this.$trough.querySelector('.view').classList;
        this.$trough.addEventListener('click', function(event) {
            event.stopPropagation();
            arrowCL.toggle('active');
            if (curtainCL.contains('active')) {
                curtainCL.remove('active');
            } else {
                curtainCL.add('active');
                // 再次打开curtain需要初始化
                if (self.withResult) {
                    // 清除搜索结果
                    self.clearResultAndRefill();
                } else {
                    // 清除级联
                    self.removeTrunkActive(self.$trunks);
                    self.$subBoard.innerHTML = '';
                }
            }
        });
        document.addEventListener('click', function() {
            if (curtainCL.contains('active')) {
                arrowCL.remove('active');
                curtainCL.remove('active');
            }
        });
    }

    $rowEvent($board) {
        const self = this;
        const $trunks = $board.querySelectorAll('.trunk');
        const $leafs = $board.querySelectorAll('.leaf');
        const $subBoard = $board.querySelector('.board');
        for (let i = 0; i < $trunks.length; i++) {
            const $trunk = $trunks[i];
            const v = $trunk.dataset.v;
            const label = $trunk.querySelector('.label').innerHTML;
            const CL = $trunk.classList;
            $trunk.addEventListener(this.eventType, function(event) {
                event.stopPropagation();
                // 遍历清除trunk的active
                self.removeTrunkActive($trunks);
                // 当前trunk变成active
                CL.add('active');
                // 构建路径
                self.buildPath(v.length, label, false);
                // 找到子数据
                const sub = self.findSubByQueue(self.data, v);
                // 填充子board
                $subBoard.innerHTML = self.renderCascade(sub);
                // 添加事件
                self.$rowEvent($subBoard);
            });
        }
        if (this.trigger === 'hover') {
            for (let i = 0; i < $leafs.length; i++) {
                $leafs[i].addEventListener('mouseenter', function() {
                    self.removeTrunkActive($trunks);
                    $subBoard.innerHTML = '';
                });
            }
            // 离开curtain
            this.$curtain.addEventListener('mouseleave', function() {
                self.removeTrunkActive(self.$trunks);
                self.$subBoard.innerHTML = '';
            });
        }
        // 树叶点击事件，结束选择
        this.endClickEvent($leafs);
    }

    endClickEvent($leafs) {
        const self = this;
        for (let i = 0; i < $leafs.length; i++) {
            const $leaf = $leafs[i];
            const length = $leaf.dataset.v.length;
            const label = $leaf.querySelector('.label').innerHTML;
            $leafs[i].addEventListener('click', function(event) {
                event.stopPropagation();
                // 渲染路径
                self.buildPath(length, label, true);
            });
        }
    }

    findSubByQueue(data, queue) {
        const n = Number.parseInt(queue.charAt(0));
        for (let i = 0; i < data.length; i++) {
            if (i === n) {
                if (queue.length > 1) {
                    return this.findSubByQueue(data[i].sub, queue.slice(1));
                } else {
                    return data[i].sub;
                }
            }
        }
    }

    buildPath(level, label, render) {
        if (this.path.length < level) {
            // 往下选择，直接push
            this.path.push(label);
        } else {
            // 退回选择，根据退回长度删除path元素，再push
            this.path = [...this.path.slice(0, level - 1), label];
        }
        if (render) {
            this.renderPath();
        }
    }

    renderPath() {
        const path = this.path.join(this.seperator);
        this.$view.innerHTML = path;
        // 回调
        this.callback(path);
        // 清空path容器
        this.path = [];
        this.$trough.click();
        this.removeTrunkActive(this.$trunks);
        this.$subBoard.innerHTML = '';
    }

    removeTrunkActive($trunks) {
        for (const $trunk of $trunks) {
            const CL = $trunk.classList;
            if (CL.contains('active')) {
                CL.remove('active');
                break;
            }
        }
    }

    iterateAllPath() {
        const self = this;
        let temp = [];
        const data = pathPush([...this.data]);
        function pathPush(data, arr = []) {
            for (const item of data) {
                item.path = [];
                // 将路径存入item中的数组
                item.path.push(...arr, item.label);
            }
            return data;
        }
        function recursive(data) {
            for (const item of data) {
                const sub = item.sub;
                if (sub) {
                    // 将下一层放入temp
                    temp.push(...pathPush(sub, item.path));
                } else {
                    // 没有下一层则路径结束
                    self.pathPool.push(item.path.join(self.seperator));
                }
            }
            if (temp.length) {
                // 重新初始化
                data = temp;
                temp = [];
                recursive(data);
            }
        }
        recursive(data);
    }

    $searchEvent() {
        const self = this;
        let timer;
        this.$search.addEventListener('input', function(event) {
            // 输入防抖
            if (timer) {
                clearTimeout(timer);
            }
            timer = setTimeout(() => {
                const value = event.target.value.trim();
                if (value) {
                    self.searchAction(value);
                } else {
                    self.clearResultAndRefill();
                }
            }, self.debounce);
        });
        // 阻止冒泡
        this.$search.addEventListener('click', function(event) {
            event.stopPropagation();
        });
    }

    searchAction(value) {
        const result = [];
        let tpl = '';
        const reg = new RegExp(value, 'i');
        for (const item of this.pathPool) {
            const match = item.match(reg);
            if (!match) {
                continue;
            }
            result.push(item);
            const index = match.index;
            let [left, center, right] = [item.slice(0, index), match[0], item.slice(index + value.length)];
            // 如果标签内第一个字符是空格，空格会被忽略
            if (right && right.startsWith(' ')) {
                right = `&nbsp;${right.trimLeft()}`;
            }
            tpl += `
                <div class="result">
                    ${left ? `<span class="s">${left}</span>` : ''}
                    <span class="s highlight">${center}</span>
                    ${right ? `<span class="s">${right}</span>` : ''}
                </div>
            `;
        }
        if (!tpl) {
            tpl = '<div class="null">No Result</div>';
        }
        this.$board.innerHTML = tpl;
        // 搜索结果填充board
        this.withResult = true;
        if (result.length) {
            this.$resultEvent(result);
        }
    }

    $resultEvent(result) {
        const self = this;
        const $results = this.$board.querySelectorAll('.result');
        let delay = 100;
        for (let i = 0; i < $results.length; i++) {
            const $result = $results[i];
            // 结果显示动画
            setTimeout(() => {
                $result.classList.add('active');
            }, delay);
            delay += 30;
            $result.addEventListener('click', function(event) {
                event.stopPropagation();
                const res = result[i];
                self.$view.innerHTML = res;
                // 回调
                self.callback(res);
                // 关闭curtain
                self.$trough.click();
                // 清除搜索结果
                self.clearResultAndRefill();
            });
        }
    }

    clearResultAndRefill() {
        this.$board.innerHTML = this.renderCascade(this.data);
        this.$trunks = this.$board.querySelectorAll('.trunk');
        this.$subBoard = this.$board.querySelector('.board');
        this.$rowEvent(this.$board);
        this.$search.value = '';
        this.withResult = false;
    }

    verifyOptions() {
        // $mount
        if (!this.$mount) {
            throw new Error('[Qing error]: 级联选择器组件无法找到挂载点');
        }
        // data
        if (Object.prototype.toString.call(this.data) !== '[object Array]') {
            this.data = [];
            console.warn('[Qing warn]: 级联选择器组件的data必须是数组');
        }
        // debounce
        if (typeof this.debounce !== 'number') {
            this.debounce = 300;
            console.warn('[Qing warn]: 级联选择器组件的debounce必须是数字');
        }
        // seperator
        if (typeof this.seperator !== 'string') {
            this.seperator = ' / ';
            console.warn('[Qing warn]: 级联选择器组件的seperator必须是字符串');
        }
        // callback
        if (typeof this.callback !== 'function') {
            this.callback = () => {};
            console.warn('[Qing warn]: 级联选择器组件的callback必须是函数');
        }
    }
}

////////// *表单公共类* //////////

class FormCommon {
    constructor() {}

    $formFormat($forms) {
        for (let i = 0; i < $forms.length; i++) {
            const $form = $forms[i];
            // 缓存display属性
            const display = getComputedStyle($form).display;
            // 替换不符合要求的tag
            if (['input', 'select', 'option', 'button'].includes($form.tagName.toLowerCase())) {
                const div = document.createElement('div');
                // 复原display属性
                if (display === 'inline' || display === 'inline-block') {
                    div.style.display = 'inline-block';
                } else {
                    div.style.display = display;
                }
                // 转移class
                for (const className of $form.classList) {
                    div.classList.add(className);
                }
                // 替换DOM的旧元素
                $form.parentElement.replaceChild(div, $form);
                // 替换数组的旧元素
                $forms[i] = div;
            } else {
                // 不允许display属性为inline
                if (display === 'inline') {
                    $form.style.display = 'inline-block';
                }
                // 清空元素里面的内容
                $form.innerHTML = '';
            }
        }
    }
}

////////// *多选框组件类* //////////

class Checkbox extends FormCommon {
    constructor({
        classes = 'checkbox',
        indeterminateIndex,
        data = [],
        callback = () => {},
    } = {}) {
        // 继承父类的this对象
        super();
        this.$mounts = [...document.querySelectorAll(`.${classes}`)];
        // 全选全不选checkbox的索引
        this.indeterminateIndex = indeterminateIndex;
        this.data = data;
        this.callback = callback;
        // checkbox数组
        this.$checkboxs = [];
        // checkbox个数
        this.count = this.$mounts.length;
        this.init();
    }

    init() {
        // 检查参数
        this.verifyOptions();
        // 检查标签是否符合要求
        this.$formFormat(this.$mounts);
        // 渲染元素
        this.render();
        // 塑造数据
        this.shapeData();
        // 渲染样式
        this.renderStyle();
        // checkbox事件
        if (this.$indeterminate) {
            this.$indeterminateEvent();
            this.$determinateEvent();
        } else {
            this.$checkboxEvent();
        }
    }

    render() {
        for (let i = 0, count = this.count; i < count; i++) {
            const $mount = this.$mounts[i];
            $mount.innerHTML = '<div class="qing qing-checkbox"></div>';
            const $checkbox = $mount.firstElementChild;
            // 取出全选全不选checkbox
            if (this.indeterminateIndex === i) {
                this.$indeterminate = $checkbox;
                // 数量减一个
                this.count--;
                delete this.indeterminateIndex;
                continue;
            }
            this.$checkboxs.push($checkbox);
        }
        delete this.$mounts;
    }

    shapeData() {
        // 锁定data的长度
        this.data.length = this.count;
        for (let i = 0; i < this.count; i++) {
            if (Object.prototype.toString.call(this.data[i]) !== '[object Object]') {
                this.data[i] = {
                    checked: false,
                    disabled: false,
                };
            } else {
                const item = this.data[i];
                if (item.checked === undefined) {
                    item.checked = false;
                }
                if (item.disabled === undefined) {
                    item.disabled = false;
                }
            }
        }
    }

    renderStyle() {
        for (let i = 0; i < this.count; i++) {
            const item = this.data[i];
            const CL = this.$checkboxs[i].classList;
            if (item.checked !== CL.contains('checked')) {
                CL.toggle('checked');
            }
            if (item.disabled !== CL.contains('disabled')) {
                CL.toggle('disabled');
            }
        }
    }

    $checkboxEvent() {
        const self = this;
        for (let i = 0; i < this.count; i++) {
            const $checkbox = this.$checkboxs[i];
            const CL = $checkbox.classList;
            $checkbox.addEventListener('click', function(event) {
                event.stopPropagation();
                // 禁用的checkbox不触发事件
                if (self.data[i].disabled === true) {
                    return;
                }
                self.$checkboxAction(CL, i);
            });
        }
    }

    $indeterminateEvent() {
        const self = this;
        const indeCL = this.$indeterminate.classList;
        this.$indeterminate.addEventListener('click', function(event) {
            event.stopPropagation();
            // 全选或者全不选
            if (indeCL.contains('checked')) {
                for (let i = 0; i < self.count; i++) {
                    // 禁用的checkbox不触发事件
                    if (self.data[i].disabled === true) {
                        continue;
                    }
                    const deCL = self.$checkboxs[i].classList;
                    deCL.contains('checked') ? deCL.remove('checked') : '';
                    // 数据变更
                    self.data[i].checked = false;
                }
            } else {
                for (let i = 0; i < self.count; i++) {
                    // 禁用的checkbox不触发事件
                    if (self.data[i].disabled === true) {
                        continue;
                    }
                    const deCL = self.$checkboxs[i].classList;
                    !deCL.contains('checked') ? deCL.add('checked') : '';
                    // 数据变更
                    self.data[i].checked = true;
                }
            }
            // 自己的action
            indeCL.contains('somechecked') ? indeCL.remove('somechecked') : '';
            indeCL.toggle('checked');
            // 全选全不选时，回调的第二个参数是'indeterminate'
            self.callback(self.data, 'indeterminate');
        });
    }

    $determinateEvent() {
        const self = this;
        const indeCL = self.$indeterminate.classList;
        for (let i = 0; i < this.count; i++) {
            const $determinate = this.$checkboxs[i];
            const deCL = $determinate.classList;
            $determinate.addEventListener('click', function(event) {
                event.stopPropagation();
                // 禁用的checkbox不触发事件
                if (self.data[i].disabled === true) {
                    return;
                }
                const checked = !deCL.contains('checked');
                let noneFlag = true;
                let allFlag = true;
                for (const $item of self.$checkboxs) {
                    if ($item === this) {
                        continue;
                    }
                    if ($item.classList.contains('checked')) {
                        noneFlag ? noneFlag = false : '';
                    } else {
                        allFlag ? allFlag = false : '';
                    }
                    if (!noneFlag && !allFlag) {
                        break;
                    }
                }
                if (checked && allFlag) {
                    indeCL.contains('somechecked') ? indeCL.remove('somechecked') : '';
                    !indeCL.contains('checked') ? indeCL.add('checked') : '';
                } else if (!checked && noneFlag) {
                    indeCL.contains('checked') ? indeCL.remove('checked') : '';
                    indeCL.contains('somechecked') ? indeCL.remove('somechecked') : '';
                } else {
                    indeCL.contains('checked') ? indeCL.remove('checked') : '';
                    !indeCL.contains('somechecked') ? indeCL.add('somechecked') : '';
                }
                // 自己的action
                self.$checkboxAction(deCL, i);
            });
        }
    }

    $checkboxAction(CL, i) {
        if (CL.contains('checked')) {
            CL.remove('checked');
            this.data[i].checked = false;
            // 回调
            this.callback(this.data, i);
        } else {
            CL.add('checked');
            this.data[i].checked = true;
            // 回调
            this.callback(this.data, i);
        }
    }

    updateData(data) {
        this.data = data;
        this.shapeData();
        this.renderStyle();
    }

    verifyOptions() {
        // $mounts
        if (!this.$mounts) {
            throw new Error('[Qing error]: 多选框组件无法找到挂载点');
        }
        // indeterminateIndex
        if (this.indeterminateIndex !== undefined) {
            if (typeof this.indeterminateIndex !== 'number') {
                delete this.indeterminateIndex;
                console.warn('[Qing warn]: 多选框组件的indeterminateIndex必须是数字');
            } else if (this.indeterminateIndex >= this.count || this.indeterminateIndex < 0) {
                this.indeterminateIndex = 0;
                console.warn('[Qing warn]: 多选框组件的indeterminateIndex必须在多选框个数之内');
            } else if (!Number.isInteger(this.indeterminateIndex)) {
                this.indeterminateIndex = 0;
                console.warn('[Qing warn]: 多选框组件的indeterminateIndex必须是整数');
            }
        }
        // data
        if (Object.prototype.toString.call(this.data) !== '[object Array]') {
            this.data = [];
            console.warn('[Qing warn]: 多选框组件的data必须是数组');
        }
        // callback
        if (typeof this.callback !== 'function') {
            this.callback = () => {};
            console.warn('[Qing warn]: 多选框组件的callback必须是函数');
        }
    }
}

////////// *单选框组件类* //////////

class Radio extends FormCommon {
    constructor({
        classes = 'radio',
        data = [],
        callback = () => {},
    } = {}) {
        // 继承父类的this对象
        super();
        this.$mounts = document.querySelectorAll(`.${classes}`);
        this.data = data;
        this.callback = callback;
        // radio数组
        this.$radios = [];
        // radio个数
        this.count = this.$mounts.length;
        // 被选中的radio索引
        this.checkedIndex;
        this.init();
    }

    init() {
        // 检查参数
        this.verifyOptions();
        // 检查标签是否符合要求
        this.$formFormat(this.$mounts);
        // 渲染元素
        this.render();
        // 塑造数据
        this.shapeData();
        // 渲染样式
        this.renderStyle();
        // radio事件
        this.$radioEvent();
    }

    render() {
        for (let i = 0, count = this.count; i < count; i++) {
            const $mount = this.$mounts[i];
            $mount.innerHTML = '<div class="qing qing-radio"></div>';
            const $radio = $mount.firstElementChild;
            this.$radios.push($radio);
        }
        delete this.$mounts;
    }

    shapeData() {
        // 锁定data的长度
        this.data.length = this.count;
        for (let i = 0; i < this.count; i++) {
            if (Object.prototype.toString.call(this.data[i]) !== '[object Object]') {
                this.data[i] = {
                    checked: false,
                    disabled: false,
                };
            } else {
                const item = this.data[i];
                // 检查disabled，如果radio同时checked和disabled，则取消checked
                if (item.disabled !== true && item.disabled !== false) {
                    item.disabled = false;
                } else if (item.disabled === true && item.checked === true) {
                    console.warn('[Qing warn]: 置灰的单选框radio不允许checked');
                    item.checked = false;
                }
                // 检查checked
                if (item.checked !== true && item.checked !== false) {
                    item.checked = false;
                } else if (item.checked === true) {
                    if (this.checkedIndex === undefined) {
                        this.checkedIndex = i;
                    } else {
                        console.warn('[Qing warn]: 单选框radio只能有一个checked');
                        item.checked = false;
                    }
                }
            }
        }
        if (this.checkedIndex === undefined) {
            console.warn('[Qing warn]: 单选框radio必须有一个checked');
            this.data[0].checked = true;
            this.checkedIndex = 0;
        }
    }

    renderStyle() {
        for (let i = 0; i < this.count; i++) {
            const item = this.data[i];
            const CL = this.$radios[i].classList;
            if (item.checked !== CL.contains('checked')) {
                CL.toggle('checked');
            }
            if (item.disabled !== CL.contains('disabled')) {
                CL.toggle('disabled');
            }
        }
    }

    $radioEvent() {
        const self = this;
        for (let i = 0; i < this.count; i++) {
            const $radio = this.$radios[i];
            const CL = $radio.classList;
            $radio.addEventListener('click', function(event) {
                event.stopPropagation();
                // 置灰的元素不触发事件
                if (self.data[i].disabled === true) {
                    return;
                }
                if (!CL.contains('checked')) {
                    // 上一个checked的action
                    self.$radios[self.checkedIndex].classList.remove('checked');
                    self.data[self.checkedIndex].checked = false;
                    // 改变指针
                    self.checkedIndex = i;
                    // 自己的action
                    CL.add('checked');
                    self.data[i].checked = true;
                    self.callback(self.data, i);
                }
            });
        }
    }

    updateData(data) {
        this.data = data;
        // 重置指针
        this.checkedIndex = undefined;
        this.shapeData();
        this.renderStyle();
    }

    verifyOptions() {
        // $mounts
        if (!this.$mounts) {
            throw new Error('[Qing error]: 单选框组件无法找到挂载点');
        }
        // data
        if (Object.prototype.toString.call(this.data) !== '[object Array]') {
            this.data = [];
            console.warn('[Qing warn]: 单选框组件的data必须是数组');
        }
        // callback
        if (typeof this.callback !== 'function') {
            this.callback = () => {};
            console.warn('[Qing warn]: 单选框组件的callback必须是函数');
        }
    }
}

////////// *开关组件类* //////////

class Switch extends FormCommon {
    constructor({
        id = 'switch',
        checked = false,
        disabled = false,
        size = 'default',
        callback = () => {},
    } = {}) {
        // 继承父类的this对象
        super();
        this.$mount = document.querySelector(`#${id}`);
        // 是否选中
        this.checked = checked;
        // 是否置灰
        this.disabled = disabled;
        // 两种规格大小，default和small
        this.size = size;
        this.callback = callback;
        this.init();
    }

    init() {
        // 检查参数
        this.verifyOptions();
        // 用一个变量缓存，确保元素被替换后，依然能够找到引用地址
        const nodeList = [this.$mount];
        this.$formFormat(nodeList);
        this.$mount = nodeList[0];
        this.render();
        this.$switchEvent();
    }

    render() {
        // 不同size的class不一样
        const classBySize = this.size === 'small' ? 'qing qing-switch switch-small' : 'qing qing-switch';
        let classList = [classBySize];
        this.checked === true ? classList.push('checked') : '';
        this.disabled === true ? classList.push('disabled') : '';
        classList = classList.join(' ');
        const tpl = `
            <div class="${classList}">
                <span class="${this.checked ? 'button checked' : 'button'}"></span>
            </div>
        `;
        this.$mount.innerHTML = tpl;
        this.$switch = this.$mount.querySelector(`.qing-switch`);
    }

    $switchEvent() {
        const self = this;
        const switchCL = this.$switch.classList;
        const buttonCL = this.$switch.querySelector('.button').classList;
        this.$switch.addEventListener('click', function(event) {
            event.stopPropagation();
            // 置灰的元素不触发事件
            if (self.disabled === true) {
                return;
            }
            buttonCL.toggle('checked');
            if (switchCL.contains('checked')) {
                switchCL.remove('checked');
                self.checked = false;
            } else {
                switchCL.add('checked');
                self.checked = true;
            }
            self.callback(self.checked);
        });
    }

    verifyOptions() {
        // $mount
        if (!this.$mount) {
            throw new Error('[Qing error]: 开关组件无法找到挂载点');
        }
        // callback
        if (typeof this.callback !== 'function') {
            this.callback = () => {};
            console.warn('[Qing warn]: 开关组件的callback必须是函数');
        }
    }
}

////////// *计数器组件类* //////////

class InputNumber extends FormCommon {
    constructor({
        id = 'input-number',
        initValue = 1,
        step = 1,
        min,
        max,
        size = 'default',
        disabled = false,
        callback = () => {},
    } = {}) {
        // 继承父类的this对象
        super();
        this.$mount = document.querySelector(`#${id}`);
        // 初始数值
        this.initValue = initValue;
        // 缓存上一个数值
        this.oldValue = this.initValue;
        // 计数器步长
        this.step = step;
        // 允许的最小值，需要考虑到初始值和步长
        min !== undefined ? this.min = (this.initValue - min) % this.step : '';
        // 允许的最大值，需要考虑到初始值和步长
        max !== undefined ? this.max = max - ((max - this.initValue) % this.step) : '';
        // 两种规格大小，default和small
        this.size = size;
        this.disabled = disabled;
        this.callback = callback;
        this.init();
    }

    init() {
        // 检查参数
        this.verifyOptions();
        // 用一个变量缓存，确保元素被替换后，依然能够找到引用地址
        const nodeList = [this.$mount];
        this.$formFormat(nodeList);
        this.$mount = nodeList[0];
        this.render();
        if (this.disabled !== true) {
            this.$decreaseEvent();
            this.$increaseEvent();
            this.$monitorEvent();
        }
    }

    render() {
        // 不同size的class不一样
        const classList = this.size === 'small' ? 'qing qing-input-number input-number-small' : 'qing qing-input-number';
        let tpl;
        if (this.disabled !== true) {
            tpl = `
                <div class="${classList}">
                    <div class="decrease">-</div>
                    <input class="monitor" type="text">
                    <div class="increase">+</div>
                </div>
            `;
        } else {
            tpl = `
                <div class="${classList}">
                    <div class="decrease disabled">-</div>
                    <input class="monitor" type="text" disabled>
                    <div class="increase disabled">+</div>
                </div>
            `;
        }
        this.$mount.innerHTML = tpl;
        this.$monitor = this.$mount.querySelector('.monitor');
        this.$decrease = this.$mount.querySelector('.decrease');
        this.$increase = this.$mount.querySelector('.increase');
        // 初始值
        this.$monitor.value = this.initValue;
    }

    $decreaseEvent() {
        const self = this;
        const deCL = this.$decrease.classList;
        const inCL = this.$increase.classList;
        this.$decrease.addEventListener('click', function(event) {
            event.stopPropagation();
            // 置灰的元素不触发事件
            if (deCL.contains('disabled')) {
                return;
            }
            const value = Number.parseInt(self.$monitor.value) - self.step;
            // 如果比min小，则disabled；同时清除另一边的disabled，如果有的话
            inCL.contains('disabled') ? inCL.remove('disabled') : '';
            if (self.min !== undefined && value <= self.min) {
                deCL.add('disabled');
            }
            self.$monitor.value = value;
            self.callback(value, self.oldValue, 'decrease');
            // 缓存该数值
            self.oldValue = value;
        });
    }

    $increaseEvent() {
        const self = this;
        const deCL = this.$decrease.classList;
        const inCL = this.$increase.classList;
        this.$increase.addEventListener('click', function(event) {
            event.stopPropagation();
            // 置灰的元素不触发事件
            if (inCL.contains('disabled')) {
                return;
            }
            const value = Number.parseInt(self.$monitor.value) + self.step;
            // 如果比max大，则disabled；同时清除另一边的disabled，如果有的话
            deCL.contains('disabled') ? deCL.remove('disabled') : '';
            if (self.max !== undefined && value >= self.max) {
                inCL.add('disabled');
            }
            self.$monitor.value = value;
            self.callback(value, self.oldValue, 'increase');
            // 缓存该数值
            self.oldValue = value;
        });
    }

    $monitorEvent() {
        const self = this;
        const deCL = this.$decrease.classList;
        const inCL = this.$increase.classList;
        this.$monitor.addEventListener('change', function(event) {
            event.stopPropagation();
            const result = this.value.trim();
            let value = Number(result);
            // 输入非数字不触发事件
            if (Number.isNaN(value)) {
                this.value = self.oldValue;
                return;
            }
            // 值为空时不触发事件
            if (result === '') {
                return;
            }
            // 判断当前值的是否超过最值
            if (self.min !== undefined && value <= self.min) {
                deCL.add('disabled');
                // 小于最小值，则等于最小值
                value = self.min;
                this.value = value;
            } else if (self.max !== undefined && value >= self.max) {
                inCL.add('disabled');
                // 大于最大值，则等于最大值
                value = self.max;
                this.value = value;
            } else {
                deCL.contains('disabled') ? deCL.remove('disabled') : '';
                inCL.contains('disabled') ? inCL.remove('disabled') : '';
            }
            // 值的增量或减量必须是step的倍数
            if (value > self.initValue) {
                const multiple = (value - self.initValue) / self.step;
                // 不是倍数时改成倍数，数字永远向初始值靠拢
                if (!Number.isInteger(multiple)) {
                    value = Number.parseInt(multiple) * self.step + self.initValue;
                    this.value = value;
                }
            } else if (value < self.initValue) {
                const multiple = (self.initValue - value) / self.step;
                // 不是倍数时改成倍数，数字永远向初始值靠拢
                if (!Number.isInteger(multiple)) {
                    value = Number.parseInt(multiple) * self.step + self.initValue;
                    this.value = value;
                }
            }
            self.callback(value, self.oldValue, 'input');
            // 缓存该数值
            self.oldValue = value;
        });
    }

    verifyOptions() {
        // $mount
        if (!this.$mount) {
            throw new Error('[Qing error]: 计数器组件无法找到挂载点');
        }
        // initValue
        if (typeof this.initValue !== 'number') {
            this.initValue = 1;
            console.warn('[Qing warn]: 计数器组件的initValue必须是数字');
        }
        // step
        if (typeof this.step !== 'number') {
            this.step = 1;
            console.warn('[Qing warn]: 计数器组件的step必须是数字');
        }
        // min和max
        if (this.min !== undefined && Number.isNaN(this.min)) {
            delete this.min;
            console.warn('[Qing warn]: 计数器组件的min必须是数字');
        }
        if (this.max !== undefined && Number.isNaN(this.max)) {
            delete this.max;
            console.warn('[Qing warn]: 计数器组件的max必须是数字');
        }
        if (this.min !== undefined && this.max !== undefined && this.min >= this.max) {
            delete this.min;
            delete this.max;
            console.warn('[Qing warn]: 计数器组件的max必须比min大');
        }
    }
}

////////// *输入框组件类* //////////

class Input extends FormCommon {
    constructor({
        id = 'input',
        eventType = 'click',
        buttonLabel = 'click',
        debounce = 300,
        placeholder,
        clearable = false,
        callback = () => {},
    } = {}) {
        // 继承父类的this对象
        super();
        this.$mount = document.querySelector(`#${id}`);
        // 事件类型：click，input，change
        this.eventType = eventType;
        // 只有事件类型是click时才有效
        this.eventType === 'click' ? this.buttonLabel = buttonLabel : '';
        // 防抖间隔，只有事件类型是input才有效
        this.eventType === 'input' ? this.debounce = debounce : '';
        this.placeholder = placeholder;
        // 清除按钮
        this.clearable = clearable;
        this.callback = callback;
        this.init();
    }

    init() {
        // 检查参数
        this.verifyOptions();
        // 用一个变量缓存，确保元素被替换后，依然能够找到引用地址
        const nodeList = [this.$mount];
        this.$formFormat(nodeList);
        this.$mount = nodeList[0];
        this.render();
        this.$inputEvent();
        if (this.clearable === true) {
            this.$clearEvent();
        }
    }

    render() {
        const classList = this.eventType === 'click' ? 'input click' : 'input';
        const tpl = `
            <div class="qing qing-input">
                <span class="wrap">
                    <input class="${classList}" type="text" placeholder="${this.placeholder}">
                    ${this.clearable === true ? '<span class="clear">+</span>' : ''}
                </span>
                ${this.eventType === 'click' ? `<button class="button">${this.buttonLabel}</button>` : ''}
            </div>
        `;
        this.$mount.innerHTML = tpl;
        this.$input = this.$mount.querySelector('.input');
        if (this.clearable === true) {
            this.$clear = this.$mount.querySelector('.clear');
        }
    }

    $inputEvent() {
        const self = this;
        if (this.eventType === 'click') {
            const $button = this.$mount.querySelector('.button');
            $button.addEventListener('click', function(event) {
                event.stopPropagation();
                self.callback(self.$input.value);
            });
        } else if (this.eventType === 'input') {
            let timer;
            this.$input.addEventListener('input', function() {
                clearTimeout(timer);
                timer = setTimeout(() => {
                    self.callback(this.value);
                }, self.debounce);
            });
        } else if (this.eventType === 'change') {
            this.$input.addEventListener('change', function() {
                self.callback(this.value);
            });
        }
        // $clear按钮的显示隐藏
        if (this.clearable === true) {
            const CL = this.$clear.classList;
            let timer;
            this.$input.addEventListener('input', function() {
                clearTimeout(timer);
                timer = setTimeout(() => {
                    if (this.value) {
                        !CL.contains('active') ? CL.add('active') : '';
                    } else {
                        CL.contains('active') ? CL.remove('active') : '';
                    }
                }, 300);
            });
        }
    }

    $clearEvent() {
        const self = this;
        const CL = this.$clear.classList;
        this.$clear.addEventListener('click', function(event) {
            event.stopPropagation();
            self.$input.value = '';
            // 重新聚焦
            self.$input.focus();
            CL.remove('active');
        });
    }

    verifyOptions() {
        // $mount
        if (!this.$mount) {
            throw new Error('[Qing error]: 输入框组件无法找到挂载点');
        }
        // eventType
        const typeList = ['input', 'change', 'click'];
        if (!typeList.includes(this.eventType)) {
            this.eventType = 'click';
            console.warn('[Qing warn]: 输入框的eventType不是有效类型');
        }
        // debounce
        if (typeof this.debounce !== 'number') {
            this.debounce = 300;
            console.warn('[Qing warn]: 输入框的debounce必须是数字');
        }
        // callback
        if (typeof this.callback !== 'function') {
            this.callback = () => {};
            console.warn('[Qing warn]: 输入框的callback必须是函数');
        }
    }
}

export {DatePicker, TimePicker, Paginator, Tree, Cascader, Checkbox, Radio, Switch, InputNumber, Input};
