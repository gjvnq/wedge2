<template>
  <div class="form-group">
    <label v-if="label">{{$t(label)}}</label>
    <select class="form-control border-input" :value="value" @input="$emit('input', $event.target.value)" :disabled="disabled">
      <option disabled selected value> --- </option>
      <option v-if="item.name !== ''" v-for="item in list" :value="item.id">{{item.name}}</option>
    </select>
  </div>
</template>
<script>
  export default {
    name: 'selector',
    props: {
      list: Array,
      value: String,
      label: String,
      disabled: {
        type: Boolean,
        default: false
      }
    }
  }
</script>