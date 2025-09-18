import fs from 'node:fs'

const componentsPath = './frontend/docs/guide/components'

export function createComponentSiderBar() {
  if (!fs.existsSync(componentsPath)) {
    console.log('components文件不存在，无法自动映射')
    return
  }
  const componentsDir = fs.readdirSync(componentsPath)
  if (componentsDir.length === 0) return

  const siderbar: {
    text: string
    link: string
  }[] = []

  for (const file of componentsDir) {
    const fileCurrentPath = componentsPath + `/${file}`
    const data = fs.readFileSync(fileCurrentPath).toString()
    const match = data.match(/sider_text="([^"]*)"/)
    const name = file.split('.')[0]
    const link = '/guide/components/' + name
    if (match && match.length > 1) {
      const text = match[1]

      siderbar.push({
        text,
        link
      })
    } else {
      siderbar.push({
        text: name,
        link
      })
    }
  }

  return siderbar
}
