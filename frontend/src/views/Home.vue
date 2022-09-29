<script setup lang="ts">
import { useI18n } from "vue-i18n";
import { useColorTheme } from "../store/color-theme";
import { EventsOn } from "../../wailsjs/runtime";

const { t, locale } = useI18n();
const { theme, changeTheme } = useColorTheme();
</script>

<template>
  <my-container class="home"> </my-container>
</template>

<style lang="scss" scoped>
.home {
}
</style>
