/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-31 09:45:52
 * @LastEditTime: 2023-04-10 14:28:32
 * @Description: Do not edit
 */
package permission

import (
	"context"
	"strconv"

	"github.com/chenke1115/go-common/functions/date"
	"github.com/chenke1115/hertz-common/pkg/errors"
	_ "github.com/chenke1115/hertz-common/pkg/errors/validate"
	"github.com/chenke1115/hertz-common/pkg/response"
	"github.com/chenke1115/hertz-permission/internal/constant/status"
	"github.com/chenke1115/hertz-permission/pkg/middleware"
	"github.com/chenke1115/hertz-permission/pkg/model"

	"github.com/cloudwego/hertz/pkg/app"
)

type ReqEditData struct {
	Name       string `json:"name,required" form:"name,required" vd:"len($)<64"`             // 权限名称[允许英文，数字，.]【必填】example("permission.add")
	Alias      string `json:"alias,required" form:"alias,required" vd:"len($)<64"`           // 别名[允许中文，英文，数字，_]【必填】example("添加权限")
	Type       string `json:"type,required" form:"type,required"`                            // 权限类型[D:目录;M:菜单;B:按钮]【必填】Enums("D", "M", "B")
	Key        string `json:"key,required" form:"key,required"`                              // 权限全局标识[即后端路由，类型为目录可空]【必填】
	Components string `json:"components,required" form:"components,required" vd:"len($)<64"` // 前端页面路径[类型为按钮可空]【必填】
	PID        int    `json:"pid" form:"pid" default:"0"`                                    // 父级ID
	Sort       int    `json:"sort" form:"sort" default:"0"`                                  // 排序[从小到大]
	Icon       string `json:"icon" form:"icon"`                                              // 图标
	Visible    int    `json:"visible" form:"visible" default:"1"`                            // 菜单状态[1:显示;0:隐藏]
	Status     int    `json:"status" form:"status" default:"1"`                              // 菜单状态[1:正常 0:停用]
	Remark     string `json:"remark" form:"remark" vd:"len($)<256"`                          // 备注
}

// EditHandler goDoc
// @Summary     修改权限
// @Description This is a api to edit permission
// @Tags        Permission【权限】
// @Accept      x-www-form-urlencoded
// @Produce     json
// @Security    authorization
// @Param       id   path     int         true "权限ID" example(1)
// @Param       data formData ReqEditData true "请求数据"
// @Success     200  {object} response.BaseResponse{data=interface{}}
// @Failure     400  {object} response.BaseResponse{data=interface{}}
// @Router      /api/permission/{id}/edit [put]
func EditHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err        error
		ID         int
		req        ReqEditData
		permission model.Permission
	)

	defer func() {
		response.HandleResponse(c, err, nil)
	}()

	// ID
	if ID, err = strconv.Atoi(c.Param("id")); err != nil {
		err = errors.Wrapf(err, status.PermissionIdMissCode)
		return
	}

	// BindAndValidate
	err = c.BindAndValidate(&req)
	if err != nil {
		err = errors.WrapCode(err, status.PermissionParamErrorCode)
		return
	}

	// Find
	if permission, err = model.GetPermissionByID(ID); err != nil {
		return
	}

	// Binding to model
	err = c.Bind(&permission)
	if err != nil {
		err = errors.Wrap(err, status.PermissionParamBindingErrorCode)
		return
	}

	// set UpdateBy and UpdateTime
	cuser, _ := middleware.GetCurrentUser(ctx, c)
	permission.UpdateBy = cuser.Account
	permission.UpdateTime = date.DateUnix()

	// Do edit
	err = permission.Edit(model.GetDB())
}
