/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-28 15:37:19
 * @LastEditTime: 2022-11-01 17:14:08
 * @Description: Do not edit
 */
package model

import (
	"errors"
	"time"

	iErrors "github.com/chenke1115/hertz-permission/internal/pkg/errors"
	gErrors "github.com/chenke1115/hertz-permission/internal/pkg/errors/gorm"

	"github.com/chenke1115/hertz-permission/internal/constant/status"
	"github.com/chenke1115/hertz-permission/internal/pkg/query"
	"gorm.io/gorm"
)

type Role struct {
	ID         int       `json:"id" gorm:"type:int(11); not null; primaryKey; autoIncrement"`
	Name       string    `json:"name" gorm:"type:varchar(64); not null; unique; comment:角色名称"`
	CreatorID  int       `json:"creator_id" gorm:"type:bigint(20); not null; unsigned; comment:创建者ID"`
	Key        string    `json:"key" gorm:"type:varchar(64); comment:角色标识[跟permission.key区分开]"`
	Status     int       `json:"status" gorm:"type:tinyint(1); default:1; comment:角色状态[1:正常 0:停用]"`
	UpdateBy   string    `json:"update_by" gorm:"type:varchar(64); comment:最后操作人"`
	UpdateTime int       `json:"update_time" gorm:"type:int(12); comment:最后操作时间戳"`
	Remark     string    `json:"remark" gorm:"type:varchar(64); comment:备注"`
	IsDel      string    `json:"is_del" gorm:"type:tinyint(1); default:1; comment:[1:正常 0:删除]"`
	CreatedAt  time.Time `json:"create_at" gorm:"type:timestamp; default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"update_at" gorm:"type:timestamp; default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

type RoleQuery struct {
	Role
	query.PaginationQuery
}

/**
 * @description: Table name
 * @return {*}
 */
func (model Role) TableName() string {
	return "role"
}

/**
 * @description: Do Search
 * @return {*}
 */
func (query RoleQuery) Search() (list *[]Role, total int64, err error) {
	// Init
	list = &[]Role{}

	// Init db-query
	tx := GetDB().Model(&Role{})

	// Set search conditions
	if query.Stime != "" {
		tx = tx.Where("`created_at` >= ?", query.Stime)
	}

	if query.Etime != "" {
		tx = tx.Where("`created_at` < ?", query.Etime)
	}

	// Get data
	total, err = crudAll(&query.PaginationQuery, tx, list)

	return
}

/**
 * @description: Do create
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (model Role) Create(tx *gorm.DB) (err error) {
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
 * @description: Do edit
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (model Role) Edit(tx *gorm.DB) (err error) {
	err = tx.Updates(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = iErrors.Wrap(err, status.RoleNotExistCode)
		} else if gErrors.IsUniqueConstraintError(err) {
			err = iErrors.Wrap(err, status.RoleParamUniqueErrCode)
		} else {
			err = iErrors.WrapCode(err, iErrors.BadRequest)
		}

		return
	}

	return
}

/**
 * @description: Do del
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (model Role) Del(tx *gorm.DB) (err error) {
	// Do del
	if err = tx.Unscoped().Delete(model).Error; err != nil {
		err = iErrors.WrapCode(err, iErrors.BadRequest)
	}

	return
}

/**
 * @description: Get role by id
 * @param {int} id
 * @return {*}
 */
func GetRoleByID(id int) (role Role, err error) {
	err = GetDB().Model(&Role{}).First(&role, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = iErrors.Wrap(err, status.RoleNotExistCode)
		} else {
			err = iErrors.WrapCode(err, iErrors.BadRequest)
		}
	}
	return
}

/**
 * @description: Get role by name
 * @param {string} name
 * @return {*}
 */
func GetRoleByName(name string) (role Role, err error) {
	err = GetDB().Model(&Role{}).First(&role, "name = ?", name).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = iErrors.Wrap(err, status.RoleNotExistCode)
		} else {
			err = iErrors.WrapCode(err, iErrors.BadRequest)
		}
	}
	return
}
