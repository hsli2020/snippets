// ==UserScript==
// @name         Auto Fill Address Form
// @namespace    http://tampermonkey.net/
// @version      0.1
// @description  Auto Fill Address Form!
// @author       You
// @match        https://www.amazon.ca/gp/buy/spc/handlers/display.html
// @match        https://www.amazon.ca/gp/buy/addressselect/handlers/display.html?hasWorkingJavascript=1
// @icon         https://www.google.com/s2/favicons?domain=amazon.ca
// @grant        GM_log
// @grant        GM_xmlhttpRequest
// ==/UserScript==

(function() {
    'use strict';

    var div = createElementFromHTML('<div style="margin-bottom:10px;float:right;width:50%">');

    var left = createElementFromHTML('<div style="float:left;width:60%;">');
    var input = createElementFromHTML('<input type="text" id="my-order-number">');
    left.append(input);

    var btn = createElementFromHTML('<button type="button">Fill Form</button>');
    btn.addEventListener("click", autoFillForm);

    div.append(left);
    div.append(btn);

    var e = document.querySelector('#add-address-popover');
    e.prepend(div);

    function createElementFromHTML(htmlString) {
        var div = document.createElement('div');
        div.innerHTML = htmlString.trim();

        // Change this to div.childNodes to support multiple top-level nodes
        return div.firstChild;
    }

    function autoFillForm() {
        var e;

        e = document.querySelector('#my-order-number');
        var orderId = e.value;
        //console.log(orderId);

        GM_xmlhttpRequest({
            "method": "GET",
            "url": "http://localhost/order/address/" + orderId,
            "onload": function (result) {
                //console.log(result.response);

                var obj = JSON.parse(result.response);
                var o = obj.data;
                //console.log(o);

                e = document.querySelector('[name=address-ui-widgets-enterAddressFullName]');
                e.value = o.contact;

                e = document.querySelector('[name=address-ui-widgets-enterAddressPhoneNumber]');
                e.value = o.phone;

                e = document.querySelector('[name=address-ui-widgets-enterAddressLine1]');
                e.value = o.address;

                e = document.querySelector('[name=address-ui-widgets-enterAddressLine2]');
                e.value = "";

                e = document.querySelector('[name=address-ui-widgets-enterAddressCity]');
                e.value = o.city;

                var province = getProvinceName(o.province);

                e = document.querySelector('[name=address-ui-widgets-enterAddressStateOrRegion]');
                e.value = province;

                e = document.querySelector('#address-ui-widgets-enterAddressStateOrRegion .a-dropdown-prompt');
                e.innerText = province;

                e = document.querySelector('[name=address-ui-widgets-enterAddressPostalCode]');
                e.value = o.postalcode;
            }
        });
    }

    function getProvinceName(code) {
        var provinces = {
            "AB" :"Alberta",
            "BC" :"British Columbia",
            "MB" :"Manitoba",
            "NB" :"New Brunswick",
            "NL" :"Newfoundland",
            "NT" :"Northwest Territories",
            "NS" :"Nova Scotia",
            "NU" :"Nunavut",
            "ON" :"Ontario",
            "PE" :"Prince Edward Island",
            "QC" :"Quebec",
            "SK" :"Saskatchewan",
            "YU" :"Yukon",
        };
        return provinces[code];
    }
})();