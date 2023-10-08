<template>
  <button class="my-button"></button>
</template>

<script setup lang="ts">
import { withDefaults } from "vue";

defineOptions({
  name: 'MyButton',
  inheritAttrs: false,
})

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
.my-button {
  display: inline-block;

  &.fixWidth {
    width: 1.3em;
    text-align: center;
  }
}
</style>
