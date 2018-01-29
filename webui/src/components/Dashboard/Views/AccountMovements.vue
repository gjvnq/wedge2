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
  const tableColumns = ['Transaction Name', 'Transaction Date', 'Amount', 'Movement Date', 'Movement Status']
  const tableColumnsStyle = ['', '', 'mono', '', '']
  const tableColumnsProperties = ['transaction_name', 'transaction_date', 'amount_human', 'local_date', 'status_text']

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
        this.load()
      },
      load () {
        this.$http.get('books/{book-id}/accounts/' + this.id + '/movements').then(response => { // Success
          this.movements = response.body
          // Make data more presentable
          var statusFixer = {
            'C': this.$t('Cancelled'),
            'D': this.$t('Done'),
            'P': this.$t('Planned')
          }
          for (var i = 0; i < this.movements.length; i++) {
            this.movements[i].status_text = statusFixer[this.movements[i].status]
            this.movements[i].amount_human = this.movements[i].amount / 1E8 + ' ' + this.movements[i].asset_code
          }
        }, response => { // Error
          console.log('ERR', response)
        })
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
        movements: [],
        tblColumns: [...tableColumns],
        tblStyles: [...tableColumnsStyle],
        tblColumnProperties: [...tableColumnsProperties],
        tblClickCallback: function (el) {
          this.$router.push('../transactions/' + el.transaction_id)
        }
      }
    }
  }

</script>
