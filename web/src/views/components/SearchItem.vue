<template>
    <nav class="navbar navbar-expand-lg navbar-light bg-white ">
      <div class="input-group cen-div">
        <select class="form-select" id="SearchMethodSelect" aria-label="Default select example" v-model="Searchtype">
          <option selected>Search</option>
          <option value="Name">Book</option>
          <option value="Author">Author</option>
        </select>
        <input type="text" id="SearchInput" v-model="searchInput" class="form-control w-50">
        <button class="btn btn-outline-success" type="button" id="BookSearch" @click="Search(Searchtype,searchInput)">Search</button>
      </div>
    </nav>
</template>

<script>
import store from "../../store"
import axios from "axios"
export default {
    name: 'SearchItem',
    data () {
      return {
        Searchtype: "Search",
        searchInput: ""
      }
    },
    methods:{
      Search (Searchtype , searchInput){
        var url ="/api/bkms/s"
        var token = store.state.accessToken;
          url = url +"/"+Searchtype.toLowerCase();
          if (Searchtype.toLowerCase() == "name"){
              axios.post(url ,
                {name: searchInput},
                { headers: { 'Authorization': `Bearer ${token}`}
                }).then( function (responce){
                    store.state.bookSearchResult = responce.data;
                })
                .catch( function (error){
                  console.log(error);
              });
          }else if (Searchtype.toLowerCase() == "author"){
            axios.post(url ,
                {author: searchInput},
                { headers: { 'Authorization': `Bearer ${token}`}
                }).then( function (responce){
                    store.state.bookSearchResult = responce.data;
                })
                .catch( function (error){
                  console.log(error);
                });
            }
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
  margin: 1px 270px 
}
</style>