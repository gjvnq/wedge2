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
                    <option v-for="account in accounts" :value="account.ID">{{account.Name}}</option>
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
              <button class="btn btn-info btn-fill btn-wd" :disabled="newAccountBtn == false" @click="addAsset">{{$t('Add Account')}}</button>
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
        <ul>
          <li>kdljfdsa</li>
        </ul>
        <div class="header">
        </div>
      </div>
    </div>
  </div>
</template>
<script>
  import PaperTable from 'components/UIComponents/PaperTable.vue'
  const tableColumns = ['Code', 'Name', 'Last Value', '']
  const tableData = []

  export default {
    components: {
      PaperTable
    },
    beforeMount () {
      this.updateAssets()
    },
    methods: {
      addAsset () {
        if (this.newAccountBtn === false) {
          return
        }
        this.newAccountBtn = false
        // Data
        var fd = {}
        fd['code'] = this.newAccountCode
        fd['name'] = this.newAccountName
        fd['places'] = this.newAccountPlaces
        // Send request
        this.$http.put('books/{book-id}/assets', fd).then(response => { // Success
          this.newAccountBtn = true
          window.book_id = fd['book_id']
          this.newAccountCode = ''
          this.newAccountName = ''
          this.newAccountPlaces = 0
        }, response => { // Error
          console.log('err', response)
          this.newAccountBtn = true
          alert(response.bodyText)
        })
      },
      updateAssets () {
        // Send request
        this.$http.get('books/{book-id}/assets').then(response => { // Success
          this.rawAssetsList = response.body
          this.tblAssets.data = this.rawAssetsList
          console.log('---')
          console.log(this.tblAssets.data)
        }, response => { // Error
          console.log('err', response)
        })
      }
    },
    data () {
      return {
        newAccountCode: '',
        newAccountName: '',
        newAccountPlaces: 0,
        newAccountBtn: true,
        rawAssetsList: [],
        tblAssets: {
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
<style>

</style>
