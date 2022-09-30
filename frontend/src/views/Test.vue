<script setup lang="ts">
import { ref } from "vue";
import {
  Hide,
  WindowMinimise,
  WindowToggleMaximise,
} from "../../wailsjs/runtime";

const count = ref(0);
const msg = ref("Wails + Vite + Vue");
</script>

<template>
  <div class="test">
    <div>
      <a href="https://wails.io/" target="_blank">
        <img
          src="../assets/wails-logo-textless.svg"
          class="logo wails"
          alt="Wails logo"
        />
      </a>
      <a href="https://vitejs.dev" target="_blank">
        <img src="../assets/vite.svg" class="logo" alt="Vite logo" />
      </a>
      <a href="https://vuejs.org/" target="_blank">
        <img src="../assets/vue.svg" class="logo vue" alt="Vue logo" />
      </a>
      <a href="https://github.com/jinyaoMa" target="_blank">
        <img src="/icon.svg" class="logo avatar" alt="Author logo" />
      </a>
    </div>
    <h1>{{ msg }}</h1>
    <div class="card">
      <button type="button" @click="count++">count is {{ count }}</button>&nbsp;
      <button @click="Hide">WindowHide</button>&nbsp;
      <button @click="WindowMinimise">WindowMinimise</button>&nbsp;
      <button @click="WindowToggleMaximise">WindowToggleMaximise</button>
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
    <p class="read-the-docs">
      Click on the Wails, Vite, Vue and My logos to learn more
    </p>
  </div>
</template>

<style lang="scss" scoped>
.test {
  width: 100%;
  max-width: 1024px;
  text-align: center;
}

a {
  font-weight: 500;
  color: #646cff;
  text-decoration: inherit;
}
a:hover {
  color: #535bf2;
  @media (prefers-color-scheme: light) {
    color: #747bff;
  }
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
  background-color: #1a1a1a;
  @media (prefers-color-scheme: light) {
    background-color: #f9f9f9;
  }
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
.logo.wails:hover {
  filter: drop-shadow(0 0 2em #e55555aa);
}
.logo.avatar:hover {
  filter: drop-shadow(0 0 2em #46adeaaa);
}
.read-the-docs {
  color: #888;
}

.card {
  padding: 2em;
}

:root.my-theme-light {
  a:hover {
    color: #747bff;
  }
  button {
    background-color: #f9f9f9;
  }
}
</style>
