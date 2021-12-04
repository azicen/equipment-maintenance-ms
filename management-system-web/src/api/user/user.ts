import {BaseRequest} from "@/api/base-request"

import {DetailedResponse, Response} from "@/api/response-data"

// 后端路由
const URL: string = "/api/v1/user"

export interface GetUserResponse {
    id: number,
    name: string,
    status: boolean,
}

export class UserApi extends BaseRequest {

    constructor() {
        super(URL);
    }

    // 很骚，但不能用
    // public addUser<R = AxiosResponse<{ id: number }>>(id: number, data: AddUserData): Promise<R> {
    //     return this.put("", data).then(res => res.data)
    // }

    /**
     * 添加用户
     * @param name 用户姓名
     * @param status 账号状态
     * @return { id: number } 创建成功后的用户ID
     */
    public addUser(name: string, status: boolean): Response<{ id: number }> {
        // 添加的用户数据
        const data = {
            name: name,
            status: status,
        }
        return this.put("", data).then((res: DetailedResponse<{ id: number }>) => res.data)
    }

    /**
     * 根据id查找用户
     * @param id 用户id
     * @return GetUserResponse 返回用户数据
     */
    public getUser(id: number): Response<{ name: string, status: boolean }> {
        return this.get(`/${id}`)
            // 消息预处理
            .then((res: DetailedResponse<{ name: string, status: boolean }>) => res.data)
    }

    /**
     * 根据id删除用户
     * @param id 要删除的用户id
     */
    public deleteUser(id: number): Response<any> {
        return this.delete(`/${id}`).then((res: DetailedResponse<any>) => res.data)
    }

    /**
     * 修改用户数据
     * @param id 需要修改的用户id,由路由/api/user/id获得
     * @param name 修改后的姓名
     * @param status 修改后的状态码
     */
    public updateUser(id: number, name: string, status: boolean): Response<any> {
        const data = {
            name: name,
            status: status,
        }
        return this.post(`/${id}`, data).then((res: DetailedResponse<any>) => res.data)
    }

    /**
     * 分页查询
     * @param n 起始页数
     * @param page 当前页数有多少条数据
     * @return GetUserResponse[] 返回用户数据集合
     */
    public getUsers(n: number, page: number): Response<GetUserResponse[]> {
        return this.get(
            `/list`,
            {
                params: {
                    n: n,
                    page: page,
                }
            },
        ).then((res: DetailedResponse<GetUserResponse[]>) => res.data)
    }

    /**
     * 将用户添加到权限组
     * @param userId 用户id
     * @param groupId 权限组id
     */
    public addUserGroup(userId: number, groupId: number): Response<any> {
        return this.put(`/${userId}/group/${groupId}`).then((res: DetailedResponse<any>) => res.data)
    }

    /**
     * 获取用户权限组
     * @param id 需要访问的用户id
     * @return { group: number }[] 返回用户所在的所有权限组
     */
    public getUserGroup(id: number): Response<{ group: number }[]> {
        return this.get(`/${id}`).then((res: DetailedResponse<{ group: number }[]>) => res.data)
    }

    /**
     * 删除指定用户的权限组
     * @param userId 用户id
     * @param groupId 权限组id
     */
    public deleteUserGroup(userId: number, groupId: number): Response<any> {
        return this.delete(`/${userId}/group/${groupId}`).then((res: DetailedResponse<any>) => res.data)
    }

    /**
     * 获取用户总数
     * @return 返回用户个数
     */
    public getUserListSize(): Response<{ size: number }> {
        return this.get('/list/size').then((res: DetailedResponse<{ size: number }>) => res.data)
    }
}