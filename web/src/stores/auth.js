import { ref } from 'vue'
import { defineStore } from 'pinia'
import request from '@/request/index.js'

export const userAuthStore = defineStore('user-auth', () => {
  const token = ref('')
  const userInfo = ref({})
  const getUserInfo = async () => {
    const res = await request.get('/user/info')
    userInfo.value = res.data
    return res.data
  }

  const login = async (data) => {
    const res = await request.post('/login', data)
    token.value = res.data.token
    localStorage.setItem('token', res.data.token)
    return res.data
  }
  const resetToken = () => {
    token.value = ''
    localStorage.removeItem('token')
  }
  return { token, userInfo, getUserInfo, login, resetToken }
})
