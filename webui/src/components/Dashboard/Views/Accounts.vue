<template>
  <div class="row">
    <div class="col-md-12">
      <div class="card">
        <div class="header">
          <h4 class="title">{{$t('Add Account')}}</h4>
        </div>
        <div class="content">
          <form @submit.prevent>
            <div class="row">
              <div class="col-md-1"></div>
              <div class="col-md-5">
                <div class="form-group">
                  <label>{{$t('Parent')}}</label>
                  <select class="form-control border-input" v-model="newAccountParent">
                    <option value="00000000-0000-0000-0000-000000000000">{{$t('No parent account')}}</option>
                    <option v-for="account in accountsList" :value="account.id">{{account.name}}</option>
                  </select>
                </div>
              </div>
              <div class="col-md-5">
                <div class="form-group">
                  <label>{{$t('Name')}}</label>
                  <input type="text" class="form-control border-input" v-model="newAccountName">
                </div>
              </div>
            </div>
            <div class="text-center">
              <button class="btn btn-info btn-fill btn-wd" :disabled="newAccountBtn == false" @click="addAccount">{{$t('Add Account')}}</button>
            </div>
            <div class="clearfix">
            </div>
          </form>
        </div>
      </div>
    </div>
    <div class="col-md-12">
      <div class="card">
        <div class="header">
          <h4 class="title">{{$t('Accounts')}}</h4>
        </div>
        <p></p>
        <tree-view :model="accountsTree">
        </tree-view>
        <div class="header">
        </div>
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
      this.updateAccounts()
    },
    methods: {
      addAccount () {
        if (this.newAccountBtn === false) {
          return
        }
        this.newAccountBtn = false
        // Data
        var fd = {}
        fd['name'] = this.newAccountName
        fd['parent_id'] = this.newAccountParent
        console.log(fd)
        // Send request
        this.$http.put('books/{book-id}/accounts', fd).then(response => { // Success
          this.newAccountBtn = true
          window.book_id = fd['book_id']
          this.newAccountName = ''
          this.updateAccounts()
        }, response => { // Error
          console.log('err', response)
          this.newAccountBtn = true
          alert(response.bodyText)
          this.updateAccounts()
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
        newAccountName: '',
        newAccountParent: '00000000-0000-0000-0000-000000000000',
        newAccountBtn: true,
        accountsTree: {},
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
