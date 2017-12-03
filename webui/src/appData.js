export default {
  AccountsList: [],
  AccountsTree: {},
  AccountsLoading: false,
  updateAccounts (that) {
    // Send request
    this.AccountsLoading = true
    that.$http.get('books/{book-id}/accounts').then(response => { // Success
      this.AccountsList = response.body
      console.log(this.AccountsList)
      this.AccountsLoading = false
    }, response => { // Error
      this.AccountsLoading = false
    })
    // Send request
    that.$http.get('books/{book-id}/accounts-tree').then(response => { // Success
      this.AccountsTree = response.body
      this.AccountsLoading = false
    }, response => { // Error
      this.AccountsLoading = false
    })
  }
}
