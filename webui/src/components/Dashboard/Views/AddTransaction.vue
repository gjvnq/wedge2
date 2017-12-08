<template>
  <div>
    <div class="row">
      <div class="col-md-12">
        <div class="card">
          <div class="content">
            <div class="row">
              <div class="col-md-6">
                <div class="form-group">
                  <label>{{$t('Name')}}</label>
                  <input type="text" class="form-control border-input" v-model="transactionName">
                </div>
              </div>
              <div class="col-md-4">
                <div class="form-group">
                  <label>{{$t('Date')}}</label>
                  <input type="date" class="form-control border-input" v-model="transactionDate">
                </div>
              </div>
              <div class="col-md-2">
                <div class="form-group">
                  <div style="height: 27px"></div>
                  <button class="btn btn-info btn-fill btn-wd" :disabled="saving == false" @click="save">{{$t('Save')}}</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="row">
      <h4>{{ $t("Movements" )}}</h4>
    </div>
    <div class="row" v-for="(movement, index) in movements">
      <div class="col-md-12">
        <div class="card">
          <div class="content">
            <movement :accountsList="accountsList" :assetsList="assetsList" v-model="movements[index]" :index="index" :delete-callback="deleteMovement"></movement>
          </div>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="text-center">
        <button class="btn btn-info btn-fill btn-wd" :disabled="saving == false" @click="addMovement">{{$t('Add Movement')}}</button>
      </div>
    </div>
    <div class="row">
      <h4>{{ $t("Items" )}}</h4>
    </div>
    <div class="row" v-for="(item, index) in items">
      <div class="col-md-12">
        <div class="card">
          <div class="content">
            <item :assetsList="assetsList" v-model="items[index]" :index="index" :delete-callback="deleteItem"></item>
          </div>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="text-center">
        <button class="btn btn-info btn-fill btn-wd" :disabled="saving == false" @click="addItem">{{$t('Add Item')}}</button>
      </div>
    </div>
    <div class="row" style="margin-top: 20px">
      <div class="text-center">
        <p>{{ $t("The save button is at the top of this page") }}</p>
      </div>
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
      addMovement () {
        this.movements.push({})
      },
      addItem () {
        this.items.push({})
      },
      save () {
      },
      deleteItem (index) {
        this.items.splice(index, 1)
      },
      deleteMovement (index) {
        this.movements.splice(index, 1)
      },
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
        saving: true,
        movements: [{}],
        items: []
      }
    }
  }

</script>
