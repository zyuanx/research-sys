import request from '@/request/index.js'

export const ping = (params) => request.get(`/ping`, { params })
export function login(data) {
  return request({
    url: '/api/user/login',
    method: 'post',
    data
  })
}

export function getInfo() {
  return request({
    url: '/api/user/info',
    method: 'get'
  })
}
