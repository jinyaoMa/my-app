<script setup lang="ts">
import { computed, ref } from 'vue'
import type { StyleValue } from 'vue'
import {} from './header'
import { type HeaderProps } from './header'
import { ContainerStickyHelper } from './@partial/enums'

defineOptions({
  name: 'MyHeader'
})

const props = withDefaults(defineProps<HeaderProps>(), {
  height: '60px'
})

// dependency injections here

const classList = computed(() => {
  return {
    [`my-header__sticky`]: props.sticky
  }
})

const styleList = computed((): StyleValue => {
  const height = typeof props.height === 'string' ? props.height : `${props.height}px`
  const stickyArray = ContainerStickyHelper.toUniqueStringArray(props.sticky)
  return {
    height,
    minHeight: height,
    ...stickyArray.reduce((a, s) => {
      a[`${s}`] = 0
      return a
    }, {} as Record<string, number>)
  }
})
</script>

<template>
  <header class="my-header" :class="classList" :style="styleList">
    <slot></slot>
  </header>
</template>
