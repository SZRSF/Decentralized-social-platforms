// // 作品模块
// import {dspFamilyList, dspWorksList} from "@/api";
//
// // state:仓库存储数据的地方
// const state ={
//     worksList:[]
// };
// // mutations: 修改state的唯一手段
// const mutations = {
//     WORKSLIST(state,worksList){
//         state.worksList = worksList
//     }
// };
// // action: 处理action,可以书写自己的业务逻辑
// const actions = {
//     // 获取作品信息
//     async getWorksList({commit}) {
//         let result = await dspWorksList();
//         if(result.code === 1000){
//             commit(" WORKSLIST",result.data);
//             console.log(result)
//             return "ok"
//         }else {
//             return Promise.reject(new Error('faile'));
//         }
//     }
// };
// // getters: 理解为计算属性，用于简化仓库数据， 让组件获取仓库的数据更加方便
// const getters = {
//     getWorksList({commit}) {
//         return commit.worksList;
//     }
// };
//
// export default {
//     state,
//         mutations,
//         actions,
//         getters
// }
