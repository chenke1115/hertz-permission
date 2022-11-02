/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-31 09:45:52
 * @LastEditTime: 2022-11-01 14:44:28
 * @Description: Do not edit
 */
package role

import (
	"context"
	"strconv"

	"github.com/chenke1115/hertz-permission/internal/constant/status"
	"github.com/chenke1115/hertz-permission/internal/pkg/errors"
	"github.com/chenke1115/hertz-permission/internal/pkg/response"
	"github.com/chenke1115/hertz-permission/pkg/model"
	"github.com/cloudwego/hertz/pkg/app"
)

type ReqEditData struct {
	Name   string `json:"name,required" form:"name,required"` //lint:ignore SA5008 ignoreCheck
	Key    string `json:"key,required" form:"key,required"`   //lint:ignore SA5008 ignoreCheck
	Status int    `json:"status" form:"status"`
	Remark string `json:"remark" form:"remark"`
}

// EditHandler goDoc
// @Summary     编辑角色
// @Description This is a api to edit role
// @Tags        RoleEdit
// @Accept      json
// @Produce     json
// @Param       id     query    int    true  "角色ID"            example(1)
// @Param       name   body     string true  "角色名"             example(“系统管理员”)
// @Param       Key    body     string true  "角色标识"            example(“SYS_ADMIN”)
// @Param       status body     int    false "角色状态[1:正常;0:停用]" Enums(1, 0)
// @Param       remark body     string false "备注"              maxlength(255)
// @Success     200    {object} response.BaseResponse{data=interface{}}
// @Failure     400    {object} response.BaseResponse{data=interface{}}
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
