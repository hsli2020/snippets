// Typescript
function $(elem: string) {
  return document.querySelector(elem) as HTMLDivElement
}
function $all(elem: string) {
  return document.querySelectorAll(elem) as unknown as HTMLDivElement[]
}
function $$(elem: string, el: string) {
  return $(elem).querySelectorAll(el)[1] as HTMLDivElement
}

// Javascript
function $(elem) {
  return document.querySelector(elem);
}
function $all(elem) {
  return document.querySelectorAll(elem);
}
function $$(elem, el) {
  return $(elem).querySelectorAll(el)[1];
}

'use strict';
 
var s = document.createElement('script');
s.setAttribute('src', 'https://lib.sinaapp.com/js/jquery/2.0.3/jquery-2.0.3.min.js');

s.onload = function() { delete_action(scroll_to_bottom) };
document.head.appendChild(s);

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

await sleep(500)
$('a[action-type="ok"]')[0].click()

function scroll_to_bottom() {
    $('html, body').animate({ scrollTop: $(document).height() }, 'slow');
}
