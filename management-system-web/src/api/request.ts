import {AxiosRequestConfig, AxiosPromise} from 'axios'

export interface Request {
    request<T>(method: string, url: string, data?: T, config?: AxiosRequestConfig<T>): AxiosPromise
}
