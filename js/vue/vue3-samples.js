https://github.com/iosamuel/vue3-samples

//======== main.js

import { createApp } from 'vue'
import App from './App.vue'

createApp(App).mount('#app')

//======== App.vue

<template>
  <Counter />
  <ToDo />
</template>

<script>
import Counter from "./components/Counter.vue";
import ToDo from "./components/ToDo.vue";

export default {
  components: {
    Counter,
    ToDo,
  },
};
</script>

//======== components/Counter.vue

<template>
  <button @click="increment">
    Count is: {{ state.count }}, double is: {{ state.double }}
  </button>
</template>

<script>
import { reactive, computed } from "vue";

export default {
  setup() {
    const state = reactive({
      count: 0,
      double: computed(() => state.count * 2),
    });
    function increment() {
      state.count++;
    }
    return {
      state,
      increment,
    };
  },
};
</script>

<style scoped>
button { font-size: 24px; padding: 14px; }
</style>

//======== components/ToDo.vue

<template>
  <div class="nav">
    <button @click="addToList()">Add new item</button>
    <div class="buttons">
      <label for="split-color">Split Color</label>
      <input id="split-color" type="checkbox" v-model="splitColor" />
    </div>
    <div class="buttons">
      <label for="fetch-list">Fetch List</label>
      <input
        id="fetch-list"
        type="checkbox"
        :disabled="shouldFetch"
        v-model="shouldFetch"
      />
    </div>
    <div class="mouse-position">X: {{ x }} | Y: {{ y }}</div>
  </div>
  <ul v-for="(todos, key) in todosList" :key="key">
    <li v-for="todo in todos" :key="todo.id" :class="`list-${key}`">
      <span>{{ todo.title }}</span>
      <button @click="deleteFromList(todo.id)">X</button>
    </li>
  </ul>
</template>

<script>
import { useTodoList } from "../composable/useTodoList";
import { useMousePosition } from "../composable/useMousePosition";
import { computed, ref } from "vue";

export default {
  setup() {
    const { todos, shouldFetch, deleteFromList, addToList } = useTodoList();
    const { x, y } = useMousePosition();
    const splitColor = ref(false);

    const todosList = computed(() => {
      if (splitColor.value) {
        const halfList = todos.value.length / 2;
        return [todos.value.slice(0, halfList), todos.value.slice(halfList)];
      }
      return [todos.value];
    });

    return {
      // useTodoList
      todosList,
      shouldFetch,
      deleteFromList,
      addToList,
      // ToDo component
      splitColor,
      // useMousePosition
      x,
      y,
    };
  },
};
</script>

<style scoped>
* { box-sizing: border-box; }
</style>

//======== composable/useMousePosition.js

import { reactive, onMounted, onUnmounted, toRefs } from "vue";

export function useMousePosition() {
  const mousePosition = reactive({
    x: 0,
    y: 0
  });

  function update(evt) {
    mousePosition.x = evt.clientX;
    mousePosition.y = evt.clientY;
  }

  onMounted(() => {
    document.addEventListener("mousemove", update);
  });

  onUnmounted(() => {
    document.removeEventListener("mousemove", update);
  });

  return {
    ...toRefs(mousePosition)
  };
}

//======== composable/useTodoList.js 

import { ref, watchEffect } from "vue";

export function useTodoList() {
  let todos = ref([
    { id: 1, title: "Hello" },
    { id: 2, title: "Composition" },
    { id: 3, title: "API" }
  ]);

  let shouldFetch = ref(false);

  watchEffect(() => {
    if (shouldFetch.value) {
      fetch("http://jsonplaceholder.typicode.com/todos")
        .then(response => response.json())
        .then(json => {
          for (let { id, title } of json) {
            todos.value.push({
              id: `api-${id}`,
              title
            });
          }
        });
    }
  });

  function deleteFromList(id) {
    todos.value = todos.value.filter(todo => {
      return todo.id !== id;
    });
  }

  function addToList() {
    todos.value.push({
      id: todos.value.length,
      title: `New item! ${todos.value.length}`
    });
  }

  return {
    todos,
    shouldFetch,
    deleteFromList,
    addToList
  };
}
