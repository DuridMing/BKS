<template>
    <button type="button" class="btn btn-outline-dark" data-bs-toggle="modal" :data-bs-target="getbtsTarget">
    <font-awesome-icon icon="pen-to-square" />
    </button>

    <!-- Modal -->
    <div class="modal fade" :id="getbtsid" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
    <div class="modal-dialog modal-dialog-centered modal-dialog-scrollable">
        <div class="modal-content">
        <div class="modal-header">
            <h5 class="modal-title" id="exampleModalLabel">book id : {{ itemMessage._id}}</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
        </div>
        <div class="modal-body">
            <div class="mb-3 row">
                <label class="col-sm-2 col-form-label">Name</label>
                <div class="col-sm-10">
                    <input class="form-control" id="itemName" :placeholder="itemMessage.name" v-model.lazy.trim="EdititemMessage.name">
                </div>
            </div>
            <div class="mb-3 row">
                <label class="col-sm-2 col-form-label">Author</label>
                <div class="col-sm-10">
                    <input class="form-control" id="itemAuthor" :placeholder="itemMessage.author" v-model.lazy.trim="EdititemMessage.author">
                </div>
            </div>
            <div class="mb-3 row">
                <label class="col-sm-2 col-form-label">EP</label>
                <div class="col-sm-10">
                    <input class="form-control" id="itemAuthor" :placeholder="itemMessage.ep" v-model.lazy.trim="EdititemMessage.ep">
                </div>
            </div>
            <div class="mb-3 row">
                <label class="col-sm-2 col-form-label">type</label>
                <div class="col-sm-10">
                    <input class="form-control" id="itemAuthor" :placeholder="itemMessage.type" v-model.lazy.trim="EdititemMessage.type">
                </div>
            </div>
            <div class="mb-3 row">
                <label class="col-sm-2 col-form-label">Type</label>
                <div class="col-sm-10">
                    <select class="form-select" aria-label="Default select example" v-model.lazy.trim="EdititemMessage.end">
                        <option selected>{{ itemMessage.end }}</option>
                        <option value="true">True</option>
                        <option value="false">False</option>
                    </select>
                </div>
            </div>
            
        </div>
        <div class="modal-footer">
            <button type="button" class="btn btn-danger " @click="Delete" data-bs-dismiss="modal" :data-bs-target="geteditinfoTarget" data-bs-toggle="modal">Delete</button>
            <button type="button" class="btn btn-secondary " data-bs-dismiss="modal">Close</button>
            <button type="button" class="btn btn-primary" @click="Modify" data-bs-dismiss="modal" :data-bs-target="geteditinfoTarget" data-bs-toggle="modal">Modify</button>
        </div>
        </div>
    </div>
    </div>
    <div class="modal fade" :id="geteditinfoid" aria-hidden="true" aria-labelledby="exampleModalToggleLabel2" tabindex="-1">
    <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
        <div class="modal-header">
            <h5 class="modal-title" id="exampleModalToggleLabel2" >{{ editsucessinfo }}</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
        </div>
        <div class="modal-body" >
            {{ editinfo }}
        </div>
        <div class="modal-footer">
            <button class="btn btn-primary"  data-bs-dismiss="modal">OK</button>
        </div>
        </div>
    </div>
    </div>



        

</template>

<script>
import store from '../../store'
import axios from "axios"

export default {
    name: 'EdiiBookItem',
    data () {
        return {
            btsTarget:"",
            idTarget:"" ,
            editsucessinfo:'',
            editinfo:'',
            EdititemMessage:{
                name : "",
                author : "",
                ep : "",
                type : "",
                end : "",
            }
        }
    },
    props:{
        id: Number,
        itemMessage: {
            type: Object,
        }
    },
    methods: {
        Modify(){
            var token = store.state.accessToken;

            for (const key in this.EdititemMessage){
                if (this.EdititemMessage[key] === ""){
                    this.EdititemMessage[key] = this.$props.itemMessage[key];
                }
            }

            console.log(this.EdititemMessage);

            axios.post('/api/bkms/modify' ,
            {
                id : this.$props.itemMessage._id,
                modifyItem: this.EdititemMessage
            },
            { headers: { 'Authorization': `Bearer ${token}`}
            }).then((response) =>{
                console.log(response.data);
                this.editsucessinfo = "Success";
                this.editinfo = "Success Modify ID: " + this.$props.itemMessage._id;

            }).catch((error) =>{
                console.log(error);
                this.editsucessinfo = "Error";
                this.editinfo = error.response.data;
            })
        },
        Delete(){
            var token = store.state.accessToken;
            axios.post('/api/bkms/delete' ,
            {
                id : this.$props.itemMessage._id,
            },
            { headers: { 'Authorization': `Bearer ${token}`}
            }).then((response) =>{
                console.log(response.data);
                this.editsucessinfo = "Success";
                this.editinfo = "Success delete ID: " + this.$props.itemMessage._id;
            }).catch((error) =>{
                console.log(error);
                this.editsucessinfo = "Error";
                this.editinfo = error.response.data;
            })
        },
    },
    computed:{
        getbtsTarget(){
            return "#model"+this.$props.id;
        },
        getbtsid(){
            return "model"+this.$props.id;
        },
        geteditinfoTarget(){
            return "#editmodel"+this.$props.id;
        },
        geteditinfoid(){
            return "editmodel"+this.$props.id;
        },

    }
}
</script>

<style>

</style>