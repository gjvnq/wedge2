<template>
  <div class="row">
    <div class="col-md-2" :class="{ 'has-error': accountErr }">
      <selector label="Account" v-model="value.account_id" :list="accountsList" @change="onChange" :disabled="disabled"/>
    </div>
    <div class="col-md-2" :class="{ 'has-error': amountErr }">
      <div class="form-group">
        <label>{{$t('Value')}}</label>
        <input type="text" class="form-control border-input" v-model.number="amount_user" @change="onChange" :disabled="disabled"/>
      </div>
    </div>
    <div class="col-md-3" :class="{ 'has-error': assetErr }">
      <asset-selector label="Currency or Asset" v-model="value.asset_id" :list="assetsList" @change="onChange" :disabled="disabled"/>
    </div>
    <div class="col-md-2" :class="{ 'has-error': dateErr }">
      <div class="form-group">
        <label>{{$t('Date')}}</label>
        <input type="date" class="form-control border-input" v-model="value.local_date" @change="onChange" :disabled="disabled"/>
      </div>
    </div>
    <div class="col-md-2" :class="{ 'has-error': statusErr }">
      <selector label="Status" v-model="value.status" :list="statusList" @change="onChange" :disabled="disabled"/>
    </div>
    <div class="col-md-1">
      <div class="form-group">
        <div style="height: 27px"></div>
        <button class="btn btn-danger btn-fill" @click="deleteMe" :disabled="disabled"><span class="ti-trash"></span></button>
      </div>
    </div>
  </div>
</template>
<script>
  import selector from 'components/UIComponents/Inputs/selector.vue'
  import assetSelector from 'components/UIComponents/Inputs/assetSelector.vue'
  import numeric from '@/numeric.js'

  export default {
    components: {
      selector,
      assetSelector
    },
    name: 'movement',
    props: {
      assetsList: Array,
      accountsList: Array,
      deleteCallback: Function,
      index: Number,
      disabled: {
        type: Boolean,
        default: false
      },
      value: {
        type: Object,
        default: function () {
          return {
            account_id: '',
            amount: 0,
            amount_user: 0,
            asset_id: '',
            local_date: new Date(),
            status: '',
            valid: false
          }
        }
      }
    },
    beforeMount () {
      if (this.value.amount !== undefined && this.value.amount !== 0) {
        this.amount_user = numeric.format(this.value.amount)
      }
    },
    methods: {
      onChange (e) {
        this.value.amount = numeric.parse(this.amount_user)
        this.$emit('change', this.value)
        if (this.validate_on_change) {
          this.validate()
        }
      },
      deleteMe () {
        if (this.deleteCallback !== undefined) {
          this.deleteCallback(this.index)
        }
      },
      validate () {
        this.accountErr = (this.value.account_id === undefined || this.value.account_id === '')
        this.amountErr = (this.value.amount === undefined)
        this.assetErr = (this.value.asset_id === undefined || this.value.asset_id === '')
        this.dateErr = (this.value.local_date === undefined || this.value.local_date.substr(0, 4) === '0000')
        this.statusErr = (this.value.status !== 'P' && this.value.status !== 'D' && this.value.status !== 'C')
        this.validate_on_change = true

        return (this.accountErr || this.amountErr || this.assetErr || this.dateErr || this.statusErr)
      }
    },
    data () {
      return {
        statusList: [
          {'id': 'P', 'name': this.$t('Planned')},
          {'id': 'D', 'name': this.$t('Done')},
          {'id': 'C', 'name': this.$t('Cancelled')}
        ],
        amount_user: '',
        validate_on_change: false,
        accountErr: false,
        amountErr: false,
        assetErr: false,
        dateErr: false,
        statusErr: false
      }
    }
  }
</script>