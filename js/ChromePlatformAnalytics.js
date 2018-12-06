// Chrome Platform Analytics
// https://github.com/GoogleChrome/chrome-platform-analytics/wiki

var ChromePlatformAnalytics = (function(){

  var service;
  var tracker;
  var userPrompted = !!localStorage.cpaUserPrompted;

  var init = function(appName, ua){
    initMessaging();
    service = analytics.getService(appName);
    tracker = service.getTracker(ua);
  };

  var initMessaging = function(){
    chrome.runtime.onMessage.addListener(
      function(request, sender, sendResponse) {
        if (typeof request !== 'object') return;
        if(request.cmd === 'cpa.sendAppView') {
          sendAppView(request.data);
        }
        else if (request.cmd === 'cpa.sendEvent') {
          sendEvent(request.data);
        }
        else if (request.cmd === 'cpa.toggle') {
          toggle(request.data);
        }
      }
    );
  };

  var setUserPrompted = function(){
    userPrompted = true;
    localStorage.cpaUserPrompted = userPrompted;
  };

  var getUserPrompted = function(){
    return userPrompted;
  };

  var sendAppView = function(view){
    if (!userPrompted) return;
    tracker.sendAppView(view);
  };

  var sendEvent = function(data){
    if (!userPrompted) return;
    tracker.sendEvent.apply(tracker, data);
  };

  var toggle = function(state){
    setUserPrompted();
    localStorage.trackingPermitted = state;
    service.getConfig().addCallback(function(config) {
      config.setTrackingPermitted(!!state);
    });
  };

  return {
    init: init,
    setUserPrompted: setUserPrompted,
    getUserPrompted: getUserPrompted,
    sendAppView: sendAppView,
    sendEvent: sendEvent,
    toggle: toggle
  };

})();

ChromePlatformAnalytics.init('us_app', 'UA-88464834-5');
