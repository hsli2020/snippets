const localStorageKey = '__bookshelf_token__'

function client(endpoint, {body, ...customConfig} = {}) {
  const token = window.localStorage.getItem(localStorageKey)
  const headers = {'content-type': 'application/json'}

  if (token) {
    headers.Authorization = `Bearer ${token}`
  }

  const config = {
    method: body ? 'POST' : 'GET',
    ...customConfig,
    headers: {
      ...headers,
      ...customConfig.headers,
    },
  }

  if (body) {
    config.body = JSON.stringify(body)
  }

  return window
    .fetch(endpoint, config)
    .then(async response => {
      if (response.status === 401) {
        logout()
        window.location.assign(window.location)
        return
      }

      const data = await response.json()
      if (response.ok) {
        return data
      } else {
        return Promise.reject(data)
      }
    })
}

function logout() { window.localStorage.removeItem(localStorageKey) }

// fetch()第二个参数的完整 API 如下。
const response = fetch(url, {
    method: "GET",
    headers: { "Content-Type": "text/plain;charset=UTF-8" },
    body: undefined,
    referrer: "about:client",
    referrerPolicy: "no-referrer-when-downgrade",
    mode: "cors", 
    credentials: "same-origin",
    cache: "default",
    redirect: "follow",
    integrity: "",
    keepalive: false,
    signal: undefined
});

// Promise 可以使用 await 语法改写，使得语义更清晰。
async function getJSON() {
  let url = 'https://api.github.com/users/ruanyf';
  try {
    let response = await fetch(url);
    return await response.json();
  } catch (error) {
    console.log('Request Failed', error);
  }
}
