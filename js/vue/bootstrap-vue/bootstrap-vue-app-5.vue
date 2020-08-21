https://codesandbox.io/examples/package/bootstrap-vue

<template>		https://codesandbox.io/s/0x608nxo10
<div id="app" style="text-align: left;margin-bottom: 200px">
  <b-container>
    <b-row>
      <div class="col-12 mb-4">
        <h2>vue-multiselect and bootstrap-vue</h2>
      </div>
      <div class="col-8 row">
        <div class="col-12">
          <Multiselect 
            v-model="dataArray" 
            :options="options"
            label="label"
            track-by="value"
            multiple
          ></Multiselect>
        </div>
        <div class="col-8 mt-4">
          <h5>Variant "primary"</h5>
          <Multiselect 
            v-model="dataArray" 
            :options="options"
            label="label"
            track-by="value"
            multiple
            variant="primary"
          ></Multiselect>
        </div>
        <div class="col-4 mt-4">
        <h5>Narrow</h5>
          <Multiselect 
            v-model="dataArray" 
            :options="options"
            label="label"
            multiple
            track-by="value"
          ></Multiselect>
        </div>
      </div>
      <div class="col-4">
        <pre>{{dataArray}}</pre>
      </div>
    </b-row>
    <hr>
    <b-row>
      <div class="col-8 row">
        <div class="col-8 mt-4">
          <h5>Singular</h5>
          <Multiselect 
          v-model="dataSingle" 
          :options="options"
          label="label"
          track-by="value"
          ></Multiselect>
        </div>
        <div class="col-4 mt-4">
          <h5>Narrow</h5>
          <Multiselect 
            v-model="dataSingle" 
            :options="options"
            label="label"
            track-by="value"
          ></Multiselect>
        </div>
      </div>
      <div class="col-4 mt-4">
        <pre>{{dataSingle}}</pre>
      </div>
    </b-row>
    <hr>
    <b-row>
      <div class="col-8 mt-4 row">
        <div class="col-12">
          <h5>Taggable and input group</h5>
          <b-input-group>
          <b-input-group-text slot="prepend">
              Taggable
          </b-input-group-text>
            <Multiselect 
              v-model="dataArrayTaggable" 
              :options="options"
              label="label"
              track-by="value"
              taggable
              @tag="addTag"
              multiple
            ></Multiselect>
          </b-input-group>
        </div>
      </div>
      <div class="col-4 mt-4">
        <pre>{{dataArrayTaggable}}</pre>
      </div>
    </b-row>
    <hr>
    <b-row>
      <div class="col-8 mt-4 row">
        <div class="col-12">
          <h5>Groups</h5>
          <Multiselect 
            v-model="dataGroups" 
            :options="optionsGroups"
            label="name"
            track-by="value"
            groupValues="components"
            groupLabel="library"
            :groupSelect="true"
            :close-on-select="false"
            multiple
          ></Multiselect>
        </div>
      </div>
      <div class="col-4 mt-4">
        <pre>{{dataGroups}}</pre>
      </div>
    </b-row>
  </b-container>
</div>
</template>

<script>
import Multiselect from "./components/Multiselect";

export default {
  name: "App",
  components: {
    Multiselect
  },
  data() {
    return {
      dataSingle: null,
      dataArray: [{ value: "dark", label: "Dark" }],
      dataArrayTaggable: [],
      dataGroups: [],
      options: [
        { value: "dark", label: "Dark" },
        { value: "light", label: "Light" },
        { value: "heavy", label: "Heavy" },
        { value: "left", label: "Left" }
      ],
      optionsGroups: [
        {
          library: "Bootstrap Vue",
          components: [
            { value: "b-select", name: "bSelect" },
            { value: "b-dropdown", name: "bDropdown" }
          ]
        },
        {
          library: "Vue Multiselect",
          components: [{ value: "multiselect", name: "Multiselect" }]
        },
        {
          library: "Bootstrap",
          components: [
            { value: "styles", name: "Styles" },
            { value: "popper.js", name: "Popper.js" }
          ]
        }
      ]
    };
  },
  methods: {
    addTag(t) {
      let newoption = { value: t, label: t };
      this.options.push(newoption);
      this.dataArrayTaggable.push(newoption);
    }
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
