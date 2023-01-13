/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-28 15:37:19
 * @LastEditTime: 2022-12-01 09:17:17
 * @Description: Do not edit
 */
package model

import (
	"errors"

	"github.com/chenke1115/go-common/functions/array"
	"github.com/chenke1115/go-common/functions/date"
	"github.com/chenke1115/go-common/functions/match"
	iErrors "github.com/chenke1115/hertz-common/pkg/errors"
	gErrors "github.com/chenke1115/hertz-common/pkg/errors/gorm"
	"github.com/chenke1115/hertz-common/pkg/query"
	"github.com/chenke1115/hertz-permission/internal/constant/status"

	"gorm.io/gorm"
)

type Role struct {
	ID         int    `json:"id" gorm:"type:int(11); not null; primaryKey; autoIncrement"`
	Name       string `json:"name" gorm:"type:varchar(64); not null; unique; comment:角色名称"`
	CreatorID  int    `json:"creator_id" gorm:"type:bigint(20); not null; unsigned; comment:创建者ID"`
	Key        string `json:"key" gorm:"type:varchar(64); unique; comment:角色标识[跟permission.key区分开]"`
	Status     int    `json:"status" gorm:"type:tinyint(1); default:1; comment:角色状态[1:正常 0:停用]"`
	UpdateBy   string `json:"update_by" gorm:"type:varchar(64); comment:最后操作人"`
	UpdateTime int    `json:"update_time" gorm:"type:int(12); comment:最后操作时间戳"`
	Remark     string `json:"remark" gorm:"type:varchar(64); comment:备注"`
	IsDel      int    `json:"is_del" gorm:"type:tinyint(1); default:0; comment:[0:正常 1:删除]"`
	DateModel
}

type RoleShow struct {
	Role
	UpdateTime string `json:"update_time"`
}

type RoleOption struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
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

func (model Role) Before() error {
	// Name
	if model.Name != "" && !match.IsNicknameString(model.Name) {
		return iErrors.New(status.RoleNameErrorCode)
	}

	// Key
	if model.Key != "" && !match.IsUserNameString(model.Key) {
		return iErrors.New(status.RoleKeyErrorCode)
	}

	// Status
	if !array.In(model.Status, []int{status.StateEnabled, status.StateInit}) {
		return iErrors.New(status.RoleStatusParamErrCode)
	}

	return nil
}

/**
 * @description: Do Search
 * @return {*}
 */
func (query RoleQuery) Search() (list *[]RoleShow, total int64, err error) {
	// Init
	list = &[]RoleShow{}
	roles := &[]Role{}

	// Init db-query
	tx := GetDB().Model(&Role{}).
		Where("`is_del` <> 1")

	// Set search conditions
	if query.Stime != "" {
		tx = tx.Where("`created_at` >= ?", query.Stime)
	}

	if query.Etime != "" {
		tx = tx.Where("`created_at` < ?", query.Etime)
	}

	// Get data
	total, err = crudAll(&query.PaginationQuery, tx, roles)
	for _, role := range *roles {
		roleShow := RoleShow{}
		roleShow.Role = role
		roleShow.UpdateTime = date.DateFormat(role.UpdateTime)
		*list = append(*list, roleShow)
	}

	return
}

/**
 * @description: Get Option
 * @return {*}
 */
func (model Role) Option() (option []RoleOption, err error) {
	err = GetDB().Model(&Role{}).Select("`id`, `name`, `key`").
		Where("`is_del` = 0 and `status` = 1").
		Scan(&option).Error
	if err != nil {
		err = iErrors.WrapCode(err, iErrors.BadRequest)
		return
	}

	return
}

/**
 * @description: Do create
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (model Role) Create(tx *gorm.DB) (err error) {
	if err = model.Before(); err != nil {
		return
	}

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
	if err = model.Before(); err != nil {
		return
	}

	err = tx.Save(&model).Error
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
	// set is_del
	model.IsDel = status.StateEnabled
	if err = tx.Updates(&model).Error; err != nil {
		err = iErrors.WrapCode(err, iErrors.BadRequest)
	}

	return
}

/**
 * @description: Get bind list
 * @return {*}
 */
func (model Role) BindList() (list []int, err error) {
	var perList []RolePermission
	rolePer := RolePermission{RoleID: model.ID}
	perList, err = rolePer.List()
	if err != nil {
		return
	}

	for _, item := range perList {
		list = append(list, item.PermissionID)
	}

	return
}

/**
 * @description: Binding permission for role
 * @param {[]int} perIDs
 * @return {*}
 */
func (model Role) Binding(db *gorm.DB, perIDs []int) (err error) {
	return db.Transaction(func(tx *gorm.DB) error {
		// Save
		for _, perID := range perIDs {
			rolePer := RolePermission{
				RoleID:       model.ID,
				PermissionID: perID,
			}
			if err = rolePer.Save(tx); err != nil {
				return err
			}
		}

		// Del
		var ids []int
		rolePer := RolePermission{RoleID: model.ID}
		ids, err = rolePer.NotIn(tx, perIDs)
		if err != nil {
			return err
		}

		if len(ids) > 0 {
			err = (&RolePermission{}).DelBatch(tx, ids)
			if err != nil {
				return err
			}
		}

		// commit
		return nil
	})
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

/**
 * @description: Is enable
 * @param {int} id
 * @return {*}
 */
func IsEnableRole(tx *gorm.DB, id int) bool {
	var role Role
	err := tx.Model(&Role{}).First(&role, "id = ? and status = 1 and is_del = 0", id).Error

	return err == nil
}
