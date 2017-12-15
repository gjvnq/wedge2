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
          <input type="text" class="form-control border-input" v-model.number="value.unit_cost" @input="calcFromUnit" :disabled="disabled">
        </div>
      </div>
      <div class="col-md-2">
        <div class="form-group">
          <label>{{$t('Quantity')}}</label>
          <input type="text" class="form-control border-input" v-model.number="value.quantity" @input="calcFromQuantity" :disabled="disabled">
        </div>
      </div>
      <div class="col-md-2">
        <div class="form-group">
          <label>{{$t('Total Value')}}</label>
          <input type="text" class="form-control border-input" v-model.number="value.total_cost" @input="calcFromTotal" :disabled="disabled">
        </div>
      </div>

      <div class="col-md-3">
        <asset-selector label="Currency or Asset" v-model="value.asset" :list="assetsList" :disabled="disabled"/>
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
        quantity: {
          type: Number,
          default: 0
        },
        total_cost: {
          type: Number,
          default: 0
        },
        asset: {
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
    methods: {
      deleteMe () {
        if (this.deleteCallback !== undefined) {
          this.deleteCallback(this.index)
        }
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