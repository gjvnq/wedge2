<template>
  <div class="row">
    <div class="col-md-12">
      <div class="card">
        <div class="header">
          <h4 class="title">{{$t('Add Transaction')}}</h4>
        </div>
        <div class="content">
          <form @submit.prevent>
            <div class="row">
              <div class="col-md-1"></div>
              <div class="col-md-5">
                <div class="form-group">
                  <label>{{$t('Name')}}</label>
                  <input type="text" class="form-control border-input" v-model="transactionName">
                </div>
              </div>
              <div class="col-md-5">
                <div class="form-group">
                  <label>{{$t('Date')}}</label>
                  <input type="date" class="form-control border-input" v-model="transactionDate">
                </div>
              </div>
            </div>
            <hr>
            <h4 class="title">{{$t('Movements')}}</h4>
            <movement :accountsList="accountsList" :assetsList="assetsList" v-for="(movement, index) in movements" v-model="movements[index]"/>
            <div class="text-center">
              <button class="btn btn-info btn-fill btn-wd" :disabled="transactionBtn == false" @click="addAccount">{{$t('Add Movement')}}</button>
            </div>
            <hr>
            <h4 class="title">{{$t('Items')}}</h4>
            <item :assetsList="assetsList" v-for="(item, index) in items" v-model="items[index]"/>
            <div class="text-center">
              <button class="btn btn-info btn-fill btn-wd" :disabled="transactionBtn == false" @click="addAccount">{{$t('Add Item')}}</button>
            </div>
            <hr>
            <div class="text-center">
              <button class="btn btn-info btn-fill btn-wd" :disabled="transactionBtn == false" @click="addAccount">{{$t('Save')}}</button>
            </div>
            <div class="clearfix">
            </div>
          </form>
        </div>
      </div>
    </div>
    <div class="col-md-12">
      <div class="card">
        a
      </div>
    </div>
  </div>
</template>
<script>
  import TreeView from 'components/UIComponents/TreeView.vue'
  import movement from 'components/UIComponents/Inputs/movement.vue'
  import item from 'components/UIComponents/Inputs/item.vue'

  export default {
    components: {
      TreeView,
      movement,
      item
    },
    methods: {
      addAccount () {
        if (this.transactionBtn === false) {
          return
        }
        this.transactionBtn = false
        // Data
        var fd = {}
        fd['name'] = this.transactionName
        fd['parent_id'] = this.transactionParent
        console.log(fd)
        // Send request
        this.$http.put('books/{book-id}/accounts', fd).then(response => { // Success
          this.transactionBtn = true
          window.book_id = fd['book_id']
          this.transactionName = ''
          this.updateAccounts
        }, response => { // Error
          console.log('err', response)
          this.transactionBtn = true
          alert(response.bodyText)
          this.updateAccounts
        })
      },
      updateAccounts () {
        this.$root.$children[0].$children[0].updateAccounts()
      }
    },
    props: ['accountsList', 'assetsList'],
    data () {
      return {
        transactionName: '',
        transactionDate: '',
        transactionBtn: true,
        movements: [
          {}
        ],
        items: [
          {}
        ]
      }
    }
  }

</script>
