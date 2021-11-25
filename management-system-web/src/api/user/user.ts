import {AxiosRequestConfig, AxiosPromise, AxiosResponse} from 'axios'
import {BaseRequest} from "@/api/base-request";

const URL: string = "/user"

export class UserApi extends BaseRequest {

    constructor() {
        super(URL);
    }

    public getUser<T = any, R = AxiosResponse<T>, D = any>(id: number, config?: AxiosRequestConfig<D>): Promise<R> {
        return this.get(`/${id}`, config)
        // return axios.get(`/user/${id}`)
    }
}