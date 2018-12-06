// Vue对象从创建前到最后死亡，在各个阶段状态改变的时候，都提供了一个钩子方法，
// 你可以注册一下，如果你希望在特定状态改变的时候干点什么的话。

//根对象
var vm = new Vue({
    //挂载点
    el: document.getElementById('root'),

    //2.使用刚刚的路由配置
    router: routerObj,

    //启动组件
    render: function (createElement) {
        return createElement(App);
    },

    //下面是Vue对象的几种状态
    beforeCreate: function () {
        console.debug('Vue对象目前状态：beforeCreate');
    },
    created: function () {
        console.debug('Vue对象目前状态：created');
    },
    beforeMount: function () {
        console.debug('Vue对象目前状态：beforeMount');
    },
    mounted: function () {
        console.debug('Vue对象目前状态：mounted');
    },
    beforeUpdate: function () {
        console.debug('Vue对象目前状态：beforeUpdate');
    },
    updated: function () {
        console.debug('Vue对象目前状态：updated');
    },
    beforeDestroy: function () {
        console.debug('Vue对象目前状态：beforeDestroy');
    },
    destroyed: function () {
        console.debug('Vue对象目前状态：destroyed');
    }
});

// 我们可以给组件的props属性添加验证，当传入的数据不符合要求时，Vue会发出警告。

Vue.component('example', {
  props: {
    // 基础类型检测 (`null` 意思是任何类型都可以)
    propA: Number,
    // 多种类型
    propB: [String, Number],
    // 必传且是字符串
    propC: {
      type: String,
      required: true
    },
    // 数字，有默认值
    propD: {
      type: Number,
      default: 100
    },
    // 数组/对象的默认值应当由一个工厂函数返回
    propE: {
      type: Object,
      default: function () {
        return { message: 'hello' }
      }
    },
    // 自定义验证函数
    propF: {
      validator: function (value) {
        return value > 10
      }
    }
  }
})

// type 可以是下面原生构造器：
// 
//     String
//     Number
//     Boolean
//     Function
//     Object
//     Array
//     Symbol
// 
// type 也可以是一个自定义构造器函数，使用 instanceof 检测。

<!-- 阻止单击事件冒泡 -->
<a v-on:click.stop="doThis"></a>

<!-- 提交事件不再重载页面 -->
<form v-on:submit.prevent="onSubmit"></form>

<!-- 修饰符可以串联  -->
<a v-on:click.stop.prevent="doThat"></a>

<!-- 只有修饰符 -->
<form v-on:submit.prevent></form>

<!-- 添加事件侦听器时使用事件捕获模式 -->
<div v-on:click.capture="doThis">...</div>

<!-- 只当事件在该元素本身（而不是子元素）触发时触发回调 -->
<div v-on:click.self="doThat">...</div>

<!-- click 事件只能点击一次，2.1.4版本新增 -->
<a v-on:click.once="doThis"></a>

<!-- 只有在 keyCode 是 13 时调用 vm.submit() -->
<input v-on:keyup.13="submit">

    <!-- 同上 -->
    <input v-on:keyup.enter="submit">

    <!-- 缩写语法 -->
    <input @keyup.enter="submit">

全部的按键别名：
    .enter
    .tab
    .delete (捕获 "删除" 和 "退格" 键)
    .esc
    .space
    .up
    .down
    .left
    .right
    .ctrl
    .alt
    .shift
    .meta

实例

<!-- Alt + C -->
<input @keyup.alt.67="clear">

<!-- Ctrl + Click -->
<div @click.ctrl="doSomething">Do something</div>
