<template>
  <i
    :class="{
      [`my-icon-${name}`]: true,
      fixWidth,
    }"
  ></i>
</template>

<script setup lang="ts" name="MyIcon">
import "@jinyaoma/my-app-icons/dist/my-icon.scss";
defineProps({
  name: {
    type: String,
    default(): string {
      return "jinyao-ma";
    },
  },
  fixWidth: {
    type: Boolean,
    default(): boolean {
      return false;
    },
  },
});
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
