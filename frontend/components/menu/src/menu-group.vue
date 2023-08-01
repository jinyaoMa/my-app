<template>
  <li class="my-menu-group" :style="style">
    <span v-if="props.title" class="my-menu-group__title" :style="titleStyle">{{
      props.title
    }}</span>
    <ul class="my-menu-group__menu">
      <slot></slot>
    </ul>
  </li>
</template>

<script setup lang="ts" name="MyMenuGroup">
import { computed, inject, StyleValue, withDefaults } from "vue";
import { Size, SizeRef } from "../../types";

const props = withDefaults(
  defineProps<{
    title?: string;
    size?: Size;
  }>(),
  {}
);
const menuSize = inject<SizeRef>("my-menu-size");

const style = computed<StyleValue>(() => {
  const size = props.size || menuSize?.value;
  let marginBottom = 1;
  switch (size) {
    case "large":
      marginBottom += 0.5;
      break;
    case "small":
      marginBottom -= 0.5;
  }
  return {
    marginBottom: marginBottom + "em",
  };
});

const titleStyle = computed<StyleValue>(() => {
  const size = props.size || menuSize?.value;
  let marginBottom = 0.5;
  switch (size) {
    case "large":
      marginBottom += 0.2;
      break;
    case "small":
      marginBottom -= 0.2;
  }
  return {
    marginBottom: marginBottom + "em",
  };
});
</script>

<style lang="scss">
.my-menu-group {
  display: flex;
  flex-direction: column;

  &__title {
    font-size: var(--my-font-size-xs);
    line-height: var(--my-line-height-sm);
    color: var(--my-color-text-secondary);
  }

  &__menu {
    display: flex;
    flex-direction: column;
    margin: 0;
    padding: 0;
    list-style: none;
  }
}
</style>
