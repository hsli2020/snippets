// index.js

// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from "vue";
import App from "./App";

Vue.config.productionTip = false;

/* eslint-disable no-new */
new Vue({
  el: "#app",
  components: { App },
  template: "<App/>"
});

// App.vue

<template>
  <div id="app">
    <UsrMsg @inputData="updateMessage" />
    <Results :msg="childData" />
  </div>
</template>

<script>
import UsrMsg from "./components/UsrMsg";
import Results from "./components/Results";

export default {
  name: "App",
  components: {
    UsrMsg,
    Results
  },
  data: function() {
    return {
      childData: ""
    };
  },
  methods: {
    updateMessage(variable) {
      this.childData = variable;
    }
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

// components/usrMsg.vue

<template>
  <div class="hello">
    <input
      placeholder="Enter Text Here"
      v-model="tempMessage"
      @keyup.enter="submit"
    />
  </div>
</template>

<script>
export default {
  name: "UsrMsg",
  data() {
    return {
      tempMessage: ""
    };
  },
  methods: {
    submit: function() {
      this.$emit("inputData", this.tempMessage);
      this.tempMessage = "";
    }
  }
};
</script>

// components/results.vue

<template>
  <div>
    <li v-for="(message, index) in messageList" :item="message" :key="index">
      {{ message }}
    </li>
  </div>
</template>

<script>
export default {
  name: "Results",
  props: {
    msg: {
      type: String
    }
  },
  data: function() {
    return {
      messageList: []
    };
  },
  watch: {
    msg: function() {
      this.messageList.push(this.msg);
    }
  }
};
</script>
