import { createRouter, createWebHistory } from 'vue-router'
import Layout from '@/layout/LayoutView.vue'

export const constantRoutes = [
  {
    path: '/login',
    name: 'login',
    meta: {
      title: '登录'
    },
    component: () => import('@/views/LoginView.vue'),
    hidden: true
  },
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [
      {
        path: '/dashboard',
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
    component: Layout,
    children: [
      {
        path: 'list',
        name: 'researchList',
        component: () => import('@/views/research/ResearchListView.vue')
      },
      {
        path: 'detail',
        name: 'researchDetail'
      },
      {
        path: 'create',
        name: 'researchCreate',
        meta: {
          title: '创建调研'
        },
        component: () => import('@/views/research/ResearchCreateView.vue')
      }
    ]
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

export const dynamicRoutes = [
  {
    path: '/admin',
    name: 'admin',
    component: Layout,
    meta: {
      roles: ['admin']
    }
  },
  {
    path: '/user',
    name: 'user',
    component: Layout,
    meta: {
      roles: ['user']
    }
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: constantRoutes
})

export default router
