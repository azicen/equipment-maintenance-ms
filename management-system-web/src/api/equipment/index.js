import axios from 'axios'

//创建设备信息
const postEquipment = form => axios.post('/api/v1/equipment', form).then(res => res.data);

//读取设备信息
const getEquipment = id => axios.get(`/api/v1/equipment/${id}`).then(res => res.data);

//修改设备信息
const putEquipment = (id, form) => axios.put(`/api/v1/equipment/${id}`, form).then(res => res.data);

//读取设备列表
const getEquipments = id => axios.get(`/api/v1/equipments/${id}`).then(res => res.data);

export {
    postEquipment,
    getEquipment,
    putEquipment,

    getEquipments,
}