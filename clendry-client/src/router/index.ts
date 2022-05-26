import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'

//Adding layouts router.
const BlankLayout = () => import("@/layouts/BlankLayout.vue")
const Layout = () => import("@/layouts/Layout.vue")

//main pages
const Dashboard = () => import('@/views/Dashboard.vue')

//drive pages
const UploadFile = () => import('@/views/drive/UploadFile.vue')

const Recent = () => import('@/views/drive/Recent.vue')
const Folder = () => import('@/views/drive/Folder.vue')
const Images = () => import('@/views/drive/Images.vue')
const Media = () => import('@/views/drive/Media.vue')
const Files = () => import('@/views/drive/Files.vue')
const Favourites = () => import('@/views/drive/Favourites.vue')
const Trash = () => import('@/views/drive/Trash.vue')

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
  },
  {
    path: 'folder/:id',
    name: 'layout.folder',
    meta: {  name: 'Folder' },
    component: Folder
  },
  {
    path: 'images',
    name: 'layout.images',
    meta: {  name: 'Images' },
    component: Images
  },
  {
    path: 'media',
    name: 'layout.media',
    meta: {  name: 'Media' },
    component: Media
  },
  {
    path: 'files',
    name: 'layout.files',
    meta: {  name: 'Files' },
    component: Files
  },
  {
    path: 'favourites',
    name: 'layout.favourites',
    meta: {  name: 'Favourites' },
    component: Favourites
  },
  {
    path: 'trash',
    name: 'layout.trash',
    meta: {  name: 'Trash' },
    component: Trash
  },
  {
    path: 'upload-file',
    name: 'layout.upload-file',
    meta: {  name: 'UploadFile' },
    component: UploadFile
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
