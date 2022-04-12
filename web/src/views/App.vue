<template>
  <navbar></navbar>
  <div class="container ">
    <home-page></home-page>
  </div>

</template>

<script>
import navbar from "./components/NavBar";
import HomePage from "./components/HomePage.vue";

export default {
  name: 'App',
  components: {
    navbar,
    HomePage
  }
}
</script>

<style>

</style>
