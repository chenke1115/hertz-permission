/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-31 09:44:07
 * @LastEditTime: 2022-11-01 14:52:05
 * @Description: Do not edit
 */
package role

import (
	"context"

	"github.com/chenke1115/hertz-permission/internal/constant/status"
	"github.com/chenke1115/hertz-permission/internal/pkg/errors"
	"github.com/chenke1115/hertz-permission/internal/pkg/response"
	"github.com/chenke1115/hertz-permission/pkg/model"

	"github.com/cloudwego/hertz/pkg/app"
)

type ReqAddData struct {
	Name   string `json:"name,required" form:"name,required"` //lint:ignore SA5008 ignoreCheck
	Key    string `json:"key,required" form:"key,required"`   //lint:ignore SA5008 ignoreCheck
	Status int    `json:"status" form:"status"`
	Remark string `json:"remark" form:"remark"`
}

// AddHandler goDoc
// @Summary     添加角色
// @Description This is a api to add role
// @Tags        RoleAdd
// @Accept      json
// @Produce     json
// @Param       name   body     string true  "角色名"             example(“系统管理员”)
// @Param       Key    body     string true  "角色标识"            example(“SYS_ADMIN”)
// @Param       status body     int    false "角色状态[1:正常;0:停用]" Enums(1, 0)
// @Param       remark body     string false "备注"              maxlength(255)
// @Success     200    {object} response.BaseResponse{data=interface{}}
// @Failure     400    {object} response.BaseResponse{data=interface{}}
// @Router      /api/role/add [post]
func AddHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err error
		req ReqAddData
	)

	defer func() {
		response.HandleResponse(c, err, nil)
	}()

	// BindAndValidate
	err = c.BindAndValidate(&req)
	if err != nil {
		err = errors.WrapCode(err, status.RoleParamErrorCode)
		return
	}

	// Binding to model
	role := &model.Role{}
	err = c.Bind(&role)
	if err != nil {
		err = errors.Wrapf(err, status.RoleParamBindingErrorCode)
		return
	}

	err = role.Create(model.GetDB())
}
