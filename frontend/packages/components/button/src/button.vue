<template>
  <button
    class="my-button"
    :class="{
      [`my-button--${currentType}`]: true,
      disabled,
    }"
    :style="style"
    @click="handleClick"
  >
    <slot></slot>
  </button>
</template>

<script setup lang="ts" name="MyButton">
import {
  computed,
  inject,
  onUpdated,
  ref,
  StyleValue,
  withDefaults,
} from "vue";
import { ResponseState, Size, SizeRef, Type } from "../../types";

const props = withDefaults(
  defineProps<{
    type?: Type;
    disabled?: boolean;
    size?: Size;
  }>(),
  {
    type: "default",
    disabled: false,
  }
);
const emit = defineEmits<{
  (event: "click", res: (state: ResponseState) => void): void;
}>();
const formSize = inject<SizeRef>("my-form-size");

const style = computed<StyleValue>(() => {
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

const currentType = ref(props.type);

let timeoutClick: number;
const handleClick = (e: MouseEvent) => {
  if (!props.disabled) {
    clearTimeout(timeoutClick);
    emit("click", (state: ResponseState) => {
      switch (state) {
        case "success":
          currentType.value = "success";
          break;
        case "warning":
          currentType.value = "warning";
          break;
        case "error":
          currentType.value = "danger";
      }
      timeoutClick = window.setTimeout(() => {
        currentType.value = props.type;
      }, 5000);
    });
  }
};
</script>

<style lang="scss">
@import "../../theme-default/variable.scss";

.my-button {
  font-size: var(--my-font-size);
  line-height: var(--my-line-height-none);
  display: inline-block;
  white-space: nowrap;
  cursor: pointer;
  border-style: solid;
  border-radius: var(--my-border-radius-sm);
  box-sizing: border-box;
  transition: 0.2s;

  &.disabled {
    cursor: not-allowed;
  }

  @each $type in $Types {
    &--#{$type} {
      @if $type == $TypeDefault {
        color: var(--my-color-text-primary);
        background-color: var(--my-color-background-base);
        border-color: var(--my-color-border-lighter);
        &:hover {
          color: var(--my-color-primary);
        }
        &.disabled {
          color: var(--my-color-text-secondary);
        }
      } @else {
        color: var(--my-color-#{$type});
        background-color: var(--my-color-#{$type}-3);
        border-color: var(--my-color-#{$type}-2);
        &:hover {
          border-color: var(--my-color-#{$type}-1);
        }
        &.disabled {
          color: var(--my-color-#{$type}-2);
          border-color: var(--my-color-#{$type}-3);
        }
      }
    }
  }
}
</style>
