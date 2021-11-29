import {AxiosRequestConfig, AxiosResponse} from 'axios'

export interface Request {
    request<T = any, R = AxiosResponse<T>, D = any>(config: AxiosRequestConfig<D>): Promise<R>
}
