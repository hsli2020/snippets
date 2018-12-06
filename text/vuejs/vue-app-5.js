http://tutorialzine.com/2016/03/5-practical-examples-for-learning-vue-js/

Example 1: Nav Menu
===================

<div id="main">
    <!-- The navigation menu will get the value of the "active" variable as a class. -->
    <!-- To stops the page from jumping when a link is clicked 
         we use the "prevent" modifier (short for preventDefault). -->

	<nav v-bind:class="active" v-on:click.prevent>

		<!-- When a link in the menu is clicked, we call the makeActive method, 
        defined in the JavaScript Vue instance. It will change the value of "active". -->

		<a href="#" class="home"     v-on:click="makeActive('home')">Home</a>
		<a href="#" class="projects" v-on:click="makeActive('projects')">Projects</a>
		<a href="#" class="services" v-on:click="makeActive('services')">Services</a>
		<a href="#" class="contact"  v-on:click="makeActive('contact')">Contact</a>
	</nav>

 	<!-- The mustache expression will be replaced with the value of "active".
 		 It will automatically update to reflect any changes. -->

	<p>You chose <b>{{active}}</b></p>
</div>

// Creating a new Vue instance and pass in an options object.
var demo = new Vue({
	
	// A DOM element to mount our view model.
	el: '#main',

    // This is the model. Define properties and give them initial values.
	data: {
		active: 'home'
	},

	// Functions we will be using.
	methods: {
		makeActive: function(item){
			// When a model is changed, the view will be automatically updated.
			this.active = item;
		}
	}
});

Example 2: Inline Editor
========================
<!-- v-cloak hides any un-compiled data bindings until the Vue instance is ready. -->
<!-- When the element is clicked the hideTooltp() method is called. -->

<div id="main" v-cloak v-on:click="hideTooltip" >

    <!-- This is the tooltip. 
         v-on:clock.stop is an event handler for clicks, with a modifier that stops event propagation.
         v-if makes sure the tooltip is shown only when the "showtooltip" variable is truthful -->

    <div class="tooltip" v-on:click.stop v-if="show_tooltip">

        <!-- v-model binds the contents of the text field with the "text_content" model.
         Any changes to the text field will automatically update the value, and
         all other bindings on the page that depend on it.  -->

        <input type="text" v-model="text_content" />
    </div>

    <!-- When the paragraph is clicked, call the "toggleTooltip" method and stop event propagation. -->

    <!-- The mustache expression will be replaced with the value of "text_content".
         It will automatically update to reflect any changes to that variable. -->

    <p v-on:click.stop="toggleTooltip">{{text_content}}</p>

</div>

// Creating a new Vue instance and pass in an options object.
var demo = new Vue({

    // A DOM element to mount our view model.
    el: '#main',

    // Define properties and give them initial values.
    data: {
        show_tooltip: false,
        text_content: 'Edit me.'
    },

    // Functions we will be using.
    methods: {
        hideTooltip: function(){
            // When a model is changed, the view will be automatically updated.
            this.show_tooltip = false;
        },
        toggleTooltip: function(){
            this.show_tooltip = !this.show_tooltip;
        }
    }
})

Example 3: Order Form
=====================
<!-- v-cloak hides any un-compiled data bindings until the Vue instance is ready. -->

<form id="main" v-cloak>

	<h1>Services</h1>

	<ul>
		<!-- Loop through the services array, assign a click handler, and set or
			 remove the "active" css class if needed -->
		<li v-for="service in services"
            v-on:click="toggleActive(service)"
            v-bind:class="{ 'active': service.active}">
			<!-- Display the name and price for every entry in the array .
                Vue.js has a built in currency filter for formatting the price -->
			{{service.name}} <span>{{service.price | currency}}</span>
		</li>
	</ul>

	<div class="total">
		<!-- Calculate the total price of all chosen services. Format it as currency. -->
		Total: <span>{{total() | currency}}</span>
	</div>

</form>

var demo = new Vue({
    el: '#main',
    data: {
    	// Define the model properties. The view will loop
        // through the services array and genreate a li
        // element for every one of its items.
        services: [
        	{ name: 'Web Development', price: 300, active:true  },
            { name: 'Design',          price: 400, active:false },
            { name: 'Integration',     price: 250, active:false },
            { name: 'Training',        price: 220, active:false }
        ]
    },
    methods: {
    	toggleActive: function(s){
            s.active = !s.active;
    	},

    	total: function(){
        	var total = 0;

        	this.services.forEach(function(s){
        		if (s.active){
        			total+= s.price;
        		}
        	});

    	   return total;
        }
    }
});

Example 4: Instant Search
=========================

<form id="main" v-cloak>

    <div class="bar">
        <!-- Create a binding between the searchString model and the text field -->
        <input type="text" v-model="searchString" placeholder="Enter your search terms" />
    </div>

    <ul>
        <!-- Render a li element for every entry in the items array. Notice
             the custom search filter "searchFor". It takes the value of the
             searchString model as an argument. -->
             
        <li v-for="i in articles | searchFor searchString">
            <a v-bind:href="i.url"><img v-bind:src="i.image" /></a>
            <p>{{i.title}}</p>
        </li>
    </ul>

</form>

// Define a custom filter called "searchFor". 

    Vue.filter('searchFor', function (value, searchString) {

        // The first parameter to this function is the data that is to be filtered.
        // The second is the string we will be searching for.

        var result = [];

        if(!searchString){
            return value;
        }

        searchString = searchString.trim().toLowerCase();

        result = value.filter(function(item){
            if(item.title.toLowerCase().indexOf(searchString) !== -1){
                return item;
            }
        })

        // Return an array with the filtered data.
        return result;
    })

    var demo = new Vue({
        el: '#main',
        data: {
            searchString: "",

            // The data model. These items would normally be requested via AJAX,
            // but are hardcoded here for simplicity.

            articles: [
                {
                    "title": "What You Need To Know About CSS Variables",
                    "url": "http://tutorialzine.com/2016/03/what-you-need-to-know-about-css-variables/",
                    "image": "http://cdn.tutorialzine.com/wp-content/uploads/2016/03/css-variables-100x100.jpg"
                }, {
                    "title": "Freebie: 4 Great Looking Pricing Tables",
                    "url": "http://tutorialzine.com/2016/02/freebie-4-great-looking-pricing-tables/",
                    "image": "http://cdn.tutorialzine.com/wp-content/uploads/2016/02/great-looking-pricing-tables-100x100.jpg"
                }, {
                    "title": "20 Interesting JavaScript and CSS Libraries for February 2016",
                    "url": "http://tutorialzine.com/2016/02/20-interesting-javascript-and-css-libraries-for-february-2016/",
                    "image": "http://cdn.tutorialzine.com/wp-content/uploads/2016/02/interesting-resources-february-100x100.jpg"
                }, {
                    "title": "Quick Tip: The Easiest Way To Make Responsive Headers",
                    "url": "http://tutorialzine.com/2016/02/quick-tip-easiest-way-to-make-responsive-headers/",
                    "image": "http://cdn.tutorialzine.com/wp-content/uploads/2016/02/quick-tip-responsive-headers-100x100.png"
                }, {
                    "title": "Learn SQL In 20 Minutes",
                    "url": "http://tutorialzine.com/2016/01/learn-sql-in-20-minutes/",
                    "image": "http://cdn.tutorialzine.com/wp-content/uploads/2016/01/learn-sql-20-minutes-100x100.png"
                }, {
                    "title": "Creating Your First Desktop App With HTML, JS and Electron",
                    "url": "http://tutorialzine.com/2015/12/creating-your-first-desktop-app-with-html-js-and-electron/",
                    "image": "http://cdn.tutorialzine.com/wp-content/uploads/2015/12/creating-your-first-desktop-app-with-electron-100x100.png"
                }
            ]
        }
    });

Example 5: Switchable Grid/List
===============================
<form id="main" v-cloak>

	<div class="bar">

		<!-- These two buttons switch the layout variable,
			 which causes the correct UL to be shown. -->

		<a class="list-icon" v-bind:class="{ 'active': layout == 'list'}" v-on:click="layout = 'list'"></a>
		<a class="grid-icon" v-bind:class="{ 'active': layout == 'grid'}" v-on:click="layout = 'grid'"></a>
	</div>

	<!-- We have two layouts. We choose which one to show depending on the "layout" binding -->

	<ul v-if="layout == 'grid'" class="grid">
		<!-- A view with big photos and no text -->
		<li v-for="a in articles">
			<a v-bind:href="a.url" target="_blank"><img v-bind:src="a.image.large" /></a>
		</li>
	</ul>

	<ul v-if="layout == 'list'" class="list">
		<!-- A compact view smaller photos and titles -->
		<li v-for="a in articles">
			<a v-bind:href="a.url" target="_blank"><img v-bind:src="a.image.small" /></a>
			<p>{{a.title}}</p>
		</li>
	</ul>

</form>

var demo = new Vue({
	el: '#main',
	data: {
        // The layout mode, possible values are "grid" or "list".
		layout: 'grid',

        articles: [{
            "title": "What You Need To Know About CSS Variables",
            "url": "http://tutorialzine.com/2016/03/what-you-need-to-know-about-css-variables/",
            "image": {
                "large": "http://cdn.tutorialzine.com/wp-content/uploads/2016/03/css-variables.jpg",
                "small": "http://cdn.tutorialzine.com/wp-content/uploads/2016/03/css-variables-150x150.jpg"
            }
        }, {
            "title": "Freebie: 4 Great Looking Pricing Tables",
            "url": "http://tutorialzine.com/2016/02/freebie-4-great-looking-pricing-tables/",
            "image": {
                "large": "http://cdn.tutorialzine.com/wp-content/uploads/2016/02/great-looking-pricing-tables.jpg",
                "small": "http://cdn.tutorialzine.com/wp-content/uploads/2016/02/great-looking-pricing-tables-150x150.jpg"
            }
        }, {
            "title": "20 Interesting JavaScript and CSS Libraries for February 2016",
            "url": "http://tutorialzine.com/2016/02/20-interesting-javascript-and-css-libraries-for-february-2016/",
            "image": {
                "large": "http://cdn.tutorialzine.com/wp-content/uploads/2016/02/interesting-resources-february.jpg",
                "small": "http://cdn.tutorialzine.com/wp-content/uploads/2016/02/interesting-resources-february-150x150.jpg"
            }
        }, {
            "title": "Quick Tip: The Easiest Way To Make Responsive Headers",
            "url": "http://tutorialzine.com/2016/02/quick-tip-easiest-way-to-make-responsive-headers/",
            "image": {
                "large": "http://cdn.tutorialzine.com/wp-content/uploads/2016/02/quick-tip-responsive-headers.png",
                "small": "http://cdn.tutorialzine.com/wp-content/uploads/2016/02/quick-tip-responsive-headers-150x150.png"
            }
        }, {
            "title": "Learn SQL In 20 Minutes",
            "url": "http://tutorialzine.com/2016/01/learn-sql-in-20-minutes/",
            "image": {
                "large": "http://cdn.tutorialzine.com/wp-content/uploads/2016/01/learn-sql-20-minutes.png",
                "small": "http://cdn.tutorialzine.com/wp-content/uploads/2016/01/learn-sql-20-minutes-150x150.png"
            }
        }, {
            "title": "Creating Your First Desktop App With HTML, JS and Electron",
            "url": "http://tutorialzine.com/2015/12/creating-your-first-desktop-app-with-html-js-and-electron/",
            "image": {
                "large": "http://cdn.tutorialzine.com/wp-content/uploads/2015/12/creating-your-first-desktop-app-with-electron.png",
                "small": "http://cdn.tutorialzine.com/wp-content/uploads/2015/12/creating-your-first-desktop-app-with-electron-150x150.png"
            }
        }]
	}
});
