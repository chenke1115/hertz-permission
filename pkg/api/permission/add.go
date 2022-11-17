/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-31 09:44:07
 * @LastEditTime: 2022-11-14 14:08:25
 * @Description: Do not edit
 */
package permission

import (
	"context"

	"github.com/chenke1115/go-common/functions/date"
	"github.com/chenke1115/hertz-common/pkg/errors"
	_ "github.com/chenke1115/hertz-common/pkg/errors/validate"
	"github.com/chenke1115/hertz-common/pkg/response"
	"github.com/chenke1115/hertz-permission/internal/constant/status"
	"github.com/chenke1115/hertz-permission/pkg/middleware"
	"github.com/chenke1115/hertz-permission/pkg/model"

	"github.com/cloudwego/hertz/pkg/app"
)

type ReqAddData struct {
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

// AddHandler goDoc
// @Summary     添加权限
// @Description This is a api to add permission
// @Tags        PermissionAdd
// @Accept      json
// @Produce     json
// @Param       pid        body     int    true  "父级ID" maximum(10) default(0)
// @Param       name       body     string true  "权限名称" maxlength(32) example("permission.add")
// @Param       alias      body     string true  "别名"   maxlength(32) example("添加权限")
// @Param       type       body     string true  "权限类型[D:目录;M:菜单;B:按钮]"   Enums("D", "M", "B")
// @Param       key        body     string false "权限全局标识[即后端路由，类型为目录可空]" maxlength(32)
// @Param       components body     string false "前端页面路径[类型为按钮可空]"       maxlength(32)
// @Param       sort       body     int    false "排序[从小到大]"              default(0)
// @Param       icon       body     string false "图标"                    maxlength(255)
// @Param       visible    body     int    false "菜单状态[1:显示;0:隐藏]"       Enums(1, 0)
// @Param       status     body     int    false "菜单状态[1:显示;0:隐藏]"       Enums(1, 0)
// @Param       remark     body     string false "备注"                    maxlength(255)
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

	// Binding to permission model
	permission := &model.Permission{}
	err = c.Bind(&permission)
	if err != nil {
		err = errors.Wrapf(err, status.PermissionParamBindingErrorCode)
		return
	}

	// set UpdateBy and UpdateTime
	cuser, _ := middleware.GetCurrentUser(ctx, c)
	permission.UpdateBy = cuser.Account
	permission.UpdateTime = date.DateUnix()

	err = permission.Create(model.GetDB())
}
