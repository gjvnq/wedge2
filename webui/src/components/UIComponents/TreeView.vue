<template>
  <li class="tree-view">
    <p>{{model.name}}</p>
    <ul v-if="isFolder">
      <tree-view
        class="tree-item"
        v-for="sub_model in model.children"
        :key="sub_model.id"
        :model="sub_model">
      </tree-view>
    </ul>
  </li>
</template>
<script>
  export default {
    name: 'tree-view',
    props: {
      model: Object
    },
    beforeMount () {
      this.$appData.accounts.first()
    },
    data: function () {
      return {}
    },
    computed: {
      isFolder: function () {
        return this.model.children &&
          this.model.children.length
      }
    }
  }

</script>
<style>
li.tree-view {
  list-style:none;
}
li.tree-view p {
  margin-bottom: 0;
}
li.tree-view ul {
  padding-left: 1.5em;
}
</style>
