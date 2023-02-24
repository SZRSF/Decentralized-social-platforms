import request from '@/utils/request'

/**
 * 获取家列表
 */
export const getFamilyList = () => {
  return request({
    method: 'GET',
    url: '/family'
  })
}
