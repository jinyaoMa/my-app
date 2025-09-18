<script setup lang="ts">
import { computed, ref } from 'vue'
import type { StyleValue } from 'vue'
import {} from './link'
import type { LinkProps } from './link'

defineOptions({
  name: 'MyLink'
})

const props = withDefaults(defineProps<LinkProps>(), {})

// dependency injections here

const classList = computed(() => {
  return {}
})

const styleList = computed((): StyleValue => {
  return {}
})
</script>

<template>
  <div class="my-link" :class="classList" :style="styleList">MyLink</div>
</template>
