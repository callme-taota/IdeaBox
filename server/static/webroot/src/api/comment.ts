import { AxiosDelete, AxiosGet, AxiosPost, type MyResponse } from './axios';

export const CreateComment = async (data: any): Promise<MyResponse> => {
    const res = await AxiosPost({
        url: "/comment/",
        data: data
    });
    return res
}

export const GetComment = async (data: any): Promise<MyResponse> => {
    const res = await AxiosGet({
        url: "/comment/",
        data: data
    });
    return res
}

export const DeleteComment = async (data: any): Promise<MyResponse> => {
    const res = await AxiosDelete({
        url: "/comment/",
        data: data
    });
    return res
}