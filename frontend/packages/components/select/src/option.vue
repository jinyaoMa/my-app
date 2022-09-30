<template>
  <option class="my-option">
    <slot></slot>
  </option>
</template>

<script setup lang="ts" name="MyOption">
import { inject, withDefaults } from "vue";
import { Size, SizeRef } from "../../types";

const props = withDefaults(
  defineProps<{
    size?: Size;
  }>(),
  {}
);
const formSize = inject<SizeRef>("my-form-size");
</script>
