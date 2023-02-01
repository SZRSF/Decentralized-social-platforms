//配置路由的地方
import Vue, {nextTick} from "vue";
import VueRouter from "vue-router";
//使用插件
Vue.use(VueRouter);
//引入路由组件
import Home from '@/views/Home'
import Serach from "@/views/Serach";
import Login from "@/views/Login";
import Register from "@/views/Register";
import User from "@/views/User";
import Works from "@/views/Works";

//配置路由
const router = new VueRouter({
    //配置路由
    routes:[
        {
            path:"/home",
            component:Home,
            meta:{show:true}
        }
        ,
        {
            path:"/search",
            component:Serach,
            meta:{show:true},
            name:"search"
        }
        ,
        {
            path:"/login",
            component:Login,
            meta:{show:false}
        }
        ,
        {
            path:"/register",
            component:Register,
            meta:{show:false}
        }
        ,
        {
            path:"/user",
            component:User,
            meta:{show:true}
        }
        ,
        {
            path: "/works",
            component:Works,
            meta:{show:true}
        },
        //重定向,在项目跑起来的时候，访问/，立马让他定向到首页
        {
            path:'*',
            redirect:'/login'
        }
    ]
})

// 挂载路由导航守卫
router.beforeEach((to, from, next) => {
    //to 将要访问的路径
    //from 代表从哪个路径跳转而来
    //next 是一个函数，表示放行
    // next()放行  next('/login‘）强制跳转

    if(to.path === '/login') return next();
    // 获取token
    const tokenStr = window.sessionStorage.getItem('token')
    if (!tokenStr) return next('/login')
    next()
})

export default router
