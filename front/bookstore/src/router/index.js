import Vue from 'vue'
import Router from 'vue-router'
// import HelloWorld from '@/components/HelloWorld'
// import Home from '@/views/Home'
// import Login from '@/views/Login'
// import Regist from '@/views/Regist'
// import Book from '@/views/Book'
// import Store from '@/views/Store'
// import Order from '@/views/Order'
// import Cart from '@/views/Cart'

Vue.use(Router)

const router = new Router({
  mode: 'hash',
  // base: process.env.BASE_URL,
  routes: [{
    path: '/',
    name: 'noPage',
    // component: () => import('@/views/Home'),
    redirect: '/Login'
  },
  {
    path: '/Home',
    name: 'Home',
    component: () => import('@/views/Home'),
    redirect: '/Home/Book',
    children: [{
      path: 'Book',
      name: 'Book',
      component: () => import('@/views/Book')
    }, {
      path: 'Store',
      name: 'Store',
      component: () => import('@/views/Store')
    }, {
      path: 'Order',
      name: 'Order',
      component: () => import('@/views/Order')
    }, {
      path: 'Cart',
      name: 'Cart',
      component: () => import('@/views/Cart')
    }]
  }, {
    path: '/Login',
    name: 'Login',
    component: () => import('@/views/Login')
  }, {
    path: '/Regist',
    name: 'Regist',
    component: () => import('@/views/Regist')
  }, {
    path: '*',
    // name: 'Regist',
    component: () => import('@/views/404')
  }
  ]
})
// vue 菜单路由重复点击报错
// https://blog.csdn.net/qq_40190624/article/details/102588258
const originalPush = Router.prototype.push
Router.prototype.push = function push (location) {
  return originalPush.call(this, location).catch(err => err)
}

// vue3.0 全局路由守卫(router.js设置)
// https://blog.csdn.net/qq_40190624/article/details/89055679
// 路由守卫
router.beforeEach((to, from, next) => {
  // console.log('to', to)
  // console.log('from', from)
  // localStorage.getItem('iview_login')
  const isLogin = !!localStorage.iview_login
  // console.log(localStorage.getItem('iview_login'))
  // console.log(isLogin)
  if (to.path === '/Login') {
    next()
  } else {
    // 是否在登录状态下
    isLogin ? next() : next('/Login')
  }
  // next()
})
export default router
