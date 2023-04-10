/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-31 09:45:52
 * @LastEditTime: 2023-04-10 14:36:20
 * @Description: Do not edit
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

type ReqEditData struct {
	Name   string `json:"name,required" form:"name,required"` // 角色名【必填】
	Key    string `json:"key,required" form:"key,required"`   // 角色标识【必填】
	Status int    `json:"status" form:"status"`               // 角色状态[1:正常;0:停用]
	Remark string `json:"remark" form:"remark"`               // 备注
}

// EditHandler goDoc
// @Summary     编辑角色
// @Description This is a api to edit role
// @Tags        Role【角色】
// @Accept      json
// @Produce     json
// @Security    authorization
// @Param       id   path     int         true "角色ID" example(1)
// @Param       data formData ReqEditData true "请求数据"
// @Success     200  {object} response.BaseResponse{data=interface{}}
// @Failure     400  {object} response.BaseResponse{data=interface{}}
// @Router      /api/role/{id}/edit [put]
func EditHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err  error
		ID   int
		req  ReqEditData
		role model.Role
	)

	// Response
	defer func() {
		response.HandleResponse(c, err, nil)
	}()

	// ID
	if ID, err = strconv.Atoi(c.Param("id")); err != nil {
		err = errors.Wrapf(err, status.RoleIdMissCode)
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

	// Binding
	err = c.Bind(&role)
	if err != nil {
		err = errors.Wrapf(err, status.RoleParamBindingErrorCode)
		return
	}

	err = role.Edit(model.GetDB())
}
