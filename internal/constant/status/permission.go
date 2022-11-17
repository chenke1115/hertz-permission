/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-09-28 09:31:13
 * @LastEditTime: 2022-11-08 11:25:20
 * @Description: Do not edit
 */
package status

import "github.com/chenke1115/hertz-common/pkg/errors"

// Permission
const (
	PermissionNotExistCode = iota + 40100
	PermissionIdMissCode
	PermissionParamErrorCode
	PermissionParamUniqueErrCode
	PermissionParamBindingErrorCode
	PermissionNameErrorCode
	PermissionAliasErrorCode
	PermissionKeyErrorCode
	PermissionTypeErrorCode
	PermissionStatusErrorCode
	PermissionIsUseErrorCode
	PermissionIsUnableErrorCode
)

// Message for codes
func init() {
	errors.NewCode(PermissionNotExistCode, "权限不存在")
	errors.NewCode(PermissionIdMissCode, "权限ID缺失")
	errors.NewCode(PermissionParamErrorCode, "权限参数错误")
	errors.NewCode(PermissionParamUniqueErrCode, "权限参数重复, 需保持唯一")
	errors.NewCode(PermissionParamBindingErrorCode, "权限参数绑定错误")
	errors.NewCode(PermissionNameErrorCode, "权限名称不合规范")
	errors.NewCode(PermissionAliasErrorCode, "权限别名不合规范")
	errors.NewCode(PermissionKeyErrorCode, "权限标识不合规范或不在路由列表内")
	errors.NewCode(PermissionTypeErrorCode, "权限类型不在枚举值内")
	errors.NewCode(PermissionStatusErrorCode, "权限状态不在枚举值内")
	errors.NewCode(PermissionIsUseErrorCode, "权限被使用中，移除后在执行该操作")
	errors.NewCode(PermissionIsUnableErrorCode, "权限不存在或状态异常")
}
