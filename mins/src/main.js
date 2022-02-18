import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import installElementPlus from './plugins/element'
import Axios from "axios";
import {ElMessageBox} from "element-plus";



const app = createApp(App)
app.config.globalProperties.Axios = Axios
installElementPlus(app)
app.use(router).mount('#app')

let token = window.localStorage.getItem("min_token");
if (token === null){
    ElMessageBox.prompt("请输入页面token","提示").then(data=>{
        if (data !== null){
            window.localStorage.setItem("min_token",data.value)
        }
    })
}

Axios.interceptors.request.use(function (config) {
    // 在发送请求之前做些什么
    console.log(config.url)
    config.headers.token = window.localStorage.getItem("min_token")
    return config;
}, function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
});

