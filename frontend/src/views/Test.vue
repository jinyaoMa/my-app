<script setup lang="ts">
import { ref } from "vue";
import {
  Hide,
  WindowMinimise,
  WindowToggleMaximise,
} from "../../wailsjs/runtime";

const count = ref(0);
const msg = ref("Vite + Vue");
</script>

<template>
  <div class="test">
    <div>
      <a href="https://vitejs.dev" target="_blank">
        <img src="/vite.svg" class="logo" alt="Vite logo" />
      </a>
      <a href="https://vuejs.org/" target="_blank">
        <img src="../assets/vue.svg" class="logo vue" alt="Vue logo" />
      </a>
      <a href="#" target="_blank">
        <img src="/icon.svg" class="logo avatar" />
      </a>
    </div>
    <h1>{{ msg }}</h1>
    <div class="card">
      <button @click="Hide">Hide</button>&nbsp;
      <button @click="WindowMinimise">WindowMinimise</button>&nbsp;
      <button @click="WindowToggleMaximise">WindowToggleMaximise</button>&nbsp;
      <button type="button" @click="count++">count is {{ count }}</button>
      <p>
        Edit
        <code>views/Test.vue</code> to test HMR
      </p>
    </div>
    <p>
      Check out
      <a href="https://vuejs.org/guide/quick-start.html#local" target="_blank"
        >create-vue</a
      >, the official Vue + Vite starter
    </p>
    <p>
      Install
      <a href="https://github.com/johnsoncodehk/volar" target="_blank">Volar</a>
      in your IDE for a better DX
    </p>
    <p class="read-the-docs">Click on the Vite and Vue logos to learn more</p>
  </div>
</template>

<style lang="scss" scoped>
.test {
  text-align: center;
}

a {
  font-weight: 500;
  color: #646cff;
  text-decoration: inherit;
}
a:hover {
  color: #535bf2;
}

h1 {
  font-size: 3.2em;
  line-height: 1.1;
}

button {
  border-radius: 8px;
  border: 1px solid transparent;
  padding: 0.6em 1.2em;
  font-size: 1em;
  font-weight: 500;
  font-family: inherit;
  background-color: #f3f2f1;
  cursor: pointer;
  transition: border-color 0.25s;
}
button:hover {
  border-color: #646cff;
}
button:focus,
button:focus-visible {
  outline: 4px auto -webkit-focus-ring-color;
}

.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
.logo.avatar:hover {
  filter: drop-shadow(0 0 2em #99d2f4);
}
.read-the-docs {
  color: #888;
}

.card {
  padding: 2em;
}
</style>
