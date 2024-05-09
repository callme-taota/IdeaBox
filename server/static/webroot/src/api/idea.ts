import { AxiosDelete, AxiosGet, AxiosPost, type MyResponse } from './axios';

export const GetIdea = async (data: any): Promise<MyResponse> => {
    const res = await AxiosGet({
        url: "/idea/ideas",
        data: data
    });
    return res
}

export const CreateIdea = async (data: any): Promise<MyResponse> => {
    const res = await AxiosPost({
        url: "/idea/",
        data: data
    });
    return res
}

export const UpdateIdeaName = async (data: any): Promise<MyResponse> => {
    const res = await AxiosPost({
        url: "/idea/name",
        data: data
    });
    return res
}

export const UpdateIdeaDescription = async (data: any): Promise<MyResponse> => {
    const res = await AxiosPost({
        url: "/idea/desc",
        data: data
    });
    return res
}

export const DeleteIdea = async (data: any): Promise<MyResponse> => {
    const res = await AxiosDelete({
        url: "/idea/",
        data: data
    });
    return res
}

export const UpdateIdea = async (data: any): Promise<MyResponse> => {
    const res = await AxiosPost({
        url: "/idea/update",
        data: data
    });
    return res
}