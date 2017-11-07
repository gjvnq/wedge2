import Vue from 'vue'
import VueI18n from 'vue-i18n'
import VueRouter from 'vue-router'
import VueResource from 'vue-resource'

// Plugins
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
  linkActiveClass: 'active'
})

// configure i18n
const i18n = new VueI18n({
  locale: 'pt', // set locale
  messages // set locale messages
})

// global library setup
Object.defineProperty(Vue.prototype, '$Chartist', {
  get () {
    return this.$root.Chartist
  }
})

/* eslint-disable no-new */
Vue.http.options.root = '/api'
Vue.http.options.root = '//localhost:8081/'
Vue.http.options.emulateJSON = false
Vue.http.headers.common['Authorization'] = 'Bearer'
new Vue({
  el: '#app',
  render: h => h(App),
  router,
  i18n,
  data: {
    Chartist: Chartist
  }
})
