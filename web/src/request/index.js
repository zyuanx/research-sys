import axios from 'axios'
import { useAuthStore } from '@/stores/auth'
const api = axios.create({
  baseURL: 'http://localhost:8080',
  timeout: 3000
})
api.interceptors.request.use(
  (config) => {
    const auth = useAuthStore()
    const token = auth.token
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (err) => {
    Promise.reject(err)
  }
)

api.interceptors.response.use(
  (res) => {
    console.log(res)
    return Promise.resolve(res.data)
  },
  (err) => {
    const res = err.response.data
    console.log(res)
    Promise.reject(err)
  }
)

export default api
