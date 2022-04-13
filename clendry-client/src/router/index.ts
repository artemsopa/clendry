import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'

//Adding layouts router.
const BlankLayout = () => import("@/layouts/BlankLayout.vue")
const Layout = () => import("@/layouts/Layout.vue")

//main pages
const Files = () => import('@/views/clendry/Files.vue')
const Pages = () => import('@/views/pages/Pages.vue')

//auth elements
const SignIn = () => import('@/views/auth/SignIn.vue')

const childRoute = () => [
  {
    path: '',
    name: 'layout.pages',
    meta: {  name: 'Pages' },
    component: Pages
  },
  {
    path: 'files',
    name: 'layout.files',
    meta: {  name: 'Files' },
    component: Files
  }
]

const authchildRoute = () =>[
  {
    path: 'sign-in',
    name: 'auth.login',
    meta: {  name: 'SignIn' },
    component: SignIn
  },
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
