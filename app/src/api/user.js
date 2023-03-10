/**
 *  用户相关请求模块
 */

import request from '@/utils/request'

export const login = data => {
  return request({
    method: 'POST',
    url: '/login',
    data
  })
}

/**
 * 获取用户自己的信息
 */
export const getUserInfo = id => {
  return request({
    method: 'GET',
    url: `/user/${id}`
    // 发送请求头数据
    // headers: {
    //   // 该接口需要授权才能访问
    //   Authorization: `Bearer ${store.state.user.token}`
    // }
  })
}

/**
 * 关注用户
 */
export const addFollow = target => {
  return request({
    method: 'POST',
    url: '/user/followings',
    data: target
  })
}

/**
 * 取消关注用户
 */
export const deleteFollow = target => {
  return request({
    method: 'DELETE',
    url: `/user/followings/${target}`
  })
}
