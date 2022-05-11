<template>
    <navbar></navbar>
    <div class="container">
        <div class="cen-div">
            <search></search>
        </div>
    </div>
</template>

<script>
import navbar from "../components/NavBar.vue"
import search from "../components/SearchItem.vue"
export default {
  name: 'SearchPage',
  components: {
      navbar,
      search
  }
}
</script>

<style scoped>
.cen-div {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 20px;
}
</style>
