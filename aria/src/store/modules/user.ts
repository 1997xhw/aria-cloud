import {defineStore} from "pinia"
import piniaPersistConfig from "@/store/helper/persist"
import "@/store/interface"

export const useUserStore = defineStore({
    id: "aria-user",
    state: (): UserState => ({
        token: "",
        username: "Aria-name"
    }),
    getters: {},
    actions: {
        // Set Token
        setToken(token: string) {
            this.token = token;
        },
        // Set setUserInfo
        setUserInfo(username: UserState["username"]) {
            this.username = username;
        }
    },
    persist: piniaPersistConfig("Aria-user")
})