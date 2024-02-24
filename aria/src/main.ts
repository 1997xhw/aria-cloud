import {createApp} from 'vue'
import './style.css'
import App from './App.vue'
import router from "@/routes";
import ScreenAdapter from "@/components/ScreenAdapter.vue";
import "./style/tailwind.css";


const app = createApp(App)
app.use(router)

app.component('screen-adapter', ScreenAdapter)
app.mount("#app")



