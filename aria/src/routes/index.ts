import {createRouter, createWebHistory, createWebHashHistory} from "vue-router";
import { staticRouter, errorRouter } from "@/routes/modules/staticRouter";
import NProgress from "@/config/nprogress";
import {useUserStore} from "@/store/modules/user.ts";
import {useAuthStore} from "@/store/modules/auth.ts";


const mode = "hash";

const routerMode = {
    hash: () => createWebHashHistory(),
    history: () => createWebHistory()
};

const router = createRouter({
    history: routerMode[mode](),
    strict: false,
    routes: [...staticRouter, ...errorRouter]
});

router.beforeEach(async (to, from, next) => {
    const userStore = useUserStore();

    // 1.NProgress 开始
    NProgress.start();

    // 2.判断是访问登陆页，有 Token 就在当前页面，没有 Token 重置路由到登陆页
    if (to.path.toLocaleLowerCase() === "/login" || to.path.toLocaleLowerCase() === "/register") {
        if (userStore.token) return next(from.fullPath);
        resetRouter();
        return next();
    }
    // 3.判断是否有 Token，没有重定向到 login 页面
    if (!userStore.token) return next({ path: "/login", replace: true });

    next();

})


/**
 * @description 重置路由
 * */
export const resetRouter = () => {
    const authStore = useAuthStore();
    authStore.flatMenuListGet.forEach(route => {
        const { name } = route;
        if (name && router.hasRoute(name)) router.removeRoute(name);
    });
};

/**
 * @description 路由跳转错误
 * */
router.onError(error => {
    NProgress.done();
    console.warn("路由错误", error.message);
});

/**
 * @description 路由跳转结束
 * */
router.afterEach(() => {
    NProgress.done();
});

// 导出
export default router