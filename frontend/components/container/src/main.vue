<template>
  <main class="my-main" :style="style">
    <slot></slot>
  </main>
</template>

<script setup lang="ts" name="MyMain">
import { computed, StyleValue } from "vue";

const style = computed<StyleValue>(() => ({
  flexGrow: 1,
}));
</script>

<style lang="scss">
.my-main {
  overflow: auto;
}
</style>
