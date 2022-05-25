import { createApp } from 'vue'
import App from './views/App.vue'
import { Routers } from './router';

import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap";

import { library } from "@fortawesome/fontawesome-svg-core";
import { faArrowRightToBracket } from '@fortawesome/free-solid-svg-icons'

library.add(faArrowRightToBracket);

const app = createApp(App);
app.use(Routers);
app.component("font-awesome-icon", FontAwesomeIcon);
app.mount('#app');
