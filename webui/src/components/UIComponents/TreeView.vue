<template>
  <li class="tree-view">
    <p v-if="html === true && formatter !== undefined" v-html="format(model)"></p>
    <p v-if="html === false && formatter !== undefined">{{ format(model) }}</p>
    <p v-if="formatter === undefined">{{ model.name }}</p>
    <ul v-if="isFolder">
      <tree-view
        class="tree-item"
        v-for="sub_model in model.children"
        :key="sub_model.id"
        :html="html"
        :formatter="formatter"
        :model="sub_model">
      </tree-view>
    </ul>
  </li>
</template>
<script>
  export default {
    name: 'tree-view',
    props: {
      model: Object,
      formatter: Function,
      html: Boolean
    },
    data: function () {
      return {}
    },
    methods: {
      format (value) {
        return this.formatter(value)
      }
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
