const fetch = (...args) => console.log(...args) // mock

function httpRequest(url, method, data) {
    const init = { method }
    switch (method) {
        case 'GET':
            if (data) url = `${url}?${new URLSearchParams(data)}`
            break
        case 'POST':
        case 'PUT':
        case 'PATCH':
            init.body = JSON.stringify(data)
    }
    return fetch(url, init)
}

function generateAPI(url) {
    // a hack, so we can use field either as property or a method
    const callable = () => {}
    callable.url = url

    return new Proxy(callable, { 
        get({ url }, propKey) {
            return (['GET', 'POST', 'PUT', 'DELETE', 'PATCH'].includes(propKey.toUpperCase())) ?
                (data) => httpRequest(url, propKey.toUpperCase(), data) :
                generateAPI(`${url}/${propKey}`)
        }, 
        apply({ url }, thisArg, [arg] = []) {
            return generateAPI( arg ? `${url}/${arg}` : url) 
        }
    })
}

// example usage
const GameAPI = generateAPI('game_api')

GameAPI.get() // GET /game_api
GameAPI.clans.get() // GET /game_api/clans
GameAPI.clans(7).get() // GET /game_api/clans/7
GameAPI.clans(7).whatever.delete() // DELETE /game_api/clans/7/whatever
GameAPI.clans.put({ whatever: 1 })

// GET game_api/tiles/public/static/3/4/2.json?turn=37038&games=wot
GameAPI.tiles.public.static(3)(4)(`${2}.json`).get({ turn: 37, games: 'wot' }) 
