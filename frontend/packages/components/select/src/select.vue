<template>
  <select
    class="my-select"
    :style="style"
    v-model="model"
    :name="props.name"
    :id="props.name"
    @change="handleChange"
  >
    <slot></slot>
  </select>
</template>

<script setup lang="ts" name="MySelect">
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
    displayValue?: string | number;
    size?: Size;
  }>(),
  {}
);
const emit = defineEmits<{
  (event: "change", value: any): void;
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

const model = ref(props.displayValue);
onUpdated(() => {
  model.value = props.displayValue;
});

const handleChange = (payload: Event) => {
  model.value != props.displayValue && emit("change", model.value);
};
</script>

<style lang="scss">
.my-select {
  display: inline-block;
  outline: none;
  border-style: solid;
  border-color: var(--my-color-border-lighter);
  border-radius: var(--my-border-radius-sm);
  background-color: var(--my-color-background-base);
  line-height: var(--my-line-height-none);

  &:hover {
    border-color: var(--my-color-border-base);
  }

  &:focus {
    border-color: var(--my-color-primary);
  }
}
</style>
