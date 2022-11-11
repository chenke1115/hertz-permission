/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-11 11:56:14
 * @LastEditTime: 2022-11-11 14:07:48
 * @Description: Do not edit
 */
package auth

import (
	"context"

	"github.com/chenke1115/hertz-permission/pkg/middleware"
	"github.com/cloudwego/hertz/pkg/app"
)

// RefreshHandler godoc
// @Summary     刷新TOKEN
// @Description This is an api to refresh token
// @Tags        Auth.RefreshToken
// @Accept      json
// @Produce     json
// @Success     200 {object} response.BaseResponse{data=interface{}}
// @Failure     400 {object} response.BaseResponse{data=interface{}}
// @Router      /api/auth/refresh_token [post]
func RefreshHandler(ctx context.Context, c *app.RequestContext) {
	// Response
	defer func() {
		middleware.JwtRefreshHandler(ctx, c)
	}()
}
