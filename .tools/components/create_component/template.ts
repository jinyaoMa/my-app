function toCapitalCase(v: string) {
  return (
    v[0].toUpperCase() +
    v
      .slice(1)
      .replace(/[-_](\w)/g, (match) => match[1].toUpperCase())
      .replace(/[-_]/g, '')
  )
}

export function createIndexTemplate(compName: string) {
  const originName = compName
  compName = toCapitalCase(compName)
  const name = `My${compName}`
  return `import { withInstall } from '../withInstall'
import ${compName} from './src/${originName}.vue'

export const ${name} = withInstall(${compName})

export default ${name}
`
}

export function createTsTemplate(compName: string) {
  compName = toCapitalCase(compName)
  return `export interface ${compName}Props {}`
}

export function createVueTemplate(compName: string) {
  const originName = compName
  const originClassName = `my-${compName}`

  compName = toCapitalCase(compName)
  const name = `My${compName}`

  return `<script setup lang="ts">
import { computed, ref } from 'vue'
import type { StyleValue } from 'vue'
import {} from './${originName}'
import type { ${compName}Props } from './${originName}'

defineOptions({
  name: '${name}'
})

const props = withDefaults(defineProps<${compName}Props>(), {})

// dependency injections here

const classList = computed(() => {
  return {}
})

const styleList = computed((): StyleValue => {
  return {}
})
</script>

<template>
  <div class="${originClassName}" :class="classList" :style="styleList">${name}</div>
</template>
`
}

export function createStyleTemplate(compName: string) {
  const originClassName = `my-${compName}`
  return `.${originClassName} {
}
`
}
