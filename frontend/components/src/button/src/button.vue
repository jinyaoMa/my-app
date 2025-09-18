<script setup lang="ts">
import { computed, ref } from 'vue'
import type { StyleValue } from 'vue'
import {} from './button'
import type { ButtonProps } from './button'
import { ColorType, ColorTypeHelper, Mode, ModeHelper, injectRwMode } from '../../@global'
import { MyIcon } from '../../icon'

defineOptions({
  name: 'MyButton'
})

const props = withDefaults(defineProps<ButtonProps>(), {
  type: ColorType.primary
})

// dependency injections here
const mode = injectRwMode(Mode.normal)

const classList = computed(() => {
  const modeString = ModeHelper.toString(mode.proxy.value)
  const typeString = ColorTypeHelper.toString(props.type)
  return {
    [`my-button__mode_${modeString}`]: modeString,
    [`my-button__type_${typeString}`]: typeString,
    [`my-button__round`]: props.round
  }
})

const styleList = computed((): StyleValue => {
  return {}
})
</script>

<template>
  <button class="my-button" :class="classList" :style="styleList">
    <MyIcon v-if="iconName" :name="iconName"></MyIcon>
    <span v-if="text" class="my-button-text">{{ text }}</span>
    <slot v-else></slot>
  </button>
</template>
