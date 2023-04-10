/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-09 15:23:25
 * @LastEditTime: 2023-02-03 11:31:24
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

type ReqEditData struct {
	Name       string `json:"name,required" form:"name,required" vd:"len($)<32"`      // 用户名 example("长歌")
	Account    string `json:"account,required" form:"account,required" vd:"email($)"` // 登陆账户 example("changge@ismart.com")
	CustomerID int    `json:"customer_id" form:"customer_id" default:"0"`             // 关联客户账户ID
}

// EditHandler goDoc
// @Summary     编辑用户
// @Description This is a api to edit user
// @Tags        User【用户】
// @Accept      x-www-form-urlencoded
// @Produce     json
// @Security    authorization
// @Param       id   path     int         true "用户ID" example(1)
// @Param       data formData ReqEditData true "请求数据"
// @Success     200  {object} response.BaseResponse{data=interface{}}
// @Failure     400  {object} response.BaseResponse{data=interface{}}
// @Router      /api/user/{:id}/edit [put]
func EditHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err      error
		ID       int
		req      ReqEditData
		userInfo model.UserInfo
	)

	// Response
	defer func() {
		response.HandleResponse(c, err, nil)
	}()

	// ID
	if ID, err = strconv.Atoi(c.Param("id")); err != nil {
		err = errors.Wrapf(err, status.UserMissIDCode)
		return
	}

	// BindAndValidate
	err = c.BindAndValidate(&req)
	if err != nil {
		err = errors.WrapCode(err, status.UserErrorParamCode)
		return
	}

	// Find
	userInfo, err = model.GetUserInfoByID(ID)
	if err != nil {
		return
	}

	// Binding to model
	err = c.Bind(&userInfo)
	if err != nil {
		err = errors.Wrapf(err, status.UserParamBindErrCode)
		return
	}

	err = userInfo.Edit(model.GetDB())
}
