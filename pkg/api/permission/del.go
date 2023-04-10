/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-31 09:46:27
 * @LastEditTime: 2023-02-03 09:58:27
 * @Description: Do not edit
 */
package permission

import (
	"context"
	"strconv"

	"github.com/chenke1115/hertz-common/pkg/errors"
	"github.com/chenke1115/hertz-common/pkg/response"
	"github.com/chenke1115/hertz-permission/internal/constant/status"
	"github.com/chenke1115/hertz-permission/pkg/model"

	"github.com/cloudwego/hertz/pkg/app"
)

// DelHandler goDoc
// @Summary     删除权限
// @Description This is a api to del permission
// @Tags        Permission【权限】
// @Accept      json
// @Produce     json
// @Security    authorization
// @Param       id  path     int true "权限ID" example(1)
// @Success     200 {object} response.BaseResponse{data=interface{}}
// @Failure     400 {object} response.BaseResponse{data=interface{}}
// @Router      /api/permission/{id}/del [delete]
func DelHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err        error
		ID         int
		permission model.Permission
	)

	// Response
	defer func() {
		response.HandleResponse(c, err, nil)
	}()

	// ID
	if ID, err = strconv.Atoi(c.Param("id")); err != nil {
		err = errors.Wrap(err, status.PermissionIdMissCode)
		return
	}

	// Find
	if permission, err = model.GetPermissionByID(ID); err != nil {
		return
	}

	err = permission.Del(model.GetDB())
}
