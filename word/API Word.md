# 权限组服务接口

| 接口描述       | 请求方法 | 接口路由                  | 请求数据                                   | 回复数据                                                     | 备注                             |
| -------------- | -------- | ------------------------- | ------------------------------------------ | :----------------------------------------------------------- | -------------------------------- |
| 创建权限组     | PUT      | /api/group                | {<br/>    name: string,<br/>}              | {<br/>    id: uint<br/>}                                     | name: 权限组名称                 |
| 获取权限组信息 | GET      | /api/group/:id            | id: uint                                   | {<br/>    name: string,<br/>}                                |                                  |
| 修改权限组信息 | POST     | /api/group/:id            | id: uint<br/>{<br/>    name: string,<br/>} | 无                                                           |                                  |
| 删除权限组     | DELETE   | /api/group/:id            | id: uint                                   | 无                                                           |                                  |
| 获取权限组列表 | GET      | /api/group/?n=:v1&page=v2 | n: uint<br/>page: uint                     | [<br/>    {<br/>        id: uint<br/>        name: string,<br/>    },<br/>    ...<br/>] | n: 每页的用户个数<br/>page: 页码 |

---

# 用户服务接口

| 接口描述           | 请求方法 | 接口路由                   | 请求数据                                                     | 回复数据                                                     | 备注                             |
| ------------------ | -------- | -------------------------- | ------------------------------------------------------------ | :----------------------------------------------------------- | -------------------------------- |
| 创建用户           | PUT      | /api/user                  | {<br/>    name: string,<br/>    passwd: string,<br/>}        | {<br/>    id: uint<br/>}                                     |                                  |
| 获取用户信息       | GET      | /api/user/:id              | id: uint                                                     | {<br/>    name: string,<br/>    status: uint,<br/>}          |                                  |
| 修改用户信息       | POST     | /api/user/:id              | id: uint<br/>{<br/>    name: string,<br/>    status: uint,<br/>} | 无                                                           |                                  |
| 删除用户           | DELETE   | /api/user/:id              | id: uint                                                     | 无                                                           |                                  |
| 获取用户信息列表   | GET      | /api/user/?n=:v1&page=v2   | n: uint<br/>page: uint                                       | [<br/>    {<br/>        id: uint<br/>        name: string,<br/>        status: uint,<br/>    },<br/>    ...<br/>] | n: 每页的用户个数<br/>page: 页码 |
| 用户加入权限组     | PUT      | /api/user/:id/group/:group | id: uint<br/>group: uint                                     | 无                                                           | id: 用户ID<br/>group: 权限组ID   |
| 获取用户权限组列表 | GET      | /api/user/:id/group        | id: uint                                                     | [<br/>    {<br/>        group: uint,<br/>    },<br/>    ...<br/>] |                                  |
| 用户移除权限组     | DELETE   | /api/user/:id/group/:group | id: uint<br/>group: uint                                     | 无                                                           |                                  |

---

# 设备数据服务接口

| 接口描述           | 请求方法 | 接口路由                        | 请求数据                                                     | 回复数据                                                     | 备注                                                         |
| ------------------ | -------- | ------------------------------- | ------------------------------------------------------------ | :----------------------------------------------------------- | ------------------------------------------------------------ |
| 创建设备           | PUT      | /api/equipment                  | {<br/>    name: string,<br/>    location: string,<br/>    status: uint,<br/>    start_date: Time,<br/>    deadline: Time,<br/>    type_id: uint,<br/>    creator_id: uint,<br/>} | {<br/>    id: uint<br/>}                                     | name: 设备名称<br/>location: 设备投放地址<br/>status: 设备状态码<br/>start_date: 开始服役时间<br/>deadline: 结束服役时间<br/>type_id: 设备类型码<br/>creator_id: 部署设备人员ID |
| 获取设备信息       | GET      | /api/equipment/:id              | id: uint                                                     | {<br/>    name: string,<br/>    location: string,<br/>    status: uint,<br/>    date: Time,<br/>    start_date: Time,<br/>    deadline: Time,<br/>    type_id: uint,<br/>    creator_id: uint,<br/>    maintainer_id: uint,<br/>} | date: 最后一次维护时间<br/>maintainer_id: 最后维护人员ID     |
| 修改设备信息       | POST     | /api/equipment/:id              | id: uint<br/>{<br/>    name: string,<br/>    location: string,<br/>    status: uint,<br/>} | 无                                                           |                                                              |
| 删除设备           | DELETE   | /api/equipment/:id              | id: uint                                                     | 无                                                           |                                                              |
| 获取设备列表       | GET      | /api/equipment/?n=:v1&page=v2   | n: uint<br/>page: uint                                       | [<br/>    {<br/>        name: string,<br/>        location: string,<br/>        status: uint,<br/>        date: Time,<br/>        start_date: Time,<br/>        deadline: Time,<br/>        type_id: uint,<br/>        creator_id: uint,<br/>        maintainer_id: uint,<br/>    },<br/>    ...<br/>] | n: 每页的用户个数<br/>page: 页码                             |
| 设备加入权限组     | PUT      | /api/equipment/:id/group/:group | id: uint<br/>group: uint                                     | 无                                                           | id: 设备ID<br/>group: 权限组ID                               |
| 获取设备权限组列表 | GET      | /api/equipment/:id/group        | id: uint                                                     | [<br/>    {<br/>        group: uint,<br/>    },<br/>    ...<br/>] |                                                              |
| 设备移除权限组     | DELETE   | /api/equipment/:id/group/:group | id: uint<br/>group: uint                                     | 无                                                           |                                                              |

---

# 维护服务接口

| 接口描述         | 请求方法 | 接口路由                     | 请求数据                                                     | 回复数据                                                     | 备注                                                         |
| ---------------- | -------- | ---------------------------- | ------------------------------------------------------------ | :----------------------------------------------------------- | ------------------------------------------------------------ |
| 创建维护信息     | PUT      | /api/maintain                | {<br/>    user_id: uint,<br/>    equipment_id: uint,<br/>    date: Time,<br/>    status: uint,<br/>    remark: string,<br/>} | {<br/>    id: uint<br/>}                                     | user_id: 维护人员ID<br/>equipment_id: 设备ID<br/>date: 维护时间<br/>status: 设备状态码<br/>remark: 备注 |
| 获取维护信息     | GET      | /api/maintain/:id            | id: uint                                                     | {<br/>    user_id: uint,<br/>    equipment_id: uint,<br/>    date: Time,<br/>    status: uint,<br/>    remark: string,<br/>} |                                                              |
| 修改维护信息     | POST     | /api/maintain/:id            | id: uint<br/>{<br/>    user_id: uint,<br/>    equipment_id: uint,<br/>    date: Time,<br/>    status: uint,<br/>    remark: string,<br/>} | 无                                                           |                                                              |
| 删除维护信息     | DELETE   | /api/maintain/:id            | id: uint                                                     | 无                                                           |                                                              |
| 获取维护信息列表 | GET      | /api/maintain/?n=:v1&page=v2 | n: uint<br/>page: uint                                       | [<br/>    {<br/>        user_id: uint,<br/>        equipment_id: uint,<br/>        date: Time,<br/>        status: uint,<br/>        remark: string,<br/>    },<br/>    ...<br/>] | n: 每页的用户个数<br/>page: 页码                             |

---























