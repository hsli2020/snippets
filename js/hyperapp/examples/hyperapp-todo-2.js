import { h, text, app } from "https://esm.run/hyperapp"

/// Utilities

const isEmptyList = (list) => list.length === 0

const createTodoItem = (value) => ({
  isEditing: false,
  lastValue: "",
  value,
})

const preventDefault = (action) => (state, event) => [
  state,
  [(dispatch) => (event.preventDefault(), dispatch(action))],
]

const withPayload = (filter) => (_, payload) => filter(payload)
const targetValue = (action) => withPayload((e) => [action, e.target.value])

/// Html

const [main, ul, li, span, label] = [
  "main",
  "ul",
  "li",
  "span",
  "label",
].map((tag) => (children) => h(tag, {}, children))

const form = ({ onsubmit, ...props }, children) =>
  h("form", { ...props, onsubmit: preventDefault(onsubmit) }, children)

const checkbox = h("input", { type: "checkbox" })

const submit = (value) => h("input", { type: "submit", value })

const input = ({ oninput, ...props }) =>
  h("input", { type: "text", oninput: targetValue(oninput), ...props })

const todoItemView = (value) => label([checkbox, span(text(value))])

const todosView = ({ title, todos }) => [
  h("h1", {}, text(title)),
  ul(todos.map((todo) => li(todoItemView(todo.value)))),
]

/// Main

const NewValue = (state, value) => ({ ...state, value })

const NewTodo = (state) =>
  isEmptyList(state.todos)
    ? state
    : {
        ...state,
        todos: state.todos.concat(createTodoItem(state.value)),
        value: "",
      }

app({
  init: {
    todos: [createTodoItem("Learn Hyperapp")],
    value: "",
  },
  view: ({ todos, value }) =>
    main([
      ...todosView({ title: "To-do list ✏️", todos }),
      form({ onsubmit: NewTodo }, [
        input({
          value,
          oninput: NewValue,
        }),
        submit("Add new"),
      ]),
    ]),
  node: document.getElementById("app"),
})
