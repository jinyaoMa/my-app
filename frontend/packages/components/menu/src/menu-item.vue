<template>
  <li
    class="my-menu-item"
    :class="{
      active: to === $route.path,
    }"
    :style="style"
    @click="router.push(to)"
  >
    <slot></slot>
  </li>
</template>

<script setup lang="ts" name="MyMenuItem">
import { StyleValue, computed, inject, withDefaults } from "vue";
import { useRouter } from "vue-router";
import { Size, SizeRef } from "../../types";

const router = useRouter();

const props = withDefaults(
  defineProps<{
    size?: Size;
    to?: string;
  }>(),
  {
    to: "#",
  }
);
const menuSize = inject<SizeRef>("my-menu-size");

const style = computed<StyleValue>(() => {
  const size = props.size || menuSize?.value;
  let lineHeight = 2.6;
  let space = "--my-space";
  let radius = "--my-border-radius";
  switch (size) {
    case "large":
      lineHeight += 0.3;
      space += "-lg";
      radius += "-lg";
      break;
    case "small":
      lineHeight -= 0.3;
      space += "-sm";
      radius += "-sm";
  }
  return {
    lineHeight,
    padding: `0 var(${space})`,
    borderRadius: `var(${radius})`,
  };
});
</script>

<style lang="scss">
.my-menu-item {
  cursor: pointer;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  transition: 0.2s;
}
</style>
