import Vue from "vue";
import Vuex from "vuex"
// 需要使用插件一次
Vue.use(Vuex);

// 引入模块仓库
import user from "@/store/user";
import family from "@/store/family";
// import works from "@/store/works";

// 对外暴露Store一个实例
export default new Vuex.Store({
    // 模块：把小仓库进行合并成为大仓库
    modules: {
        user,
        family,
        // works
    },
});



// state:仓库存储数据的地方
// const state ={};
// mutations: 修改state的唯一手段
// const mutations = {};
// action: 处理action,可以书写自己的业务逻辑
// const actions = {};
// getters: 理解为计算属性，用于简化仓库数据， 让组件获取仓库的数据更加方便
// const getters = {};
