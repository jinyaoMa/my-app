<template>
  <form class="my-form" @submit="handleSubmit">
    <slot></slot>
  </form>
</template>

<script setup lang="ts" name="MyForm">
import { onUpdated, provide, Ref, ref, withDefaults } from "vue";
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
    labelWidth: "",
  }
);

const formSize = ref<Size>(props.size);
const formLabelPosition = ref<LabelPosition>(props.labelPosition);
const formLabelWidth = ref<string>(props.labelWidth);
provide<SizeRef>("my-form-size", formSize);
provide<LabelPositionRef>("my-form-label-position", formLabelPosition);
provide<Ref<string>>("my-form-label-width", formLabelWidth);
onUpdated(() => {
  formSize.value = props.size;
  formLabelPosition.value = props.labelPosition;
  formLabelWidth.value = props.labelWidth;
});

const handleSubmit = (payload: Event) => {
  payload.preventDefault();
};
</script>
