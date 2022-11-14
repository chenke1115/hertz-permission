/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-11 10:55:15
 * @LastEditTime: 2022-11-14 13:50:10
 * @Description: Do not edit
 */
package auth

import (
	"context"

	"github.com/chenke1115/hertz-permission/internal/constant/status"
	"github.com/chenke1115/hertz-permission/internal/pkg/errors"
	"github.com/chenke1115/hertz-permission/internal/pkg/response"
	_ "github.com/chenke1115/hertz-permission/internal/pkg/validate"
	"github.com/chenke1115/hertz-permission/pkg/middleware"

	"github.com/cloudwego/hertz/pkg/app"
)

type ReqLoginData struct {
	Username string `json:"username,required" form:"username,required"`                       //lint:ignore SA5008 ignoreCheck
	Password string `json:"password,required" form:"password,required" vd:"checkPassword($)"` //lint:ignore SA5008 ignoreCheck
}

// LoginHandler godoc
// @Summary     登陆
// @Group       auth
// @Description This is an api to login
// @Tags        Auth.Login
// @Accept      json
// @Produce     json
// @Param       username body     string true "用户名[超管初始账户：admin]"    maxlength(32)
// @Param       password body     string true "密码[超管初始密码：Admin123!]" maxlength(32)
// @Success     200      {object} response.BaseResponse{data=interface{}}
// @Failure     400      {object} response.BaseResponse{data=interface{}}
// @Router      /api/auth/login [post]
func LoginHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err error
		req ReqLoginData
	)

	// Response
	defer func() {
		if err != nil {
			response.HandleResponse(c, err, nil)
			return
		}

		// Jwt
		middleware.JwtLoginHandler(ctx, c)
	}()

	// BindAndValidate
	err = c.BindAndValidate(&req)
	if err != nil {
		err = errors.WrapCode(err, status.UserErrorParamCode)
		c.Abort()
		return
	}
}
