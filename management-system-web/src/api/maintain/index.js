import axios from 'axios'

//创建维护信息
const postMaintain = form => axios.post('/api/v1/maintain', form).then(res => res.data);

//读取维护信息
const getMaintain = id => axios.get(`/api/v1/maintain/${id}`).then(res => res.data);

export {
    postMaintain,
    getMaintain,
}