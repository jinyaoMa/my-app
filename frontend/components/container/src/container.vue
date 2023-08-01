<template>
  <section class="my-container" :style="style">
    <slot></slot>
    <div class="my-container__loading" :class="{ loading }"></div>
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
    loading?: boolean;
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
  position: relative;

  &__loading {
    position: absolute;
    top: 0;
    left: 0;
    height: 100%;
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    pointer-events: none;
    background-color: transparent;
    transition: background-color 0.2s;

    &.loading {
      pointer-events: all;
      background-color: var(--my-color-background-transparent);

      &::before {
        content: "";
        display: block;
        width: 10vmin;
        height: 10vmin;
        border-top: 2vmin solid var(--my-color-text-primary);
        border-left: 2vmin solid var(--my-color-text-primary);
        border-right: 2vmin solid var(--my-color-text-primary);
        border-bottom: 2vmin solid transparent;
        border-radius: 50%;
        animation: loading 1.2s linear infinite;
      }
    }

    @keyframes loading {
      0% {
        border-top-color: var(--my-color-primary);
        border-left-color: var(--my-color-primary);
        border-right-color: var(--my-color-primary);
        transform: rotate(0deg);
      }
      50% {
        border-top-color: var(--my-color-primary-2);
        border-left-color: var(--my-color-primary-2);
        border-right-color: var(--my-color-primary-2);
      }
      100% {
        border-top-color: var(--my-color-primary);
        border-left-color: var(--my-color-primary);
        border-right-color: var(--my-color-primary);
        transform: rotate(360deg);
      }
    }
  }
}
</style>
