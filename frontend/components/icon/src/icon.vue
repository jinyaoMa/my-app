<template>
  <i class="my-icon" :class="{
    [`my-icon-${props.name}`]: true,
    [`my-icon__fw`]: props.fixWidth,
  }"></i>
</template>

<script setup lang="ts">
import { withDefaults } from "vue";

defineOptions({
  name: 'MyIcon',
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
@use "@jinyaoma/my-app-icons/dist/my-icon.scss";

.my-icon {
  display: inline-block;

  &.my-icon__fw { // fix width
    width: 1.3em;
    text-align: center;
  }
}
</style>
