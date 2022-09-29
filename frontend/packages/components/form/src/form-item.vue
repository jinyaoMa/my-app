<template>
  <div class="my-form-item">
    <label
      class="my-form-item__label"
      v-if="props.label"
      :for="props.for"
      :style="labelStyle"
      >{{ props.label }}</label
    >
    <slot></slot>
  </div>
</template>

<script setup lang="ts" name="MyFormItem">
import { computed, StyleValue, withDefaults } from "vue";
import { LabelPosition } from "../../types";

const props = withDefaults(
  defineProps<{
    label?: string;
    for?: string;
    labelPosition?: LabelPosition;
    labelWidth?: string;
  }>(),
  {}
);

const labelStyle = computed<StyleValue>(() => {
  const width = props.labelWidth;
  return {
    width,
  };
});
</script>

<style lang="scss">
.my-form-item {
  &:not(:last-child) {
    margin-bottom: 0.5em;
  }

  &__label {
    display: inline-block;
  }
}
</style>
