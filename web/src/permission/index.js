import router from '@/router'
import { useAuthStore } from '@/stores/auth'
import { usePermissionStore } from '@/stores/permission'

const whiteList = ['/login']

router.beforeEach(async (to, from, next) => {
  console.log('to', to)
  console.log('from', from)
  document.title = to.meta.title
  const auth = useAuthStore()
  const token = auth.token
  console.log('token', token)
  if (token) {
    if (to.path === '/login1') {
      next({ path: '/' })
    } else {
      const roles = auth.userInfo.roles
      const hasRoles = roles && roles.length > 0
      if (hasRoles) {
        next()
      } else {
        try {
          const userInfo = await auth.getUserInfo()
          const roles = userInfo.roles.map((item) => item.title)
          const permission = usePermissionStore()
          const accessRoutes = await permission.getRoute(roles)
          router.addRoute(accessRoutes)
          next({ ...to, replace: true })
        } catch (error) {
          console.log('error', error)
          await auth.resetToken()
          next(`/login?redirect=${to.path}`)
        }
      }
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
})

// function hasPermission(roles, route) {
//   if (route.meta && route.meta.roles) {
//     return roles.some((role) => route.meta.roles.includes(role))
//   } else {
//     return true
//   }
// }
