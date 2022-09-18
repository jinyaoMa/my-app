<script setup lang="ts" name="MyIcon">
defineProps<{
  name: string;
  fixWidth: boolean;
}>();
</script>

<template>
  <i
    :class="`my-icon-${name || 'jinyao-ma'}${fixWidth ? ' fix-width' : ''}`"
  ></i>
</template>
