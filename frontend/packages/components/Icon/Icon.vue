<script setup lang="ts" name="MyIcon">
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

<template>
  <i
    :class="{
      [`my-icon-${name}`]: true,
      fixWidth,
    }"
  ></i>
</template>

<style lang="scss">
[class*="my-icon-"] {
  display: inline-block;
  &.fixWidth {
    width: 1.3em;
    text-align: center;
  }
}
</style>
