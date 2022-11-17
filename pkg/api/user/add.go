/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-09 10:14:57
 * @LastEditTime: 2022-11-09 14:54:02
 * @Description: Do not edit
 */
package user

import (
	"context"

	"github.com/chenke1115/hertz-common/pkg/errors"
	"github.com/chenke1115/hertz-common/pkg/response"
	"github.com/chenke1115/hertz-permission/internal/constant/status"
	"github.com/chenke1115/hertz-permission/pkg/model"

	"github.com/cloudwego/hertz/pkg/app"
)

type ReqAddData struct {
	Name       string `json:"name,required" form:"name,required" vd:"len($)<32"`      //lint:ignore SA5008 ignoreCheck
	Account    string `json:"account,required" form:"account,required" vd:"email($)"` //lint:ignore SA5008 ignoreCheck
	CustomerID int    `json:"customer_id" form:"customer_id" default:"0"`
}

// AddHandler goDoc
// @Summary     添加用户
// @Description This is a api to add user
// @Tags        UserAdd
// @Accept      json
// @Produce     json
// @Param       name        body     string true  "用户名"             example("长歌")
// @Param       account     body     string true  "登陆账户"            example("changge@ismart.com")
// @Param       customer_id body     int    false "关联客户账户ID" default(0)
// @Success     200         {object} response.BaseResponse{data=interface{}}
// @Failure     400         {object} response.BaseResponse{data=interface{}}
// @Router      /api/user/add [post]
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
		err = errors.WrapCode(err, status.RoleParamErrorCode)
		return
	}

	// Binding to user model
	userInfo := &model.UserInfo{}
	err = c.Bind(&userInfo)
	if err != nil {
		err = errors.Wrapf(err, status.UserParamBindErrCode)
		return
	}

	// Create
	err = userInfo.Create(model.GetDB())
}
