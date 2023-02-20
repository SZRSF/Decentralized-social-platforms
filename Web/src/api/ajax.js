// 对于axios进行二次封装
import axios from "axios";
import store from "@/store";

// 1:利用axios对象的方法create, 去创建一个axios实例
// 2:request就是axios， 只不过稍微配置一下
const requests = axios.create({
    // 配置对象
    // 基础路径, 发请求的时候，路径会出现api
    baseURL: "/api/v1",
    // 请求超时的时间5s
    timeout: 5000,
});

//请求拦截器: 在发请求之前， 请求拦截器可以检测到， 可以在请求发出之前做一些事情
requests.interceptors.request.use((config) => {
    // config：配置对象， 对象里面有一个属性很摘要， headers请求头


    //需要携带token带给服务器
    if(store.state.user.token){
        config.headers.Authorization = store.state.user.token;
    }
    return config;
})

// 相应拦截器
requests.interceptors.response.use((res) => {
    // 成功地回调函数：服务器相应数据回来以后，响应拦截器可以检测到， 可以做一些事情
    return res.data;
}, (error)=> {
    //响应失败的回调函数
    return Promise.reject(new error('faile'));
});

//对外暴露
export default requests;
