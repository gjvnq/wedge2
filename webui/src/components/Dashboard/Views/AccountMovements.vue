<template>
  <div>
    <account-tree-view :model="thisAccount" :level="0" :single-card="true"></account-tree-view>
    <div class="row">
      <div class="col-md-12">
        <div class="card">
          <div class="content">
            <h4 class="title">{{$t('Movements')}}</h4>
            <paper-table :data="movements" :columns="tblColumns" :columnsStyles="tblStyles" :click_callback="tblClickCallback" :columnsProperties="tblColumnProperties"/>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
  import AccountTreeView from 'components/UIComponents/AccountTreeView.vue'
  import PaperTable from 'components/UIComponents/PaperTable.vue'
  const tableColumns = ['Transaction Name', 'Transaction Date', 'Movement Amount', 'Movement Date', 'Movement Status']
  const tableColumnsStyle = ['', '', 'mono', '', '']
  const tableColumnsProperties = ['name', 'local_date', 'total']

  export default {
    components: {
      AccountTreeView,
      PaperTable
    },
    beforeMount () {
      this.start(this.$route.params.acc_id)
    },
    watch: {
      '$route' (to, from) {
        this.start(to.params.acc_id)
      }
    },
    methods: {
      start (id) {
        this.id = id
        this.$store.dispatch('updateAssets')
        this.$store.dispatch('updateAccounts')
      }
    },
    computed: {
      thisAccount () {
        var account = this.$store.state.accountsMap[this.id]
        if (account === undefined) {
          account = {}
        }
        return account
      }
    },
    data () {
      return {
        newAccountName: '',
        newAccountParent: '00000000-0000-0000-0000-000000000000',
        newAccountBtn: true,
        id: {},
        tblColumns: [...tableColumns],
        tblStyles: [...tableColumnsStyle],
        tblColumnProperties: [...tableColumnsProperties],
        tblClickCallback: function (el) {
          console.log(el.id)
          this.$router.push('transactions/' + el.id)
        }
      }
    }
  }

</script>
