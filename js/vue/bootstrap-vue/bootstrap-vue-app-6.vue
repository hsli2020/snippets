https://codesandbox.io/examples/package/bootstrap-vue

<template>		https://codesandbox.io/s/lyzqn4m659
  <div id="app">
    <b-container fluid
                 tag="main"
                 id="app"
                 v-cloak>
    <b-navbar toggleable="md"
              class="mb-3"
              type="light"
              variant="light">
      <b-nav-toggle target="nav_collapse"></b-nav-toggle>
      <b-navbar-brand href="#">NavBar</b-navbar-brand>
      <b-collapse is-nav
                  id="nav_collapse">
        <b-navbar-nav>
          <b-nav-item href="#">Link</b-nav-item>
          <b-nav-item href="#"
                      disabled>Disabled</b-nav-item>
        </b-navbar-nav>
        <!-- Right aligned nav items -->
        <b-navbar-nav class="ml-auto">
          <b-nav-form>
            <b-form-input size="sm"
                          class="mr-sm-2"
                          type="text"
                          placeholder="Search" />
            <b-button size="sm"
                      class="my-2 my-sm-0"
                      type="submit">Search</b-button>
          </b-nav-form>
          <b-nav-item-dropdown text="Lang"
                               right>
            <b-dropdown-item href="#">EN</b-dropdown-item>
            <b-dropdown-item href="#">ES</b-dropdown-item>
            <b-dropdown-item href="#">RU</b-dropdown-item>
            <b-dropdown-item href="#">FA</b-dropdown-item>
          </b-nav-item-dropdown>
          <b-nav-item-dropdown right>
            <!-- Using button-content slot -->
            <template slot="button-content">
              <em>User</em>
            </template>
            <b-dropdown-item href="#">Profile</b-dropdown-item>
            <b-dropdown-item href="#">Sign out</b-dropdown-item>
          </b-nav-item-dropdown>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
    <header class="text-center">
      <a href="https://github.com/alexsasharegan/vue-transmit"
         target="_blank">
        <img class="img--logo"
             src="https://raw.githubusercontent.com/alexsasharegan/vue-transmit/master/docs/logo.png">
      </a>
      <h1 class="mb-5"><code>&lt;vue-transmit&gt;</code></h1>
      <b-form-checkbox class="mb-3"
                       v-model="options.uploadMultiple">
        Upload Multiple
      </b-form-checkbox>
    </header>
    <b-container tag="main">
      <b-row>
        <b-col cols="3"></b-col>
        <b-col cols="6">
          <vue-transmit ref="uploader"
                        upload-area-classes="vh-20"
                        drag-class="dragging"
                        v-bind="options"
                        @added-file="onAddedFile"
                        @success="onUploadSuccess"
                        @success-multiple="onUploadSuccessMulti"
                        @timeout="onError"
                        @error="onError">
            <flex-col align-v="center"
                      class="h-100">
              <flex-row align-h="center">
                <b-btn variant="primary"
                       @click="triggerBrowse"
                       class="w-50">
                  Upload Files
                </b-btn>
              </flex-row>
            </flex-col>
            <template slot-scope="{ uploadingFiles }"
                      slot="files">
              <flex-row v-for="file in uploadingFiles"
                        :key="file.id"
                        align-v="center"
                        no-wrap
                        class="w-100 my-5"
                        style="height: 100px;">
                <img v-show="file.dataUrl"
                     :src="file.dataUrl"
                     :alt="file.name"
                     class="img-fluid w-25">
                <b-progress :value="file.upload.progress"
                            show-progress
                            :precision="2"
                            :variant="file.upload.progress === 100 ? 'success' : 'warning'"
                            :animated="file.upload.progress === 100"
                            class="ml-2 w-100"></b-progress>
              </flex-row>
            </template>
          </vue-transmit>
        </b-col>
      </b-row>
      <b-row class="my-3">
        <b-col v-for="file in files"
               :key="file.id"
               cols="4">
          <b-card :title="file.name"
                  :sub-title="file.type"
                  :img-src="file.src"
                  :img-alt="file.name"
                  img-top>
            <pre>{{ file | json }}</pre>
          </b-card>
        </b-col>
      </b-row>
    </b-container>
    <b-modal v-model="showModal"
             title="File Upload: Error">
      <p class="bg-danger text-white p-3 my-2"
         v-html="error"></p>
    </b-modal>
  </b-container>
  </div>
</template>

<script>
import { VueTransmit } from "vue-transmit";

export default {
  name: "App",
  components: {
    "vue-transmit": VueTransmit
  },
  data() {
    return {
      options: {
        acceptedFileTypes: ["image/*"],
        clickable: false,
        accept: this.accept,
        uploadMultiple: true,
        maxConcurrentUploads: 4,
        adapterOptions: {
          url: "/",
          timeout: 3000,
          errUploadError: xhr => xhr.response.message
        }
      },
      files: [],
      showModal: false,
      error: "",
      count: 0
    };
  },
  methods: {
    triggerBrowse() {
      this.$refs.uploader.triggerBrowseFiles();
    },
    onAddedFile(file) {
      console.log(
        this.$refs.uploader.inputEl.value,
        this.$refs.uploader.inputEl.files
      );
    },
    onUploadSuccess(file, res) {
      console.log(res);
      if (!this.options.uploadMultiple) {
        file.src = res.url[0];
        this.files.push(file);
      }
    },
    onUploadSuccessMulti(files, res) {
      console.log(...arguments);
      for (let i = 0; i < files.length; i++) {
        files[i].src = res.url[i];
        this.files.push(files[i]);
      }
    },
    onError(file, errorMsg) {
      this.error = errorMsg;
      this.showModal = true;
    },
    listen(event) {
      this.$refs.uploader.$on(event, (...args) => {
        console.log(event);
        for (let arg of args) {
          // console.log(`${typeof arg}: ${JSON.stringify(arg, undefined, 2)}`)
          console.log(arg);
        }
      });
    },
    accept(file, done) {
      this.count++;
      console.log(JSON.stringify(file, undefined, 2));
      done();
    }
  },
  filters: {
    json(value) {
      return JSON.stringify(value, null, 2);
    }
  },
  mounted() {
    [
      "drop",
      "drag-start",
      "drag-end",
      "drag-enter",
      "drag-over",
      "drag-leave",
      "accepted-file",
      "rejected-file",
      "accept-complete",
      "added-file",
      "added-files",
      "removed-file",
      "thumbnail",
      "error",
      "error-multiple",
      "processing",
      "processing-multiple",
      "upload-progress",
      "total-upload-progress",
      "sending",
      "sending-multiple",
      "success",
      "success-multiple",
      "canceled",
      "canceled-multiple",
      "complete",
      "complete-multiple",
      "reset",
      "max-files-exceeded",
      "max-files-reached",
      "queue-complete"
    ].forEach(this.listen);
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
