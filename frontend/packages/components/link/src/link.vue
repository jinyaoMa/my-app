<template>
  <a
    class="my-link"
    :class="{
      [`my-link--${props.type}`]: true,
      underline: !disabled,
      disabled,
    }"
    :href="href"
    @click="handleClick"
  >
    <slot></slot>
  </a>
</template>

<script setup lang="ts" name="MyLink">
import { computed, withDefaults } from "vue";
import { Type } from "../../types";

const props = withDefaults(
  defineProps<{
    type?: Type;
    underline?: boolean;
    disabled?: boolean;
    href?: string;
  }>(),
  {
    type: "default",
    underline: true,
    disabled: false,
    href: "",
  }
);
const emit = defineEmits<{
  (event: "click", e: MouseEvent): void;
}>();

const href = computed(() => {
  return props.disabled || !props.href ? undefined : props.href;
});

const handleClick = (e: MouseEvent) => {
  props.disabled || emit("click", e);
};
</script>

<style lang="scss">
@import "../../theme-default/variable.scss";

.my-link {
  font-size: var(--my-font-size);
  line-height: var(--my-line-height);
  text-decoration: none;
  display: inline-block;
  position: relative;
  white-space: nowrap;
  cursor: pointer;

  &.underline {
    &::before {
      content: "";
      display: block;
      position: absolute;
      bottom: 0;
      right: 0;
      width: 0%;
      border-bottom: var(--my-border-width-sm) solid;
      transition: width 0.2s ease-in;
    }
    &:hover::before {
      left: 0;
      width: 100%;
    }
  }

  &.disabled {
    cursor: not-allowed;
  }

  @each $type in $Types {
    &--#{$type} {
      @if $type == $TypeDefault {
        color: var(--my-color-text-primary);
        &.underline::before {
          border-bottom-color: var(--my-color-primary);
        }
        &:hover {
          color: var(--my-color-primary);
        }
        &.disabled {
          color: var(--my-color-text-secondary);
        }
      } @else {
        color: var(--my-color-#{$type});
        &.underline::before {
          border-bottom-color: var(--my-color-#{$type});
        }
        &:hover {
          color: var(--my-color-#{$type}-1);
        }
        &.disabled {
          color: var(--my-color-#{$type}-2);
        }
      }
    }
  }
}
</style>
