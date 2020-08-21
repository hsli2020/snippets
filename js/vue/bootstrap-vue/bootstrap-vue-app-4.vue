https://codesandbox.io/examples/package/bootstrap-vue

<template>		https://codesandbox.io/s/jp4xmmv1v
    <div>
      <div v-if="!hasRecords" style="text-align: center"><br><br>LOADING DATA...</div>
      <div v-if="hasRecords">
        <b-table :items="records" :fields="column" striped hover :current-page="currentPage" :per-page="perPage">
          <template slot="HEAD_selected" slot-scope="data">
              <input type="checkbox" @click.stop v-model="selectAll" @change="toggleSelectAll" />
          </template>
          <template slot="selected" slot-scope="data">
              <input type="checkbox" v-model="data.item.selected" @change="selectRow(data.item)" />_rowVariant = {{ data.item._rowVariant }}
          </template>
        </b-table>
        <b-row>
          <b-col md="6" class="my-1">
              <b-pagination :total-rows="totalRows" :per-page="perPage" v-model="currentPage" class="my-0" />
          </b-col>
          <b-col md="6" class="my-1">
              <b-form-group horizontal label="Per page" class="mb-0">
                  <b-form-select :options="pageOptions" v-model="perPage" />
              </b-form-group>
          </b-col>
        </b-row>
      </div>
    </div>
</template>

<script>
import json from "../data/data.json";

export default {
  data() {
    return {
      selectAll: false,
      records: [],
      perPage: 5,
      currentPage: 1,
      pageOptions: [5, 10, 15],
      column: [
        {
          key: "selected",
          sortable: false,
          label: "",
          class: "options-column"
        },
        {
          key: "name",
          sortable: true,
          label: "Log File Name"
        },
        {
          key: "fileSize",
          sortable: true,
          label: "File Size",
          class: "text-right options-column"
        },
        {
          key: "lastModified",
          sortable: true,
          label: "Last Modified Date",
          class: "text-right options-column"
        },
        {
          sortable: false,
          label: "Options",
          class: "options-column"
        }
      ]
    };
  },
  computed: {
    hasRecords() {
      return this.records.length > 0;
    },
    totalRows() {
      return this.records.length;
    }
  },
  methods: {
    selectRow(item) {
      if (item.selected) {
        item._rowVariant = "info";
      } else {
        item._rowVariant = "default";
        if (this.selectAll) {
          this.selectAll = false;
        }
      }
    },
    toggleSelectAll() {
      if (this.selectAll) {
        for (var i = 0; i < this.records.length; i++) {
          var updatingItem = this.records[i];
          updatingItem.selected = true;
          updatingItem._rowVariant = "info";
          this.$set(this.records, i, updatingItem);
        }
      } else {
        for (var i = 0; i < this.records.length; i++) {
          var updatingItem = this.records[i];
          updatingItem.selected = false;
          updatingItem._rowVariant = "default";
          this.$set(this.records, i, updatingItem);
        }
      }
    }
  },
  components: {},
  mounted() {
    var vm = this;
    setTimeout(function() {
      vm.records = json;
    }, 1000);
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
