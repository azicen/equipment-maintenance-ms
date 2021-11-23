import axios, {AxiosInstance, AxiosRequestConfig, AxiosPromise, AxiosResponse} from 'axios';
import {Request} from "@/api/request";

export abstract class BaseRequest implements Request {

    // 接口请求路径
    protected url: string
    // 请求超时的毫秒数(0 表示无超时时间)
    protected timeout: number = 1000

    // axios实例
    protected instance: AxiosInstance

    protected constructor(url: string, timeout?: number) {
        this.url = url
        if (typeof timeout === 'number') {
            this.timeout = timeout
        }
        // 创建一个axios实例
        this.instance = axios.create()
    }

    public get<T>(url: string, data?: T, config?: AxiosRequestConfig<T>): AxiosPromise {
        return this.request('GET', url, data, config)
    }

    public delete<T>(url: string, data?: T, config?: AxiosRequestConfig<T>): AxiosPromise {
        return this.request('DELETE', url, data, config)
    }

    public post<T>(url: string, data?: T, config?: AxiosRequestConfig<T>): AxiosPromise {
        return this.request('POST', url, data, config)
    }

    public put<T>(url: string, data?: T, config?: AxiosRequestConfig<T>): AxiosPromise {
        return this.request('PUT', url, data, config)
    }

    public head<T>(url: string, data?: T, config?: AxiosRequestConfig<T>): AxiosPromise {
        return this.request('HEAD', url, data, config)
    }

    public options<T>(url: string, data?: T, config?: AxiosRequestConfig<T>): AxiosPromise {
        return this.request('OPTIONS', url, data, config)
    }

    public patch<T>(url: string, data?: T, config?: AxiosRequestConfig<T>): AxiosPromise {
        return this.request('PATCH', url, data, config)
    }

    public request<T>(method: string, url: string, data?: T, config?: AxiosRequestConfig<T>): AxiosPromise {
        let conf: AxiosRequestConfig<T> = {}
        if (typeof config !== undefined) {
            conf = <AxiosRequestConfig<T>>config
        }
        if (typeof data !== undefined) {
            Object.assign(<AxiosRequestConfig>{data: data}, config)
        }
        Object.assign(<AxiosRequestConfig>{
            url: this.url + url,
            method: method,
            timeout: this.timeout,
        }, config)
        return this.instance.request(conf)
    }
}
