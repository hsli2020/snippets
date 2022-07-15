// ==UserScript==
// @name         zhihutime - 知乎时间跨度
// @version      0.3
// @description  展示回答或者文章的发表时间距今的时间跨度
// @author       lucienlugeek@gmail.com
// @match        https://www.zhihu.com/question/*
// @match        https://zhuanlan.zhihu.com/*
// @icon         https://www.google.com/s2/favicons?sz=64&domain=zhihu.com
// @grant        GM_addStyle
// @require      https://cdnjs.cloudflare.com/ajax/libs/moment.js/2.29.3/moment.min.js
// @license      MIT
// @namespace https://greasyfork.org/users/909394
// ==/UserScript==
 
(function() {
    'use strict';
 
    // 问题与答案，并且可以点击【查看所有答案】的页面
    const Q_A = /^https:\/\/www.zhihu.com\/question\/\d{1,}\/answer\/\d{1,}$/;
    // 纯粹查看所有答案的页面
    const Q = /^https:\/\/www.zhihu.com\/question\/\d{1,}$/;
    // 知乎专栏页面
    const Z = /^https:\/\/zhuanlan.zhihu.com.*$/;
 
    const style = `
        .warning-date {
            -webkit-text-size-adjust: 100%;
            word-break: normal;
            font-size: 1.4rem;
            margin-bottom: 10px;
            display: flex;
            -webkit-box-align: center;
            align-items: center;
            padding: 10px 16px;
            background-color: #fcf0bf;
            color: rgba(0, 0, 0, 0.57);
            line-height: 1.5;
            font-weight: 600;
            border-radius: 3px;
        }
 
        .warning-date-year, .warning-date-month {
            background-color: #fcf0bf;
        }
        .warning-date-day, .warning-date-hour {
            background-color: #cafcca;
        }
        `;
 
    GM_addStyle(style);
 
    // render time warning info by given date and node
    const render_dom = function(date, answerNode) {
        // const date_modified = answer.querySelector(`meta[itemprop='dateModified']`).getAttribute('content');
        const duration = moment.duration(moment().diff(moment(date)));
        const node = document.createElement('div');
        // year
        const y = duration.get('year');
        // month
        const m = duration.get('month');
        // day
        const d = duration.get('day');
        // hour
        const h = duration.get('hour');
 
        node.classList.add('warning-date');
 
        if (y !== 0) {
            // years ago
            node.classList.add('warning-date-year');
            node.innerHTML = `⚠️ 回答于${y}年前`;
        } else if (m !== 0) {
            // months ago
            node.classList.add('warning-date-month');
            node.innerHTML = `⚠️ 回答于${m}个月前`;
        } else if (d !== 0) {
            // some days ago
            node.classList.add('warning-date-day');
            node.innerHTML = `ℹ 回答于${d}天前`;
        } else if (h !== 0) {
            // just some hours ago
            node.classList.add('warning-date-hour');
            node.innerHTML = `ℹ 回答于${h}小时前`;
        }
 
        // insert date warning before the main content.
        const first = answerNode.firstChild;
        if (first.classList[0] !== 'warning-date') {
            answerNode.insertBefore(node, answerNode.firstChild);
        }
    };
 
    // render datetime warning before the main content
    const render = function() {
        // console.log('render', new Date().getTime());
        const answerItems = Array.from(document.querySelectorAll('.AnswerItem'));
        answerItems.map(answer => {
            const date_modified = answer.querySelector(`meta[itemprop='dateModified']`).getAttribute('content');
            render_dom(date_modified, answer);
        });
    };
 
    // 知乎专栏页面
    if (Z.test(document.URL)) {
        const t = document.querySelector('.ContentItem-time').innerText;
        const first_space_index = t.indexOf(' ');
        const date = t.slice(first_space_index);
        render_dom(date, document.querySelector('.Post-Header'));
        return;
    }
 
    let observer = null;
 
    // execute render after 1000ms delay.
    const delay = function() {
        setTimeout(function() {
            if (observer) {
                // 停止观察
                observer.disconnect();
            }
            render();
            createObserver();
        }, 1000);
    }
 
    const createObserver = function() {
        // 选择需要观察变动的节点
        let targetNodeSelectorString = '';
        if (Q.test(document.URL)) {
            targetNodeSelectorString = `.AnswersNavWrapper [role='list']`;
        } else if (Q_A.test(document.URL)) {
            targetNodeSelectorString = `.Question-mainColumn`;
            // 点击【查看所有答案】按钮时，触发render()方法
            const viewAllEle = Array.from(document.querySelectorAll('.ViewAll-QuestionMainAction'));
            viewAllEle.map(m => {
                m.removeEventListener('click', delay);
                m.addEventListener('click', delay);
            });
        } else {
            return;
        }
 
        const targetNode = document.querySelector(targetNodeSelectorString);
 
        // 观察器的配置（需要观察什么变动）
        const config = { attributes: false, childList: true, subtree: false };
 
        // 当观察到变动时执行的回调函数
        const callback = function(mutationsList, observer) {
            // Use traditional 'for loops' for IE 11
            for(let mutation of mutationsList) {
                if (mutation.type === 'childList') {
                    render();
                }
            }
        };
 
        // 创建一个观察器实例并传入回调函数
        observer = new MutationObserver(callback);
 
        // 以上述配置开始观察目标节点
        observer.observe(targetNode, config);
    }
 
    render();
    createObserver();
 
    window.addEventListener('popstate', function () {
        // console.log('url state changed!');
        if (observer) {
            // 停止观察
            observer.disconnect();
        }
        render();
        createObserver();
    });
 
})();