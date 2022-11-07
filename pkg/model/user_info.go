/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-07 16:22:05
 * @LastEditTime: 2022-11-07 18:32:54
 * @Description: Do not edit
 */
package model

import (
	"errors"
	"time"

	"github.com/chenke1115/hertz-permission/internal/constant/status"
	iErrors "github.com/chenke1115/hertz-permission/internal/pkg/errors"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID         int       `json:"id" gorm:"type:int(11); primaryKey; autoIncrement"`
	Name       string    `json:"name" gorm:"type:varchar(32); not null; comment:用户名"`
	Account    string    `json:"account" gorm:"varchar(32); unique; not null; comment:登录账户"`
	CustomerID int       `json:"customer_id" gorm:"type:int(11); index; not null; comment:客户ID"`
	CreatedAt  time.Time `json:"create_at" gorm:"type:timestamp; default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"update_at" gorm:"type:timestamp; default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

/**
 * @description: Table name
 * @return {*}
 */
func (model UserInfo) TableName() string {
	return "user_info"
}

/**
 * @description: Assign role for user
 * @param {[]int} roleIDs
 * @return {*}
 */
func (model UserInfo) AssignRole(db *gorm.DB, roleIDs []int) (err error) {
	return db.Transaction(func(tx *gorm.DB) error {
		// Save
		for _, roleID := range roleIDs {
			userRole := UserRole{
				UID:    model.ID,
				RoleID: roleID,
			}
			if err = userRole.Save(tx); err != nil {
				return err
			}
		}

		// Del
		var ids []int
		userRole := UserRole{UID: model.ID}
		ids, err = userRole.NotIn(tx, roleIDs)
		if err != nil {
			return err
		}

		if len(ids) > 0 {
			err = (&UserRole{}).DelBatch(tx, ids)
			if err != nil {
				return err
			}
		}

		// commit
		return nil
	})
}

/**
 * @description: Get by ID
 * @param {int} id
 * @return {*}
 */
func GetUserInfoByID(id int) (userInfo UserInfo, err error) {
	err = GetDB().Model(&UserInfo{}).First(&userInfo, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = iErrors.Wrap(err, status.UserNotExistCode)
		} else {
			err = iErrors.WrapCode(err, iErrors.BadRequest)
		}
	}
	return
}
