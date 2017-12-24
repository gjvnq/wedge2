<template>
  <div class="row">
    <div class="col-md-12">
      <div class="card">
        <div class="header">
          <h4 class="title">{{$t('Add Currency or Asset')}}</h4>
        </div>
        <div class="content">
          <form @submit.prevent>
            <div class="row">
              <div class="col-md-1"></div>
              <div class="col-md-3">
                <div class="form-group">
                  <label>{{$t('Code')}}</label>
                  <input type="text" class="form-control border-input" v-model="newAssetCode">
                </div>
              </div>
              <div class="col-md-7">
                <div class="form-group">
                  <label>{{$t('Name')}}</label>
                  <input type="text" class="form-control border-input" v-model="newAssetName">
                </div>
              </div>
            </div>
            <div class="text-center">
              <button class="btn btn-info btn-fill btn-wd" :disabled="newAssetBtn == false" @click="addAsset">{{$t('Add Currency or Asset')}}</button>
            </div>
            <div class="clearfix">
            </div>
          </form>
        </div>
      </div>
    </div>
    <div class="col-md-12">
      <div class="card">
        <paper-table :title="$t(tblAssets.title)" :sub-title="$t(tblAssets.subTitle)" :data="tblAssetsData" :columns="tblAssets.columns" :columnsStyles="tblAssets.columnsStyles" :click_callback="tblAssets.click_callback">
        </paper-table>
      </div>
    </div>
  </div>
</template>
<script>
  import PaperTable from 'components/UIComponents/PaperTable.vue'
  const tableColumns = ['Code', 'Name', 'Last Value']
  const tableColumnsStyle = ['mono', '', '']

  export default {
    components: {
      PaperTable
    },
    beforeMount () {
      this.updateAssets()
    },
    methods: {
      addAsset () {
        if (this.newAssetBtn === false) {
          return
        }
        this.newAssetBtn = false
        // Data
        var fd = {}
        fd['code'] = this.newAssetCode
        fd['name'] = this.newAssetName
        // Send request
        this.$http.put('books/{book-id}/assets', fd).then(response => { // Success
          this.newAssetBtn = true
          window.book_id = fd['book_id']
          this.newAssetCode = ''
          this.newAssetName = ''
          this.updateAssets()
        }, response => { // Error
          console.log('err', response)
          this.newAssetBtn = true
          alert(response.bodyText)
        })
      },
      updateAssets () {
        this.$store.dispatch('updateAssets')
      }
    },
    props: ['assetsList'],
    computed: {
      tblAssetsData () {
        return this.$store.state.assets
      }
    },
    data () {
      return {
        newAssetCode: '',
        newAssetName: '',
        newAssetBtn: true,
        rawAssetsList: [],
        tblAssets: {
          title: 'Currencies & Assets',
          subTitle: '',
          click_callback: function (obj) {
            console.log('Currency ' + obj.code + ' clicked')
          },
          columns: [...tableColumns],
          columnsStyles: [...tableColumnsStyle]
        }
      }
    }
  }

</script>
<style>

</style>
