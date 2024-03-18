import {createApp} from 'vue'
import './style.css'
import App from './App.vue'
import router from "@/routes";
import ScreenAdapter from "@/components/ScreenAdapter.vue";
import "@/styles/tailwind.css";
// element css
import "element-plus/dist/index.css";
// element dark css
import "element-plus/theme-chalk/dark/css-vars.css";
// custom element dark css
import "@/styles/element-dark.scss";
// custom element css
import "@/styles/element.scss";
// custom element css
import "@/styles/theme/header.ts";
// element icons
import * as Icons from "@element-plus/icons-vue";
// element plus
import ElementPlus from "element-plus";
// iconfont css
import "@/assets/iconfont/iconfont.scss";
// font css
import "@/assets/fonts/font.scss";

import pinia from "@/stores";

import VueCompareImage from "vue3-compare-image";



const app = createApp(App)
Object.keys(Icons).forEach(key => {
    app.component(key, Icons[key as keyof typeof Icons]);
});

app.use(router)
app.use(pinia)
app.use(ElementPlus)
app.use(VueCompareImage)
app.component('screen-adapter', ScreenAdapter)
app.mount("#app")



