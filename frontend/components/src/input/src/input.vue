<script setup lang="ts">
import { computed, ref } from 'vue'
import type { StyleValue } from 'vue'
import {} from './input'
import type { InputProps } from './input'

defineOptions({
  name: 'MyInput'
})

const props = withDefaults(defineProps<InputProps>(), {})

// dependency injections here

const classList = computed(() => {
  return {}
})

const styleList = computed((): StyleValue => {
  return {}
})
</script>

<template>
  <div class="my-input" :class="classList" :style="styleList">MyInput</div>
</template>
