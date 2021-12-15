export function camelToDash(str) {
  return str.replace(/([a-z])([A-Z])/g, '$1-$2').toLowerCase();
}

export function pascalToDash(str) {
  str = str[0].toLowerCase() + str.slice(1);
  return camelToDash(str);
}

export function dispatch(host, eventType, options = {}) {
  return host.dispatchEvent(new CustomEvent(eventType, { bubbles: false, ...options }));
}

export function shadyCSS(fn, fallback) {
  const shady = window.ShadyCSS;

  /* istanbul ignore next */
  if (shady && !shady.nativeShadow) {
    return fn(shady);
  }

  return fallback;
}

export function stringifyElement(element) {
  const tagName = String(element.tagName).toLowerCase();
  return `<${tagName}>`;
}

export const IS_IE = 'ActiveXObject' in window;

// ==========

// precompute "00" to "FF"
var decToHex = [];
for (var i = 0; i < 16; i++) {
  for (var j = 0; j < 16; j++) {
    decToHex[i * 16 + j] = i.toString(16) + j.toString(16);
  }
}

function clamp(v, min, max) {
  return Math.min(max, Math.max(min, v));
}

function percent(s) {
  return parseFloat(s) / 100;
}

function getRgbHslContent(styleString) {
  var start = styleString.indexOf('(', 3);
  var end = styleString.indexOf(')', start + 1);
  var parts = styleString.substring(start + 1, end).split(',');
  // add alpha if needed
  if (parts.length != 4 || styleString.charAt(3) != 'a') {
    parts[3] = 1;
  }
  return parts;
}

function hslToRgb(parts){
  var r, g, b, h, s, l;
  h = parseFloat(parts[0]) / 360 % 360;
  if (h < 0)
    h++;
  s = clamp(percent(parts[1]), 0, 1);
  l = clamp(percent(parts[2]), 0, 1);
  if (s == 0) {
    r = g = b = l; // achromatic
  } else {
    var q = l < 0.5 ? l * (1 + s) : l + s - l * s;
    var p = 2 * l - q;
    r = hueToRgb(p, q, h + 1 / 3);
    g = hueToRgb(p, q, h);
    b = hueToRgb(p, q, h - 1 / 3);
  }

  return '#' + decToHex[Math.floor(r * 255)] +
      decToHex[Math.floor(g * 255)] +
      decToHex[Math.floor(b * 255)];
}

function hueToRgb(m1, m2, h) {
  if (h < 0)
    h++;
  if (h > 1)
    h--;

  if (6 * h < 1)
    return m1 + (m2 - m1) * 6 * h;
  else if (2 * h < 1)
    return m2;
  else if (3 * h < 2)
    return m1 + (m2 - m1) * (2 / 3 - h) * 6;
  else
    return m1;
}
