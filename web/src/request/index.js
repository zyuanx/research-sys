import axios from 'axios'
import { message } from 'ant-design-vue'
import { userAuthStore } from '@/stores/auth'
const api = axios.create({
  baseURL: 'http://localhost:8080',
  timeout: 3000
})
const auth = userAuthStore()
api.interceptors.request.use(
  (config) => {
    // console.log(config)
    if (auth.token) {
      config.headers.Authorization = `Bearer ${auth.token}`
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
    message.error(res.message)
    Promise.reject(err)
  }
)

export default api
