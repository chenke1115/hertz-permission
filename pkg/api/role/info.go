/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-01 11:46:59
 * @LastEditTime: 2023-02-03 10:09:39
 * @Description: Do not edit
 */
package role

import (
	"context"
	"strconv"

	"github.com/chenke1115/go-common/functions/date"
	"github.com/chenke1115/hertz-common/pkg/errors"
	"github.com/chenke1115/hertz-common/pkg/response"
	"github.com/chenke1115/hertz-permission/internal/constant/status"
	"github.com/chenke1115/hertz-permission/pkg/model"

	"github.com/cloudwego/hertz/pkg/app"
)

type RespInfo struct {
	model.Role
	UpdateTime string `json:"update_time"`
}

// InfoHandler goDoc
// @Summary     角色详情
// @Description This is a api of role info
// @Tags        Role【角色】
// @Accept      json
// @Produce     json
// @Security    authorization
// @Param       id  path     int true "角色ID" example(1)
// @Success     200 {object} response.BaseResponse{data=role.RespInfo{}}
// @Failure     400 {object} response.BaseResponse{data=interface{}}
// @Router      /api/role/{id}/info [get]
func InfoHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err  error
		ID   int
		resp RespInfo
	)

	// Response
	defer func() {
		if err != nil {
			resp = RespInfo{}
		}
		response.HandleResponse(c, err, &resp)
	}()

	// ID
	if ID, err = strconv.Atoi(c.Param("id")); err != nil {
		err = errors.Wrap(err, status.RoleIdMissCode)
		return
	}

	resp.Role, err = model.GetRoleByID(ID)
	if err != nil {
		return
	}

	resp.UpdateTime = date.DateFormat(resp.Role.UpdateTime)
}
