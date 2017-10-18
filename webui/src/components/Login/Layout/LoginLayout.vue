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
                  <label for="inBookId">BookId</label>
                  <input type="text" class="form-control" id="inBookId" placeholder="Ex: 00000000-0000-0000-0000-00000000000">
                </div>
                <div class="form-group">
                  <label for="inBookPassword">Password</label>
                  <input type="password" class="form-control" id="inBookPassword" placeholder="Password">
                </div>
                <button type="submit" class="btn btn-default" v-on:click="login">Login</button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default{
    props: {
    },
    methods: {
      login () {
        alert('hi')
      }
    }
  }
</script>
