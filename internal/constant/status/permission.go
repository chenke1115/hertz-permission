/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-09-28 09:31:13
 * @LastEditTime: 2022-10-28 14:08:27
 * @Description: Do not edit
 */
package status

import "github.com/chenke1115/ismart-permission/internal/pkg/errors"

// Permission
const (
	PermissionNotExistCode = iota + 40100
	PermissionIdMissCode
	PermissionParamErrorCode
	PermissionParamUniqueErrCode
	PermissionParamBindingErrorCode
	PermissionNameErrorCode
)

// Message for codes
func init() {
	errors.NewCode(PermissionNotExistCode, "权限不存在")
	errors.NewCode(PermissionIdMissCode, "权限ID缺失")
	errors.NewCode(PermissionParamErrorCode, "权限参数错误")
	errors.NewCode(PermissionParamUniqueErrCode, "权限参数重复, 需保持唯一")
	errors.NewCode(PermissionParamBindingErrorCode, "权限参数绑定错误")
	errors.NewCode(PermissionNameErrorCode, "权限不合规范或不在路由列表内")
}
