<script setup lang="ts">
import { computed, ref } from 'vue'
import type { StyleValue } from 'vue'
import {} from './main'
import { type MainProps } from './main'
import { ContainerStickyHelper } from './@partial/enums'

defineOptions({
  name: 'MyMain'
})

const props = withDefaults(defineProps<MainProps>(), {})

// dependency injections here

const classList = computed(() => {
  return {
    [`my-main__sticky`]: props.sticky
  }
})

const styleList = computed((): StyleValue => {
  const stickyArray = ContainerStickyHelper.toUniqueStringArray(props.sticky)
  return {
    ...stickyArray.reduce((a, s) => {
      a[`${s}`] = 0
      return a
    }, {} as Record<string, number>)
  }
})
</script>

<template>
  <main class="my-main" :class="classList" :style="styleList">
    <slot></slot>
  </main>
</template>
