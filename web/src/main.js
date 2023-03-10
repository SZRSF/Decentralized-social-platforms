/*
  该文件是整个项目的入口文件
 */
//引入Vue
import Vue from 'vue'
//引入App组件，它是所有组件的夫组件
import App from './App.vue'
//引入路由
import router from "@/router";
//引入ElementUI组件库
import ElementUI from 'element-ui';
//引入ElementUI全部样式
import 'element-ui/lib/theme-chalk/index.css';
//导入全局样式表
import './assets/css/global.css'
//引入仓库
import store from "@/store";

//应用
Vue.use(ElementUI)

//创建Vue实例对象 ---vm
new Vue({
  //将App组件放入容器中
  render: h => h(App),
  //注册路由:地下的写法KV一致省略v[router小写的]
  //注册路由信息：当这里书写router的时候，组件身上就拥有$route,$router属性
  router,
  //注册仓库，组件实例身上会多一个属性$store属性
  store
}).$mount('#app')
