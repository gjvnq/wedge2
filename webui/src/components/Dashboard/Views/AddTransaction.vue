<template>
  <div>
    <div class="row">
      <div class="card">
        <div class="content">
          <div class="row">
            <div class="col-md-6">
              <div class="form-group">
                <label>{{$t('Name')}}</label>
                <input type="text" class="form-control border-input" v-model="transactionName" :disabled="saving">
              </div>
            </div>
            <div class="col-md-4">
              <div class="form-group">
                <label>{{$t('Date')}}</label>
                <input type="date" class="form-control border-input" v-model="transactionDate" :disabled="saving">
              </div>
            </div>
            <div class="col-md-2">
              <div class="form-group">
                <div style="height: 27px"></div>
                <button class="btn btn-info btn-fill btn-wd" :disabled="saving == true" @click="save">{{$t('Save')}}</button>
              </div>
            </div>
          </div>
          <div class="text-center">
            <span id="msgErrConn" class="label label-warning" :class="{ hide: !flagErrConn }" v-t="'Failed to comunicate with the server :-('"></span>
            <span id="msgErrOk" class="label label-success" :class="{ hide: !flagOk }" v-t="'Successfully added transaction'"></span>
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
            <movement :accountsList="accountsList" :assetsList="assetsList" v-model="movements[index]" :index="index" :delete-callback="deleteMovement" @change="autoSetAssetMov" :disabled="saving"></movement>
          </div>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="text-center">
        <button class="btn btn-info btn-fill btn-wd" :disabled="saving == true" @click="addMovement">{{$t('Add Movement')}}</button>
      </div>
    </div>
    <div class="row">
      <h4>{{ $t("Items" )}}</h4>
    </div>
    <div class="row" v-for="(item, index) in items">
      <div class="col-md-12">
        <div class="card">
          <div class="content">
            <item :assetsList="assetsList" v-model="items[index]" :index="index" :delete-callback="deleteItem" :disabled="saving"></item>
          </div>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="text-center">
        <button class="btn btn-info btn-fill btn-wd" :disabled="saving == true" @click="addItem">{{$t('Add Item')}}</button>
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
    methods: {
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
      },
      save () {
        if (this.saving === false) {
          return
        }
        this.saving = false
        // Data
        var fd = {}
        fd['name'] = this.transactionName
        fd['local_date'] = this.transactionDate
        fd['movements'] = []
        for (let movement of this.movements) {
          let fd2 = {}
          fd2['account_id'] = movement.account
          fd2['asset_id'] = movement.asset
          fd2['amount'] = movement.amount
          fd2['status'] = movement.status
          fd2['local_date'] = movement.date
          fd2['amount'] = movement.amount
          fd['movements'].push(fd2)
        }
        fd['items'] = []
        for (let item of this.items) {
          let fd2 = {}
          fd2['name'] = item.name
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
          this.saving = true
          this.updateTransactions()
        }, response => { // Error
          console.log('err', response)
          this.saving = true
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
        this.$root.$children[0].$children[0].updateTransactions()
      }
    },
    props: ['accountsList', 'assetsList'],
    data () {
      return {
        transactionName: '',
        transactionDate: '',
        saving: false,
        default_asset: '',
        movements: [{}],
        items: [{}, {}],
        flagOk: false,
        flagErrConn: false
      }
    }
  }

</script>
