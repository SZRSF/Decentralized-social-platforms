import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

// 路由表
const routes = [
  {
    path: '/login',
    name: 'login',
    component: () => import('@/views/login')
  },
  {
    path: '/',
    name: 'layout',
    redirect: 'home',
    component: () => import('@/views/layout'),
    children: [
      {
        path: '/home', // 默认子路由,只能有一个
        name: 'home',
        component: () => import('@/views/home')
      },
      {
        path: '/visit',
        name: 'visit',
        component: () => import('@/views/visit')
      },
      {
        path: '/message',
        name: 'message',
        component: () => import('@/views/message')
      },
      {
        path: '/mine',
        name: 'mine',
        component: () => import('@/views/mine')
      },
      {
        path: '/post',
        name: 'post',
        component: () => import('@/views/post')
      }
    ]
  },
  {
    path: '/works/:worksId',
    name: 'works',
    component: () => import('@/views/works'),
    props: true // 开启 Props 传参， 说白了就是把路由参数映射到组件的 props 数据中
  }
]

const router = new VueRouter({
  routes
})

export default router
