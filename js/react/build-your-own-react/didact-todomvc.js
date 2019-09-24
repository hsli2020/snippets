(function () {
'use strict';

const TEXT_ELEMENT = "TEXT ELEMENT";

function createElement(type, config, ...args) {
  const props = Object.assign({}, config);
  const hasChildren = args.length > 0;
  const rawChildren = hasChildren ? [].concat(...args) : [];
  props.children = rawChildren.filter(c => c != null && c !== false).map(c => c instanceof Object ? c : createTextElement(c));
  return { type, props };
}

function createTextElement(value) {
  return createElement(TEXT_ELEMENT, { nodeValue: value });
}

function createDom(element) {
  return element.type === "TEXT ELEMENT" ? document.createTextNode("") : document.createElement(element.type);
}

function appendChildren(dom, childDoms) {
  childDoms.forEach(childDom => dom.appendChild(childDom));
}

function appendChild(dom, childDom) {
  dom.appendChild(childDom);
}

function removeChild(dom, childDom) {
  dom.removeChild(childDom);
}

function replaceChild(newChildDom, oldChildDom) {
  const parentDom = oldChildDom.parentNode;
  parentDom.replaceChild(newChildDom, oldChildDom);
}

function updateAttributes(dom, prevProps, nextProps) {
  // Remove events
  Object.keys(prevProps).filter(isEvent).forEach(name => {
    const eventType = name.toLowerCase().substring(2);
    dom.removeEventListener(eventType, prevProps[name]);
  });

  // Remove attributes
  Object.keys(prevProps).filter(isAttribute).forEach(name => {
    dom[name] = null;
  });

  // Set attributes
  Object.keys(nextProps).filter(isAttribute).forEach(name => {
    const value = nextProps[name];
    if (name === "autoFocus") {
      if (value === false) {
        dom.blur();
      } else {
        dom.focus();
      }
    } else if (value != null && value !== false) {
      dom[name] = value;
    }
  });

  // Set events
  Object.keys(nextProps).filter(isEvent).forEach(name => {
    const eventType = name.toLowerCase().substring(2);
    dom.addEventListener(eventType, nextProps[name]);
  });
}

const isEvent = name => name.startsWith("on");
const isAttribute = name => !isEvent(name) && name != "children";

function reconcile(parentDom, instance, element) {
  if (instance == null) {
    //Create element
    const newInstance = instantiate(element);
    appendChild(parentDom, newInstance.dom);
    return newInstance;
  } else if (element == null) {
    //Remove element
    removeChild(parentDom, instance.dom);
    return null;
  } else if (instance.element.type !== element.type) {
    //Replace instance
    const newInstance = instantiate(element);
    replaceChild(newInstance.dom, instance.dom);
    return newInstance;
  } else if (typeof element.type === "string") {
    //Update dom instance
    const nextChildElements = element.props.children || [];
    updateAttributes(instance.dom, instance.element.props, element.props);
    instance.childInstances = reconcileChildren(instance, nextChildElements);
    instance.element = element;
    return instance;
  } else {
    //Update composite instance
    instance.publicInstance.props = element.props;
    const childElement = instance.publicInstance.render();
    const oldChildInstance = instance.childInstance;
    const childInstance = reconcile(parentDom, oldChildInstance, childElement);
    instance.dom = childInstance.dom;
    instance.childInstance = childInstance;
    instance.element = element;
    return instance;
  }
}

function reconcileChildren(instance, nextChildElements) {
  const dom = instance.dom;
  const childInstances = instance.childInstances;
  const newChildInstances = [];
  const count = Math.max(childInstances.length, nextChildElements.length);
  for (let i = 0; i < count; i++) {
    const childInstance = childInstances[i];
    const element = nextChildElements[i];
    const newChildInstance = reconcile(dom, childInstance, element);
    newChildInstances.push(newChildInstance);
  }
  return newChildInstances.filter(instance => instance != null);
}

function instantiate(element) {
  const isDomElement = typeof element.type === "string";
  const instance = {};

  if (isDomElement) {
    const childElements = element.props.children || [];
    const childInstances = childElements.map(instantiate);
    const childDoms = childInstances.map(childInstance => childInstance.dom);

    const dom = createDom(element);
    updateAttributes(dom, [], element.props);
    appendChildren(dom, childDoms);

    instance.dom = dom;
    instance.element = element;
    instance.childInstances = childInstances;
  } else {
    const publicInstance = Component.__create(element, instance);
    const childElement = publicInstance.render();
    const childInstance = instantiate(childElement);

    instance.dom = childInstance.dom;
    instance.element = element;
    instance.childInstance = childInstance;
    instance.publicInstance = publicInstance;
  }

  return instance;
}

class Component {
  static __create(element, internalInstance) {
    const { type, props } = element;
    const publicInstance = new type(props);
    publicInstance.__internalInstance = internalInstance;
    return publicInstance;
  }

  constructor(props) {
    this.props = props;
    this.state = this.state || {};
    this.__internalInstance = null; // Set by __create
  }

  setState(partialState) {
    this.state = Object.assign({}, this.state, partialState);
    updateInstance(this.__internalInstance);
  }
}

function updateInstance(internalInstance) {
  const parentDom = internalInstance.dom.parentNode;
  const element = internalInstance.element;
  reconcile(parentDom, internalInstance, element);
}

function render(element, container) {
  const instance = reconcile(container, null, element);
}

function uuid() {
	let uuid = '';
	for (let i = 0; i < 32; i++) {
		let random = Math.random() * 16 | 0;
		if (i === 8 || i === 12 || i === 16 || i === 20) {
			uuid += '-';
		}
		uuid += (i === 12 ? 4 : i === 16 ? random & 3 | 8 : random).toString(16);
	}
	return uuid;
}

function pluralize(count, word) {
	return count === 1 ? word : word + 's';
}

function store(namespace, data) {
	if (data) return localStorage[namespace] = JSON.stringify(data);

	let store = localStorage[namespace];
	return store && JSON.parse(store) || [];
}

var _extends = Object.assign || function (target) {
  for (var i = 1; i < arguments.length; i++) {
    var source = arguments[i];

    for (var key in source) {
      if (Object.prototype.hasOwnProperty.call(source, key)) {
        target[key] = source[key];
      }
    }
  }

  return target;
};

var get = function get(object, property, receiver) {
  if (object === null) object = Function.prototype;
  var desc = Object.getOwnPropertyDescriptor(object, property);

  if (desc === undefined) {
    var parent = Object.getPrototypeOf(object);

    if (parent === null) {
      return undefined;
    } else {
      return get(parent, property, receiver);
    }
  } else if ("value" in desc) {
    return desc.value;
  } else {
    var getter = desc.get;

    if (getter === undefined) {
      return undefined;
    }

    return getter.call(receiver);
  }
};

















var set = function set(object, property, value, receiver) {
  var desc = Object.getOwnPropertyDescriptor(object, property);

  if (desc === undefined) {
    var parent = Object.getPrototypeOf(object);

    if (parent !== null) {
      set(parent, property, value, receiver);
    }
  } else if ("value" in desc && desc.writable) {
    desc.value = value;
  } else {
    var setter = desc.set;

    if (setter !== undefined) {
      setter.call(receiver, value);
    }
  }

  return value;
};

class TodoModel {
	constructor(key, sub) {
		this.key = key;
		this.todos = store(key) || [];
		this.onChanges = [sub];
	}

	inform() {
		store(this.key, this.todos);
		this.onChanges.forEach(cb => cb());
	}

	addTodo(title) {
		this.todos = this.todos.concat({
			id: uuid(),
			title,
			completed: false
		});
		this.inform();
	}

	toggleAll(completed) {
		this.todos = this.todos.map(todo => _extends({}, todo, { completed }));
		this.inform();
	}

	toggle(todoToToggle) {
		this.todos = this.todos.map(todo => todo !== todoToToggle ? todo : _extends({}, todo, { completed: !todo.completed }));
		this.inform();
	}

	destroy(todo) {
		this.todos = this.todos.filter(t => t !== todo);
		this.inform();
	}

	save(todoToSave, title) {
		this.todos = this.todos.map(todo => todo !== todoToSave ? todo : _extends({}, todo, { title }));
		this.inform();
	}

	clearCompleted() {
		this.todos = this.todos.filter(todo => !todo.completed);
		this.inform();
	}
}

class TodoFooter extends Component {
	render() {
		let { nowShowing, count, completedCount, onClearCompleted } = this.props;
		return createElement(
			'footer',
			{ className: 'footer' },
			createElement(
				'span',
				{ className: 'todo-count' },
				createElement(
					'strong',
					null,
					count
				),
				' ',
				pluralize(count, 'item'),
				' left'
			),
			createElement(
				'ul',
				{ className: 'filters' },
				createElement(
					'li',
					null,
					createElement(
						'a',
						{ href: '#/', className: nowShowing == 'all' && 'selected' },
						'All'
					)
				),
				' ',
				createElement(
					'li',
					null,
					createElement(
						'a',
						{ href: '#/active', className: nowShowing == 'active' && 'selected' },
						'Active'
					)
				),
				' ',
				createElement(
					'li',
					null,
					createElement(
						'a',
						{ href: '#/completed', className: nowShowing == 'completed' && 'selected' },
						'Completed'
					)
				)
			),
			completedCount > 0 && createElement(
				'button',
				{ className: 'clear-completed', onClick: onClearCompleted },
				'Clear completed'
			)
		);
	}
}

const ESCAPE_KEY = 27;
const ENTER_KEY$1 = 13;

class TodoItem extends Component {
	constructor(...args) {
		var _temp;

		return _temp = super(...args), this.handleSubmit = () => {
			let { onSave, onDestroy, todo } = this.props,
			    val = this.state.editText.trim();
			if (val) {
				onSave(todo, val);
				this.setState({ editText: val });
			} else {
				onDestroy(todo);
			}
		}, this.handleEdit = () => {
			let { onEdit, todo } = this.props;
			onEdit(todo);
			this.setState({ editText: todo.title });
		}, this.toggle = e => {
			let { onToggle, todo } = this.props;
			onToggle(todo);
			e.preventDefault();
		}, this.handleKeyDown = e => {
			if (e.which === ESCAPE_KEY) {
				let { todo } = this.props;
				this.setState({ editText: todo.title });
				this.props.onCancel(todo);
			} else if (e.which === ENTER_KEY$1) {
				this.handleSubmit();
			}
		}, this.handleDestroy = () => {
			this.props.onDestroy(this.props.todo);
		}, this.updateEditText = e => {
			this.setState({ editText: e.target.value });
		}, _temp;
	}

	// shouldComponentUpdate({ todo, editing, editText }) {
	// 	return (
	// 		todo !== this.props.todo ||
	// 		editing !== this.props.editing ||
	// 		editText !== this.state.editText
	// 	);
	// }

	render() {
		let { todo: { title, completed }, onToggle, onDestroy, editing } = this.props;
		let { editText } = this.state;
		let className = completed ? "completed" : "";
		className += editing ? " editing" : "";

		return createElement(
			"li",
			{ className: className },
			createElement(
				"div",
				{ className: "view" },
				createElement("input", {
					className: "toggle",
					type: "checkbox",
					checked: completed,
					onChange: this.toggle
				}),
				createElement(
					"label",
					{ onDblClick: this.handleEdit },
					title
				),
				createElement("button", { className: "destroy", onClick: this.handleDestroy })
			),
			editing && createElement("input", {
				className: "edit",
				autoFocus: true,
				value: editText,
				onBlur: this.handleSubmit,
				onInput: this.updateEditText,
				onKeyDown: this.handleKeyDown
			})
		);
	}
}

const ENTER_KEY = 13;

const FILTERS = {
	all: todo => true,
	active: todo => !todo.completed,
	completed: todo => todo.completed
};

class App extends Component {
	constructor() {
		super();

		this.handleNewTodoKeyDown = e => {
			if (e.keyCode !== ENTER_KEY) return;
			e.preventDefault();

			let val = this.state.newTodo.trim();
			if (val) {
				this.model.addTodo(val);
				this.setState({ newTodo: '' });
			}
		};

		this.updateNewTodo = e => {
			this.setState({ newTodo: e.target.value });
		};

		this.toggleAll = event => {
			let checked = event.target.checked;
			this.model.toggleAll(checked);
		};

		this.toggle = todo => {
			this.model.toggle(todo);
		};

		this.destroy = todo => {
			this.model.destroy(todo);
		};

		this.edit = todo => {
			this.setState({ editing: todo.id });
		};

		this.save = (todoToSave, text) => {
			this.model.save(todoToSave, text);
			this.setState({ editing: null });
		};

		this.cancel = () => {
			this.setState({ editing: null });
		};

		this.clearCompleted = () => {
			this.model.clearCompleted();
		};

		this.model = new TodoModel('didact-todos', () => this.setState({}));
		addEventListener('hashchange', this.handleRoute.bind(this));
		let nowShowing = this.getRoute();
		this.state = {
			nowShowing,
			newTodo: "",
			editing: ""
		};
	}

	getRoute() {
		let nowShowing = String(location.hash || '').split('/').pop();
		if (!FILTERS[nowShowing]) {
			nowShowing = 'all';
		}
		return nowShowing;
	}

	handleRoute() {
		let nowShowing = this.getRoute();
		this.setState({ nowShowing });
	}

	render() {
		let { nowShowing = ALL_TODOS, newTodo, editing } = this.state;
		let { todos } = this.model,
		    shownTodos = todos.filter(FILTERS[nowShowing]),
		    activeTodoCount = todos.reduce((a, todo) => a + (todo.completed ? 0 : 1), 0),
		    completedCount = todos.length - activeTodoCount;

		return createElement(
			'div',
			null,
			createElement(
				'header',
				{ className: 'header' },
				createElement(
					'h1',
					null,
					'todos'
				),
				createElement('input', {
					className: 'new-todo',
					placeholder: 'What needs to be done?',
					value: newTodo,
					onKeyDown: this.handleNewTodoKeyDown,
					onInput: this.updateNewTodo,
					autoFocus: true
				})
			),
			todos.length ? createElement(
				'section',
				{ className: 'main' },
				createElement('input', {
					className: 'toggle-all',
					type: 'checkbox',
					onChange: this.toggleAll,
					checked: activeTodoCount === 0
				}),
				createElement(
					'ul',
					{ className: 'todo-list' },
					shownTodos.map(todo => createElement(TodoItem, {
						todo: todo,
						onToggle: this.toggle,
						onDestroy: this.destroy,
						onEdit: this.edit,
						editing: editing === todo.id,
						onSave: this.save,
						onCancel: this.cancel
					}))
				)
			) : null,
			activeTodoCount || completedCount ? createElement(TodoFooter, {
				count: activeTodoCount,
				completedCount: completedCount,
				nowShowing: nowShowing,
				onClearCompleted: this.clearCompleted
			}) : null
		);
	}
}

// import 'todomvc-common';
// import 'todomvc-common/base.css';
// import 'todomvc-app-css/index.css';

render(createElement(App, null), document.querySelector('.todoapp'));

}());
//# sourceMappingURL=app.js.map
