import { createApp } from 'vue'
import './asset/style/index.scss'
import App from './App.vue'
import router from './router/index'

// 全局引入组件
import MyDesign from '@jinyaoma/my-components'

const app = createApp(App)
app.use(MyDesign)
app.use(router)
app.mount('#app')
