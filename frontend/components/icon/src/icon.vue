<template>
  <i class="my-icon" :class="{
    [`my-icon-${props.name}`]: true,
    fixWidth: props.fixWidth,
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
@import "@jinyaoma/my-app-icons/dist/my-icon.scss";

.my-icon {
  &.fixWidth {
    width: 1.3em;
    text-align: center;
  }
}
</style>
