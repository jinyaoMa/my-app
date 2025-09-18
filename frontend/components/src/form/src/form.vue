<script setup lang="ts">
import { computed, ref } from 'vue'
import type { StyleValue } from 'vue'
import {} from './form'
import type { FormProps } from './form'

defineOptions({
  name: 'MyForm'
})

const props = withDefaults(defineProps<FormProps>(), {})

// dependency injections here

const classList = computed(() => {
  return {}
})

const styleList = computed((): StyleValue => {
  return {}
})
</script>

<template>
  <div class="my-form" :class="classList" :style="styleList">MyForm</div>
</template>
