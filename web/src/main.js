import { createApp } from 'vue'
import App from './views/App.vue'
import { Routers } from './router';

import '@fortawesome/fontawesome-free/css/all.css'
import '@fortawesome/fontawesome-free/js/all.js'

import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap";



const app = createApp(App)
app.use(Routers)

app.mount('#app');
