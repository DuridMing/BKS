<template>
<div class="table-responsive">
    <table class="table table-striped table-hover text-center">
        <thead>
            <tr class="d-flwx">
                <th class="col-2">#</th>
                <th class="col-3">Book</th>
                <th class="col-2">Author</th>
                <th class="col-1">EP</th>
                <th class="col-2">Type</th>
                <th class="col-2">Edit</th>
            </tr>
        </thead>
        <tbody>
            <template :key="index" v-for="item, index in book">
            <tr >
                <td>{{ index+1 }}</td>
                <td>{{ item.name }}</td>
                <td>{{ item.author }}</td>
                <td>{{ item.ep }}</td>
                <td>{{ item.type }}</td>
                <td><edit-book-item :id="index" :itemMessage="item"></edit-book-item></td>
            </tr>
            </template>
        </tbody>
    </table>
</div>

</template>

<script>
import store from  "../../store"
import EditBookItem from "./EditBookItem.vue";
export default {
    name: 'BookResult',
    components: {
        EditBookItem,
    },
    computed:{
        book(){
            return store.state.bookSearchResult;
        }
    }
}
</script>

<style>

</style>