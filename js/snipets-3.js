function loadJS(url) {
    let script = document.createElement('script');
    script.src = url;
    document.body.appendChild(script);
}

function loadJSAsync(url) {
    let script = document.createElement('script');
    script.src = url;
    script.async = true;
    document.body.appendChild(script);
}

//----------------------------------------------------------
// one$ / all$ / new$

function $one(sel) { return document.querySelector(sel) }
function $all(sel) { return document.querySelectorAll(sel) }

function $new(element, attribute, inner) {
  if (typeof(element) === "undefined") {
    return false;
  }
  if (typeof(inner) === "undefined") {
    inner = "";
  }
  var el = document.createElement(element);
  if (typeof(attribute) === 'object') {
    for (var key in attribute) {
      el.setAttribute(key, attribute[key]);
    }
  }
  if (!Array.isArray(inner)) {
    inner = [inner];
  }
  for (var k = 0; k < inner.length; k++) {
    if (inner[k].tagName) {
      el.appendChild(inner[k]);
    } else {
      el.appendChild(document.createTextNode(inner[k]));
    }
  }
  return el;
}

$new("a", {"href":"http://google.com","style":"color:#FFF;background:#333;"}, "google");
// <a href="http://google.com" style="color:#FFF;background:#333;">google</a>

var google = $new("a", {"href":"http://google.com"}, "google"),
    youtube = $new("a", {"href":"http://youtube.com"}, "youtube"),
    facebook = $new("a", {"href":"http://facebook.com"}, "facebook"),
    links_conteiner = $new("div", {"id":"links"}, [google,youtube,facebook]);

//<div id="links">
//    <a href="http://google.com">google</a>
//    <a href="http://youtube.com">youtube</a>
//    <a href="http://facebook.com">facebook</a>
//</div>

//----------------------------------------------------------

// p.before(h1)
// Element.before(node1, node2, ... nodeN)
// Element.before(str1, str2, ... strN)
// Element.insertAdjacentHTML('beforebegin', '<button>Start</button>');

function insertBefore(newNode, existingNode) {
    existingNode.parentNode.insertBefore(newNode, existingNode);
}

// h1.after(p)
// Element.after(node1, node2, ... nodeN)
// Element.after(str1, str2, ... strN)
// Element.insertAdjacentHTML('afterend', '<button>Stop</button>');

function insertAfter(newNode, existingNode) {
    existingNode.parentNode.insertBefore(newNode, existingNode.nextSibling);
}
