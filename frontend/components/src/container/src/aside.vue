<script setup lang="ts">
import { computed, ref } from 'vue'
import type { StyleValue } from 'vue'
import {} from './aside'
import { type AsideProps } from './aside'
import { ContainerStickyHelper } from './@partial/enums'

defineOptions({
  name: 'MyAside'
})

const props = withDefaults(defineProps<AsideProps>(), {
  width: '280px'
})

// dependency injections here

const classList = computed(() => {
  return {
    [`my-aside__sticky`]: props.sticky
  }
})

const styleList = computed((): StyleValue => {
  const width = typeof props.width === 'string' ? props.width : `${props.width}px`
  const stickyArray = ContainerStickyHelper.toUniqueStringArray(props.sticky)
  return {
    width,
    ...stickyArray.reduce((a, s) => {
      a[`${s}`] = 0
      return a
    }, {} as Record<string, number>)
  }
})
</script>

<template>
  <aside class="my-aside" :class="classList" :style="styleList">
    <slot></slot>
  </aside>
</template>
