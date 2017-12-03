<template>
  <div class="row">
    <div class="col-md-12">
      <div class="card">
        <div class="header">
          <h4 class="title">{{$t('Add Transaction')}}</h4>
        </div>
        <div class="content">
          <form @submit.prevent>
            <div class="row">
              <div class="col-md-1"></div>
              <div class="col-md-5">
                <div class="form-group">
                  <label>{{$t('Name')}}</label>
                  <input type="text" class="form-control border-input" v-model="transactionName">
                </div>
              </div>
              <div class="col-md-5">
                <div class="form-group">
                  <label>{{$t('Date')}}</label>
                  <input type="date" class="form-control border-input" v-model="transactionDate">
                </div>
              </div>
            </div>
            <hr>
            <h4 class="title">{{$t('Movements')}}</h4>
            <div class="row">
              <div class="col-md-1"></div>
              <div class="col-md-2">
                <div class="form-group">
                  <label>{{$t('Account')}}</label>
                  <input type="text" class="form-control border-input" v-model="movementAccount">
                </div>
              </div>
              <div class="col-md-2">
                <div class="form-group">
                  <label>{{$t('Value')}}</label>
                  <input type="text" class="form-control border-input" v-model="movementValue">
                </div>
              </div>
              <div class="col-md-2">
                <div class="form-group">
                  <label>{{$t('Currency or Asset')}}</label>
                  <input type="text" class="form-control border-input" v-model="movementAsset">
                </div>
              </div>
              <div class="col-md-2">
                <div class="form-group">
                  <label>{{$t('Date')}}</label>
                  <input type="date" class="form-control border-input" v-model="movementDate">
                </div>
              </div>
              <div class="col-md-2">
                <div class="form-group">
                  <label>{{$t('Status')}}</label>
                  <select class="form-control border-input" v-model="movementStatus">
                    <option value="P">{{$t('Planned')}}</option>
                    <option value="D">{{$t('Done')}}</option>
                    <option value="C">{{$t('Cancelled')}}</option>
                  </select>
                </div>
              </div>
            </div>

            <hr>
            <h4 class="title">{{$t('Items')}}</h4>
            <div class="text-center">
              <button class="btn btn-info btn-fill btn-wd" :disabled="transactionBtn == false" @click="addAccount">{{$t('Add Account')}}</button>
            </div>
            <div class="clearfix">
            </div>
          </form>
        </div>
      </div>
    </div>
    <div class="col-md-12">
      <div class="card">
        {{hack}}
      </div>
    </div>
  </div>
</template>
<script>
  import PaperTable from 'components/UIComponents/PaperTable.vue'
  import TreeView from 'components/UIComponents/TreeView.vue'
  const tableColumns = ['Code', 'Name', 'Last Value', '']
  const tableData = []

  export default {
    components: {
      PaperTable,
      TreeView
    },
    beforeMount () {
      this.$appData.accounts.first()
    },
    methods: {
      addAccount () {
        if (this.transactionBtn === false) {
          return
        }
        this.transactionBtn = false
        // Data
        var fd = {}
        fd['name'] = this.transactionName
        fd['parent_id'] = this.transactionParent
        console.log(fd)
        // Send request
        this.$http.put('books/{book-id}/accounts', fd).then(response => { // Success
          this.transactionBtn = true
          window.book_id = fd['book_id']
          this.transactionName = ''
          this.$appData.accounts.update()
        }, response => { // Error
          console.log('err', response)
          this.transactionBtn = true
          alert(response.bodyText)
          this.$appData.accounts.update()
        })
      },
      updateAccounts () {
        // Send request
        this.$http.get('books/{book-id}/accounts').then(response => { // Success
          this.accountsList = response.body
        }, response => { // Error
          console.log('err', response)
        })
        // Send request
        this.$http.get('books/{book-id}/accounts-tree').then(response => { // Success
          this.accountsTree = response.body
        }, response => { // Error
          console.log('err', response)
        })
      }
    },
    data () {
      return {
        accountsList: [],
        transactionName: '',
        transactionDate: '',
        transactionBtn: true,
        movementAccount: '00000000-0000-0000-0000-000000000000',
        movementAsset: '00000000-0000-0000-0000-000000000000',
        movementDate: '',
        movementStatus: '',
        movementValue: 0,
        accountsTree: {},
        hack: this.$appData.accounts.tree,
        tblAccounts: {
          title: 'Accounts',
          subTitle: '',
          click_callback: function (obj) {
            console.log('Currency ' + obj.code + ' clicked')
          },
          columns: [...tableColumns],
          data: [...tableData]
        }
      }
    }
  }

</script>
