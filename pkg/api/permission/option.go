/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-01 17:53:08
 * @LastEditTime: 2023-01-09 14:54:44
 * @Description: Do not edit
 */
package permission

import (
	"context"

	"github.com/chenke1115/go-common/functions/array"
	"github.com/chenke1115/hertz-common/global"
	"github.com/chenke1115/hertz-common/pkg/response"
	"github.com/chenke1115/hertz-permission/pkg/model"

	"github.com/cloudwego/hertz/pkg/app"
)

// OptionHandler goDoc
// @Summary     权限下拉选项
// @Description This is a api of permission option
// @Tags        PermissionOption
// @Accept      json
// @Produce     json
// @Success     200 {object} response.BaseResponse{data=[]model.PermissionOption{}}
// @Failure     400 {object} response.BaseResponse{data=interface{}}
// @Router      /api/permission/option [get]
func OptionHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err  error
		resp []model.PermissionOption
	)

	// Response
	defer func() {
		if err != nil {
			resp = []model.PermissionOption{}
		}

		response.HandleResponse(c, err, &resp)
	}()

	option := model.Permission{}
	resp, err = option.Option()
}

// MenuOptionHandler goDoc
// @Summary     上级菜单下拉选项
// @Description This is a api of menu option. [选项键值就是上级菜单ID]
// @Tags        MenuOption
// @Accept      json
// @Produce     json
// @Success     200 {object} response.BaseResponse{data=map[int]string{}}
// @Failure     400 {object} response.BaseResponse{data=interface{}}
// @Router      /api/permission/menu [get]
func MenuOptionHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err  error
		resp map[int]string
	)

	// Response
	defer func() {
		if err != nil {
			resp = map[int]string{}
		}

		response.HandleResponse(c, err, &resp)
	}()

	option := model.Permission{}
	resp, err = option.MenuOption()
}

// RouteHandler goDoc
// @Summary     路由下拉选项[未添加为权限的路由]
// @Description This is a api of route option
// @Tags        PermissionRouteOption
// @Accept      json
// @Produce     json
// @Security    authorization
// @Success     200 {object} response.BaseResponse{data=map[int]string{}}
// @Failure     400 {object} response.BaseResponse{data=interface{}}
// @Router      /api/permission/route [get]
func RouteHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err  error
		resp map[int]string
	)

	defer func() {
		if err != nil {
			resp = map[int]string{}
		}

		response.HandleResponse(c, err, resp)
	}()

	yetKeys, _ := model.GetPermissionKeys()

	i := 0
	resp = map[int]string{}
	for _, v := range global.RouteInfo {
		if v.Path != "" && !array.In(v.Path, yetKeys) {
			resp[i] = v.Path
		}
		i++
	}
}
