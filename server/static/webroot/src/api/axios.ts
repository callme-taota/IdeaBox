import axios, { type AxiosResponse, type AxiosRequestConfig } from 'axios';

interface RequestParams {
    url: string;
    data?: any;
    config?: AxiosRequestConfig;
}

export interface ResponseData {
    code: number;
    data: any;
    ok: boolean;
    msg: string;
}

export interface MyResponse<T = any> extends AxiosResponse {
    ok: boolean;
    msg: string;
    data: T;
}

const instance = axios.create({
    baseURL: 'https://idea.callmetaota.fun/api',
    // baseURL: 'http://localhost:3061/api',
    timeout: 5000,
    withCredentials: true,
});

export const AxiosPost = async ({ url, data, config }: RequestParams): Promise<MyResponse<ResponseData>> => {
    try {
        const response = await instance.post(url, data, config);
        return response.data;
    } catch (error) {
        if (error instanceof Error) throw new Error(error.message);
        throw new Error("Invalid request");
    }
};

export const AxiosGet = async ({ url, data, config }: RequestParams): Promise<MyResponse<ResponseData>> => {
    try {
        const response = await instance.get(url, { params: data, ...config });
        return response.data;
    } catch (error) {
        if (error instanceof Error) throw new Error(error.message);
        throw new Error("Invalid request");
    }
};

export const AxiosDelete = async ({ url, data, config }: RequestParams): Promise<MyResponse<ResponseData>> => {
    try {
        const response = await instance.delete(url, { data, ...config });
        return response.data;
    } catch (error) {
        if (error instanceof Error) throw new Error(error.message);
        throw new Error("Invalid request");
    }
};
