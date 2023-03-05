import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

// 加载 Vant 核心组件库
import Vant from 'vant'

// 加载 Vant 全局样式
import 'vant/lib/index.css'
// 引入Vant 全部样式
import 'vant/lib/index.less'

// 加载动态设计 REM 基准值
import 'amfe-flexible'

import qs from 'qs'

// 加载全局样式
import './styles/index.less'

// 加载tinymce
import tinymce from 'tinymce'
import VueTinymce from '@packy-tang/vue-tinymce'

// 加载 dayjs 初始化配置
import './utils/datjs'

Vue.prototype.$tinymce = tinymce // 将全局tinymce对象指向给Vue作用域下
Vue.use(VueTinymce) // 安装vue的tinymce组件

// 注册使用 Vant 组件库
Vue.use(Vant)

Vue.prototype.$qs = qs
Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
