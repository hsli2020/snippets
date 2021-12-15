import { h, text, app } from "https://esm.run/hyperapp"

const div = (props, children) => h("div", props, children)
const main = (...children) => h("main", {}, children)
const button = (props, ...children) => h("button", props, children)

const computer = {
  "+": (a, b) => a + b,
  "-": (a, b) => a - b,
  "×": (a, b) => a * b,
  "÷": (a, b) => a / b,
}

const initialState = {
  fn: "",
  carry: 0,
  value: 0,
  hasCarry: false,
}

const Clear = () => initialState

const NewDigit = (state, number) => ({
  ...state,
  hasCarry: false,
  value: (state.hasCarry ? 0 : state.value) * 10 + number,
})

const NewFn = (state, fn) => ({
  ...state,
  fn,
  hasCarry: true,
  carry: state.value,
  value:
    state.hasCarry || !state.fn
      ? state.value
      : computer[state.fn](state.carry, state.value),
})

const Equal = (state) => ({
  ...state,
  hasCarry: true,
  carry: state.hasCarry ? state.carry : state.value,
  value: state.fn
    ? computer[state.fn](
        state.hasCarry ? state.value : state.carry,
        state.hasCarry ? state.carry : state.value
      )
    : state.value,
})

const displayView = (value) => div({ class: "display" }, text(value))

const keysView = (...children) => div({ class: "keys" }, children)

const fnView = (props) =>
  props.keys.map((fn) =>
    button({ class: "function", onclick: [NewFn, fn] }, text(fn))
  )

const digitsView = ({ digits }) =>
  digits.map((digit) =>
    button(
      { class: { zero: digit === 0 }, onclick: [NewDigit, digit] },
      text(digit)
    )
  )

const acView = button({ onclick: Clear }, text("AC"))
const eqView = button({ onclick: Equal, class: "equal" }, text("="))

app({
  init: initialState,
  view: (state) =>
    main(
      displayView(state.value),
      keysView(
        ...fnView({ keys: Object.keys(computer) }),
        ...digitsView({ digits: [7, 8, 9, 4, 5, 6, 1, 2, 3, 0] }),
        acView,
        eqView
      )
    ),
  node: document.getElementById("app"),
})
