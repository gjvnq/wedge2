<template>
  <div class="row">
    <div class="col-md-12">
      <div class="card">
        <div class="header">
          <h4 class="title">{{$t('Transactions')}} <router-link to="transactions/add">({{$t('Add')}})</router-link></h4>
        </div>
        <p v-for="(transaction, index) in transactions">{{ transaction }}</p>
        <div class="header">
        </div>
      </div>
    </div>
  </div>
</template>
<script>
  export default {
    components: {
    },
    beforeMount () {
      this.loadTransactions()
    },
    methods: {
      loadTransactions () {
        if (this.transactions_loading) {
          return
        }
        this.transactions_loading = true
        this.$http.get('books/{book-id}/transactions').then(response => { // Success
          this.transactions = response.body
          this.transactions_loading = false
        }, response => { // Error
          console.log('err', response)
          this.transactions_loading = false
        })
      }
    },
    props: [],
    data () {
      return {
        transactions: [],
        transactions_loading: false
      }
    }
  }

</script>
