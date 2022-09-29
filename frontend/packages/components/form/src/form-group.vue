<template>
  <fieldset class="my-form-group" :style="style">
    <legend class="my-form-group__legend">{{ props.legend }}</legend>
    <slot></slot>
  </fieldset>
</template>

<script setup lang="ts" name="MyFormGroup">
import { computed, inject, StyleValue, withDefaults } from "vue";
import { Size, SizeRef } from "../../types";

const props = withDefaults(
  defineProps<{
    legend?: string;
    size?: Size;
    labelWidth?: string;
  }>(),
  {}
);
const formSize = inject<SizeRef>("my-form-size");

const style = computed<StyleValue>(() => {
  const size = props.size || formSize?.value;
  let width = "--my-border-width";
  let radius = "--my-border-radius";
  let space = "--my-space";
  switch (size) {
    case "large":
      width += "-lg";
      radius += "-lg";
      space += "-lg";
      break;
    case "small":
      width += "-sm";
      radius += "-sm";
      space += "-sm";
  }
  return {
    borderWidth: `var(${width})`,
    borderRadius: `var(${radius})`,
    padding: `var(${space}) calc(0.4em + var(${space})) calc(0.275em + var(${space}))`,
    margin: `0 var(${width}) var(${space})`,
  };
});
</script>

<style lang="scss">
.my-form-group {
  border-style: solid;
  border-color: var(--my-color-border-lighter);

  &__legend {
    font-size: var(--my-font-size-xs);
    color: var(--my-color-text-secondary);
  }
}
</style>
