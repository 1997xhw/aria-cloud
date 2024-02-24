import request from "@/utils/request.ts"
export function login(data :any) {
    return request({
        headers: {
            'Content-Type': 'application/json;charset=UTF-8',  //指定消息格式
        },
        url: '/login',
        method: 'post',
        data: data,
    })
}

export default {

}