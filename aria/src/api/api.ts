import request from "@/utils/request.ts"
import authMenuList from '@/assets/json/authMenuList.json';
import authButtonList from '@/assets/json/authButtonList.json';
// import {data} from "autoprefixer";

export function login(data: any) {
    return request({
        headers: {
            'Content-Type': 'application/json;charset=UTF-8',  //指定消息格式
        },
        url: '/login',
        method: 'post',
        data: data,
    })
}

export function uploadFile(data: FormData) {
    return request({
        url: '/aria/file/upload',
        method: 'post',
        data: data
        // transformRequest:[function (data, headers) {
        //     delete headers[ 'Content-Type']
        //     return data
        // }
        // ]
    })
}

export const uloadFile = (data: any, onUploadProgress: any) => {
    return request({
        url: '/aria/file/upload',
        method: 'post',
        data,
        onUploadProgress: onUploadProgress
    })
}
export const verifyToken = (token: string, username: string) => {
    return request({
        url: '/verify?token=' + token + '&username=' + username,
        method: 'get',
    })
}

export const getFileAllList = (token: string, username: string) => {
    return request({
        url: '/aria/file/allList?token=' + token + '&username=' + username,
        method: 'get'
    })
}


export const getAuthMenuListApi = () => {
    // return GET<MenuOptions[]>(PORT1 + `/menu/list`, {}, { loading: false });
    // 如果想让菜单变为本地数据，注释上一行代码，并引入本地 authMenuList.json 数据
    return authMenuList;
};

export const getAuthButtonListApi = () => {
    return authButtonList;
};

export const DeleteFileOne = (data: FormData) => {
    return request({
        url: "/aria/file/delete",
        method: "POST",
        data: data
    })
}

export const DownloadFileOss = (filesha: string, filename: string) => {
    return request({
        url: "/aria/file/downloadOss?filesha=" + filesha + "&filename=" + filename,
        method: "GET",
        responseType: 'blob',
    })
}


export function register(data: any) {
    return request({
        url: '/register',
        method: 'post',
        data: data,
    })
}

export const logoutApi = () => {
    // return http.post(PORT1 + `/logout`);
    return
};

export default {}