package interfaces

import (
    "gorm.io/gorm"
)

// 驱动接口
type Driver interface {
    // 初始化配置
    Init(map[string]interface{}) Driver

    // 设置配置
    WithConfig(map[string]interface{}) Driver

    // 获取配置
    GetConfig(...string) interface{}

    // 连接
    GetConnection() *gorm.DB

    // 关闭
    Close()
}
