import './assets/main.css'
import 'element-plus/dist/index.css'
import { createApp } from 'vue'

import App from './App.vue'
import router from './router'
import ElementPlus from 'element-plus'
import { store } from '@/stores'

import '@/permission'

const app = createApp(App)

app.use(ElementPlus)
app.use(store)
app.use(router)

app.mount('#app')
