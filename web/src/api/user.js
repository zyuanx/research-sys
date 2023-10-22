import request from '@/request/index'

export function changePassword(data) {
  return request({
    url: '/api/user/change/password',
    method: 'put',
    data
  })
}
