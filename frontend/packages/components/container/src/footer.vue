<template>
  <footer class="my-footer" :style="style">
    <slot></slot>
  </footer>
</template>

<script setup lang="ts" name="MyFooter">
import type { StyleValue } from "vue";
import { computed, withDefaults } from "vue";

const props = withDefaults(
  defineProps<{
    height?: string;
  }>(),
  {
    height: "auto",
  }
);

const style = computed<StyleValue>(() => ({
  height: props.height,
}));
</script>
