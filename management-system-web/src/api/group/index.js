import axios from 'axios'

//创建权限组信息
const postGroup = form => axios.post('/api/v1/group', form).then(res => res.data);

//读取权限组信息
const getGroup = id => axios.get(`/api/v1/group/${id}`).then(res => res.data);

//修改权限组信息
const putGroup = (id, form) => axios.put(`/api/v1/group/${id}`, form).then(res => res.data);

//读取权限组列表
const getGroups = () => axios.get(`/api/v1/groups`).then(res => res.data);

export {
    postGroup,
    getGroup,
    putGroup,

    getGroups,
}