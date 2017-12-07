<template>
  <div>
    <div class="header">
      <slot name="header">
        <h4 class="title">{{title}}</h4>
        <p class="category">{{subTitle}}</p>
      </slot>
    </div>
    <div class="content table-responsive table-full-width">
      <table class="table" :class="tableClass">
        <thead>
          <th v-for="column in columns">{{$t(column)}}</th>
          <th></th>
        </thead>
        <tbody>
          <tr v-for="item in data" :class="trClass" @click="click_event(item)">
            <td v-for="(column, col_index) in columns" :class="itemClass(col_index)" v-if="hasValue(item, column)">{{itemValue(item, column)}}</td>
            <td v-html="itemValue(item, '_extra')"></td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
<script>
  export default {
    props: {
      columns: Array,
      columnsStyles: Array,
      data: Array,
      click_callback: Function,
      type: {
        type: String, // striped | hover
        default: 'hover'
      },
      title: {
        type: String,
        default: ''
      },
      subTitle: {
        type: String,
        default: ''
      }
    },
    computed: {
      tableClass () {
        return `table-${this.type}`
      },
      trClass () {
        if (this.click_callback !== undefined) {
          return 'click-cursor'
        }
        return ''
      }
    },
    methods: {
      hasValue (item, column) {
        return item[column.toLowerCase()] !== 'undefined'
      },
      itemValue (item, column) {
        return item[column.toLowerCase()]
      },
      itemClass (index) {
        if (this.columnsStyles === undefined) {
          return ''
        }
        return this.columnsStyles[index]
      },
      click_event (obj) {
        if (this.click_callback !== undefined) {
          this.click_callback(obj)
        }
      }
    }
  }

</script>
<style>

</style>
