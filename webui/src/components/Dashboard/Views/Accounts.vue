<template>
  <div>
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
                    <label>{{$t('Parent account')}}</label>
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
    </div>
    <account-tree-view :model="accountsTree" :level="0"></account-tree-view>
  </div>
</template>
<script>
  import AccountTreeView from 'components/UIComponents/AccountTreeView.vue'

  export default {
    components: {
      AccountTreeView
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
        this.$store.dispatch('updateAccounts')
      }
    },
    computed: {
      accountsList () {
        return this.$store.state.accounts
      },
      accountsTree () {
        return this.$store.state.accountsTree
      }
    },
    data () {
      return {
        newAccountName: '',
        newAccountParent: '00000000-0000-0000-0000-000000000000',
        newAccountBtn: true
      }
    }
  }

</script>
