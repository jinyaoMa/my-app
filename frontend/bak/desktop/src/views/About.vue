<script setup lang="ts">
import { computed } from "vue";
import { useI18n } from "vue-i18n";
// import README from "../../../README.md";
// import README_ZH from "../../../README.zh.md";

const { locale } = useI18n();

const readme = computed(() => {
  // switch (locale.value) {
  //   case "zh":
  //     return README_ZH;
  // }
  // return README;
});
</script>

<template>
  <my-container class="about">
    <!-- <component :is="readme"></component> -->
  </my-container>
</template>

<style lang="scss" scoped>
.about {
  width: 100%;
  max-width: 1024px;
}
</style>
