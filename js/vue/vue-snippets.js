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


