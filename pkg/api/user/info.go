/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-11 16:11:52
 * @LastEditTime: 2023-01-09 14:55:39
 * @Description: Do not edit
 */
package user

import (
	"context"

	"github.com/chenke1115/hertz-common/pkg/response"
	"github.com/chenke1115/hertz-permission/pkg/middleware"
	"github.com/chenke1115/hertz-permission/pkg/model"

	"github.com/cloudwego/hertz/pkg/app"
)

// CurrentUserHandler godoc
// @Summary     当前用户信息
// @Description This is an api to get current user info
// @Tags        CurrentUser
// @Accept      json
// @Produce     json
// @Security    authorization
// @Success     200 {object} response.BaseResponse{data=model.CurrentUser{}}
// @Failure     400 {object} response.BaseResponse{data=interface{}}
// @Router      /api/user/current [get]
func CurrentUserHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err         error
		currentUser *model.CurrentUser
	)

	// Response
	defer func() {
		if err != nil {
			currentUser = &model.CurrentUser{}
		}

		response.HandleResponse(c, err, &currentUser)
	}()

	// Clean permission cache
	err = middleware.CleanCurrentUserCache(ctx, c)
	if err != nil {
		return
	}

	currentUser, err = middleware.GetCurrentUser(ctx, c)
}
