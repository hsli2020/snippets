import { h, text, app } from "https://esm.run/hyperapp"

const withPayload = (filter) => (_, payload) => filter(payload)

const input = (oninput, props) => h("input", { oninput, ...props })
const title = (title) => h("h1", {}, text(title))
const main = (...children) => h("main", {}, children)

const SetText = (state, message) => ({ ...state, message })

app({
  init: { message: "" },
  view: (state) =>
    main(
      title(state.message.trim() === "" ? "🤷‍♂️" : state.message),
      input(
        withPayload(({ target }) => [SetText, target.value]), { 
          placeholder: "Type in something...",
          value: state.message,
          type: "text",
        }
      )
    ),
  node: document.getElementById("app"),
})
