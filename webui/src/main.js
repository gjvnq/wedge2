import Vue from 'vue'
import Vuex from 'vuex'
import VueI18n from 'vue-i18n'
import VueRouter from 'vue-router'
import VueResource from 'vue-resource'
import VueShortkey from 'vue-shortkey'

// Plugins
import storePlugin from './storePlugin'
import GlobalComponents from './globalComponents'
import GlobalDirectives from './globalDirectives'
import Notifications from './components/UIComponents/NotificationPlugin'
import SideBar from './components/UIComponents/SidebarPlugin'
import App from './App'

// router setup
import routes from './routes/routes'

// library imports
import Chartist from 'chartist'
import 'bootstrap/dist/css/bootstrap.css'
import './assets/sass/paper-dashboard.scss'
import 'es6-promise/auto'

// load translator
import {messages} from './i18n.js'

// plugin setup
Vue.use(Vuex)
Vue.use(storePlugin)
Vue.use(VueShortkey)
Vue.use(VueI18n)
Vue.use(VueRouter)
Vue.use(VueResource)
Vue.use(GlobalComponents)
Vue.use(GlobalDirectives)
Vue.use(Notifications)
Vue.use(SideBar)

// configure router
const router = new VueRouter({
  routes, // short for routes: routes
  linkActiveClass: 'active',
  scrollBehavior (to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { x: 0, y: 0 }
    }
  }
})

// configure i18n
const i18n = new VueI18n({
  locale: 'pt',
  fallbackLocale: 'en',
  messages
})

// global library setup
Object.defineProperty(Vue.prototype, '$Chartist', {
  get () {
    return this.$root.Chartist
  }
})

/* eslint-disable no-new */
Vue.http.options.root = window.location.origin + '/wedge/api/'
if (window.location.hostname === 'localhost') {
  Vue.http.options.root = '//localhost:8081/'
}
Vue.http.options.book_id = 'no-book'
Vue.http.options.emulateJSON = false
Vue.http.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('wedge-Authorization')
Vue.http.options.book_id = localStorage.getItem('wedge-BookId')

var first401 = false

Vue.http.interceptors.push(function (request, next) {
  request.url = request.url.replace('{book-id}', Vue.http.options.book_id)
  next(function (response) {
    if (response.status === 401 && first401 === false) {
      first401 = true
      TheVue.$notifications.notify(
        {
          message: TheVue.$t('You have been logged out. Going back to the login page in 5 seconds...'),
          icon: 'ti-alert',
          horizontalAlign: 'center',
          verticalAlign: 'top',
          type: 'danger'
        })
      setTimeout(function () {
        TheVue.$router.push('/login')
        first401 = false
      }, 5000)
    }
  })
})

var TheVue = new Vue({
  el: '#app',
  render: h => h(App),
  router,
  i18n,
  data: {
    Chartist: Chartist
  }
})
