// src/components/Modal.vue
<script>
  export default {
    name: 'Modal',
    methods: {
      close() {
        this.$emit('close');
      },
    },
  };
</script>

<template>
  <transition name="modal-fade">
    <div class="modal-backdrop">
      <div class="modal" role="dialog"
        aria-labelledby="modalTitle" aria-describedby="modalDescription"
      >
        <header class="modal-header" id="modalTitle">
          <slot name="header">
            This is the default tile!
          </slot>
          <button type="button" class="btn-close" @click="close" aria-label="Close modal">
            x
          </button>
        </header>

        <section class="modal-body" id="modalDescription">
          <slot name="body">
            This is the default body!
          </slot>
        </section>

        <footer class="modal-footer">
          <slot name="footer">
            This is the default footer!
          </slot>

          <button type="button" class="btn-green" @click="close" aria-label="Close modal">
            Close me!
          </button>
        </footer>
      </div>
    </div>
  </transition>
</template>

// src/App.vue

<template>
  <div id="app">
    <button type="button" class="btn" @click="showModal">Open Modal!</button>

    <Modal v-show="isModalVisible" @close="closeModal">
      <template v-slot:header>
        This is a new modal header.
      </template>

      <template v-slot:body>
        This is a new modal body.
      </template>

      <template v-slot:footer>
        This is a new modal footer.
      </template>
    </Modal>
  </div>
</template>

<script>
  import modal from './components/Modal.vue';

  export default {
    name: 'App',
    components: {
      Modal,
    },
    data() {
      return {
        isModalVisible: false,
      };
    },
    methods: {
      showModal() {
        this.isModalVisible = true;
      },
      closeModal() {
        this.isModalVisible = false;
      }
    }
  };
</script>
