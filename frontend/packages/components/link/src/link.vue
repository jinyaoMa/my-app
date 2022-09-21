<template>
  <a class="my-link"></a>
</template>

<script setup lang="ts" name="MyLink">
import { withDefaults } from "vue";

const props = withDefaults(
  defineProps<{
    type?: "primary" | "success" | "warning" | "info" | "danger" | "default";
    underline?: boolean;
  }>(),
  {
    type: "default",
    underline: true,
  }
);
</script>

<style lang="scss">
.my-link {
  text-decoration: none;
}
</style>
