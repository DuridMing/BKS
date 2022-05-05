<template>
  <navbar></navbar>
  <div class="container">
    <div id="HomeLogo" class="cen-div">
      <img alt="Vue logo" height="250" width="250" src="../../assets/logo.png">
      <h3>The Book Managment System</h3>
    </div>
  </div>
</template>

<script>
import navbar from "./NavBar.vue"
export default {
  name: 'HelloWorld',
  components:{
    navbar
  },
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 35px 0 0;
}
.cen-div {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
