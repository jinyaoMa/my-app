import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import MyAppDesign from "@jinyaoma/my-app-design";

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(MyAppDesign)

app.mount('#app')
