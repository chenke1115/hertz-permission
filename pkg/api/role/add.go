/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-31 09:44:07
 * @LastEditTime: 2023-04-10 14:33:58
 * @Description: Do not edit
 */
package role

import (
	"context"

	"github.com/chenke1115/go-common/functions/date"
	"github.com/chenke1115/hertz-common/pkg/errors"
	"github.com/chenke1115/hertz-common/pkg/response"
	"github.com/chenke1115/hertz-permission/internal/constant/status"
	"github.com/chenke1115/hertz-permission/pkg/middleware"
	"github.com/chenke1115/hertz-permission/pkg/model"
	"github.com/cloudwego/hertz/pkg/app"
)

type ReqAddData struct {
	Name   string `json:"name,required" form:"name,required" vd:"len($)<32"` // 角色名【必填】 example(“系统管理员”)
	Key    string `json:"key,required" form:"key,required" vd:"len($)<32"`   // 角色标识【必填】 example(“SYS_ADMIN”)
	Status int    `json:"status" form:"status"`                              // 角色状态[1:正常;0:停用]
	Remark string `json:"remark" form:"remark" vd:"len($)<256"`              // 备注
}

// AddHandler goDoc
// @Summary     添加角色
// @Description This is a api to add role
// @Tags        Role【角色】
// @Accept      x-www-form-urlencoded
// @Produce     json
// @Security    authorization
// @Param       data formData ReqAddData true "请求数据"
// @Success     200  {object} response.BaseResponse{data=interface{}}
// @Failure     400  {object} response.BaseResponse{data=interface{}}
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

	cuser, _ := middleware.GetCurrentUser(ctx, c)
	role.CreatorID = cuser.ID
	role.UpdateBy = cuser.Account
	role.UpdateTime = date.DateUnix()

	err = role.Create(model.GetDB())
}
