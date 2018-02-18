<template>
  <div>
    <div class="row">
      <div class="col-md-3" :class="{ 'has-error': nameErr }">
        <div class="form-group">
          <label>{{$t('Name')}}</label>
          <input type="text" class="form-control border-input" v-model="value.name" :disabled="disabled" @change="onChange">
        </div>
      </div>
      <div class="col-md-2" :class="{ 'has-error': unitErr }">
        <div class="form-group">
          <label>{{$t('Unit Value')}}</label>
          <input type="text" class="form-control border-input" v-model.number="unit_cost_user" :disabled="disabled" @change="onChange">
        </div>
      </div>
      <div class="col-md-2" :class="{ 'has-error': quantityErr }">
        <div class="form-group">
          <label>{{$t('Quantity')}}</label>
          <input type="text" class="form-control border-input" v-model.number="value.quantity" :disabled="disabled" @change="onChange">
        </div>
      </div>
      <div class="col-md-2" :class="{ 'has-error': totalErr }">
        <div class="form-group">
          <label>{{$t('Total Value')}}</label>
          <input type="text" class="form-control border-input" v-model.number="total_cost_user" :disabled="disabled" @change="onChange">
        </div>
      </div>

      <div class="col-md-3" :class="{ 'has-error': assetErr }">
        <asset-selector label="Currency or Asset" v-model="value.asset_id" :list="assetsList" :disabled="disabled" @change="onChange"/>
      </div>
    </div>
    <div class="row">
      <div class="col-md-7" :class="{ 'has-error': tagsErr }">
        <div class="form-group">
          <label>{{$t('Tags (comma separated)')}}</label>
          <input type="text" class="form-control border-input" v-model="tags_user" :disabled="disabled" @change="onChange">
        </div>
      </div>
      <div class="col-md-2" :class="{ 'has-error': startErr }">
        <div class="form-group">
          <label>{{$t('Period Start')}}</label>
          <date-input v-model="value.period_start" :disabled="disabled" @change="onChange"/>
        </div>
      </div>
      <div class="col-md-2" :class="{ 'has-error': endErr }">
        <div class="form-group">
          <label>{{$t('Period End')}}</label>
          <date-input v-model="value.period_end" :disabled="disabled" @change="onChange"/>
        </div>
      </div>
      <div class="col-md-1">
        <div class="form-group">
          <div style="height: 27px"></div>
          <button class="btn btn-danger btn-fill" @click="deleteMe" :disabled="disabled"><span class="ti-trash"></span></button>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
  import selector from 'components/UIComponents/Inputs/selector.vue'
  import assetSelector from 'components/UIComponents/Inputs/assetSelector.vue'
  import dateInput from 'components/UIComponents/Inputs/dateInput.vue'

  export default {
    components: {
      selector,
      assetSelector,
      dateInput
    },
    name: 'item',
    props: {
      assetsList: Array,
      accountsList: Array,
      deleteCallback: Function,
      disabled: {
        type: Boolean,
        default: false
      },
      index: Number,
      value: {
        type: Object,
        default: function () {
          return {
            value: '',
            unit_cost: 0,
            quantity: 0,
            total_cost: 0,
            asset_id: '',
            period_start: '',
            period_end: '',
            tags: []
          }
        }
      }
    },
    beforeMount () {
      if (this.value.unit_cost !== undefined && this.value.unit_cost !== 0) {
        this.unit_cost_user = this.value.unit_cost / 1E8
      }
      if (this.value.total_cost !== undefined && this.value.total_cost !== 0) {
        this.total_cost_user = this.value.total_cost / 1E8
      }
      if (this.value.tags === undefined) {
        this.value.tags = []
      } else {
        this.tags_user = this.value.tags.join(', ')
      }
    },
    methods: {
      onChange (e) {
        this.value.unit_cost = Math.floor(this.unit_cost_user * 1E8)
        this.value.total_cost = Math.floor(this.total_cost_user * 1E8)
        this.computeTags()
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
      computeTags () {
        this.value.tags = this.tags_user.split(',')
      },
      validate () {
        this.nameErr = (this.value.name === undefined || this.value.name === '')
        this.unitErr = isNaN(this.unit_cost_user)
        this.quantityErr = isNaN(this.value.quantity)
        this.totalErr = isNaN(this.total_cost_user)
        this.assetErr = (this.value.asset_id === undefined || this.value.asset_id === '')
        this.tagsErr = false
        this.startErr = false
        this.endErr = false
        this.validate_on_change = true
        return (this.nameErr || this.unitErr || this.quantityErr || this.totalErr || this.assetErr || this.tagsErr || this.startErr)
      }
    },
    data () {
      return {
        tags_user: '',
        validate_on_change: false,
        nameErr: false,
        unitErr: false,
        quantityErr: false,
        totalErr: false,
        assetErr: false,
        tagsErr: false,
        startErr: false,
        endErr: false,
        unit_cost_user: '',
        total_cost_user: ''
      }
    }
  }
</script>