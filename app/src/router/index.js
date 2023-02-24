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
      }
    ]
  }
]

const router = new VueRouter({
  routes
})

export default router
