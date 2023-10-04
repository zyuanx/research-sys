import { ref } from 'vue'
import { defineStore } from 'pinia'
import { constantRoutes, dynamicRoutes } from '@/router'
function makeDynamicRoute(r, roles) {
  const routes = []

  r.forEach((item) => {
    if (item.hidden) {
      return
    }
    let flag = false
    if (item.meta && item.meta.roles) {
      if (roles.some((role) => item.meta.roles.includes(role))) {
        flag = true
      }
    }
    if (!flag) {
      return
    }
    if (item?.children?.length || 0 === 0) {
      routes.push(item)
    } else {
      const newItem = { ...item }
      newItem.children = makeDynamicRoute(item.children, roles)
      routes.push(newItem)
    }
  })
  return routes
}

export const usePermissionStore = defineStore(
  'user-permission',
  () => {
    const router = ref([])
    const test = ref('test')
    router.value = constantRoutes

    function getRoute(roles) {
      const dyRouter = makeDynamicRoute(dynamicRoutes, roles)
      router.value.push(...dyRouter)
      return dyRouter
    }
    return { router, test, getRoute }
  },
  {
    persist: {
      storage: localStorage
    }
  }
)
