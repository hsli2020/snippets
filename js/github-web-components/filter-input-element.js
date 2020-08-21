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
    class FilterInputElement extends HTMLElement {
        constructor() {
            super();
            this.currentQuery = null;
            this.filter = null;
            this.debounceInputChange = debounce(() => filterResults(this, true));
            this.boundFilterResults = () => {
                filterResults(this, false);
            };
        }
        static get observedAttributes() {
            return ['aria-owns'];
        }
        attributeChangedCallback(name, oldValue) {
            if (oldValue && name === 'aria-owns') {
                filterResults(this, false);
            }
        }
        connectedCallback() {
            const input = this.input;
            if (!input)
                return;
            input.setAttribute('autocomplete', 'off');
            input.setAttribute('spellcheck', 'false');
            input.addEventListener('focus', this.boundFilterResults);
            input.addEventListener('change', this.boundFilterResults);
            input.addEventListener('input', this.debounceInputChange);
        }
        disconnectedCallback() {
            const input = this.input;
            if (!input)
                return;
            input.removeEventListener('focus', this.boundFilterResults);
            input.removeEventListener('change', this.boundFilterResults);
            input.removeEventListener('input', this.debounceInputChange);
        }
        get input() {
            const input = this.querySelector('input');
            return input instanceof HTMLInputElement ? input : null;
        }
        reset() {
            const input = this.input;
            if (input) {
                input.value = '';
                input.dispatchEvent(new Event('change', { bubbles: true }));
            }
        }
    }
    async function filterResults(filterInput, checkCurrentQuery = false) {
        const input = filterInput.input;
        if (!input)
            return;
        const query = input.value.trim();
        const id = filterInput.getAttribute('aria-owns');
        if (!id)
            return;
        const container = document.getElementById(id);
        if (!container)
            return;
        const list = container.hasAttribute('data-filter-list') ? container : container.querySelector('[data-filter-list]');
        if (!list)
            return;
        filterInput.dispatchEvent(new CustomEvent('filter-input-start', {
            bubbles: true
        }));
        if (checkCurrentQuery && filterInput.currentQuery === query)
            return;
        filterInput.currentQuery = query;
        const filter = filterInput.filter || matchSubstring;
        const total = list.childElementCount;
        let count = 0;
        let hideNew = false;
        for (const item of Array.from(list.children)) {
            if (!(item instanceof HTMLElement))
                continue;
            const itemText = getText(item);
            const result = filter(item, itemText, query);
            if (result.hideNew === true)
                hideNew = result.hideNew;
            item.hidden = !result.match;
            if (result.match)
                count++;
        }
        const newItem = container.querySelector('[data-filter-new-item]');
        const showCreateOption = !!newItem && query.length > 0 && !hideNew;
        if (newItem instanceof HTMLElement) {
            newItem.hidden = !showCreateOption;
            if (showCreateOption)
                updateNewItem(newItem, query);
        }
        toggleBlankslate(container, count > 0 || showCreateOption);
        filterInput.dispatchEvent(new CustomEvent('filter-input-updated', {
            bubbles: true,
            detail: {
                count,
                total
            }
        }));
    }
    function matchSubstring(_item, itemText, query) {
        const match = itemText.toLowerCase().indexOf(query.toLowerCase()) !== -1;
        return {
            match,
            hideNew: itemText === query
        };
    }
    function getText(filterableItem) {
        const target = filterableItem.querySelector('[data-filter-item-text]') || filterableItem;
        return (target.textContent || '').trim();
    }
    function updateNewItem(newItem, query) {
        const newItemText = newItem.querySelector('[data-filter-new-item-text]');
        if (newItemText)
            newItemText.textContent = query;
        const newItemValue = newItem.querySelector('[data-filter-new-item-value]');
        if (newItemValue instanceof HTMLInputElement || newItemValue instanceof HTMLButtonElement) {
            newItemValue.value = query;
        }
    }
    function toggleBlankslate(container, force) {
        const emptyState = container.querySelector('[data-filter-empty-state]');
        if (emptyState instanceof HTMLElement)
            emptyState.hidden = force;
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
    exports.default = FilterInputElement;
    if (!window.customElements.get('filter-input')) {
        window.FilterInputElement = FilterInputElement;
        window.customElements.define('filter-input', FilterInputElement);
    }
});
