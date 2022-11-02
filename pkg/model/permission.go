/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-28 11:47:56
 * @LastEditTime: 2022-11-02 14:31:49
 * @Description: Do not edit
 */
package model

import (
	"errors"
	"strings"
	"time"

	iErrors "github.com/chenke1115/ismart-permission/internal/pkg/errors"
	gErrors "github.com/chenke1115/ismart-permission/internal/pkg/errors/gorm"

	"github.com/chenke1115/ismart-permission/internal/constant/global"
	"github.com/chenke1115/ismart-permission/internal/constant/status"
	"gorm.io/gorm"
)

type Permission struct {
	ID         int       `json:"id" gorm:"type:int(11); primaryKey; autoIncrement"`
	PID        int       `json:"pid" gorm:"column:pid; type:int(11); index; comment:父级ID"`
	Name       string    `json:"name" gorm:"type:varchar(64); not null; unique; comment:权限名称"`
	Alias      string    `json:"alias" gorm:"type:varchar(64); not null; unique; comment:别名"`
	Key        string    `json:"key" gorm:"type:varchar(64); unique; comment:权限全局标识[即路由, 类型为目录可空]"`
	Components string    `json:"components" gorm:"type:varchar(64); comment:前端页面路径[类型为按钮可空]"`
	Sort       int       `json:"sort" gorm:"type:int(4); default:0; comment:排序[从小到大]"`
	Type       string    `json:"type" gorm:"type:char(1); comment:权限类型[D:目录 M:菜单 B:按钮]"`
	Icon       string    `json:"icon" gorm:"type:varchar(255); comment:图标"`
	Visible    int       `json:"visible" gorm:"type:tinyint(1); default:1; comment:菜单状态[1:显示 0:隐藏]"`
	Status     int       `json:"status" gorm:"type:tinyint(1); default:1; comment:菜单状态[1:正常 0:停用]"`
	UpdateBy   string    `json:"update_by" gorm:"type:varchar(64); comment:最后操作人"`
	UpdateTime int       `json:"update_time" gorm:"type:int(12); comment:最后操作时间戳"`
	Remark     string    `json:"remark" gorm:"type:varchar(64); comment:备注"`
	CreatedAt  time.Time `json:"create_at" gorm:"type:timestamp; default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"update_at" gorm:"type:timestamp; default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`

	// ignore
	Child []Permission `json:"child" gorm:"-"`
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
	// Check key
	if !IsValidRoute(model.Key) {
		return iErrors.New(status.PermissionKeyErrorCode)
	}

	return nil
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
	err = tx.Updates(&model).Error
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
	// TODO

	// Do del
	if err = tx.Unscoped().Delete(model).Error; err != nil {
		err = iErrors.WrapCode(err, iErrors.BadRequest)
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
