<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { useColorTheme } from "../store/color-theme";
import { EventsOn } from "../../wailsjs/runtime";
import README from "../../../README.md";

const { t, locale } = useI18n();
const { theme, changeTheme } = useColorTheme();
</script>

<template>
  <my-container class="about">
    <README></README>
  </my-container>
</template>

<style lang="scss" scoped>
.about {
  width: 100%;
  max-width: 1024px;
}
</style>
