function throttle(callback, wait = 0, {
  start = true,
  middle = true,
  once = false
} = {}) {
  var last = 0;
  var timer;
  var cancelled = false;

  var fn = function fn(...args) {
    if (cancelled) return;
    var delta = Date.now() - last;
    last = Date.now();

    if (start) {
      //eslint-disable-next-line flowtype/no-flow-fix-me-comments
      // $FlowFixMe this isn't a const
      start = false;
      callback(...args);
      if (once) fn.cancel();
    } else if (middle && delta < wait || !middle) {
      clearTimeout(timer);
      timer = setTimeout(function () {
        last = Date.now();
        callback(...args);
        if (once) fn.cancel();
      }, !middle ? wait : wait - delta);
    }
  };

  fn.cancel = function () {
    clearTimeout(timer);
    cancelled = true;
  };

  return fn;
}
function debounce(callback, wait = 0, {
  start = false,
  middle = false,
  once = false
} = {}) {
  return throttle(callback, wait, {
    start: start,
    middle: middle,
    once: once
  });
}

const states = new WeakMap();
class AutoCheckElement extends HTMLElement {
    connectedCallback() {
        const input = this.input;
        if (!input)
            return;
        const checker = debounce(check.bind(null, this), 300);
        const state = { check: checker, controller: null };
        states.set(this, state);
        input.addEventListener('input', setLoadingState);
        input.addEventListener('input', checker);
        input.autocomplete = 'off';
        input.spellcheck = false;
    }
    disconnectedCallback() {
        const input = this.input;
        if (!input)
            return;
        const state = states.get(this);
        if (!state)
            return;
        states.delete(this);
        input.removeEventListener('input', setLoadingState);
        input.removeEventListener('input', state.check);
        input.setCustomValidity('');
    }
    attributeChangedCallback(name) {
        if (name === 'required') {
            const input = this.input;
            if (!input)
                return;
            input.required = this.required;
        }
    }
    static get observedAttributes() {
        return ['required'];
    }
    get input() {
        return this.querySelector('input');
    }
    get src() {
        const src = this.getAttribute('src');
        if (!src)
            return '';
        const link = this.ownerDocument.createElement('a');
        link.href = src;
        return link.href;
    }
    set src(value) {
        this.setAttribute('src', value);
    }
    get csrf() {
        const csrfElement = this.querySelector('[data-csrf]');
        return this.getAttribute('csrf') || (csrfElement instanceof HTMLInputElement && csrfElement.value) || '';
    }
    set csrf(value) {
        this.setAttribute('csrf', value);
    }
    get required() {
        return this.hasAttribute('required');
    }
    set required(required) {
        if (required) {
            this.setAttribute('required', '');
        }
        else {
            this.removeAttribute('required');
        }
    }
}
function setLoadingState(event) {
    const input = event.currentTarget;
    if (!(input instanceof HTMLInputElement))
        return;
    const autoCheckElement = input.closest('auto-check');
    if (!(autoCheckElement instanceof AutoCheckElement))
        return;
    const src = autoCheckElement.src;
    const csrf = autoCheckElement.csrf;
    const state = states.get(autoCheckElement);
    if (!src || !csrf || !state) {
        return;
    }
    let message = 'Verifying…';
    const setValidity = (text) => (message = text);
    input.dispatchEvent(new CustomEvent('auto-check-start', {
        bubbles: true,
        detail: { setValidity }
    }));
    if (autoCheckElement.required) {
        input.setCustomValidity(message);
    }
}
function makeAbortController() {
    if ('AbortController' in window) {
        return new AbortController();
    }
    return {
        signal: null,
        abort() {
        }
    };
}
async function fetchWithNetworkEvents(el, url, options) {
    try {
        const response = await fetch(url, options);
        el.dispatchEvent(new CustomEvent('load'));
        el.dispatchEvent(new CustomEvent('loadend'));
        return response;
    }
    catch (error) {
        if (error.name !== 'AbortError') {
            el.dispatchEvent(new CustomEvent('error'));
            el.dispatchEvent(new CustomEvent('loadend'));
        }
        throw error;
    }
}
async function check(autoCheckElement) {
    const input = autoCheckElement.input;
    if (!input) {
        return;
    }
    const src = autoCheckElement.src;
    const csrf = autoCheckElement.csrf;
    const state = states.get(autoCheckElement);
    if (!src || !csrf || !state) {
        if (autoCheckElement.required) {
            input.setCustomValidity('');
        }
        return;
    }
    if (!input.value.trim()) {
        if (autoCheckElement.required) {
            input.setCustomValidity('');
        }
        return;
    }
    const body = new FormData();
    body.append('authenticity_token', csrf);
    body.append('value', input.value);
    input.dispatchEvent(new CustomEvent('auto-check-send', {
        bubbles: true,
        detail: { body }
    }));
    if (state.controller) {
        state.controller.abort();
    }
    else {
        autoCheckElement.dispatchEvent(new CustomEvent('loadstart'));
    }
    state.controller = makeAbortController();
    try {
        const response = await fetchWithNetworkEvents(autoCheckElement, src, {
            credentials: 'same-origin',
            signal: state.controller.signal,
            method: 'POST',
            body
        });
        if (response.ok) {
            processSuccess(response, input, autoCheckElement.required);
        }
        else {
            processFailure(response, input, autoCheckElement.required);
        }
        state.controller = null;
        input.dispatchEvent(new CustomEvent('auto-check-complete', { bubbles: true }));
    }
    catch (error) {
        if (error.name !== 'AbortError') {
            state.controller = null;
            input.dispatchEvent(new CustomEvent('auto-check-complete', { bubbles: true }));
        }
    }
}
function processSuccess(response, input, required) {
    if (required) {
        input.setCustomValidity('');
    }
    input.dispatchEvent(new CustomEvent('auto-check-success', {
        bubbles: true,
        detail: {
            response: response.clone()
        }
    }));
}
function processFailure(response, input, required) {
    let message = 'Validation failed';
    const setValidity = (text) => (message = text);
    input.dispatchEvent(new CustomEvent('auto-check-error', {
        bubbles: true,
        detail: {
            response: response.clone(),
            setValidity
        }
    }));
    if (required) {
        input.setCustomValidity(message);
    }
}
if (!window.customElements.get('auto-check')) {
    window.AutoCheckElement = AutoCheckElement;
    window.customElements.define('auto-check', AutoCheckElement);
}

export default AutoCheckElement;
