/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-09 16:43:09
 * @LastEditTime: 2023-02-03 11:31:32
 * @Description: Do not edit
 */
package user

import (
	"context"
	"strconv"

	"github.com/chenke1115/hertz-common/pkg/errors"
	"github.com/chenke1115/hertz-common/pkg/response"
	"github.com/chenke1115/hertz-permission/internal/constant/status"
	"github.com/chenke1115/hertz-permission/pkg/model"

	"github.com/cloudwego/hertz/pkg/app"
)

type ReqStateData struct {
	Status int `json:"status,requried" form:"status,requried"` // 用户状态[1:启用 0:失效]【必填】
}

// StateHandler goDoc
// @Summary     用户状态变更
// @Description This is a api of edit user status
// @Tags        User【用户】
// @Accept      x-www-form-urlencoded
// @Produce     json
// @Security    authorization
// @Param       user_id path     int          true "用户列表返回的UserID" example(1)
// @Param       data    formData ReqStateData true "请求数据"
// @Success     200     {object} response.BaseResponse{data=interface{}}
// @Failure     400     {object} response.BaseResponse{data=interface{}}
// @Router      /api/user/{user_id}/state [put]
func StateHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err  error
		ID   int
		req  ReqStateData
		user model.User
	)

	// Response
	defer func() {
		response.HandleResponse(c, err, nil)
	}()

	// ID
	if ID, err = strconv.Atoi(c.Param("id")); err != nil {
		err = errors.WrapCode(err, status.UserMissIDCode)
		return
	}

	// BindAndValidate
	err = c.BindAndValidate(&req)
	if err != nil {
		err = errors.WrapCode(err, status.UserErrorParamCode)
		return
	}

	// Find
	if user, err = model.GetUserByID(ID); err != nil {
		return
	}

	user.Status = req.Status
	err = user.Edit(model.GetDB())
}
