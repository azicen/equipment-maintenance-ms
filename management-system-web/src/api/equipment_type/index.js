import axios from 'axios'

//创建设备类型信息
const postEquipmentType = form => axios.post('/api/v1/equipment_type', form).then(res => res.data);

//读取设备类型信息
const getEquipmentType = id => axios.get(`/api/v1/equipment_type/${id}`).then(res => res.data);

//修改设备类型信息
const putEquipmentType = (id, form) => axios.put(`/api/v1/equipment_type/${id}`, form).then(res => res.data);

//删除设备类型信息
const delEquipmentType = id => axios.delete(`/api/v1/equipment_type/${id}`).then(res => res.data);

//读取设备类型列表
const getEquipmentTypes = () => axios.get(`/api/v1/equipment_types`).then(res => res.data);

export {
    postEquipmentType,
    getEquipmentType,
    putEquipmentType,
    delEquipmentType,

    getEquipmentTypes,
}