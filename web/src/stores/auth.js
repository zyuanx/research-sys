import { ref } from 'vue'
import { defineStore } from 'pinia'

export const userAuthStore = defineStore('user-auth', () => {
  const token = ref('')

  return { token }
})
