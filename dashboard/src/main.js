import { createApp } from 'vue'
import App from './App.vue'

import './assets/css/app.css'

import router from './router'

//const app = createApp(App)
const app = createApp(App);
app.use(router);
app.mount('#app');
