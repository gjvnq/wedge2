<template>
  <div class="form-group">
    <label v-if="label">{{$t(label)}}</label>
    <input type="date" class="form-control border-input" :value="value" @input="$emit('input', $event.target.value)" v-on="$listeners" :disabled="disabled">
  </div>
</template>
<script>
  export default {
    name: 'date-input',
    props: {
      value: String,
      label: String,
      disabled: {
        type: Boolean,
        default: false
      }
    }
  }
</script>