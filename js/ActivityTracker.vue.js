// https://codesandbox.io/s/activity-tracker-mbgix
// App.vue
<template>
  <div id="app">
    <div class="content" :class="{ 'is-faded': isInactive}">
      <img class="logo" width="25%" src="./assets/logo.png">
      <p>User is inactive = {{ isInactive }}</p>
    </div>
  </div>
</template>

<script>
import {
  INACTIVE_USER_TIME_THRESHOLD,
  USER_ACTIVITY_THROTTLER_TIME
} from "@/constants";

export default {
  name: "App",

  data() {
    return {
      isInactive: false,
      userActivityThrottlerTimeout: null,
      userActivityTimeout: null
    };
  },

  methods: {
    activateActivityTracker() {
      window.addEventListener("mousemove", this.userActivityThrottler);
      window.addEventListener("scroll", this.userActivityThrottler);
      window.addEventListener("keydown", this.userActivityThrottler);
      window.addEventListener("resize", this.userActivityThrottler);
    },

    deactivateActivityTracker() {
      window.removeEventListener("mousemove", this.userActivityThrottler);
      window.removeEventListener("scroll", this.userActivityThrottler);
      window.removeEventListener("keydown", this.userActivityThrottler);
      window.removeEventListener("resize", this.userActivityThrottler);
    },

    resetUserActivityTimeout() {
      clearTimeout(this.userActivityTimeout);

      this.userActivityTimeout = setTimeout(() => {
        this.userActivityThrottler();
        this.inactiveUserAction();
      }, INACTIVE_USER_TIME_THRESHOLD);
    },

    userActivityThrottler() {
      if (this.isInactive) {
        this.isInactive = false;
      }

      if (!this.userActivityThrottlerTimeout) {
        this.userActivityThrottlerTimeout = setTimeout(() => {
          this.resetUserActivityTimeout();
          clearTimeout(this.userActivityThrottlerTimeout);
          this.userActivityThrottlerTimeout = null;
        }, USER_ACTIVITY_THROTTLER_TIME);
      }
    },

    inactiveUserAction() {
      this.isInactive = true;
    }
  },

  beforeMount() {
    this.activateActivityTracker();
  },

  beforeDestroy() {
    this.deactivateActivityTracker();
    clearTimeout(this.userActivityTimeout);
    clearTimeout(this.userActivityThrottlerTimeout);
  }
};
</script>

<style lang="scss" scoped>
.content {
  font-family: sans-serif;
  font-weight: 100;
  height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.is-faded {
  opacity: 0.25;
  transition: opacity 500ms ease-in-out;
}
</style>

// constants.js
export const INACTIVE_USER_TIME_THRESHOLD = 2000;
export const USER_ACTIVITY_THROTTLER_TIME = 1000;

// main.js
import Vue from "vue";
import App from "./App.vue";

Vue.config.productionTip = false;

new Vue({
  render: h => h(App)
}).$mount("#app");
