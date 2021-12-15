----------------------------------------
AppModal.vue
----------------------------------------

<template>
  <div class="modal" v-if="showModal">
    <div v-if="showModal" class="modal-content">
      <div class="modal-header">
        <slot name="header"></slot>
      </div>
      <hr>

      <div class="modal-body">
        <slot name="body"></slot>
      </div>
      <hr>

      <div class="modal-footer">
        <button @click="closeModal">Close</button>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'app-modal',
    props: {
        showModal: Boolean
    },
    methods: {
        closeModal() {
            this.$emit('clicked');
        }
    },
  };
</script>

<style>
    .modal {
        position: fixed;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        width: 600px;
        max-width: 100%;
        height: 400px;
        max-height: 100%;
        background: #FFFFFF;
        box-shadow: 2px 2px 20px 1px;
        overflow-x: auto;
        display: flex;
        flex-direction: column;
    }
</style>

----------------------------------------
App.vue
----------------------------------------

<template>
  <div id="app">
    <h1>Vue Modal Tutorial</h1>
    <button @click="openModal" v-if="!showModal">Open Modal</button>

    <app-modal v-if="showModal" :showModal=showModal @clicked="onChildClick">
      <div slot="header">
        <h3 class="modal-title">
          CodeMix
        </h3>
      </div>

      <div slot="body">
        <p>
          With CodeMix, you can join the modern web movement right from your Eclipse IDE!
        </p>
      </div>
    </app-modal>

  </div>
</template>

<script>
import AppModal from './components/AppModal';

export default {
  components: {
    AppModal
  },
  data() {
    return {
      showModal: false
    }
  },
  methods: {
    openModal() {
      this.showModal = true;
    },
    onChildClick () {
      this.showModal = false;
    }
  },
}
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
