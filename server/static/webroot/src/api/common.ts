import { AxiosDelete, AxiosGet, AxiosPost, type MyResponse } from './axios';

export const Auth = async (data: any): Promise<MyResponse> => {
    const res = await AxiosPost({
        url: "/auth",
        data: data
    });
    return res
}