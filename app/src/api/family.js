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

/**
 * 获取我的家列表
 */
export const getMyFamilyList = userId => {
  return request({
    method: 'GET',
    url: `/family/myFamily/${userId}`
  })
}
