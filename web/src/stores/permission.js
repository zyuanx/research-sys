import { ref, h } from 'vue'
import { defineStore } from 'pinia'
import { constantRoutes } from '@/router'

function getItem(label, key, icon, children, type) {
  return {
    key,
    icon,
    children,
    label,
    type
  }
}

function makeRoute(r) {
  const routes = []
  console.log(r)

  r.forEach((item) => {
    if (item.hidden) {
      return
    }
    const label = item.meta && item.meta.title ? item.meta.title : item.name
    if (item.children?.length === 0) {
      routes.push(getItem(label, item.path, () => h(item.meta.icon), null, null))
    } else {
      routes.push(
        getItem(label, item.path, () => h(item.meta.icon), makeRoute(item.children), 'group')
      )
    }
  })
  return routes
}

export const usePermissionStore = defineStore('permission', () => {
  console.log('constantRoutes', constantRoutes)
  const routes = ref(makeRoute(constantRoutes))
  return { routes }
})
