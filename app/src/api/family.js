import request from '@/utils/request'

/**
 * 获取家列表
 */
export const getFamilyList = params => {
  return request({
    method: 'GET',
    url: '/family',
    params
  })
}
