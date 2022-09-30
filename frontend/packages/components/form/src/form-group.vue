<template>
  <fieldset class="my-form-group" :style="style">
    <legend class="my-form-group__legend" :style="legendStyle">
      {{ props.legend }}
    </legend>
    <slot></slot>
  </fieldset>
</template>

<script setup lang="ts" name="MyFormGroup">
import {
  computed,
  inject,
  onUpdated,
  provide,
  ref,
  Ref,
  StyleValue,
  withDefaults,
} from "vue";
import { SizeRef } from "../../types";

const props = withDefaults(
  defineProps<{
    legend?: string;
    labelWidth?: string;
  }>(),
  {}
);
const formSize = inject<SizeRef>("my-form-size");
const formLabelWidth = props.labelWidth
  ? ref<string>(props.labelWidth)
  : inject<Ref<string>>("my-form-label-width") || ref<string>("");
provide<Ref<string>>("my-form-label-width", formLabelWidth);
onUpdated(() => {
  formLabelWidth.value = props.labelWidth || formLabelWidth.value;
});

const style = computed<StyleValue>(() => {
  const size = formSize?.value;
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
    padding: `calc(0.35em + var(${space})) calc(0.75em + var(${space})) calc(0.625em + var(${space}))`,
    margin: `0 var(${width}) var(${space})`,
  };
});

const legendStyle = computed<StyleValue>(() => {
  const size = formSize?.value;
  let space = "--my-space";
  switch (size) {
    case "large":
      space += "-lg";
      break;
    case "small":
      space += "-sm";
  }
  return {
    padding: `0 var(${space})`,
  };
});
</script>

<style lang="scss">
.my-form-group {
  border-style: solid;
  border-color: var(--my-color-border-extra-light);

  &__legend {
    font-size: var(--my-font-size-xs);
    color: var(--my-color-text-secondary);
  }
}
</style>
