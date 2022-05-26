<template>
    <form class="shadow-cus pas-div">
      <div class="mb-3 row ">
        <label  class="col-sm-2 col-form-label" >Name</label>
        <div class="col-sm-10"> 
          <input type="text" class="form-control" v-model.lazy.trim="AddBookItemInfo.name">
        </div>   
      </div>
     <div class="mb-3 row ">
        <label  class="col-sm-2 col-form-label">Author</label>
        <div class="col-sm-10"> 
          <input type="text" class="form-control" v-model.lazy.trim="AddBookItemInfo.author">
        </div>   
      </div>
      <div class="mb-3 row ">
        <label  class="col-sm-2 col-form-label">Episode</label>
        <div class="col-sm-10"> 
          <input type="text" class="form-control" v-model.lazy.trim="AddBookItemInfo.ep">
        </div>   
      </div>
      <div class="mb-3 row ">
        <label  class="col-sm-2 col-form-label">Type</label>
        <div class="col-sm-10"> 
          <input type="text" class="form-control" v-model.lazy.trim="AddBookItemInfo.type">
        </div>   
      </div>
      <div class="mb-3 row ">
        <label  class="col-sm-2 col-form-label">End</label>
        <div class="col-sm-10"> 
           <select class="form-select" aria-label="Default select example" v-model.lazy.trim="AddBookItemInfo.end">
            <option selected>Is it finished?</option>
            <option value="true">True</option>
            <option value="false">False</option>
          </select>
        </div> 
    
      </div>
      <button type="button" class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#AddbookModal">
        <font-awesome-icon icon="circle-plus" />
        Add The Book
      </button>
    </form>


    <!-- Modal -->
    <div class="modal fade" id="AddbookModal" tabindex="-1" aria-labelledby="AddBookLabel" aria-hidden="true">
      <div class="modal-dialog modal-dialog-centered modal-dialog-scrollable">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="AddBookLabel">Final Check</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            <form>
              <div class="mb-3 row ">
                <label  class="col-sm-2 col-form-label" >Name</label>
                <div class="col-sm-10"> 
                  <input type="text" class="form-control-plaintext" readonly :value="AddBookItemInfo.name">
                </div>   
              </div>
              <div class="mb-3 row ">
                <label  class="col-sm-2 col-form-label" >Author</label>
                <div class="col-sm-10"> 
                  <input type="text" class="form-control-plaintext" readonly :value="AddBookItemInfo.author">
                </div>   
              </div>
              <div class="mb-3 row ">
                <label  class="col-sm-2 col-form-label" >Episode</label>
                <div class="col-sm-10"> 
                  <input type="text" class="form-control-plaintext" readonly :value="AddBookItemInfo.ep">
                </div>   
              </div>
              <div class="mb-3 row ">
                <label  class="col-sm-2 col-form-label" >Type</label>
                <div class="col-sm-10"> 
                  <input type="text" class="form-control-plaintext" readonly :value="AddBookItemInfo.type">
                </div>   
              </div>
              <div class="mb-3 row ">
                <label  class="col-sm-2 col-form-label" >End</label>
                <div class="col-sm-10"> 
                  <input type="text" class="form-control-plaintext" readonly :value="getendinfo">
                </div>   
              </div>
            </form>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
            <button type="button" class="btn btn-primary" data-bs-dismiss="modal" data-bs-toggle="modal" data-bs-target="#AddBookInfoMadal" @click="AddBook">
            <font-awesome-icon icon="circle-plus" />
            Add
            </button>
          </div>
        </div>
      </div>
    </div>

  <!-- info modal0 -->
  <div class="modal fade" id="AddBookInfoMadal" aria-hidden="true" aria-labelledby="exampleModalToggleLabel2" tabindex="-1">
  <div class="modal-dialog modal-dialog-centered">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title" id="exampleModalToggleLabel2">{{ addbooksuccessinfo }}</h5>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
      </div>
      <div class="modal-body">
        {{ addbookinfo }}
      </div>
      <div class="modal-footer">
        <button class="btn btn-primary" data-bs-dismiss="modal">OK</button>
      </div>
    </div>
  </div>
</div>

</template>

<script>
import store from '../../store'
import axios from "axios"
export default {
    name: 'AddBookItem',
    data () {
      return {
        addbooksuccessinfo: '',
        addbookinfo:'',
        AddBookItemInfo: {
            name:'',
            author:'',
            ep:'',
            type:'',
            end: Boolean,
        },
      }
    },
    methods:{
      AddBook() {
        var token = store.state.accessToken;
        var payload = this.AddBookItemInfo
        if (this.AddBookItemInfo.end == true){
          payload.end = true;
        }else{
          payload.end = false;
        }
        axios.post('/api/bkms/add' ,payload,
            { headers: { 
              'Authorization': `Bearer ${token}`,
              'Content-Type': 'application/json'
              }}
            ).then((response) =>{
                console.log(response.data);
                this.addbooksuccessinfo = "Success";
                this.addbookinfo = response.data.message;
            }).catch((error) =>{
                console.log(error);
                this.addbooksuccessinfo = "Error";
                this.addbookinfo = error.response.data;
            })
      
      }
    },
    computed:{
      getendinfo(){
        return this.AddBookItemInfo.end.toString();
      }
    }

}
</script>

<style>
.cen-div {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
.shadow-cus{
  box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.2), 0 2.5rem 5rem 0 rgba(0, 0, 0, 0.1)
}
.pas-div{
  background-color: white;
  padding:25px 10px 20px 10px;
}
</style>