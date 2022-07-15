https://blog.logrocket.com/build-table-component-scratch-vue3-bootstrap/

<template>
<div class="container text-center  mt-5 mb-5">
    <h1 class="mt-5 fw-bolder text-success "> Student's Database </h1>
    <div class="table-responsive my-5">
      <!-- The table component -->
      <Table :fields='fields' :studentData ="studentData"></Table>
    </div>
</div>
</template>

<script>
import Table from './components/Table.vue' // Importing the table component
import "bootstrap/dist/css/bootstrap.min.css"; //importing bootstrap 5

export default {
  name: 'App',

  components: {
    Table
  },

  setup(){
    //An array of values for the data
    const studentData = [
      {ID:"01", Name: "Abiola Esther", Course:"Computer Science", Gender:"Female", Age:"17"},
      {ID:"02", Name: "Robert V. Kratz", Course:"Philosophy", Gender:"Male", Age:'19'},
      {ID:"03", Name: "Kristen Anderson", Course:"Economics", Gender:"Female", Age:'20'},
      {ID:"04", Name: "Adam Simon", Course:"Food science", Gender:"Male", Age:'21'},
      {ID:"05", Name: "Daisy Katherine", Course:"Business studies", Gender:"Female", Age:'22'},  
    ]

    const fields = [ 'ID','Name','Course','Gender','Age' ]

    return {
      studentData, fields
    }
  },
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

// ============================ Table.vue ============================

<template>
  <div class="searchBar">
    <!-- Filter Search -->
      <div class="input-group mb-5">
        <input type="search" class="form-control" v-model='searchQuery'
			placeholder="Student's Name"
			aria-label="Recipient's username"
			aria-describedby="button-addon2">
      </div>
  </div>

  <table id="tableComponent" class="table table-bordered table-striped">
  <caption> A Responsive, Accessible Table Component</caption>
    <thead>
      <tr>
        <!-- loop through each value of the fields to get the table header -->
        <th  v-for="field in fields" :key='field' @click="sortTable(field)" > 
          {{field}} <i class="bi bi-sort-alpha-down" aria-label='Sort Icon'></i>
        </th>
      </tr>
    </thead>
    <tbody>
      <!-- Loop through the list get the each student data -->
      <tr v-for="item in filteredList" :key='item'>
        <td v-for="field in fields" :key='field'>{{item[field]}}</td>
      </tr>
    </tbody>
  </table> 
</template>

<script>
import {computed,ref} from "vue";
// Importing  the lodash library 
import { sortBy} from 'lodash';

export default {
  name: 'TableComponent',
  props:{
      studentData:{ type: Array },
      fields:{ type: Array }
  },
  
  setup(props) {
    let sort = ref(false);
    let updatedList =  ref([])
    let searchQuery = ref("");
    
    // a function to sort the table
    const sortTable = (col) => {
      sort.value = true
       // Use of _.sortBy() method
      updatedList.value = sortBy(props.studentData,col)
    }

    const sortedList = computed(() => {
      if (sort.value) {
         return updatedList.value
      }
      else {
         return props.studentData;
      }
    });

    // Filter Search
    const filteredList = computed(() => {
      return sortedList.value.filter((product) => {
        return (
          product.Name.toLowerCase().indexOf(searchQuery.value.toLowerCase()) != -1
        );
      });
    });   
      
    return {sortedList, sortTable,searchQuery,filteredList}
  }
}
</script>

<style scoped>
  table th:hover { background:#f2f2f2; }
</style>
