<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { useColorTheme } from "./store/color-theme";
import {
  EventsOn,
  Hide,
  WindowMinimise,
  WindowToggleMaximise,
} from "../wailsjs/runtime/runtime";

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
  <div>
    <a href="#" target="_blank">
      <img src="/icon.svg" class="logo avatar" />
    </a>
  </div>
  <div>
    <button @click="Hide">Hide</button>
    <button @click="WindowMinimise">WindowMinimise</button>
    <button @click="WindowToggleMaximise">WindowToggleMaximise</button>
  </div>
  <div>{{ t("lang") }}</div>
  <div>{{ colorTheme.theme }}</div>
  <div>
    <a href="https://vitejs.dev" target="_blank">
      <img src="/vite.svg" class="logo" alt="Vite logo" />
    </a>
    <a href="https://vuejs.org/" target="_blank">
      <img src="./assets/vue.svg" class="logo vue" alt="Vue logo" />
    </a>
  </div>
  <my-hello-world msg="Vite + Vue" />
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
</style>
