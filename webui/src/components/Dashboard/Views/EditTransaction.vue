<template>
  <div>
    <div class="row">
      <div class="card">
        <div class="content">
          <div class="row">
            <div class="col-md-6" :class="{ 'has-error': nameErr }">
              <div class="form-group">
                <label>{{$t('Name')}}</label>
                <input type="text" class="form-control border-input" v-model="value.name" :disabled="saving || all_disabled" @input="validateBasic">
              </div>
            </div>
            <div class="col-md-3" :class="{ 'has-error': dateErr }">
              <div class="form-group">
                <label>{{$t('Date')}}</label>
                <input type="date" class="form-control border-input" v-model="value.local_date" :disabled="saving || all_disabled" @input="validateBasic" v-shortkey="['ctrl', 'alt', 'c']" @shortkey="setMovDates">
              </div>
            </div>
            <div class="col-md-1">
              <div class="form-group" v-if="value.id !== undefined">
                <div style="height: 27px"></div>
                <button class="btn btn-danger btn-fill" @click="deleteMe" :disabled="saving || all_disabled"><span class="ti-trash"></span></button>
              </div>
            </div>
            <div class="col-md-2">
              <div class="form-group">
                <div style="height: 27px"></div>
                <button class="btn btn-info btn-fill btn-wd" :disabled="saving || all_disabled" @click="save">{{$t('Save')}}</button>
              </div>
            </div>
          </div>
          <div class="row">
            <div class="col-md-12 text-center">
              <p>{{ $t('Tip: use Ctrl+Alt+C while on the date field to copy its value to all movements.') }}</p>
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
    <div class="row" v-for="(movement, index) in value.movements">
      <div class="col-md-12">
        <div class="card">
          <div class="content">
            <movement ref="movements" :accountsList="accountsList" :assetsList="assetsList" v-model="value.movements[index]" :index="index" :delete-callback="deleteMovement" @change="autoSetAssetMov" :disabled="saving || all_disabled"></movement>
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
    <div class="row" v-for="(item, index) in value.items">
      <div class="col-md-12">
        <div class="card">
          <div class="content">
            <item ref="items" :assetsList="assetsList" v-model="value.items[index]" :index="index" :delete-callback="deleteItem" :disabled="saving || all_disabled"></item>
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
        this.value = {
          id: undefined,
          name: '',
          local_date: undefined,
          movements: [{}],
          items: []
        }
        this.saving = false
        this.default_asset = ''
        this.flagOk = false
        this.flagErrConn = false
        this.flagLoading = false
        this.nameErr = false
        this.dateErr = false
        this.warn_zero_sum = false
        this.all_disabled = false
        if (this.$route.params.tr_id !== 'new') {
          this.value.id = this.$route.params.tr_id
          this.all_disabled = true
          this.value.movements = []
          this.value.items = []
          this.loadAndFillTransaction(this.value.id)
        }
      },
      deleteMe () {
        if (!confirm(this.$t('Do you really want to DELETE this transaction?'))) {
          return
        }
        this.saving = true
        this.$http.delete('books/{book-id}/transactions/' + this.value.id, this.value).then(response => { // Success
          this.saving = false
          this.notifyVue('success', 'ti-trash', 'Transaction removed')
          this.updateTransactions()
        }, response => { // Error
          console.log('err', response)
          this.saving = false
          this.notifyVue('danger', 'ti-alert', this.$t('Failed to talk to server') + ': ' + response.bodyText)
          this.updateTransactions()
        })
      },
      notifyVue (kind, icon, msg) {
        if (!['info', 'success', 'warning', 'danger'].includes(kind)) {
          kind = 'info'
        }
        this.$notifications.notify(
          {
            message: this.$t(msg),
            icon: icon,
            horizontalAlign: 'center',
            verticalAlign: 'top',
            type: kind
          })
      },
      validateBasic () {
        this.nameErr = (this.value.name === undefined || this.value.name === '')
        this.dateErr = (this.value.local_date === undefined || this.value.local_date.substr(0, 4) === '0000')
        return !(this.nameErr || this.dateErr)
      },
      validate () {
        var failedValidation = false

        failedValidation |= !this.validateBasic()

        if (this.value.movements.length > 0) {
          for (let mov of this.$refs.movements) {
            failedValidation |= mov.validate()
          }
        }
        if (this.value.items.length > 0) {
          for (let item of this.$refs.items) {
            failedValidation |= item.validate()
          }
        }

        return !failedValidation
      },
      loadAndFillTransaction (id) {
        if (this.flagLoading) {
          return
        }
        this.flagLoading = true
        this.$http.get('books/{book-id}/transactions/' + id).then(response => { // Success
          if (this.value.id === response.body.id) {
            this.value = response.body
          }
          this.all_disabled = false
          this.flagLoading = false
        }, response => { // Error
          console.log('err', response)
          if (this.value.id !== '') {
            this.flagLoading = false
            this.flagErrConn = true
          }
        })
      },
      addMovement () {
        var newMov = {}
        var len = this.value.movements.length
        if (len > 0) {
          var oldMov = this.value.movements[len - 1]
          newMov.asset_id = oldMov.asset_id
          newMov.account_id = oldMov.account_id
          newMov.status = oldMov.status
          newMov.local_date = oldMov.local_date
        }
        this.value.movements.push(newMov)
      },
      addItem () {
        this.value.items.push({asset: this.default_asset})
      },
      deleteItem (index) {
        this.value.items.splice(index, 1)
        this.validate()
      },
      deleteMovement (index) {
        this.value.movements.splice(index, 1)
        this.validate()
      },
      autoSetAssetMov () {
        if (this.value.movements.length === 0) {
          // Array is too small
          return
        }

        let asset = this.value.movements[0].asset
        for (let movement of this.value.movements) {
          if (movement.asset !== asset) {
            // Assets are not consistent
            this.default_asset = ''
            return
          }
        }
        this.default_asset = asset

        // Check for sum
        this.warn_zero_sum = false
        if (this.value.movements.length === 2) {
          var prod = this.value.movements[0].amount * this.value.movements[1].amount
          var sum = this.value.movements[0].amount + this.value.movements[1].amount
          var equalAssets = (this.value.movements[0].asset_id === this.value.movements[1].asset_id)
          if (prod < 0 && sum !== 0 && equalAssets) {
            this.warn_zero_sum = true
          }
        }
      },
      save () {
        // Do not save twice at the same time
        if (this.saving === true) {
          return
        }
        this.saving = false
        // Validate before saving
        if (!this.validate()) {
          return
        }
        // Send request
        this.$http.put('books/{book-id}/transactions', this.value).then(response => { // Success
          this.saving = false
          this.notifyVue('success', 'ti-save', 'Transaction edited')
          this.updateTransactions()
        }, response => { // Error
          console.log('err', response)
          this.saving = false
          this.notifyVue('danger', 'ti-alert', this.$t('Failed to talk to server') + ': ' + response.bodyText)
          this.updateTransactions()
        })
      },
      clear () {
        this.movments = [{}]
        this.value.items = []
        this.transactionName = ''
        this.transactionDate = ''
        this.flagErrConn = false
        this.flagOk = false
      },
      updateTransactions () {
      },
      setMovDates () {
        for (var i = 0; i < this.value.movements.length; i++) {
          this.$set(this.value.movements[i], 'local_date', this.value.local_date)
        }
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
        value: {
          id: undefined,
          name: '',
          local_date: undefined,
          movements: [{}],
          items: []
        },
        saving: false,
        default_asset: '',
        flagOk: false,
        flagErrConn: false,
        flagLoading: false,
        warn_zero_sum: false,
        all_disabled: false,
        nameErr: false,
        dateErr: false
      }
    }
  }

</script>
