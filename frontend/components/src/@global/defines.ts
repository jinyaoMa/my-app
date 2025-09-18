import { readonly, ref, Ref } from 'vue'
import { Mode } from './enums'

export interface RwProxy<T> {
  /**
   * @description readonly proxy of reactive ref value
   */
  proxy: Readonly<Ref<T, T>>
  /**
   * @description update the reactive ref value
   */
  update: (m: T) => void
}

export const defineRwMode = (mode: Mode): RwProxy<Mode> => {
  const modeRef = ref(mode)
  return {
    proxy: readonly(modeRef),
    update: (mode: Mode) => {
      modeRef.value = mode
    }
  }
}
