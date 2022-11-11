/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-08 16:52:22
 * @LastEditTime: 2022-11-10 14:57:03
 * @Description: Do not edit
 */
package model

import (
	"errors"
	"time"

	"github.com/chenke1115/hertz-permission/internal/constant/status"
	iErrors "github.com/chenke1115/hertz-permission/internal/pkg/errors"
	gErrors "github.com/chenke1115/hertz-permission/internal/pkg/errors/gorm"

	"gorm.io/gorm"
)

type User struct {
	ID            int       `json:"id" gorm:"type:int(11); primaryKey; autoIncrement"`
	Password      string    `json:"password" gorm:"type:varchar(128); not null; comment:密码"`
	AccountID     int       `json:"account_id" gorm:"unique; not null; comment:账户ID"`
	ActivateToken string    `json:"activate_token" gorm:"type:varchar(255); comment:生效TOKEN"`
	ResetToken    string    `json:"reset_token" gorm:"type:varchar(255); comment:重置TOKEN"`
	Status        int       `json:"enabled" gorm:"type:char(1); index; default:1; comment:生效状态[1:启用 0:失效]"`
	CreatedAt     time.Time `json:"created_at" gorm:"type:timestamp; default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"type:timestamp; default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

/**
 * @description: Table name
 * @return {*}
 */
func (model User) TableName() string {
	return "user"
}

/**
 * @description: Create user
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (model User) Create(tx *gorm.DB) (err error) {
	err = tx.Create(&model).Error
	if err != nil {
		if gErrors.IsUniqueConstraintError(err) {
			err = iErrors.Wrap(err, status.UserParamUniqueErrCode)
		} else {
			err = iErrors.WrapCode(err, iErrors.BadRequest)
		}
	}
	return
}

/**
 * @description: Edit
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (model User) Edit(tx *gorm.DB) (err error) {
	err = tx.Save(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = iErrors.Wrap(err, status.UserNotExistCode)
		} else if gErrors.IsUniqueConstraintError(err) {
			err = iErrors.Wrap(err, status.UserParamUniqueErrCode)
		} else {
			err = iErrors.WrapCode(err, iErrors.BadRequest)
		}
	}
	return
}

/**
 * @description: Del
 * @param {int} id
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (model User) Del(tx *gorm.DB) (err error) {
	err = tx.Unscoped().Delete(model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = iErrors.Wrap(err, status.UserNotExistCode)
		} else {
			err = iErrors.WrapCode(err, iErrors.BadRequest)
		}
		return
	}
	return
}

/**
 * @description: Get by ID
 * @param {int} id
 * @return {*}
 */
func GetUserByID(id int) (user User, err error) {
	err = GetDB().Model(&User{}).First(&user, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = iErrors.Wrap(err, status.UserNotExistCode)
		} else {
			err = iErrors.WrapCode(err, iErrors.BadRequest)
		}
	}
	return
}
