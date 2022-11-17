/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-28 14:07:47
 * @LastEditTime: 2022-11-08 11:27:24
 * @Description: Do not edit
 */
package status

import "github.com/chenke1115/hertz-common/pkg/errors"

// Role
const (
	RoleNotExistCode = iota + 40200
	RoleIdMissCode
	RoleParamErrorCode
	RoleNameErrorCode
	RoleKeyErrorCode
	RoleStatusParamErrCode
	RoleParamUniqueErrCode
	RoleParamBindingErrorCode
	RoleDelFixedValErrorCode
	RoleDelInUseRecordCode
	RoleStatusErrorCode
)

// Message for codes
func init() {
	errors.NewCode(RoleNotExistCode, "角色不存在")
	errors.NewCode(RoleIdMissCode, "角色ID缺失")
	errors.NewCode(RoleParamErrorCode, "角色参数错误")
	errors.NewCode(RoleNameErrorCode, "角色名称不合规范")
	errors.NewCode(RoleKeyErrorCode, "角色标识不合规范")
	errors.NewCode(RoleStatusParamErrCode, "角色状态不在枚举值内")
	errors.NewCode(RoleParamUniqueErrCode, "角色参数重复, 需保持唯一")
	errors.NewCode(RoleParamBindingErrorCode, "角色参数绑定错误")
	errors.NewCode(RoleDelFixedValErrorCode, "固定角色不能删除")
	errors.NewCode(RoleDelInUseRecordCode, "有用户正在使用该角色，请先变更")
	errors.NewCode(RoleStatusErrorCode, "角色状态异常")
}
