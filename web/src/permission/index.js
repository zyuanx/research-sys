import router from '@/router'
import { userAuthStore } from '@/stores/auth'

const whiteList = ['/login']

router.beforeEach(async (to, from, next) => {
  console.log('to', to)
  console.log('from', from)
  console.log('next', next)
  document.title = to.meta.title
  const auth = userAuthStore()
  console.log('auth', auth)
  const token = localStorage.getItem('token')
  if (token) {
    if (to.path === '/login') {
      next({ path: '/' })
    } else {
      // 获取用户角色
    }
  } else {
    let pass = false
    for (let i = 0; i < whiteList.length; i++) {
      if (to.path.search(whiteList[i]) !== -1) {
        pass = true
        break
      }
    }
    if (pass) {
      next()
    } else {
      next(`/login?redirect=${to.path}`)
    }
  }
  next()
})

// function hasPermission(roles, route) {
//   if (route.meta && route.meta.roles) {
//     return roles.some((role) => route.meta.roles.includes(role))
//   } else {
//     return true
//   }
// }
