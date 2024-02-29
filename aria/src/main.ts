import {createApp} from 'vue'
import './style.css'
import App from './App.vue'
import router from "@/routes";
import ScreenAdapter from "@/components/ScreenAdapter.vue";
import "./style/tailwind.css";
import "element-plus/dist/index.css";
import pinia from "@/store";

const app = createApp(App)
app.use(router)
app.use(pinia)
app.component('screen-adapter', ScreenAdapter)
app.mount("#app")



