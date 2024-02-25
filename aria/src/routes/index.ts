import { createRouter, createWebHistory } from "vue-router";
const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/login',
            name: 'login',
            component: () => import('@/views/home/signin.vue')
        },
        {
            path: '/register',
            name: 'register',
            component: () => import('@/views/home/signup.vue')
        }
    ]
})
// 导出
export default router