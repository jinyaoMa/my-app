<template>
  <aside class="my-aside" :style="style">
    <slot></slot>
  </aside>
</template>

<script setup lang="ts" name="MyAside">
import { computed, StyleValue, withDefaults } from "vue";

const props = withDefaults(
  defineProps<{
    width?: string;
  }>(),
  {
    width: "280px",
  }
);

const style = computed<StyleValue>(() => ({
  width: props.width,
  minWidth: props.width,
}));
</script>

<style lang="scss">
.my-aside {
  overflow: auto;
}
</style>
