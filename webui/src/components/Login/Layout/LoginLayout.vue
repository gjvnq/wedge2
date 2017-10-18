<template>
  <div class="contact-us full-screen">
    <div class="wrapper wrapper-full-page section content bg-nude">
      <div class="">
        <div class="container">
          <div class="row">
            <div class="col-md-8 col-md-offset-2 text-center">
              <h2 class="title text-info">wedge²</h2>
              <form>
                <div class="form-group">
                  <label for="inBookId" v-t="'BookId'"></label>
                  <input type="text" class="form-control" id="inBookId">
                </div>
                <div class="form-group">
                  <label for="inBookPassword" v-t="'Password'"></label>
                  <input type="password" class="form-control" id="inBookPassword">
                </div>
                <button type="submit" class="btn btn-default" v-on:click="login" v-t="'Login'"></button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  var Cleave = require('cleave.js')
  export default {
    methods: {
      login () {
        alert('hi')
      }
    },
    mounted: function () {
      this.clv = new Cleave('#inBookId', {
        delimiters: ['-', '-', '-', '-', '-'],
        blocks: [8, 4, 4, 4, 12],
        uppercase: true
      })
    },
    beforeDestory: function () {
      this.clv.destroy()
    }
  }
</script>
