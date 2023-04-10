/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-08 10:11:54
 * @LastEditTime: 2023-04-10 14:34:19
 * @Description: Do not edits
 */
package role

import (
	"context"
	"strconv"

	"github.com/chenke1115/hertz-common/pkg/errors"
	"github.com/chenke1115/hertz-common/pkg/response"
	"github.com/chenke1115/hertz-permission/internal/constant/status"
	"github.com/chenke1115/hertz-permission/pkg/model"
	"github.com/cloudwego/hertz/pkg/app"
)

type ReqBindData struct {
	PermissionIDs []int `json:"permission_ids,required" gorm:"permission_ids,required"` // 权限ID数组【必填】example([1,2,3])
}

// BindHandler goDoc
// @Summary     绑定权限
// @Description This is a api to binding permission for role
// @Tags        Role【角色】
// @Accept      x-www-form-urlencoded
// @Produce     json
// @Security    authorization
// @Param       id   path     int         true "角色ID" example(1)
// @Param       data formData ReqBindData true "请求数据"
// @Success     200  {object} response.BaseResponse{data=interface{}}
// @Failure     400  {object} response.BaseResponse{data=interface{}}
// @Router      /api/role/{id}/bind [post]
func BindHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err  error
		req  ReqBindData
		ID   int
		role model.Role
	)

	// Response
	defer func() {
		response.HandleResponse(c, err, nil)
	}()

	// ID
	if ID, err = strconv.Atoi(c.Param("id")); err != nil {
		err = errors.Wrap(err, status.RoleIdMissCode)
		return
	}

	// BindAndValidate
	err = c.BindAndValidate(&req)
	if err != nil {
		err = errors.WrapCode(err, status.RoleParamErrorCode)
		return
	}

	// Find
	if role, err = model.GetRoleByID(ID); err != nil {
		return
	}

	// Do binding
	err = role.Binding(model.GetDB(), req.PermissionIDs)
}
