<template>
  <li class="my-menu-group" :style="style">
    <span class="my-menu-group__title" :style="titleStyle">{{
      props.title
    }}</span>
    <ul class="my-menu-group__menu">
      <slot></slot>
    </ul>
  </li>
</template>

<script setup lang="ts" name="MyMenuGroup">
import { computed, inject, ref, StyleValue, withDefaults } from "vue";
import { Size, SizeRef } from "../../types";

const props = withDefaults(
  defineProps<{
    title: string;
    size?: Size;
  }>(),
  {
    title: "[Title]",
  }
);
const menuSize = inject<SizeRef>("my-menu-size") || ref(props.size);

const style = computed<StyleValue>(() => {
  const size = props.size || menuSize.value;
  let marginBottom = 0.6;
  switch (size) {
    case "large":
      marginBottom += 0.3;
      break;
    case "small":
      marginBottom -= 0.3;
  }
  return {
    marginBottom: marginBottom + "em",
  };
});

const titleStyle = computed<StyleValue>(() => {
  const size = props.size || menuSize.value;
  let lineHeight = 2.6;
  switch (size) {
    case "large":
      lineHeight += 0.3;
      break;
    case "small":
      lineHeight -= 0.3;
  }
  return {
    lineHeight,
  };
});
</script>

<style lang="scss">
.my-menu-group {
  display: flex;
  flex-direction: column;

  &__title {
    font-size: var(--my-font-size-xs);
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
