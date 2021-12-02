import axios, {AxiosInstance, AxiosRequestConfig, AxiosResponse} from 'axios';
import {Request} from "@/api/request";

export interface MiddlewareFunction<D> {
    <D = any>(config: AxiosRequestConfig<D>): AxiosRequestConfig<D>
}

export abstract class BaseRequest implements Request {

    // axios实例  (单例模式)
    private static instance: AxiosInstance
    // 获取 getInstance 实例
    public static getInstance() {
        // 实例不存在时创建
        if (!BaseRequest.instance) {
            BaseRequest.instance = axios.create()
        }
        return BaseRequest.instance
    }

    // 接口请求路径
    protected url: string

    protected constructor(url: string) {
        this.url = url
        // 创建一个axios实例
        this.updateToken()
    }

    // 更新全局标头token
    public updateToken() {
        // 全局标头token
        BaseRequest.getInstance().interceptors.request.use(function (config) {
            // 是否存在token
            const token = localStorage.getItem('token')
            if (token === null) {
                return config
            }
            // 是否存在标头headers
            if (config.headers === undefined) {
                config.headers = {}
            }
            // 在标头中加入token
            config.headers['authorization'] = token
            return config;
        }, function (error) {
            return Promise.reject(error);
        });
    }

    public get<T = any, D = any>(url: string, config?: AxiosRequestConfig<D>): Promise<AxiosResponse<T>> {
        return this.baseRequest(config, this.baseInfoMiddleware('GET', url))
    }

    public delete<T = any>(url: string, config?: AxiosRequestConfig): Promise<AxiosResponse<T>> {
        return this.baseRequest(config, this.baseInfoMiddleware('DELETE', url))
    }

    public head<T = any, D = any>(url: string, config?: AxiosRequestConfig<D>): Promise<AxiosResponse<T>> {
        return this.baseRequest(config, this.baseInfoMiddleware('HEAD', url))
    }

    public options<T = any, D = any>(url: string, config?: AxiosRequestConfig<D>): Promise<AxiosResponse<T>> {
        return this.baseRequest(config, this.baseInfoMiddleware('OPTIONS', url))
    }

    public post<T = any, D = any>
    (url: string, data?: D, config?: AxiosRequestConfig<D>): Promise<AxiosResponse<T>> {
        return this.baseRequest(
            config,
            this.baseInfoMiddleware('POST', url),
            this.addDataMiddleware(data),
        )
    }

    public put<T = any, D = any>
    (url: string, data?: D, config?: AxiosRequestConfig<D>): Promise<AxiosResponse<T>> {
        return this.baseRequest(
            config,
            this.baseInfoMiddleware('PUT', url),
            this.addDataMiddleware(data),
        )
    }

    public patch<T = any, D = any>
    (url: string, data?: D, config?: AxiosRequestConfig<D>): Promise<AxiosResponse<T>> {
        return this.baseRequest(
            config,
            this.baseInfoMiddleware('PATCH', url),
            this.addDataMiddleware(data),
        )
    }

    public baseInfoMiddleware<D = any>(method: string, url: string): MiddlewareFunction<D> {
        const mainUrl = this.url
        return function <D>(config: AxiosRequestConfig<D>): AxiosRequestConfig<D> {
            return Object.assign(<AxiosRequestConfig>{
                url: mainUrl + url,
                method: method,
            }, config)
        }
    }

    public addDataMiddleware<D>(data: D): MiddlewareFunction<D> {
        return function <D>(config: AxiosRequestConfig<D>): AxiosRequestConfig<D> {
            return Object.assign({
                data: data,
            }, config)
        }
    }

    public baseRequest<T = any, D = any>
    (config?: AxiosRequestConfig<D>, ...middleware: MiddlewareFunction<D>[]): Promise<AxiosResponse<T>> {
        let conf: AxiosRequestConfig<D> = (typeof config === 'undefined') ? {} : config
        for (let i = 0; i < middleware.length; i++) {
            conf = middleware[i](conf)
        }
        return this.request(conf)
    }

    public request<T = any, R = AxiosResponse<T>, D = any>(config: AxiosRequestConfig<D>): Promise<R> {
        return BaseRequest.getInstance().request(config)
    }
}
