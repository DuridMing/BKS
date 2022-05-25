<template>
    <navbar></navbar>
    <search></search>
    <div class="container">
      <div class="table-div">
        <bookresult></bookresult>
      </div>
    </div>
</template>

<script>
import navbar from "../components/NavBar.vue"
import search from "../components/SearchItem.vue"
import bookresult from "../components/BookResultItem.vue"
export default {
  name: 'SearchPage',
  components: {
      navbar,
      search,
      bookresult
  }
}
</script>

<style scoped>
.table-div{
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin: 10px 0px;
}
</style>
