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
