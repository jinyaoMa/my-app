<template>
  <main class="my-main" :style="style">
    <slot></slot>
  </main>
</template>

<script lang="ts" setup name="MyMain">
import type { StyleValue } from "vue";
import { computed } from "vue";

const style = computed<StyleValue>(() => ({
  flexGrow: 1,
}));
</script>

<style lang="scss">
.my-main {
  overflow: auto;
}
</style>
