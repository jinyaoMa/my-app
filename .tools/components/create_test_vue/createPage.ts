import fs from 'node:fs'

const pagesPath = './frontend/playground/src/pages'

export function createPage(commands: any[]) {
  const template = `<script setup lang="ts">

</script>

<template>
  <div>
    component
  </div>
</template>

<style scoped lang="scss">

</style>
`

  for (let i = 0; i < commands.length; i++) {
    const path = `${pagesPath}/${commands[i]}`
    const target = `${path}/index.vue`

    fs.mkdir(
      path,
      {
        recursive: true
      },
      (err) => {
        if (!err) {
          fs.writeFile(target, template, (err) => {
            if (!err) {
              console.log(`[O] ${target} 新测试组件生成完成`)
            } else {
              console.log(`[X] ${target} 新测试组件生成错误`, err)
            }
          })
        } else {
          console.log(`[X] ${target} 新测试组件创建失败`, err)
          return
        }
      }
    )
  }
}
