<template>
  <nav class="navbar navbar-expand-lg navbar-light bg-light">
    <div class="container-fluid">
        <router-link class="navbar-brand" :to="{name: 'home'}">
            <img src="../..//assets/logo.png" alt="" width="35" height="35">
            BKMS
        </router-link>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarSupportedContent">
            <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                <li class="nav-item">
                    <router-link class="nav-link active" aria-current="page" :to="{name:'dashboard'}">Dashboard</router-link>
                </li>
                <li class="nav-item">
                    <router-link class="nav-link active" aria-current="page" :to="{name:'search'}">Search</router-link>
                </li>
            </ul>

            <span >
              <div class="btn-group">
                    <button type="button" class="btn btn-light dropdown-toggle" data-bs-toggle="dropdown" data-bs-display="static" aria-expanded="false">
                        Login
                    </button>
                    <ul class="dropdown-menu dropdown-menu-lg-end">
                        <li><button class="dropdown-item" type="button">Action</button></li>
                        <li><button class="dropdown-item" type="button">Another action</button></li>
                        <li><button class="dropdown-item" type="button">Something else here</button></li>
                    </ul>
                </div>
            </span>
        </div>
    </div>
  </nav>
</template>

<script>
export default {
    name: 'NavBar',
}
</script>

<style>

</style>