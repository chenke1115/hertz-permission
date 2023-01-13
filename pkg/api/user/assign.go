/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-01 17:38:22
 * @LastEditTime: 2023-01-09 14:55:31
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
	RoleIDs []int `json:"role_ids,required" form:"role_ids,required"` //lint:ignore SA5008 ignoreCheck
}

// AssignHandler goDoc
// @Summary     角色分配
// @Description This is a api to assign role for user
// @Tags        UserRoleAssign
// @Accept      json
// @Produce     json
// @Security    authorization
// @Param       id       query    int   true "账户ID"   example(1)
// @Param       role_ids body     array true "权限ID数组" example([1, 3])
// @Success     200      {object} response.BaseResponse{data=interface{}}
// @Failure     400      {object} response.BaseResponse{data=interface{}}
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
