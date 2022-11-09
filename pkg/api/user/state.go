/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-09 16:43:09
 * @LastEditTime: 2022-11-09 17:24:16
 * @Description: Do not edit
 */
package user

import (
	"context"

	"github.com/chenke1115/hertz-permission/internal/pkg/response"

	"github.com/cloudwego/hertz/pkg/app"
)

type ReqStateData struct {
	Status int `json:"status,requried" form:"status,requried"` //lint:ignore SA5008 ignoreCheck
}

// StateHandler goDoc
// @Summary     用户状态变更
// @Description This is a api of edit user status
// @Tags        UserStateEdit
// @Accept      json
// @Produce     json
// @Param       user_id query    int true "用户列表返回的UserID"   example(1)
// @Param       status  body     int true "用户状态[1:启用 0:失效]" Enums(1, 0)
// @Success     200     {object} response.BaseResponse{data=interface{}}
// @Failure     400     {object} response.BaseResponse{data=interface{}}
// @Router      /api/user/{user_id}/state [put]
func StateHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err error
	)

	// Response
	defer func() {
		response.HandleResponse(c, err, nil)
	}()
}
