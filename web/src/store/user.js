import {dspUserLogin} from "@/api";
// 登录与注册模块
const state = {
    code: "",
    token:'',
    user:{
        "username": "",
        "password": "",
        "phone_num": "",
        "emil": "",
        "gender": "",
        "head_img": "",
        "create_time": "",
        "user_id": 0,
        "invite_id": 0,
        "time_stamp": 0
    }
};
// mutations是唯一修改state的地方
const mutations={
    USERLOGIN(state,data){
        state.token = "Bearer " + data.token;
        state.user = data.user
    }
};
// actions 用户处理派发actions的地方，可以书写异步语句、自己逻辑的地方
const actions = {
    //登录业务
    async userLogin({ commit }, data) {
        let result = await dspUserLogin(data);
        //服务器下发token, 用户唯一标识符
        //将来通过带 token 找服务器要用户信息进行展示
        if(result.code === 1000){
            commit("USERLOGIN",result.data);
            return 'ok'
        }else {
            return Promise.reject(new Error('faile'));
        }
    },
};
const getters = {
    donUsers (state) {
        return state.user
    }
};


export default {
    state,
    mutations,
    actions,
    getters
}
