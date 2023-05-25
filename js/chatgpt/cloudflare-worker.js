// https://github.com/barretlee/cloudflare-proxy

addEventListener("fetch", (event) => {
  event.respondWith(handleRequest(event.request));
});

async function handleRequest(request) {
  const url = new URL(request.url);
  const fetchAPI = request.url.replace(url.host, 'api.openai.com');
  let body;
  if (request.method === 'POST') {
    body = await request.json();
  }
  const authKey = request.headers.get('Authorization');
  if (!authKey) {
    return new Response("Not allowed", {
      status: 403
    });
  }

  const payload = {
    method: request.method,
    headers: {
      "Content-Type": "application/json",
      Authorization: authKey,
    },
    body: typeof body === 'object' ? JSON.stringify(body) : '{}',
  };

  if (['HEAD', 'GET'].includes(request.method)) {
    delete payload.body;
  }

  if (body && body.stream && body.stream === false) {
    const results = await response.json();
    return new Response(JSON.stringify(results), {
      status: response.status,
      headers: {
        "Content-Type": "application/json",
      },
    });
  } else {
    const response = await fetch(fetchAPI, payload);
    return new Response(response.body, {
      status: response.status,
      statusText: response.statusText,
      headers: response.headers,
    });
  }
}
