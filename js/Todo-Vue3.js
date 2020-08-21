// App.vue ==================================================

<script setup>
import { ref } from 'vue'
import TodoList from './TodoList.vue'
  
const visibility = ref('all')
const themeColor = ref('#045975')
</script>

<template>
  <label><input type="radio"  v-model="visibility" value="all"> All</label>
  <label><input type="radio"  v-model="visibility" value="active"> Active</label>
  <label><input type="radio"  v-model="visibility" value="completed"> Completed</label>
  <label><input type="color" v-model="themeColor"> theme color</label>
  <TodoList :visibility="visibility" :themeColor="themeColor" />
</template>

<style scoped>
  label { margin-right: 1em; }
</style>

// TodoList.vue ==================================================

<script setup>
import { reactive, ref, computed } from 'vue'
import Todo from './Todo.vue'
import initialTodos from './initialTodos.js'

const props = defineProps(['visibility', 'themeColor'])

const todos = reactive(initialTodos)

let filterCalls = 0
const filteredTodos = computed(() => {
  filterCalls++
	return props.visibility === 'all' ? todos : todos.filter(todo => {
  	return props.visibility === 'active' ? !todo.done : todo.done
	})
})

const toggle = todo => todo.done = !todo.done

const addInput = ref()
const add = () => {
  todos.push({ text: addInput.value, done: false })
  addInput.value = ''
}
</script>

<template>
  <p>
    filter was called
    <span class="filter-call">{{ filterCalls }}</span> times
  </p>
  <ul>
    <Todo v-for="todo of filteredTodos"
          :key="todo.text"
          :todo="todo"
          @change="toggle" />
  </ul>
  <form @submit.prevent="add">
    <input placeholder="add todo" ref="addInput">
    <button type="submit">Add</button>
  </form>
</template>

<style scoped>
  .filter-call {
    background-color: #f66;
    color: #fff;
    padding: 0.2em 0.5em;
    border-radius: 4px;
  }
  button {
    color: #fff;
    background-color: v-bind('props.themeColor')
  }
</style>

// Todo.vue ==================================================

<script setup>
import { ref, onBeforeUpdate } from 'vue'

defineProps(['todo'])
defineEmits(['change'])

let updates = 0
onBeforeUpdate(() => updates++)
</script>

<template>
  <li>
  	<label>
      <input type="checkbox"
             :checked="todo.done"
             @change="$emit('change', todo)">
      {{ todo.text }}
    </label>
	  <span>Updated {{ updates }} times</span>
  </li>
</template>

<style scoped>
  span {
    background-color: yellow;
    border-radius: 8px;
    display: inline-block;
    font-size: 0.8em;
    padding: 0.2em 0.5em;
    margin-left: 1em;
  }
</style>

// initialTodos.js ==================================================

export default [
  { text: 'one', done: false },
  { text: 'two', done: false },
  { text: 'three', done: false }
]
