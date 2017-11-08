import Sidebar from './SideBar.vue'

const SidebarStore = {
  showSidebar: false,
  sidebarLinks: [
    {
      name: 'Summary',
      icon: 'ti-panel',
      path: '/book/summary'
    },
    {
      name: 'Accounts',
      icon: 'ti-wallet',
      path: '/book/accounts'
    },
    {
      name: 'Currencies & Assets',
      icon: 'ti-gift',
      path: '/book/assets'
    },
    {
      name: 'Transactions',
      icon: 'ti-shopping-cart',
      path: '/book/transactions'
    }
  ]
}

const SidebarPlugin = {

  install (Vue) {
    Vue.mixin({
      data () {
        return {
          sidebarStore: SidebarStore
        }
      }
    })

    Object.defineProperty(Vue.prototype, '$sidebar', {
      get () {
        return this.$root.sidebarStore
      }
    })
    Vue.component('side-bar', Sidebar)
  }
}

export default SidebarPlugin
