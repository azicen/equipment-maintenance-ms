import axios from 'axios'

//创建维护信息
const postMaintain = form => axios.post('/api/v1/maintain', form).then(res => res.data);

//读取维护信息
const getMaintain = id => axios.get(`/api/v1/maintain/${id}`).then(res => res.data);

//删除维护信息
const delMaintain = id => axios.delete(`/api/v1/maintain/${id}`).then(res => res.data);

//读取维护信息列表
const getMaintains = (id, userId) => axios.get(`/api/v1/maintains?id=${id}&user_id=${userId}`).then(res => res.data);

export {
    postMaintain,
    getMaintain,
    delMaintain,

    getMaintains,
}