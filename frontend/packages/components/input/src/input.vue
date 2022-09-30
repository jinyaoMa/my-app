<template>
  <div class="my-input" :style="style">
    <input
      type="text"
      class="my-input__field"
      :class="{
        disabled,
      }"
      :style="fieldStyle"
      v-model="model"
      :name="props.name"
      :id="props.name"
    />
    <slot name="append"></slot>
  </div>
</template>

<script setup lang="ts" name="MyInput">
import {
  computed,
  inject,
  onUpdated,
  ref,
  StyleValue,
  withDefaults,
} from "vue";
import { Size, SizeRef } from "../../types";

const props = withDefaults(
  defineProps<{
    name: string;
    displayValue?: string;
    size?: Size;
    disabled?: boolean;
  }>(),
  {
    disabled: false,
  }
);
const formSize = inject<SizeRef>("my-form-size");

const style = computed<StyleValue>(() => {
  const size = props.size || formSize?.value;
  let space = "--my-space";
  switch (size) {
    case "large":
      space += "-lg";
      break;
    case "small":
      space += "-sm";
  }
  return {
    gap: `var(${space})`,
  };
});

const fieldStyle = computed<StyleValue>(() => {
  const size = props.size || formSize?.value;
  let space = "--my-space";
  let spaceStartEnd = 0.6;
  let width = "--my-border-width";
  switch (size) {
    case "large":
      space += "-lg";
      spaceStartEnd += 0.3;
      width += "-lg";
      break;
    case "small":
      space += "-sm";
      spaceStartEnd -= 0.3;
      width += "-sm";
  }
  return {
    padding: `${spaceStartEnd}em var(${space})`,
    borderWidth: `var(${width})`,
  };
});

const model = ref(props.displayValue);
onUpdated(() => {
  model.value = props.displayValue;
});
</script>

<style lang="scss">
.my-input {
  box-sizing: border-box;
  max-width: 100%;
  display: flex;
  flex-direction: row;

  &__field {
    flex-grow: 1;
    box-sizing: border-box;
    display: inline-block;
    border-style: solid;
    border-color: var(--my-color-border-lighter);
    border-radius: var(--my-border-radius-sm);
    background-color: var(--my-color-background-base);
    line-height: var(--my-line-height-none);
    outline: none;

    &:not(.disabled):hover {
      border-color: var(--my-color-border-base);
    }

    &:not(.disabled):focus {
      border-color: var(--my-color-primary);
    }

    &.disabled {
      color: var(--my-color-text-placeholder);
      background-color: var(--my-color-border-extra-light);
      cursor: not-allowed;
    }
  }
}
</style>
