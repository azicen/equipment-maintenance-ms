# 用户前后端交互接口

### 创建用户

路由地址：/api/user

**请求数据**

```json
{
    "name": "string",
    "passwd": "string",
}
```



### 获取用户基本信息

路由地址：/api/user/:id

**回复数据**

```json
{
	"name": "string",
	"status": "uint",
	"group": "uint",
}
```



### 修改用户基本信息

路由地址：/api/user/:id

**请求数据**

```json
{
    "name": "string",
    "status": "uint",
    "group": "uint",
}
```



### 删除用户

路由地址：/api/user/:id



### 获取用户基本信息列表

路由地址：/api/users/?begin=:begin&end=:end

begin:	 uint	// 起始id

end:		uint	// 结束id

**回复数据**

```json
[
    {
        "name": "string",
        "status": "uint",
        "group": "uint",
	},
    ...
]

```

