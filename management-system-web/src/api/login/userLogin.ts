import {BaseRequest} from "@/api/base-request";
import {DetailedResponse, Response} from "@/api/response-data";

const URL: string ="/login"

export interface UserLogin{
    id: number,
    passwd: string,
}

export class UserLogin extends BaseRequest{
    constructor() {
        super(URL);
    }

    public Login(id:number,passwd:string):Response<any>{
        const data={
            id: id,
            passwd: passwd,
        }
        return this.post("",data).then((res: DetailedResponse<any>) => res.data)
    }

}