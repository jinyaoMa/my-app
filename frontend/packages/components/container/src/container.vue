<template>
  <section class="my-container" :style="style">
    <slot></slot>
  </section>
</template>

<script setup lang="ts" name="MyContainer">
import {
  Component,
  StyleValue,
  VNode,
  computed,
  useSlots,
  withDefaults,
} from "vue";
import { Direction } from "../../types";

const props = withDefaults(
  defineProps<{
    direction?: Direction;
    height?: string;
  }>(),
  {
    direction: "vertical",
    height: "auto",
  }
);
const slots = useSlots();

const isHorizontal = computed(() => {
  if (props.direction === "horizontal") {
    return true;
  }
  if (slots && slots.default) {
    const vNodes: VNode[] = slots.default();
    return vNodes.some((vNode) => {
      const tag = (vNode.type as Component).name;
      return tag === "MyAside";
    });
  }
  return false;
});

const style = computed<StyleValue>(() => ({
  flexDirection: isHorizontal.value ? "row" : "column",
  height: props.height,
}));
</script>

<style lang="scss">
@import "../../theme-default/mixin.scss";

.my-container {
  @include font-family();
  display: flex;
}
</style>
