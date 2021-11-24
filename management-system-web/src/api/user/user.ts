import {AxiosRequestConfig, AxiosPromise} from 'axios'
import {BaseRequest} from "@/api/base-request";

const URL: string = "/user"

export class UserApi extends BaseRequest {

    constructor() {
        super(URL);
    }

    public getUser(id: number, config?: AxiosRequestConfig) : AxiosPromise {
        return this.get(`/${id}`, null, config)
    }
}