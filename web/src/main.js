import { createApp } from 'vue'
import App from './views/App.vue'
import { Routers } from './router';


import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap";

const app = createApp(App).use(Routers);

app.mount('#app');
