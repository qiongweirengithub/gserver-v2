web接口api
### 域名
http://test.gserver.v2/
http://gserver.v2/

### 失败统一返回
```
{
    status：false
    code: 错误码
    errmsg：错误原因
}
```
### 注册
- uri:     /register/
- param

| name | type | remark |
| ---- | ---- | ---- |
| account | string | 帐号,唯一 |
|pwd | string | 密码|

- 返回值
```
{
    status：true
    data：{
        name: string 帐号名称
    }
}
```
### 登录
- uri:     /login/
- param

| name | type | remark |
| ---- | ---- | ---- |
| account | string | 帐号,唯一 |
|pwd | string | 密码|

- 返回值
```
{
    status：true
    data：{
        name: string 帐号名称
        token：string 登录token
    }
}
```

### 登录
- uri:     /logout/
- param

| name | type | remark |
| ---- | ---- | ---- |
| token | string | 登录token|


- 返回值
```
{
    status：true
    data：{
        name: string 帐号名称
    }
}
```

### 查询网关信息
- uri:     /gateinfo/load
- param

| name | type | remark |
| ---- | ---- | ---- |
| token | string | 登录token|

- 返回值
```
{
    status：true
    data：{
        gates: [g1,g2]  // 网关ip列表
    }
}
```