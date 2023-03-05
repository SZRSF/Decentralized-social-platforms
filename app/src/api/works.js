/**
 * 文章请求模块
 */

import request from '@/utils/request'
/**
 * 请求获取文章列表
 */
export const getWorksList = params => {
  return request({
    method: 'GET',
    url: '/posts/',
    params
  })
}

/**
 * 请求发布文章
 */
export const postWorks = params => {
  return request({
    method: 'POST',
    url: '/post',
    data: params
  })
}

/**
 * 图片上传
 */
export const postImage = params => {
  return request(
    {
      method: 'POST',
      url: '/post/image',
      data: params
    })
}

/**
 * 根据id获取文章
 */
export const getWorksById = id => {
  return request({
    method: 'GET',
    url: `/post/${id}`
  })
}
