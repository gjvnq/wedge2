<template>
  <div class="wrapper">
    <side-bar type="sidebar" :sidebar-links="$sidebar.sidebarLinks">

    </side-bar>
    <notifications>

    </notifications>
    <div class="main-panel">
      <top-navbar></top-navbar>

      <dashboard-content :globoids="globoids" @click.native="toggleSidebar">

      </dashboard-content>

      <content-footer></content-footer>
    </div>
  </div>
</template>
<style lang="scss">

</style>
<script>
  import TopNavbar from './TopNavbar.vue'
  import ContentFooter from './ContentFooter.vue'
  import DashboardContent from './Content.vue'
  export default {
    components: {
      TopNavbar,
      ContentFooter,
      DashboardContent
    },
    beforeMount () {
      this.updateAccounts()
      this.updateAssets()
    },
    methods: {
      toggleSidebar () {
        if (this.$sidebar.showSidebar) {
          this.$sidebar.displaySidebar(false)
        }
      },
      updateAccounts () {
        // Send request
        this.globoids.accountsLoading = true
        this.$http.get('books/{book-id}/accounts').then(response => { // Success
          this.globoids.accountsList = response.body
          this.globoids.accountsLoading = false
        }, response => { // Error
          this.globoids.accountsLoading = false
        })
        // Send request
        this.$http.get('books/{book-id}/accounts-tree').then(response => { // Success
          this.globoids.accountsTree = response.body
          this.globoids.accountsLoading = false
        }, response => { // Error
          this.globoids.accountsLoading = false
        })
      },
      updateAssets () {
        // Send request
        this.globoids.assetsLoading = true
        this.$http.get('books/{book-id}/assets').then(response => { // Success
          this.globoids.assetsList = response.body
          this.globoids.assetsLoading = false
        }, response => { // Error
          this.globoids.assetsLoading = false
        })
      }
    },
    data () {
      return {
        globoids: {
          accountsLoading: false,
          accountsList: [1, 2, 3],
          accountsTree: {}
        }
      }
    }
  }

</script>
