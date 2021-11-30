import {AxiosResponse} from "axios";

// 封装的消息模型
export interface ResponseData<D> {
    code: number,
    msg: string,
    data: D,
}

// 对详细信息预处理后的信息
export interface Response<D> extends Promise<ResponseData<D>> {
}

// 详细的回复信息
export interface DetailedResponse<D> extends AxiosResponse<ResponseData<D>> {
}