import '../style/index.scss'
import type { App } from 'vue'
import * as components from './components'
export * from './components'

const install = (app: App) => {
  for (const n in components) {
    app.use((components as any)[n])
  }
}

export default install
