/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-11 11:55:43
 * @LastEditTime: 2023-04-10 14:26:09
 * @Description: Do not edit
 */
package auth

import (
	"context"

	"github.com/chenke1115/hertz-permission/pkg/middleware"
	"github.com/cloudwego/hertz/pkg/app"
)

// LogoutHandler godoc
// @Summary     退出登陆
// @Description This is an api to logout
// @Tags        Auth【授权】
// @Accept      json
// @Produce     json
// @Security    authorization
// @Success     200 {object} response.BaseResponse{data=interface{}}
// @Failure     400 {object} response.BaseResponse{data=interface{}}
// @Router      /api/auth/logout [post]
func LogoutHandler(ctx context.Context, c *app.RequestContext) {
	// Response
	defer func() {
		middleware.Jwt().LogoutHandler(ctx, c)
	}()
}
