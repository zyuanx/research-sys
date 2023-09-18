import { createRouter, createWebHistory } from 'vue-router'
import Layout from '@/layout/LayoutView.vue'

export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/LoginView.vue'),
    hidden: true
  },

  {
    path: '/',
    name: 'home',
    component: Layout,
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/DashboardView.vue'),
        meta: {
          title: '工作台',
          icon: 'dashboard'
        }
      }
    ]
  },
  {
    path: '/research',
    name: 'research',
    component: () => import('@/views/research/ResearchView.vue')
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (About.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import('../views/AboutView.vue')
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: constantRoutes
})

export default router
