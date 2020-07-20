'use strict';

var browser = browser;

(function(browser) {
  var browser = browser;
  var promises = true;

  browser.storage.local.get(name).then(function(items) {
      resolve(items[name]);
  });

  browser.storage.local.set(settingObject);
  browser.storage.local.get(settingNames).then(ensureSettings);

  browser.storage.local.remove(host).then(function() {
      browser.tabs.reload({bypassCache: true});
  });

  browser.storage.local.set(item).then(function() {
      browser.tabs.reload({bypassCache: true});
  });

  browser.tabs.query({active:true}).then(function(tab) {
      toggleJSState(tab[0]);
  });

  browser.runtime.openOptionsPage();

  browser.commands.onCommand.addListener(commandListener);
  browser.commands.onCommand.removeListener(commandListener);

  browser.webRequest.onHeadersReceived.addListener(
    addHeader,
    {
      urls: listenUrls,
      types: ['main_frame', 'sub_frame']
    },
    ['blocking', 'responseHeaders']
  );

  browser.tabs.onUpdated.addListener(function(tabId, changeInfo, tab) {
  });

  if (typeof browser.browserAction.setIcon !== 'undefined') {
      browser.browserAction.setIcon(getIcon(jsEnabled, tabId, url));
  }
  browser.browserAction.setTitle({
      title: (jsEnabled ? 'Disable' : 'Enable') + ' Javascript',
      tabId: tabId
  });
  browser.browserAction.onClicked.addListener(toggleJSState);

  browser.tabs.executeScript(
      tabId, { file: '/background/content.js' }
  );

  browser.tabs.onCreated.addListener(function(tab) {
      var url = tab.url;
      var host = new URL(url).hostname;
  });

  browser.menus.create({
      id: 'settings',
      title: browser.i18n.getMessage('menuItemSettings'),
      contexts: ['browser_action']
  });
  browser.menus.create({
      id: 'toggle-js',
      title: 'Toggle JavaScript',
      contexts: ['page']
  });

  browser.menus.remove('toggle-js');

  browser.menus.onClicked.addListener(function(info, tab) {
    switch (info.menuItemId) {
    }
  });

  browser.runtime.onMessage.addListener(function(request, sender, sendResponse) {
    switch (request.type) {
    }
  });

  browser.runtime.onInstalled.addListener(function(details) {
    if (promises) {
      browser.storage.local.get('setting-version').then(handleUpdate);
    } else {
      browser.storage.local.get('setting-version', handleUpdate);
    }
  });

  function handleUpdate(settingValues) {
    var manifest = browser.runtime.getManifest();
    var prevVersion = '1.0.0';
    var thisVersion = manifest.version;
    var anyChange = false;
    var majorChange = false;
    var minorChange = false;
    var patchChange = false;

    if (settingValues.hasOwnProperty('setting-version')) {
      prevVersion = settingValues['setting-version'];
    }

    var prevParts = prevVersion.split('.');
    var thisParts = thisVersion.split('.');

    if (prevParts[0] !== thisParts[0]) {
      anyChange = true;
      majorChange = true;
    } else if (prevParts[1] !== thisParts[1]) {
      anyChange = true;
      minorChange = true;
    } else if (prevParts[2] !== thisParts[2]) {
      anyChange = true;
      patchChange = true;
    }

    if (majorChange || minorChange) {
      // We have a major or minor web extension update, show the about page.
      browser.tabs.create({url: './pages/about.html'});
    }

    if (anyChange) {
      // The web extension was updated, store our new version into our local storage.
      var settingObject = {};
      settingObject['setting-version'] = thisVersion;
      browser.storage.local.set(settingObject);
    }

    if (anyChange && parseInt(prevParts[0]) <= 2) {
      // Make sure we don't have black- or whitelisted any empty url's.
      browser.storage.local.remove('');
    }
  }
})(browser);
