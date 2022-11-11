/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-31 09:46:27
 * @LastEditTime: 2022-11-10 15:25:56
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

// DelHandler goDoc
// @Summary     删除角色
// @Description This is a api to del role
// @Tags        RoleDel
// @Accept      json
// @Produce     json
// @Param       id  query    int true "角色ID" example(1)
// @Success     200 {object} response.BaseResponse{data=interface{}}
// @Failure     400 {object} response.BaseResponse{data=interface{}}
// @Router      /api/role/{id}/del [delete]
func DelHandler(ctx context.Context, c *app.RequestContext) {
	var (
		err  error
		ID   int
		role model.Role
	)

	// Response
	defer func() {
		response.HandleResponse(c, err, nil)
	}()

	// ID
	if ID, err = strconv.Atoi(c.Param("id")); err != nil {
		err = errors.Wrap(err, status.RoleIdMissCode)
		return
	}

	// Find
	if role, err = model.GetRoleByID(ID); err != nil {
		return
	}

	err = role.Del(model.GetDB())
}
