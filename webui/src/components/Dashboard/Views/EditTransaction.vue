<template>
  <div>
    <div class="row">
      <div class="card">
        <div class="content">
          <div class="row">
            <div class="col-md-6">
              <div class="form-group">
                <label>{{$t('Name')}}</label>
                <input type="text" class="form-control border-input" v-model="transactionName" :disabled="saving || all_disabled">
              </div>
            </div>
            <div class="col-md-4">
              <div class="form-group">
                <label>{{$t('Date')}}</label>
                <input type="date" class="form-control border-input" v-model="transactionDate" :disabled="saving || all_disabled">
              </div>
            </div>
            <div class="col-md-2">
              <div class="form-group">
                <div style="height: 27px"></div>
                <button class="btn btn-info btn-fill btn-wd" :disabled="saving == true || all_disabled" @click="save">{{$t('Save')}}</button>
              </div>
            </div>
          </div>
          <div class="text-center">
            <span id="msgErrConn" class="label label-warning" :class="{ hide: !flagErrConn }" v-t="'Failed to comunicate with the server :-('"></span>
            <span id="msgErrOk" class="label label-success" :class="{ hide: !flagOk }" v-t="'Successfully added transaction'"></span>
            <span class="label label-info" :class="{ hide: !flagLoading }" v-t="'Loading transaction...'"></span>
          </div>
        </div>
      </div>
    </div>
    <div class="row">
      <h4>{{ $t("Movements" )}}</h4>
    </div>
    <div class="row" v-for="(movement, index) in movements">
      <div class="col-md-12">
        <div class="card">
          <div class="content">
            <movement :accountsList="accountsList" :assetsList="assetsList" v-model="movements[index]" :index="index" :delete-callback="deleteMovement" @change="autoSetAssetMov" :disabled="saving || all_disabled"></movement>
          </div>
        </div>
      </div>
    </div>
    <div class="row" style="margin-bottom: 15px">
      <div class="text-center">
        <p class="text-warning" v-if="warn_zero_sum">{{ $t("Shouldn't the sum be zero?") }}</p>
      </div>
    </div>
    <div class="row">
      <div class="text-center">
        <button class="btn btn-info btn-fill btn-wd" :disabled="saving == true || all_disabled" @click="addMovement">{{$t('Add Movement')}}</button>
      </div>
    </div>
    <div class="row">
      <h4>{{ $t("Items" )}}</h4>
    </div>
    <div class="row" v-for="(item, index) in items">
      <div class="col-md-12">
        <div class="card">
          <div class="content">
            <item :assetsList="assetsList" v-model="items[index]" :index="index" :delete-callback="deleteItem" :disabled="saving || all_disabled"></item>
          </div>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="text-center">
        <button class="btn btn-info btn-fill btn-wd" :disabled="saving == true || all_disabled" @click="addItem">{{$t('Add Item')}}</button>
      </div>
    </div>
    <div class="row" style="margin-top: 20px">
      <div class="text-center">
        <p>{{ $t("The save button is at the top of this page") }}</p>
      </div>
    </div>
  </div>
</div>
</template>
<script>
  import TreeView from 'components/UIComponents/TreeView.vue'
  import movement from 'components/UIComponents/Inputs/movement.vue'
  import item from 'components/UIComponents/Inputs/item.vue'

  export default {
    components: {
      TreeView,
      movement,
      item
    },
    beforeMount () {
      this.$store.dispatch('updateAssets')
      this.$store.dispatch('updateAccounts')
      this.start()
    },
    watch: {
      '$route' (to, from) {
        this.start()
      }
    },
    methods: {
      start () {
        this.transactionID = ''
        this.transactionName = ''
        this.transactionDate = ''
        this.saving = false
        this.default_asset = ''
        this.movements = [{}]
        this.items = []
        this.flagOk = false
        this.flagErrConn = false
        this.flagLoading = false
        this.warn_zero_sum = false
        this.all_disabled = false
        if (this.$route.params.tr_id !== 'new') {
          this.transactionID = this.$route.params.tr_id
          this.all_disabled = true
          this.movements = []
          this.items = []
          this.loadAndFillTransaction(this.transactionID)
        }
      },
      loadAndFillTransaction (id) {
        if (this.flagLoading) {
          return
        }
        this.flagLoading = true
        this.$http.get('books/{book-id}/transactions/' + id).then(response => { // Success
          if (this.transactionID === response.body.id) {
            this.transactionName = response.body.name
            this.transactionDate = response.body.local_date
            this.movements = response.body.movements
            this.items = response.body.items
          }
          this.all_disabled = false
          this.flagLoading = false
        }, response => { // Error
          console.log('err', response)
          if (this.transactionID !== '') {
            this.flagLoading = false
            this.flagErrConn = true
          }
        })
      },
      addMovement () {
        this.movements.push({})
      },
      addItem () {
        this.items.push({asset: this.default_asset})
      },
      deleteItem (index) {
        this.items.splice(index, 1)
      },
      deleteMovement (index) {
        this.movements.splice(index, 1)
      },
      autoSetAssetMov () {
        if (this.movements.length === 0) {
          // Array is too small
          return
        }

        let asset = this.movements[0].asset
        for (let movement of this.movements) {
          if (movement.asset !== asset) {
            // Assets are not consistent
            this.default_asset = ''
            return
          }
        }
        this.default_asset = asset

        // Check for sum
        this.warn_zero_sum = false
        if (this.movements.length === 2) {
          var prod = this.movements[0].amount * this.movements[1].amount
          var sum = this.movements[0].amount + this.movements[1].amount
          if (prod < 0 && sum !== 0) {
            this.warn_zero_sum = true
          }
        }
      },
      save () {
        if (this.saving === true) {
          return
        }
        this.saving = false
        // Data
        var fd = {}
        fd['id'] = this.transactionID
        fd['name'] = this.transactionName
        fd['local_date'] = this.transactionDate
        fd['movements'] = []
        for (let movement of this.movements) {
          let fd2 = {}
          fd2['account_id'] = movement.account_id
          fd2['asset_id'] = movement.asset_id
          fd2['amount'] = movement.amount
          fd2['status'] = movement.status
          fd2['local_date'] = movement.local_date
          console.log(movement.local_date.slice(0, 4))
          if (movement.local_date.slice(0, 4) === '0000') {
            return
          }
          fd['movements'].push(fd2)
        }
        fd['items'] = []
        for (let item of this.items) {
          let fd2 = {}
          fd2['name'] = item.name
          fd2['asset_id'] = item.asset_id
          fd2['unit_cost'] = item.unit_cost
          fd2['total_cost'] = item.total_cost
          fd2['quantity'] = item.quantity
          fd2['period_start'] = item.start
          fd2['period_end'] = item.end
          fd2['tags'] = item.tags_list
          fd['items'].push(fd2)
        }
        // Send request
        this.$http.put('books/{book-id}/transactions', fd).then(response => { // Success
          this.saving = false
          this.updateTransactions()
        }, response => { // Error
          console.log('err', response)
          this.saving = false
          alert(response.bodyText)
          this.updateTransactions()
        })
      },
      clear () {
        this.movments = [{}]
        this.items = []
        this.transactionName = ''
        this.transactionDate = ''
        this.flagErrConn = false
        this.flagOk = false
      },
      updateTransactions () {
        this.$store.dispatch('updateTransactions')
      }
    },
    computed: {
      assetsList () {
        return this.$store.state.assets
      },
      accountsList () {
        return this.$store.state.accounts
      }
    },
    data () {
      return {
        transactionID: '',
        transactionName: '',
        transactionDate: '',
        saving: false,
        default_asset: '',
        movements: [{}],
        items: [],
        flagOk: false,
        flagErrConn: false,
        flagLoading: false,
        warn_zero_sum: false,
        all_disabled: false
      }
    }
  }

</script>
