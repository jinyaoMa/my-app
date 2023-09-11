<template></template>

<script setup lang="ts" name="MyIcon">
import "@jinyaoma/my-app-icons/dist/my-icon.scss";
import { withDefaults } from "vue";

const props = withDefaults(
  defineProps<{
    name?: string;
    fixWidth?: boolean;
  }>(),
  {
    name: "jinyao-ma",
    fixWidth: false,
  }
);
</script>

<style lang="scss">
[class*="my-icon-"] {
  display: inline-block;
  &.fixWidth {
    width: 1.3em;
    text-align: center;
  }
}
</style>
