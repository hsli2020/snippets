<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App by kagol"/>
    <List :data-source="dataList" />
    <Pagination :default-current="defaultCurrent" :default-page-size="defaultPageSize" :total="total" @change="onChange" />
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue';
import Pagination from '@/components/pagination/Pagination.vue';
import List from './List.vue';
import { lists } from '@/db';
import { chunk } from '@/util';

export default {
  name: 'home',
  components: {
    HelloWorld,
    Pagination,
    List,
  },
  data() {
    return {
      defaultCurrent: 1,
      defaultPageSize: 4,
      total: lists.length,
      dataList: [],
    }
  },
  created() {
    this.setList(this.defaultCurrent, this.defaultPageSize);
  },
  methods: {
    onChange(current) {
      this.setList(current, this.defaultPageSize);
    },
    setList: function(current, pageSize) {
      this.dataList = chunk(lists, pageSize)[current - 1];
    }
  }
};
</script>
