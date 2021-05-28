import { createApp } from 'vue'
import App from './App.vue'
import installElementPlus from './plugins/element'
import router from './router'
import '@/assets/css/global.css'

const app = createApp(App).use(router)
installElementPlus(app)
app.mount('#app')