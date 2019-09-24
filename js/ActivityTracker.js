// https://codesandbox.io/s/activity-tracker-vanilla-js-mvlnc
// constants.js
export const INACTIVE_USER_TIME_THRESHOLD = 2000;
export const USER_ACTIVITY_THROTTLER_TIME = 1000;

// index.js
import { INACTIVE_USER_TIME_THRESHOLD, USER_ACTIVITY_THROTTLER_TIME } from "./constants";

document.getElementById("app").innerHTML = `<h1>User is inactive = false</h1>`;

let userActivityTimeout = null;
let userActivityThrottlerTimeout = null;
let isInactive = false;

activateActivityTracker();

function activateActivityTracker() {
  window.addEventListener("mousemove", userActivityThrottler);
  window.addEventListener("scroll", userActivityThrottler);
  window.addEventListener("keydown", userActivityThrottler);
  window.addEventListener("resize", userActivityThrottler);
  window.addEventListener("beforeunload", deactivateActivityTracker);
}

function deactivateActivityTracker() {
  window.removeEventListener("mousemove", userActivityThrottler);
  window.removeEventListener("scroll", userActivityThrottler);
  window.removeEventListener("keydown", userActivityThrottler);
  window.removeEventListener("resize", userActivityThrottler);
  window.removeEventListener("beforeunload", deactivateActivityTracker);
}

function resetUserActivityTimeout() {
  clearTimeout(userActivityTimeout);

  userActivityTimeout = setTimeout(() => {
    userActivityThrottler();
    inactiveUserAction();
  }, INACTIVE_USER_TIME_THRESHOLD);
}

function userActivityThrottler() {
  if (isInactive) {
    isInactive = false;
    document.getElementById("app").innerHTML = `<h1>User is inactive = false</h1>`;
    resetUserActivityTimeout();
  }

  if (!userActivityThrottlerTimeout) {
    userActivityThrottlerTimeout = setTimeout(() => {
      resetUserActivityTimeout();
      clearTimeout(userActivityThrottlerTimeout);
    }, USER_ACTIVITY_THROTTLER_TIME);
  }
}

function inactiveUserAction() {
  isInactive = true;
  document.getElementById("app").innerHTML = `<h1>User is inactive = true</h1>`;
}
