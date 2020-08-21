// define Vue component in HTML file
<script type="text/x-template" id="search-box-template">
    <form>
        <input type="text">
    </form>
</script>

<script>
Vue.component('search-box', {
    props: [ "query", "items", "results" ],
    template: "#search-box-template"
})
</script>

// define Vue component in Javascript file
const search_box_template = `
    <form>
        <input type="text">
    </form>
`;

Vue.component('search-box', {
    props: [ "query", "items", "results" ],
    template: search_box_template,
    methods: {
        refresh: function() { }
        update: function() { }
    },
    watch: {
        activeKey: function(newValue) {
            this.editType = newValue.type;
            this.refresh();
        }
    }
})

// ----------------------------------------------------------
// 模版中不需要用 this
// 代码中必须用 this, eg: this.items, this.refresh()

<div id="app">
    <search-box ref="searchbox" @search-result="showSearchResults"></search-box>

    <search-results :query="q" :results="searchResults" :item="key"
        @select-item="selectItem"></search-box>

    <cmd-view :command="cmdText" :value="cmdOutput"></cmd-view>
    <key-view :name="key" :type="keyType" :value="keyValue" @edit-key="editKey"></key-view>
    <key-edit :active-key="activeKey" @run-command="runCommand"></key-edit>
</div>

Vue.prototype.$http = axios;
window.bus = new Vue({});  // 用于组件之间通讯

this.$http.get('localhost/api/url').then(res => { // 箭头函数, 没有this问题
    this.$refs.searchbox.command = '';
    bus.$emit('update-state');  // 通知别的组件
    this.$emit('search', { name, type, value });  // 通知自己/父组件
});

this.$http.get('localhost/api/url').then(function(res) { // 普通函数，有this问题
    this.$emit('search', { keyword });
}.bind(this));  // 必须绑定 this

bus.$on('update-state', function() {  // 接收通知并处理
    if (!history.pushState) return;
    history.pushState(null, url);
});

window.onpopstate = function(e) { /* do sth */ }
