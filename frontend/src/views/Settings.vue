<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { useColorTheme } from "../store/color-theme";
import { EventsOn } from "../../wailsjs/runtime";

const { t, locale } = useI18n();
const colorTheme = useColorTheme();
</script>

<template>
  <my-container class="settings">
    <my-main>
      <h2>{{ t("menu.main.settings") }}</h2>
    </my-main>
  </my-container>
</template>

<style lang="scss" scoped>
.settings {
  max-width: 800px;
}
</style>
