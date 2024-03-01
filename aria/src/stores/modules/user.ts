import {defineStore} from "pinia"
import piniaPersistConfig from "@/stores/helper/persist"
// import {UserState} from "@/stores/interface";
export interface UserState {
    token: string;
    username: string;
}
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
            this.username= username;
        }
    },
    persist: piniaPersistConfig("aria-user")
})