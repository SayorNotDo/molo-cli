import { createApp } from 'vue';
import { createPinia } from 'pinia';
import router from './router';
import ElementPlus from 'element-plus'
import App from './App.vue';
import 'element-plus/dist/index.css';
// import './assets/style/global.less'


const app = createApp(App);
const pinia = createPinia();

app.use(pinia);
app.use(router);
app.use(ElementPlus)

app.mount('#app');
