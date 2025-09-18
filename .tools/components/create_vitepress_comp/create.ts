import fs from 'node:fs'

export function createVitepressComp(compName: string) {
  const dir = `./frontend/docs/guide/components/${compName}.md`
  const template = `---
sider_text="侧边栏名称"
---`

  fs.writeFile(dir, template, (err) => {
    if (!err) {
      console.log(`[O] ${dir} 组件文档创建成功`)
    } else {
      console.log(`[X] ${dir} 组件文档创建失败`, err)
    }
  })
}
