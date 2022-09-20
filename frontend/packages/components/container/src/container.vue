<template>
  <section class="my-container" :style="style">
    <slot></slot>
  </section>
</template>

<script setup lang="ts" name="MyContainer">
import type { Component, StyleValue, VNode } from "vue";
import { computed, useSlots, withDefaults } from "vue";

const props = withDefaults(
  defineProps<{
    direction?: "vertical" | "horizontal";
  }>(),
  {
    direction: "vertical",
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
}));
</script>

<style lang="scss">
@import "../../theme-default/mixin.scss";

.my-container {
  @include font-family();
  display: flex;
}
</style>
