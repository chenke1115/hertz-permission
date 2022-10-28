/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-28 16:15:00
 * @LastEditTime: 2022-10-28 17:08:19
 * @Description: Do not edit
 */
package model

import "time"

type UserRole struct {
	ID        int       `json:"id" gorm:"type:int(11); primaryKey; autoIncrement"`
	UID       int       `json:"uid" gorm:"type:int(11); unsigned; not null; index; uniqueIndex:user_role_unique; comment:用户ID"`
	RoleID    int       `json:"role_id" gorm:"type:int(11); not null;index; uniqueIndex:user_role_unique; comment:角色ID"`
	CreatedAt time.Time `json:"create_at" gorm:"type:timestamp; default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"update_at" gorm:"type:timestamp; default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

/**
 * @description: Table name
 * @return {*}
 */
func (model UserRole) TableName() string {
	return "user_role"
}
