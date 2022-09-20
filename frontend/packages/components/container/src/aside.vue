<template>
  <aside class="my-aside" :style="style">
    <slot></slot>
  </aside>
</template>

<script lang="ts" setup name="MyAside">
import type { StyleValue } from "vue";
import { computed, withDefaults } from "vue";

const props = withDefaults(
  defineProps<{
    width: string;
  }>(),
  {
    width: "280px",
  }
);

const style = computed<StyleValue>(() => ({
  width: props.width,
}));
</script>
