<template>
  <div>
    <div class="row">
      <div class="col-md-3">
        <div class="form-group">
          <label>{{$t('Name')}}</label>
          <input type="text" class="form-control border-input" v-model="value.name" @input="updateValue($event.target.value)">
        </div>
      </div>
      <div class="col-md-2">
        <div class="form-group">
          <label>{{$t('Unit Value')}}</label>
          <input type="text" class="form-control border-input" v-model.number="value.unit" @input="updateValue($event.target.value)">
        </div>
      </div>
      <div class="col-md-2">
        <div class="form-group">
          <label>{{$t('Quantity')}}</label>
          <input type="text" class="form-control border-input" v-model.number="value.quantity" @input="updateValue($event.target.value)">
        </div>
      </div>
      <div class="col-md-2">
        <div class="form-group">
          <label>{{$t('Total Value')}}</label>
          <input type="text" class="form-control border-input" v-model.number="value.total" @input="updateValue($event.target.value)">
        </div>
      </div>

      <div class="col-md-3">
        <asset-selector label="Currency or Asset" v-model="value.asset" :list="assetsList" @input="updateValue($event)"/>
      </div>
    </div>
    <div class="row">
      <div class="col-md-7">
        <div class="form-group">
          <label>{{$t('Tags (comma separated)')}}</label>
          <input type="text" class="form-control border-input" v-model="value.tags" @input="updateValue($event.target.value)">
        </div>
      </div>
      <div class="col-md-2">
        <div class="form-group">
          <label>{{$t('Period Start')}}</label>
          <date-input v-model="value.start" @input="updateValue($event)"/>
        </div>
      </div>
      <div class="col-md-2">
        <div class="form-group">
          <label>{{$t('Period End')}}</label>
          <date-input v-model="value.end" @input="updateValue($event)"/>
        </div>
      </div>
      <div class="col-md-1">
        <div class="form-group">
          <div style="height: 27px"></div>
          <button class="btn btn-danger btn-fill" @click="deleteMe"><span class="ti-trash"></span></button>
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
      index: Number,
      value: {
        name: String,
        unit: Number,
        quantity: Number,
        total: Number,
        asset: String,
        start: Date,
        end: Date
      }
    },
    methods: {
      updateValue () {
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
      }
    }
  }
</script>