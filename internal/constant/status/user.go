/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-09-01 18:22:13
 * @LastEditTime: 2022-11-09 14:22:29
 * @Description: Do not edit
 */
package status

import "github.com/chenke1115/hertz-permission/internal/pkg/errors"

// Status code of user
const (
	UserNotExistCode = iota + 40000
	UserMissIDCode
	UserMissParamCode
	UserErrorParamCode
	UserIncorrectOldPasswordCode
	UserPermissionDeniedCode
	UserUpdateFailCode
	UserParamUniqueErrCode
	UserParamBindErrCode
)

// Message for code
func init() {
	errors.NewCode(UserNotExistCode, "用户不存在")
	errors.NewCode(UserMissIDCode, "用户ID缺失")
	errors.NewCode(UserMissParamCode, "用户参数缺失")
	errors.NewCode(UserErrorParamCode, "用户参数验证错误")
	errors.NewCode(UserIncorrectOldPasswordCode, "旧密码不正确")
	errors.NewCode(UserPermissionDeniedCode, "当前用户权限不足")
	errors.NewCode(UserUpdateFailCode, "用户变更失败")
	errors.NewCode(UserParamUniqueErrCode, "用户参数重复, 需保持唯一")
	errors.NewCode(UserParamBindErrCode, "用户参数绑定失败")
}
