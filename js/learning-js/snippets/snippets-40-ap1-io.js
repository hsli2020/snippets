<script type="text/javascript">

var isMobile = {
    Android: function() {
        return navigator.userAgent.match(/Android/i);
    },
    BlackBerry: function() {
        return navigator.userAgent.match(/BlackBerry/i);
    },
    iOS: function() {
        return navigator.userAgent.match(/iPhone|iPad|iPod/i);
    },
    Opera: function() {
        return navigator.userAgent.match(/Opera Mini/i);
    },
    Windows: function() {
        return navigator.userAgent.match(/IEMobile/i);
    },
    any: function() {
        return (isMobile.Android() || isMobile.BlackBerry() || 
                isMobile.iOS() || isMobile.Opera() || isMobile.Windows());
    }
};

$(document).ready(function() {

	resizeLogo();

    // if( !(isMobile.any()) ){  }
    
	setTimeout(function () {
		$(".blackMask").fadeOut(550);
		resizeLogo();
	}, 900);

})

window.onresize = function(event) {
	resizeLogo();
};

function resizeLogo() {
	var parentHeight = $(".fullscreen-bg").height();
    var parentWidth = $(".fullscreen-bg").width();
	$(".videoLogo").css("height", parentHeight +"px");
    $(".videoLogo").css("width", parentWidth +"px");
}

</script>
