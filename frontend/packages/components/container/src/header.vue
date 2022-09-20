<template>
  <header class="my-header" :style="style">
    <slot></slot>
  </header>
</template>

<script setup lang="ts" name="MyHeader">
import type { StyleValue } from "vue";
import { computed, withDefaults } from "vue";

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
