package model

import (
    "gorm.io/gorm"
    "lakego-admin/lakego/database"
)

// 管理员管理分组
type AuthGroupAccess struct {
    AdminId string `gorm:"column:admin_id;size:32;not null;index;" json:"admin_id"`
    GroupId string `gorm:"column:group_id;size:32;not null;index;" json:"group_id"`
}

func NewAuthGroupAccess() *gorm.DB {
    return database.GetDB().Model(&AuthGroupAccess{})
}
