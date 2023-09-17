import axios from 'axios'
import { message } from 'ant-design-vue'
const api = axios.create({
  baseURL: 'http://localhost:8080',
  timeout: 3000
})

api.interceptors.request.use(
  (config) => {
    // console.log(config)
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
