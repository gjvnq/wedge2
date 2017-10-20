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
                <div class="form-group">
                  <p id="msgErrConn" class="label label-warning hide" v-t="'Failed to comunicate with the server :-('"></p>
                </div>
                <div class="form-group">
                  <p id="msgErrPass" class="label label-warning hide" v-t="'Wrong password :-('"></p>
                </div>
                <button id="btnLogin" type="submit" class="btn btn-default" v-on:click="login" v-t="'Login'"></button>
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
      showError (err) {
        var el = null
        if (err === 'conn') {
          el = document.querySelector('#msgErrConn')
        }
        if (err === 'pass') {
          el = document.querySelector('#msgErrPass')
        }
        el.classList.remove('hide')
      },
      clearErrors () {
        document.querySelector('#msgErrPass').classList.add('hide')
        document.querySelector('#msgErrConn').classList.add('hide')
      },
      lockBtn () {
        document.querySelector('#btnLogin').disabled = true
      },
      unlockBtn () {
        document.querySelector('#btnLogin').disabled = false
      },
      login () {
        this.lockBtn()
        var fd = new FormData()
        fd.append('bookId', document.querySelector('#inBookId').value)
        fd.append('password', document.querySelector('#inBookPassword').value)
        this.clearErrors()
        this.$http.post('auth', fd).then(response => {
          // Success
          this.unlockBtn()
          console.log('suc', response)
        }, response => {
          // Error
          this.unlockBtn()
          console.log('err', response)
          if (response.status !== 403) {
            this.showError('conn')
          } else {
            this.showError('pass')
          }
        })
      }
    },
    mounted: function () {
      this.clv = new Cleave('#inBookId', {
        delimiters: ['-', '-', '-', '-', '-'],
        blocks: [8, 4, 4, 4, 12],
        lowercase: true
      })
    },
    beforeDestory: function () {
      this.clv.destroy()
    }
  }
</script>
