/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-28 15:55:40
 * @LastEditTime: 2022-11-01 15:34:07
 * @Description: Do not edit
 */
package model

import (
	"github.com/chenke1115/hertz-permission/internal/constant/status"
	iErrors "github.com/chenke1115/hertz-permission/internal/pkg/errors"
	gErrors "github.com/chenke1115/hertz-permission/internal/pkg/errors/gorm"

	"gorm.io/gorm"
)

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

/**
 * @description: Do create
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (model RolePermission) Create(tx *gorm.DB) (err error) {
	err = tx.Create(&model).Error
	if err != nil {
		if gErrors.IsUniqueConstraintError(err) {
			err = iErrors.Wrap(err, status.RoleParamUniqueErrCode)
		} else {
			err = iErrors.WrapCode(err, iErrors.BadRequest)
		}
	}
	return
}

/**
 * @description: Do del
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (model RolePermission) Del(tx *gorm.DB) (err error) {
	// Do del
	if err = tx.Unscoped().Delete(model).Error; err != nil {
		err = iErrors.WrapCode(err, iErrors.BadRequest)
	}

	return
}
