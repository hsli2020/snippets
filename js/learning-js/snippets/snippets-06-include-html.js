<div class="include" data-include="header"></div>
<div class="include" data-include="footer"></div>

<script>
  $(function(){
    var includes = $('.include');
    jQuery.each(includes, function(){
      var file = 'views/' + $(this).data('include') + '.html';
      $(this).load(file);
    });
  });
</script>


window.onload = function() {
	var elements = document.getElementsByTagName('*'),
		i;
	for (i in elements) {
		if (elements[i].hasAttribute && elements[i].hasAttribute('data-include')) {
			fragment(elements[i], elements[i].getAttribute('data-include'));
		}
	}
	function fragment(el, url) {
		var localTest = /^(?:file):/,
			xmlhttp = new XMLHttpRequest(),
			status = 0;

		xmlhttp.onreadystatechange = function() {
			/* if we are on a local protocol, and we have response text, we'll assume
 *  				things were sucessful */
			if (xmlhttp.readyState == 4) {
				status = xmlhttp.status;
			}
			if (localTest.test(location.href) && xmlhttp.responseText) {
				status = 200;
			}
			if (xmlhttp.readyState == 4 && status == 200) {
				el.outerHTML = xmlhttp.responseText;
			}
		}

		try { 
			xmlhttp.open("GET", url, true);
			xmlhttp.send();
		} catch(err) {
			/* todo catch error */
		}
	}
}


<iframe src="/path/to/file.html" seamless></iframe>
iframe[seamless] {
    border: none;
}


<div class="banner" title="banner.html"></div>
$(".banner").each(function(){
    var inc=$(this);
    $.get(inc.attr("title"), function(data){
        inc.replaceWith(data);
    });
});
