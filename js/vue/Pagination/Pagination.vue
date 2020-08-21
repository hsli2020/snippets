<template>
  <div class="m-pagination">
    <Button class="btn-prev" @click="setPage(current - 1)">&lt;</Button>
    <Pager :total-page="totalPage" :default-current="current" @change="onChange"></Pager>
    <Button class="btn-next" @click="setPage(current + 1)">></Button>
  </div>
</template>

<script>
// @ is an alias to /src
import Button from './Button.vue';
import Pager from './Pager.vue';

export default {
  name: 'Pagination',
  components: {
    Button,
    Pager,
  },
  // 接口定义 props
  props: {
    defaultCurrent: {
      type: Number,
      // default: 2
    },
    defaultPageSize: {
      type: Number,
      // default: 3
    },
    total: {
      type: Number,
      // default: 24
    },
  },
  data() {
    return {
      current: this.defaultCurrent,
    }
  },
  computed: {
    totalPage: function () {
      return Math.ceil(this.total / this.defaultPageSize);
    },
  },
  methods: {
    setPage(page) {
      if (page < 1) return;
      if (page > this.totalPage) return;
      this.current = page;
      this.$emit('change', this.current);
    },
    onChange(current) {
      this.current = current;
      this.$emit('change', this.current);
    }
  }
};
</script>

<style scoped lang="scss">
@import '@/assets/common.scss';

.m-pagination {
  display: flex;
  align-items: center;
  .btn-prev, .btn-next {
    @include page-button;
    border: none;
  }
}
</style>
