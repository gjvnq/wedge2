<template>
  <div class="row">
    <div class="col-md-12">
      <div class="card">
        <div class="header">
          <h4 class="title">{{$t('Transactions')}} <router-link to="transactions/add">({{$t('Add')}})</router-link></h4>
        </div>
        <p></p>
<!--         <tree-view :model="accountsTree">
        </tree-view> -->
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
    methods: {
    },
    props: [],
    data () {
      return {
      }
    }
  }

</script>
