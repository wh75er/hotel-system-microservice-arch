import { createApp } from 'vue'
import App from './App.vue'
import mitt from 'mitt'
import User from './userInstance.js'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from './router.js'
import GatewayClient from "@/grpc-client";

const emitter = mitt()

const app = createApp(App)
    .use(router)
    .use(ElementPlus)

app.config.globalProperties.emitter = emitter;
app.config.globalProperties.gatewayClient = new GatewayClient()
app.config.globalProperties.userSingletone = new User()
app.mount('#app')
