// index.js
import React from "react";
import Vue from "vue"; // "vue-toy";
import DemoGrid from "./DemoGrid";
import "./styles.css";

const vue = new Vue({
  el: "#root",
  render() {
    return (
      <>
        <form id="search">
          Search{" "}
          <input
            name="query"
            value={this.searchQuery}
            onChange={e => (this.searchQuery = e.target.value)}
          />
        </form>
        <DemoGrid
          heroes={this.gridData}
          columns={this.gridColumns}
          filterKey={this.searchQuery}
        />
        <br />
        <a target="_blank" href="https://github.com/bplok20010/vue-toy">
          View Github
        </a>
      </>
    );
  },
  data: () => {
    return {
      searchQuery: "",
      gridColumns: ["name", "power"],
      gridData: [
        { name: "Chuck Norris", power: Infinity },
        { name: "Bruce Lee", power: 9000 },
        { name: "Jackie Chan", power: 7000 },
        { name: "Jet Li", power: 8000 }
      ]
    };
  }
});

console.log(vue);

// DemoGrid.js
import React from "react";
import classnames from "classnames";
import Vue from "vue-toy";

function capitalize(str) {
  return str.charAt(0).toUpperCase() + str.slice(1);
}

const DemoGrid = Vue.component({
  props: ["heroes", "columns", "filterKey"],
  data: function() {
    var sortOrders = {};
    this.columns.forEach(function(key) {
      sortOrders[key] = 1;
    });
    return {
      sortKey: "",
      sortOrders: sortOrders
    };
  },
  computed: {
    filteredHeroes: function() {
      var sortKey = this.sortKey;
      var filterKey = this.filterKey && this.filterKey.toLowerCase();
      var order = this.sortOrders[sortKey] || 1;
      var heroes = this.heroes;
      if (filterKey) {
        heroes = heroes.filter(function(row) {
          return Object.keys(row).some(function(key) {
            return (
              String(row[key])
                .toLowerCase()
                .indexOf(filterKey) > -1
            );
          });
        });
      }
      if (sortKey) {
        heroes = heroes.slice().sort(function(a, b) {
          a = a[sortKey];
          b = b[sortKey];
          return (a === b ? 0 : a > b ? 1 : -1) * order;
        });
      }
      return heroes;
    }
  },
  methods: {
    sortBy: function(key) {
      this.sortKey = key;
      this.sortOrders[key] = this.sortOrders[key] * -1;
    }
  },
  render() {
    return (
      <table>
        <thead>
          <tr>
            {this.columns.map(key => (
              <th
                onClick={() => this.sortBy(key)}
                className={classnames({
                  active: this.sortKey === key
                })}
              >
                {capitalize(key)}
                <span
                  className={classnames({
                    arrow: true,
                    asc: this.sortOrders[key] > 0,
                    dsc: this.sortOrders[key] <= 0
                  })}
                />
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {this.filteredHeroes.map((entry, index) => (
            <tr key={index}>
              {this.columns.map(key => (
                <td key={key}>{entry[key]}</td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    );
  }
});

export default DemoGrid;

// style.css
body {
  font-family: Helvetica Neue, Arial, sans-serif;
  font-size: 14px;
  color: #444;
}

.demo {
  width: 500px;
  margin: 30px auto;
}

table {
  border: 2px solid #42b983;
  border-radius: 3px;
  background-color: #fff;
}

th {
  background-color: #42b983;
  color: rgba(255, 255, 255, 0.66);
  cursor: pointer;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

td {
  background-color: #f9f9f9;
}

th,
td {
  min-width: 120px;
  padding: 10px 20px;
}

th.active {
  color: #fff;
}

th.active .arrow {
  opacity: 1;
}

.arrow {
  display: inline-block;
  vertical-align: middle;
  width: 0;
  height: 0;
  margin-left: 5px;
  opacity: 0.66;
}

.arrow.asc {
  border-left: 4px solid transparent;
  border-right: 4px solid transparent;
  border-bottom: 4px solid #fff;
}

.arrow.dsc {
  border-left: 4px solid transparent;
  border-right: 4px solid transparent;
  border-top: 4px solid #fff;
}
