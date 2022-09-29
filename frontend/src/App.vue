<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { useColorTheme } from "./store/color-theme";
import { EventsOn } from "../wailsjs/runtime";
import Sidebar from "./components/Sidebar.vue";

const { locale } = useI18n();
EventsOn("onDisplayLanguageChanged", (lang: string) => {
  console.log("onDisplayLanguageChanged", lang);
  locale.value = lang;
});
const { changeTheme } = useColorTheme();
EventsOn("onColorThemeChanged", (theme: string) => {
  console.log("onColorThemeChanged", theme);
  changeTheme(theme);
});
</script>

<template>
  <my-container class="app" height="100vh">
    <my-aside width="260px">
      <Sidebar></Sidebar>
    </my-aside>
    <my-main class="app-main">
      <router-view></router-view>
    </my-main>
  </my-container>
</template>

<style lang="scss" scoped>
.app {
  box-sizing: border-box;
  border-top: var(--my-border-width) solid var(--my-color-border-extra-light);
  &-main {
    padding: 2em;
    display: flex;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;
  }
}
</style>
