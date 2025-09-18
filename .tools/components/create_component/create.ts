import fs from 'node:fs'
import { createIndexTemplate, createTsTemplate, createVueTemplate, createStyleTemplate } from './template.ts'

export function createComponent(componentsPath: string, compName: string) {
  const compRoot = `${componentsPath}/${compName}`

  if (!componentExist(compRoot)) {
    console.log(`[X] ${compName} 组件存在: ${compRoot}`)
    return
  }

  const dirName = `${compRoot}/src/style`
  const indexTemplate = createIndexTemplate(compName)
  const tsTemplate = createTsTemplate(compName)
  const vueTemplate = createVueTemplate(compName)
  const styleTemplate = createStyleTemplate(compName)
  const indexTsName = `${compRoot}/index.ts`
  const tsName = `${compRoot}/src/${compName}.ts`
  const vueName = `${compRoot}/src/${compName}.vue`
  const styleName = `${dirName}/index.scss`

  insertComponent(`${componentsPath}/components.ts`, compName)

  fs.mkdir(
    dirName,
    {
      recursive: true
    },
    (err) => {
      if (!err) {
        createFile(indexTsName, indexTemplate)
        createFile(tsName, tsTemplate)
        createFile(vueName, vueTemplate)
        createFile(styleName, styleTemplate)
      } else {
        console.log('[X] 创建组件失败', err)
      }
    }
  )
  const mainStylePath = './frontend/components/style/index.scss'
  const importStyleContent = `@use '../src/${compName}/src/style/index.scss' as *;\n`
  fs.writeFile(mainStylePath, importStyleContent, { encoding: 'utf8', flag: 'a' }, (err) => {
    if (err) {
      console.log(`[X] ${mainStylePath} 全局样式引入局部样式失败`)
    } else {
      console.log(`[O] ${mainStylePath} 全局样式引入局部样式成功`)
    }
  })
}

function createFile(dir: string, template: string) {
  fs.writeFile(dir, template, (err) => {
    if (!err) {
      console.log(`[O] ${dir} 组件文件创建成功`)
    } else {
      console.log(`[X] ${dir} 组件文件创建失败`, err)
    }
  })
}

export function componentExist(compRoot: string) {
  if (fs.existsSync(compRoot)) {
    return false
  } else {
    return true
  }
}

function insertComponent(path: string, compName: string) {
  const content = `export * from './${compName}'\n`
  fs.writeFile(path, content, { encoding: 'utf8', flag: 'a' }, (err) => {
    if (err) {
      console.log(`[X] ${path} 注入组件失败`)
    } else {
      console.log(`[O] ${path} 注入组件成功`)
    }
  })
}
