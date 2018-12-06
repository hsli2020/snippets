'use strict';

// convert object to query string
function stringifyObject(json) {
  var keys = Object.keys(json);

  return '?' +
      keys.map(function(k) {
          return encodeURIComponent(k) + '=' +
              encodeURIComponent(json[k]);
      }).join('&');
}

function randomString(n) {
  var s = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
  return Array(n).join().split(',').map(() => s.charAt(Math.floor(Math.random() * s.length))).join('');
}

const M = 1000000
const K = 1000
const rx = new RegExp('\\.0$');
const commaRx = new RegExp('(\\d+)(\\d{3})');

function formatPretty(num) {
  let decimals = 0;

  if (num >= M) {
    num /= M
    decimals = 3 - ((Math.round(num) + "").length) || 0;
    return (num.toFixed(decimals > -1 ? decimals : 0).replace(rx, '') + 'M').replace('.00', '');
  }

  if (num >= (K * 10)) {
    num /= K
    decimals = 3 - ((Math.round(num) + "").length) || 0;
    return num.toFixed(decimals).replace(rx, '') + 'K';
  }

  return formatWithComma(num);
}

function formatWithComma(nStr) {
	nStr += '';

  if(nStr.length < 4 ) {
    return nStr;
  }

  var	x = nStr.split('.');
	var x1 = x[0];
	var x2 = x.length > 1 ? '.' + x[1] : '';
	while (commaRx.test(x1)) {
		x1 = x1.replace(commaRx, '$1' + ',' + '$2');
	}
	return x1 + x2;
}

function formatDuration(seconds) {
  seconds = Math.round(seconds);
  var date = new Date(null);
  date.setSeconds(seconds); // specify value for SECONDS here
  return date.toISOString().substr(14, 5);
}

function formatPercentage(p) {
   return Math.round(p*100) + "%";
}

var Client = {};
Client.request = function(url, args) {
  args = args || {};
  args.credentials = 'same-origin'
  args.headers = args.headers || {};
  args.headers['Accept'] = 'application/json';

  if( args.method && args.method === 'POST') {
    args.headers['Content-Type'] = 'application/json';

    if(args.data) {
      if( typeof(args.data) !== "string") {
        args.data = JSON.stringify(args.data)
      }
      args.body = args.data
      delete args.data
    }
  }

  // trim leading slash from URL
  url = (url[0] === '/') ? url.substring(1) : url;

  return window.fetch(`api/${url}`, args)
    .then(handleRequestErrors)
    .then(parseJSON)
    .then(parseData)
}

function handleRequestErrors(r) {
  // if response is not JSON (eg timeout), throw a generic error
  if (! r.ok && r.headers.get("Content-Type") !== "application/json") {
    throw { code: "request_error", message: "An error occurred" }
  }

  return r
}

function parseJSON(r) {
  return r.json()
}

function parseData(d) {

  // if JSON response contains an Error property, use that as error code
  // Message is generic here, so that individual components can set their 
  // own specific messages based on the error code
  if(d.Error) {
    throw { code: d.Error, message: "An error occurred" }
  }

  return d.Data
}
