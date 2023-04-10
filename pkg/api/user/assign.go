/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-01 17:38:22
 * @LastEditTime: 2023-02-03 11:31:21
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

type ReqAssignData struct {
	RoleIDs []int `json:"role_ids,required" form:"role_ids,required"` // 权限ID数组【必填】
}

// AssignHandler goDoc
// @Summary     角色分配
// @Description This is a api to assign role for user
// @Tags        User【用户】
// @Accept      x-www-form-urlencoded
// @Produce     json
// @Security    authorization
// @Param       id   path     int           true "账户ID" example(1)
// @Param       data formData ReqAssignData true "请求数据"
// @Success     200  {object} response.BaseResponse{data=interface{}}
// @Failure     400  {object} response.BaseResponse{data=interface{}}
// @Router      /api/user/{id}/assign [post]
func AssignHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err      error
		ID       int
		req      ReqAssignData
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

	// User exist
	if userInfo, err = model.GetUserInfoByID(ID); err != nil {
		return
	}

	err = userInfo.AssignRole(model.GetDB(), req.RoleIDs)
}
