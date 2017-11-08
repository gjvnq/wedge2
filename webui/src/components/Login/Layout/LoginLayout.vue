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
                  <label for="inBookId" v-t="'Book'"></label>
                  <select class="form-control" id="inBookId" v-model="selected_book">
                    <option v-for="book in books" :value="book.ID">{{book.Name}} ({{book.ID}})</option>
  </select>
                  </select>
                </div>
                <div class="form-group">
                  <label for="inBookPassword" v-t="'Password'"></label>
                  <input type="password" class="form-control" id="inBookPassword">
                </div>
                <div class="form-group">
                  <p id="msgErr404" class="label label-warning hide" v-t="'No such book :-('"></p>
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
  import Vue from 'vue'
  export default {
    methods: {
      showError (err) {
        var el = null
        if (err === 'conn') {
          el = document.querySelector('#msgErrConn')
        }
        if (err === '404') {
          el = document.querySelector('#msgErr404')
        }
        if (err === 'password') {
          el = document.querySelector('#msgErrPass')
        }
        el.classList.remove('hide')
      },
      clearErrors () {
        document.querySelector('#msgErrPass').classList.add('hide')
        document.querySelector('#msgErrConn').classList.add('hide')
        document.querySelector('#msgErr404').classList.add('hide')
      },
      lockBtn () {
        document.querySelector('#btnLogin').disabled = true
      },
      unlockBtn () {
        document.querySelector('#btnLogin').disabled = false
      },
      login () {
        this.lockBtn()
        var fd = {}
        fd['book_id'] = document.querySelector('#inBookId').value
        fd['password'] = document.querySelector('#inBookPassword').value
        this.clearErrors()
        this.$http.post('auth', fd).then(response => {
          // Success
          this.unlockBtn()
          Vue.http.headers.common['Authorization'] = 'Bearer ' + response.bodyText
          console.log('suc', response)
          console.log('suc', this.$router)
          this.$router.push('book')
        }, response => {
          // Error
          this.unlockBtn()
          console.log('err', response)
          if (response.status === 404) {
            this.showError('404')
          } else if (response.status === 403 || response.status === 401) {
            this.showError('password')
          } else {
            this.showError('conn')
          }
        })
      },
      list_books () {
        this.$http.get('books').then(response => {
          console.log('suc', response)
          this.books = response.body
          console.log(this.books)
        }, response => {
          console.log('err', response)
        })
      }
    },
    mounted () {
      this.list_books()
    },
    data () {
      return {
        books: [],
        selected_book: null
      }
    }
  }
</script>
