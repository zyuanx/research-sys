import { ref } from 'vue'
import { defineStore } from 'pinia'
import { login, getInfo } from '@/api/auth'

export const useAuthStore = defineStore(
  'user-auth',
  () => {
    const token = ref('')
    const userInfo = ref({})
    const userLogin = async (data) => {
      const res = await login(data)
      token.value = res.data.token
      return res.data
    }
    const getUserInfo = async () => {
      const res = await getInfo()
      userInfo.value = res.data
      return res.data
    }

    const resetToken = () => {
      token.value = ''
      userInfo.value = {}
    }
    const userLogout = () => {
      resetToken()
    }
    return { token, userInfo, userLogin, getUserInfo, resetToken, userLogout }
  },
  {
    persist: {
      storage: localStorage
    }
  }
)
