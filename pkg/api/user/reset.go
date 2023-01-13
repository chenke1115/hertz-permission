/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-09 16:43:39
 * @LastEditTime: 2023-01-09 14:55:47
 * @Description: Do not edit
 */
package user

import (
	"context"
	"strconv"

	"github.com/chenke1115/go-common/configs"
	"github.com/chenke1115/go-common/functions/hash"
	"github.com/chenke1115/hertz-common/pkg/errors"
	"github.com/chenke1115/hertz-common/pkg/response"
	_ "github.com/chenke1115/hertz-common/pkg/validate"
	"github.com/chenke1115/hertz-permission/internal/constant/status"
	"github.com/chenke1115/hertz-permission/pkg/model"

	"github.com/cloudwego/hertz/pkg/app"
)

type ReqResetData struct {
	OldPassword     string `json:"old_password,required" form:"old_password,required"`                                              //lint:ignore SA5008 ignoreCheck
	Password        string `json:"password,required" form:"password,required" vd:"checkPassword($)"`                                //lint:ignore SA5008 ignoreCheck
	ConfirmPassword string `json:"confirm_password,required" form:"confirm_password,required" vd:"confirmPassword($, (Password)$)"` //lint:ignore SA5008 ignoreCheck
}

// ResetHandler goDoc
// @Summary     用户密码重置
// @Description This is a api of reset user password
// @Tags        UserResetPassword
// @Accept      json
// @Produce     json
// @Security    authorization
// @Param       user_id          query    int    true "用户列表返回的UserID" example(1)
// @Param       old_password     body     string true "旧密码" example("Admin123!")
// @Param       password         body     string true "新密码" example("Admin123#")
// @Param       confirm_password body     string true "确认新密码" example("Admin123#")
// @Success     200              {object} response.BaseResponse{data=interface{}}
// @Failure     400              {object} response.BaseResponse{data=interface{}}
// @Router      /api/user/{user_id}/reset [put]
func ResetHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err  error
		ID   int
		req  ReqResetData
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

	// Check old password
	userConf := configs.GetConf().App.User
	if !hash.CheckHashedPassword(req.OldPassword, userConf.Password.Salt, user.Password) {
		err = errors.New(status.UserIncorrectOldPasswordCode)
		return
	}

	user.Password = hash.GetHashedPassword(req.Password, userConf.Password.Salt)
	err = user.Edit(model.GetDB())
}
