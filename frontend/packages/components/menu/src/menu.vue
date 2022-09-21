<template>
  <nav class="my-menu" :style="style">
    <slot></slot>
  </nav>
</template>

<script setup lang="ts" name="MyMenu">
import type { StyleValue } from "vue";
import { computed, withDefaults } from "vue";

const props = withDefaults(
  defineProps<{
    direction?: "vertical" | "horizontal";
  }>(),
  {
    direction: "vertical",
  }
);

const style = computed<StyleValue>(() => ({
  flexDirection: props.direction === "horizontal" ? "row" : "column",
}));
</script>

<style lang="scss">
.my-menu {
  display: flex;
}
</style>
