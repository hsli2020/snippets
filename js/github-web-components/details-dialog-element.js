const CLOSE_ATTR = 'data-close-dialog';
const CLOSE_SELECTOR = `[${CLOSE_ATTR}]`;
function autofocus(el) {
    let autofocusElement = Array.from(el.querySelectorAll('[autofocus]')).filter(focusable)[0];
    if (!autofocusElement) {
        autofocusElement = el;
        el.setAttribute('tabindex', '-1');
    }
    autofocusElement.focus();
}
function keydown(event) {
    const details = event.currentTarget;
    if (!(details instanceof Element))
        return;
    if (event.key === 'Escape' || event.key === 'Esc') {
        toggleDetails(details, false);
        event.stopPropagation();
    }
    else if (event.key === 'Tab') {
        restrictTabBehavior(event);
    }
}
function focusable(el) {
    return el.tabIndex >= 0 && !el.disabled && visible(el);
}
function visible(el) {
    return (!el.hidden &&
        (!el.type || el.type !== 'hidden') &&
        (el.offsetWidth > 0 || el.offsetHeight > 0));
}
function restrictTabBehavior(event) {
    if (!(event.currentTarget instanceof Element))
        return;
    const dialog = event.currentTarget.querySelector('details-dialog');
    if (!dialog)
        return;
    event.preventDefault();
    const elements = Array.from(dialog.querySelectorAll('*')).filter(focusable);
    if (elements.length === 0)
        return;
    const movement = event.shiftKey ? -1 : 1;
    const root = dialog.getRootNode();
    const currentFocus = dialog.contains(root.activeElement) ? root.activeElement : null;
    let targetIndex = movement === -1 ? -1 : 0;
    if (currentFocus instanceof HTMLElement) {
        const currentIndex = elements.indexOf(currentFocus);
        if (currentIndex !== -1) {
            targetIndex = currentIndex + movement;
        }
    }
    if (targetIndex < 0) {
        targetIndex = elements.length - 1;
    }
    else {
        targetIndex = targetIndex % elements.length;
    }
    elements[targetIndex].focus();
}
function allowClosingDialog(details) {
    const dialog = details.querySelector('details-dialog');
    if (!(dialog instanceof DetailsDialogElement))
        return true;
    return dialog.dispatchEvent(new CustomEvent('details-dialog-close', {
        bubbles: true,
        cancelable: true
    }));
}
function onSummaryClick(event) {
    if (!(event.currentTarget instanceof Element))
        return;
    const details = event.currentTarget.closest('details');
    if (!details || !details.hasAttribute('open'))
        return;
    if (!allowClosingDialog(details)) {
        event.preventDefault();
        event.stopPropagation();
    }
}
function toggle(event) {
    const details = event.currentTarget;
    if (!(details instanceof Element))
        return;
    const dialog = details.querySelector('details-dialog');
    if (!(dialog instanceof DetailsDialogElement))
        return;
    if (details.hasAttribute('open')) {
        const root = 'getRootNode' in dialog ? dialog.getRootNode() : document;
        if (root.activeElement instanceof HTMLElement) {
            initialized.set(dialog, { details, activeElement: root.activeElement });
        }
        autofocus(dialog);
        details.addEventListener('keydown', keydown);
    }
    else {
        for (const form of dialog.querySelectorAll('form')) {
            form.reset();
        }
        const focusElement = findFocusElement(details, dialog);
        if (focusElement)
            focusElement.focus();
        details.removeEventListener('keydown', keydown);
    }
}
function findFocusElement(details, dialog) {
    const state = initialized.get(dialog);
    if (state && state.activeElement instanceof HTMLElement) {
        return state.activeElement;
    }
    else {
        return details.querySelector('summary');
    }
}
function toggleDetails(details, open) {
    if (open === details.hasAttribute('open'))
        return;
    if (open) {
        details.setAttribute('open', '');
    }
    else if (allowClosingDialog(details)) {
        details.removeAttribute('open');
    }
}
function loadIncludeFragment(event) {
    const details = event.currentTarget;
    if (!(details instanceof Element))
        return;
    const dialog = details.querySelector('details-dialog');
    if (!(dialog instanceof DetailsDialogElement))
        return;
    const loader = dialog.querySelector('include-fragment:not([src])');
    if (!loader)
        return;
    const src = dialog.src;
    if (src === null)
        return;
    loader.addEventListener('loadend', () => {
        if (details.hasAttribute('open'))
            autofocus(dialog);
    });
    loader.setAttribute('src', src);
    removeIncludeFragmentEventListeners(details);
}
function updateIncludeFragmentEventListeners(details, src, preload) {
    removeIncludeFragmentEventListeners(details);
    if (src) {
        details.addEventListener('toggle', loadIncludeFragment, { once: true });
    }
    if (src && preload) {
        details.addEventListener('mouseover', loadIncludeFragment, { once: true });
    }
}
function removeIncludeFragmentEventListeners(details) {
    details.removeEventListener('toggle', loadIncludeFragment);
    details.removeEventListener('mouseover', loadIncludeFragment);
}
const initialized = new WeakMap();
class DetailsDialogElement extends HTMLElement {
    static get CLOSE_ATTR() {
        return CLOSE_ATTR;
    }
    static get CLOSE_SELECTOR() {
        return CLOSE_SELECTOR;
    }
    constructor() {
        super();
        initialized.set(this, { details: null, activeElement: null });
        this.addEventListener('click', function ({ target }) {
            if (!(target instanceof Element))
                return;
            const details = target.closest('details');
            if (details && target.closest(CLOSE_SELECTOR)) {
                toggleDetails(details, false);
            }
        });
    }
    get src() {
        return this.getAttribute('src');
    }
    set src(value) {
        this.setAttribute('src', value || '');
    }
    get preload() {
        return this.hasAttribute('preload');
    }
    set preload(value) {
        value ? this.setAttribute('preload', '') : this.removeAttribute('preload');
    }
    connectedCallback() {
        this.setAttribute('role', 'dialog');
        this.setAttribute('aria-modal', 'true');
        const state = initialized.get(this);
        if (!state)
            return;
        const details = this.parentElement;
        if (!details)
            return;
        const summary = details.querySelector('summary');
        if (summary) {
            if (!summary.hasAttribute('role'))
                summary.setAttribute('role', 'button');
            summary.addEventListener('click', onSummaryClick, { capture: true });
        }
        details.addEventListener('toggle', toggle);
        state.details = details;
        updateIncludeFragmentEventListeners(details, this.src, this.preload);
    }
    disconnectedCallback() {
        const state = initialized.get(this);
        if (!state)
            return;
        const { details } = state;
        if (!details)
            return;
        details.removeEventListener('toggle', toggle);
        removeIncludeFragmentEventListeners(details);
        const summary = details.querySelector('summary');
        if (summary) {
            summary.removeEventListener('click', onSummaryClick, { capture: true });
        }
        state.details = null;
    }
    toggle(open) {
        const state = initialized.get(this);
        if (!state)
            return;
        const { details } = state;
        if (!details)
            return;
        toggleDetails(details, open);
    }
    static get observedAttributes() {
        return ['src', 'preload'];
    }
    attributeChangedCallback() {
        const state = initialized.get(this);
        if (!state)
            return;
        const { details } = state;
        if (!details)
            return;
        updateIncludeFragmentEventListeners(details, this.src, this.preload);
    }
}
export default DetailsDialogElement;
if (!window.customElements.get('details-dialog')) {
    window.DetailsDialogElement = DetailsDialogElement;
    window.customElements.define('details-dialog', DetailsDialogElement);
}
