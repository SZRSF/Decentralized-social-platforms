import Vue from 'vue'
import Vuex from 'vuex'
import { getItem, setItem } from '@/utils/storage'

Vue.use(Vuex)
const TOKEN_KEY = 'DSP_USER'

export default new Vuex.Store({
  state: {
    // 一个对象， 存储当前登录用户信息（token等数据）
    user: getItem(TOKEN_KEY)
  },
  mutations: {
    setUser (state, data) {
      state.user = data

      // 为了防止刷新， 我们需要将数据备份到本地存储
      setItem(TOKEN_KEY, state.user)
    }
  },
  actions: {
  },
  modules: {
  },
  getters: {
  }
})
