## lakego-admin 后台管理系统


### 项目介绍

*  `lakego-admin` 是基于 `gin` 的后台开发框架，完全api接口化，适用于前后端分离的项目
*  基于 `JWT` 的用户登录态管理
*  权限判断基于 `go-casbin` 的 `RBAC` 授权
*  使用 `Swagger` 作为 API 文档管理
*  本项目为 `后台api服务`


### 环境要求

 - Go >= 1.18
 - Gorm >= v1.21.10
 - Redis


### 截图预览

<table>
    <tr>
        <td width="50%">
            <center>
                <img alt="登录" src="https://user-images.githubusercontent.com/24578855/151009218-d544fcb1-973d-42e4-a3b0-1ae72ea6a088.png" />
            </center>
        </td>
        <td width="50%">
            <center>
                <img alt="控制台" src="https://user-images.githubusercontent.com/24578855/151192881-72510e1d-88db-4db3-b730-a741fd981fd7.png" />
            </center>
        </td>
        <td width="50%">
            <center>
                <img alt="操作日志" src="https://user-images.githubusercontent.com/24578855/168456599-8401a6ef-9b8a-4fd8-bb30-3978bf4b0ec7.png" />
            </center>
        </td>
    </tr>
    
    <tr>
        <td width="50%">
            <center>
                <img alt="管理员" src="https://user-images.githubusercontent.com/24578855/168456604-c4dddd71-4b70-496b-ba2e-752e69932571.png" />
            </center>
        </td>
        <td width="50%">
            <center>
                <img alt="用户组" src="https://user-images.githubusercontent.com/24578855/168456611-1f7fcdb6-e2af-4f8f-8572-227cd4096b61.png" />
            </center>
        </td>
        <td width="50%">
            <center>
                <img alt="权限路由" src="https://user-images.githubusercontent.com/24578855/168456618-c4ab5e26-7e89-4bb5-bb25-3299a5a70c3d.png" />
            </center>
        </td>
    </tr>
</table>

更多截图
[Lakego Admin 后台截图](https://github.com/deatil/lakego-admin/issues/1)


### 安装步骤

1. 首先克隆项目到本地

```
git clone https://github.com/deatil/lakego-admin.git
```

2. 然后配置数据库等相关配置，配置位置

```
/config
```

3. 最后运行下面的命令安装系统

```go
go run main.go lakego-admin:install
```

4. 运行下面的命令创建附件软链接

```go
go run main.go lakego:storage-link
```

5. 权限规则导入。导入的权限规则需要重新设置层级关系和名称内容

```go
go run main.go lakego-admin:import-route
```

6. 运行测试

```go
go run main.go
```

6. 后台登录账号及密码：`admin` / `123456`


### 特别鸣谢

感谢以下的项目,排名不分先后

 - github.com/gin-gonic/gin

 - gorm.io/gorm

 - github.com/golang-jwt/jwt

 - github.com/casbin/casbin

 - github.com/spf13/cobra

 - github.com/go-redis/redis


### 开源协议

*  `lakego-admin` 遵循 `Apache2` 开源协议发布，在保留本系统版权的情况下提供个人及商业免费使用。


### 版权

*  该系统所属版权归 deatil(https://github.com/deatil) 所有。
