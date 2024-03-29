/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-01 10:33:30
 * @LastEditTime: 2023-04-10 14:29:21
 * @Description: Do not edit
 */
package permission

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
	model.Permission
	UpdateTime string `json:"update_time"`
}

// InfoHandler goDoc
// @Summary     权限详情
// @Description This is a api of permission info
// @Tags        Permission【权限】
// @Accept      json
// @Produce     json
// @Security    authorization
// @Param       id  path     int true "权限ID" example(1)
// @Success     200 {object} response.BaseResponse{data=permission.RespInfo{}}
// @Failure     400 {object} response.BaseResponse{data=interface{}}
// @Router      /api/permission/{id}/info [get]
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
		err = errors.Wrap(err, status.PermissionIdMissCode)
		return
	}

	resp.Permission, err = model.GetPermissionByID(ID)
	if err != nil {
		return
	}

	resp.UpdateTime = date.DateFormat(resp.Permission.UpdateTime)
}
