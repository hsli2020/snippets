<template>
  <ul class="m-list">
    <li v-for="list in lists" :key="list.id">
      {{ list.name }}
    </li>
  </ul>
</template>

<script>
export default {
  name: 'List',
  props: {
    dataSource: {
      type: Array
    }
  },
  data() {
    return {
      lists: this.dataSource
    }
  },
  watch: {
    // 对dataSource进行监听，如果发生变化
    dataSource: {
      handler(newValue, oldValue) {
        this.lists = newValue;
      }
    }
  }
};
</script>
