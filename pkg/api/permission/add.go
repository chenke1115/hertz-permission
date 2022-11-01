/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-31 09:44:07
 * @LastEditTime: 2022-11-01 10:24:52
 * @Description: Do not edit
 */
package permission

import (
	"context"

	"github.com/chenke1115/ismart-permission/internal/constant/status"
	"github.com/chenke1115/ismart-permission/internal/pkg/errors"
	"github.com/chenke1115/ismart-permission/internal/pkg/response"
	"github.com/chenke1115/ismart-permission/pkg/model"

	"github.com/cloudwego/hertz/pkg/app"
)

type ReqAddData struct {
	PID        int    `json:"pid,required" form:"pid,required"`     //lint:ignore SA5008 ignoreCheck
	Name       string `json:"name,required" form:"name,required"`   //lint:ignore SA5008 ignoreCheck
	Alias      string `json:"alias,required" form:"alias,required"` //lint:ignore SA5008 ignoreCheck
	Key        string `json:"key" form:"key"`
	Components string `json:"components" form:"components"`
	Sort       int    `json:"sort" form:"sort"`
	Icon       string `json:"icon" form:"icon"`
	Visible    int    `json:"visible" form:"visible"`
	Status     int    `json:"status" form:"status"`
	Remark     string `json:"remark" form:"remark"`
}

// AddHandler goDoc
// @Summary     添加权限
// @Description This is a api to add permission
// @Tags        PermissionAdd
// @Accept      json
// @Produce     json
// @Param       pid        body     int    true  "父级ID" maximum(10) example(1)
// @Param       name       body     string true  "权限名称" maxlength(32) example("permission.add")
// @Param       alias      body     string true  "别名"   maxlength(32) example("添加权限")
// @Param       key        body     string false "权限全局标识[即路由，类型为目录可空]" maxlength(32)
// @Param       components body     string false "前端页面路径[类型为按钮可空]"     maxlength(32)
// @Param       sort       body     int    false "排序[从小到大]"            default(0)
// @Param       icon       body     string false "图标"                  maxlength(255)
// @Param       visible    body     int    false "菜单状态[1:显示;0:隐藏]"     Enums(1, 0)
// @Param       status     body     int    false "菜单状态[1:显示;0:隐藏]"     Enums(1, 0)
// @Param       remark     body     string false "备注"                  maxlength(255)
// @Success     200        {object} response.BaseResponse{data=interface{}}
// @Failure     400        {object} response.BaseResponse{data=interface{}}
// @Router      /api/permission/add [post]
func AddHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err error
		req ReqAddData
	)

	// Response
	defer func() {
		response.HandleResponse(c, err, nil)
	}()

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

	// Binding to permission model
	permission := &model.Permission{}
	err = c.Bind(&permission)
	if err != nil {
		err = errors.Wrapf(err, status.PermissionParamBindingErrorCode)
		return
	}

	if err = permission.Create(model.GetDB()); err != nil {
		return
	}
}
