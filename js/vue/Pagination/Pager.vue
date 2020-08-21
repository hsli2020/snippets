<template>
  <ul class="x-pager">
    <!-- 首页 -->
    <li 
      class="number" 
      :class="{ active: current == 1 }" 
      @click="setPage(1)"
    >1</li>

    <!-- 左更多按钮 -->
    <li 
      class="more left"
      v-if="totalPage > centerSize + 2 && current >= centerSize"
      @click="setPage(current - jumpSize)"
    ></li>

    <!-- 中间页码组 -->
    <li 
      class="number"
      :class="{ active: current === page }"
      v-for="(page, index) in centerPages"
      :key="index"
      @click="setPage(page)"
    >{{ page }}</li>

    <!-- 右更多按钮 -->
    <li
      class="more right"
      v-if="totalPage > centerSize + 2 && current <= totalPage - centerSize + 1"
      @click="setPage(current + jumpSize)"
    ></li>
    
    <!-- 尾页 -->
    <li 
      class="number" 
      :class="{ active: current == totalPage }"
      v-if="totalPage !== 1" 
      @click="setPage(totalPage)"
    >{{ totalPage }}</li>
  </ul>
</template>

<script>
import Vue from 'vue';
import { generatePages } from '@/util';

export default {
  name: 'Pager',
  props: {
    totalPage: Number,
    defaultCurrent: Number,
  },
  data() {
    return {
      centerSize: 5,
      jumpSize: 5,
      current: this.defaultCurrent,
      pages: generatePages(this.totalPage),
    }
  },
  watch: {
    defaultCurrent: {
      handler(newValue, oldValue) {
        this.current = newValue;
      }
    }
  },
  computed: {
    centerPages: function() {
      let centerPage = this.current;
      if (this.current > this.totalPage - 3) {
        centerPage = this.totalPage - 3;
      }
      if (this.current < 4) {
        centerPage = 4;
      }
      if (this.totalPage <= this.centerSize + 2) {
        const centerArr = [];
        for (let i = 2; i < this.totalPage; i++) {
          centerArr.push(i);
        }
        return centerArr;
      } else {
        const centerArr = [];
        for (let i = centerPage - 2; i <= centerPage + 2; i++) {
          centerArr.push(i);
        }
        return centerArr;
      }
    }
  },
  methods: {
    setPage(page) {
      this.current = page;
      this.$emit('change', this.current);
    }
  },
};
</script>

<style scoped lang="scss">
@import '@/assets/common.scss';

.x-pager {
  list-style: none;
  padding: 0;
  li {
    @include page-button;

    &.active {
      background-color: $blue;
      color: #fff;
      cursor: default;
    }

    &:not(.active):hover {
      color: $blue;
    }

    &.more {
      &::before {
        content: '...';
      }
      &.left:hover::before {
        content: '<<';
      }
      &.right:hover::before {
        content: '>>';
      }
    }
  }
}
</style>
