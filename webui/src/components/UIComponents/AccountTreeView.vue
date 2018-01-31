<template>
  <div>
    <div class="row">
      <div :class="spacer_class()">
      </div>
      <div :class="size_class()">
        <div class="card">
          <div class="header">
            <h4 class="title">{{ get_name(model) }}
              <router-link v-if="model.id !== '00000000-0000-0000-0000-000000000000' && !singleCard" :to="'/book/accounts/'+model.id">({{$t('See movements')}})</router-link></h4>
          </div>
          <div class="content">
            <p v-html="local_totals(model)"></p>
            <p v-if="isFolder || singleCard" v-html="sum_totals(model)"></p>
          </div>
        </div>
      </div>
    </div>
    <account-tree-view
      v-if="isFolder"
      v-for="sub_model in model.children"
      :key="sub_model.id"
      :level="level+1"
      :model="sub_model">
    </account-tree-view>
  </div>
</template>
<script>
  import {_} from 'vue-underscore'
  import numeric from '@/numeric.js'

  export default {
    name: 'account-tree-view',
    props: {
      level: Number,
      model: Object,
      singleCard: Boolean
    },
    data: function () {
      return {}
    },
    methods: {
      spacer_class () {
        return 'col-md-' + this.level
      },
      size_class () {
        return 'col-md-' + (12 - this.level)
      },
      get_name (model) {
        if (model.id === '00000000-0000-0000-0000-000000000000') {
          return this.$t('Total')
        }
        return model.name
      },
      local_totals (acc) {
        if (acc.name === '') {
          return ''
        }
        return this.$t('This account only:') + ' ' + this.generic_totals(acc, 'local_balance_codes')
      },
      sum_totals (acc) {
        return this.$t('Including children accounts:') + ' ' + this.generic_totals(acc, 'total_balance_codes')
      },
      generic_totals (acc, key) {
        var balance = ''
        for (var assetCode in acc[key]) {
          var num = numeric.format(acc[key][assetCode])
          if (num >= 0) {
            num = '<span class="text-success nowrap">+' + num
          } else {
            num = '<span class="text-danger nowrap">' + num
          }
          if (balance === '') {
            balance += num + ' ' + _.escape(assetCode) + '</span>'
          } else {
            balance += ' • ' + num + ' ' + _.escape(assetCode) + '</span>'
          }
        }
        return balance
      }
    },
    computed: {
      isFolder: function () {
        return this.model.children &&
          this.model.children.length !== 0
      }
    }
  }

</script>
<style>
li.tree-view {
  list-style:none;
}
li.tree-view p {
  margin-bottom: 0;
}
li.tree-view ul {
  padding-left: 1.5em;
}
</style>
