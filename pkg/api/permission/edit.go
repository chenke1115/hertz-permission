/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-31 09:45:52
 * @LastEditTime: 2022-11-01 17:08:33
 * @Description: Do not edit
 */
package permission

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
	PID        int    `json:"pid,required" form:"pid,required"`     //lint:ignore SA5008 ignoreCheck
	Name       string `json:"name,required" form:"name,required"`   //lint:ignore SA5008 ignoreCheck
	Alias      string `json:"alias,required" form:"alias,required"` //lint:ignore SA5008 ignoreCheck
	Type       string `json:"type,required" form:"type,required"`   //lint:ignore SA5008 ignoreCheck
	Key        string `json:"key" form:"key"`
	Components string `json:"components" form:"components"`
	Sort       int    `json:"sort" form:"sort"`
	Icon       string `json:"icon" form:"icon"`
	Visible    int    `json:"visible" form:"visible"`
	Status     int    `json:"status" form:"status"`
	Remark     string `json:"remark" form:"remark"`
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

	// Check route
	if !model.IsValidRoute(req.Key) {
		err = errors.New(status.PermissionKeyErrorCode)
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

	// Do edit
	err = permission.Edit(model.GetDB())
}
