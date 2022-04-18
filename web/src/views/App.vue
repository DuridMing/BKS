<template>
  <navbar></navbar>
  <div class="container">
    <router-view></router-view>
  </div>

</template>

<script>
import navbar from "./components/NavBar";

export default {
  name: 'App',
  components: {
    navbar,
  }
}
</script>

<style>

</style>
