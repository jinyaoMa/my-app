<script setup lang="ts">
import { computed, ref } from 'vue'
import type { StyleValue } from 'vue'
import type { IconProps } from './icon'
import { IconSize, IconSizeHelper } from './@partial/enums'
import '@jinyaoma/my-icons/dist/my-icon.scss'

defineOptions({
  name: 'MyIcon'
})

const props = withDefaults(defineProps<IconProps>(), {
  name: 'jinyao-ma',
  size: IconSize.regular
})

// dependency injections here

const classList = computed(() => {
  return {}
})

const styleList = computed((): StyleValue => {
  return {
    width: props.fixWidth
  }
})

const iconClassList = computed(() => {
  const sizeString = IconSizeHelper.toString(props.size)
  return {
    [`my-icon-${props.name}`]: props.name,
    [`my-icon__${sizeString}`]: sizeString
  }
})
</script>

<template>
  <span class="my-icon" :class="classList" :style="styleList">
    <i :class="iconClassList"></i>
  </span>
</template>
