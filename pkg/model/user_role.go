/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-28 16:15:00
 * @LastEditTime: 2022-11-14 15:30:23
 * @Description: Do not edit
 */
package model

import (
	"time"

	"github.com/chenke1115/hertz-permission/internal/constant/status"
	iErrors "github.com/chenke1115/hertz-permission/internal/pkg/errors"
	gErrors "github.com/chenke1115/hertz-permission/internal/pkg/errors/gorm"

	"gorm.io/gorm"
)

type UserRole struct {
	ID        int       `json:"id" gorm:"type:int(11); primaryKey; autoIncrement"`
	UID       int       `json:"uid" gorm:"type:int(11); unsigned; not null; index; uniqueIndex:user_role_unique; comment:用户ID"`
	RoleID    int       `json:"role_id" gorm:"type:int(11); not null;index; uniqueIndex:user_role_unique; comment:角色ID"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp; default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp; default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

/**
 * @description: Table name
 * @return {*}
 */
func (model UserRole) TableName() string {
	return "user_role"
}

/**
 * @description: Do create
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (model UserRole) Create(tx *gorm.DB) (err error) {
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
func (model UserRole) Del(tx *gorm.DB) (err error) {
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
func (model UserRole) DelBatch(tx *gorm.DB, ids []int) (err error) {
	// Do del
	if err = tx.Unscoped().Where("id in (?)", ids).Delete(model).Error; err != nil {
		err = iErrors.WrapCode(err, iErrors.BadRequest)
	}

	return
}

/**
 * @description: Is exist
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (model UserRole) IsExist(tx *gorm.DB) bool {
	var userRole UserRole
	err := tx.Where(&model).First(&userRole).Error
	return err == nil
}

/**
 * @description: Not In
 * @param {*gorm.DB} tx
 * @param {[]int} roleIDs
 * @return {*}
 */
func (model UserRole) NotIn(tx *gorm.DB, roleIDs []int) (ids []int, err error) {
	err = tx.Select("id").Model(&UserRole{}).Where(&model).
		Where("role_id not in (?)", roleIDs).
		Scan(&ids).Error
	return
}

/**
 * @description: Do save
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (model UserRole) Save(tx *gorm.DB) (err error) {
	if !IsEnableRole(tx, model.RoleID) {
		return iErrors.New(status.RoleStatusErrorCode)
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
 * @description: Get roles
 * @param {int} uid
 * @return {*}
 */
func GetRolesByUID(uid int) (roles []string, err error) {
	err = GetDB().Model(&UserRole{}).
		Select("role.key").
		Where("user_role.uid = ? ", uid).
		Joins("inner join role on user_role.role_id = role.id").
		Scan(&roles).Error
	return
}

/**
 * @description: Get permissions by uid
 * @param {int} uid
 * @return {*}
 */
func GetPermissionsByUID(uid int) (permissions []string, err error) {
	var roleIDs []int
	err = GetDB().Model(&UserRole{}).
		Select("role_id").
		Where("uid = ?", uid).
		Scan(&roleIDs).Error
	if len(roleIDs) < 1 {
		return
	}

	permissions, err = GetPermissionsByRoleIDs(roleIDs)
	return
}
