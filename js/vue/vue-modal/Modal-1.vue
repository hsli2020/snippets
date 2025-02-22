<template>
  <div v-if="mount">
    <portal :selector="appendTo">
      <transition
        name="vm-backdrop-transition"
        :enter-active-class="bgInClass"
        :leave-active-class="bgOutClass"
      >
        <div
          v-show="show"
          :data-vm-backdrop-id="id"
          class="vm-backdrop"
          :class="bgClass"
          :style="{ 'z-index': zIndex - 1 }"
        ></div>
      </transition>
      <transition
        name="vm-transition"
        :enter-active-class="inClass"
        :leave-active-class="outClass"
        @before-enter="beforeOpen"
        @enter="opening"
        @after-enter="afterOpen"
        @before-leave="beforeClose"
        @leave="closing"
        @after-leave="afterClose"
      >
        <div
          v-show="show"
          ref="vm-wrapper"
          :data-vm-wrapper-id="id"
          tabindex="-1"
          class="vm-wrapper"
          :class="wrapperClass"
          :style="{ 'z-index': zIndex, cursor: enableClose ? 'pointer' : 'default' }"
          role="dialog"
          :aria-label="title"
          aria-modal="true"
          @click="clickOutside($event)"
          @keydown="keydown($event)"
        >
          <div
            ref="vm"
            class="vm"
            :data-vm-id="id"
            :class="modalClass"
            :style="modalStyle"
          >
            <slot name="titlebar">
              <div class="vm-titlebar">
                <h3 class="vm-title">
                  {{ title }}
                </h3>
                <button
                  v-if="enableClose"
                  type="button"
                  class="vm-btn-close"
                  @click.prevent="close"
                ></button>
              </div>
            </slot>
            <slot name="content">
              <div class="vm-content">
                <slot></slot>
              </div>
            </slot>
          </div>
        </div>
      </transition>
    </portal>
  </div>
</template>

<script>
import { Portal } from '@linusborg/vue-simple-portal';

const TYPE_CSS = {
  type: [String, Object, Array],
  default: ''
};
const FOCUSABLE_ELEMENTS = 'a[href], area[href], input:not([disabled]), select:not([disabled]), textarea:not([disabled]), button:not([disabled]), [tabindex]:not([tabindex="-1"])';
let animatingZIndex = 0;

export default {
  name: 'VueModal',
  components: {
    Portal
  },
  model: {
    prop: 'basedOn',
    event: 'close'
  },
  props: {
    title: {
      type: String,
      default: ''
    },
    baseZindex: {
      type: Number,
      default: 1051
    },
    bgClass: TYPE_CSS,
    wrapperClass: TYPE_CSS,
    modalClass: TYPE_CSS,
    modalStyle: TYPE_CSS,
    inClass: Object.assign({}, TYPE_CSS, { default: 'vm-fadeIn' }),
    outClass: Object.assign({}, TYPE_CSS, { default: 'vm-fadeOut' }),
    bgInClass: Object.assign({}, TYPE_CSS, { default: 'vm-fadeIn' }),
    bgOutClass: Object.assign({}, TYPE_CSS, { default: 'vm-fadeOut' }),
    appendTo: {
      type: String,
      default: 'body'
    },
    live: {
      type: Boolean,
      default: false
    },
    enableClose: {
      type: Boolean,
      default: true
    },
    basedOn: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      zIndex: 0,
      id: null,
      show: false,
      mount: false,
      elToFocus: null
    };
  },
  created() {
    if (this.live) {
      this.mount = true;
    }
  },
  mounted() {
    this.id = 'vm-' + this._uid;
    this.$watch(
      'basedOn',
      function(newVal) {
        if (newVal) {
          this.mount = true;
          this.$nextTick(() => {
            this.show = true;
          });
        } else {
          this.show = false;
        }
      },
      {
        immediate: true
      }
    );
  },
  beforeDestroy() {
    this.elToFocus = null;
  },
  methods: {
    close() {
      if (this.enableClose === true) {
        this.$emit('close', false);
      }
    },
    clickOutside(e) {
      if (e.target === this.$refs['vm-wrapper']) {
        this.close();
      }
    },
    keydown(e) {
      if (e.which === 27) {
        this.close();
      }
      if (e.which === 9) {
        // Get only visible elements
        const all = [].slice.call(this.$refs['vm-wrapper'].querySelectorAll(FOCUSABLE_ELEMENTS)).filter(function(el) {
          return !!(el.offsetWidth || el.offsetHeight || el.getClientRects().length);
        });
        if (e.shiftKey) {
          if (e.target === all[0] || e.target === this.$refs['vm-wrapper']) {
            e.preventDefault();
            all[all.length - 1].focus();
          }
        } else {
          if (e.target === all[all.length - 1]) {
            e.preventDefault();
            all[0].focus();
          }
        }
      }
    },
    getAllVisibleWrappers() {
      return [].slice.call(document.querySelectorAll('[data-vm-wrapper-id]')).filter(w => w.display !== 'none');
    },
    getTopZindex() {
      return this.getAllVisibleWrappers().reduce((acc, curr) => {
        return parseInt(curr.style.zIndex) > acc ? parseInt(curr.style.zIndex) : acc;
      }, 0);
    },
    handleFocus(wrapper) {
      const autofocus = wrapper.querySelector('[autofocus]');
      if (autofocus) {
        autofocus.focus();
      } else {
        const focusable = wrapper.querySelectorAll(FOCUSABLE_ELEMENTS);
        focusable.length ? focusable[0].focus() : wrapper.focus();
      }
    },
    beforeOpen() {
      // console.log('beforeOpen');
      this.elToFocus = document.activeElement;
      const lastZindex = this.getTopZindex();
      if (animatingZIndex) {
        this.zIndex = animatingZIndex + 2;
      } else {
        this.zIndex = lastZindex === 0 ? this.baseZindex : lastZindex + 2;
      }
      animatingZIndex = this.zIndex;
      this.$emit('before-open');
    },
    opening() {
      // console.log('opening');
      this.$emit('opening');
    },
    afterOpen() {
      // console.log('afterOpen');
      this.handleFocus(this.$refs['vm-wrapper']);
      this.$emit('after-open');
    },
    beforeClose() {
      // console.log('beforeClose');
      this.$emit('before-close');
    },
    closing() {
      // console.log('closing');
      this.$emit('closing');
    },
    afterClose() {
      // console.log('afterClose');
      this.zIndex = 0;
      if (!this.live) {
        this.mount = false;
      }
      this.$nextTick(() => {
        window.requestAnimationFrame(() => {
          const lastZindex = this.getTopZindex();
          if (lastZindex > 0) {
            const all = this.getAllVisibleWrappers();
            for (let i = 0; i < all.length; i++) {
              const wrapper = all[i];
              if (parseInt(wrapper.style.zIndex) === lastZindex) {
                if (wrapper.contains(this.elToFocus)) {
                  this.elToFocus.focus();
                } else {
                  // console.log(wrapper);
                  this.handleFocus(wrapper);
                }
                break;
              }
            }
          } else {
            if (document.body.contains(this.elToFocus)) {
              this.elToFocus.focus();
            }
          }
          animatingZIndex = 0;
          this.$emit('after-close');
        });
      });
    }
  }
};
</script>
