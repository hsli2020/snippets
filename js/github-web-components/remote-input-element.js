(function (factory) {
    if (typeof module === "object" && typeof module.exports === "object") {
        var v = factory(require, exports);
        if (v !== undefined) module.exports = v;
    }
    else if (typeof define === "function" && define.amd) {
        define(["require", "exports"], factory);
    }
})(function (require, exports) {
    "use strict";
    Object.defineProperty(exports, "__esModule", { value: true });
    const states = new WeakMap();
    class RemoteInputElement extends HTMLElement {
        constructor() {
            super();
            const fetch = fetchResults.bind(null, this, true);
            const state = { currentQuery: null, oninput: debounce(fetch), fetch, controller: null };
            states.set(this, state);
        }
        static get observedAttributes() {
            return ['src'];
        }
        attributeChangedCallback(name, oldValue) {
            if (oldValue && name === 'src') {
                fetchResults(this, false);
            }
        }
        connectedCallback() {
            const input = this.input;
            if (!input)
                return;
            input.setAttribute('autocomplete', 'off');
            input.setAttribute('spellcheck', 'false');
            const state = states.get(this);
            if (!state)
                return;
            input.addEventListener('focus', state.fetch);
            input.addEventListener('change', state.fetch);
            input.addEventListener('input', state.oninput);
        }
        disconnectedCallback() {
            const input = this.input;
            if (!input)
                return;
            const state = states.get(this);
            if (!state)
                return;
            input.removeEventListener('focus', state.fetch);
            input.removeEventListener('change', state.fetch);
            input.removeEventListener('input', state.oninput);
        }
        get input() {
            const input = this.querySelector('input, textarea');
            return input instanceof HTMLInputElement || input instanceof HTMLTextAreaElement ? input : null;
        }
        get src() {
            return this.getAttribute('src') || '';
        }
        set src(url) {
            this.setAttribute('src', url);
        }
    }
    function makeAbortController() {
        if ('AbortController' in window) {
            return new AbortController();
        }
        return { signal: null, abort() { } };
    }
    async function fetchResults(remoteInput, checkCurrentQuery) {
        const input = remoteInput.input;
        if (!input)
            return;
        const state = states.get(remoteInput);
        if (!state)
            return;
        const query = input.value;
        if (checkCurrentQuery && state.currentQuery === query)
            return;
        state.currentQuery = query;
        const src = remoteInput.src;
        if (!src)
            return;
        const resultsContainer = document.getElementById(remoteInput.getAttribute('aria-owns') || '');
        if (!resultsContainer)
            return;
        const url = new URL(src, window.location.href);
        const params = new URLSearchParams(url.search);
        params.append(remoteInput.getAttribute('param') || 'q', query);
        url.search = params.toString();
        if (state.controller) {
            state.controller.abort();
        }
        else {
            remoteInput.dispatchEvent(new CustomEvent('loadstart'));
            remoteInput.setAttribute('loading', '');
        }
        state.controller = makeAbortController();
        let response;
        let html = '';
        try {
            response = await fetchWithNetworkEvents(remoteInput, url.toString(), {
                signal: state.controller.signal,
                credentials: 'same-origin',
                headers: { accept: 'text/fragment+html' }
            });
            html = await response.text();
            remoteInput.removeAttribute('loading');
            state.controller = null;
        }
        catch (error) {
            if (error.name !== 'AbortError') {
                remoteInput.removeAttribute('loading');
                state.controller = null;
            }
            return;
        }
        if (response && response.ok) {
            resultsContainer.innerHTML = html;
            remoteInput.dispatchEvent(new CustomEvent('remote-input-success', { bubbles: true }));
        }
        else {
            remoteInput.dispatchEvent(new CustomEvent('remote-input-error', { bubbles: true }));
        }
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
    function debounce(callback) {
        let timeout;
        return function () {
            clearTimeout(timeout);
            timeout = setTimeout(() => {
                clearTimeout(timeout);
                callback();
            }, 300);
        };
    }
    exports.default = RemoteInputElement;
    if (!window.customElements.get('remote-input')) {
        window.RemoteInputElement = RemoteInputElement;
        window.customElements.define('remote-input', RemoteInputElement);
    }
});
