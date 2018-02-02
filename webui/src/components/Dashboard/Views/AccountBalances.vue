<template>
  <div>
    <div class="row">
      <div class="col-md-12">
        <div class="card">
          <div class="content">
            <h4 class="title">{{thisAccount.name}}</h4>
            <paper-table :data="balances" :columns="tblColumns" :columnsStyles="tblStyles" :columnsProperties="tblColumnProperties"/>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
  import AccountTreeView from 'components/UIComponents/AccountTreeView.vue'
  import PaperTable from 'components/UIComponents/PaperTable.vue'
  import numeric from '@/numeric.js'

  const tableColumns = ['Date', 'Balance']
  const tableColumnsStyle = ['', 'mono']
  const tableColumnsProperties = ['date', 'txt']

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
        this.$http.get('books/{book-id}/accounts/' + this.id + '/balance/1600-01-01/9999-12-31').then(response => { // Success
          this.balances = response.body
          // Make data more presentable
          for (var i = 0; i < this.balances.length; i++) {
            var txt = ''
            for (let code in this.balances[i].total_codes) {
              txt += numeric.format(this.balances[i].total_codes[code]) + '\u00A0' + code + '; '
            }
            this.balances[i].txt = txt
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
        balances: [],
        tblColumns: [...tableColumns],
        tblStyles: [...tableColumnsStyle],
        tblColumnProperties: [...tableColumnsProperties]
      }
    }
  }

</script>
