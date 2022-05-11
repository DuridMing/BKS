<template>
  <navbar></navbar>
  <div class="container">
    <div class="cen-div">
      <dashboard></dashboard>
    </div>
  </div>
</template>

<script>
import navbar from "../components/NavBar.vue"
import dashboard from "../components/DashboardItem.vue"
export default {
  name: 'DashPage',
  components: {
    navbar,
    dashboard
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
