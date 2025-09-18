import { computed, ComputedGetter, inject, provide, ref, Ref, watchEffect } from 'vue'
import { Mode, ModeHelper } from './enums'
import type { ModeKey } from './enums'
import { defineRwMode, type RwProxy } from './defines'

const ROOT: symbol = Symbol('key to check if component is root')

export const injectAndCheckRoot = (): boolean => {
  const isRoot = inject(ROOT, true)
  if (isRoot) provide(ROOT, false)
  return isRoot
}

const MODE: symbol = Symbol('key to get provided mode')

export const injectRwMode = (def: Mode): RwProxy<Mode> => {
  return inject(MODE, defineRwMode(def))
}

export const provideRwMode = (rwMode: RwProxy<Mode>) => {
  provide(MODE, rwMode)
}

export const injectAndComputeMode = (
  def: Mode,
  propModeGetter: ComputedGetter<Mode | ModeKey | undefined>,
  override = false
) => {
  const propMode = computed(propModeGetter)
  const mode = defineRwMode(def)
  const injectedMode = injectRwMode(def)
  watchEffect(() => {
    if (propMode.value !== undefined) {
      // props.mode set, use props.mode
      mode.update(ModeHelper.toEnum(propMode.value))
      if (override) provideRwMode(mode) // provide new mode
    } else {
      // props.mode unset, use injected mode
      mode.update(injectedMode.proxy.value)
      if (override) provideRwMode(injectedMode) // provide injected mode
    }
  })
  return { mode, injectedMode }
}
