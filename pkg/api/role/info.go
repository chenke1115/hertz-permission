/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-01 11:46:59
 * @LastEditTime: 2022-11-01 15:10:36
 * @Description: Do not edit
 */
package role

import (
	"context"
	"strconv"

	"github.com/chenke1115/hertz-permission/internal/constant/status"
	"github.com/chenke1115/hertz-permission/internal/pkg/errors"
	"github.com/chenke1115/hertz-permission/internal/pkg/response"
	"github.com/chenke1115/hertz-permission/pkg/model"

	"github.com/cloudwego/hertz/pkg/app"
)

type RespInfo struct {
	Role model.Role `json:"role"`
}

// InfoHandler goDoc
// @Summary     角色详情
// @Description This is a api of role info
// @Tags        RoleInfo
// @Accept      json
// @Produce     json
// @Param       id  query    int true "角色ID" example(1)
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
}
