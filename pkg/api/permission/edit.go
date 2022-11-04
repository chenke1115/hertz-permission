/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-31 09:45:52
 * @LastEditTime: 2022-11-03 16:22:52
 * @Description: Do not edit
 */
package permission

import (
	"context"
	"strconv"

	"github.com/chenke1115/hertz-permission/internal/constant/status"
	"github.com/chenke1115/hertz-permission/internal/pkg/date"
	"github.com/chenke1115/hertz-permission/internal/pkg/errors"
	_ "github.com/chenke1115/hertz-permission/internal/pkg/errors/validate"
	"github.com/chenke1115/hertz-permission/internal/pkg/response"
	"github.com/chenke1115/hertz-permission/pkg/model"

	"github.com/cloudwego/hertz/pkg/app"
)

type ReqEditData struct {
	Name       string `json:"name,required" form:"name,required" vd:"len($)<64"`             //lint:ignore SA5008 ignoreCheck
	Alias      string `json:"alias,required" form:"alias,required" vd:"len($)<64"`           //lint:ignore SA5008 ignoreCheck
	Type       string `json:"type,required" form:"type,required"`                            //lint:ignore SA5008 ignoreCheck
	Key        string `json:"key,required" form:"key,required"`                              //lint:ignore SA5008 ignoreCheck
	Components string `json:"components,required" form:"components,required" vd:"len($)<64"` //lint:ignore SA5008 ignoreCheck
	PID        int    `json:"pid" form:"pid" default:"0"`
	Sort       int    `json:"sort" form:"sort" default:"0"`
	Icon       string `json:"icon" form:"icon"`
	Visible    int    `json:"visible" form:"visible" default:"1"`
	Status     int    `json:"status" form:"status" default:"1"`
	Remark     string `json:"remark" form:"remark" vd:"len($)<256"`
}

// EditHandler goDoc
// @Summary     修改权限
// @Description This is a api to edit permission
// @Tags        PermissionEdit
// @Accept      json
// @Produce     json
// @Param       id         query    int    true  "权限ID" example(1)
// @Param       pid        body     int    true  "父级ID" example(1)
// @Param       name       body     string true  "权限名称" maxlength(32) example("permission.add")
// @Param       alias      body     string true  "别名"   maxlength(32) example("添加权限")
// @Param       type       body     string true  "权限类型[D:目录;M:菜单;B:按钮]"   Enums("D", "M", "B")
// @Param       key        body     string false "权限全局标识[即路由，类型为目录可空]" maxlength(32)
// @Param       components body     string false "前端页面路径[类型为按钮可空]"     maxlength(32)
// @Param       sort       body     int    false "排序[从小到大]"            default(0)
// @Param       icon       body     string false "图标"                  maxlength(255)
// @Param       visible    body     int    false "菜单状态[1:显示;0:隐藏]"     Enums(1, 0)
// @Param       status     body     int    false "菜单状态[1:显示;0:隐藏]"     Enums(1, 0)
// @Param       remark     body     string false "备注"                  maxlength(255)
// @Success     200        {object} response.BaseResponse{data=interface{}}
// @Failure     400        {object} response.BaseResponse{data=interface{}}
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
	permission.UpdateBy = "admin" // TODO
	permission.UpdateTime = date.DateUnix()

	// Do edit
	err = permission.Edit(model.GetDB())
}
