<template>
  <header class="my-header" :style="style">
    <slot></slot>
  </header>
</template>

<script setup lang="ts" name="MyHeader">
import { computed, StyleValue, withDefaults } from "vue";

const props = withDefaults(
  defineProps<{
    height?: string;
  }>(),
  {
    height: "80px",
  }
);

const style = computed<StyleValue>(() => ({
  height: props.height,
}));
</script>

<style lang="scss">
.my-header {
  overflow: visible;
}
</style>
