<script setup lang="ts">
import { computed } from "vue";
import { useI18n } from "vue-i18n";
import { useLoading } from "./store/loading";
import { useColorTheme } from "./store/color-theme";
import { EventsOn } from "../wailsjs/runtime";
import { GetOptions } from "../wailsjs/go/service/service";
import Sidebar from "./components/Sidebar.vue";

const storeLoading = useLoading();
const loading = computed(() => storeLoading.loading);

const { locale } = useI18n();
const { changeTheme } = useColorTheme();
GetOptions().then((config: config.Config) => {
  locale.value = config.DisplayLanguage;
  changeTheme(config.ColorTheme);
  storeLoading.endLoading();
});
EventsOn("onDisplayLanguageChanged", (lang: string) => {
  locale.value = lang;
  storeLoading.endLoading();
});
EventsOn("onColorThemeChanged", (theme: string) => {
  changeTheme(theme);
  storeLoading.endLoading();
});
</script>

<template>
  <my-container class="app" height="100vh" :loading="loading">
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
