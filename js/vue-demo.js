new Vue({
    el: '#app',
    delimiters: ['{:', ':}'],
    components: {
        autocomplete: Vue2Autocomplete
    },
    data: {
        nextPage: 0,
        showSpinner: 1,
        blogs: [],
        picked: [],
    },
    methods: {
        select(id) {
            if (this.picked.indexOf(id) == -1) {
                this.picked.push(id)
            } else {
                this.picked.splice(this.picked.indexOf(id), 1)
            }
        }
        toggle: function() {
            this.menuOpen = !this.menuOpen;
        },
        itemSelected: function(data) {
            ga('send', 'pageview', data.searchUrl);
            window.location.href = data.url;
        },
        processJsonData: function(json) {
            return json.data;
        }
    },
    watch: {
        picked(value) {
            if (value.length > 2) {
                this.picked.shift()
            }
        }
    },
    created: function() {
        this.$http.get('/api/blog.json').then(function(data) {
            if (data.body.meta.pagination.total_pages > 1)
                this.nextPage = 2;
            this.showSpinner = 0;
        });
    },
    computed: {
        message() {
            if (this.picked.length < 2) {
                return '(check what you need)'
            } else if (this.picked.indexOf(0) === -1) {
                return '(cheap & quality will take a long time)'
            } else if (this.picked.indexOf(1) === -1) {
                return '(fast & quality costs money)'
            } else if (this.picked.indexOf(2) === -1) {
                return '(fast & cheap can never be good)'
            }
        }
    }
});
