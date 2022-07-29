const axios = {
  get(url, { headers = {} }) {
    return new Promise((resolve, reject) => {
      const xhr = window.XMLHttpRequest
          ? new XMLHttpRequest() : new ActiveXObject("Microsoft.XMLHTTP");

      xhr.open("get", url);
      xhr.withCredentials = true;
      // Object.entries(headers).forEach(([key,value]) => xhr.setRequestHeader(key, value));

      xhr.send(null);
      xhr.onreadystatechange = function () {
        if (this.readyState === 4) {
          const data = JSON.parse(this.responseText);
          return resolve({ data })
        }
      }
      xhr.onerror = function (error) { return reject(error); }
    })
  },
  post(url, data, { headers = {} }) {
    return new Promise((resolve, reject) => {
      const xhr = window.XMLHttpRequest
          ? new XMLHttpRequest() : new ActiveXObject("Microsoft.XMLHTTP");

      xhr.open("post", url);
      xhr.withCredentials = true;
      xhr.setRequestHeader("Content-Type","application/json");

      // Object.entries(headers).forEach(([key,value]) => xhr.setRequestHeader(key, value));

      xhr.send(JSON.stringify(data));
      xhr.onreadystatechange = function () {
        if (this.readyState === 4) {
          const data = JSON.parse(this.responseText);
          return resolve({ data })
        }
      }
      xhr.onerror = function (error) { return reject(error); }
    })
  }
}

async function run() {
    const spinList = await getSpinList();
    //test
    // const spinList = [ { spinType: 'tweet-spin', spinMultiplier: 1 }, 
    // { spinType: 'super-spin', spinMultiplier: 1 } ]  ;

    log('spinList -> ', spinList);

    if (!spinList.length) { log('no spin'); return; }

    // [ { spinType: 'tweet-spin', spinMultiplier: 1 } ]  
    const splitSpinList = [];

    spinList.forEach(({ spinType, spinMultiplier }) => {
        new Array(spinMultiplier).fill(0).forEach(() => {
            splitSpinList.push({ spinType, spinMultiplier: 1 })
        })
    })

    log('splitSpinList -> ', splitSpinList);

    for (let i = 0; i < splitSpinList.length; i++) {
        const item = splitSpinList[i];
        const isSuper = item.spinType === 'super-spin';
        log(`${i} send spinType ${item.spinType} ...`);

        try {
            const url = `https://waitlist-api.prod.blast.io/v1${
                isSuper ? '/spins/sample-superspin' : '/spins/execute'}`
            const spinRes = await axios.post(url, isSuper ? {} : splitSpinList[i], {
                // headers: { 'cookie': `accessToken=${accessToken}`, }
            });

            if(!isSuper) {
                const { data: { result: { numberPoints, rarity } } } = spinRes;
                log(`${i} => rarity ${rarity} => numberPoints ${numberPoints}`);
            }else {
                const { data: { spinMultiplier = 1, spinType = 'super-spin' } } = spinRes;
                
                const url = `https://waitlist-api.prod.blast.io/v1/spins/execute`
                const spinRes1 = await axios.post(url, {
                    spinMultiplier: spinMultiplier || 1, 
                    spinType: spinType || 'super-spin' }, {});

                const { data: { result: { numberPoints, rarity } } } = spinRes1;
                log(`${i} => rarity ${rarity} => numberPoints ${numberPoints}`);
            }
        } catch (error) {
            log('❌ spin failed',
                error.response?.data?.message || 'unknown reason', error.toString());
        }
    }
}

function getSpinList() {
    log('get spinList ...');

    return new Promise(resolve => {
        const request = async () => {
            try {
                const { data: { spins } } = 
                    await axios.get('https://waitlist-api.prod.blast.io/v1/user/dashboard', {
                    // headers: { 'cookie': `accessToken=${accessToken}`, }
                });

                const spinList = Object.entries(spins).map(([key, value]) => ({
                    spinType: key, spinMultiplier: value
                })).filter(item => item.spinMultiplier > 0)

                resolve(spinList);
            } catch (error) {
                log('get spinList error', error.toString(), ' retring...');
                setTimeout(request, 2000);
            }
        }
        request();
    })
}

function log(...msgs) { const date = new Date(); console.log(`${getDateTime()}`, ...msgs); }

function getDateTime() {
    const date = new Date();
    const now = date.getTime();
    const ms = String(now).slice(-3);

    return `[${date.toLocaleString()}.${ms}]`;
}

run();
