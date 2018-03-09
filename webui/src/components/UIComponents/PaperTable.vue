<template>
  <div>
    <div class="header" v-if="title != '' && subTitle != ''">
      <slot name="header">
        <h4 class="title" v-if="title != ''">{{title}}</h4>
        <p class="category" v-if="subTitle != ''">{{subTitle}}</p>
      </slot>
    </div>
    <div class="content table-responsive table-full-width paper-table">
      <table class="table" :class="tableClass">
        <thead>
          <th v-for="column in columns">{{$t(column)}}</th>
          <th></th>
        </thead>
        <tbody>
          <tr v-for="item in data" :class="trClass" @click="click_event(item)">
            <td v-for="(column, col_index) in columns" :class="itemClass(col_index)" v-if="hasValue(item, column, col_index)">{{itemValue(item, column, col_index)}}</td>
            <td v-html="itemValue(item, '_extra')"></td>
            <td v-if="hasValue(item, '_link_txt') && hasValue(item, '_link_addr')"><router-link :to="itemValue(item, '_link_addr')">{{$t(itemValue(item, '_link_txt'))}}</router-link></td>
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
      columnsProperties: {
        type: Array,
        default: () => []
      },
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
      hasValue (item, column, index) {
        return this.itemValue(item, column, index) !== undefined
      },
      itemValue (item, column, index) {
        if (this.columnsProperties.length !== 0 && typeof index === 'number') {
          return item[this.columnsProperties[index]]
        }
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
