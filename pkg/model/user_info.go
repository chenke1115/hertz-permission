/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-07 16:22:05
 * @LastEditTime: 2022-11-14 11:30:01
 * @Description: Do not edit
 */
package model

import (
	"errors"
	"strings"
	"time"

	"github.com/chenke1115/hertz-permission/internal/constant/consts"
	"github.com/chenke1115/hertz-permission/internal/constant/status"
	"github.com/chenke1115/hertz-permission/internal/constant/types"
	"github.com/chenke1115/hertz-permission/internal/pkg/array"
	iErrors "github.com/chenke1115/hertz-permission/internal/pkg/errors"
	gErrors "github.com/chenke1115/hertz-permission/internal/pkg/errors/gorm"
	"github.com/chenke1115/hertz-permission/internal/pkg/hash"
	"github.com/chenke1115/hertz-permission/internal/pkg/match"
	"github.com/chenke1115/hertz-permission/internal/pkg/query"

	"gorm.io/gorm"
)

type UserInfo struct {
	ID         int       `json:"id" gorm:"type:int(11); primaryKey; autoIncrement"`
	Name       string    `json:"name" gorm:"type:varchar(32); not null; comment:用户名"`
	Account    string    `json:"account" gorm:"type:varchar(32); unique; not null; comment:登录账户"`
	CustomerID int       `json:"customer_id" gorm:"type:int(11); index; not null; comment:客户ID"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:timestamp; default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:timestamp; default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

type APIUser struct {
	UserInfo
	UserID int `json:"user_id"`
	Status int `json:"status"`
}

type UserQuery struct {
	APIUser
	query.PaginationQuery
	Status string `json:"status" form:"status"`
}

type CurrentUser struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Account     string   `json:"account"`
	CustomerID  int      `json:"customer_id"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

/**
 * @description: Table name
 * @return {*}
 */
func (model UserInfo) TableName() string {
	return "user_info"
}

/**
 * @description: Before operation
 * @return {*}
 */
func (model UserInfo) Before() error {
	// Name
	if model.Name != "" && !match.IsNicknameString(model.Name) {
		return iErrors.New(status.UserErrorParamCode)
	}

	return nil
}

/**
 * @description: Do Search
 * @return {*}
 */
func (q UserQuery) Search() (list *[]APIUser, total int64, err error) {
	// Init
	list = &[]APIUser{}

	// Init db-query
	tx := GetDB().Model(&UserInfo{}).
		Select("`user_info`.*, `user`.`id` as user_id, `user`.`status` as status").
		Joins("inner join `user` on `user`.`account_id` = `user_info`.`id`")

	// Set search conditions
	if q.Stime != "" {
		tx = tx.Where("`user_info`.`created_at` >= ?", q.Stime)
	}

	if q.Etime != "" {
		tx = tx.Where("`user_info`.`created_at` < ?", q.Etime)
	}

	if q.Status != "" {
		tx = tx.Where("`user`.`status` = ?", q.Status)
	} else {
		tx = tx.Where("`user`.`status` = ?", status.StateEnabled)
	}

	// Get data
	total, err = crudAll(&q.PaginationQuery, tx, list)
	return
}

/**
 * @description: Create
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (model UserInfo) Create(tx *gorm.DB) (err error) {
	if err = model.Before(); err != nil {
		return
	}

	err = tx.Transaction(func(db *gorm.DB) error {
		if err = tx.Create(&model).Error; err != nil {
			if gErrors.IsUniqueConstraintError(err) {
				err = iErrors.Wrap(err, status.UserParamUniqueErrCode)
			} else {
				err = iErrors.WrapCode(err, iErrors.BadRequest)
			}
			return err
		}

		// create user
		var user = User{
			Password:  hash.GetHashedPassword(consts.InitPassword, consts.Salt),
			AccountID: model.ID,
		}
		if err = user.Create(tx); err != nil {
			return err
		}

		// commit
		return nil
	})
	return
}

/**
 * @description: Edit
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (model UserInfo) Edit(tx *gorm.DB) (err error) {
	if err = model.Before(); err != nil {
		return
	}

	err = tx.Updates(&model).Error
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
 * @param {*gorm.DB} tx
 * @return {*}
 */
func (model UserInfo) Del(tx *gorm.DB) (err error) {
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

/**
 * @description: Check
 * @param {string} username
 * @param {string} hashedPassword
 * @return {*}
 */
func CheckUsernameAndPassword(username string, hashedPassword string) (cuser CurrentUser, err error) {
	var userInfo UserInfo
	err = GetDB().Model(UserInfo{}).
		Select("user_info.*").
		Where("user_info.account = ? and user.password = ?", username, hashedPassword).
		Joins("inner join user on user.account_id=user_info.id").
		First(&userInfo).Error
	if err != nil {
		return
	}

	cuser.ID = userInfo.ID
	cuser.Name = userInfo.Name
	cuser.Account = userInfo.Account
	cuser.CustomerID = userInfo.CustomerID
	cuser.Roles, err = GetRolesByUID(cuser.ID)
	return
}

/**
 * @description: Judge the ID is owner id or not
 * @param {int} id
 * @return {*}
 */
func (cuser CurrentUser) IsOwner(id int) bool {
	return cuser.CustomerID == id
}

/**
 * @description: Judge the user is super user or not
 * @return {*}
 */
func (cuser CurrentUser) IsSuperUser() bool {
	for _, role := range cuser.Roles {
		if array.InArray(role, types.SuperRoleArr) {
			return true
		}
	}

	return false
}

/**
 * @description: Check role permission
 * @param {string} permissionFunc
 * @return {*}
 */
func (cuser CurrentUser) CheckPermission(permissionFunc string) bool {
	permissions := cuser.Permissions
	for _, permission := range permissions {
		if strings.HasPrefix(permissionFunc, permission) {
			return true
		}
	}

	return false
}
