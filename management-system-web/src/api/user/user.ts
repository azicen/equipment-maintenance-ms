import {AxiosRequestConfig, AxiosResponse} from 'axios'
import {BaseRequest} from "@/api/base-request";

const URL: string = "/user"

export class UserApi extends BaseRequest {

    constructor() {
        super(URL);
    }

    public addUser<T = any, R = AxiosResponse<T>, D = any>
    (id: number, data: D, config?: AxiosRequestConfig<D>): Promise<R> {
        return this.post("", data, config)
    }

    public getUser<T = any, R = AxiosResponse<T>, D = any>(id: number, config?: AxiosRequestConfig<D>): Promise<R> {
        return this.get(`/${id}`, config)
    }
}