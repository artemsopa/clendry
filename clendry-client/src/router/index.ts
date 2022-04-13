import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'

//Adding layouts router.
const BlankLayout = () => import("@/layouts/BlankLayout.vue")
const Layout = () => import("@/layouts/Layout.vue")

//main pages
const Dashboard = () => import('@/views/Dashboard.vue')
const Recent = () => import('@/views/drive/Recent.vue')

//auth elements
const SignIn = () => import('@/views/auth/SignIn.vue')
const SignUp = () => import('@/views/auth/SignUp.vue')

const childRoute = () => [
  {
    path: '',
    name: 'layout.dashboard',
    meta: {  name: 'Dashboard' },
    component: Dashboard
  },
  {
    path: 'recent',
    name: 'layout.recent',
    meta: {  name: 'Recent' },
    component: Recent
  }
]

const authchildRoute = () =>[
  {
    path: 'sign-in',
    name: 'auth.login',
    meta: {  name: 'SignIn' },
    component: SignIn
  },
  {
    path: 'sign-up',
    name: 'auth.register',
    meta: {  name: 'SignUp' },
    component: SignUp
  }
]

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: '',
    component: Layout,
    children: childRoute()
  },
  {
    path: '/auth',
    name: 'auth',
    component: BlankLayout,
    children: authchildRoute()
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
