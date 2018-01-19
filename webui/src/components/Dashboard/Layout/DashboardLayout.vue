<template>
  <div class="wrapper">
    <side-bar type="sidebar" :sidebar-links="$sidebar.sidebarLinks">

    </side-bar>
    <notifications>

    </notifications>
    <div class="main-panel">
      <top-navbar></top-navbar>

      <dashboard-content @click.native="toggleSidebar">
      </dashboard-content>

      <content-footer @click.native="toggleSidebar"></content-footer>
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
    },
    methods: {
      toggleSidebar (e) {
        if (this.$sidebar.showSidebar) {
          this.$sidebar.showSidebar = false
        }
      }
    },
    data () {
      return {}
    }
  }

</script>
