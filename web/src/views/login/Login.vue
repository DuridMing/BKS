<template>
    <div class="container">
      <div class="footer-60">
        <div class="cen-div">
          <img alt="Vue logo" height="150" width="150" src="../../assets/logo-tr.png">
        </div>
        <div class="box-cen">
            <div class="pas-div shadow-cus">
              <span style="padding-bottom: 10px;">{{ login_error_message }}</span>
                <div class="mb-3 row">
                  <label  class="col-sm-2 col-form-label" >Account</label>
                  <div class="col-sm-10">
                    <input  class="form-control" id="inputAccount" v-model.lazy="Account">
                  </div>
                </div>
                <div class="mb-3 row">
                  <label  class="col-sm-2 col-form-label">Password</label>
                  <div class="col-sm-10">
                    <input type="password" class="form-control" id="inputPassword" v-model.lazy="Password">
                  </div>
                </div>
                <button class="btn btn-primary" @click="login(Account , Password)">
                  <font-awesome-icon icon="arrow-right-to-bracket" />
                  <span >&nbsp;</span>
                  Login
                </button>
            </div> 
        </div>
      </div>
    </div>

</template>

<script>
import store from "../../store"
import authcrypto from "./auth"
import axios from "axios"

export default {
  name: 'LoginPage',
  data () {
    return {
      Account: "",
      Password: "",
      login_error_message: "",
    }
  },
  methods:{
    login(Account, Password) {
      var url = '/api/login';
      let hashPassword = authcrypto.encrypt(Password);
      var data = this;
      axios.post(url ,
        {username:Account , password:hashPassword})
        .then( function (responce){
          var access_token = responce.data.access_token;
          store.state.accessToken = access_token;
          var refresh_token = responce.data.refresh_token;
          store.state.refreshToken = refresh_token;
          store.state.isAuthenticated = true;
          store.state.userName = Account;
          console.log(store.state.userName);
          data.$router.push({name: 'home'});
        })
        .catch( function (error){
          console.log(error);
          console.log(error.response.status);
          console.log(error.response.data);
          if (error.response.status  == 422){
            data.login_error_message = error.response.data;
          }
          if ( error.response.status == 401){
            data.login_error_message = error.response.data;
          }
        });
    }
    
  },
  computed:{
    
  }
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
  padding:25px 10px 20px 10px;
}
</style>
