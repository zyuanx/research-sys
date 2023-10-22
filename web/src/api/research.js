import request from '@/request/index'

export function createResearch(data) {
  return request({
    url: '/api/research',
    method: 'post',
    data
  })
}
