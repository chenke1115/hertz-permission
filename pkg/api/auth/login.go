/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-11 10:55:15
 * @LastEditTime: 2023-07-27 18:26:45
 * @Description: Do not edit
 */
package auth

import (
	"context"

	"github.com/chenke1115/hertz-common/pkg/errors"
	"github.com/chenke1115/hertz-common/pkg/response"
	_ "github.com/chenke1115/hertz-common/pkg/validate"
	"github.com/chenke1115/hertz-permission/internal/constant/status"
	"github.com/chenke1115/hertz-permission/pkg/middleware"
	"github.com/cloudwego/hertz/pkg/app"
)

type ReqLoginData struct {
	Username string `json:"username,required" form:"username,required"`                       // 用户名[超管初始账户：admin]【必填】 maxlength(32)
	Password string `json:"password,required" form:"password,required" vd:"checkPassword($)"` // 密码[超管初始密码：Admin123!]【必填】 maxlength(32)
}

// LoginHandler godoc
// @Summary     登陆
// @Group       auth
// @Description This is an api to login
// @Tags        Auth【授权】
// @Accept      x-www-form-urlencoded
// @Produce     json
// @Param       data formData ReqLoginData true "请求数据"
// @Success     200  {object} response.BaseResponse{data=interface{}}
// @Failure     400  {object} response.BaseResponse{data=interface{}}
// @Router      /api/auth/login [post]
func LoginHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err error
		req ReqLoginData
	)

	// Response
	defer func() {
		if err != nil {
			response.HandleResponseWithStatus(c, errors.BadRequest, err, nil)
			return
		}
	}()

	// BindAndValidate
	err = c.BindAndValidate(&req)
	if err != nil {
		err = errors.Newf(status.UserLoginErrCode)
		return
	}

	// Jwt
	middleware.Jwt().LoginHandler(ctx, c)
}
