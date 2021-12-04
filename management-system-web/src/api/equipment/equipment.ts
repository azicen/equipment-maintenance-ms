import {BaseRequest} from "@/api/base-request"

import {DetailedResponse, Response} from "@/api/response-data"

// 后端路由
const URL: string = "/api/v1/equipment"

export interface GetEquipmentResponse {
    id: number,
    name: string,
    location: string,
    status: string,
    date: number,
    start_date: number,
    deadline: number,
    type_id: number,
    user_id: number,
    maintainer_id: number, // 最后维护者 ID
}

export class EquipmentApi extends BaseRequest {

    constructor() {
        super(URL);
    }

    /**
     * 创建设备
     */
    public addEquipment(name: string, location: string, status: string,
                        deadLine: number, typeId: number): Response<{ id: number }> {
        // 添加的用户数据
        const data = {
            name: name,
            location: location,
            status: status,
            deadline: deadLine,
            type_id: typeId,
        }
        return this.put("", data).then((res: DetailedResponse<{ id: number }>) => res.data)
    }

    /**
     * 根据id删除设备
     * @param id 要删除的用户id
     */
    public deleteUser(id: number): Response<any> {
        return this.delete(`/${id}`).then((res: DetailedResponse<any>) => res.data)
    }

    /**
     * 分页查询
     * @param n 起始页数
     * @param page 当前页数有多少条数据
     * @return GetEquipmentResponse[] 返回用户数据集合
     */
    public getEquipments(n: number, page: number): Response<GetEquipmentResponse[]> {
        return this.get(
            `/list`,
            {
                params: {
                    n: n,
                    page: page,
                }
            },
        ).then((res: DetailedResponse<GetEquipmentResponse[]>) => res.data)
    }

    /**
     * 获取总数
     * @return 返回设备个数
     */
    public getEquipmentListSize(): Response<{ size: number }> {
        return this.get('/list/size').then((res: DetailedResponse<{ size: number }>) => res.data)
    }
}