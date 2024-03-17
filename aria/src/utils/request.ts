import axios, { AxiosRequestConfig } from 'axios'
import {showMessage} from "@/utils/status.ts";
import {ElMessage} from "element-plus";
import router from "@/routes";
import {LOGIN_URL} from "@/config"; // 引入axios

const instacne = axios.create({
    timeout: 5000,
});
//我们可以更方便地在项目中使用axios，并且统一了响应错误的处理方式和请求url的规范化
interface OptionsTypes extends AxiosRequestConfig {//对axios库的请求配置进行了扩展,添加自定义属性
    ifHandleError?: boolean//一个布尔型参数，表示是否需要处理响应错误
    prefix?: string//个字符串型参数，表示请求接口的前缀
}
const request = async function (opt: OptionsTypes): Promise<any> {
    let options: OptionsTypes = {
        method: 'get',
        ifHandleError: true,
        prefix: '',
        ...opt,
        // baseURL: '/dev-api'//因为下面`vite.confing.ts`中以代理
        baseURL: import.meta.env.VITE_NODE_ENV === 'dev' ? '/dev-api' : '/prod-api'//因为下面`vite.confing.ts`中以代理

    }
    try {
        const res: any = await instacne(options)
        //对返回结果data进行操作判断
        return res
    } catch (err) {
        return err
    }
}
// 请求拦截器
instacne.interceptors.request.use(
    config => {
        // 添加token等操作
        if (config.method === 'get') {
            config.headers = Object.assign(
                {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json; charset=UTF-8'
                }
            )
        } else if (config.method === 'post') {
            // config.headers = Object.assign(
            //     {
            //         'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
            //         // 'Content-Type': 'application/json; charset=UTF-8'
            //     }
            // )
        }
        return Promise.resolve(config)
    },
    error => {
        return Promise.reject(error);
    }
);
// 响应拦截器
instacne.interceptors.response.use(
    response => {
        // 数据处理等操作
        return response;
    },
    error => {
        const { response } = error;
        if (response) {
            // 请求已发出，但是不在2xx的范围
            ElMessage.warning(showMessage(response.status)); // 传入响应码，匹配响应码对应信息

            return Promise.reject(response.data);
        } else {
            ElMessage.warning("网络连接异常,请稍后再试!");
        }
    }
);
export default request;
