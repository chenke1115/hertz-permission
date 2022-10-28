/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-28 15:55:40
 * @LastEditTime: 2022-10-28 17:08:30
 * @Description: Do not edit
 */
package model

type RolePermission struct {
	ID           int `json:"id" gorm:"type:int(11); primaryKey; autoIncrement"`
	PermissionID int `json:"permission_id" gorm:"type:int(11); not null; uniqueIndex:role_permission_unique; comment:权限ID"`
	RoleID       int `json:"role_id" gorm:"type:int(11); not null; uniqueIndex:role_permission_unique; comment:角色ID"`
}

/**
 * @description: Table name
 * @return {*}
 */
func (model RolePermission) TableName() string {
	return "role_permission"
}
