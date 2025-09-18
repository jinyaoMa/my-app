<script setup lang="ts">
import { computed, ref } from 'vue'
import type { StyleValue } from 'vue'
import {} from './footer'
import { type FooterProps } from './footer'
import { ContainerStickyHelper } from './@partial/enums'

defineOptions({
  name: 'MyFooter'
})

const props = withDefaults(defineProps<FooterProps>(), {})

// dependency injections here

const classList = computed(() => {
  return {
    [`my-footer__sticky`]: props.sticky
  }
})

const styleList = computed((): StyleValue => {
  const height =
    props.height === undefined ? undefined : typeof props.height === 'string' ? props.height : `${props.height}px`
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
  <footer class="my-footer" :class="classList" :style="styleList">
    <slot></slot>
  </footer>
</template>
