import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

// import MyDesign from '@jinyaoma/my-design'
import MyDesign from '@jinyaoma/my-components'

const app = createApp(App)

app.use(MyDesign)
app.use(createPinia())
app.use(router)

app.mount('#app')
