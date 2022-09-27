<template>
  <ul
    class="my-menu"
    :class="{
      [`my-menu--${props.type}`]: true,
    }"
    :style="style"
  >
    <slot></slot>
  </ul>
</template>

<script setup lang="ts" name="MyMenu">
import {
  computed,
  withDefaults,
  provide,
  ref,
  StyleValue,
  onUpdated,
} from "vue";
import { Type, Size, SizeRef, Direction } from "../../types";

const props = withDefaults(
  defineProps<{
    direction?: Direction;
    type?: Type;
    size?: Size;
    height?: string;
  }>(),
  {
    direction: "vertical",
    type: "default",
    size: "default",
    height: "auto",
  }
);

const menuSize = ref<Size>(props.size);
provide<SizeRef>("my-menu-size", menuSize);
onUpdated(() => {
  menuSize.value = props.size;
});

const style = computed<StyleValue>(() => ({
  flexDirection: props.direction === "horizontal" ? "row" : "column",
  height: props.height,
}));
</script>

<style lang="scss">
@import "../../theme-default/variable.scss";

.my-menu {
  display: flex;
  margin: 0;
  padding: 0;
  list-style: none;

  @each $type in $Types {
    &--#{$type} {
      @if $type == $TypeDefault {
        color: var(--my-color-text-regular);
        .my-menu-item {
          &:hover {
            color: var(--my-color-text-primary);
          }
          &.active {
            color: var(--my-color-white);
            background-color: var(--my-color-text-primary);
          }
        }
      } @else {
        color: var(--my-color-text-primary);
        .my-menu-item {
          &:hover {
            color: var(--my-color-#{$type});
          }
          &.active {
            color: var(--my-color-white);
            background-color: var(--my-color-#{$type});
          }
        }
      }
    }
  }
}
</style>
