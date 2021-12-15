直接在F12开发者工具里进行调试。 有些网页不用jQuery，为了方便，我们需要自己将jQuery导入到页面中，
可以将下面的代码复制到浏览器控制台中。

var jq = document.createElement('script');
jq.src = "https://cdn.staticfile.org/jquery/3.4.1/jquery.min.js";
document.getElementsByTagName('head')[0].appendChild(jq);

https://segmentfault.com/a/1190000021654926

// ==UserScript==
// @name         remind_me_vagrant_update
// @namespace    https://github.com/techstay/myscripts
// @version      0.1
// @description  remind me if vagrant support virtualbox
// @author       techstay
// @match        *
// @require      https://cdn.staticfile.org/jquery/3.4.1/jquery.min.js
// @connect      vagrantup.com
// @grant        GM_setValue
// @grant        GM_getValue
// @grant        GM_setClipboard
// @grant        GM_log
// @grant        GM_xmlhttpRequest
// @grant        unsafeWindow
// @grant        window.close
// @grant        window.focus
// ==/UserScript==

(function () {
    'use strict';

    const CHECKED_DATE = 'checkedDate';

    function checkDateEquals(a, b) {
        return a.getFullYear() === b.getFullYear() &&
            a.getMonth() === b.getMonth() &&
            a.getDate() === b.getDate();
    }

    function checkVagrantVersion() {
        GM_setValue(CHECKED_DATE, new Date());
        GM_xmlhttpRequest({
            "method": "GET",
            "url": "https://www.vagrantup.com/",
            "headers": {
                "user-agent": 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36'
            },
            "onload": function (result) {
                var list = jQuery.parseHTML(result.response);
                jQuery.each(list, function (i, el) {
                    if (el.nodeName == 'HEADER') {
                        var header = jQuery.parseXML(el.innerHTML);
                        var version = header.getElementsByTagName('a')[1].textContent.replace('Download ', '');
                        if (version != '2.2.6') {
                            alert('Vagrant update!');
                        }
                        return false;
                    }
                });
            }
        });
    }

    var today = new Date();
    var lastCheckedDay = new Date(GM_getValue(CHECKED_DATE, new Date('2006-1-1')));
    if (!checkDateEquals(lastCheckedDay, today)) {
        checkVagrantVersion();
    }
})();


// ==UserScript==
// @name         copy_jianshu_to_csdn_and_segmentfault
// @namespace    https://github.com/techstay/myscripts
// @version      0.1
// @description  将简书文章复制到csdn和思否编辑器中
// @author       techstay
// @match        https://editor.csdn.net/md/
// @match        https://segmentfault.com/write
// @match        https://www.jianshu.com/writer*
// @require      https://cdn.staticfile.org/jquery/3.4.1/jquery.min.js
// @require      https://cdn.bootcss.com/jqueryui/1.12.1/jquery-ui.min.js
// @grant        GM_setValue
// @grant        GM_getValue
// @grant        GM_deleteValue
// @grant        unsafeWindow
// @grant        GM_setClipboard
// @grant        window.close
// @grant        window.focus
// @grant        GM_openInTab
// ==/UserScript==

(function () {
    'use strict';

    const SF_URL = 'https://segmentfault.com/write'
    const CSDN_URL = 'https://editor.csdn.net/md/'

    const SF_TITLE = 'sf_title'
    const SF_CONTENT = 'sf_content'
    const CSDN_TITLE = 'csdn_title'
    const CSDN_CONTENT = 'csdn_content'

    function saveArticle() {
        GM_setValue(CSDN_TITLE, $('._24i7u').val())
        GM_setValue(CSDN_CONTENT, $('#arthur-editor').val())
        GM_setValue(SF_TITLE, $('._24i7u').val())
        GM_setValue(SF_CONTENT, $('#arthur-editor').val())
    }

    function copyToCsdn() {
        var title = GM_getValue(CSDN_TITLE, '')
        var content = GM_getValue(CSDN_CONTENT, '')
        if (title != '' && content != '') {
            $('.article-bar__title').delay(2000).queue(function () {
                $('.article-bar__title').val(title)
                $('.editor__inner').text(content)
                GM_deleteValue(CSDN_TITLE)
                GM_deleteValue(CSDN_CONTENT)
                $(this).dequeue()
            })
        }
    }

    function copyToSegmentFault() {
        $(document).ready(function () {
            var title = GM_getValue(SF_TITLE, '')
            var content = GM_getValue(SF_CONTENT, '')
            if (title != '' && content != '') {
                $('#title').delay(2000).queue(function () {
                    $('#title').val(title)
                    GM_setClipboard(content, 'text')
                    GM_deleteValue(SF_TITLE)
                    GM_deleteValue(SF_CONTENT)
                    $(this).dequeue()
                })
            }
        })
    }

    function addCopyButton() {
        $('body').append('<div id="copyToCS">双击复制到CSDN和思否</div>')
        $('#copyToCS').css('width', '200px')
        $('#copyToCS').css('position', 'absolute')
        $('#copyToCS').css('top', '70px')
        $('#copyToCS').css('left', '350px')
        $('#copyToCS').css('background-color', '#28a745')
        $('#copyToCS').css('color', 'white')
        $('#copyToCS').css('font-size', 'large')
        $('#copyToCS').css('z-index', 100)
        $('#copyToCS').css('border-radius', '25px')
        $('#copyToCS').css('text-align', 'center')
        $('#copyToCS').dblclick(function () {
            saveArticle()
            GM_openInTab(SF_URL, true)
            GM_openInTab(CSDN_URL, true)
        })
        $('#copyToCS').draggable()
    }

    $(document).ready(function () {
        if (window.location.href.startsWith('https://www.jianshu.com')) {
            addCopyButton()
        } else if (window.location.href.startsWith(SF_URL)) {
            copyToSegmentFault()
        } else if (window.location.href.startsWith(CSDN_URL)) {
            copyToCsdn()
        }
    })
})()
