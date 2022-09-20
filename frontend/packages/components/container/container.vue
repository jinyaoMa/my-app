<script setup lang="ts" name="MyContainer">
defineProps({
  direction: {
    type: String,
    default(): string {
      return "vertical";
    },
  },
});
</script>

<template>
  <div class="my-container">
    <slot></slot>
  </div>
</template>

<style lang="scss">
@import "../theme-default/mixin.scss";

.my-container {
  @include font-family();
}
</style>
