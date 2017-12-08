<template>
  <div class="row">
    <div class="col-md-2">
      <selector label="Account" v-model="value.account" :list="accountsList"/>
    </div>
    <div class="col-md-2">
      <div class="form-group">
        <label>{{$t('Value')}}</label>
        <input type="text" class="form-control border-input" :value.number="value.amount" @input="update('amount', $event.target.value)">
      </div>
    </div>
    <div class="col-md-3">
      <asset-selector label="Currency or Asset" v-model="value.asset" :list="assetsList"/>
    </div>
    <div class="col-md-2">
      <div class="form-group">
        <label>{{$t('Date')}}</label>
        <input type="date" class="form-control border-input" v-model="value.date"/>
      </div>
    </div>
    <div class="col-md-2">
      <selector label="Status" v-model="value.status" :list="statusList"/>
    </div>
    <div class="col-md-1">
      <div class="form-group">
        <div style="height: 27px"></div>
        <button class="btn btn-danger btn-fill" @click="deleteMe"><span class="ti-trash"></span></button>
      </div>
    </div>
  </div>
</template>
<script>
  import selector from 'components/UIComponents/Inputs/selector.vue'
  import assetSelector from 'components/UIComponents/Inputs/assetSelector.vue'

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
      value: {
        account: {
          type: String,
          default: ''
        },
        amount: {
          type: Number,
          default: 0
        },
        asset: {
          type: String,
          default: ''
        },
        date: {
          type: Date,
          default: new Date()
        },
        status: {
          type: String,
          default: ''
        }
      }
    },
    methods: {
      update (field, value) {
        if (field === 'amount') {
          value = Number(value)
        }

        console.log(field, value)
        this.$set(this.value, field, value)
        console.log('input', this.value)
        this.$emit('input', this.value)
      },
      deleteMe () {
        if (this.deleteCallback !== undefined) {
          this.deleteCallback(this.index)
        }
      }
    },
    data () {
      return {
        statusList: [
          {'id': 'P', 'name': this.$t('Planned')},
          {'id': 'D', 'name': this.$t('Done')},
          {'id': 'C', 'name': this.$t('Cancelled')}
        ]
      }
    }
  }
</script>