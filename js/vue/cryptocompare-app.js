// script.js  https://github.com/rdegges/cryptocompare

// The API we're using for grabbing metadata about each cryptocurrency
// (including logo images). The service can be found at:
// https://www.cryptocompare.com/api/
let CRYPTOCOMPARE_API_URI = "https://min-api.cryptocompare.com";
let CRYPTOCOMPARE_URI = "https://www.cryptocompare.com";

// The API we're using for grabbing cryptocurrency prices.  The service can be
// found at: https://coinmarketcap.com/api/
let COINMARKETCAP_API_URI = "https://api.coinmarketcap.com";

let app = new Vue({
  el: "#app",
  data: {
    coins: {},
    coinData: {}
  },
  methods: {
    /**
     * Load up all cryptocurrency data.  This data is used to find what logos
     * each currency has, so we can display things in a friendly way.
     */
    getCoinData: function() {
      let self = this;

      axios.get(CRYPTOCOMPARE_API_URI + "/data/all/coinlist")
        .then((resp) => {
          this.coinData = resp.data.Data;
          this.getCoins();
        })
        .catch((err) => {
          this.getCoins();
          console.error(err);
        });
    },

    /**
     * Get the top 10 cryptocurrencies by value.  This data is refreshed each 5
     * minutes by the backing API service.
     */
    getCoins: function() {
      let self = this;

      axios.get(COINMARKETCAP_API_URI + "/v2/ticker/?limit=100")
        .then((resp) => {
          this.coins = resp.data.data;
        })
        .catch((err) => {
          console.error(err);
        });
    },

    /**
     * Given a cryptocurrency ticket symbol, return the currency's logo image.
     */
    getCoinImage: function(symbol) {
      // These two symbols don't match up across API services. I'm manually
      // replacing these here so I can find the correct image for the currency.
      //
      // In the future, it would be nice to find a more generic way of searching
      // for currency images
      symbol = (symbol === "MIOTA" ? "IOT" : symbol);
      symbol = (symbol === "VERI" ? "VRM" : symbol);
      
      try {
          return CRYPTOCOMPARE_URI + this.coinData[symbol].ImageUrl;
      } catch(err) {
        console.log(err);
      }
    },

    /**
     * Return a CSS color (either red or green) depending on whether or
     * not the value passed in is negative or positive.
     */
    getColor: (num) => {
      return num > 0 ? "color:green;" : "color:red;";
    },
  },

  /**
   * Using this lifecycle hook, we'll populate all of the cryptocurrency data as
   * soon as the page is loaded a single time.
   */
  created: function () {
    this.getCoinData();
  }
});

/**
 * Once the page has been loaded and all of our app stuff is working, we'll
 * start polling for new cryptocurrency data every minute.
 *
 * This is sufficiently dynamic because the API's we're relying on are updating
 * their prices every 5 minutes, so checking every minute is sufficient.
 */
let UPDATE_INTERVAL = 60 * 1000;
setInterval(() => { app.getCoins(); }, UPDATE_INTERVAL);

// index.html
<div id="app">
  <table class="table table-hover">
    <thead>
      <tr>
        <td>Rank</td>
        <td>Name</td>
        <td>Symbol</td>
        <td>Price (USD)</td>
        <td>1H</td>
        <td>1D</td>
        <td>1W</td>
        <td>Market Cap (USD)</td>
    </thead>
    <tbody>
      <tr v-cloak v-for="coin in coins">
        <td>{{ coin.rank }}</td>
        <td><img v-bind:src="getCoinImage(coin.symbol)"> {{ coin.name }}</td>
        <td>{{ coin.symbol }}</td>
        <td>{{ coin.quotes.USD.price  | currency }}</td>
        <td v-bind:style="getColor(coin.quotes.USD.percent_change_1h)">
          <span v-if="coin.quotes.USD.percent_change_1h > 0">+</span>
            {{ coin.quotes.USD.percent_change_1h }}%
        </td>
        <td v-bind:style="getColor(coin.quotes.USD.percent_change_24h)">
          <span v-if="coin.quotes.USD.percent_change_24h > 0">+</span>
            {{ coin.quotes.USD.percent_change_24h }}%
        </td>
        <td v-bind:style="getColor(coin.quotes.USD.percent_change_7d)">
          <span v-if="coin.quotes.USD.percent_change_7d > 0">+</span>
            {{ coin.quotes.USD.percent_change_7d }}%
        </td>
        <td>{{ coin.quotes.USD.market_cap | currency }}</td>
      </tr>
    </tbody>
  </table>
</div>
