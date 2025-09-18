<script setup lang="ts">
import { computed, ref, useSlots } from 'vue'
import type { StyleValue, Component } from 'vue'
import type { ContainerProps } from './container'
import { Mode, injectAndCheckRoot, injectAndComputeMode } from '../../@global'
import { ContainerDirection, ContainerDirectionHelper } from './@partial/enums'

defineOptions({
  name: 'MyContainer'
})

const props = withDefaults(defineProps<ContainerProps>(), {})

// dependency injections here
const isRoot = injectAndCheckRoot()
const { mode } = injectAndComputeMode(Mode.normal, () => props.mode, true)

const slots = useSlots()
const directionString = computed(() => {
  if (props.direction !== undefined) {
    return ContainerDirectionHelper.toString(props.direction)
  }
  if (
    slots &&
    slots.default &&
    slots.default({}).some((vNode) => {
      const tag = (vNode.type as Component).name
      return tag === 'MyAside'
    })
  ) {
    return ContainerDirectionHelper.toString(ContainerDirection.horizontal)
  }
  return ContainerDirectionHelper.toString(ContainerDirection.vertical)
})

const classList = computed(() => {
  return {
    'css-var-default': isRoot,
    [`my-container__direction_${directionString.value}`]: directionString.value,
    'my-container__scrollable': props.scrollable
  }
})

const styleList = computed((): StyleValue => {
  const height = props.viewport ? '100vh' : typeof props.height === 'string' ? props.height : `${props.height}px`
  const width = props.viewport ? '100vw' : typeof props.width === 'string' ? props.width : `${props.width}px`
  return {
    height,
    width
  }
})
</script>

<template>
  <div class="my-container" :class="classList" :style="styleList">
    <slot></slot>
  </div>
</template>
