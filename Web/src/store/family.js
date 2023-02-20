// 家模块
import {dspFamilyList} from "@/api";

const state = {
    familyList: []
};
// mutations是唯一修改state的地方
const mutations={
    FAMILYLIST(state,familyList){
        state.familyList = familyList;
    }
};
// actions 用户处理派发actions的地方，可以书写异步语句、自己逻辑的地方
const actions = {
   // 获取家信息
    async getFamilyList({commit}) {
        let result = await dspFamilyList();
        if(result.code === 1000){
            commit("FAMILYLIST",result.data);
            console.log(result)
            return "ok"
        }else {
            return Promise.reject(new Error('faile'));
        }
    }
};
const getters = {
    getFamilyLists(state) {
        return state.familyList;
    }
};


export default {
    state,
    mutations,
    actions,
    getters
}
