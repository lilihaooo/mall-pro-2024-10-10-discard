import 'element-plus/es/components/message/style/css'
import 'element-plus/es/components/loading/style/css'
import 'element-plus/es/components/notification/style/css'
import 'element-plus/es/components/message-box/style/css'
import 'element-plus/theme-chalk/dark/css-vars.css'
import './style/element_visiable.scss'

import { createApp } from 'vue'
// 引入admin-vue前端初始化相关内容
import './core/admin-vue'
// 引入封装的router
import router from '@/router/index'
import '@/permission'
import run from '@/core/admin-vue.js'
import auth from '@/directive/auth'
import {store} from '@/pinia'
import App from './App.vue'
import { initDom } from './utils/positionToCode'



initDom()


const app = createApp(App)
app.config.productionTip = false



app
    .use(run)
    .use(store)
    .use(auth)
    .use(router)
    app.mount('#app')

export default app
