import { h, text, app } from "https://esm.run/hyperapp"

const GIPHY = "https://api.giphy.com/v1/gifs/search"

/// Http FX

const request = (dispatch, { parse, url, err, ok, ...props }) =>
  fetch(url, props)
    .then((resp) => (resp.ok ? resp : Promise.reject(resp)))
    .then((resp) => resp[parse]())
    .then((result) => dispatch(ok(result)))
    .catch((error) => dispatch(err(error)))

const getJson = (url, props) => [request, { parse: "json", url, ...props }]

/// Html helpers

const input = (oninput, props) =>
  h("input", { oninput: (_, { target }) => oninput(target.value), ...props })

const title = (title) => h("h1", {}, text(title))
const main = (...children) => h("main", {}, children)
const img = (src) => h("img", { src })
const p = (p) => h("p", {}, text(p))

/// Main

const downloadGif = (query) =>
  getJson(`${GIPHY}?q=${query}&api_key=${env.dataset.apikey}`, {
    err: (error) => [GotError, error.error],
    ok: ({ data: [first] }) => [GotURL, first && first.images.original.url],
  })

const GotError = (state, error) => ({
  ...state,
  isFetching: false,
  error: error ? error : "Unexpected error, try again later?",
  url: "",
})

const GotURL = (state, url) => ({
  ...state,
  isFetching: false,
  url: state.query ? url : "",
})

const GetURL = (state, query) => [
  { ...state, isFetching: true, query, error: "", url: "" },
  downloadGif(query),
]

app({
  init: {
    isFetching: false,
    query: "",
    error: "",
    url: "",
  },
  view: (state) =>
    main(
      title("GIF Search 💬💁‍♂️"),
      input((value) => [GetURL, value.trim()], {
        placeholder: "Search GIFs...",
        type: "text",
      }),
      state.error ? p(state.error) : img(state.url)
    ),
  node: document.getElementById("app"),
})
