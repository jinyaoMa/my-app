import DefaultTheme from 'vitepress/theme'
import './style/index.scss'

// 全局引入组件库（开发环境）
import MyDesign from '../../../components'
import '../../../components/style/index.scss'

export default {
  extends: DefaultTheme, // or ...DefaultTheme
  enhanceApp({ app }) {
    app.use(MyDesign)
  }
}
