import chalk from 'chalk'
import { createComponent } from '../create_component/create.ts'
import { createPage } from '../create_test_vue/createPage.ts'
import { createVitepressComp } from '../create_vitepress_comp/create.ts'

const componentsPath = './frontend/components/src'

function create() {
  const commands = process.argv.splice(2)

  if (commands.length === 0) {
    console.log(chalk.red('缺少必要参数'))
    return
  }
  for (const item of commands) {
    createComponent(componentsPath, item)
    createVitepressComp(item)
  }
  createPage(commands)
}

create()
