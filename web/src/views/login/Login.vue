<template>
    <div class="container">
      <div class="footer-60">
        <div class="cen-div">
          <img alt="Vue logo" height="150" width="150" src="../../assets/logo-tr.png">
        </div>
        <div class="box-cen">
            <div class="pas-div shadow-cus">
                <div class="mb-3 row">
                  <label  class="col-sm-2 col-form-label">Account</label>
                  <div class="col-sm-10">
                    <input type="password" class="form-control" id="inputAccount">
                  </div>
                </div>
                <div class="mb-3 row">
                  <label  class="col-sm-2 col-form-label">Password</label>
                  <div class="col-sm-10">
                    <input type="password" class="form-control" id="inputPassword">
                  </div>
                </div>
                <button class="btn btn-primary" type="submit">
                  <i class="fa-solid fa-arrow-right-to-bracket"></i>
                  Login
                </button>
            </div> 
        </div>
      </div>
    </div>

</template>

<script>
export default {
  name: 'LoginItem'
}
</script>

<style scoped >

.cen-div {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.box-cen{
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  margin: 10px 60px 10px 60px;
  padding: 10px 220px 0px 220px;
}
.shadow-cus{
  box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.2), 0 2.5rem 5rem 0 rgba(0, 0, 0, 0.1)
}
.pas-div{
  background-color: white;
  padding:30px 10px 20px 10px;
}
</style>
