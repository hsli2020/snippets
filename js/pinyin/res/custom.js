jQuery(document).ready(function(){

	jQuery('#menus .menu1 ul li').each(function(index, value) {

		content = menu[index+1];
		jQuery(this).attr('id', 'menu' + index);

		if(index != 7){
			var sub = '<ul class="submenu">';
			var n = content.length;
			for(var i=0; i<n; i++) {
				sub += '<li><a href="' + content[i][1] + '">' + content[i][0] + '</a></li>'; 
			}
			sub += '</ul>';

			jQuery(this).append(sub);
		}

	});

});

jQuery(function(){

	jQuery('#menus .menu1').slicknav({
		label: 'MENU',
		duration: 1000,
		easingOpen: "easeOutBounce", //available with jQuery UI
	});

	jQuery('#menus').prepend(jQuery('.slicknav_menu'));

});