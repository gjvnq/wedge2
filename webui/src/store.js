import Vue from 'vue'
import Vuex from 'vuex'
import VueResource from 'vue-resource'

Vue.use(Vuex)
Vue.use(VueResource)

export default new Vuex.Store({
  strict: true,
  state: {
    assets: [],
    assetsByID: {},
    assetsLoading: false,
    accounts: [],
    accountsMap: {},
    accountsTree: {},
    accountsLoading: false
  },
  mutations: {
    setAssets (state, newAssets) {
      state.assets = newAssets
      state.assetsByID = {}
      for (var i = 0; i < newAssets.length; i++) {
        var asset = newAssets[i]
        state.assetsByID[asset.id] = asset
      }
    },
    setAssetsLoading (state, flag) {
      state.assetsLoading = flag
    },
    setAccounts (state, newAccounts) {
      state.accounts = newAccounts
      state.accountsMap = {}
      for (var i = 0; i < newAccounts.length; i++) {
        var account = newAccounts[i]
        state.accountsMap[account.id] = account
      }
    },
    setAccountsTree (state, newAccountsTree) {
      state.accountsTree = newAccountsTree
    },
    setAccountsLoading (state, flag) {
      state.accountsLoading = flag
    }
  },
  actions: {
    updateAssets ({ commit, state }) {
      if (state.assetsLoading) {
        return
      }
      // Send request
      commit('setAssetsLoading', true)
      Vue.http.get('books/{book-id}/assets').then(response => { // Success
        commit('setAssets', response.body)
        commit('setAssetsLoading', false)
      }, response => { // Error
        commit('setAssetsLoading', false)
      })
    },
    updateAccounts ({ commit, state }) {
      if (state.accountsLoading) {
        return
      }
      // Send request
      commit('setAccountsLoading', true)
      Vue.http.get('books/{book-id}/accounts').then(response => { // Success
        commit('setAccounts', response.body)
        commit('setAccountsLoading', false)
      }, response => { // Error
        commit('setAccountsLoading', false)
      })
      // Send request
      Vue.http.get('books/{book-id}/accounts-tree').then(response => { // Success
        commit('setAccountsTree', response.body)
        commit('setAccountsLoading', false)
      }, response => { // Error
        commit('setAccountsLoading', false)
      })
    }
  }
})
