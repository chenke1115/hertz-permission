/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-28 15:55:40
 * @LastEditTime: 2022-11-14 15:52:24
 * @Description: Do not edit
 */
package model

import (
	"github.com/chenke1115/go-common/functions/array"
	iErrors "github.com/chenke1115/hertz-common/pkg/errors"
	gErrors "github.com/chenke1115/hertz-common/pkg/errors/gorm"
	"github.com/chenke1115/hertz-permission/internal/constant/status"

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
 * @description: List
 * @return {*}
 */
func (model RolePermission) List() (list []RolePermission, err error) {
	err = GetDB().Model(&RolePermission{}).Where(&model).Scan(&list).Error
	return
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

/**
 * @description: Del batch
 * @param {*gorm.DB} tx
 * @param {[]int} ids
 * @return {*}
 */
func (model RolePermission) DelBatch(tx *gorm.DB, ids []int) (err error) {
	// Do del
	if err = tx.Unscoped().Where("id in (?)", ids).Delete(model).Error; err != nil {
		err = iErrors.WrapCode(err, iErrors.BadRequest)
	}

	return
}

/**
 * @description: Not In
 * @param {*gorm.DB} tx
 * @param {[]int} roleIDs
 * @return {*}
 */
func (model RolePermission) NotIn(tx *gorm.DB, perIDs []int) (ids []int, err error) {
	err = tx.Select("id").Model(&RolePermission{}).Where(&model).
		Where("permission_id not in (?)", perIDs).
		Scan(&ids).Error
	return
}

/**
 * @description: Is exist
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (model RolePermission) IsExist(tx *gorm.DB) bool {
	err := tx.Where(&model).First(&RolePermission{}).Error
	return err == nil
}

/**
 * @description: Do save
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (model RolePermission) Save(tx *gorm.DB) (err error) {
	if !IsEnablePermission(tx, model.PermissionID) {
		return iErrors.New(status.PermissionIsUnableErrorCode)
	}

	if !model.IsExist(tx) {
		err = tx.Create(&model).Error
		if err != nil {
			if gErrors.IsUniqueConstraintError(err) {
				err = iErrors.Wrap(err, status.RoleParamUniqueErrCode)
			} else {
				err = iErrors.WrapCode(err, iErrors.BadRequest)
			}
			return
		}
	}

	return
}

/**
 * @description: check permission_id is exist
 * @param {int} permission_id
 * @return {*}
 */
func IsExistPermissionID(permission_id int) bool {
	err := GetDB().Model(&RolePermission{}).
		First(&RolePermission{}, "permission_id = ?", permission_id).Error
	return err == nil
}

/**
 * @description: Do not edit
 * @param {[]int} ids
 * @return {*}
 */
func GetPermissionsByRoleIDs(ids []int) (permissions []string, err error) {
	var perIDs []int
	err = GetDB().Model(&RolePermission{}).
		Select("permission_id").
		Where("role_id in (?)", ids).
		Scan(&perIDs).Error
	if len(perIDs) < 1 {
		return
	}

	permissions, err = GetPermissionKeysByIDs(array.UniqueArray(perIDs))
	return
}
