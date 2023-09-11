<template>
  <div
    class="my-input"
    :class="{
      hasPrepend,
      hasAppend,
    }"
  >
    <span class="my-input__prepend" :style="prependStyle">
      <slot name="prepend"></slot>
    </span>
    <input
      class="my-input__field"
      :class="{
        disabled,
      }"
      :style="fieldStyle"
      :type="props.type"
      :name="props.name"
      :id="props.name"
      :placeholder="placeholder"
      :step="step"
      :min="min"
      :max="max"
      :disabled="disabled"
      v-model="value"
      @change="handleChange"
    />
    <span class="my-input__append" :style="appendStyle">
      <slot name="append"></slot>
    </span>
  </div>
</template>

<script setup lang="ts" name="MyInput">
import { computed, inject, StyleValue, useSlots, withDefaults } from "vue";
import { Size, SizeRef } from "../../types";

const props = withDefaults(
  defineProps<{
    type?: "text" | "number" | "checkbox" | "password";
    name?: string;
    modelValue: string | number | boolean;
    placeholder?: string;
    size?: Size;
    disabled?: boolean;
    width?: string;
    step?: number;
    min?: number;
    max?: number;
  }>(),
  {
    type: "text",
    disabled: false,
    width: "auto",
    step: 1,
  }
);
const emit = defineEmits<{
  (event: "update:modelValue", value: string | number | boolean): void;
  (event: "change", payload: Event): void;
}>();
const formSize = inject<SizeRef>("my-form-size");

const slots = useSlots();
const hasPrepend = computed(() => {
  if (slots && slots.prepend) {
    return true;
  }
  return false;
});
const hasAppend = computed(() => {
  if (slots && slots.append) {
    return true;
  }
  return false;
});

const value = computed({
  get(): string | number | boolean {
    return props.modelValue;
  },
  set(v: string | number | boolean) {
    emit("update:modelValue", v);
  },
});

const prependStyle = computed<StyleValue>(() => {
  if (!hasPrepend.value) {
    return {};
  }
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
    borderWidth: `var(${width}) 0 var(${width}) var(${width})`,
  };
});

const appendStyle = computed<StyleValue>(() => {
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
    marginLeft: hasAppend.value ? `var(${space})` : undefined,
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
  if (props.type === "checkbox") {
    let extra = 0.25;
    spaceStartEnd += extra;
    return {
      padding: `${spaceStartEnd}em calc(var(${space}) + ${extra * 3}em)`,
      margin: `${extra}em 0`,
      borderRadius: `${spaceStartEnd + extra * 3}em`,
      borderWidth: `var(${width})`,
    };
  }
  return {
    padding: `${spaceStartEnd}em var(${space})`,
    borderWidth: `var(${width})`,
    width: props.width,
  };
});

const handleChange = (payload: Event) => {
  props.disabled || emit("change", payload);
};
</script>

<style lang="scss">
.my-input {
  box-sizing: border-box;
  max-width: 100%;
  display: flex;
  flex-direction: row;

  &.hasPrepend &__prepend {
    box-sizing: border-box;
    display: inline-flex;
    align-items: center;
    font-weight: bold;
    border-style: solid;
    border-color: var(--my-color-border-lighter);
    border-radius: var(--my-border-radius-sm) 0 0 var(--my-border-radius-sm);
    color: var(--my-color-text-placeholder);
    background-color: var(--my-color-border-extra-light);
    line-height: var(--my-line-height-none);
  }

  &.hasAppend &__append {
    display: inline-flex;
  }

  &.hasPrepend &__field {
    border-radius: 0 var(--my-border-radius-sm) var(--my-border-radius-sm) 0;
  }

  &__field {
    box-sizing: border-box;
    display: inline-block;
    border-style: solid;
    border-color: var(--my-color-border-lighter);
    border-radius: var(--my-border-radius-sm);
    color: var(--my-color-text-primary);
    background-color: var(--my-color-background-base);
    line-height: var(--my-line-height-none);
    outline: none;

    &::placeholder {
      color: var(--my-color-text-placeholder);
    }

    &[type="checkbox"] {
      appearance: none;
      position: relative;
      border-color: var(--my-color-danger-1);
      background-color: var(--my-color-danger-3);

      &::before {
        content: "";
        display: block;
        position: absolute;
        top: 0;
        left: 0;
        height: 100%;
        width: 50%;
        border-radius: 50%;
        transform: scale(0.6);
        background-color: var(--my-color-danger);
        transition: left 0.2s;
      }

      &:hover {
        border-color: var(--my-color-danger);
        &::before {
          background-color: var(--my-color-danger);
        }
      }

      &:checked {
        border-color: var(--my-color-success-1);
        background-color: var(--my-color-success-3);
        &::before {
          background-color: var(--my-color-success);
          left: 50%;
        }
        &:hover {
          border-color: var(--my-color-success);
          &::before {
            background-color: var(--my-color-success);
          }
        }
      }
    }

    &:not(.disabled):not([type="checkbox"]):hover {
      border-color: var(--my-color-border-base);
    }

    &:not(.disabled):not([type="checkbox"]):focus {
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
