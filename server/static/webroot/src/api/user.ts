import { AxiosGet, type MyResponse } from './axios';

export const GetUser = async (data: any): Promise<MyResponse> => {
    const res = await AxiosGet({
        url: "/user/",
        data: data
    });
    return res
}

export const GetUserList = async (data: any): Promise<MyResponse> => {
    const res = await AxiosGet({
        url: "/user/users",
        data: data
    });
    return res
}
