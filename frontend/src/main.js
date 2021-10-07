import { createApp } from 'vue'
import App from './App.vue'
import mitt from 'mitt'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from './router.js'

const emitter = mitt()

const app = createApp(App)
    .use(router)
    .use(ElementPlus)

app.config.globalProperties.emitter = emitter;
app.mount('#app')
