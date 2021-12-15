const http = new class {
    //constructor() { }

    get(url) {
        return this.request('GET', url, null);
    }

    post(url, data) {
        return this.request('POST', url, data);
    }

    put(url, data) {
        return this.request('PUT', url, data);
    }

    delete(url, data) {
        return this.request('DELETE', url, data);
    }

    async request(method, url, data) {
        let headers = {'content-type': 'application/json'};
        let token = ''; //window.localStorage.getItem('token')

        let config = {
            method: method,
            headers: headers,
        }

        if (data) {
            config.body = JSON.stringify(data)
        }

        if (token) {
            headers.Authorization = `Bearer ${token}`
        }

        try {
            let res = await fetch(url, config);
            return await res.json();
        } catch (error) {
            console.log('Request Failed', error);
        }
    }
}

const url = 'http://localhost:9000/api/date/exclude';

http.get(url).then(json => console.log(json));

var id;

await http.post(url, {
    date: '2020-02-23',
    note: 'TEST',
    user: 'root'
}).then(json => {
    console.log(json);
    id = json.data.id;
});

console.log(id);

await http.put(url, {
    id:   id,
    date: '2020-02-22',
    note: 'TEST-ROOT',
}).then(json => console.log(json));

http.delete(url, {id: id}).then(json => console.log(json));
