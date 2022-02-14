import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import installElementPlus from './plugins/element'
import Axios from "axios";


const app = createApp(App)
app.config.globalProperties.Axios = Axios
installElementPlus(app)
app.use(router).mount('#app')