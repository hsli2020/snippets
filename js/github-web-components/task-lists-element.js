const sortHandlers = new WeakMap();
let state = null;
function isDragging() {
    return !!state;
}
function sortable(el, sortStarted, sortFinished) {
    sortHandlers.set(el, { sortStarted, sortFinished });
    el.addEventListener('dragstart', onDragStart);
    el.addEventListener('dragenter', onDragEnter);
    el.addEventListener('dragend', onDragEnd);
    el.addEventListener('drop', onDrop);
    el.addEventListener('dragover', onDragOver);
}
function isBefore(item1, item2) {
    if (item1.parentNode === item2.parentNode) {
        let node = item1;
        while (node) {
            if (node === item2)
                return true;
            node = node.previousElementSibling;
        }
    }
    return false;
}
function isSameContainer(item1, item2) {
    return item1.closest('task-lists') === item2.closest('task-lists');
}
function onDragStart(event) {
    if (event.currentTarget !== event.target)
        return;
    const target = event.currentTarget;
    if (!(target instanceof Element))
        return;
    const sourceList = target.closest('.contains-task-list');
    if (!sourceList)
        return;
    target.classList.add('is-ghost');
    if (event.dataTransfer) {
        event.dataTransfer.setData('text/plain', (target.textContent || '').trim());
    }
    if (!target.parentElement)
        return;
    const siblings = Array.from(target.parentElement.children);
    const sourceIndex = siblings.indexOf(target);
    const handlers = sortHandlers.get(target);
    if (handlers) {
        handlers.sortStarted(sourceList);
    }
    state = {
        didDrop: false,
        dragging: target,
        dropzone: target,
        sourceList,
        sourceSibling: siblings[sourceIndex + 1] || null,
        sourceIndex
    };
}
function onDragEnter(event) {
    if (!state)
        return;
    const dropzone = event.currentTarget;
    if (!(dropzone instanceof Element))
        return;
    if (!isSameContainer(state.dragging, dropzone)) {
        event.stopPropagation();
        return;
    }
    event.preventDefault();
    if (event.dataTransfer) {
        event.dataTransfer.dropEffect = 'move';
    }
    if (state.dropzone === dropzone)
        return;
    state.dragging.classList.add('is-dragging');
    state.dropzone = dropzone;
    if (isBefore(state.dragging, dropzone)) {
        dropzone.before(state.dragging);
    }
    else {
        dropzone.after(state.dragging);
    }
}
function onDrop(event) {
    if (!state)
        return;
    event.preventDefault();
    event.stopPropagation();
    const dropzone = event.currentTarget;
    if (!(dropzone instanceof Element))
        return;
    state.didDrop = true;
    if (!state.dragging.parentElement)
        return;
    let newIndex = Array.from(state.dragging.parentElement.children).indexOf(state.dragging);
    const currentList = dropzone.closest('.contains-task-list');
    if (!currentList)
        return;
    if (state.sourceIndex === newIndex && state.sourceList === currentList)
        return;
    if (state.sourceList === currentList && state.sourceIndex < newIndex) {
        newIndex++;
    }
    const src = { list: state.sourceList, index: state.sourceIndex };
    const dst = { list: currentList, index: newIndex };
    const handlers = sortHandlers.get(state.dragging);
    if (handlers) {
        handlers.sortFinished({ src, dst });
    }
}
function onDragEnd() {
    if (!state)
        return;
    state.dragging.classList.remove('is-dragging');
    state.dragging.classList.remove('is-ghost');
    if (!state.didDrop) {
        state.sourceList.insertBefore(state.dragging, state.sourceSibling);
    }
    state = null;
}
function onDragOver(event) {
    if (!state)
        return;
    const dropzone = event.currentTarget;
    if (!(dropzone instanceof Element))
        return;
    if (!isSameContainer(state.dragging, dropzone)) {
        event.stopPropagation();
        return;
    }
    event.preventDefault();
    if (event.dataTransfer) {
        event.dataTransfer.dropEffect = 'move';
    }
}

const observers = new WeakMap();
class TaskListsElement extends HTMLElement {
    constructor() {
        super();
        this.addEventListener('change', (event) => {
            const checkbox = event.target;
            if (!(checkbox instanceof HTMLInputElement))
                return;
            if (!checkbox.classList.contains('task-list-item-checkbox'))
                return;
            this.dispatchEvent(new CustomEvent('task-lists-check', {
                bubbles: true,
                detail: {
                    position: position(checkbox),
                    checked: checkbox.checked
                }
            }));
        });
        observers.set(this, new MutationObserver(syncState.bind(null, this)));
    }
    connectedCallback() {
        const observer = observers.get(this);
        if (observer) {
            observer.observe(this, { childList: true, subtree: true });
        }
        syncState(this);
    }
    disconnectedCallback() {
        const observer = observers.get(this);
        if (observer) {
            observer.disconnect();
        }
    }
    get disabled() {
        return this.hasAttribute('disabled');
    }
    set disabled(value) {
        if (value) {
            this.setAttribute('disabled', '');
        }
        else {
            this.removeAttribute('disabled');
        }
    }
    get sortable() {
        return this.hasAttribute('sortable');
    }
    set sortable(value) {
        if (value) {
            this.setAttribute('sortable', '');
        }
        else {
            this.removeAttribute('sortable');
        }
    }
    static get observedAttributes() {
        return ['disabled'];
    }
    attributeChangedCallback(name, oldValue, newValue) {
        if (oldValue === newValue)
            return;
        switch (name) {
            case 'disabled':
                syncDisabled(this);
                break;
        }
    }
}
const handleTemplate = document.createElement('template');
handleTemplate.innerHTML = `
  <span class="handle">
    <svg class="drag-handle" aria-hidden="true" width="16" height="16">
      <path d="M10 13a1 1 0 100-2 1 1 0 000 2zm-4 0a1 1 0 100-2 1 1 0 000 2zm1-5a1 1 0 11-2 0 1 1 0 012 0zm3 1a1 1 0 100-2 1 1 0 000 2zm1-5a1 1 0 11-2 0 1 1 0 012 0zM6 5a1 1 0 100-2 1 1 0 000 2z"/>
    </svg>
  </span>`;
const initialized = new WeakMap();
function initItem(el) {
    if (initialized.get(el))
        return;
    initialized.set(el, true);
    const currentTaskList = el.closest('task-lists');
    if (!(currentTaskList instanceof TaskListsElement))
        return;
    if (currentTaskList.querySelectorAll('.task-list-item').length <= 1)
        return;
    const fragment = handleTemplate.content.cloneNode(true);
    const handle = fragment.querySelector('.handle');
    el.prepend(fragment);
    if (!handle)
        throw new Error('handle not found');
    handle.addEventListener('mouseenter', onHandleMouseOver);
    handle.addEventListener('mouseleave', onHandleMouseOut);
    sortable(el, onSortStart, onSorted);
    el.addEventListener('mouseenter', onListItemMouseOver);
    el.addEventListener('mouseleave', onListItemMouseOut);
}
function onListItemMouseOver(event) {
    const item = event.currentTarget;
    if (!(item instanceof Element))
        return;
    const list = item.closest('task-lists');
    if (!(list instanceof TaskListsElement))
        return;
    if (list.sortable && !list.disabled) {
        item.classList.add('hovered');
    }
}
function onListItemMouseOut(event) {
    const item = event.currentTarget;
    if (!(item instanceof Element))
        return;
    item.classList.remove('hovered');
}
function position(checkbox) {
    const list = taskList(checkbox);
    if (!list)
        throw new Error('.contains-task-list not found');
    const item = checkbox.closest('.task-list-item');
    const index = item ? Array.from(list.children).indexOf(item) : -1;
    return [listIndex(list), index];
}
function taskList(el) {
    const parent = el.parentElement;
    return parent ? parent.closest('.contains-task-list') : null;
}
function isRootTaskList(el) {
    return taskList(el) === rootTaskList(el);
}
function rootTaskList(node) {
    const list = taskList(node);
    return list ? rootTaskList(list) || list : null;
}
function syncState(list) {
    const items = list.querySelectorAll('.contains-task-list > .task-list-item');
    for (const el of items) {
        if (isRootTaskList(el)) {
            initItem(el);
        }
    }
    syncDisabled(list);
}
function syncDisabled(list) {
    for (const el of list.querySelectorAll('.task-list-item')) {
        el.classList.toggle('enabled', !list.disabled);
    }
    for (const el of list.querySelectorAll('.task-list-item-checkbox')) {
        if (el instanceof HTMLInputElement) {
            el.disabled = list.disabled;
        }
    }
}
function listIndex(list) {
    const container = list.closest('task-lists');
    if (!container)
        throw new Error('parent not found');
    return Array.from(container.querySelectorAll('ol, ul')).indexOf(list);
}
const originalLists = new WeakMap();
function onSortStart(srcList) {
    const container = srcList.closest('task-lists');
    if (!container)
        throw new Error('parent not found');
    originalLists.set(container, Array.from(container.querySelectorAll('ol, ul')));
}
function onSorted({ src, dst }) {
    const container = src.list.closest('task-lists');
    if (!container)
        return;
    const lists = originalLists.get(container);
    if (!lists)
        return;
    originalLists.delete(container);
    container.dispatchEvent(new CustomEvent('task-lists-move', {
        bubbles: true,
        detail: {
            src: [lists.indexOf(src.list), src.index],
            dst: [lists.indexOf(dst.list), dst.index]
        }
    }));
}
function onHandleMouseOver(event) {
    const target = event.currentTarget;
    if (!(target instanceof Element))
        return;
    const item = target.closest('.task-list-item');
    if (!item)
        return;
    const list = item.closest('task-lists');
    if (!(list instanceof TaskListsElement))
        return;
    if (list.sortable && !list.disabled) {
        item.setAttribute('draggable', 'true');
    }
}
function onHandleMouseOut(event) {
    if (isDragging())
        return;
    const target = event.currentTarget;
    if (!(target instanceof Element))
        return;
    const item = target.closest('.task-list-item');
    if (!item)
        return;
    item.setAttribute('draggable', 'false');
}

if (!window.customElements.get('task-lists')) {
    window.TaskListsElement = TaskListsElement;
    window.customElements.define('task-lists', TaskListsElement);
}

export default TaskListsElement;
