<template>
  <div class="row">
    <div class="col-md-12">
      <div class="card">
        <div class="content">
          <h4 class="title">{{$t('Transactions')}} <router-link to="transactions/add">({{$t('Add')}})</router-link></h4>
          <paper-table :data="transactions" :columns="tblColumns" :columnsStyles="tblStyles" :click_callback="tblClickCallback" :columnsProperties="tblColumnProperties"/>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
  import PaperTable from 'components/UIComponents/PaperTable.vue'
  const tableColumns = ['Name', 'Date', 'Total']
  const tableColumnsStyle = ['', '', 'mono']
  const tableColumnsProperties = ['name', 'local_date', 'total']

  export default {
    components: {
      PaperTable
    },
    beforeMount () {
      this.$store.dispatch('updateAssets')
      this.loadTransactions()
    },
    methods: {
      loadTransactions () {
        if (this.transactions_loading) {
          return
        }
        this.transactions_loading = true
        this.$http.get('books/{book-id}/transactions').then(response => { // Success
          this.transactions = response.body
          for (var i = 0; i < this.transactions.length; i++) {
            var str = ''
            for (var id in this.transactions[i].totals) {
              var amount = this.transactions[i].totals[id] / 1E8
              var code = '?'
              if (this.assetsByID[id] !== undefined) {
                code = this.assetsByID[id].code
              } else {
                console.log('Unknown asset: ' + id)
                continue
              }
              if (str !== '') {
                str += '; '
              }
              str += amount + ' ' + code
            }
            this.transactions[i].total = str
          }
          this.transactions_loading = false
        }, response => { // Error
          console.log('err', response)
          this.transactions_loading = false
        })
      }
    },
    props: [],
    computed: {
      assetsByID () {
        return this.$store.state.assetsByID
      }
    },
    data () {
      return {
        transactions: [],
        transactions_loading: false,
        tblColumns: [...tableColumns],
        tblStyles: [...tableColumnsStyle],
        tblColumnProperties: [...tableColumnsProperties],
        tblClickCallback: undefined
      }
    }
  }

</script>
