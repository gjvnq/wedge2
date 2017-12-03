import VueResource from 'vue-resource'

export default function install (Vue, options) {
  Vue.appData = {}
  Vue.appData.accounts = {}
  Vue.appData.accounts.tree = {}
  Vue.appData.accounts.list = []
  Vue.appData.accounts.loading = false
  Vue.appData.accounts.update = function () {
    // Send request
    Vue.appData.accounts.loading = true
    Vue.http.get('books/{book-id}/accounts').then(response => { // Success
      Vue.appData.accounts.list = response.body
      Vue.appData.accounts.loading = false
    }, response => { // Error
      Vue.appData.accounts.loading = false
    })
    // Send request
    Vue.http.get('books/{book-id}/accounts-tree').then(response => { // Success
      Vue.appData.accounts.tree = response.body
      Vue.appData.accounts.loading = false
    }, response => { // Error
      Vue.appData.accounts.loading = false
    })
  }

  Vue.use(VueResource)

  Vue.prototype.$appData = {}
  Vue.prototype.$appData.accounts = {}
  Vue.prototype.$appData.accounts.tree = function () {
    console.log('hi')
    return Vue.appData.accounts.tree
  }
  Vue.prototype.$appData.accounts.list = function () {
    return Vue.appData.accounts.list
  }
  Vue.prototype.$appData.accounts.first = function () {
    if (Vue.appData.accounts.tree === {} || Vue.appData.accounts.list.length === 0) {
      if (!Vue.appData.accounts.loading) {
        return Vue.appData.accounts.update()
      }
    }
  }
  Vue.prototype.$appData.accounts.update = function () {
    return Vue.appData.accounts.update()
  }
}
