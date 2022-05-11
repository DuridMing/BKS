
import { createStore } from "vuex";
import createPersistedState from 'vuex-persistedstate';

// vuex-persisted not working
// I don't know why

const store = createStore({
        state:{
            accessToken:'',
            refreshToken:'',
            userName:'',
            isAuthenticated : false,
        },
        plugins: [
            createPersistedState()
        ],
        mutations:{

        },
        actions:{

        },
        modules:{

        },

});

export default store
