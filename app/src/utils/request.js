/**
 * 请求模块
 */
import axios from 'axios'
import store from '@/store'
import JSONBig from 'json-bigint'

// JSONBig 可以处理数据中超过 JavaScript 安全整数范围的问题
// JSONBig.parse()  // 把 JSON 格式的字符串转为 JavaScript 对象
// JSONBig。stringify() // 把 JavaScript 对象转换为JSON 格式的字符串

const request = axios.create({
  baseURL: '/api/v1', // 接口的基准路径
  // 自定义后端返回的原始数据
  // data: 后端返回的原始数据，说白了就是 JSON 格式的字符串
  transformResponse: [function (data) {
    try {
      return JSONBig.parse(data)
    } catch (err) {
      return data
    }
  // axios 默认会在内部这样来处理后端返回的数据
  // return JSON.parse(data)
  }]
})

// 请求拦截器
// Add a request interceptor
request.interceptors.request.use(function (config) {
  // 请求发起会经过这里
  // config: 本次请求的请求配置对象
  const { user } = store.state
  if (user && user.token) {
    config.headers.Authorization = `Bearer ${user.token}`
  }

  // 注意，这里务必返回 config 配置对象， 否则请求九停在这里出不去
  return config
}, function (error) {
  // 如果请求出错了（还没有发出去） 会进入这里.
  return Promise.reject(error)
})

// 响应拦截器

export default request
