<template>
  <form class="my-form">
    <slot></slot>
  </form>
</template>

<script setup lang="ts" name="MyForm">
import { onUpdated, provide, ref, withDefaults } from "vue";
import { LabelPosition, LabelPositionRef, Size, SizeRef } from "../../types";

const props = withDefaults(
  defineProps<{
    size?: Size;
    labelPosition?: LabelPosition;
    labelWidth?: string;
  }>(),
  {
    size: "default",
    labelPosition: "left",
    labelWidth: "auto",
  }
);

const formSize = ref<Size>(props.size);
const formLabelPosition = ref<LabelPosition>(props.labelPosition);
provide<SizeRef>("my-form-size", formSize);
provide<LabelPositionRef>("my-form-label-position", formLabelPosition);
onUpdated(() => {
  formSize.value = props.size;
  formLabelPosition.value = props.labelPosition;
});
</script>
