import {useUserStore} from "@/stores/modules/user.ts";
import {useAuthStore} from "@/stores/modules/auth.ts";
import router from "@/routes";
import {LOGIN_URL} from "@/config";
import {RouteRecordRaw} from "vue-router";

const modules = import.meta.glob("@/views/**/*.vue");

export const initDynamicRouter = async () => {
    const userStore = useUserStore();
    const authStore = useAuthStore();
    try {
        // 1.获取菜单列表 && 按钮权限列表
        await authStore.getAuthMenuList();

        // 3.添加动态路由
        authStore.flatMenuListGet.forEach(item => {
            item.children && delete item.children;
            if (item.component && typeof item.component == "string") {
                item.component = modules["/src/views" + item.component + ".vue"];
            }
            if (item.meta.isFull) {
                router.addRoute(item as unknown as RouteRecordRaw);
            } else {
                router.addRoute("layout", item as unknown as RouteRecordRaw);
            }
        });

    } catch (error) {
        userStore.setToken("")
        userStore.setUserInfo("")
        router.replace(LOGIN_URL)
        return Promise.reject(error)
    }
}