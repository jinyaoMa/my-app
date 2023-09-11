<template>
  <div class="my-form-item" :style="style">
    <label class="my-form-item__label" :for="props.for" :style="labelStyle">
      {{ props.label }}
    </label>
    <div class="my-form-item__inner">
      <slot></slot>
    </div>
  </div>
</template>

<script setup lang="ts" name="MyFormItem">
import { computed, inject, Ref, StyleValue, withDefaults } from "vue";
import { LabelPosition, LabelPositionRef, Size, SizeRef } from "../../types";

const props = withDefaults(
  defineProps<{
    label?: string;
    for?: string;
    size?: Size;
    labelPosition?: LabelPosition;
    labelWidth?: string;
  }>(),
  {}
);
const formSize = inject<SizeRef>("my-form-size");
const formLabelPosition = inject<LabelPositionRef>("my-form-label-position");
const formLabelWidth = inject<Ref<string>>("my-form-label-width");

const style = computed<StyleValue>(() => {
  const size = props.size || formSize?.value;
  const pos = props.labelPosition || formLabelPosition?.value;
  let space = "--my-space";
  switch (size) {
    case "large":
      space += "-lg";
      break;
    case "small":
      space += "-sm";
  }
  return {
    flexDirection: pos === "top" ? "column" : "row",
    alignItems: pos === "top" ? "start" : "center",
    marginBottom: `var(${space})`,
  };
});

const labelStyle = computed<StyleValue>(() => {
  const size = props.size || formSize?.value;
  const width = props.labelWidth || formLabelWidth?.value;
  let space = "--my-space";
  switch (size) {
    case "large":
      space += "-lg";
      break;
    case "small":
      space += "-sm";
  }
  return {
    marginRight: `var(${space})`,
    width,
    minWidth: width,
  };
});
</script>

<style lang="scss">
.my-form-item {
  display: flex;
  align-items: center;

  &:last-child {
    margin-bottom: 0 !important;
  }

  &__label {
    display: inline-block;
  }

  &__inner {
    flex-grow: 1;

    > .my-button:not(:last-child) {
      margin-right: 0.85em;
    }
  }
}
</style>
