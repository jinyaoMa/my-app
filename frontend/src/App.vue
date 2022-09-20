<script setup lang="ts">
import { ref } from "vue";
import { useI18n } from "vue-i18n";
import { useColorTheme } from "./store/color-theme";
import {
  EventsOn,
  Hide,
  WindowMinimise,
  WindowToggleMaximise,
} from "../wailsjs/runtime/runtime";

const count = ref(0);
const msg = ref("Vite + Vue");

const { t, locale } = useI18n();
EventsOn("onDisplayLanguageChanged", (lang: string) => {
  console.log("onDisplayLanguageChanged", lang);
  locale.value = lang;
});
const colorTheme = useColorTheme();
EventsOn("onColorThemeChanged", (theme: string) => {
  console.log("onColorThemeChanged", theme);
  colorTheme.theme = theme;
});
</script>

<template>
  <my-container
    :style="{
      height: '100vh',
    }"
  >
    <my-aside>
      <my-container
        :style="{
          gap: '10px',
        }"
      >
        <a href="#" target="_blank">
          <img src="/icon.svg" class="logo avatar" />
        </a>
        <div>
          <my-icon fix-width></my-icon>
        </div>
        <button @click="Hide">Hide</button>
        <button @click="WindowMinimise">WindowMinimise</button>
        <button @click="WindowToggleMaximise">WindowToggleMaximise</button>
        <div>{{ t("lang") }}</div>
        <div>{{ colorTheme.theme }}</div>
      </my-container>
    </my-aside>
    <my-main>
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src="/vite.svg" class="logo" alt="Vite logo" />
        </a>
        <a href="https://vuejs.org/" target="_blank">
          <img src="./assets/vue.svg" class="logo vue" alt="Vue logo" />
        </a>
      </div>
      <h1>{{ msg }}</h1>
      <div class="card">
        <button type="button" @click="count++">count is {{ count }}</button>
        <p>
          Edit
          <code>components/HelloWorld.vue</code> to test HMR
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
        <a href="https://github.com/johnsoncodehk/volar" target="_blank"
          >Volar</a
        >
        in your IDE for a better DX
      </p>
      <p class="read-the-docs">Click on the Vite and Vue logos to learn more</p>
    </my-main>
  </my-container>
</template>

<style lang="scss">
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
</style>
