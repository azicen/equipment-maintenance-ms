import {BaseRequest} from "@/api/base-request";
import {DetailedResponse, Response} from "@/api/response-data";

const URL: string = "/login"

export interface Login {
    id: number,
    passwd: string,
}

export class LoginApi extends BaseRequest {
    constructor() {
        super(URL);
    }

    public login(id: number, passwd: string): Response<any> {
        const data = {
            id: id,
            passwd: passwd,
        }

        return new Promise<any>((resolve, reject) => {
            // post 登录接口
            this.post('', data)
                // 请求成功
                .then((res: DetailedResponse<any>) => {
                    // 如果后端回复code 200，并且存在token
                    if (res.data.code == 200 && res.headers['authorization'] !== undefined) {
                        //将 token 储存到localStorage
                        localStorage.setItem('token', res.headers['authorization'])
                        // 刷新实例Token
                        this.updateToken()
                    }
                    resolve(res.data)
                })
                // 请求失败
                .catch((error) => {
                    reject(error)
                })
        })
    }

}