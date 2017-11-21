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
                  <input type="password" class="form-control" id="inBookPassword" v-model="password">
                </div>
                <div class="form-group">
                  <p id="msgErr404" class="label label-warning" :class="{ hide: flagHideErr404 }" v-t="'No such book :-('"></p>
                </div>
                <div class="form-group">
                  <p id="msgErrConn" class="label label-warning" :class="{ hide: flagHideErrConn }" v-t="'Failed to comunicate with the server :-('"></p>
                </div>
                <div class="form-group">
                  <p id="msgErrPass" class="label label-warning" :class="{ hide: flagHideErrPass }" v-t="'Wrong password :-('"></p>
                </div>
                <button id="btnLogin" type="submit" class="btn btn-default" :disabled="flagBtn == false" v-on:click="login" v-t="'Login'"></button>
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
      clearErrors () {
        this.flagHideErr404 = true
        this.flagHideErrPass = true
        this.flagHideErrConn = true
      },
      login () {
        // UI
        this.clearErrors()
        this.flagBtn = false
        // Data
        var fd = {}
        fd['book_id'] = this.selected_book
        fd['password'] = this.password
        // Send request
        this.$http.post('auth', fd).then(response => { // Success
          this.flagBtn = true
          Vue.http.headers.common['Authorization'] = 'Bearer ' + response.bodyText
          this.$router.push('book')
        }, response => { // Error
          console.log('err', response)
          this.flagBtn = true
          if (response.status === 404) {
            this.flagHideErr404 = false
          } else if (response.status === 403 || response.status === 401) {
            this.flagHideErrPass = false
          } else {
            this.flagHideErrConn = false
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
        selected_book: null,
        password: '',
        flagHideErr404: true,
        flagHideErrPass: true,
        flagHideErrConn: true,
        flagBtn: true
      }
    }
  }
</script>
