/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-28 11:47:56
 * @LastEditTime: 2022-12-26 16:21:03
 * @Description: Do not edit
 */
package model

import (
	"errors"
	"fmt"
	"strings"

	"github.com/chenke1115/go-common/functions/array"
	"github.com/chenke1115/go-common/functions/conver"
	"github.com/chenke1115/go-common/functions/date"
	"github.com/chenke1115/go-common/functions/match"
	"github.com/chenke1115/hertz-common/global"
	iErrors "github.com/chenke1115/hertz-common/pkg/errors"
	gErrors "github.com/chenke1115/hertz-common/pkg/errors/gorm"
	"github.com/chenke1115/hertz-common/pkg/query"
	"github.com/chenke1115/hertz-permission/internal/constant/status"
	"github.com/chenke1115/hertz-permission/internal/constant/types"

	"gorm.io/gorm"
)

type Permission struct {
	ID         int    `json:"id" gorm:"type:int(11); primaryKey; autoIncrement"`
	PID        int    `json:"pid" gorm:"column:pid; type:int(11); index; comment:父级ID"`
	Name       string `json:"name" gorm:"type:varchar(64); not null; unique; comment:权限名称"`
	Alias      string `json:"alias" gorm:"type:varchar(64); not null; unique; comment:别名"`
	Key        string `json:"key" gorm:"type:varchar(64); comment:权限全局标识[即路由, 类型为目录可空]"`
	Components string `json:"components" gorm:"type:varchar(64); comment:前端页面路径[类型为按钮可空]"`
	Sort       int    `json:"sort" gorm:"type:int(4); default:0; comment:排序[从小到大]"`
	Type       string `json:"type" gorm:"type:char(1); comment:权限类型[D:目录 M:菜单 B:按钮]"`
	Icon       string `json:"icon" gorm:"type:varchar(255); comment:图标"`
	Visible    int    `json:"visible" gorm:"type:tinyint(1); default:1; comment:菜单状态[1:显示 0:隐藏]"`
	Status     int    `json:"status" gorm:"type:tinyint(1); default:1; comment:菜单状态[1:正常 0:停用]"`
	UpdateBy   string `json:"update_by" gorm:"type:varchar(64); comment:最后操作人"`
	UpdateTime int    `json:"update_time" gorm:"type:int(12); comment:最后操作时间戳"`
	Remark     string `json:"remark" gorm:"type:varchar(64); comment:备注"`
	DateModel
}

type PermissionShow struct {
	Permission
	UpdateTime string `json:"update_time"`
}

type PermissionOption struct {
	ID    int                `json:"id"`
	Name  string             `json:"name"`
	Alias string             `json:"alias"`
	Show  string             `json:"show"`
	Child []PermissionOption `json:"child" gorm:"-"`
}

type PermissionQuery struct {
	PermissionShow
	query.PaginationQuery
}

/**
 * @description: Table name
 * @return {*}
 */
func (model Permission) TableName() string {
	return "permission"
}

/**
 * @description: Before operating
 * @return {*}
 */
func (model Permission) Before() error {
	// Type
	if model.Type != "" && !array.In(model.Type, conver.StructToArray(types.PermissionTypeArr)) {
		return iErrors.New(status.PermissionTypeErrorCode)
	}

	// Name
	if model.Name != "" && !match.IsKeyString(model.Name) {
		return iErrors.New(status.PermissionNameErrorCode)
	}

	// Alias
	if model.Alias != "" && !match.IsNicknameString(model.Alias) {
		return iErrors.New(status.PermissionAliasErrorCode)
	}

	// Key
	if (model.Type != types.PermissionTypeArr.Dir && model.Key == "") ||
		(model.Key != "" && !IsValidRoute(model.Key)) {
		return iErrors.New(status.PermissionKeyErrorCode)
	}

	// Components
	if model.Type != types.PermissionTypeArr.Button && model.Components == "" {
		return iErrors.Newf(status.PermissionParamErrorCode, "该类型下，前端路径不能为空")
	}

	// Visible and Status
	if !array.In(model.Visible, []int{status.StateEnabled, status.StateInit}) ||
		!array.In(model.Status, []int{status.StateEnabled, status.StateInit}) {
		return iErrors.New(status.PermissionStatusErrorCode)
	}

	return nil
}

/**
 * @description: Do Search
 * @return {*}
 */
func (query PermissionQuery) Search() (list *[]PermissionShow, total int64, err error) {
	// Init
	list = &[]PermissionShow{}
	permissions := &[]Permission{}

	// Init db-query
	tx := GetDB().Model(&Permission{})

	// Set search conditions
	if query.Stime != "" {
		tx = tx.Where("`created_at` >= ?", query.Stime)
	}

	if query.Etime != "" {
		tx = tx.Where("`created_at` < ?", query.Etime)
	}

	// Get data
	total, err = crudAll(&query.PaginationQuery, tx, permissions)
	for _, permission := range *permissions {
		perShow := PermissionShow{}
		perShow.Permission = permission
		perShow.UpdateTime = date.DateFormat(permission.UpdateTime)
		*list = append(*list, perShow)
	}

	return
}

/**
 * @description: Get for permission option
 * @param {[]Permission} list
 * @param {error} err
 * @return {*}
 */
func (model Permission) Option() (list []PermissionOption, err error) {
	return model.GetOption(0)
}

/**
 * @description: Get permission for menu option
 * @return {*}
 */
func (model Permission) MenuOption() (list map[int]string, err error) {
	permissions := []Permission{}
	err = GetDB().Model(&Permission{}).
		Select("id, name, alias").
		Where("status = 1").
		Find(&permissions).Error
	if err != nil {
		err = iErrors.WrapCode(err, iErrors.BadRequest)
		return
	}

	list = map[int]string{}
	for _, v := range permissions {
		list[v.ID] = v.Alias + "[" + v.Name + "]"
	}

	return
}

/**
 * @description: Do create
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (model Permission) Create(tx *gorm.DB) (err error) {
	if err = model.Before(); err != nil {
		return
	}

	err = tx.Create(&model).Error
	if err != nil {
		if gErrors.IsUniqueConstraintError(err) {
			err = iErrors.Wrap(err, status.PermissionParamUniqueErrCode)
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
func (model Permission) Edit(tx *gorm.DB) (err error) {
	if err = model.Before(); err != nil {
		return
	}

	err = tx.Save(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = iErrors.Wrap(err, status.PermissionNotExistCode)
		} else if gErrors.IsUniqueConstraintError(err) {
			err = iErrors.Wrap(err, status.PermissionParamUniqueErrCode)
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
func (model Permission) Del(tx *gorm.DB) (err error) {
	// Is in use
	if IsExistPermissionID(model.ID) {
		err = iErrors.New(status.PermissionIsUseErrorCode)
		return
	}

	// Do del
	if err = tx.Unscoped().Delete(model).Error; err != nil {
		err = iErrors.WrapCode(err, iErrors.BadRequest)
	}

	return
}

/**
 * @description: Get for permission option
 * @param {[]Permission} list
 * @param {error} err
 * @return {*}
 */
func (model Permission) GetOption(pid int) (list []PermissionOption, err error) {
	err = GetDB().Model(&Permission{}).
		Select("id, pid, name, alias").
		Where("pid = ? and status = 1", pid).
		Scan(&list).Error
	if err != nil {
		err = iErrors.WrapCode(err, iErrors.BadRequest)
		return
	}

	for k, v := range list {
		var c []PermissionOption
		c, err = model.GetOption(v.ID)
		if err != nil {
			err = iErrors.WrapCode(err, iErrors.BadRequest)
			return
		}

		list[k].Show = fmt.Sprintf("%s[%s]", v.Alias, v.Name)
		if len(c) > 0 {
			list[k].Child = c
		} else {
			list[k].Child = []PermissionOption{}
		}
	}

	return
}

/**
 * @description: Get by ID
 * @param {int} id
 * @return {*}
 */
func GetPermissionByID(id int) (permission Permission, err error) {
	err = GetDB().Model(&Permission{}).First(&permission, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = iErrors.Wrap(err, status.PermissionNotExistCode)
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
func IsEnablePermission(tx *gorm.DB, id int) bool {
	err := tx.Model(&Permission{}).First(&Permission{}, "id = ? and status = 1", id).Error
	return err == nil
}

/**
 * @description: Verify that the route is valid
 * @param {string} key
 * @return {*}
 */
func IsValidRoute(key string) bool {
	for _, route := range global.RouteInfo {
		if strings.HasPrefix(key, route.Path) && strings.HasSuffix(key, route.Path) {
			return true
		}
	}

	return false
}

/**
 * @description: Get permission keys by ids
 * @param {[]int} ids
 * @return {*}
 */
func GetPermissionKeysByIDs(ids []int) (perKeys []string, err error) {
	err = GetDB().Model(&Permission{}).
		Select("key").
		Where("id in (?) and `key` <> ''", ids).
		Where("status = 1").
		Scan(&perKeys).Error
	return
}

/**
 * @description: Get permissions by ids
 * @param {[]int} ids
 * @return {*}
 */
func GetPermissionsByIDs(ids []int) (permissions []Permission, err error) {
	err = GetDB().Model(&Permission{}).
		Where("id in (?) and `key` <> ''", ids).
		Where("status = 1").
		Scan(&permissions).Error
	return
}

/**
 * @description: Get all
 * @return {*}
 */
func GetAllPermissions() (permissions []Permission, err error) {
	err = GetDB().Model(&Permission{}).
		Where("status = 1").
		Scan(&permissions).Error
	return
}

/**
 * @description: Get keys
 * @return {*}
 */
func GetPermissionKeys() (perKeys []string, err error) {
	err = GetDB().Model(&Permission{}).
		Select("key").
		Group("key").
		Scan(&perKeys).Error

	return
}
