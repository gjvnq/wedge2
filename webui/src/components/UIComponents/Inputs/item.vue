<template>
  <div>
    <div class="row">
      <div class="col-md-3">
        <div class="form-group">
          <label>{{$t('Name')}}</label>
          <input type="text" class="form-control border-input" v-model="value.name" :disabled="disabled">
        </div>
      </div>
      <div class="col-md-2">
        <div class="form-group">
          <label>{{$t('Unit Value')}}</label>
          <input type="text" class="form-control border-input" v-model.number="value.unit_cost_user" @input="calcFromUnit" :disabled="disabled" @change="onChange">
        </div>
      </div>
      <div class="col-md-2">
        <div class="form-group">
          <label>{{$t('Quantity')}}</label>
          <input type="text" class="form-control border-input" v-model.number="value.quantity" @input="calcFromQuantity" :disabled="disabled" @change="onChange">
        </div>
      </div>
      <div class="col-md-2">
        <div class="form-group">
          <label>{{$t('Total Value')}}</label>
          <input type="text" class="form-control border-input" v-model.number="value.total_cost_user" @input="calcFromTotal" :disabled="disabled" @change="onChange">
        </div>
      </div>

      <div class="col-md-3">
        <asset-selector label="Currency or Asset" v-model="value.asset_id" :list="assetsList" :disabled="disabled" @change="onChange"/>
      </div>
    </div>
    <div class="row">
      <div class="col-md-7">
        <div class="form-group">
          <label>{{$t('Tags (comma separated)')}}</label>
          <input type="text" class="form-control border-input" v-model="value.tags" @input="computeTags" :disabled="disabled">
        </div>
      </div>
      <div class="col-md-2">
        <div class="form-group">
          <label>{{$t('Period Start')}}</label>
          <date-input v-model="value.start" :disabled="disabled"/>
        </div>
      </div>
      <div class="col-md-2">
        <div class="form-group">
          <label>{{$t('Period End')}}</label>
          <date-input v-model="value.end" :disabled="disabled"/>
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
        name: {
          type: String,
          default: ''
        },
        unit_cost: {
          type: Number,
          default: 0
        },
        unit_cost_user: {
          type: Number,
          default: 0
        },
        quantity: {
          type: Number,
          default: 0
        },
        total_cost: {
          type: Number,
          default: 0
        },
        total_cost_user: {
          type: Number,
          default: 0
        },
        asset_id: {
          type: String,
          default: ''
        },
        start: {
          type: String,
          default: ''
        },
        end: {
          type: String,
          default: ''
        },
        tags: {
          type: Array,
          default: []
        }
      }
    },
    beforeMount () {
      if (this.value.unit_cost !== undefined && this.value.unit_cost !== 0) {
        this.setUnitCost(this.value.unit_cost)
      }
      if (this.value.total_cost !== undefined && this.value.total_cost !== 0) {
        this.setTotalCost(this.value.total_cost)
      }
    },
    methods: {
      onChange (e) {
        this.value.total_cost_user = Math.floor(this.value.total_cost * 1E8)
        this.$emit('change', this.value)
      },
      deleteMe () {
        if (this.deleteCallback !== undefined) {
          this.deleteCallback(this.index)
        }
      },
      setUnitCost (unitCost) {
        this.value.unit_cost = unitCost
        this.value.unit_cost_user = this.value.unit_cost / 1E8
      },
      setTotalCost (totalCost) {
        this.value.total_cost = totalCost
        this.value.total_cost_user = this.value.total_cost / 1E8
      },
      calcFromUnit () {
        if (this.value.quantity !== undefined && this.value.quantity !== 0) {
          this.value.total_cost = this.value.unit_cost * this.value.quantity
        }
        if (this.value.total_cost !== undefined && this.value.total_cost !== 0) {
          this.value.quantity = this.value.total_cost / this.value.unit_cost
        }
      },
      calcFromQuantity () {
        if (this.value.total_cost !== undefined && this.value.total_cost !== 0) {
          this.value.unit_cost = this.value.total_cost / this.value.quantity
        }
        if (this.value.unit_cost !== undefined && this.value.unit_cost !== 0) {
          this.value.total_cost = this.value.unit_cost * this.value.quantity
        }
      },
      calcFromTotal () {
        if (this.value.quantity !== undefined && this.value.quantity !== 0) {
          this.value.unit_cost = this.value.total_cost / this.value.quantity
        }
        if (this.value.unit_cost !== undefined && this.value.unit_cost !== 0) {
          this.value.quantity = this.value.total_cost / this.value.unit_cost
        }
      },
      computeTags () {
        if (this.tags === undefined) {
          this.tags = ''
        }
        this.tags_list = this.tags.split(',')
      }
    },
    data () {
      return {
        tags_list: []
      }
    }
  }
</script>