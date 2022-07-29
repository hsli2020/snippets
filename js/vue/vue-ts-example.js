// Play with it here: http://vite.vue.new

// ========== index.html

<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <link rel="icon" type="image/svg+xml" href="/vite.svg" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Vite + Vue + TS</title>
  </head>
  <body>
    <div id="app"></div>
    <script type="module" src="/src/main.ts"></script>
  </body>
</html>

// ========== main.ts

import { createApp } from 'vue'
import './style.css'
import App from './App.vue'

createApp(App).mount('#app')

// ========== App.vue

<script setup lang="ts">
import HelloWorld from "./components/HelloWorld.vue";
</script>

<template>
  <div>
    <a href="https://vitejs.dev"></a>
    <a href="https://vuejs.org/"></a>
  </div>
  <HelloWorld msg="Vite + Vue" />
</template>

<style scoped>
.logo { height: 6em; padding: 1.5em; will-change: filter; transition: filter 300ms; }
.logo:hover { filter: drop-shadow(0 0 2em #646cffaa); }
.logo.vue:hover { filter: drop-shadow(0 0 2em #42b883aa); }
</style>

// ========== components/HelloWorld.vue

<script setup lang="ts">
import { ref } from "vue";

defineProps<{ msg: string }>();

const count = ref(0);
</script>

<template>
  <h1>{{ msg }}</h1>

  <button type="button" @click="count++">count is {{ count }}</button>

  Install <a href="https://github.com/johnsoncodehk/volar">Volar</a> in your IDE for a better DX
</template>

<style scoped></style>
