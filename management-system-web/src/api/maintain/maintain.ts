import {BaseRequest} from "@/api/base-request"

import {DetailedResponse, Response} from "@/api/response-data"

// 后端路由
const URL: string = "/api/v1/maintain"

export class MaintainApi extends BaseRequest {

    constructor() {
        super(URL);
    }

    /**
     * 添加维护记录
     */
    public addMaintain(equipment_id: number, date: number, status: string, remark: string,): Response<{ id: number }> {
        // 添加的用户数据
        const data = {
            equipment_id: equipment_id,
            date: date,
            status: status,
            remark: remark,
        }
        return this.put("", data).then((res: DetailedResponse<{ id: number }>) => res.data)
    }
}