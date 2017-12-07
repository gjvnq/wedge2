<template>
  <div class="row">
    <div class="col-md-2">
      <selector label="Account" v-model="value.account" :list="accountsList" @input="updateValue($event)"/>
    </div>
    <div class="col-md-2">
      <div class="form-group">
        <label>{{$t('Value')}}</label>
        <input type="text" class="form-control border-input" v-model.number="value.ammount" @input="updateValue($event.target.value)">
      </div>
    </div>
    <div class="col-md-3">
      <asset-selector label="Currency or Asset" v-model="value.asset" :list="assetsList" @input="updateValue($event)"/>
    </div>
    <div class="col-md-2">
      <div class="form-group">
        <label>{{$t('Date')}}</label>
        <input type="date" class="form-control border-input" v-model="value.date" @input="updateValue($event.target.value)"/>
      </div>
    </div>
    <div class="col-md-2">
      <selector label="Status" v-model="value.status" :list="statusList" @input="updateValue($event)"/>
    </div>
    <div class="col-md-1">
      <div class="form-group">
        <div style="height: 27px"></div>
        <button class="btn btn-danger btn-fill"><span class="ti-trash"></span></button>
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
      value: {
        account: String,
        ammount: Number,
        asset: String,
        date: Date,
        status: String
      }
    },
    methods: {
      updateValue () {
        this.$emit('input', this.value)
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