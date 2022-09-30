<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { useColorTheme } from "../store/color-theme";
import { EventsOn } from "../../wailsjs/runtime";
import README from "../../../README.md";
import README_ZH from "../../../README.zh.md";
import { computed } from "vue";

const { t, locale } = useI18n();
const { theme, changeTheme } = useColorTheme();

const readme = computed(() => {
  switch (locale.value) {
    case "zh":
      return README_ZH;
  }
  return README;
});
</script>

<template>
  <my-container class="about">
    <component :is="readme"></component>
  </my-container>
</template>

<style lang="scss" scoped>
.about {
  width: 100%;
  max-width: 1024px;
}
</style>
