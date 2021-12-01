import axios, {AxiosInstance, AxiosRequestConfig, AxiosResponse} from 'axios';
import {Request} from "@/api/request";

export interface ModifyFunction<D> {
    <D = any>(config: AxiosRequestConfig<D>): AxiosRequestConfig<D>
}

export abstract class BaseRequest implements Request {

    // 接口请求路径
    protected url: string

    // axios实例
    protected instance: AxiosInstance

    protected constructor(url: string) {
        this.url = url
        // 创建一个axios实例
        this.instance = axios.create()
    }

    public get<T = any, D = any>(url: string, config?: AxiosRequestConfig<D>): Promise<AxiosResponse<T>> {
        return this.baseRequest(config, this.modifyBase('GET', url))
    }

    public delete<T = any>(url: string, config?: AxiosRequestConfig): Promise<AxiosResponse<T>> {
        return this.baseRequest(config, this.modifyBase('DELETE', url))
    }

    public head<T = any, D = any>(url: string, config?: AxiosRequestConfig<D>): Promise<AxiosResponse<T>> {
        return this.baseRequest(config, this.modifyBase('HEAD', url))
    }

    public options<T = any, D = any>(url: string, config?: AxiosRequestConfig<D>): Promise<AxiosResponse<T>> {
        return this.baseRequest(config, this.modifyBase('OPTIONS', url))
    }

    public post<T = any, D = any>
    (url: string, data?: D, config?: AxiosRequestConfig<D>): Promise<AxiosResponse<T>> {
        return this.baseRequest(
            config,
            this.modifyBase('POST', url),
            this.modifyData(data),
        )
    }

    public put<T = any, D = any>
    (url: string, data?: D, config?: AxiosRequestConfig<D>): Promise<AxiosResponse<T>> {
        return this.baseRequest(
            config,
            this.modifyBase('PUT', url),
            this.modifyData(data),
        )
    }

    public patch<T = any, D = any>
    (url: string, data?: D, config?: AxiosRequestConfig<D>): Promise<AxiosResponse<T>> {
        return this.baseRequest(
            config,
            this.modifyBase('PATCH', url),
            this.modifyData(data),
        )
    }

    public modifyBase<D = any>(method: string, url: string): ModifyFunction<D> {
        const mainUrl = this.url
        return function <D>(config: AxiosRequestConfig<D>): AxiosRequestConfig<D> {
            return Object.assign(<AxiosRequestConfig>{
                url: mainUrl + url,
                method: method,
            }, config)
        }
    }

    public modifyData<D>(data: D): ModifyFunction<D> {
        return function <D>(config: AxiosRequestConfig<D>): AxiosRequestConfig<D> {
            return Object.assign({
                data: data,
            }, config)
        }
    }

    public baseRequest<T = any, D = any>
    (config?: AxiosRequestConfig<D>, ...modifies: ModifyFunction<D>[]): Promise<AxiosResponse<T>> {
        let conf: AxiosRequestConfig<D> = (typeof config === 'undefined') ? {} : config
        for (let i = 0; i < modifies.length; i++) {
            conf = modifies[i](conf)
        }
        return this.request(conf)
    }

    public request<T = any, R = AxiosResponse<T>, D = any>(config: AxiosRequestConfig<D>): Promise<R> {
        return this.instance.request(config)
    }
}
