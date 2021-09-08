package controller

import (
    "strings"
    "encoding/json"
    "github.com/gin-gonic/gin"

    "lakego-admin/lakego/collection"
    "lakego-admin/lakego/support/cast"
    "lakego-admin/lakego/facade/config"

    "lakego-admin/admin/model"
    "lakego-admin/admin/model/scope"
    "lakego-admin/admin/auth/admin"
    adminValidate "lakego-admin/admin/validate/admin"
    adminRepository "lakego-admin/admin/repository/admin"
)

/**
 * 管理员
 *
 * @create 2021-9-2
 * @author deatil
 */
type Admin struct {
    Base
}

/**
 * 列表
 */
func (control *Admin) Index(ctx *gin.Context) {
    // 模型
    adminModel := model.NewAdmin()

    // 排序
    order := ctx.DefaultQuery("order", "id__DESC")
    orders := strings.SplitN(order, "__", 2)
    if orders[0] != "id" ||
        orders[0] != "name" ||
        orders[0] != "last_active" ||
        orders[0] != "add_time" {
        orders[0] = "id"
    }

    adminModel = adminModel.Order(orders[0] + " " + orders[1])

    // 搜索条件
    searchword := ctx.DefaultQuery("searchword", "")
    if searchword != "" {
        searchword = "%" + searchword + "%"

        adminModel = adminModel.
            Or("name LIKE ?", searchword).
            Or("nickname LIKE ?", searchword).
            Or("email LIKE ?", searchword)
    }

    // 时间条件
    startTime := ctx.DefaultQuery("start_time", "")
    if startTime != "" {
        adminModel = adminModel.Where("add_time >= ?", control.FormatDate(startTime))
    }

    endTime := ctx.DefaultQuery("end_time", "")
    if endTime != "" {
        adminModel = adminModel.Where("add_time <= ?", control.FormatDate(endTime))
    }

    status := control.SwitchStatus(ctx.DefaultQuery("status", ""))
    if status != -1 {
        adminModel = adminModel.Where("status = ?", status)
    }

    // 分页相关
    start := ctx.DefaultQuery("start", "0")
    limit := ctx.DefaultQuery("limit", "10")

    newStart := cast.ToInt(start)
    newLimit := cast.ToInt(limit)

    adminModel = adminModel.
        Offset(newStart).
        Limit(newLimit)

    list := make([]map[string]interface{}, 0)

    // 列表
    adminModel = adminModel.
        Select([]string{
            "id", "name", "nickname",
            "email", "avatar",
            "is_root", "status",
            "last_active", "last_ip",
            "update_time", "update_ip",
            "add_time", "add_ip",
        }).
        Find(&list)

    var total int64

    // 总数
    err := adminModel.
        Offset(-1).
        Limit(-1).
        Count(&total).
        Error
    if err != nil {
        control.Error(ctx, "获取失败")
        return
    }

    // 数据输出
    control.SuccessWithData(ctx, "获取成功", gin.H{
        "start": start,
        "limit": limit,
        "total": total,
        "list": list,
    })
}

/**
 * 详情
 */
func (control *Admin) Detail(ctx *gin.Context) {
    id := ctx.Param("id")
    if id == "" {
        control.Error(ctx, "账号ID不能为空")
        return
    }

    var info = model.Admin{}

    // 附件模型
    err := model.NewAdmin().
        Where("id = ?", id).
        Preload("Groups").
        First(&info).
        Error
    if err != nil {
        control.Error(ctx, "账号不存在")
        return
    }

    // 结构体转map
    data, _ := json.Marshal(&info)
    adminData := map[string]interface{}{}
    json.Unmarshal(data, &adminData)

    newInfoGroups:= collection.Collect(adminData["Groups"]).
        Select("id", "parentid", "title", "description").
        ToMapArray()

    avatar := model.AttachmentUrl(adminData["avatar"].(string))

    newInfo := collection.Collect(adminData).
        Only([]string{
            "id", "name", "nickname", "email",
            "is_root", "status",
            "last_active", "last_ip",
            "update_time", "update_ip",
            "add_time", "add_ip",
        }).
        ToMap()

    newInfo["groups"] = newInfoGroups
    newInfo["avatar"] = avatar

    // 数据输出
    control.SuccessWithData(ctx, "获取成功", newInfo)
}

/**
 * 管理员权限
 */
func (control *Admin) Rules(ctx *gin.Context) {
    id := ctx.Param("id")
    if id == "" {
        control.Error(ctx, "账号ID不能为空")
        return
    }

    var info = model.Admin{}

    // 附件模型
    err := model.NewAdmin().
        Scopes(scope.AdminWithAccess(ctx, []string{})).
        Where("id = ?", id).
        Preload("Groups").
        First(&info).
        Error
    if err != nil {
        control.Error(ctx, "账号不存在")
        return
    }

    // 结构体转map
    data, _ := json.Marshal(&info)
    adminData := map[string]interface{}{}
    json.Unmarshal(data, &adminData)

    groupids := collection.Collect(adminData["Groups"]).
        Pluck("id").
        ToStringArray()

    rules := adminRepository.GetRules(groupids)

    // 数据输出
    control.SuccessWithData(ctx, "获取成功", gin.H{
        "list": rules,
    })
}

/**
 * 删除
 */
func (control *Admin) Delete(ctx *gin.Context) {
    id := ctx.Param("id")
    if id == "" {
        control.Error(ctx, "账号ID不能为空")
        return
    }

    adminId, _ := ctx.Get("admin_id")
    if id == adminId.(string) {
        control.Error(ctx, "你不能删除自己的账号")
        return
    }

    result := map[string]interface{}{}

    // 附件模型
    err := model.NewAdmin().
        Where("id = ?", id).
        First(&result).
        Error
    if err != nil || len(result) < 1 {
        control.Error(ctx, "账号信息不存在")
        return
    }

    authAdminId := config.New("auth").GetString("Auth.AdminId")
    if authAdminId == adminId.(string) {
        control.Error(ctx, "当前账号不能被删除")
        return
    }

    // 删除
    err2 := model.NewAdmin().
        Delete(&model.Admin{
            ID: id,
        }).
        Error
    if err2 != nil {
        control.Error(ctx, "账号删除失败")
        return
    }

    // 数据输出
    control.Success(ctx, "账号删除成功")
}

/**
 * 添加
 */
func (control *Admin) Create(ctx *gin.Context) {
    // 接收数据
    post := make(map[string]interface{})
    ctx.BindJSON(&post)

    validateErr := adminValidate.Login(post)
    if validateErr != "" {
        control.Error(ctx, validateErr)
        return
    }

    status := 0
    if post["status"].(string) == 1 {
        status = 1
    }

    // 附件模型
    result := map[string]interface{}{}
    err := model.NewAdmin().
        Where("name = ?", post["name"].(string)).
        Or("email = ?", post["email"].(string))
        First(&result).
        Error
    if !(err != nil || len(result) < 1) {
        control.Error(ctx, "邮箱或者账号已经存在")
        return
    }

    insertData := model.Admin{
        Name: post["name"].(string),
        Nickname: post["nickname"].(string),
        Email: post["email"].(string),
        Introduce: post["introduce"].(string),
        Status: status,
    }

    adminInfo := ctx.Get("admin").(*admin.Admin)
    groupChildrenIds := adminInfo.GetGroupChildrenIds()
    if len(groupChildrenIds) < 1 {
        control.Error(ctx, "当前账号不能创建子账号")
        return
    }

    err := model.NewDB().
        Create(&insertData).
        Error
    if err != nil {
        control.Error(ctx, "添加账号失败")
        return
    }

    model.NewDB().Create(&model.AuthGroupAccess{
        AdminId: insertData.ID,
        GroupId: groupChildrenIds[0],
    })

    // 数据输出
    control.SuccessWithData(ctx, "添加账号成功", gin.H{
        "id": insertData.ID,
    })
}

/**
 * 更新
 */
func (control *Admin) Update(ctx *gin.Context) {
    id := ctx.Param("id")
    if id == "" {
        control.Error(ctx, "账号ID不能为空")
        return
    }

    adminId, _ := ctx.Get("admin_id")
    if id == adminId.(string) {
        control.Error(ctx, "你不能修改自己的账号")
        return
    }

    // 查询
    result := map[string]interface{}{}
    err := model.NewAdmin().
        Where("id = ?", id).
        First(&result).
        Error
    if err != nil || len(result) < 1 {
        control.Error(ctx, "账号信息不存在")
        return
    }

    // 接收数据
    post := make(map[string]interface{})
    ctx.BindJSON(&post)

    validateErr := adminValidate.Update(post)
    if validateErr != "" {
        control.Error(ctx, validateErr)
        return
    }

    status := 0
    if post["status"].(string) == 1 {
        status = 1
    }

    // 链接db
    db := model.NewDB()

    // 验证
    result2 := map[string]interface{}{}
    err2 := model.NewAdmin().
        Where(db.Where("id = ?", id).Where("name = ?", post["name"].(string))).
        Or(db.Where("id = ?", id).Where("email = ?", post["email"].(string))).
        First(&result2).
        Error
    if !(err2 != nil || len(result2) < 1) {
        control.Error(ctx, "管理员账号或者邮箱已经存在")
        return
    }

    err3 := model.NewAdmin().
        Where("id = ?", id).
        Updates(map[string]interface{}{
            "name": post["name"].(string),
            "nickname": post["nickname"].(string),
            "email": post["email"].(string),
            "introduce": post["introduce"].(string),
            "status": status,
        }).
        Error
    if err3 != nil {
        control.Error(ctx, "账号修改失败")
        return
    }

    // 数据输出
    control.Success(ctx, "账号修改成功")
}

/**
 * 修改头像
 */
func (control *Admin) UpdateAvatar(ctx *gin.Context) {

    // 数据输出
    control.SuccessWithData(ctx, "获取成功", gin.H{})
}

/**
 * 修改密码
 */
func (control *Admin) UpdatePasssword(ctx *gin.Context) {

    // 数据输出
    control.SuccessWithData(ctx, "获取成功", gin.H{})
}

/**
 * 授权
 */
func (control *Admin) Access(ctx *gin.Context) {

    // 数据输出
    control.SuccessWithData(ctx, "获取成功", gin.H{})
}

/**
 * 启用
 */
func (control *Admin) Enable(ctx *gin.Context) {

    // 数据输出
    control.SuccessWithData(ctx, "获取成功", gin.H{})
}

/**
 * 禁用
 */
func (control *Admin) Disable(ctx *gin.Context) {

    // 数据输出
    control.SuccessWithData(ctx, "获取成功", gin.H{})
}

/**
 * 退出
 */
func (control *Admin) Logout(ctx *gin.Context) {

    // 数据输出
    control.SuccessWithData(ctx, "获取成功", gin.H{})
}

